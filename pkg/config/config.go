package config

func FetchDomain(configMap map[string]interface{}) (domainList []string){
    for domain := range configMap{
        if (domain != "regions"){
            domainList = append(domainList, domain)
        }
    }
    return domainList
}
