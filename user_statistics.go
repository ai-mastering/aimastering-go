/*
 * AI Mastering API
 *
 * This is a AI Mastering API document. You can use the mastering feature of [AI Mastering](https://aimastering.com) through this API.
 *
 * API version: 1.0.0
 * Contact: info@bakuage.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package aimastering

type UserStatistics struct {

	RegistrationInvitationCount int32 `json:"registration_invitation_count,omitempty"`

	SubscriptionInvitationCount int32 `json:"subscription_invitation_count,omitempty"`

	MasteringCount int32 `json:"mastering_count,omitempty"`

	MonthlyRegistrationInvitationCount int32 `json:"monthly_registration_invitation_count,omitempty"`

	MonthlySubscriptionInvitationCount int32 `json:"monthly_subscription_invitation_count,omitempty"`

	MonthlyMasteringCount int32 `json:"monthly_mastering_count,omitempty"`
}
