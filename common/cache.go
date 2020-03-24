// Copyright (c) 2019-2020 The Zcash developers
// Distributed under the MIT software license, see the accompanying
// file COPYING or https://www.opensource.org/licenses/mit-license.php .
package common

import (
	"bytes"
	"encoding/binary"
	"hash/fnv"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"github.com/golang/protobuf/proto"
	"github.com/zcash/lightwalletd/walletrpc"
)

type blockCacheEntry struct {
	data []byte
}

// BlockCache contains a consecutive set of recent compact blocks in marshalled form.
type BlockCache struct {
	lengthsName, blocksName string // pathnames
	lengthsFile, blocksFile *os.File
	starts                  []int64 // Starting offset of each block within blocksFile
	firstBlock              int     // height of the first block in the cache (usually Sapling activation)
	nextBlock               int     // height of the first block not in the cache
	latestHash              []byte  // hash of the most recent (highest height) block, for detecting reorgs.
	mutex                   sync.RWMutex
}

func (c *BlockCache) GetNextHeight() int {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.nextBlock
}

func (c *BlockCache) GetLatestHash() []byte {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.latestHash
}

//  HashMismatch indicates if the given prev-hash doesn't match the most recent block's hash
// so reorgs can be detected.
func (c *BlockCache) HashMismatch(prevhash []byte) bool {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.latestHash != nil && !bytes.Equal(c.latestHash, prevhash)
}

// Start over, so that all blocks get re-downloaded (file corruption)
func (c *BlockCache) truncZero() {
	if err := c.lengthsFile.Truncate(0); err != nil {
		Log.Fatal("truncate lengths file failed: ", err)
	}
	if err := c.blocksFile.Truncate(0); err != nil {
		Log.Fatal("truncate blocks file failed: ", err)
	}
	c.Sync()
	c.starts = nil
	c.starts = append(c.starts, 0)
	c.nextBlock = c.firstBlock
	c.latestHash = nil
}

// not including the checksum
func (c *BlockCache) blockLength(height int) int {
	index := height - c.firstBlock
	return int(c.starts[index+1] - c.starts[index] - 8)
}

// Calculate the 8-byte checksum that precedes each block in the blocks file.
func checksum(height int, b []byte) []byte {
	h := make([]byte, 8)
	binary.LittleEndian.PutUint64(h, uint64(height))
	cs := fnv.New64a()
	cs.Write(h)
	cs.Write(b)
	return cs.Sum(nil)
}

func (c *BlockCache) readBlockFromDisk(height int) *walletrpc.CompactBlock {
	blockLen := c.blockLength(height)
	b := make([]byte, blockLen+8)
	offset := c.starts[height-c.firstBlock]
	n, err := c.blocksFile.ReadAt(b, offset)
	if err != nil || n != len(b) {
		Log.Warning("blocks read offset: ", offset, " failed: ", n, err)
		return nil
	}
	diskcs := b[:8]
	b = b[8 : blockLen+8]
	if !bytes.Equal(checksum(height, b), diskcs) {
		Log.Warning("bad block checksum at height: ", height, " offset: ", offset)
		return nil
	}
	block := &walletrpc.CompactBlock{}
	err = proto.Unmarshal(b, block)
	if err != nil {
		// Could be file corruption.
		Log.Warning("blocks unmarshal at offset: ", offset, " failed: ", err)
		return nil
	}
	// TODO COINBASE-HEIGHT: restore this check after coinbase height is fixed
	if false && int(block.Height) != height {
		// Could be file corruption.
		Log.Warning("block unexpected height at height ", height, " offset: ", offset)
		return nil
	}
	return block
}

func (c *BlockCache) readBlock(height int) *walletrpc.CompactBlock {
	b := c.readBlockFromDisk(height)
	if b == nil {
		// Some kind of disk corruption occured, start over
		c.truncZero()
	}
	return b
}

func (c *BlockCache) setLatestHash() {
	c.latestHash = nil
	// There is at least one block; get the last block's hash
	if c.nextBlock > c.firstBlock {
		// At least one block remains; get the last block's hash
		block := c.readBlock(c.nextBlock - 1)
		if block == nil {
			return
		}
		c.latestHash = make([]byte, len(block.Hash))
		copy(c.latestHash, block.Hash)
	}
}

