
package main

import (
	"net/rpc/jsonrpc"

	"github.com/rs/zerolog/log"
)

type Args struct {
	ProductId   int
	ProductName string
}

func main() {
	client, err := jsonrpc.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err.Error())
	}
	args := Args{ProductId: 1,ProductName: "new product"}
	var result int
	if err := client.Call("Shop.Add", args, &result); err != nil {
		log.Fatal().Err(err).Msg("Can't call calculator method")
	}
	log.Info().Msgf("%v + %v = %v", args.ProductId, args.ProductName, result)
}
