// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// TeamsCatalogApp undocumented
type TeamsCatalogApp struct {
	// Entity is the base model of TeamsCatalogApp
	Entity
	// ExternalID undocumented
	ExternalID *string `json:"externalId,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// DistributionMethod undocumented
	DistributionMethod *TeamsAppDistributionMethod `json:"distributionMethod,omitempty"`
}
