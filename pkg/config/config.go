package config

func FetchDomain(config map[string]interface{}) (domainList []string){
    for domain := range config{
        if (domain != "regions"){
            domainList = append(domainList, domain)
        }
    }
    return domainList
}
