package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"strings"

	"github.com/hashicorp/go-plugin"
	"github.com/mwantia/nautilus-debug/pkg/core"
	"github.com/mwantia/nautilus/pkg/shared"
)

var (
	Address = flag.String("address", "", "Define the address used to serve this plugin over network.")
	Network = flag.String("network", "tcp", "Define the network used to serve this plugin over network. (Default: tcp)")
)

func main() {
	flag.Parse()

	impl := core.NewImpl()

	if IsNetwork() {
		listener, err := net.Listen(*Network, *Address)
		if err != nil {
			log.Fatalf("Failed to start listener: %v", err)
		}

		defer listener.Close()

		fmt.Printf("Plugin server listening on %s://%s\n", *Network, *Address)

		server := rpc.NewServer()
		err = server.RegisterName("Plugin", &shared.NautilusRPCServer{
			Impl: impl,
		})
		if err != nil {
			log.Fatalf("Failed to register RPC server: %v", err)
		}

		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Printf("Failed to accept connection: %v", err)
			}
			go server.ServeConn(conn)
		}
	} else {
		plugin.Serve(&plugin.ServeConfig{
			HandshakeConfig: shared.Handshake,
			Plugins: map[string]plugin.Plugin{
				"pipeline": &shared.NautilusPipelinePlugin{
					Impl: impl,
				},
			},
		})
	}
}

func IsNetwork() bool {
	return strings.TrimSpace(*Network) != "" && strings.TrimSpace(*Address) != ""
}
