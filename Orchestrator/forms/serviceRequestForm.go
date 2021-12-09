package forms

type ServiceRequestForm struct {
	RequesterSystem struct {
		SystemName         string `json:"systemName"`
		Address            string `json:"address"`
		Port               int    `json:"port"`
		AuthenticationInfo string `json:"authenticationInfo"`
	} `json:"requesterSystem"`
	RequestedService struct {
		ServiceDefinitionRequirement string   `json:"serviceDefinitionRequirement"`
		InterfaceRequirements        []string `json:"interfaceRequirements"`
		SecurityRequirements         []string `json:"securityRequirements"`
		MetadataRequirements         struct {
			AdditionalProp1 string `json:"additionalProp1"`
			AdditionalProp2 string `json:"additionalProp2"`
			AdditionalProp3 string `json:"additionalProp3"`
		} `json:"metadataRequirements"`
		VersionRequirement    int `json:"versionRequirement"`
		MaxVersionRequirement int `json:"maxVersionRequirement"`
		MinVersionRequirement int `json:"minVersionRequirement"`
	} `json:"requestedService"`
	PreferredProviders []struct {
		ProviderCloud struct {
			Operator string `json:"operator"`
			Name     string `json:"name"`
		} `json:"providerCloud"`
		ProviderSystem struct {
			SystemName string `json:"systemName"`
			Address    string `json:"address"`
			Port       int    `json:"port"`
		} `json:"providerSystem"`
	} `json:"preferredProviders"`
	OrchestrationFlags struct {
		AdditionalProp1 bool `json:"additionalProp1"`
		AdditionalProp2 bool `json:"additionalProp2"`
		AdditionalProp3 bool `json:"additionalProp3"`
	} `json:"orchestrationFlags"`
}