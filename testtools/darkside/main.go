// This program fetches the given range of blocks, and the sapling transactions within
// those blocks, from a locally-running zcashd. The zcashd can be running mainnet or
// testnet, but is expected to be mainnet. The two output files contain the *compact*
// blocks in the given range, and the *full* transactions in that range. These files
// are then read by lightwalletd when running in --darkside-very-insecure (test) mode
// so it can serve these compact blocks and transactions without requiring zcashd.
//
// The range of blocks in the current source-controlled versions of these generated
// files is: 663150 - 663250 (because this range contains several of the wallet team's
// "developer wallet" transactions, so enables them to test their code).
//
// This program is based on common/common.go: getBlockFromRPC().
package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/zcash/lightwalletd/frontend"
	"github.com/zcash/lightwalletd/parser"
	"github.com/zcash/lightwalletd/walletrpc"
	"os"
	"strconv"
)

func main() {
	if err := genfiles(); err != "" {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func genfiles() string {
	if len(os.Args) != 4 {
		return fmt.Sprint("usage:", os.Args[0], "zcash.conf start-offset end-offset")
	}
	start, err := strconv.ParseInt(os.Args[2], 10, 32)
	if err != nil {
		return fmt.Sprint("bad start-offset:", err)
	}
	end, err := strconv.ParseInt(os.Args[3], 10, 32)
	if err != nil {
		return fmt.Sprint("bad end-offset:", err)
	}
	rpcClient, err := frontend.NewZRPCFromConf(os.Args[1])
	if err != nil {
		return fmt.Sprint("rpcClient.RawRequest error:", err)
	}

	// This file will have the compact blocks, one per line
	fcblocks, err := os.OpenFile("testdata/darkside-cblocks", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Sprint("open testdata/darkside-cblocks failed:", err)
	}
	// This file will have the full transactions, one per line
	ftx, err := os.OpenFile("testdata/darkside-transactions", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Sprint("open testdata/darkside-transactions failed:", err)
	}

	for height := start; height <= end; height++ {
		params := make([]json.RawMessage, 2)
		params[0] = json.RawMessage("\"" + strconv.Itoa(int(height)) + "\"")
		params[1] = json.RawMessage("0") // non-verbose (raw hex)
		result, rpcErr := rpcClient.RawRequest("getblock", params)
		if rpcErr != nil {
			return fmt.Sprint("rpcClient.RawRequest error: ", rpcErr)
		}

		var blockDataHex string
		err := json.Unmarshal(result, &blockDataHex)
		if err != nil {
			return fmt.Sprint("error reading JSON response", rpcErr)
		}

		blockData, err := hex.DecodeString(blockDataHex)
		if err != nil {
			return fmt.Sprint("error decoding getblock output", rpcErr)
		}

		block := parser.NewBlock()
		rest, err := block.ParseFromSlice(blockData)
		if err != nil {
			return fmt.Sprint("error parsing block", rpcErr)
		}
		if len(rest) != 0 {
			return fmt.Sprint("received overlong message", rpcErr)
		}

		if block.GetHeight() != int(height) {
			return fmt.Sprint("received unexpected height block", rpcErr)
		}

		cb := block.ToCompact()
		m, err := json.Marshal(cb)
		if err != nil {
			return fmt.Sprint("Marshal err: ", err)
		}
		fcblocks.WriteString(string(m) + "\n")

		for _, tx := range block.Transactions() {
			// This is simpler than blocks, because we're writing full
			// transactions as recieved from zcashd, not compact transactions.
			if tx.HasSaplingElements() {
				dhash := hex.EncodeToString(tx.GetDisplayHash())
				params := []json.RawMessage{
					json.RawMessage("\"" + dhash + "\""),
					json.RawMessage("1"),
				}
				result, rpcErr := rpcClient.RawRequest("getrawtransaction", params)
				if rpcErr != nil {
					return fmt.Sprint("rpcClient.RawRequest GetTransaction error: ", rpcErr)
				}
				ftx.WriteString(string(result) + "\n")
			}
		}

		// the following isn't necessary, but it demonstrates how
		// to deserialize m (file contents) back to a compact block
		var cb1 walletrpc.CompactBlock
		err = json.Unmarshal(m, &cb1)
		if err != nil {
			return fmt.Sprint("u: err: ", err, cb1)
		}
		if cb.Height != cb1.Height {
			return fmt.Sprint("not equal")
		}
	}
	return ""
}
