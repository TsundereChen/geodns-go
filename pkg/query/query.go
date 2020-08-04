package query

import (
    _ "net"
    "github.com/miekg/dns"
    "github.com/TsundereChen/geodns-go/pkg/handler"
)

var (
    C *bool
)

func HandleFunction(w dns.ResponseWriter, r *dns.Msg){
    m := new(dns.Msg)
    m.SetReply(r)
    m.Compress = *C
    m.Authoritative = true
    domain := m.Question[0].Name
    switch r.Question[0].Qtype{
        case dns.TypeTXT:
            rr := handler.TypeTXTHandler(domain)
            m.Answer = []dns.RR{rr}
            break
        case dns.TypeA:
            rr := handler.TypeAHandler(domain)
            m.Answer = []dns.RR{rr}
            break
        case dns.TypeAAAA:
            rr := handler.TypeAAAAHandler(domain)
            m.Answer = []dns.RR{rr}
            break
        case dns.TypeCNAME:
            rr := handler.TypeCNAMEHandler(domain)
            m.Answer = []dns.RR{rr}
            break
        default:
    }
    w.WriteMsg(m)
}
