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

type Audio struct {

	Id int32 `json:"id,omitempty"`

	FileResourceId int32 `json:"file_resource_id,omitempty"`

	UserId int32 `json:"user_id,omitempty"`

	Name string `json:"name,omitempty"`

	CreatedByUser bool `json:"created_by_user,omitempty"`

	Status string `json:"status,omitempty"`

	FailureReason string `json:"failure_reason,omitempty"`

	ProbeJson string `json:"probe_json,omitempty"`

	Rms float32 `json:"rms,omitempty"`

	Peak float32 `json:"peak,omitempty"`

	TruePeak float32 `json:"true_peak,omitempty"`

	LowpassTruePeak15khz float32 `json:"lowpass_true_peak_15khz,omitempty"`

	Loudness float32 `json:"loudness,omitempty"`

	Dynamics float32 `json:"dynamics,omitempty"`

	Sharpness float32 `json:"sharpness,omitempty"`

	Space float32 `json:"space,omitempty"`

	LoudnessRange float32 `json:"loudness_range,omitempty"`

	Drr float32 `json:"drr,omitempty"`

	SoundQuality float32 `json:"sound_quality,omitempty"`

	SoundQuality2 float32 `json:"sound_quality2,omitempty"`

	Dissonance float32 `json:"dissonance,omitempty"`

	Frames int32 `json:"frames,omitempty"`

	SampleRate int32 `json:"sample_rate,omitempty"`

	Channels int32 `json:"channels,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`

	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
