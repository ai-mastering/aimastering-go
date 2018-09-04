# \AudioApi

All URIs are relative to *https://aimastering.com:443/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAudio**](AudioApi.md#CreateAudio) | **Post** /audios | Create a new audio.
[**DownloadAudio**](AudioApi.md#DownloadAudio) | **Get** /audios/{id}/download | Download an audio data by id.
[**DownloadAudioByToken**](AudioApi.md#DownloadAudioByToken) | **Get** /audios/download_by_token | Download an audio data by audio_download_token.
[**GetAudio**](AudioApi.md#GetAudio) | **Get** /audios/{id} | Get an audio by id.
[**GetAudioAnalysis**](AudioApi.md#GetAudioAnalysis) | **Get** /audios/{id}/analysis | Get an audio analysis by id.
[**GetAudioDownloadToken**](AudioApi.md#GetAudioDownloadToken) | **Get** /audios/{id}/download_token | Get an audio download token by id.
[**ListAudios**](AudioApi.md#ListAudios) | **Get** /audios | Get all audios accessable.


# **CreateAudio**
> Audio CreateAudio(ctx, optional)
Create a new audio.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File**| The file to upload. | 

### Return type

[**Audio**](Audio.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DownloadAudio**
> string DownloadAudio(ctx, id)
Download an audio data by id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Audio id | 

### Return type

**string**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DownloadAudioByToken**
> string DownloadAudioByToken(ctx, downloadToken)
Download an audio data by audio_download_token.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **downloadToken** | **string**| Audio download token | 

### Return type

**string**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAudio**
> Audio GetAudio(ctx, id)
Get an audio by id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Audio id | 

### Return type

[**Audio**](Audio.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAudioAnalysis**
> AudioAnalysis GetAudioAnalysis(ctx, id)
Get an audio analysis by id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Audio id | 

### Return type

[**AudioAnalysis**](AudioAnalysis.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAudioDownloadToken**
> AudioDownloadToken GetAudioDownloadToken(ctx, id)
Get an audio download token by id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Audio id | 

### Return type

[**AudioDownloadToken**](AudioDownloadToken.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAudios**
> []Audio ListAudios(ctx, )
Get all audios accessable.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Audio**](Audio.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

