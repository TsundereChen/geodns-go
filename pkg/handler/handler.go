package handler

import (
	"github.com/TsundereChen/geodns-go/pkg/config"
	"github.com/TsundereChen/geodns-go/pkg/fetch"
	"github.com/miekg/dns"
	_ "net"
	"strings"
)

func DNSHandler(fqdn string, questionType uint16) (rr dns.RR) {
	// Get the subdomain information first
	var value string
	for k := range config.ConfigMap {
		if strings.Contains(fqdn, k) {
			// Split FQDN into domain and subdomain
			subdomain := fetch.FetchSubDomainName(fqdn, k)
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
