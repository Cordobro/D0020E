package forms

type ServiceRequestForm struct {
	RequesterSystem    RequesterSystem      `json:"requesterSystem"`
	RequestedService   RequestedService     `json:"requestedService"`
	PreferredProviders []PreferredProviders `json:"preferredProviders"`
	OrchestrationFlags OrchestrationFlags   `json:"orchestrationFlags"`
}
type RequesterSystem struct {
	SystemName         string `json:"systemName"`
	Address            string `json:"address"`
	Port               int    `json:"port"`
	AuthenticationInfo string `json:"authenticationInfo"`
}

/* type metadataRequirements struct {
	AdditionalProp1 string `json:"additionalProp1"`
	AdditionalProp2 string `json:"additionalProp2"`
	AdditionalProp3 string `json:"additionalProp3"`
} */

type RequestedService struct {
	ServiceDefinitionRequirement string   `json:"serviceDefinitionRequirement"`
	InterfaceRequirements        []string `json:"interfaceRequirements"`
	SecurityRequirements         []string `json:"securityRequirements"`
	MetadataRequirements         []string `json:"metadataRequirementsGo"`
	VersionRequirement           int      `json:"versionRequirement"`
	MaxVersionRequirement        int      `json:"maxVersionRequirement"`
	MinVersionRequirement        int      `json:"minVersionRequirement"`
}

type ProviderCloud struct {
	Operator string `json:"operator"`
	Name     string `json:"name"`
}
type ProviderSystem struct {
	SystemName string `json:"systemName"`
	Address    string `json:"address"`
	Port       int    `json:"port"`
}
type PreferredProviders struct {
	ProviderCloud  ProviderCloud  `json:"providerCloud"`
	ProviderSystem ProviderSystem `json:"providerSystem"`
}
type OrchestrationFlags struct {
	AdditionalProp1 bool `json:"additionalProp1"`
	AdditionalProp2 bool `json:"additionalProp2"`
	AdditionalProp3 bool `json:"additionalProp3"`
}
