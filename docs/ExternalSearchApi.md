# \ExternalSearchApi

All URIs are relative to *https://aimastering.com:443/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchExternal**](ExternalSearchApi.md#SearchExternal) | **Get** /external_search | Search external music and get name, url, thumbnails, etc.


# **SearchExternal**
> ExternalSearchResult SearchExternal(ctx, query, country)
Search external music and get name, url, thumbnails, etc.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **query** | **string**| Search query | 
  **country** | **string**| Country ex. US, JP, etc | 

### Return type

[**ExternalSearchResult**](ExternalSearchResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

