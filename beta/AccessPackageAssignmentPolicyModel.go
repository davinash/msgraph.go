// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// AccessPackageAssignmentPolicy undocumented
type AccessPackageAssignmentPolicy struct {
	// Entity is the base model of AccessPackageAssignmentPolicy
	Entity
	// UserType undocumented
	UserType *string `json:"userType,omitempty"`
	// AccessPackageID undocumented
	AccessPackageID *string `json:"accessPackageId,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// IsEnabled undocumented
	IsEnabled *bool `json:"isEnabled,omitempty"`
	// CanExtend undocumented
	CanExtend *bool `json:"canExtend,omitempty"`
	// DurationInDays undocumented
	DurationInDays *int `json:"durationInDays,omitempty"`
	// ExpirationDateTime undocumented
	ExpirationDateTime *time.Time `json:"expirationDateTime,omitempty"`
	// CreatedBy undocumented
	CreatedBy *string `json:"createdBy,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// ModifiedBy undocumented
	ModifiedBy *string `json:"modifiedBy,omitempty"`
	// ModifiedDateTime undocumented
	ModifiedDateTime *time.Time `json:"modifiedDateTime,omitempty"`
	// AccessPackage undocumented
	AccessPackage *AccessPackage `json:"accessPackage,omitempty"`
	// AccessPackageCatalog undocumented
	AccessPackageCatalog *AccessPackageCatalog `json:"accessPackageCatalog,omitempty"`
}
