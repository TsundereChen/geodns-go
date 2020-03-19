package main

import (
    "flag"
    "log"
    "os"
    "os/signal"
    "syscall"

    "github.com/miekg/dns"
)

var (
    c *string
    g *string
    p *int
    D *bool
    a *string
    dummySite string
)

func main() {
    defaultOptions()
    flag.Parse()

    dummySite = "example.com"

    dns.HandleFunc(dummySite, dummyHandleFunction)

    log.Printf("Starting DNS server...\n")

    go serve(p, "tcp", a)
    go serve(p, "udp", a)

    sig := make(chan os.Signal)
    signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
    s := <-sig
    log.Fatalf("Signal (%v) received, stopping\n", s)

}
