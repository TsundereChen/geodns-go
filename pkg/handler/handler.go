package handler

import (
	"fmt"
	"github.com/TsundereChen/geodns-go/pkg/config"
	"github.com/TsundereChen/geodns-go/pkg/fetch"
	"github.com/miekg/dns"
	"net"
	"strings"
    "log"
)

func DNSHandler(fqdn string, questionType uint16, sourceAddress net.IP, IPv4 bool) (rr dns.RR) {
	// Get the subdomain information first
    record, err := config.GeoDB.Country(sourceAddress)
    if err != nil {
        log.Panic(err)
    }
	if *(config.Debug) == true {
        fmt.Printf("Source IP => %s\n", sourceAddress.String())
        fmt.Printf("Source Country => %s\n", record.Country.IsoCode)
		fmt.Printf("handler.DNSHandler handling request %s, question type %s\n", fqdn, dns.TypeToString[questionType])
	}
	var value string
	value = ""
	for k := range config.ConfigMap {
		if strings.Contains(fqdn, k) {
			// Split FQDN into domain and subdomain
			subdomain := fetch.FetchSubDomainName(fqdn, k)
			if *(config.Debug) == true {
				fmt.Printf("Searching subdomain %s in %s\n", subdomain, k)
			}
			// Now find if record exists.
			rrData := fetch.FetchRR(config.ConfigMap[k])
			for rrName := range rrData {
				if rrName == subdomain {
					// Check if rr type match
					rrType := fetch.FetchRrType(rrData[rrName])
					if typeChecker(rrType, questionType) {
						// Match
						// Use default value first
						value = fetch.FetchDefaultValue(rrData[rrName])
					} else {
						continue
					}
				}
			}
		}
	}
	rr = RrGenerator(questionType, fqdn, value)
	return rr
}

func typeChecker(rrType string, questionType uint16) (res bool) {
	return rrType == dns.TypeToString[questionType]
}
