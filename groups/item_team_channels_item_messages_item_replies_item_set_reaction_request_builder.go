package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemTeamChannelsItemMessagesItemRepliesItemSetReactionRequestBuilder provides operations to call the setReaction method.
type ItemTeamChannelsItemMessagesItemRepliesItemSetReactionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTeamChannelsItemMessagesItemRepliesItemSetReactionRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamChannelsItemMessagesItemRepliesItemSetReactionRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemTeamChannelsItemMessagesItemRepliesItemSetReactionRequestBuilderInternal instantiates a new SetReactionRequestBuilder and sets the default values.
func NewItemTeamChannelsItemMessagesItemRepliesItemSetReactionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamChannelsItemMessagesItemRepliesItemSetReactionRequestBuilder) {
    m := &ItemTeamChannelsItemMessagesItemRepliesItemSetReactionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/team/channels/{channel%2Did}/messages/{chatMessage%2Did}/replies/{chatMessage%2Did1}/setReaction", pathParameters),
    }
    return m
}
// NewItemTeamChannelsItemMessagesItemRepliesItemSetReactionRequestBuilder instantiates a new SetReactionRequestBuilder and sets the default values.
func NewItemTeamChannelsItemMessagesItemRepliesItemSetReactionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamChannelsItemMessagesItemRepliesItemSetReactionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamChannelsItemMessagesItemRepliesItemSetReactionRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action setReaction
func (m *ItemTeamChannelsItemMessagesItemRepliesItemSetReactionRequestBuilder) Post(ctx context.Context, body ItemTeamChannelsItemMessagesItemRepliesItemSetReactionPostRequestBodyable, requestConfiguration *ItemTeamChannelsItemMessagesItemRepliesItemSetReactionRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation invoke action setReaction
func (m *ItemTeamChannelsItemMessagesItemRepliesItemSetReactionRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemTeamChannelsItemMessagesItemRepliesItemSetReactionPostRequestBodyable, requestConfiguration *ItemTeamChannelsItemMessagesItemRepliesItemSetReactionRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
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
