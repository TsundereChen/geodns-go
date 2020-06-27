package handler

import (
    "net"
    "github.com/miekg/dns"
)

func TypeAAAAHandler(domain string)(rr *dns.AAAA){
    rr.Hdr = dns.RR_Header{
        Name: domain,
        Rrtype: dns.TypeAAAA,
        Class: dns.ClassINET,
        Ttl: 3600}
    rr.AAAA = net.ParseIP("2001:4860:4860::8888")
    return rr
}
