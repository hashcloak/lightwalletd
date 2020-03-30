package common

import (
	"bufio"
	"encoding/json"
	"os"
	"time"

	"github.com/zcash/lightwalletd/walletrpc"
)

var DarksideEnable bool

var darksideState struct {
	ldinfo                walletrpc.LightdInfo
	cblocks               []*walletrpc.CompactBlock
	transactions          map[string][]byte // txid: txdata
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

	{
		testBlocks, err := os.Open("./testdata/darkside-cblocks")
		if err != nil {
			Log.Fatal("Error loading darksidewalletd blocks")
		}
		scan := bufio.NewScanner(testBlocks)
		for scan.Scan() { // each line is a JSON-formatted compact block
			var cb walletrpc.CompactBlock
			err = json.Unmarshal(scan.Bytes(), &cb)
			if err != nil {
				Log.Fatal("Error unmarshalling block from darkside-cblocks")
			}
			// Pretend the first block in the file is the Sapling height.
			if darksideState.ldinfo.SaplingActivationHeight == 0 {
				darksideState.ldinfo.SaplingActivationHeight = cb.Height
				darksideState.ldinfo.BlockHeight = cb.Height
			}
			darksideState.cblocks = append(darksideState.cblocks, &cb)
		}
	}

	{
		darksideState.transactions = make(map[string][]byte)
		testTx, err := os.Open("./testdata/darkside-transactions")
		if err != nil {
			Log.Fatal("Error loading default darksidewalletd transactions")
		}
		scan := bufio.NewScanner(testTx)
		for scan.Scan() { // each line
			txJSON := scan.Bytes()
			var txinfo interface{}
			err = json.Unmarshal(txJSON, &txinfo)
			if err != nil {
				Log.Fatal("Error unmarshalling block from darkside-transactions")
			}
			txid := txinfo.(map[string]interface{})["txid"].(string)
			darksideState.transactions[txid] = make([]byte, len(txJSON))
			copy(darksideState.transactions[txid], txJSON)
		}
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

func DarksideGetTransaction(txid string) []byte {
	if t, ok := darksideState.transactions[txid]; ok {
		return t
	}
	return nil
}

// Setting reorg to X means simulates a new version of X, but it has the same
// predecessor (parent), i.e., its hash changes but not its prevhash.
// All blocks following X have both their hash and prevhash updated.
func DarksideSetState(state *walletrpc.DarksideState) {
	sapling := darksideState.ldinfo.SaplingActivationHeight
	if state.LatestHeight < sapling {
		Log.Fatal("DarksideSetState: latestHeight ", state.LatestHeight, " can't be less than sapling ", sapling)
	}
	if int(state.LatestHeight-sapling) >= len(darksideState.cblocks) {
		Log.Fatal("DarksideSetState: latestHeight ", state.LatestHeight, " is too high, max ", int(sapling)+len(darksideState.cblocks)-1)
	}
	if int(state.ReorgHeight-sapling) >= len(darksideState.cblocks) {
		Log.Fatal("DarksideSetState: reorgHeight ", state.ReorgHeight, " is too high, max ", int(sapling)+len(darksideState.cblocks)-1)
	}
	if state.ReorgHeight > 0 {
		if state.ReorgHeight < sapling {
			Log.Fatal("DarksideSetState: reorgHeight ", state.ReorgHeight, " can't be less than sapling ", sapling)
		}
		var prevHash *byte
		for h := int(state.ReorgHeight - sapling); h < len(darksideState.cblocks); h++ {
			cblock := darksideState.cblocks[h]
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
