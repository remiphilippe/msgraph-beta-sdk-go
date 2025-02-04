package identityprotection

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// RiskyServicePrincipalsConfirmCompromisedRequestBuilder provides operations to call the confirmCompromised method.
type RiskyServicePrincipalsConfirmCompromisedRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// RiskyServicePrincipalsConfirmCompromisedRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RiskyServicePrincipalsConfirmCompromisedRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewRiskyServicePrincipalsConfirmCompromisedRequestBuilderInternal instantiates a new ConfirmCompromisedRequestBuilder and sets the default values.
func NewRiskyServicePrincipalsConfirmCompromisedRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RiskyServicePrincipalsConfirmCompromisedRequestBuilder) {
    m := &RiskyServicePrincipalsConfirmCompromisedRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityProtection/riskyServicePrincipals/microsoft.graph.confirmCompromised";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRiskyServicePrincipalsConfirmCompromisedRequestBuilder instantiates a new ConfirmCompromisedRequestBuilder and sets the default values.
func NewRiskyServicePrincipalsConfirmCompromisedRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RiskyServicePrincipalsConfirmCompromisedRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRiskyServicePrincipalsConfirmCompromisedRequestBuilderInternal(urlParams, requestAdapter)
}
// Post confirm one or more riskyServicePrincipal objects as compromised. This action sets the targeted service principal account's risk level to `high`.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/riskyserviceprincipal-confirmcompromised?view=graph-rest-1.0
func (m *RiskyServicePrincipalsConfirmCompromisedRequestBuilder) Post(ctx context.Context, body RiskyServicePrincipalsConfirmCompromisedPostRequestBodyable, requestConfiguration *RiskyServicePrincipalsConfirmCompromisedRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation confirm one or more riskyServicePrincipal objects as compromised. This action sets the targeted service principal account's risk level to `high`.
func (m *RiskyServicePrincipalsConfirmCompromisedRequestBuilder) ToPostRequestInformation(ctx context.Context, body RiskyServicePrincipalsConfirmCompromisedPostRequestBodyable, requestConfiguration *RiskyServicePrincipalsConfirmCompromisedRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
