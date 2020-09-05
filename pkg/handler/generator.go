package handler
import (
    "github.com/miekg/dns"
    rrLib "github.com/TsundereChen/geodns-go/handler/pkg/rr"
)

func RrGenerator(requestType uint16, fqdn string, value string) (rr dns.RR){
    switch requestType {
        case dns.TypeA:
            rr = rrLib.TypeA(fqdn, value)
            break
        case dns.TypeAAAA:
            rr = new(dns.AAAA)
            break
        case dns.TypeCNAME:
            rr = new (dns.CNAME)
            break
        case dns.TypeTXT:
            rr = new(dns.TXT)
            break
        default:
            rr = new(dns.NULL)
            break
    }
    return rr
}
