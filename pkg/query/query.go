package query

import (
	"log"
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
		v4            bool
		TCP           bool
		sourceAddress net.IP
	)
	m := new(dns.Msg)
	m.SetReply(r)
	m.Authoritative = true

	if ip, ok := w.RemoteAddr().(*net.UDPAddr); ok {
		sourceAddress = ip.IP
		v4 = sourceAddress.To4() != nil
		TCP = false
	}
	if ip, ok := w.RemoteAddr().(*net.TCPAddr); ok {
		sourceAddress = ip.IP
		v4 = sourceAddress.To4() != nil
		TCP = true
	}
	if *config.Debug {
		if TCP {
			log.Printf("Got incoming request TCP.\n")
		} else {
			log.Printf("Got incoming request UDP.\n")
		}
	}
	fqdn := strings.ToLower(m.Question[0].Name)
	rr := handler.DNSHandler(fqdn, r.Question[0].Qtype, sourceAddress, v4)
    if rr != nil {
        // Only add Answer part if rr isn't nil
	    m.Answer = []dns.RR{rr}
    }
    m.Ns = append(m.Ns, handler.SOAHandler(m.Question[0].Name))
	w.WriteMsg(m)
}
