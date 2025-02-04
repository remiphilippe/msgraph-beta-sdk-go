package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// IdentityRequestBuilder provides operations to manage the identityContainer singleton.
type IdentityRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// IdentityRequestBuilderGetQueryParameters get identity
type IdentityRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IdentityRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IdentityRequestBuilderGetQueryParameters
}
// IdentityRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ApiConnectors provides operations to manage the apiConnectors property of the microsoft.graph.identityContainer entity.
func (m *IdentityRequestBuilder) ApiConnectors()(*ApiConnectorsRequestBuilder) {
    return NewApiConnectorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ApiConnectorsById provides operations to manage the apiConnectors property of the microsoft.graph.identityContainer entity.
func (m *IdentityRequestBuilder) ApiConnectorsById(id string)(*ApiConnectorsIdentityApiConnectorItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["identityApiConnector%2Did"] = id
    }
    return NewApiConnectorsIdentityApiConnectorItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AuthenticationEventListeners provides operations to manage the authenticationEventListeners property of the microsoft.graph.identityContainer entity.
func (m *IdentityRequestBuilder) AuthenticationEventListeners()(*AuthenticationEventListenersRequestBuilder) {
    return NewAuthenticationEventListenersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AuthenticationEventListenersById provides operations to manage the authenticationEventListeners property of the microsoft.graph.identityContainer entity.
func (m *IdentityRequestBuilder) AuthenticationEventListenersById(id string)(*AuthenticationEventListenersAuthenticationEventListenerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["authenticationEventListener%2Did"] = id
    }
    return NewAuthenticationEventListenersAuthenticationEventListenerItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// B2cUserFlows provides operations to manage the b2cUserFlows property of the microsoft.graph.identityContainer entity.
func (m *IdentityRequestBuilder) B2cUserFlows()(*B2cUserFlowsRequestBuilder) {
    return NewB2cUserFlowsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// B2cUserFlowsById provides operations to manage the b2cUserFlows property of the microsoft.graph.identityContainer entity.
func (m *IdentityRequestBuilder) B2cUserFlowsById(id string)(*B2cUserFlowsB2cIdentityUserFlowItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["b2cIdentityUserFlow%2Did"] = id
    }
    return NewB2cUserFlowsB2cIdentityUserFlowItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// B2xUserFlows provides operations to manage the b2xUserFlows property of the microsoft.graph.identityContainer entity.
func (m *IdentityRequestBuilder) B2xUserFlows()(*B2xUserFlowsRequestBuilder) {
    return NewB2xUserFlowsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// B2xUserFlowsById provides operations to manage the b2xUserFlows property of the microsoft.graph.identityContainer entity.
func (m *IdentityRequestBuilder) B2xUserFlowsById(id string)(*B2xUserFlowsB2xIdentityUserFlowItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["b2xIdentityUserFlow%2Did"] = id
    }
    return NewB2xUserFlowsB2xIdentityUserFlowItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ConditionalAccess provides operations to manage the conditionalAccess property of the microsoft.graph.identityContainer entity.
func (m *IdentityRequestBuilder) ConditionalAccess()(*ConditionalAccessRequestBuilder) {
    return NewConditionalAccessRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewIdentityRequestBuilderInternal instantiates a new IdentityRequestBuilder and sets the default values.
func NewIdentityRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityRequestBuilder) {
    m := &IdentityRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identity{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIdentityRequestBuilder instantiates a new IdentityRequestBuilder and sets the default values.
func NewIdentityRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityRequestBuilderInternal(urlParams, requestAdapter)
}
// ContinuousAccessEvaluationPolicy provides operations to manage the continuousAccessEvaluationPolicy property of the microsoft.graph.identityContainer entity.
func (m *IdentityRequestBuilder) ContinuousAccessEvaluationPolicy()(*ContinuousAccessEvaluationPolicyRequestBuilder) {
    return NewContinuousAccessEvaluationPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CustomAuthenticationExtensions provides operations to manage the customAuthenticationExtensions property of the microsoft.graph.identityContainer entity.
func (m *IdentityRequestBuilder) CustomAuthenticationExtensions()(*CustomAuthenticationExtensionsRequestBuilder) {
    return NewCustomAuthenticationExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CustomAuthenticationExtensionsById provides operations to manage the customAuthenticationExtensions property of the microsoft.graph.identityContainer entity.
func (m *IdentityRequestBuilder) CustomAuthenticationExtensionsById(id string)(*CustomAuthenticationExtensionsCustomAuthenticationExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["customAuthenticationExtension%2Did"] = id
    }
    return NewCustomAuthenticationExtensionsCustomAuthenticationExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get identity
func (m *IdentityRequestBuilder) Get(ctx context.Context, requestConfiguration *IdentityRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityContainerable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateIdentityContainerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityContainerable), nil
}
// IdentityProviders provides operations to manage the identityProviders property of the microsoft.graph.identityContainer entity.
func (m *IdentityRequestBuilder) IdentityProviders()(*IdentityProvidersRequestBuilder) {
    return NewIdentityProvidersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IdentityProvidersById provides operations to manage the identityProviders property of the microsoft.graph.identityContainer entity.
func (m *IdentityRequestBuilder) IdentityProvidersById(id string)(*IdentityProvidersIdentityProviderBaseItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["identityProviderBase%2Did"] = id
    }
    return NewIdentityProvidersIdentityProviderBaseItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update identity
func (m *IdentityRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityContainerable, requestConfiguration *IdentityRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityContainerable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateIdentityContainerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityContainerable), nil
}
// ToGetRequestInformation get identity
func (m *IdentityRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *IdentityRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update identity
func (m *IdentityRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentityContainerable, requestConfiguration *IdentityRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// UserFlowAttributes provides operations to manage the userFlowAttributes property of the microsoft.graph.identityContainer entity.
func (m *IdentityRequestBuilder) UserFlowAttributes()(*UserFlowAttributesRequestBuilder) {
    return NewUserFlowAttributesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserFlowAttributesById provides operations to manage the userFlowAttributes property of the microsoft.graph.identityContainer entity.
func (m *IdentityRequestBuilder) UserFlowAttributesById(id string)(*UserFlowAttributesIdentityUserFlowAttributeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["identityUserFlowAttribute%2Did"] = id
    }
    return NewUserFlowAttributesIdentityUserFlowAttributeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UserFlows provides operations to manage the userFlows property of the microsoft.graph.identityContainer entity.
func (m *IdentityRequestBuilder) UserFlows()(*UserFlowsRequestBuilder) {
    return NewUserFlowsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserFlowsById provides operations to manage the userFlows property of the microsoft.graph.identityContainer entity.
func (m *IdentityRequestBuilder) UserFlowsById(id string)(*UserFlowsIdentityUserFlowItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["identityUserFlow%2Did"] = id
    }
    return NewUserFlowsIdentityUserFlowItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
