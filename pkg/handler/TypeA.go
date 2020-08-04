package handler

import (
    "net"
    "github.com/miekg/dns"
)

func TypeAHandler(domain string)(rr *dns.A){
    rr.Hdr = dns.RR_Header{
        Name: domain,
        Rrtype: dns.TypeA,
        Class: dns.ClassINET,
        Ttl: 3600}
    rr.A = net.IPv4(127,0,0,1)
    return rr
}
