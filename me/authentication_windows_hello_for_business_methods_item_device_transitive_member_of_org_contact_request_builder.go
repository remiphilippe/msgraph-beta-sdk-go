package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfOrgContactRequestBuilder casts the previous resource to orgContact.
type AuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfOrgContactRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfOrgContactRequestBuilderGetQueryParameters get the items of type microsoft.graph.orgContact in the microsoft.graph.directoryObject collection
type AuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfOrgContactRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// AuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfOrgContactRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfOrgContactRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfOrgContactRequestBuilderGetQueryParameters
}
// NewAuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfOrgContactRequestBuilderInternal instantiates a new OrgContactRequestBuilder and sets the default values.
func NewAuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfOrgContactRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfOrgContactRequestBuilder) {
    m := &AuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfOrgContactRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/authentication/windowsHelloForBusinessMethods/{windowsHelloForBusinessAuthenticationMethod%2Did}/device/transitiveMemberOf/microsoft.graph.orgContact{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfOrgContactRequestBuilder instantiates a new OrgContactRequestBuilder and sets the default values.
func NewAuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfOrgContactRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfOrgContactRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfOrgContactRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *AuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfOrgContactRequestBuilder) Count()(*AuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfOrgContactCountRequestBuilder) {
    return NewAuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfOrgContactCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get the items of type microsoft.graph.orgContact in the microsoft.graph.directoryObject collection
func (m *AuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfOrgContactRequestBuilder) Get(ctx context.Context, requestConfiguration *AuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfOrgContactRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OrgContactCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOrgContactCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OrgContactCollectionResponseable), nil
}
// ToGetRequestInformation get the items of type microsoft.graph.orgContact in the microsoft.graph.directoryObject collection
func (m *AuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfOrgContactRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AuthenticationWindowsHelloForBusinessMethodsItemDeviceTransitiveMemberOfOrgContactRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
