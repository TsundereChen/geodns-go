package handler

import (
    "github.com/miekg/dns"
)

func TypeCNAMEHandler(domain string)(rr *dns.CNAME){
    rr.Hdr = dns.RR_Header{
        Name: domain,
        Rrtype: dns.TypeCNAME,
        Class: dns.ClassINET,
        Ttl: 3600}
    rr.Target = "www.google.com"
    return rr
}
