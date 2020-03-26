// This program generates the hex form of the compact blocks in the given
// range, one per line. It requires zcashd to be running. This program
// generated the contents of testdata/darkside-cblocks (on mainnet).
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
	if len(os.Args) != 4 {
		fmt.Println("usage:", os.Args[0], "zcash.conf start-offset end-offset")
		os.Exit(1)
	}
	start, err := strconv.ParseInt(os.Args[2], 10, 32)
	if err != nil {
		fmt.Println("bad start-offset:", err)
		os.Exit(1)
	}
	end, err := strconv.ParseInt(os.Args[3], 10, 32)
	if err != nil {
		fmt.Println("bad end-offset:", err)
		os.Exit(1)
	}
	rpcClient, err := frontend.NewZRPCFromConf(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "rpcClient.RawRequest error: ", err)
	}

	// The first line in the file is the start height, in decimal
	fmt.Println(start)

	for height := start; height <= end; height++ {
		params := make([]json.RawMessage, 2)
		params[0] = json.RawMessage("\"" + strconv.Itoa(int(height)) + "\"")
		params[1] = json.RawMessage("0") // non-verbose (raw hex)
		result, rpcErr := rpcClient.RawRequest("getblock", params)
		if rpcErr != nil {
			fmt.Fprintln(os.Stderr, "rpcClient.RawRequest error: ", rpcErr)
			os.Exit(1)
		}

		var blockDataHex string
		err := json.Unmarshal(result, &blockDataHex)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error reading JSON response", rpcErr)
			os.Exit(1)
		}

		blockData, err := hex.DecodeString(blockDataHex)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error decoding getblock output", rpcErr)
			os.Exit(1)
		}

		block := parser.NewBlock()
		rest, err := block.ParseFromSlice(blockData)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error parsing block", rpcErr)
			os.Exit(1)
		}
		if len(rest) != 0 {
			fmt.Fprintln(os.Stderr, "received overlong message", rpcErr)
			os.Exit(1)
		}

		if block.GetHeight() != int(height) {
			fmt.Fprintln(os.Stderr, "received unexpected height block", rpcErr)
			os.Exit(1)
		}

		cb := block.ToCompact()
		m, err := json.Marshal(cb)
		if err != nil {
			fmt.Println("Marshal err: ", err)
			os.Exit(1)
		}
		e := hex.EncodeToString(m)
		fmt.Println(e)

		// the following isn't necessary, but it demonstrates how
		// to deserialize e (file contents) back to a compact block
		d, err := hex.DecodeString(e)
		if err != nil {
			fmt.Println("Decode err: ", err)
			os.Exit(1)
		}
		var cb1 walletrpc.CompactBlock
		err = json.Unmarshal(d, &cb1)
		if err != nil {
			fmt.Println("u: err: ", err, cb1)
			os.Exit(1)
		}
		if cb.Height != cb1.Height {
			fmt.Println("not equal")
			os.Exit(1)
		}
	}
}
