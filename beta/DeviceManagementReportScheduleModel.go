// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// DeviceManagementReportSchedule Entity representing a schedule for which reports are delivered
type DeviceManagementReportSchedule struct {
	// Entity is the base model of DeviceManagementReportSchedule
	Entity
	// ReportScheduleName Name of the schedule
	ReportScheduleName *string `json:"reportScheduleName,omitempty"`
	// Subject Subject of the scheduled reports that are delivered
	Subject *string `json:"subject,omitempty"`
	// Emails Emails to which the scheduled reports are delivered
	Emails []string `json:"emails,omitempty"`
	// Recurrence Frequency of scheduled report delivery
	Recurrence *DeviceManagementScheduledReportRecurrence `json:"recurrence,omitempty"`
	// StartDateTime Time that the delivery of the scheduled reports starts
	StartDateTime *time.Time `json:"startDateTime,omitempty"`
	// EndDateTime Time that the delivery of the scheduled reports ends
	EndDateTime *time.Time `json:"endDateTime,omitempty"`
	// UserID The Id of the User who created the report
	UserID *string `json:"userId,omitempty"`
	// ReportName Name of the report
	ReportName *string `json:"reportName,omitempty"`
	// Filter Filters applied on the report
	Filter *string `json:"filter,omitempty"`
	// Select Columns selected from the report
	Select []string `json:"select,omitempty"`
	// OrderBy Ordering of columns in the report
	OrderBy []string `json:"orderBy,omitempty"`
	// Format Format of the scheduled report
	Format *DeviceManagementReportFileFormat `json:"format,omitempty"`
}
