package main

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	"github.com/rs/zerolog/log"
)

// 1. Створити клієнт і сервер, використовуючи json-rpc:
// Сервер виконує роль кошика інтернет-магазина. Кошик очікує id і name товара. Додає до кошику, оновлює і видаляє.
// Клієнт виконує роль клієнта і працює із кошиком: додате, видаляє і редагує товари
type Args struct {
	A, B int
}
type Shop int

func (c *Shop) Add(args *Args, result *int) error {
	log.Info().Msg("Operation Add")
	*result = args.A + args.B
	return nil
}

func main() {
	log.Info().Msgf("Starting RPC server")
	var Shop Shop

	rpcServer := rpc.NewServer()
	rpcServer.Register(&Shop)
	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatal().Err(err).Msg("Can't create tcp listener")
	}

	log.Info().Msg("Started RPC server")

	for {
		conn, err := listener.Accept()

		if err != nil {
			continue
		}

		log.Info().Msg("Accepted new connection")

		go rpcServer.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
