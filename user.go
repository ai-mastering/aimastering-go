/*
 * AI Mastering API
 *
 * This is a AI Mastering API document. You can use the mastering feature of [AI Mastering](https://aimastering.com) through this API.
 *
 * API version: 1.0.0
 * Contact: aimasteringcom@gmail.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package aimastering

import (
	"time"
)

type User struct {

	Id int32 `json:"id,omitempty"`

	AgreedTermsOfService bool `json:"agreed_terms_of_service,omitempty"`

	AuthId string `json:"auth_id,omitempty"`

	AuthProvider string `json:"auth_provider,omitempty"`

	Email string `json:"email,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`

	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
