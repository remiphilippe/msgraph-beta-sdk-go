package contacts

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemMemberOfRequestBuilder provides operations to manage the memberOf property of the microsoft.graph.orgContact entity.
type ItemMemberOfRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemMemberOfRequestBuilderGetQueryParameters groups that this contact is a member of. Read-only. Nullable. Supports $expand.
type ItemMemberOfRequestBuilderGetQueryParameters struct {
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
// ItemMemberOfRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMemberOfRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemMemberOfRequestBuilderGetQueryParameters
}
// Application casts the previous resource to application.
func (m *ItemMemberOfRequestBuilder) Application()(*ItemMemberOfApplicationRequestBuilder) {
    return NewItemMemberOfApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewItemMemberOfRequestBuilderInternal instantiates a new MemberOfRequestBuilder and sets the default values.
func NewItemMemberOfRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMemberOfRequestBuilder) {
    m := &ItemMemberOfRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/contacts/{orgContact%2Did}/memberOf{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewItemMemberOfRequestBuilder instantiates a new MemberOfRequestBuilder and sets the default values.
func NewItemMemberOfRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMemberOfRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMemberOfRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *ItemMemberOfRequestBuilder) Count()(*ItemMemberOfCountRequestBuilder) {
    return NewItemMemberOfCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Device casts the previous resource to device.
func (m *ItemMemberOfRequestBuilder) Device()(*ItemMemberOfDeviceRequestBuilder) {
    return NewItemMemberOfDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get groups that this contact is a member of. Read-only. Nullable. Supports $expand.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/orgcontact-list-memberof?view=graph-rest-1.0
func (m *ItemMemberOfRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemMemberOfRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDirectoryObjectCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable), nil
}
// Group casts the previous resource to group.
func (m *ItemMemberOfRequestBuilder) Group()(*ItemMemberOfGroupRequestBuilder) {
    return NewItemMemberOfGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OrgContact casts the previous resource to orgContact.
func (m *ItemMemberOfRequestBuilder) OrgContact()(*ItemMemberOfOrgContactRequestBuilder) {
    return NewItemMemberOfOrgContactRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal casts the previous resource to servicePrincipal.
func (m *ItemMemberOfRequestBuilder) ServicePrincipal()(*ItemMemberOfServicePrincipalRequestBuilder) {
    return NewItemMemberOfServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ToGetRequestInformation groups that this contact is a member of. Read-only. Nullable. Supports $expand.
func (m *ItemMemberOfRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemMemberOfRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// User casts the previous resource to user.
func (m *ItemMemberOfRequestBuilder) User()(*ItemMemberOfUserRequestBuilder) {
    return NewItemMemberOfUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
