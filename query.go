package main

import (
    "net"
    "log"
    "github.com/miekg/dns"
)

func dummyHandleFunction(w dns.ResponseWriter, r *dns.Msg){
    m := new(dns.Msg)
    m.SetReply(r)
    m.Authoritative = true
    domain := m.Question[0].Name
    log.Printf("Received query request for %s\n", domain)
    rr := new(dns.A)
    rr.Hdr = dns.RR_Header{Name: domain,
                            Rrtype: dns.TypeA,
                            Class: dns.ClassINET,
                            Ttl: 3600}
    rr.A = net.IPv4(127,0,0,1)
    m.Answer = []dns.RR{rr}
    w.WriteMsg(m)
}
