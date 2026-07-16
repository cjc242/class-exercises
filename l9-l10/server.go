package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"sync"
	"time"
)

type PeerServer struct {
	Address string
	Client  *rpc.Client
}

type Args struct {
	GossipLive map[string]int
	Round      int
	Sender     string
}

type Server struct {
	live    map[string]int
	lock    sync.Mutex
	Round   int
	Address string
	peers   []PeerServer
}

func (t *Server) Heartbeat(args *Args, reply *int) error {

}

func (t *Server) sendHeartbeat(to PeerServer) {

}

func (t *Server) GenerateReport() {

}

func main() {

	server := new(Server)
	rpc.Register(server)
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	go http.Serve(l, nil)

	my_address := "10.239.244.33"
	server.Address = my_address
	server.Round = 0
	server.peers = make([]PeerServer, 0)
	server.live = make(map[string]int)
	peer_addresses := []string{"10.193.25.197", "10.239.38.177", "10.239.246.218"}

	time.Sleep(10 * time.Second) // WAIT to start other servers

	for _, addr := range peer_addresses {
		if addr == my_address {
			continue
		}
		client, err := rpc.DialHTTP("tcp", addr)
		if err != nil {
			log.Fatal("dialing:", err)
		}
		server.peers = append(server.peers, PeerServer{addr, client})
	}

	/*
		TODO: call send heartbeats to a random server every second
			- NOTE: ensure that this code is non-blocking!
		TODO: call generate report every 5 seconds
	*/

}
