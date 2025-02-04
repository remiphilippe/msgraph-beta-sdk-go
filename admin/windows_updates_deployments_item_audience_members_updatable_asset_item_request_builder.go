package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c "github.com/microsoftgraph/msgraph-beta-sdk-go/models/windowsupdates"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilder provides operations to manage the members property of the microsoft.graph.windowsUpdates.deploymentAudience entity.
type WindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// WindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// WindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilderGetQueryParameters specifies the assets to include in the audience.
type WindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// WindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilderGetQueryParameters
}
// WindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AddMembers provides operations to call the addMembers method.
func (m *WindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilder) AddMembers()(*WindowsUpdatesDeploymentsItemAudienceMembersItemAddMembersRequestBuilder) {
    return NewWindowsUpdatesDeploymentsItemAudienceMembersItemAddMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AddMembersById provides operations to call the addMembersById method.
func (m *WindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilder) AddMembersById()(*WindowsUpdatesDeploymentsItemAudienceMembersItemAddMembersByIdRequestBuilder) {
    return NewWindowsUpdatesDeploymentsItemAudienceMembersItemAddMembersByIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewWindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilderInternal instantiates a new UpdatableAssetItemRequestBuilder and sets the default values.
func NewWindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilder) {
    m := &WindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/admin/windows/updates/deployments/{deployment%2Did}/audience/members/{updatableAsset%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewWindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilder instantiates a new UpdatableAssetItemRequestBuilder and sets the default values.
func NewWindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property members for admin
func (m *WindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *WindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get specifies the assets to include in the audience.
func (m *WindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilder) Get(ctx context.Context, requestConfiguration *WindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilderGetRequestConfiguration)(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.CreateUpdatableAssetFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable), nil
}
// Patch update the navigation property members in admin
func (m *WindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilder) Patch(ctx context.Context, body i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable, requestConfiguration *WindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilderPatchRequestConfiguration)(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.CreateUpdatableAssetFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable), nil
}
// RemoveMembers provides operations to call the removeMembers method.
func (m *WindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilder) RemoveMembers()(*WindowsUpdatesDeploymentsItemAudienceMembersItemRemoveMembersRequestBuilder) {
    return NewWindowsUpdatesDeploymentsItemAudienceMembersItemRemoveMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RemoveMembersById provides operations to call the removeMembersById method.
func (m *WindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilder) RemoveMembersById()(*WindowsUpdatesDeploymentsItemAudienceMembersItemRemoveMembersByIdRequestBuilder) {
    return NewWindowsUpdatesDeploymentsItemAudienceMembersItemRemoveMembersByIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ToDeleteRequestInformation delete navigation property members for admin
func (m *WindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *WindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation specifies the assets to include in the audience.
func (m *WindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property members in admin
func (m *WindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable, requestConfiguration *WindowsUpdatesDeploymentsItemAudienceMembersUpdatableAssetItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
