package config

import (
	"github.com/oschwald/geoip2-golang"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var (
	ConfigMap         map[string]interface{}
    ServerMapping     map[interface{}]interface{}
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
