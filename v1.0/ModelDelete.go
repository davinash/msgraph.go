// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// DeleteUserFromSharedAppleDeviceActionResult undocumented
type DeleteUserFromSharedAppleDeviceActionResult struct {
	// DeviceActionResult is the base model of DeleteUserFromSharedAppleDeviceActionResult
	DeviceActionResult
	// UserPrincipalName User principal name of the user to be deleted
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
}
