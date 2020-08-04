package serve

import (
	"github.com/miekg/dns"
	"log"
	"strconv"
)

func Serve(port *int, connType string, address *string) {
	srv := &dns.Server{Addr: *address + ":" + strconv.Itoa(*port), Net: connType}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Failed to set %s listener %s\n", connType, err.Error())
	}
}
