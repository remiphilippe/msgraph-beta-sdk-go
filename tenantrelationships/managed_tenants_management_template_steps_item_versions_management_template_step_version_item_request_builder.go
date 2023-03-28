package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
)

// ManagedTenantsManagementTemplateStepsItemVersionsManagementTemplateStepVersionItemRequestBuilder provides operations to manage the versions property of the microsoft.graph.managedTenants.managementTemplateStep entity.
type ManagedTenantsManagementTemplateStepsItemVersionsManagementTemplateStepVersionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedTenantsManagementTemplateStepsItemVersionsManagementTemplateStepVersionItemRequestBuilderGetQueryParameters get versions from tenantRelationships
type ManagedTenantsManagementTemplateStepsItemVersionsManagementTemplateStepVersionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ManagedTenantsManagementTemplateStepsItemVersionsManagementTemplateStepVersionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedTenantsManagementTemplateStepsItemVersionsManagementTemplateStepVersionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedTenantsManagementTemplateStepsItemVersionsManagementTemplateStepVersionItemRequestBuilderGetQueryParameters
}
// NewManagedTenantsManagementTemplateStepsItemVersionsManagementTemplateStepVersionItemRequestBuilderInternal instantiates a new ManagementTemplateStepVersionItemRequestBuilder and sets the default values.
func NewManagedTenantsManagementTemplateStepsItemVersionsManagementTemplateStepVersionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsManagementTemplateStepsItemVersionsManagementTemplateStepVersionItemRequestBuilder) {
    m := &ManagedTenantsManagementTemplateStepsItemVersionsManagementTemplateStepVersionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/managedTenants/managementTemplateSteps/{managementTemplateStep%2Did}/versions/{managementTemplateStepVersion%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewManagedTenantsManagementTemplateStepsItemVersionsManagementTemplateStepVersionItemRequestBuilder instantiates a new ManagementTemplateStepVersionItemRequestBuilder and sets the default values.
func NewManagedTenantsManagementTemplateStepsItemVersionsManagementTemplateStepVersionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsManagementTemplateStepsItemVersionsManagementTemplateStepVersionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedTenantsManagementTemplateStepsItemVersionsManagementTemplateStepVersionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get versions from tenantRelationships
func (m *ManagedTenantsManagementTemplateStepsItemVersionsManagementTemplateStepVersionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedTenantsManagementTemplateStepsItemVersionsManagementTemplateStepVersionItemRequestBuilderGetRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementTemplateStepVersionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateManagementTemplateStepVersionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementTemplateStepVersionable), nil
}
// ToGetRequestInformation get versions from tenantRelationships
func (m *ManagedTenantsManagementTemplateStepsItemVersionsManagementTemplateStepVersionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManagedTenantsManagementTemplateStepsItemVersionsManagementTemplateStepVersionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
