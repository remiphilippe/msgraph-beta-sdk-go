package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/termstore"
    if37ca3748c081d626a200e3354fca5c0d75d6c98d76383d181b55ada257e45e1 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/termstore/groups/item/sets"
    i079fbfa6a17671510d081186b3a4d407748bf5f48e641d244923ea9ba9101107 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/sites/item/termstore/groups/item/sets/item"
)

// GroupItemRequestBuilder provides operations to manage the groups property of the microsoft.graph.termStore.store entity.
type GroupItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GroupItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// GroupItemRequestBuilderGetQueryParameters collection of all groups available in the term store.
type GroupItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// GroupItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GroupItemRequestBuilderGetQueryParameters
}
// GroupItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGroupItemRequestBuilderInternal instantiates a new GroupItemRequestBuilder and sets the default values.
func NewGroupItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupItemRequestBuilder) {
    m := &GroupItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/termStore/groups/{group%2Did1}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGroupItemRequestBuilder instantiates a new GroupItemRequestBuilder and sets the default values.
func NewGroupItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property groups for groups
func (m *GroupItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *GroupItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation collection of all groups available in the term store.
func (m *GroupItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *GroupItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property groups in groups
func (m *GroupItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.Groupable, requestConfiguration *GroupItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property groups for groups
func (m *GroupItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *GroupItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get collection of all groups available in the term store.
func (m *GroupItemRequestBuilder) Get(ctx context.Context, requestConfiguration *GroupItemRequestBuilderGetRequestConfiguration)(i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.Groupable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.CreateGroupFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.Groupable), nil
}
// Patch update the navigation property groups in groups
func (m *GroupItemRequestBuilder) Patch(ctx context.Context, body i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.Groupable, requestConfiguration *GroupItemRequestBuilderPatchRequestConfiguration)(i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.Groupable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.CreateGroupFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i45fc41673b99130d86c1854da651a8f416ed902eef3acbecd5738f9ef72690a8.Groupable), nil
}
// Sets provides operations to manage the sets property of the microsoft.graph.termStore.group entity.
func (m *GroupItemRequestBuilder) Sets()(*if37ca3748c081d626a200e3354fca5c0d75d6c98d76383d181b55ada257e45e1.SetsRequestBuilder) {
    return if37ca3748c081d626a200e3354fca5c0d75d6c98d76383d181b55ada257e45e1.NewSetsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SetsById provides operations to manage the sets property of the microsoft.graph.termStore.group entity.
func (m *GroupItemRequestBuilder) SetsById(id string)(*i079fbfa6a17671510d081186b3a4d407748bf5f48e641d244923ea9ba9101107.SetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["set%2Did"] = id
    }
    return i079fbfa6a17671510d081186b3a4d407748bf5f48e641d244923ea9ba9101107.NewSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
