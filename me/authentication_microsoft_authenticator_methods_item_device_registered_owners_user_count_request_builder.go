package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersUserCountRequestBuilder provides operations to count the resources in the collection.
type AuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersUserCountRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersUserCountRequestBuilderGetQueryParameters get the number of the resource
type AuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersUserCountRequestBuilderGetQueryParameters struct {
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
}
// AuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersUserCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersUserCountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersUserCountRequestBuilderGetQueryParameters
}
// NewAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersUserCountRequestBuilderInternal instantiates a new CountRequestBuilder and sets the default values.
func NewAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersUserCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersUserCountRequestBuilder) {
    m := &AuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersUserCountRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/authentication/microsoftAuthenticatorMethods/{microsoftAuthenticatorAuthenticationMethod%2Did}/device/registeredOwners/microsoft.graph.user/$count{?%24search,%24filter}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersUserCountRequestBuilder instantiates a new CountRequestBuilder and sets the default values.
func NewAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersUserCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersUserCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersUserCountRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the number of the resource
func (m *AuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersUserCountRequestBuilder) Get(ctx context.Context, requestConfiguration *AuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersUserCountRequestBuilderGetRequestConfiguration)(*int32, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendPrimitive(ctx, requestInfo, "int32", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*int32), nil
}
// ToGetRequestInformation get the number of the resource
func (m *AuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersUserCountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersUserCountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "text/plain")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
