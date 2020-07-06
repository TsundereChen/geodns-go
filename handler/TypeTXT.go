package handler

import (
    "github.com/miekg/dns"
)

func TypeTXTHandler(domain string)(rr *dns.TXT){
    rr.Hdr = dns.RR_Header{
        Name: domain,
        Rrtype: dns.TypeTXT,
        Class: dns.ClassINET,
        Ttl: 3600}
    rr.Txt = []string{"sample string"}
    return rr
}
