package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
)

// ManagedTenantsManagementActionTenantDeploymentStatusesChangeDeploymentStatusRequestBuilder provides operations to call the changeDeploymentStatus method.
type ManagedTenantsManagementActionTenantDeploymentStatusesChangeDeploymentStatusRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ManagedTenantsManagementActionTenantDeploymentStatusesChangeDeploymentStatusRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedTenantsManagementActionTenantDeploymentStatusesChangeDeploymentStatusRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewManagedTenantsManagementActionTenantDeploymentStatusesChangeDeploymentStatusRequestBuilderInternal instantiates a new ChangeDeploymentStatusRequestBuilder and sets the default values.
func NewManagedTenantsManagementActionTenantDeploymentStatusesChangeDeploymentStatusRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsManagementActionTenantDeploymentStatusesChangeDeploymentStatusRequestBuilder) {
    m := &ManagedTenantsManagementActionTenantDeploymentStatusesChangeDeploymentStatusRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/tenantRelationships/managedTenants/managementActionTenantDeploymentStatuses/microsoft.graph.managedTenants.changeDeploymentStatus";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewManagedTenantsManagementActionTenantDeploymentStatusesChangeDeploymentStatusRequestBuilder instantiates a new ChangeDeploymentStatusRequestBuilder and sets the default values.
func NewManagedTenantsManagementActionTenantDeploymentStatusesChangeDeploymentStatusRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsManagementActionTenantDeploymentStatusesChangeDeploymentStatusRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedTenantsManagementActionTenantDeploymentStatusesChangeDeploymentStatusRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action changeDeploymentStatus
func (m *ManagedTenantsManagementActionTenantDeploymentStatusesChangeDeploymentStatusRequestBuilder) Post(ctx context.Context, body ManagedTenantsManagementActionTenantDeploymentStatusesChangeDeploymentStatusPostRequestBodyable, requestConfiguration *ManagedTenantsManagementActionTenantDeploymentStatusesChangeDeploymentStatusRequestBuilderPostRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementActionDeploymentStatusable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateManagementActionDeploymentStatusFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementActionDeploymentStatusable), nil
}
// ToPostRequestInformation invoke action changeDeploymentStatus
func (m *ManagedTenantsManagementActionTenantDeploymentStatusesChangeDeploymentStatusRequestBuilder) ToPostRequestInformation(ctx context.Context, body ManagedTenantsManagementActionTenantDeploymentStatusesChangeDeploymentStatusPostRequestBodyable, requestConfiguration *ManagedTenantsManagementActionTenantDeploymentStatusesChangeDeploymentStatusRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
