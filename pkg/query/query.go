package query

import (
	_ "fmt"
	"net"
	"strings"

	"github.com/TsundereChen/geodns-go/pkg/config"
	"github.com/TsundereChen/geodns-go/pkg/handler"
	"github.com/miekg/dns"
)

func RegisterDomain() {
	var domainList = config.FetchDomain(config.ConfigMap)
	for i := range domainList {
		if domainList[i] != "regions" {
			dns.HandleFunc(domainList[i], HandleFunction)
		}
	}
}

func HandleFunction(w dns.ResponseWriter, r *dns.Msg) {
    var (
        v4 bool
        sourceAddress net.IP
    )
	m := new(dns.Msg)
	m.SetReply(r)
	m.Authoritative = true

    if ip, ok := w.RemoteAddr().(*net.UDPAddr); ok {
        sourceAddress = ip.IP
        v4 = sourceAddress.To4() != nil
    }
    if ip, ok := w.RemoteAddr().(*net.TCPAddr); ok {
        sourceAddress = ip.IP
        v4 = sourceAddress.To4() != nil
    }
	fqdn := strings.ToLower(m.Question[0].Name)
	rr := handler.DNSHandler(fqdn, r.Question[0].Qtype, sourceAddress, v4)
	m.Answer = []dns.RR{rr}
	w.WriteMsg(m)
}
