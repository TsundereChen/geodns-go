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
            rr = rrLib.TypeAAAA(fqdn, value)
            break
        case dns.TypeCNAME:
            rr = rrLib.TypeCNAME(fqdn, value)
            break
        case dns.TypeTXT:
            rr = rrLib.TypeTXT(fqdn, value)
            break
        default:
            rr = new(dns.NULL)
            break
    }
    return rr
}
