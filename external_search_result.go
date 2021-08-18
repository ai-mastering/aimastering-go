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

type ExternalSearchResult struct {

	Itunes []ExternalSearchResultItunes `json:"itunes,omitempty"`

	Youtube []ExternalSearchResultYoutube `json:"youtube,omitempty"`
}
