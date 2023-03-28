package external

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a "github.com/microsoftgraph/msgraph-beta-sdk-go/models/industrydata"
)

// IndustryDataRequestBuilder provides operations to manage the industryData property of the microsoft.graph.externalConnectors.external entity.
type IndustryDataRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// IndustryDataRequestBuilderGetQueryParameters get industryData from external
type IndustryDataRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IndustryDataRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IndustryDataRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IndustryDataRequestBuilderGetQueryParameters
}
// NewIndustryDataRequestBuilderInternal instantiates a new IndustryDataRequestBuilder and sets the default values.
func NewIndustryDataRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IndustryDataRequestBuilder) {
    m := &IndustryDataRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/external/industryData{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewIndustryDataRequestBuilder instantiates a new IndustryDataRequestBuilder and sets the default values.
func NewIndustryDataRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IndustryDataRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIndustryDataRequestBuilderInternal(urlParams, requestAdapter)
}
// DataConnectors provides operations to manage the dataConnectors property of the microsoft.graph.industryData.industryDataRoot entity.
func (m *IndustryDataRequestBuilder) DataConnectors()(*IndustryDataDataConnectorsRequestBuilder) {
    return NewIndustryDataDataConnectorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DataConnectorsById provides operations to manage the dataConnectors property of the microsoft.graph.industryData.industryDataRoot entity.
func (m *IndustryDataRequestBuilder) DataConnectorsById(id string)(*IndustryDataDataConnectorsIndustryDataConnectorItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["industryDataConnector%2Did"] = id
    }
    return NewIndustryDataDataConnectorsIndustryDataConnectorItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// Get get industryData from external
func (m *IndustryDataRequestBuilder) Get(ctx context.Context, requestConfiguration *IndustryDataRequestBuilderGetRequestConfiguration)(id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.IndustryDataRootable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.CreateIndustryDataRootFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.IndustryDataRootable), nil
}
// InboundFlows provides operations to manage the inboundFlows property of the microsoft.graph.industryData.industryDataRoot entity.
func (m *IndustryDataRequestBuilder) InboundFlows()(*IndustryDataInboundFlowsRequestBuilder) {
    return NewIndustryDataInboundFlowsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// InboundFlowsById provides operations to manage the inboundFlows property of the microsoft.graph.industryData.industryDataRoot entity.
func (m *IndustryDataRequestBuilder) InboundFlowsById(id string)(*IndustryDataInboundFlowsInboundFlowItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["inboundFlow%2Did"] = id
    }
    return NewIndustryDataInboundFlowsInboundFlowItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// Operations provides operations to manage the operations property of the microsoft.graph.industryData.industryDataRoot entity.
func (m *IndustryDataRequestBuilder) Operations()(*IndustryDataOperationsRequestBuilder) {
    return NewIndustryDataOperationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// OperationsById provides operations to manage the operations property of the microsoft.graph.industryData.industryDataRoot entity.
func (m *IndustryDataRequestBuilder) OperationsById(id string)(*IndustryDataOperationsLongRunningOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["longRunningOperation%2Did"] = id
    }
    return NewIndustryDataOperationsLongRunningOperationItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ReferenceDefinitions provides operations to manage the referenceDefinitions property of the microsoft.graph.industryData.industryDataRoot entity.
func (m *IndustryDataRequestBuilder) ReferenceDefinitions()(*IndustryDataReferenceDefinitionsRequestBuilder) {
    return NewIndustryDataReferenceDefinitionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ReferenceDefinitionsById provides operations to manage the referenceDefinitions property of the microsoft.graph.industryData.industryDataRoot entity.
func (m *IndustryDataRequestBuilder) ReferenceDefinitionsById(id string)(*IndustryDataReferenceDefinitionsReferenceDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["referenceDefinition%2Did"] = id
    }
    return NewIndustryDataReferenceDefinitionsReferenceDefinitionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// RoleGroups provides operations to manage the roleGroups property of the microsoft.graph.industryData.industryDataRoot entity.
func (m *IndustryDataRequestBuilder) RoleGroups()(*IndustryDataRoleGroupsRequestBuilder) {
    return NewIndustryDataRoleGroupsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RoleGroupsById provides operations to manage the roleGroups property of the microsoft.graph.industryData.industryDataRoot entity.
func (m *IndustryDataRequestBuilder) RoleGroupsById(id string)(*IndustryDataRoleGroupsRoleGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["roleGroup%2Did"] = id
    }
    return NewIndustryDataRoleGroupsRoleGroupItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// Runs provides operations to manage the runs property of the microsoft.graph.industryData.industryDataRoot entity.
func (m *IndustryDataRequestBuilder) Runs()(*IndustryDataRunsRequestBuilder) {
    return NewIndustryDataRunsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RunsById provides operations to manage the runs property of the microsoft.graph.industryData.industryDataRoot entity.
func (m *IndustryDataRequestBuilder) RunsById(id string)(*IndustryDataRunsIndustryDataRunItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["industryDataRun%2Did"] = id
    }
    return NewIndustryDataRunsIndustryDataRunItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// SourceSystems provides operations to manage the sourceSystems property of the microsoft.graph.industryData.industryDataRoot entity.
func (m *IndustryDataRequestBuilder) SourceSystems()(*IndustryDataSourceSystemsRequestBuilder) {
    return NewIndustryDataSourceSystemsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SourceSystemsById provides operations to manage the sourceSystems property of the microsoft.graph.industryData.industryDataRoot entity.
func (m *IndustryDataRequestBuilder) SourceSystemsById(id string)(*IndustryDataSourceSystemsSourceSystemDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sourceSystemDefinition%2Did"] = id
    }
    return NewIndustryDataSourceSystemsSourceSystemDefinitionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get industryData from external
func (m *IndustryDataRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *IndustryDataRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
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
// Years provides operations to manage the years property of the microsoft.graph.industryData.industryDataRoot entity.
func (m *IndustryDataRequestBuilder) Years()(*IndustryDataYearsRequestBuilder) {
    return NewIndustryDataYearsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// YearsById provides operations to manage the years property of the microsoft.graph.industryData.industryDataRoot entity.
func (m *IndustryDataRequestBuilder) YearsById(id string)(*IndustryDataYearsYearTimePeriodDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["yearTimePeriodDefinition%2Did"] = id
    }
    return NewIndustryDataYearsYearTimePeriodDefinitionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
