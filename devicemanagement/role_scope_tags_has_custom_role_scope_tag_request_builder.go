package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// RoleScopeTagsHasCustomRoleScopeTagRequestBuilder provides operations to call the hasCustomRoleScopeTag method.
type RoleScopeTagsHasCustomRoleScopeTagRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// RoleScopeTagsHasCustomRoleScopeTagRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleScopeTagsHasCustomRoleScopeTagRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewRoleScopeTagsHasCustomRoleScopeTagRequestBuilderInternal instantiates a new HasCustomRoleScopeTagRequestBuilder and sets the default values.
func NewRoleScopeTagsHasCustomRoleScopeTagRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleScopeTagsHasCustomRoleScopeTagRequestBuilder) {
    m := &RoleScopeTagsHasCustomRoleScopeTagRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/roleScopeTags/microsoft.graph.hasCustomRoleScopeTag()";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRoleScopeTagsHasCustomRoleScopeTagRequestBuilder instantiates a new HasCustomRoleScopeTagRequestBuilder and sets the default values.
func NewRoleScopeTagsHasCustomRoleScopeTagRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleScopeTagsHasCustomRoleScopeTagRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRoleScopeTagsHasCustomRoleScopeTagRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function hasCustomRoleScopeTag
func (m *RoleScopeTagsHasCustomRoleScopeTagRequestBuilder) Get(ctx context.Context, requestConfiguration *RoleScopeTagsHasCustomRoleScopeTagRequestBuilderGetRequestConfiguration)(RoleScopeTagsHasCustomRoleScopeTagResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateRoleScopeTagsHasCustomRoleScopeTagResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(RoleScopeTagsHasCustomRoleScopeTagResponseable), nil
}
// ToGetRequestInformation invoke function hasCustomRoleScopeTag
func (m *RoleScopeTagsHasCustomRoleScopeTagRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *RoleScopeTagsHasCustomRoleScopeTagRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
