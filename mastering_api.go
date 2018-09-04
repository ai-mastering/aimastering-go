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
	"io/ioutil"
	"net/url"
	"net/http"
	"strings"
	"golang.org/x/net/context"
	"encoding/json"
	"fmt"
	"log"
)

// Linger please
var (
	_ context.Context
)

type MasteringApiService service


/* MasteringApiService Cancel a mastering by id.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param id Mastering id
 @return Mastering*/
func (a *MasteringApiService) CancelMastering(ctx context.Context, id int32) (Mastering,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Mastering
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/masterings/{id}/cancel"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if id < 1 {
		return successPayload, nil, reportError("id must be greater than 1")
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* MasteringApiService Create a new mastering.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param inputAudioId Input audio id
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "mode" (string) Mode
     @param "bassPreservation" (bool) This parameter represents if the bass preservation is enabled.
     @param "mastering" (bool) This parameter represents if the mastering is enabled. This parameter is effective only when the mode is \&quot;default\&quot; or \&quot;custom\&quot;.
     @param "preset" (string) This parameter is effective only when the mode is \&quot;preset\&quot;.
     @param "targetLoudness" (float32) This parameter represents the target loudness of the output audio in dB. This parameter is effective only when the mode is \&quot;custom\&quot;.
     @param "masteringMatchingLevel" (float32) This parameter represents the mastering reference matching level. This parameter is effective only when the mode is \&quot;custom\&quot; and the mastering is enabled.
     @param "masteringReverb" (bool) This parameter represents if the mastering reverb is enabled. This parameter is effective only when the mode is \&quot;custom\&quot; and the mastering is enabled.
     @param "masteringReverbGain" (float32) This parameter represents the mastering reverb gain relative to the dry sound in dB. This parameter is effective only when the mode is \&quot;custom\&quot; and the mastering is \&quot;true\&quot; and the mastering_reverb is \&quot;true\&quot;.
     @param "referenceAudioId" (int32) Reference audio id. This parameter is effective only when the mode is \&quot;custom\&quot; and the mastering is enabled.
     @param "lowCutFreq" (float32) This parameter represents the low cut freq  of the output audio in Hz. This parameter is effective only when the mode is \&quot;custom\&quot;.
     @param "highCutFreq" (float32) This parameter represents the target loudness of the output audio in Hz. This parameter is effective only when the mode is \&quot;custom\&quot;.
     @param "sampleRate" (int32) This parameter represents the sample rate of the output audio in dB. This parameter is effective only when the mode is \&quot;custom\&quot;.
     @param "bitDepth" (int32) This parameter represents the bit depth of the output audio in dB. This parameter is effective only when the mode is \&quot;custom\&quot;.
     @param "outputFormat" (string) This parameter represents the format of the output audio. This parameter is effective only when the mode is \&quot;custom\&quot;.
     @param "forPreview" (bool) If this is true, the mastering is treated for preview purpose (ex. not purchasable, not publishable, short lifetime).
     @param "startAt" (float32) Partial mastering start at.
     @param "endAt" (float32) Partial mastering end at.
     @param "isBakuage" (bool) Deprecated. For backward compatibility.
 @return Mastering*/
func (a *MasteringApiService) CreateMastering(ctx context.Context, inputAudioId int32, localVarOptionals map[string]interface{}) (Mastering,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Mastering
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/masterings"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if inputAudioId < 1 {
		return successPayload, nil, reportError("inputAudioId must be greater than 1")
	}
	if err := typeCheckParameter(localVarOptionals["mode"], "string", "mode"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["bassPreservation"], "bool", "bassPreservation"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["mastering"], "bool", "mastering"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["preset"], "string", "preset"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["targetLoudness"], "float32", "targetLoudness"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["masteringMatchingLevel"], "float32", "masteringMatchingLevel"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["masteringReverb"], "bool", "masteringReverb"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["masteringReverbGain"], "float32", "masteringReverbGain"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["referenceAudioId"], "int32", "referenceAudioId"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["lowCutFreq"], "float32", "lowCutFreq"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["highCutFreq"], "float32", "highCutFreq"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["sampleRate"], "int32", "sampleRate"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["bitDepth"], "int32", "bitDepth"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["outputFormat"], "string", "outputFormat"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["forPreview"], "bool", "forPreview"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["startAt"], "float32", "startAt"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["endAt"], "float32", "endAt"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["isBakuage"], "bool", "isBakuage"); err != nil {
		return successPayload, nil, err
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "multipart/form-data",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarTempParam, localVarOk := localVarOptionals["mode"].(string); localVarOk {
		log.Println(localVarTempParam)
		log.Println(parameterToString(localVarTempParam, ""))
		localVarFormParams.Add("mode", parameterToString(localVarTempParam, ""))
	}
	localVarFormParams.Add("input_audio_id", parameterToString(inputAudioId, ""))
	if localVarTempParam, localVarOk := localVarOptionals["bassPreservation"].(bool); localVarOk {
		localVarFormParams.Add("bass_preservation", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["mastering"].(bool); localVarOk {
		localVarFormParams.Add("mastering", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["preset"].(string); localVarOk {
		localVarFormParams.Add("preset", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["targetLoudness"].(float32); localVarOk {
		localVarFormParams.Add("target_loudness", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["masteringMatchingLevel"].(float32); localVarOk {
		localVarFormParams.Add("mastering_matching_level", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["masteringReverb"].(bool); localVarOk {
		localVarFormParams.Add("mastering_reverb", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["masteringReverbGain"].(float32); localVarOk {
		localVarFormParams.Add("mastering_reverb_gain", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["referenceAudioId"].(int32); localVarOk {
		localVarFormParams.Add("reference_audio_id", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["lowCutFreq"].(float32); localVarOk {
		localVarFormParams.Add("low_cut_freq", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["highCutFreq"].(float32); localVarOk {
		localVarFormParams.Add("high_cut_freq", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["sampleRate"].(int32); localVarOk {
		localVarFormParams.Add("sample_rate", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["bitDepth"].(int32); localVarOk {
		localVarFormParams.Add("bit_depth", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["outputFormat"].(string); localVarOk {
		localVarFormParams.Add("output_format", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["forPreview"].(bool); localVarOk {
		localVarFormParams.Add("for_preview", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["startAt"].(float32); localVarOk {
		localVarFormParams.Add("start_at", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["endAt"].(float32); localVarOk {
		localVarFormParams.Add("end_at", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["isBakuage"].(bool); localVarOk {
		localVarFormParams.Add("is_bakuage", parameterToString(localVarTempParam, ""))
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}
	log.Println(r)

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* MasteringApiService Delete mastering.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param id Mastering id
 @return Mastering*/
func (a *MasteringApiService) DeleteMastering(ctx context.Context, id int32) (Mastering,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Mastering
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/masterings/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if id < 1 {
		return successPayload, nil, reportError("id must be greater than 1")
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* MasteringApiService Get a mastering by id.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param id Mastering id
 @return Mastering*/
func (a *MasteringApiService) GetMastering(ctx context.Context, id int32) (Mastering,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Mastering
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/masterings/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if id < 1 {
		return successPayload, nil, reportError("id must be greater than 1")
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* MasteringApiService Review a mastering by id.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param id Mastering id
 @return Mastering*/
func (a *MasteringApiService) GetMasteringUnlockProduct(ctx context.Context, id int32) (Mastering,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Mastering
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/masterings/{id}/unlock_product"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if id < 1 {
		return successPayload, nil, reportError("id must be greater than 1")
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* MasteringApiService Get all accessable masterings.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @return []Mastering*/
func (a *MasteringApiService) ListMasterings(ctx context.Context) ([]Mastering,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []Mastering
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/masterings"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* MasteringApiService Publish a mastering by id.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param id Mastering id
 @param accessToken This parameter represents if the access token of the publishment service API.
 @param message This parameter represents the publishment message.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "service" (string) This parameter represents the publishment service.
     @param "accessTokenSecret" (string) This parameter represents the access token secret of the publishment service API. This parameter is effective only when the service is \&quot;twitter\&quot;.
 @return Mastering*/
func (a *MasteringApiService) PublishMastering(ctx context.Context, id int32, accessToken string, message string, localVarOptionals map[string]interface{}) (Mastering,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Mastering
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/masterings/{id}/publish"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if id < 1 {
		return successPayload, nil, reportError("id must be greater than 1")
	}
	if err := typeCheckParameter(localVarOptionals["service"], "string", "service"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["accessTokenSecret"], "string", "accessTokenSecret"); err != nil {
		return successPayload, nil, err
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "multipart/form-data",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarTempParam, localVarOk := localVarOptionals["service"].(string); localVarOk {
		localVarFormParams.Add("service", parameterToString(localVarTempParam, ""))
	}
	localVarFormParams.Add("access_token", parameterToString(accessToken, ""))
	if localVarTempParam, localVarOk := localVarOptionals["accessTokenSecret"].(string); localVarOk {
		localVarFormParams.Add("access_token_secret", parameterToString(localVarTempParam, ""))
	}
	localVarFormParams.Add("message", parameterToString(message, ""))
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* MasteringApiService Review a mastering by id.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param id Mastering id
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "reviewComment" (string) This parameter represents the review comment.
     @param "reviewScore" (float32) This parameter represents the review score.
 @return Mastering*/
func (a *MasteringApiService) ReviewMastering(ctx context.Context, id int32, localVarOptionals map[string]interface{}) (Mastering,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Mastering
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/masterings/{id}/review"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if id < 1 {
		return successPayload, nil, reportError("id must be greater than 1")
	}
	if err := typeCheckParameter(localVarOptionals["reviewComment"], "string", "reviewComment"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["reviewScore"], "float32", "reviewScore"); err != nil {
		return successPayload, nil, err
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "multipart/form-data",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarTempParam, localVarOk := localVarOptionals["reviewComment"].(string); localVarOk {
		localVarFormParams.Add("review_comment", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["reviewScore"].(float32); localVarOk {
		localVarFormParams.Add("review_score", parameterToString(localVarTempParam, ""))
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}
