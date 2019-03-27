# \MasteringApi

All URIs are relative to *https://bakuage.com:443/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelMastering**](MasteringApi.md#CancelMastering) | **Put** /masterings/{id}/cancel | Cancel a mastering by id.
[**CreateMastering**](MasteringApi.md#CreateMastering) | **Post** /masterings | Create a new mastering.
[**DeleteMastering**](MasteringApi.md#DeleteMastering) | **Delete** /masterings/{id} | Delete mastering.
[**FreeUnlockMastering**](MasteringApi.md#FreeUnlockMastering) | **Put** /masterings/{id}/free_unlock | Free unlock a mastering by id.
[**GetMastering**](MasteringApi.md#GetMastering) | **Get** /masterings/{id} | Get a mastering by id.
[**GetMasteringUnlockProduct**](MasteringApi.md#GetMasteringUnlockProduct) | **Get** /masterings/{id}/unlock_product | Review a mastering by id.
[**ListMasterings**](MasteringApi.md#ListMasterings) | **Get** /masterings | Get all accessable masterings.
[**PublishMastering**](MasteringApi.md#PublishMastering) | **Post** /masterings/{id}/publish | Publish a mastering by id.
[**ReviewMastering**](MasteringApi.md#ReviewMastering) | **Put** /masterings/{id}/review | Review a mastering by id.
[**UpdateMastering**](MasteringApi.md#UpdateMastering) | **Put** /masterings/{id} | Update a mastering.


# **CancelMastering**
> Mastering CancelMastering(ctx, id)
Cancel a mastering by id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Mastering id | 

### Return type

[**Mastering**](Mastering.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMastering**
> Mastering CreateMastering(ctx, inputAudioId, optional)
Create a new mastering.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **inputAudioId** | **int32**| Input audio id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputAudioId** | **int32**| Input audio id | 
 **mode** | **string**| Mode | [default to default]
 **bassPreservation** | **bool**| This parameter represents if the bass preservation is enabled. | [default to false]
 **mastering** | **bool**| This parameter represents if the mastering is enabled. This parameter is effective only when the mode is \&quot;default\&quot; or \&quot;custom\&quot;. | [default to false]
 **masteringAlgorithm** | **string**|  | [default to v2]
 **noiseReduction** | **bool**| This parameter represents if the nosie reduction is enabled. This parameter is effective only when the mode is \&quot;custom\&quot;. | [default to false]
 **preset** | **string**| This parameter is effective only when the mode is \&quot;custom\&quot;. | [default to general]
 **targetLoudness** | **float32**| This parameter represents the target loudness of the output audio in dB. This parameter is effective only when the mode is \&quot;custom\&quot;. | [default to -5]
 **targetLoudnessMode** | **string**|  | [default to loudness]
 **masteringMatchingLevel** | **float32**| This parameter represents the mastering reference matching level. This parameter is effective only when the mode is \&quot;custom\&quot; and the mastering is enabled. | [default to 0.5]
 **masteringReverb** | **bool**| This parameter represents if the mastering reverb is enabled. This parameter is effective only when the mode is \&quot;custom\&quot; and the mastering is enabled. | [default to false]
 **masteringReverbGain** | **float32**| This parameter represents the mastering reverb gain relative to the dry sound in dB. This parameter is effective only when the mode is \&quot;custom\&quot; and the mastering is \&quot;true\&quot; and the mastering_reverb is \&quot;true\&quot;. | [default to -36]
 **referenceAudioId** | **int32**| Reference audio id. This parameter is effective only when the mode is \&quot;custom\&quot; and the mastering is enabled. | 
 **lowCutFreq** | **float32**| This parameter represents the low cut freq  of the output audio in Hz. This parameter is effective only when the mode is \&quot;custom\&quot;. | [default to 20]
 **highCutFreq** | **float32**| This parameter represents the high cut freq of the output audio in Hz. This parameter is effective only when the mode is \&quot;custom\&quot;. | [default to 20000]
 **ceiling** | **float32**|  | [default to 0]
 **ceilingMode** | **string**|  | [default to 0]
 **oversample** | **float32**|  | [default to 1]
 **sampleRate** | **int32**| This parameter represents the sample rate of the output audio in dB. This parameter is effective only when the mode is \&quot;custom\&quot;. | [default to 44100]
 **bitDepth** | **int32**| This parameter represents the bit depth of the output audio in dB. This parameter is effective only when the mode is \&quot;custom\&quot;. | [default to 16]
 **outputFormat** | **string**| This parameter represents the format of the output audio. This parameter is effective only when the mode is \&quot;custom\&quot;. | [default to wav]
 **forPreview** | **bool**| If this is true, the mastering is treated for preview purpose (ex. not purchasable, not publishable, short lifetime).  | [default to false]
 **startAt** | **float32**| Partial mastering start at.  | [default to 0]
 **endAt** | **float32**| Partial mastering end at.  | [default to -1]
 **videoTitle** | **string**| This parameter represents the title of output video. | [default to ]

### Return type

[**Mastering**](Mastering.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMastering**
> Mastering DeleteMastering(ctx, id)
Delete mastering.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Mastering id | 

### Return type

[**Mastering**](Mastering.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FreeUnlockMastering**
> Mastering FreeUnlockMastering(ctx, id)
Free unlock a mastering by id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Mastering id | 

### Return type

[**Mastering**](Mastering.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMastering**
> Mastering GetMastering(ctx, id)
Get a mastering by id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Mastering id | 

### Return type

[**Mastering**](Mastering.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMasteringUnlockProduct**
> Mastering GetMasteringUnlockProduct(ctx, id)
Review a mastering by id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Mastering id | 

### Return type

[**Mastering**](Mastering.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMasterings**
> []Mastering ListMasterings(ctx, )
Get all accessable masterings.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Mastering**](Mastering.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PublishMastering**
> Mastering PublishMastering(ctx, id, accessToken, message, optional)
Publish a mastering by id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Mastering id | 
  **accessToken** | **string**| This parameter represents if the access token of the publishment service API. | 
  **message** | **string**| This parameter represents the publishment message. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| Mastering id | 
 **accessToken** | **string**| This parameter represents if the access token of the publishment service API. | 
 **message** | **string**| This parameter represents the publishment message. | 
 **service** | **string**| This parameter represents the publishment service. | 
 **accessTokenSecret** | **string**| This parameter represents the access token secret of the publishment service API. This parameter is effective only when the service is \&quot;twitter\&quot;. | 

### Return type

[**Mastering**](Mastering.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReviewMastering**
> Mastering ReviewMastering(ctx, id, optional)
Review a mastering by id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Mastering id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| Mastering id | 
 **reviewComment** | **string**| This parameter represents the review comment. | 
 **reviewScore** | **float32**| This parameter represents the review score. | 

### Return type

[**Mastering**](Mastering.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMastering**
> Mastering UpdateMastering(ctx, id, optional)
Update a mastering.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Mastering id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| Mastering id | 
 **preserved** | **bool**| Disable auto delete. | 

### Return type

[**Mastering**](Mastering.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

