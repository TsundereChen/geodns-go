package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var (
	ConfigMap map[string]interface{}
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
