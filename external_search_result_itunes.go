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

type ExternalSearchResultItunes struct {

	Id string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	Description string `json:"description,omitempty"`

	Url string `json:"url,omitempty"`

	ThumbnailUrl string `json:"thumbnail_url,omitempty"`

	PreviewUrl string `json:"preview_url,omitempty"`

	AlbumName string `json:"album_name,omitempty"`

	AlbumUrl string `json:"album_url,omitempty"`

	ArtistName string `json:"artist_name,omitempty"`

	ArtistUrl string `json:"artist_url,omitempty"`

	TrackNumber int32 `json:"track_number,omitempty"`
}
