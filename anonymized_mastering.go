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

type AnonymizedMastering struct {

	UserId string `json:"user_id,omitempty"`

	UserAuthProvider string `json:"user_auth_provider,omitempty"`

	Mode string `json:"mode,omitempty"`

	Status string `json:"status,omitempty"`

	FailureReason string `json:"failure_reason,omitempty"`

	TargetLoudness float32 `json:"target_loudness,omitempty"`

	OutputFormat string `json:"output_format,omitempty"`

	Preset string `json:"preset,omitempty"`

	BitDepth int32 `json:"bit_depth,omitempty"`

	SampleRate int32 `json:"sample_rate,omitempty"`

	ReviewScore float32 `json:"review_score,omitempty"`

	MasteringMatchingLevel float32 `json:"mastering_matching_level,omitempty"`

	Mastering bool `json:"mastering,omitempty"`

	Paid bool `json:"paid,omitempty"`

	PaymentService string `json:"payment_service,omitempty"`

	RetryCount int32 `json:"retry_count,omitempty"`

	MasteringReverb bool `json:"mastering_reverb,omitempty"`

	MasteringReverbGain float32 `json:"mastering_reverb_gain,omitempty"`

	LowCutFreq float32 `json:"low_cut_freq,omitempty"`

	HighCutFreq float32 `json:"high_cut_freq,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`

	UpdatedAt time.Time `json:"updated_at,omitempty"`
}