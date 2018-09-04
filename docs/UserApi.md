# \UserApi

All URIs are relative to *https://aimastering.com:443/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSelf**](UserApi.md#GetSelf) | **Get** /users/self | Get self user.
[**UpdateSelf**](UserApi.md#UpdateSelf) | **Put** /users/self | Update self user.


# **GetSelf**
> User GetSelf(ctx, )
Get self user.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**User**](User.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSelf**
> User UpdateSelf(ctx, optional)
Update self user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **agreedTermsOfService** | **bool**| Whether you agreed terms of service. | 
 **email** | **string**| The email. | 

### Return type

[**User**](User.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

