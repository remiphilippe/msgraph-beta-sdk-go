package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PrivilegedAccessGroupRequestBuilder provides operations to manage the group property of the microsoft.graph.privilegedAccessRoot entity.
type PrivilegedAccessGroupRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PrivilegedAccessGroupRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedAccessGroupRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// PrivilegedAccessGroupRequestBuilderGetQueryParameters get group from identityGovernance
type PrivilegedAccessGroupRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PrivilegedAccessGroupRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedAccessGroupRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PrivilegedAccessGroupRequestBuilderGetQueryParameters
}
// PrivilegedAccessGroupRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedAccessGroupRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AssignmentScheduleInstances provides operations to manage the assignmentScheduleInstances property of the microsoft.graph.privilegedAccessGroup entity.
func (m *PrivilegedAccessGroupRequestBuilder) AssignmentScheduleInstances()(*PrivilegedAccessGroupAssignmentScheduleInstancesRequestBuilder) {
    return NewPrivilegedAccessGroupAssignmentScheduleInstancesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AssignmentScheduleInstancesById provides operations to manage the assignmentScheduleInstances property of the microsoft.graph.privilegedAccessGroup entity.
func (m *PrivilegedAccessGroupRequestBuilder) AssignmentScheduleInstancesById(id string)(*PrivilegedAccessGroupAssignmentScheduleInstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["privilegedAccessGroupAssignmentScheduleInstance%2Did"] = id
    }
    return NewPrivilegedAccessGroupAssignmentScheduleInstancesPrivilegedAccessGroupAssignmentScheduleInstanceItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// AssignmentScheduleRequests provides operations to manage the assignmentScheduleRequests property of the microsoft.graph.privilegedAccessGroup entity.
func (m *PrivilegedAccessGroupRequestBuilder) AssignmentScheduleRequests()(*PrivilegedAccessGroupAssignmentScheduleRequestsRequestBuilder) {
    return NewPrivilegedAccessGroupAssignmentScheduleRequestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AssignmentScheduleRequestsById provides operations to manage the assignmentScheduleRequests property of the microsoft.graph.privilegedAccessGroup entity.
func (m *PrivilegedAccessGroupRequestBuilder) AssignmentScheduleRequestsById(id string)(*PrivilegedAccessGroupAssignmentScheduleRequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["privilegedAccessGroupAssignmentScheduleRequest%2Did"] = id
    }
    return NewPrivilegedAccessGroupAssignmentScheduleRequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// AssignmentSchedules provides operations to manage the assignmentSchedules property of the microsoft.graph.privilegedAccessGroup entity.
func (m *PrivilegedAccessGroupRequestBuilder) AssignmentSchedules()(*PrivilegedAccessGroupAssignmentSchedulesRequestBuilder) {
    return NewPrivilegedAccessGroupAssignmentSchedulesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AssignmentSchedulesById provides operations to manage the assignmentSchedules property of the microsoft.graph.privilegedAccessGroup entity.
func (m *PrivilegedAccessGroupRequestBuilder) AssignmentSchedulesById(id string)(*PrivilegedAccessGroupAssignmentSchedulesPrivilegedAccessGroupAssignmentScheduleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["privilegedAccessGroupAssignmentSchedule%2Did"] = id
    }
    return NewPrivilegedAccessGroupAssignmentSchedulesPrivilegedAccessGroupAssignmentScheduleItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewPrivilegedAccessGroupRequestBuilderInternal instantiates a new GroupRequestBuilder and sets the default values.
func NewPrivilegedAccessGroupRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedAccessGroupRequestBuilder) {
    m := &PrivilegedAccessGroupRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/privilegedAccess/group{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewPrivilegedAccessGroupRequestBuilder instantiates a new GroupRequestBuilder and sets the default values.
func NewPrivilegedAccessGroupRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedAccessGroupRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrivilegedAccessGroupRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property group for identityGovernance
func (m *PrivilegedAccessGroupRequestBuilder) Delete(ctx context.Context, requestConfiguration *PrivilegedAccessGroupRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// EligibilityScheduleInstances provides operations to manage the eligibilityScheduleInstances property of the microsoft.graph.privilegedAccessGroup entity.
func (m *PrivilegedAccessGroupRequestBuilder) EligibilityScheduleInstances()(*PrivilegedAccessGroupEligibilityScheduleInstancesRequestBuilder) {
    return NewPrivilegedAccessGroupEligibilityScheduleInstancesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// EligibilityScheduleInstancesById provides operations to manage the eligibilityScheduleInstances property of the microsoft.graph.privilegedAccessGroup entity.
func (m *PrivilegedAccessGroupRequestBuilder) EligibilityScheduleInstancesById(id string)(*PrivilegedAccessGroupEligibilityScheduleInstancesPrivilegedAccessGroupEligibilityScheduleInstanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["privilegedAccessGroupEligibilityScheduleInstance%2Did"] = id
    }
    return NewPrivilegedAccessGroupEligibilityScheduleInstancesPrivilegedAccessGroupEligibilityScheduleInstanceItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// EligibilityScheduleRequests provides operations to manage the eligibilityScheduleRequests property of the microsoft.graph.privilegedAccessGroup entity.
func (m *PrivilegedAccessGroupRequestBuilder) EligibilityScheduleRequests()(*PrivilegedAccessGroupEligibilityScheduleRequestsRequestBuilder) {
    return NewPrivilegedAccessGroupEligibilityScheduleRequestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// EligibilityScheduleRequestsById provides operations to manage the eligibilityScheduleRequests property of the microsoft.graph.privilegedAccessGroup entity.
func (m *PrivilegedAccessGroupRequestBuilder) EligibilityScheduleRequestsById(id string)(*PrivilegedAccessGroupEligibilityScheduleRequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["privilegedAccessGroupEligibilityScheduleRequest%2Did"] = id
    }
    return NewPrivilegedAccessGroupEligibilityScheduleRequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// EligibilitySchedules provides operations to manage the eligibilitySchedules property of the microsoft.graph.privilegedAccessGroup entity.
func (m *PrivilegedAccessGroupRequestBuilder) EligibilitySchedules()(*PrivilegedAccessGroupEligibilitySchedulesRequestBuilder) {
    return NewPrivilegedAccessGroupEligibilitySchedulesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// EligibilitySchedulesById provides operations to manage the eligibilitySchedules property of the microsoft.graph.privilegedAccessGroup entity.
func (m *PrivilegedAccessGroupRequestBuilder) EligibilitySchedulesById(id string)(*PrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["privilegedAccessGroupEligibilitySchedule%2Did"] = id
    }
    return NewPrivilegedAccessGroupEligibilitySchedulesPrivilegedAccessGroupEligibilityScheduleItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// Get get group from identityGovernance
func (m *PrivilegedAccessGroupRequestBuilder) Get(ctx context.Context, requestConfiguration *PrivilegedAccessGroupRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedAccessGroupable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePrivilegedAccessGroupFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedAccessGroupable), nil
}
// Patch update the navigation property group in identityGovernance
func (m *PrivilegedAccessGroupRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedAccessGroupable, requestConfiguration *PrivilegedAccessGroupRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedAccessGroupable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePrivilegedAccessGroupFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedAccessGroupable), nil
}
// ToDeleteRequestInformation delete navigation property group for identityGovernance
func (m *PrivilegedAccessGroupRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *PrivilegedAccessGroupRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation get group from identityGovernance
func (m *PrivilegedAccessGroupRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PrivilegedAccessGroupRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property group in identityGovernance
func (m *PrivilegedAccessGroupRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedAccessGroupable, requestConfiguration *PrivilegedAccessGroupRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
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
