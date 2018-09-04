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

type LibraryAudioLike struct {

	// Audio analysis id
	Id int32 `json:"id,omitempty"`

	// Audio id
	LibraryAudioId int32 `json:"library_audio_id,omitempty"`

	// User id
	UserId int32 `json:"user_id,omitempty"`
}