package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
)

// ManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilder provides operations to manage the managementTemplateStepVersions property of the microsoft.graph.managedTenants.managedTenant entity.
type ManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilderGetQueryParameters get managementTemplateStepVersions from tenantRelationships
type ManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilderGetQueryParameters
}
// ManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AcceptedFor provides operations to manage the acceptedFor property of the microsoft.graph.managedTenants.managementTemplateStepVersion entity.
func (m *ManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilder) AcceptedFor()(*ManagedTenantsManagementTemplateStepVersionsItemAcceptedForRequestBuilder) {
    return NewManagedTenantsManagementTemplateStepVersionsItemAcceptedForRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilderInternal instantiates a new ManagementTemplateStepVersionItemRequestBuilder and sets the default values.
func NewManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilder) {
    m := &ManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/managedTenants/managementTemplateStepVersions/{managementTemplateStepVersion%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilder instantiates a new ManagementTemplateStepVersionItemRequestBuilder and sets the default values.
func NewManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property managementTemplateStepVersions for tenantRelationships
func (m *ManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Deployments provides operations to manage the deployments property of the microsoft.graph.managedTenants.managementTemplateStepVersion entity.
func (m *ManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilder) Deployments()(*ManagedTenantsManagementTemplateStepVersionsItemDeploymentsRequestBuilder) {
    return NewManagedTenantsManagementTemplateStepVersionsItemDeploymentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeploymentsById provides operations to manage the deployments property of the microsoft.graph.managedTenants.managementTemplateStepVersion entity.
func (m *ManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilder) DeploymentsById(id string)(*ManagedTenantsManagementTemplateStepVersionsItemDeploymentsManagementTemplateStepDeploymentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managementTemplateStepDeployment%2Did"] = id
    }
    return NewManagedTenantsManagementTemplateStepVersionsItemDeploymentsManagementTemplateStepDeploymentItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// Get get managementTemplateStepVersions from tenantRelationships
func (m *ManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilderGetRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementTemplateStepVersionable, error) {
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
// Patch update the navigation property managementTemplateStepVersions in tenantRelationships
func (m *ManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilder) Patch(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementTemplateStepVersionable, requestConfiguration *ManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilderPatchRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementTemplateStepVersionable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// TemplateStep provides operations to manage the templateStep property of the microsoft.graph.managedTenants.managementTemplateStepVersion entity.
func (m *ManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilder) TemplateStep()(*ManagedTenantsManagementTemplateStepVersionsItemTemplateStepRequestBuilder) {
    return NewManagedTenantsManagementTemplateStepVersionsItemTemplateStepRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property managementTemplateStepVersions for tenantRelationships
func (m *ManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation get managementTemplateStepVersions from tenantRelationships
func (m *ManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property managementTemplateStepVersions in tenantRelationships
func (m *ManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementTemplateStepVersionable, requestConfiguration *ManagedTenantsManagementTemplateStepVersionsManagementTemplateStepVersionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
