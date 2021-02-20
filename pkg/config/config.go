package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
    "github.com/oschwald/geoip2-golang"
)

var (
	ConfigMap         map[string]interface{}
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
}

func FetchDomain(configMap map[string]interface{}) (domainList []string) {
	for domain := range configMap {
		if domain != "regions" {
			domainList = append(domainList, domain)
		}
	}
	return domainList
}
