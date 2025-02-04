package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemThreadsItemPostsItemInReplyToRequestBuilder provides operations to manage the inReplyTo property of the microsoft.graph.post entity.
type ItemThreadsItemPostsItemInReplyToRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemThreadsItemPostsItemInReplyToRequestBuilderGetQueryParameters the earlier post that this post is replying to in the conversationThread. Read-only. Supports $expand.
type ItemThreadsItemPostsItemInReplyToRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemThreadsItemPostsItemInReplyToRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemThreadsItemPostsItemInReplyToRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemThreadsItemPostsItemInReplyToRequestBuilderGetQueryParameters
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.post entity.
func (m *ItemThreadsItemPostsItemInReplyToRequestBuilder) Attachments()(*ItemThreadsItemPostsItemInReplyToAttachmentsRequestBuilder) {
    return NewItemThreadsItemPostsItemInReplyToAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById provides operations to manage the attachments property of the microsoft.graph.post entity.
func (m *ItemThreadsItemPostsItemInReplyToRequestBuilder) AttachmentsById(id string)(*ItemThreadsItemPostsItemInReplyToAttachmentsAttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return NewItemThreadsItemPostsItemInReplyToAttachmentsAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewItemThreadsItemPostsItemInReplyToRequestBuilderInternal instantiates a new InReplyToRequestBuilder and sets the default values.
func NewItemThreadsItemPostsItemInReplyToRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemThreadsItemPostsItemInReplyToRequestBuilder) {
    m := &ItemThreadsItemPostsItemInReplyToRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/threads/{conversationThread%2Did}/posts/{post%2Did}/inReplyTo{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewItemThreadsItemPostsItemInReplyToRequestBuilder instantiates a new InReplyToRequestBuilder and sets the default values.
func NewItemThreadsItemPostsItemInReplyToRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemThreadsItemPostsItemInReplyToRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemThreadsItemPostsItemInReplyToRequestBuilderInternal(urlParams, requestAdapter)
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.post entity.
func (m *ItemThreadsItemPostsItemInReplyToRequestBuilder) Extensions()(*ItemThreadsItemPostsItemInReplyToExtensionsRequestBuilder) {
    return NewItemThreadsItemPostsItemInReplyToExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.post entity.
func (m *ItemThreadsItemPostsItemInReplyToRequestBuilder) ExtensionsById(id string)(*ItemThreadsItemPostsItemInReplyToExtensionsExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return NewItemThreadsItemPostsItemInReplyToExtensionsExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward provides operations to call the forward method.
func (m *ItemThreadsItemPostsItemInReplyToRequestBuilder) Forward()(*ItemThreadsItemPostsItemInReplyToForwardRequestBuilder) {
    return NewItemThreadsItemPostsItemInReplyToForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the earlier post that this post is replying to in the conversationThread. Read-only. Supports $expand.
func (m *ItemThreadsItemPostsItemInReplyToRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemThreadsItemPostsItemInReplyToRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Postable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePostFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Postable), nil
}
// Mentions provides operations to manage the mentions property of the microsoft.graph.post entity.
func (m *ItemThreadsItemPostsItemInReplyToRequestBuilder) Mentions()(*ItemThreadsItemPostsItemInReplyToMentionsRequestBuilder) {
    return NewItemThreadsItemPostsItemInReplyToMentionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MentionsById provides operations to manage the mentions property of the microsoft.graph.post entity.
func (m *ItemThreadsItemPostsItemInReplyToRequestBuilder) MentionsById(id string)(*ItemThreadsItemPostsItemInReplyToMentionsMentionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mention%2Did"] = id
    }
    return NewItemThreadsItemPostsItemInReplyToMentionsMentionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MultiValueExtendedProperties provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.post entity.
func (m *ItemThreadsItemPostsItemInReplyToRequestBuilder) MultiValueExtendedProperties()(*ItemThreadsItemPostsItemInReplyToMultiValueExtendedPropertiesRequestBuilder) {
    return NewItemThreadsItemPostsItemInReplyToMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.post entity.
func (m *ItemThreadsItemPostsItemInReplyToRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ItemThreadsItemPostsItemInReplyToMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return NewItemThreadsItemPostsItemInReplyToMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Reply provides operations to call the reply method.
func (m *ItemThreadsItemPostsItemInReplyToRequestBuilder) Reply()(*ItemThreadsItemPostsItemInReplyToReplyRequestBuilder) {
    return NewItemThreadsItemPostsItemInReplyToReplyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedProperties provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.post entity.
func (m *ItemThreadsItemPostsItemInReplyToRequestBuilder) SingleValueExtendedProperties()(*ItemThreadsItemPostsItemInReplyToSingleValueExtendedPropertiesRequestBuilder) {
    return NewItemThreadsItemPostsItemInReplyToSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.post entity.
func (m *ItemThreadsItemPostsItemInReplyToRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ItemThreadsItemPostsItemInReplyToSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return NewItemThreadsItemPostsItemInReplyToSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ToGetRequestInformation the earlier post that this post is replying to in the conversationThread. Read-only. Supports $expand.
func (m *ItemThreadsItemPostsItemInReplyToRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemThreadsItemPostsItemInReplyToRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
