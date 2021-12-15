package forms

type TokenResult struct {
	TokenData []TokenData `json:"tokenData"`
}
type Tokens struct {
	AdditionalProp1 string `json:"additionalProp1"`
	AdditionalProp2 string `json:"additionalProp2"`
	AdditionalProp3 string `json:"additionalProp3"`
}
type TokenData struct {
	ProviderAddress string `json:"providerAddress"`
	ProviderName    string `json:"providerName"`
	ProviderPort    int    `json:"providerPort"`
	Tokens          Tokens `json:"tokens"`
}
