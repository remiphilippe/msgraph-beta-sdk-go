package app

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CallsItemRejectRequestBuilder provides operations to call the reject method.
type CallsItemRejectRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CallsItemRejectRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallsItemRejectRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCallsItemRejectRequestBuilderInternal instantiates a new RejectRequestBuilder and sets the default values.
func NewCallsItemRejectRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallsItemRejectRequestBuilder) {
    m := &CallsItemRejectRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/app/calls/{call%2Did}/microsoft.graph.reject";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCallsItemRejectRequestBuilder instantiates a new RejectRequestBuilder and sets the default values.
func NewCallsItemRejectRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallsItemRejectRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCallsItemRejectRequestBuilderInternal(urlParams, requestAdapter)
}
// Post enable a bot to reject an incoming call. The incoming call request can be an invite from a participant in a group call or a peer-to-peer call. If an invite to a group call is received, the notification will contain the **chatInfo** and **meetingInfo** parameters. The bot is expected to answer or reject the call before the call times out. The current timeout value is 15 seconds. This API does not end existing calls that have already been answered. Use delete call to end a call.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/call-reject?view=graph-rest-1.0
func (m *CallsItemRejectRequestBuilder) Post(ctx context.Context, body CallsItemRejectPostRequestBodyable, requestConfiguration *CallsItemRejectRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation enable a bot to reject an incoming call. The incoming call request can be an invite from a participant in a group call or a peer-to-peer call. If an invite to a group call is received, the notification will contain the **chatInfo** and **meetingInfo** parameters. The bot is expected to answer or reject the call before the call times out. The current timeout value is 15 seconds. This API does not end existing calls that have already been answered. Use delete call to end a call.
func (m *CallsItemRejectRequestBuilder) ToPostRequestInformation(ctx context.Context, body CallsItemRejectPostRequestBodyable, requestConfiguration *CallsItemRejectRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
