package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/identitygovernance"
)

// LifecycleWorkflowsDeletedItemsWorkflowsItemRunsRunItemRequestBuilder provides operations to manage the runs property of the microsoft.graph.identityGovernance.workflow entity.
type LifecycleWorkflowsDeletedItemsWorkflowsItemRunsRunItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// LifecycleWorkflowsDeletedItemsWorkflowsItemRunsRunItemRequestBuilderGetQueryParameters get runs from identityGovernance
type LifecycleWorkflowsDeletedItemsWorkflowsItemRunsRunItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// LifecycleWorkflowsDeletedItemsWorkflowsItemRunsRunItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleWorkflowsDeletedItemsWorkflowsItemRunsRunItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LifecycleWorkflowsDeletedItemsWorkflowsItemRunsRunItemRequestBuilderGetQueryParameters
}
// NewLifecycleWorkflowsDeletedItemsWorkflowsItemRunsRunItemRequestBuilderInternal instantiates a new RunItemRequestBuilder and sets the default values.
func NewLifecycleWorkflowsDeletedItemsWorkflowsItemRunsRunItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleWorkflowsDeletedItemsWorkflowsItemRunsRunItemRequestBuilder) {
    m := &LifecycleWorkflowsDeletedItemsWorkflowsItemRunsRunItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/lifecycleWorkflows/deletedItems/workflows/{workflow%2Did}/runs/{run%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewLifecycleWorkflowsDeletedItemsWorkflowsItemRunsRunItemRequestBuilder instantiates a new RunItemRequestBuilder and sets the default values.
func NewLifecycleWorkflowsDeletedItemsWorkflowsItemRunsRunItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleWorkflowsDeletedItemsWorkflowsItemRunsRunItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLifecycleWorkflowsDeletedItemsWorkflowsItemRunsRunItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get runs from identityGovernance
func (m *LifecycleWorkflowsDeletedItemsWorkflowsItemRunsRunItemRequestBuilder) Get(ctx context.Context, requestConfiguration *LifecycleWorkflowsDeletedItemsWorkflowsItemRunsRunItemRequestBuilderGetRequestConfiguration)(i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.Runable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.CreateRunFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i45fdec8a8c1f65ca74c5cf52921d432ad02ee300dbbd24b25f33cc8ecf6a1a91.Runable), nil
}
// TaskProcessingResults provides operations to manage the taskProcessingResults property of the microsoft.graph.identityGovernance.run entity.
func (m *LifecycleWorkflowsDeletedItemsWorkflowsItemRunsRunItemRequestBuilder) TaskProcessingResults()(*LifecycleWorkflowsDeletedItemsWorkflowsItemRunsItemTaskProcessingResultsRequestBuilder) {
    return NewLifecycleWorkflowsDeletedItemsWorkflowsItemRunsItemTaskProcessingResultsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TaskProcessingResultsById provides operations to manage the taskProcessingResults property of the microsoft.graph.identityGovernance.run entity.
func (m *LifecycleWorkflowsDeletedItemsWorkflowsItemRunsRunItemRequestBuilder) TaskProcessingResultsById(id string)(*LifecycleWorkflowsDeletedItemsWorkflowsItemRunsItemTaskProcessingResultsTaskProcessingResultItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["taskProcessingResult%2Did"] = id
    }
    return NewLifecycleWorkflowsDeletedItemsWorkflowsItemRunsItemTaskProcessingResultsTaskProcessingResultItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ToGetRequestInformation get runs from identityGovernance
func (m *LifecycleWorkflowsDeletedItemsWorkflowsItemRunsRunItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LifecycleWorkflowsDeletedItemsWorkflowsItemRunsRunItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// UserProcessingResults provides operations to manage the userProcessingResults property of the microsoft.graph.identityGovernance.run entity.
func (m *LifecycleWorkflowsDeletedItemsWorkflowsItemRunsRunItemRequestBuilder) UserProcessingResults()(*LifecycleWorkflowsDeletedItemsWorkflowsItemRunsItemUserProcessingResultsRequestBuilder) {
    return NewLifecycleWorkflowsDeletedItemsWorkflowsItemRunsItemUserProcessingResultsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserProcessingResultsById provides operations to manage the userProcessingResults property of the microsoft.graph.identityGovernance.run entity.
func (m *LifecycleWorkflowsDeletedItemsWorkflowsItemRunsRunItemRequestBuilder) UserProcessingResultsById(id string)(*LifecycleWorkflowsDeletedItemsWorkflowsItemRunsItemUserProcessingResultsUserProcessingResultItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userProcessingResult%2Did"] = id
    }
    return NewLifecycleWorkflowsDeletedItemsWorkflowsItemRunsItemUserProcessingResultsUserProcessingResultItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
