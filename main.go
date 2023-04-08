package main

import (
	"flag"
	"fmt"
	Server "github.com/mhthrh/TimeSeriesDb/server"
	"log"
	"net/http"
)

var (
	ip   string
	port int
)

func main() {

	flag.StringVar(&ip, "IP", "localhost", "ip address")
	flag.IntVar(&port, "PORT", 8585, "port number")
	flag.Parse()
	fmt.Printf("server running on ip:%s and port %d", ip, port)
	server := http.Server{
		Addr:    fmt.Sprintf("%s:%d", ip, port),
		Handler: Server.RunServer(),
	}
	log.Fatalln(server.ListenAndServe())
}
