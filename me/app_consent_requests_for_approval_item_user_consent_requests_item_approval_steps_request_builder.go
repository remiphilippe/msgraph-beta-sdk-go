package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AppConsentRequestsForApprovalItemUserConsentRequestsItemApprovalStepsRequestBuilder provides operations to manage the steps property of the microsoft.graph.approval entity.
type AppConsentRequestsForApprovalItemUserConsentRequestsItemApprovalStepsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AppConsentRequestsForApprovalItemUserConsentRequestsItemApprovalStepsRequestBuilderGetQueryParameters in Azure AD entitlement management, lists the approvalStep objects associated with an approval object.  This call can be made by an approver, providing the identifier of the access package assignment request.
type AppConsentRequestsForApprovalItemUserConsentRequestsItemApprovalStepsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// AppConsentRequestsForApprovalItemUserConsentRequestsItemApprovalStepsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AppConsentRequestsForApprovalItemUserConsentRequestsItemApprovalStepsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AppConsentRequestsForApprovalItemUserConsentRequestsItemApprovalStepsRequestBuilderGetQueryParameters
}
// AppConsentRequestsForApprovalItemUserConsentRequestsItemApprovalStepsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AppConsentRequestsForApprovalItemUserConsentRequestsItemApprovalStepsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAppConsentRequestsForApprovalItemUserConsentRequestsItemApprovalStepsRequestBuilderInternal instantiates a new StepsRequestBuilder and sets the default values.
func NewAppConsentRequestsForApprovalItemUserConsentRequestsItemApprovalStepsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AppConsentRequestsForApprovalItemUserConsentRequestsItemApprovalStepsRequestBuilder) {
    m := &AppConsentRequestsForApprovalItemUserConsentRequestsItemApprovalStepsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/appConsentRequestsForApproval/{appConsentRequest%2Did}/userConsentRequests/{userConsentRequest%2Did}/approval/steps{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAppConsentRequestsForApprovalItemUserConsentRequestsItemApprovalStepsRequestBuilder instantiates a new StepsRequestBuilder and sets the default values.
func NewAppConsentRequestsForApprovalItemUserConsentRequestsItemApprovalStepsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AppConsentRequestsForApprovalItemUserConsentRequestsItemApprovalStepsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAppConsentRequestsForApprovalItemUserConsentRequestsItemApprovalStepsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *AppConsentRequestsForApprovalItemUserConsentRequestsItemApprovalStepsRequestBuilder) Count()(*AppConsentRequestsForApprovalItemUserConsentRequestsItemApprovalStepsCountRequestBuilder) {
    return NewAppConsentRequestsForApprovalItemUserConsentRequestsItemApprovalStepsCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get in Azure AD entitlement management, lists the approvalStep objects associated with an approval object.  This call can be made by an approver, providing the identifier of the access package assignment request.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/approval-list-steps?view=graph-rest-1.0
func (m *AppConsentRequestsForApprovalItemUserConsentRequestsItemApprovalStepsRequestBuilder) Get(ctx context.Context, requestConfiguration *AppConsentRequestsForApprovalItemUserConsentRequestsItemApprovalStepsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ApprovalStepCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateApprovalStepCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ApprovalStepCollectionResponseable), nil
}
// Post create new navigation property to steps for me
func (m *AppConsentRequestsForApprovalItemUserConsentRequestsItemApprovalStepsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ApprovalStepable, requestConfiguration *AppConsentRequestsForApprovalItemUserConsentRequestsItemApprovalStepsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ApprovalStepable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateApprovalStepFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ApprovalStepable), nil
}
// ToGetRequestInformation in Azure AD entitlement management, lists the approvalStep objects associated with an approval object.  This call can be made by an approver, providing the identifier of the access package assignment request.
func (m *AppConsentRequestsForApprovalItemUserConsentRequestsItemApprovalStepsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AppConsentRequestsForApprovalItemUserConsentRequestsItemApprovalStepsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to steps for me
func (m *AppConsentRequestsForApprovalItemUserConsentRequestsItemApprovalStepsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ApprovalStepable, requestConfiguration *AppConsentRequestsForApprovalItemUserConsentRequestsItemApprovalStepsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
