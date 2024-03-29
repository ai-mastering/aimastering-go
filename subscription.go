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

import (
	"time"
)

type Subscription struct {

	Id int32 `json:"id,omitempty"`

	UserId int32 `json:"user_id,omitempty"`

	StripeSubscriptionId string `json:"stripe_subscription_id,omitempty"`

	CurrentPeriodStartAt time.Time `json:"current_period_start_at,omitempty"`

	CurrentPeriodEndAt time.Time `json:"current_period_end_at,omitempty"`

	Canceled bool `json:"canceled,omitempty"`

	CancelAtPeriodEnd bool `json:"cancel_at_period_end,omitempty"`

	TrialEnd time.Time `json:"trial_end,omitempty"`

	Status string `json:"status,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`

	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
