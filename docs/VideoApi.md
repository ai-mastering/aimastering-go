# \VideoApi

All URIs are relative to *https://api.bakuage.com:443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadVideo**](VideoApi.md#DownloadVideo) | **Get** /videos/{id}/download | Download an video data by id.
[**DownloadVideoByToken**](VideoApi.md#DownloadVideoByToken) | **Get** /videos/download_by_token | Download an video data by video_download_token.
[**GetVideo**](VideoApi.md#GetVideo) | **Get** /videos/{id} | Get an video by id.
[**GetVideoDownloadToken**](VideoApi.md#GetVideoDownloadToken) | **Get** /videos/{id}/download_token | Get an video download token by id.
[**ListVideos**](VideoApi.md#ListVideos) | **Get** /videos | Get all videos accessable.


# **DownloadVideo**
> string DownloadVideo(ctx, id)
Download an video data by id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Video id | 

### Return type

**string**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DownloadVideoByToken**
> string DownloadVideoByToken(ctx, downloadToken)
Download an video data by video_download_token.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **downloadToken** | **string**| Video download token | 

### Return type

**string**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVideo**
> Video GetVideo(ctx, id)
Get an video by id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Video id | 

### Return type

[**Video**](Video.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVideoDownloadToken**
> VideoDownloadToken GetVideoDownloadToken(ctx, id)
Get an video download token by id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Video id | 

### Return type

[**VideoDownloadToken**](VideoDownloadToken.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListVideos**
> []Video ListVideos(ctx, )
Get all videos accessable.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Video**](Video.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

