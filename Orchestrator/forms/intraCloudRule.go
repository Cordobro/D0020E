package forms

/********************************* IntraCloudRule ******************************************/

type IntraCloudRule struct {
	Cloud                       Cloud                         `json:"cloud"`
	ProviderIdsWithInterfaceIds []ProviderIdsWithInterfaceIds `json:"providerIdsWithInterfaceIds"`
	ServiceDefinition           string                        `json:"serviceDefinition"`
}

type Cloud struct {
	AuthenticationInfo string `json:"authenticationInfo"`
	GatekeeperRelayIds []int  `json:"gatekeeperRelayIds"`
	GatewayRelayIds    []int  `json:"gatewayRelayIds"`
	Name               string `json:"name"`
	Neighbor           bool   `json:"neighbor"`
	Operator           string `json:"operator"`
	Secure             bool   `json:"secure"`
}

type ProviderIdsWithInterfaceIds struct {
	ID     int   `json:"id"`
	IDList []int `json:"idList"`
}
