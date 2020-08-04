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
    var value string;
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
                    if(rrType == recordType){
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
    rr.Hdr = dns.RR_Header{
        Name: fqdn,
        Rrtype: dns.TypeA,
        Class: dns.ClassINET,
        Ttl: 3600}
    rr.A = net.ParseIP(value)
    return rr
}
