package main

import (
    "flag"
)

func defaultOptions(){
    c = flag.String("c", "/etc/geodns/config.yml", "the location of the configuration file of DNS server")
    g = flag.String("g", "/etc/geodns/geolite2-city.mmdb", "the location of GeoLite2/GeoIP2 city MMDB")
    p = flag.Int("p", 8053, "which port to listen")
    D = flag.Bool("D", false, "enable debug mode to print out more information while running the server")
    a = flag.String("a", "127.0.0.1", "which address to listen for the request")
}
