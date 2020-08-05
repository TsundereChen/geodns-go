package handler

import (
	"fmt"
	"net"
	"strings"

	"github.com/TsundereChen/geodns-go/pkg/config"
	"github.com/TsundereChen/geodns-go/pkg/fetch"
	"github.com/miekg/dns"
)

func DNSHandler(fqdn string, recordType string) (rr *dns.A) {
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
					if rrType == recordType {
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
	rr = new(dns.A)
	// Create RR according to request
	rr.Hdr = dns.RR_Header{
		Name:   fqdn,
		Rrtype: dns.TypeA,
		Class:  dns.ClassINET,
		Ttl:    3600}
	rr.A = net.ParseIP(value)
	return rr
}
