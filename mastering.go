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

type Mastering struct {

	Id int32 `json:"id,omitempty"`

	UserId int32 `json:"user_id,omitempty"`

	InputAudioId int32 `json:"input_audio_id,omitempty"`

	OutputAudioId int32 `json:"output_audio_id,omitempty"`

	OutputVideoId int32 `json:"output_video_id,omitempty"`

	ReferenceAudioId int32 `json:"reference_audio_id,omitempty"`

	Mode string `json:"mode,omitempty"`

	Status string `json:"status,omitempty"`

	FailureReason string `json:"failure_reason,omitempty"`

	TargetLoudness float32 `json:"target_loudness,omitempty"`

	OutputFormat string `json:"output_format,omitempty"`

	Preset string `json:"preset,omitempty"`

	BitDepth int32 `json:"bit_depth,omitempty"`

	SampleRate int32 `json:"sample_rate,omitempty"`

	ReviewComment string `json:"review_comment,omitempty"`

	ReviewScore float32 `json:"review_score,omitempty"`

	MasteringMatchingLevel float32 `json:"mastering_matching_level,omitempty"`

	Progression float32 `json:"progression,omitempty"`

	BassPreservation bool `json:"bass_preservation,omitempty"`

	Mastering bool `json:"mastering,omitempty"`

	RetryCount int32 `json:"retry_count,omitempty"`

	MasteringReverb bool `json:"mastering_reverb,omitempty"`

	MasteringReverbGain float32 `json:"mastering_reverb_gain,omitempty"`

	LowCutFreq float32 `json:"low_cut_freq,omitempty"`

	HighCutFreq float32 `json:"high_cut_freq,omitempty"`

	ExpireAt time.Time `json:"expire_at,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`

	UpdatedAt time.Time `json:"updated_at,omitempty"`
}