package fetch

import (
    "strings"
)

func FetchRR(config interface{}) (rrData map[string]interface{}){
    return config.(map[string]interface{})["rr"].(map[string]interface{})
}

func FetchSubDomainName(fqdn string, domain string) (subDomain string){
    return strings.Split(strings.Split(fqdn, domain)[0], ".")[0]
}

func FetchRrType(rrData interface{}) (rrType string){
    return rrData.(map[string]interface{})["type"].(string)
}
