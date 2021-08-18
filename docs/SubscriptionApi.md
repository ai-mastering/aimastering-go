# \SubscriptionApi

All URIs are relative to *https://api.bakuage.com:443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelSubscription**](SubscriptionApi.md#CancelSubscription) | **Put** /subscriptions/{id}/cancel | Cancel a subscription by id.
[**CancelSubscriptionCancellation**](SubscriptionApi.md#CancelSubscriptionCancellation) | **Put** /subscriptions/{id}/cancel_cancellation | Cancel the subscription cancellation  by id.
[**CreateSubscription**](SubscriptionApi.md#CreateSubscription) | **Post** /subscriptions | Create a new subscription.
[**GetSubscription**](SubscriptionApi.md#GetSubscription) | **Get** /subscriptions/{id} | Get a subscription by id.
[**ListSubscriptions**](SubscriptionApi.md#ListSubscriptions) | **Get** /subscriptions | Get all accessable subscriptions.


# **CancelSubscription**
> Subscription CancelSubscription(ctx, id)
Cancel a subscription by id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Subscription id | 

### Return type

[**Subscription**](Subscription.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CancelSubscriptionCancellation**
> Subscription CancelSubscriptionCancellation(ctx, id)
Cancel the subscription cancellation  by id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Subscription id | 

### Return type

[**Subscription**](Subscription.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSubscription**
> Subscription CreateSubscription(ctx, service, optional)
Create a new subscription.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **service** | **string**| This parameter represents the payment message. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **service** | **string**| This parameter represents the payment message. | 
 **stripePlanId** | **string**| The Stripe plan id. This parameter is effective only when the service is \&quot;stripe\&quot;. | 
 **token** | **string**| This parameter represents the card token. This parameter is effective only when the service is \&quot;stripe\&quot;. | 
 **affiliateId** | **string**| Affiliate id of inviter user. | 

### Return type

[**Subscription**](Subscription.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSubscription**
> Subscription GetSubscription(ctx, id)
Get a subscription by id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Subscription id | 

### Return type

[**Subscription**](Subscription.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSubscriptions**
> []Subscription ListSubscriptions(ctx, )
Get all accessable subscriptions.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Subscription**](Subscription.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

