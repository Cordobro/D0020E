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
type Provider struct {
	Address            string `json:"address"`
	AuthenticationInfo string `json:"authenticationInfo"`
	Port               int    `json:"port"`
	SystemName         string `json:"systemName"`
}
type Providers struct {
	Provider          Provider `json:"provider"`
	ServiceInterfaces []string `json:"serviceInterfaces"`
}
