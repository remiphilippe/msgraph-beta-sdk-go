package managementactiontenantdeploymentstatuses

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/managedtenants"
    id50624c28ac1a98a4cd1cc6f61b4dc20c4a7524a54c2bbd7c8cbd10f86ded55f "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managementactiontenantdeploymentstatuses/changedeploymentstatus"
    id677fd2cf57f8c7a79f86b3b17441e9d65b000e7afc795c0b7e5e43f29feb8f7 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managementactiontenantdeploymentstatuses/count"
)

// ManagementActionTenantDeploymentStatusesRequestBuilder provides operations to manage the managementActionTenantDeploymentStatuses property of the microsoft.graph.managedTenants.managedTenant entity.
type ManagementActionTenantDeploymentStatusesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ManagementActionTenantDeploymentStatusesRequestBuilderGetOptions options for Get
type ManagementActionTenantDeploymentStatusesRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ManagementActionTenantDeploymentStatusesRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ManagementActionTenantDeploymentStatusesRequestBuilderGetQueryParameters the tenant level status of management actions across managed tenants.
type ManagementActionTenantDeploymentStatusesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool;
    // Expand related entities
    Expand []string;
    // Filter items by property values
    Filter *string;
    // Order items by property values
    Orderby []string;
    // Search items by search phrases
    Search *string;
    // Select properties to be returned
    Select []string;
    // Skip the first n items
    Skip *int32;
    // Show only the first n items
    Top *int32;
}
// ManagementActionTenantDeploymentStatusesRequestBuilderPostOptions options for Post
type ManagementActionTenantDeploymentStatusesRequestBuilderPostOptions struct {
    // 
    Body i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementActionTenantDeploymentStatusable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ManagementActionTenantDeploymentStatusesRequestBuilder) ChangeDeploymentStatus()(*id50624c28ac1a98a4cd1cc6f61b4dc20c4a7524a54c2bbd7c8cbd10f86ded55f.ChangeDeploymentStatusRequestBuilder) {
    return id50624c28ac1a98a4cd1cc6f61b4dc20c4a7524a54c2bbd7c8cbd10f86ded55f.NewChangeDeploymentStatusRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewManagementActionTenantDeploymentStatusesRequestBuilderInternal instantiates a new ManagementActionTenantDeploymentStatusesRequestBuilder and sets the default values.
func NewManagementActionTenantDeploymentStatusesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagementActionTenantDeploymentStatusesRequestBuilder) {
    m := &ManagementActionTenantDeploymentStatusesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/tenantRelationships/managedTenants/managementActionTenantDeploymentStatuses{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewManagementActionTenantDeploymentStatusesRequestBuilder instantiates a new ManagementActionTenantDeploymentStatusesRequestBuilder and sets the default values.
func NewManagementActionTenantDeploymentStatusesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagementActionTenantDeploymentStatusesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagementActionTenantDeploymentStatusesRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ManagementActionTenantDeploymentStatusesRequestBuilder) Count()(*id677fd2cf57f8c7a79f86b3b17441e9d65b000e7afc795c0b7e5e43f29feb8f7.CountRequestBuilder) {
    return id677fd2cf57f8c7a79f86b3b17441e9d65b000e7afc795c0b7e5e43f29feb8f7.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the tenant level status of management actions across managed tenants.
func (m *ManagementActionTenantDeploymentStatusesRequestBuilder) CreateGetRequestInformation(options *ManagementActionTenantDeploymentStatusesRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePostRequestInformation create new navigation property to managementActionTenantDeploymentStatuses for tenantRelationships
func (m *ManagementActionTenantDeploymentStatusesRequestBuilder) CreatePostRequestInformation(options *ManagementActionTenantDeploymentStatusesRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.POST
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Get the tenant level status of management actions across managed tenants.
func (m *ManagementActionTenantDeploymentStatusesRequestBuilder) Get(options *ManagementActionTenantDeploymentStatusesRequestBuilderGetOptions)(i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementActionTenantDeploymentStatusCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.CreateManagementActionTenantDeploymentStatusCollectionResponseFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementActionTenantDeploymentStatusCollectionResponseable), nil
}
// Post create new navigation property to managementActionTenantDeploymentStatuses for tenantRelationships
func (m *ManagementActionTenantDeploymentStatusesRequestBuilder) Post(options *ManagementActionTenantDeploymentStatusesRequestBuilderPostOptions)(i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementActionTenantDeploymentStatusable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.CreateManagementActionTenantDeploymentStatusFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementActionTenantDeploymentStatusable), nil
}
