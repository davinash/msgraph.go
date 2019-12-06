// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// Device undocumented
type Device struct {
	// DirectoryObject is the base model of Device
	DirectoryObject
	// AccountEnabled undocumented
	AccountEnabled *bool `json:"accountEnabled,omitempty"`
	// AlternativeSecurityIDs undocumented
	AlternativeSecurityIDs []AlternativeSecurityID `json:"alternativeSecurityIds,omitempty"`
	// ApproximateLastSignInDateTime undocumented
	ApproximateLastSignInDateTime *time.Time `json:"approximateLastSignInDateTime,omitempty"`
	// ComplianceExpirationDateTime undocumented
	ComplianceExpirationDateTime *time.Time `json:"complianceExpirationDateTime,omitempty"`
	// DeviceID undocumented
	DeviceID *string `json:"deviceId,omitempty"`
	// DeviceMetadata undocumented
	DeviceMetadata *string `json:"deviceMetadata,omitempty"`
	// DeviceVersion undocumented
	DeviceVersion *int `json:"deviceVersion,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// IsCompliant undocumented
	IsCompliant *bool `json:"isCompliant,omitempty"`
	// IsManaged undocumented
	IsManaged *bool `json:"isManaged,omitempty"`
	// MdmAppID undocumented
	MdmAppID *string `json:"mdmAppId,omitempty"`
	// OnPremisesLastSyncDateTime undocumented
	OnPremisesLastSyncDateTime *time.Time `json:"onPremisesLastSyncDateTime,omitempty"`
	// OnPremisesSyncEnabled undocumented
	OnPremisesSyncEnabled *bool `json:"onPremisesSyncEnabled,omitempty"`
	// OperatingSystem undocumented
	OperatingSystem *string `json:"operatingSystem,omitempty"`
	// OperatingSystemVersion undocumented
	OperatingSystemVersion *string `json:"operatingSystemVersion,omitempty"`
	// PhysicalIDs undocumented
	PhysicalIDs []string `json:"physicalIds,omitempty"`
	// ProfileType undocumented
	ProfileType *string `json:"profileType,omitempty"`
	// SystemLabels undocumented
	SystemLabels []string `json:"systemLabels,omitempty"`
	// TrustType undocumented
	TrustType *string `json:"trustType,omitempty"`
	// MemberOf undocumented
	MemberOf []DirectoryObject `json:"memberOf,omitempty"`
	// RegisteredOwners undocumented
	RegisteredOwners []DirectoryObject `json:"registeredOwners,omitempty"`
	// RegisteredUsers undocumented
	RegisteredUsers []DirectoryObject `json:"registeredUsers,omitempty"`
	// TransitiveMemberOf undocumented
	TransitiveMemberOf []DirectoryObject `json:"transitiveMemberOf,omitempty"`
	// Extensions undocumented
	Extensions []Extension `json:"extensions,omitempty"`
}
