package forms

type InterCloudRule struct {
	Cloud struct {
		AuthenticationInfo string `json:"authenticationInfo"`
		GatekeeperRelayIds []int  `json:"gatekeeperRelayIds"`
		GatewayRelayIds    []int  `json:"gatewayRelayIds"`
		Name               string `json:"name"`
		Neighbor           bool   `json:"neighbor"`
		Operator           string `json:"operator"`
		Secure             bool   `json:"secure"`
	} `json:"cloud"`

	ProviderIdsWithInterfaceIds []struct {
		ID     int   `json:"id"`
		IDList []int `json:"idList"`
	} `json:"providerIdsWithInterfaceIds"`

	ServiceDefinition string `json:"serviceDefinition"`
}
