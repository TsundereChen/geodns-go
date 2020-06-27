package main

import (
//    "net"
    "github.com/miekg/dns"
    "github.com/TsundereChen/geodns-go/handler"
)

func handleFunction(w dns.ResponseWriter, r *dns.Msg){
    m := new(dns.Msg)
    m.SetReply(r)
    m.Compress = *C
    m.Authoritative = true
    domain := m.Question[0].Name
    switch r.Question[0].Qtype{
        case dns.TypeTXT:

        case dns.TypeA:
            rr := handler.TypeAHandler(domain)
            m.Answer = []dns.RR{rr}
            break
        case dns.TypeAAAA:

        case dns.TypeCNAME:

        default:
    }
    w.WriteMsg(m)
}
