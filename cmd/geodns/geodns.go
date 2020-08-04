package geodns

import (
    "flag"
    "io/ioutil"
    "log"
    "os"
    "os/signal"
    "syscall"

    "gopkg.in/yaml.v2"
    "github.com/miekg/dns"
)

var (
    c *string
    g *string
    p *int
    D *bool
    a *string
    C *bool
    config map[string]interface{}
)

func main() {
    defaultOptions()
    flag.Parse()

    // Read in config.yaml
    var configFileRaw, err = ioutil.ReadFile(*c)
    if err != nil {
        panic(err)
    }
    yaml.Unmarshal([]byte(configFileRaw), &config)

    // Add domain into handleFunction
    var domainList = fetchDomain(config)
    for i := range domainList {
        dns.HandleFunc(domainList[i], handleFunction)
    }


    log.Printf("Starting DNS server...\n")

    go serve(p, "tcp", a)
    go serve(p, "udp", a)

    sig := make(chan os.Signal)
    signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
    s := <-sig
    log.Fatalf("Signal (%v) received, stopping\n", s)
}
