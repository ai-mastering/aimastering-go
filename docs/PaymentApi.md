# \PaymentApi

All URIs are relative to *https://bakuage.com:443/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePayment**](PaymentApi.md#CreatePayment) | **Post** /payments | Create a new payment.
[**ExecutePayment**](PaymentApi.md#ExecutePayment) | **Put** /payments/{id}/execute | Execute a payment by id.
[**GetPayment**](PaymentApi.md#GetPayment) | **Get** /payments/{id} | Get a payment by id.
[**ListPayments**](PaymentApi.md#ListPayments) | **Get** /payments | Get all accessable payments.


# **CreatePayment**
> Payment CreatePayment(ctx, productToken, service, optional)
Create a new payment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **productToken** | **string**| This parameter represents the product token. | 
  **service** | **string**| This parameter represents the payment message. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productToken** | **string**| This parameter represents the product token. | 
 **service** | **string**| This parameter represents the payment message. | 
 **token** | **string**| This parameter represents the card token. This parameter is effective only when the service is \&quot;stripe\&quot;. | 

### Return type

[**Payment**](Payment.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExecutePayment**
> Payment ExecutePayment(ctx, id, payerId)
Execute a payment by id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Payment id | 
  **payerId** | **string**| This parameter represents the card token. This parameter is effective only when the service is \&quot;paypal\&quot;. | 

### Return type

[**Payment**](Payment.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPayment**
> Payment GetPayment(ctx, id)
Get a payment by id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**| Payment id | 

### Return type

[**Payment**](Payment.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPayments**
> []Payment ListPayments(ctx, )
Get all accessable payments.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Payment**](Payment.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

