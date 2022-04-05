package identityprotection

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i0e267c8100f2aba2f7153bd084faad7f539542ced6c643bbda5c3501e2ccc9a4 "github.com/microsoftgraph/msgraph-beta-sdk-go/identityprotection/riskdetections"
    i19f24f05457f3b35ed5a71197a5407212d3b7238b612e5bf489fbc4202271989 "github.com/microsoftgraph/msgraph-beta-sdk-go/identityprotection/riskyusers"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i98ea4cb17d390592303b9e25c8f92ccaf566e5d10c2cabc9f160512904151e07 "github.com/microsoftgraph/msgraph-beta-sdk-go/identityprotection/riskyserviceprincipals"
    ic7b59228365548f0fce94cdce65ae97f7b560a35a266ece91414c55ca4e5c249 "github.com/microsoftgraph/msgraph-beta-sdk-go/identityprotection/serviceprincipalriskdetections"
    i294f460d075f1fad271e61484dbf4976eb35e90151e8cbeea5c33955169dbb83 "github.com/microsoftgraph/msgraph-beta-sdk-go/identityprotection/riskdetections/item"
    i3b4ddc0707ce5dc4f8bc784d8e4305e5ba05b0af6b7287ebf1b1abc4f5796032 "github.com/microsoftgraph/msgraph-beta-sdk-go/identityprotection/serviceprincipalriskdetections/item"
    i5b186cdb023bab2640affb6c67369c3f8028451aa7845fa6d84e60b0d9b0fc05 "github.com/microsoftgraph/msgraph-beta-sdk-go/identityprotection/riskyusers/item"
    i93ffa285a6083330a0c11899f97deea0e0a67ea5245ac02339ac77d9f35ce82b "github.com/microsoftgraph/msgraph-beta-sdk-go/identityprotection/riskyserviceprincipals/item"
)

// IdentityProtectionRequestBuilder provides operations to manage the identityProtectionRoot singleton.
type IdentityProtectionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// IdentityProtectionRequestBuilderGetOptions options for Get
type IdentityProtectionRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *IdentityProtectionRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// IdentityProtectionRequestBuilderGetQueryParameters get identityProtection
type IdentityProtectionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// IdentityProtectionRequestBuilderPatchOptions options for Patch
type IdentityProtectionRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityProtectionRootable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// NewIdentityProtectionRequestBuilderInternal instantiates a new IdentityProtectionRequestBuilder and sets the default values.
func NewIdentityProtectionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityProtectionRequestBuilder) {
    m := &IdentityProtectionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityProtection{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIdentityProtectionRequestBuilder instantiates a new IdentityProtectionRequestBuilder and sets the default values.
func NewIdentityProtectionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityProtectionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityProtectionRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get identityProtection
func (m *IdentityProtectionRequestBuilder) CreateGetRequestInformation(options *IdentityProtectionRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update identityProtection
func (m *IdentityProtectionRequestBuilder) CreatePatchRequestInformation(options *IdentityProtectionRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Get get identityProtection
func (m *IdentityProtectionRequestBuilder) Get(options *IdentityProtectionRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityProtectionRootable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateIdentityProtectionRootFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityProtectionRootable), nil
}
// Patch update identityProtection
func (m *IdentityProtectionRequestBuilder) Patch(options *IdentityProtectionRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// RiskDetections the riskDetections property
func (m *IdentityProtectionRequestBuilder) RiskDetections()(*i0e267c8100f2aba2f7153bd084faad7f539542ced6c643bbda5c3501e2ccc9a4.RiskDetectionsRequestBuilder) {
    return i0e267c8100f2aba2f7153bd084faad7f539542ced6c643bbda5c3501e2ccc9a4.NewRiskDetectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RiskDetectionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityProtection.riskDetections.item collection
func (m *IdentityProtectionRequestBuilder) RiskDetectionsById(id string)(*i294f460d075f1fad271e61484dbf4976eb35e90151e8cbeea5c33955169dbb83.RiskDetectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["riskDetection_id"] = id
    }
    return i294f460d075f1fad271e61484dbf4976eb35e90151e8cbeea5c33955169dbb83.NewRiskDetectionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RiskyServicePrincipals the riskyServicePrincipals property
func (m *IdentityProtectionRequestBuilder) RiskyServicePrincipals()(*i98ea4cb17d390592303b9e25c8f92ccaf566e5d10c2cabc9f160512904151e07.RiskyServicePrincipalsRequestBuilder) {
    return i98ea4cb17d390592303b9e25c8f92ccaf566e5d10c2cabc9f160512904151e07.NewRiskyServicePrincipalsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RiskyServicePrincipalsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityProtection.riskyServicePrincipals.item collection
func (m *IdentityProtectionRequestBuilder) RiskyServicePrincipalsById(id string)(*i93ffa285a6083330a0c11899f97deea0e0a67ea5245ac02339ac77d9f35ce82b.RiskyServicePrincipalItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["riskyServicePrincipal_id"] = id
    }
    return i93ffa285a6083330a0c11899f97deea0e0a67ea5245ac02339ac77d9f35ce82b.NewRiskyServicePrincipalItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RiskyUsers the riskyUsers property
func (m *IdentityProtectionRequestBuilder) RiskyUsers()(*i19f24f05457f3b35ed5a71197a5407212d3b7238b612e5bf489fbc4202271989.RiskyUsersRequestBuilder) {
    return i19f24f05457f3b35ed5a71197a5407212d3b7238b612e5bf489fbc4202271989.NewRiskyUsersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RiskyUsersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityProtection.riskyUsers.item collection
func (m *IdentityProtectionRequestBuilder) RiskyUsersById(id string)(*i5b186cdb023bab2640affb6c67369c3f8028451aa7845fa6d84e60b0d9b0fc05.RiskyUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["riskyUser_id"] = id
    }
    return i5b186cdb023bab2640affb6c67369c3f8028451aa7845fa6d84e60b0d9b0fc05.NewRiskyUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ServicePrincipalRiskDetections the servicePrincipalRiskDetections property
func (m *IdentityProtectionRequestBuilder) ServicePrincipalRiskDetections()(*ic7b59228365548f0fce94cdce65ae97f7b560a35a266ece91414c55ca4e5c249.ServicePrincipalRiskDetectionsRequestBuilder) {
    return ic7b59228365548f0fce94cdce65ae97f7b560a35a266ece91414c55ca4e5c249.NewServicePrincipalRiskDetectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipalRiskDetectionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identityProtection.servicePrincipalRiskDetections.item collection
func (m *IdentityProtectionRequestBuilder) ServicePrincipalRiskDetectionsById(id string)(*i3b4ddc0707ce5dc4f8bc784d8e4305e5ba05b0af6b7287ebf1b1abc4f5796032.ServicePrincipalRiskDetectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["servicePrincipalRiskDetection_id"] = id
    }
    return i3b4ddc0707ce5dc4f8bc784d8e4305e5ba05b0af6b7287ebf1b1abc4f5796032.NewServicePrincipalRiskDetectionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
