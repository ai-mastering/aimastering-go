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

type Payment struct {

	Id int32 `json:"id,omitempty"`

	Service string `json:"service,omitempty"`

	ProductGiven bool `json:"product_given,omitempty"`

	Product *interface{} `json:"product,omitempty"`

	TransactionId string `json:"transaction_id,omitempty"`

	TransactionDetail *interface{} `json:"transaction_detail,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`
}