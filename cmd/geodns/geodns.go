package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/TsundereChen/geodns-go/pkg/config"
	"github.com/TsundereChen/geodns-go/pkg/query"
	server "github.com/TsundereChen/geodns-go/pkg/serve"
)

var (
	c         *string
	g         *string
	p         *int
	D         *bool
	a         *string
	configMap map[string]interface{}
)

func defaultOptions() {
	c = flag.String("c", "/etc/geodns/config.yml", "the location of the configuration file of DNS server")
	g = flag.String("g", "/etc/geodns/geolite2-city.mmdb", "the location of GeoLite2/GeoIP2 city MMDB")
	p = flag.Int("p", 8053, "which port to listen")
	D = flag.Bool("D", false, "enable debug mode to print out more information while running the server")
	a = flag.String("a", "127.0.0.1", "which address to listen for the request")
}

func main() {
	defaultOptions()
	flag.Parse()

	// Initial configMap
	config.FetchConfigMap(c)

	// Register domain
	query.RegisterDomain()

	log.Printf("Starting DNS server...\n")

	go server.Serve(p, "tcp", a)
	go server.Serve(p, "udp", a)

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	s := <-sig
	log.Fatalf("Signal (%v) received, stopping\n", s)
}
