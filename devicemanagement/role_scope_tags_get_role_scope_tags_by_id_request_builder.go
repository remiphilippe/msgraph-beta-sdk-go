package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// RoleScopeTagsGetRoleScopeTagsByIdRequestBuilder provides operations to call the getRoleScopeTagsById method.
type RoleScopeTagsGetRoleScopeTagsByIdRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RoleScopeTagsGetRoleScopeTagsByIdRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleScopeTagsGetRoleScopeTagsByIdRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewRoleScopeTagsGetRoleScopeTagsByIdRequestBuilderInternal instantiates a new GetRoleScopeTagsByIdRequestBuilder and sets the default values.
func NewRoleScopeTagsGetRoleScopeTagsByIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleScopeTagsGetRoleScopeTagsByIdRequestBuilder) {
    m := &RoleScopeTagsGetRoleScopeTagsByIdRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/roleScopeTags/getRoleScopeTagsById", pathParameters),
    }
    return m
}
// NewRoleScopeTagsGetRoleScopeTagsByIdRequestBuilder instantiates a new GetRoleScopeTagsByIdRequestBuilder and sets the default values.
func NewRoleScopeTagsGetRoleScopeTagsByIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleScopeTagsGetRoleScopeTagsByIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRoleScopeTagsGetRoleScopeTagsByIdRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action getRoleScopeTagsById
func (m *RoleScopeTagsGetRoleScopeTagsByIdRequestBuilder) Post(ctx context.Context, body RoleScopeTagsGetRoleScopeTagsByIdPostRequestBodyable, requestConfiguration *RoleScopeTagsGetRoleScopeTagsByIdRequestBuilderPostRequestConfiguration)(RoleScopeTagsGetRoleScopeTagsByIdResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateRoleScopeTagsGetRoleScopeTagsByIdResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(RoleScopeTagsGetRoleScopeTagsByIdResponseable), nil
}
// ToPostRequestInformation invoke action getRoleScopeTagsById
func (m *RoleScopeTagsGetRoleScopeTagsByIdRequestBuilder) ToPostRequestInformation(ctx context.Context, body RoleScopeTagsGetRoleScopeTagsByIdPostRequestBodyable, requestConfiguration *RoleScopeTagsGetRoleScopeTagsByIdRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
