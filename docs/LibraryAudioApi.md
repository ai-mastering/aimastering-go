# \LibraryAudioApi

All URIs are relative to *https://api.bakuage.com:443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLibraryAudio**](LibraryAudioApi.md#CreateLibraryAudio) | **Post** /library_audios | Create a new library audio.
[**CreateLibraryAudioLike**](LibraryAudioApi.md#CreateLibraryAudioLike) | **Post** /library_audios/{id}/like | Create a new library audio like.
[**DeleteLibraryAudio**](LibraryAudioApi.md#DeleteLibraryAudio) | **Delete** /library_audios/{id} | Delete library audio.
[**GetLibraryAudio**](LibraryAudioApi.md#GetLibraryAudio) | **Get** /library_audios/{id} | Get a library audio by id.
[**GetLibraryAudioAnalysis**](LibraryAudioApi.md#GetLibraryAudioAnalysis) | **Get** /library_audios/{id}/analysis | Get a library audio analysis by id.
[**ListLibraryAudios**](LibraryAudioApi.md#ListLibraryAudios) | **Get** /library_audios | Get all library audios accessable.
[**UpdateLibraryAudio**](LibraryAudioApi.md#UpdateLibraryAudio) | **Put** /library_audios/{id} | Update library audio.


# **CreateLibraryAudio**
> LibraryAudio CreateLibraryAudio(ctx, optional)
Create a new library audio.

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

[**LibraryAudio**](LibraryAudio.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateLibraryAudioLike**
> LibraryAudioLike CreateLibraryAudioLike(ctx, id)
Create a new library audio like.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Library audio id | 

### Return type

[**LibraryAudioLike**](LibraryAudioLike.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLibraryAudio**
> LibraryAudio DeleteLibraryAudio(ctx, id)
Delete library audio.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Library audio id | 

### Return type

[**LibraryAudio**](LibraryAudio.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLibraryAudio**
> LibraryAudio GetLibraryAudio(ctx, id)
Get a library audio by id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Library audio id | 

### Return type

[**LibraryAudio**](LibraryAudio.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLibraryAudioAnalysis**
> LibraryAudioAnalysis GetLibraryAudioAnalysis(ctx, id)
Get a library audio analysis by id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Library audio id | 

### Return type

[**LibraryAudioAnalysis**](LibraryAudioAnalysis.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListLibraryAudios**
> []LibraryAudio ListLibraryAudios(ctx, )
Get all library audios accessable.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]LibraryAudio**](LibraryAudio.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateLibraryAudio**
> LibraryAudio UpdateLibraryAudio(ctx, id, optional)
Update library audio.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Library audio id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| Library audio id | 
 **isPublic** | **bool**| Whether the library audio is public. | 

### Return type

[**LibraryAudio**](LibraryAudio.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

