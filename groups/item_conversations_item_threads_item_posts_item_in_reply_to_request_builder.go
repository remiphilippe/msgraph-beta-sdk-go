package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilder provides operations to manage the inReplyTo property of the microsoft.graph.post entity.
type ItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilderGetQueryParameters the earlier post that this post is replying to in the conversationThread. Read-only. Supports $expand.
type ItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilderGetQueryParameters
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.post entity.
func (m *ItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilder) Attachments()(*ItemConversationsItemThreadsItemPostsItemInReplyToAttachmentsRequestBuilder) {
    return NewItemConversationsItemThreadsItemPostsItemInReplyToAttachmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AttachmentsById provides operations to manage the attachments property of the microsoft.graph.post entity.
func (m *ItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilder) AttachmentsById(id string)(*ItemConversationsItemThreadsItemPostsItemInReplyToAttachmentsAttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return NewItemConversationsItemThreadsItemPostsItemInReplyToAttachmentsAttachmentItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilderInternal instantiates a new InReplyToRequestBuilder and sets the default values.
func NewItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilder) {
    m := &ItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/conversations/{conversation%2Did}/threads/{conversationThread%2Did}/posts/{post%2Did}/inReplyTo{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilder instantiates a new InReplyToRequestBuilder and sets the default values.
func NewItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilderInternal(urlParams, requestAdapter)
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.post entity.
func (m *ItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilder) Extensions()(*ItemConversationsItemThreadsItemPostsItemInReplyToExtensionsRequestBuilder) {
    return NewItemConversationsItemThreadsItemPostsItemInReplyToExtensionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.post entity.
func (m *ItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilder) ExtensionsById(id string)(*ItemConversationsItemThreadsItemPostsItemInReplyToExtensionsExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return NewItemConversationsItemThreadsItemPostsItemInReplyToExtensionsExtensionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// Forward provides operations to call the forward method.
func (m *ItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilder) Forward()(*ItemConversationsItemThreadsItemPostsItemInReplyToForwardRequestBuilder) {
    return NewItemConversationsItemThreadsItemPostsItemInReplyToForwardRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the earlier post that this post is replying to in the conversationThread. Read-only. Supports $expand.
func (m *ItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Postable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePostFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Postable), nil
}
// Mentions provides operations to manage the mentions property of the microsoft.graph.post entity.
func (m *ItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilder) Mentions()(*ItemConversationsItemThreadsItemPostsItemInReplyToMentionsRequestBuilder) {
    return NewItemConversationsItemThreadsItemPostsItemInReplyToMentionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MentionsById provides operations to manage the mentions property of the microsoft.graph.post entity.
func (m *ItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilder) MentionsById(id string)(*ItemConversationsItemThreadsItemPostsItemInReplyToMentionsMentionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mention%2Did"] = id
    }
    return NewItemConversationsItemThreadsItemPostsItemInReplyToMentionsMentionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// MultiValueExtendedProperties provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.post entity.
func (m *ItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilder) MultiValueExtendedProperties()(*ItemConversationsItemThreadsItemPostsItemInReplyToMultiValueExtendedPropertiesRequestBuilder) {
    return NewItemConversationsItemThreadsItemPostsItemInReplyToMultiValueExtendedPropertiesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.post entity.
func (m *ItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ItemConversationsItemThreadsItemPostsItemInReplyToMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return NewItemConversationsItemThreadsItemPostsItemInReplyToMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// Reply provides operations to call the reply method.
func (m *ItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilder) Reply()(*ItemConversationsItemThreadsItemPostsItemInReplyToReplyRequestBuilder) {
    return NewItemConversationsItemThreadsItemPostsItemInReplyToReplyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SingleValueExtendedProperties provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.post entity.
func (m *ItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilder) SingleValueExtendedProperties()(*ItemConversationsItemThreadsItemPostsItemInReplyToSingleValueExtendedPropertiesRequestBuilder) {
    return NewItemConversationsItemThreadsItemPostsItemInReplyToSingleValueExtendedPropertiesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.post entity.
func (m *ItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ItemConversationsItemThreadsItemPostsItemInReplyToSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return NewItemConversationsItemThreadsItemPostsItemInReplyToSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation the earlier post that this post is replying to in the conversationThread. Read-only. Supports $expand.
func (m *ItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemConversationsItemThreadsItemPostsItemInReplyToRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
