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
	"github.com/oschwald/geoip2-golang"
    "github.com/smallnest/weighted"
)

func defaultOptions() {
	config.ConfigLocation = flag.String("c", "/etc/geodns/config.yml", "the location of the configuration file of DNS server")
	config.GeoLiteDBLocation = flag.String("g", "/etc/geodns/geolite2-city.mmdb", "the location of GeoLite2/GeoIP2 city MMDB")
	config.Port = flag.Int("p", 8053, "which port to listen")
	config.Debug = flag.Bool("D", false, "enable debug mode to print out more information while running the server")
	config.ListenAddress = flag.String("a", "127.0.0.1", "which address to listen for the request")
}

func main() {
	defaultOptions()
	flag.Parse()

	// Initial configMap
	config.FetchConfigMap(config.ConfigLocation)

	// Register domain
	query.RegisterDomain()

    // Initialize WeightedRR map
    config.WeightedRR = make(map[string]*weighted.RRW)
    // Register Weighted Round-Robin DNS records
    config.RegisterWeightedRRRecords()

	config.GeoDB, _ = geoip2.Open(*config.GeoLiteDBLocation)

	log.Printf("Starting DNS server...\n")

	go server.Serve(config.Port, "tcp", config.ListenAddress)
	go server.Serve(config.Port, "udp", config.ListenAddress)

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	s := <-sig
	log.Fatalf("Signal (%v) received, stopping\n", s)
}
