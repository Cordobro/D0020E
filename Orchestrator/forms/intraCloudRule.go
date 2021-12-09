package forms

/********************************* IntraCloudRule ******************************************/

type IntraCloudRule struct {
	Consumer struct {
		Address            string `json:"address"`
		AuthenticationInfo string `json:"authenticationInfo"`
		Port               int    `json:"port"`
		SystemName         string `json:"systemName"`
	} `json:"consumer"`
	ProviderIdsWithInterfaceIds []struct {
		ID     int   `json:"id"`
		IDList []int `json:"idList"`
	} `json:"providerIdsWithInterfaceIds"`
	ServiceDefinitionID int `json:"serviceDefinitionId"`
}
