package query

import (
    _ "net"
    "strings"
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
    fqdn := m.Question[0].Name
    fqdn = strings.ToLower(fqdn)
    switch r.Question[0].Qtype{
        case dns.TypeTXT:
            rr := handler.DNSHandler(fqdn, "TXT")
            m.Answer = []dns.RR{rr}
            break
        case dns.TypeA:
            rr := handler.DNSHandler(fqdn, "A")
            m.Answer = []dns.RR{rr}
            break
        case dns.TypeAAAA:
            rr := handler.DNSHandler(fqdn, "AAAA")
            m.Answer = []dns.RR{rr}
            break
        case dns.TypeCNAME:
            rr := handler.DNSHandler(fqdn, "CNAME")
            m.Answer = []dns.RR{rr}
            break
        default:
            break
    }
    w.WriteMsg(m)
}