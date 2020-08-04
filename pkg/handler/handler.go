package handler

import (
    "net"
    "strings"
    "github.com/miekg/dns"
    "github.com/TsundereChen/geodns-go/pkg/fetch"
)

var (
    configMap map[string]interface{}
)

func DNSHandler(fqdn string, recordType string)(rr *dns.A){
    // Get the subdomain information first
    for k := range configMap {
        if(strings.Contains(fqdn, k)){
            // Split FQDN into domain and subdomain
            subdomain := fetch.FetchSubDomainName(fqdn, k)
            // Now find if record exists.
            rrData := fetch.FetchRR(configMap[k])
            for rrName := range rrData {
                if(rrName == subdomain){
                    // Check if rr type match
                    rrType := fetch.FetchRrType(rrData[rrName])
                    if(rrType == "A"){
                        // Match
                    } else {
                        continue
                    }
                }
            }
        }
    }
    rr.Hdr = dns.RR_Header{
        Name: fqdn,
        Rrtype: dns.TypeA,
        Class: dns.ClassINET,
        Ttl: 3600}
    rr.A = net.IPv4(127,0,0,1)
    return rr
}
