package rr

import (
    "github.com/miekg/dns"
    "net"
)

func TypeAAAA(fqdn string, value string) (Rr *dns.AAAA){
    Rr = new(dns.AAAA)
    Rr.Hdr = dns.RR_Header{
        Name: fqdn,
        Rrtype: dns.TypeAAAA,
        Class: dns.ClassINET,
        Ttl: 3600}
    Rr.AAAA = net.ParseIP(value)
    return Rr
}