// NewBlockCache returns an instance of a block cache object.
func NewBlockCache(dbPath string, chainName string, startHeight int, redownload bool, backup int) *BlockCache {
	c := &BlockCache{}
	c.firstBlock = startHeight
	c.nextBlock = startHeight
	c.lengthsName, c.blocksName = dbFileNames(dbPath, chainName)
	var err error
	if err := os.MkdirAll(filepath.Join(dbPath, chainName), 0755); err != nil {
		Log.Fatal("mkdir ", dbPath, " failed: ", err)
	}
	c.blocksFile, err = os.OpenFile(c.blocksName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		Log.Fatal("open ", c.blocksName, " failed: ", err)
	}
	c.lengthsFile, err = os.OpenFile(c.lengthsName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		Log.Fatal("open ", c.lengthsName, " failed: ", err)
	}
	if redownload {
		c.truncZero()
	}
	lengths, err := ioutil.ReadFile(c.lengthsName)
	if err != nil {
		Log.Fatal("read ", c.lengthsName, " failed: ", err)
	}
	if (len(lengths) % 4) > 0 {
		// some kind of corruption, start over
		Log.Warningln("lengths file has unexpected length, re-downloading")
		c.truncZero()
		lengths = nil
	}

	// Because the ends of these files are more likely to have corruption
	// due to (for example) unclean shutdown, let's discard the last several
	// blocks (re-download them from zcashd).
	n := int(len(lengths)/4) - backup
	if n < 0 {
		n = 0
	}

	// The last entry in starts[] is where to write the next block.
	var offset int64
	c.starts = nil
	c.starts = append(c.starts, 0)
	var i int
	for i = 0; i < n; i++ {
		length := binary.LittleEndian.Uint32(lengths[i*4 : (i+1)*4])
		if length < 78 || length > 8*1000*1000 {
			// start over
			Log.Warningln("lengths file has impossible value, re-downloading")
			c.truncZero()
			offset = 0
			break
		}
		offset += int64(length) + 8
		c.starts = append(c.starts, offset)
		if i == 0 {
			// first block, sanity check it has Sapling activation height
			block := c.readBlock(c.firstBlock)
			if block == nil || int(block.Height) != startHeight {
				// start over
				Log.Warningln("first block looks wrong, re-downloading")
				c.truncZero()
				offset = 0
				break
			}
		}
		c.nextBlock++
	}

	// Make sure the files have the correct length (if unclean shutdown)
	if err := c.lengthsFile.Truncate(int64(i * 4)); err != nil {
		Log.Fatal("truncate blocks file failed: ", err)
	}
	if err := c.blocksFile.Truncate(offset); err != nil {
		Log.Fatal("truncate blocks file failed: ", err)
	}
	c.Sync()
	c.setLatestHash()
	return c
}

func dbFileNames(dbPath string, chainName string) (string, string) {
	return filepath.Join(dbPath, chainName, "lengths"),
		filepath.Join(dbPath, chainName, "blocks")
}

// Add adds the given block to the cache at the given height, returning true
// if a reorg was detected.
func (c *BlockCache) Add(height int, block *walletrpc.CompactBlock) error {
	// Invariant: m[firstBlock..nextBlock) are valid.
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if height > c.nextBlock {
		// Cache has been reset (for example, checksum error)
		return nil
	}
	if height < c.firstBlock {
		// Should never try to add a block before Sapling activation height
		Log.Fatal("cache.Add height below Sapling: ", height)
		return nil
	}
	if height < c.nextBlock {
		// Should never try to "backup" (call Reorg() instead).
		Log.Fatal("cache.Add height going backwards: ", height)
		return nil
	}
	bheight := int(block.Height)

	// TODO COINBASE-HEIGHT: restore this check after coinbase height is fixed
	if false && bheight != height {
		// This could only happen if zcashd returned the wrong
		// block (not the height we requested).
		Log.Fatal("cache.Add wrong height: ", bheight, " expecting: ", height)
		return nil
	}

	// Add the new block and its length to the db files.
	data, err := proto.Marshal(block)
	if err != nil {
		return err
	}
	_, err = c.blocksFile.Write(append(checksum(height, data), data...))
	if err != nil {
		Log.Fatal("blocks write failed: ", err)
	}
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, uint32(len(data)))
	_, err = c.lengthsFile.Write(b)
	if err != nil {
		Log.Fatal("lengths write failed: ", err)
	}

	// update the in-memory variables
	offset := c.starts[len(c.starts)-1]
	c.starts = append(c.starts, offset+int64(len(data)+8))

	if c.latestHash == nil {
		c.latestHash = make([]byte, len(block.Hash))
	}
	copy(c.latestHash, block.Hash)
	c.nextBlock++
	// Invariant: m[firstBlock..nextBlock) are valid.
	return nil
}

// Reorg resets nextBlock (the block that should be Add()ed next)
// downward to the given height.
func (c *BlockCache) Reorg(height int) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	// Allow the caller not to have to worry about Sapling start height.
	if height < c.firstBlock {
		height = c.firstBlock
	}
	if height >= c.nextBlock {
		// Timing window, ignore this request
		return
	}
	// Remove the end of the cache.
	c.nextBlock = height
	newCacheLen := height - c.firstBlock
	c.starts = c.starts[:newCacheLen+1]

	if err := c.lengthsFile.Truncate(int64(4 * newCacheLen)); err != nil {
		Log.Fatal("truncate failed: ", err)
	}
	if err := c.blocksFile.Truncate(c.starts[newCacheLen]); err != nil {
		Log.Fatal("truncate failed: ", err)
	}
	c.setLatestHash()
}

// Get returns the compact block at the requested height if it's
// in the cache, else nil.
func (c *BlockCache) Get(height int) *walletrpc.CompactBlock {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	if height < c.firstBlock || height >= c.nextBlock {
		return nil
	}
	return c.readBlock(height)
}

// GetLatestHeight returns the height of the most recent block, or -1
// if the cache is empty.
func (c *BlockCache) GetLatestHeight() int {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	if c.firstBlock == c.nextBlock {
		return -1
	}
	return c.nextBlock - 1
}

func (c *BlockCache) Sync() {
	c.lengthsFile.Sync()
	c.blocksFile.Sync()
}

// Currently used only for testing.
func (c *BlockCache) Close() {
	// Some operating system require you to close files before you can remove them.
	if c.lengthsFile != nil {
		c.lengthsFile.Close()
		c.lengthsFile = nil
	}
	if c.blocksFile != nil {
		c.blocksFile.Close()
		c.blocksFile = nil
	}
}
