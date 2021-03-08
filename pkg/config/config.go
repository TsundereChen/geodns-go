package config

import (
	"crypto/sha1"
	"github.com/TsundereChen/geodns-go/pkg/fetch"
	"github.com/oschwald/geoip2-golang"
	"github.com/smallnest/weighted"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var (
	ConfigMap         map[string]interface{}
	ServerMapping     map[interface{}]interface{}
	WeightedRR        map[string]*weighted.RRW
	ConfigLocation    *string
	GeoLiteDBLocation *string
	Port              *int
	Debug             *bool
	ListenAddress     *string
	GeoDB             *geoip2.Reader
)

func FetchConfigMap(c *string) {
	var configFile, err = ioutil.ReadFile(*c)
	if err != nil {
		panic(err)
	}
	yaml.Unmarshal([]byte(configFile), &ConfigMap)
	ServerMapping = ConfigMap["regions"].(map[interface{}]interface{})
}

func FetchDomain(configMap map[string]interface{}) (domainList []string) {
	for domain := range configMap {
		if domain != "regions" {
			domainList = append(domainList, domain)
		}
	}
	return domainList
}

func RegisterWeightedRRRecords() {
	for domain := range ConfigMap {
		if domain != "regions" {
			rrData := fetch.FetchRR(ConfigMap[domain])
			for rr := range rrData {
				rules := fetch.FetchRules(rrData[rr])
				for rule := range rules {
					values := rules[rule].(map[interface{}]interface{})["value"].([]interface{})
					region := rules[rule].(map[interface{}]interface{})["region"].(string)
					time := rules[rule].(map[interface{}]interface{})["time"].(string)
					hashString := Hash(rr.(string) + domain + region + time)
					WeightedRR[hashString] = &weighted.RRW{}
					for value := range values {
						valueMap := values[value].(map[interface{}]interface{})
						for k, v := range valueMap {
							WeightedRR[hashString].Add(k, v.(int))
						}
					}
				}
			}
		}
	}
}

func Hash(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	return string(bs)
}
