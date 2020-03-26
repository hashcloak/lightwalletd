package common

import (
	"bufio"
	"encoding/hex"
	"encoding/json"
	"os"
	"strconv"
	"time"

	"github.com/zcash/lightwalletd/walletrpc"
)

var DarksideEnable bool

var darksideState struct {
	ldinfo                walletrpc.LightdInfo
	cblocks               []walletrpc.CompactBlock
	incoming_transactions [][]byte
	server_start          time.Time
}

func DarksideInit() {
	DarksideEnable = true
	// Hard-code for mainnet, at least for now
	darksideState.ldinfo.Version = "0"
	darksideState.ldinfo.Vendor = "ECC Darkside"
	darksideState.ldinfo.TaddrSupport = true
	darksideState.ldinfo.ChainName = "main"
	darksideState.ldinfo.ConsensusBranchId = "2bb40e60"

	testBlocks, err := os.Open("./testdata/darkside-cblocks")
	if err != nil {
		Log.Fatal("Error loading default darksidewalletd blocks")
	}
	scan := bufio.NewScanner(testBlocks)
	var height int
	for scan.Scan() { // each line (first line height, then blocks)
		if height == 0 {
			// first line is starting height of the blocks that follow
			height, err = strconv.Atoi(string(scan.Bytes()))
			if err != nil {
				Log.Fatal("Error loading BlockHeight from darkside-cblocks")
			}
			darksideState.ldinfo.SaplingActivationHeight = uint64(height)
			darksideState.ldinfo.BlockHeight = uint64(height)
			continue
		}
		d, err := hex.DecodeString(string(scan.Bytes()))
		if err != nil {
			Log.Fatal("Error decoding block from darkside-cblocks")
		}
		var cb walletrpc.CompactBlock
		err = json.Unmarshal(d, &cb)
		if err != nil {
			Log.Fatal("Error unmarshalling block from darkside-cblocks")
		}
		if int(cb.Height) != height {
			Log.Fatal("Error, unexpected coinbase height from darkside-cblocks")
		}
		height++
		darksideState.cblocks = append(darksideState.cblocks, cb)
	}
}

func DarksideGetSaplingInfo() (int, int, string, string) {
	return int(darksideState.ldinfo.SaplingActivationHeight),
		int(darksideState.ldinfo.BlockHeight),
		darksideState.ldinfo.ChainName,
		darksideState.ldinfo.ConsensusBranchId
}

func DarksideGetLightdInfo() *walletrpc.LightdInfo {
	return &darksideState.ldinfo
}

// Setting reorg to X means simulates a new version of X, but it has the same
// predecessor (parent), i.e., its hash changes but not its prevhash.
// All blocks following X have both their hash and prevhash updated.
func DarksideSetState(state *walletrpc.DarksideState) {
	sapling := darksideState.ldinfo.SaplingActivationHeight
	if state.LatestHeight < sapling {
		Log.Fatal("DarksideSetState: latestHeight can't be less than sapling")
	}
	if int(state.LatestHeight-sapling) >= len(darksideState.cblocks) {
		Log.Fatal("DarksideSetState: latestHeight too high")
	}
	if int(state.ReorgHeight-sapling) >= len(darksideState.cblocks) {
		Log.Fatal("DarksideSetState: reorgHeight too high")
	}
	if state.ReorgHeight > 0 {
		if state.ReorgHeight < sapling {
			Log.Fatal("DarksideSetState: reorgHeight can't be less than sapling")
		}
		var prevHash *byte
		for h := int(state.ReorgHeight - sapling); h < len(darksideState.cblocks); h++ {
			cblock := &darksideState.cblocks[h]
			// Changing the hash in any way simulates a new version of the block.
			cblock.Hash[0]++
			if prevHash != nil {
				cblock.PrevHash[0] = *prevHash
			}
			prevHash = &cblock.Hash[0]
		}
	}
	darksideState.ldinfo.BlockHeight = state.LatestHeight
}
