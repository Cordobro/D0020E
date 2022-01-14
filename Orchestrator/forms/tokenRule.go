package forms

type TokenRule struct {
	Consumer      Consumer      `json:"consumer"`
	ConsumerCloud ConsumerCloud `json:"consumerCloud"`
	Duration      int           `json:"duration"`
	Providers     []Providers   `json:"providers"`
	Service       string        `json:"service"`
}
type Consumer struct {
	Address            string `json:"address"`
	AuthenticationInfo string `json:"authenticationInfo"`
	Port               int    `json:"port"`
	SystemName         string `json:"systemName"`
}
type ConsumerCloud struct {
	AuthenticationInfo string `json:"authenticationInfo"`
	GatekeeperRelayIds []int  `json:"gatekeeperRelayIds"`
	GatewayRelayIds    []int  `json:"gatewayRelayIds"`
	Name               string `json:"name"`
	Neighbor           bool   `json:"neighbor"`
	Operator           string `json:"operator"`
	Secure             bool   `json:"secure"`
}
type Provider2 struct {
	Address            string `json:"address"`
	AuthenticationInfo string `json:"authenticationInfo"`
	Port               int    `json:"port"`
	SystemName         string `json:"systemName"`
}
type Providers struct {
	Provider          Provider2 `json:"provider"`
	ServiceInterfaces []string  `json:"serviceInterfaces"`
}

func ConstructTokenRule(srf *ServiceRequestForm) *TokenRule {

	var tokenRule TokenRule

	//Consumer
	tokenRule.Consumer.Address = srf.RequesterSystem.Address
	tokenRule.Consumer.AuthenticationInfo = srf.RequesterSystem.AuthenticationInfo
	tokenRule.Consumer.Port = srf.RequesterSystem.Port
	tokenRule.Consumer.SystemName = srf.RequesterSystem.SystemName

	//ConsumerCloud
	//Where do we get these?
	/*
		tokenRule.ConsumerCloud.AuthenticationInfo =
		tokenRule.ConsumerCloud.GatekeeperRelayIds =
		tokenRule.ConsumerCloud.GatewayRelayIds =
		tokenRule.ConsumerCloud.Name =
		tokenRule.ConsumerCloud.Neighbor =
		tokenRule.ConsumerCloud.Operator =
		tokenRule.ConsumerCloud.Secure =
	*/

	//Duration
	//tokenRule.Duration =

	//Providers
	//provider

	//Service
	//tokenRule.Service

	return &tokenRule
}
