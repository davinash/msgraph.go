// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// EmailFileAssessmentRequestObject undocumented
type EmailFileAssessmentRequestObject struct {
	// ThreatAssessmentRequestObject is the base model of EmailFileAssessmentRequestObject
	ThreatAssessmentRequestObject
	// RecipientEmail undocumented
	RecipientEmail *string `json:"recipientEmail,omitempty"`
	// DestinationRoutingReason undocumented
	DestinationRoutingReason *MailDestinationRoutingReason `json:"destinationRoutingReason,omitempty"`
	// ContentData undocumented
	ContentData *string `json:"contentData,omitempty"`
}
