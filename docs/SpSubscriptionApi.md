# \SpSubscriptionApi

All URIs are relative to *https://api.bakuage.com:443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSpSubscription**](SpSubscriptionApi.md#CreateSpSubscription) | **Post** /sp_subscriptions | Create a new smartphone subscription.
[**ListSpSubscriptions**](SpSubscriptionApi.md#ListSpSubscriptions) | **Get** /sp_subscriptions | Get all accessable smartphone subscriptions.


# **CreateSpSubscription**
> SpSubscription CreateSpSubscription(ctx, service, optional)
Create a new smartphone subscription.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **service** | **string**| Service. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **service** | **string**| Service. | 
 **receipt** | **string**| Base64 encoded app store receipt. This parameter is effective only when the service is \&quot;appstore\&quot;. | 

### Return type

[**SpSubscription**](SpSubscription.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSpSubscriptions**
> []SpSubscription ListSpSubscriptions(ctx, )
Get all accessable smartphone subscriptions.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]SpSubscription**](SpSubscription.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

