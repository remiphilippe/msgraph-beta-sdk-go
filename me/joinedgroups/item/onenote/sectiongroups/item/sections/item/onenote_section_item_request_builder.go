package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i01d2ceb22ae4503a7efa4f3f42be4b20150f1563ff1abdfb3145e218691d452c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/onenote/sectiongroups/item/sections/item/copytonotebook"
    i3a84f7d928c1d37380c3fa16f26be7f80a66bb39cdb6a9fd2ce42f71b629fa16 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/onenote/sectiongroups/item/sections/item/parentnotebook"
    i6e050e505bc8d2e93652aa5d0e458faf97a9419d3a49820c824d5e3e102309c2 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/onenote/sectiongroups/item/sections/item/copytosectiongroup"
    id9ebf62531fc96b30c87bc4a394dde5231b93becab88dd2dd0edaacdd20e95a7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/onenote/sectiongroups/item/sections/item/pages"
    if80d49b49f94cd162dbf0da009c23761d509d92a5c9fbbefc81e95e47605d475 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/onenote/sectiongroups/item/sections/item/parentsectiongroup"
    i6b77e4f5602b1fbbf8a15351ee900d3bfe6cab2432ca6d791653d2021c01975d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/item/onenote/sectiongroups/item/sections/item/pages/item"
)

// OnenoteSectionItemRequestBuilder provides operations to manage the sections property of the microsoft.graph.sectionGroup entity.
type OnenoteSectionItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// OnenoteSectionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnenoteSectionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// OnenoteSectionItemRequestBuilderGetQueryParameters the sections in the section group. Read-only. Nullable.
type OnenoteSectionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// OnenoteSectionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnenoteSectionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *OnenoteSectionItemRequestBuilderGetQueryParameters
}
// OnenoteSectionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OnenoteSectionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewOnenoteSectionItemRequestBuilderInternal instantiates a new OnenoteSectionItemRequestBuilder and sets the default values.
func NewOnenoteSectionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnenoteSectionItemRequestBuilder) {
    m := &OnenoteSectionItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/joinedGroups/{group%2Did}/onenote/sectionGroups/{sectionGroup%2Did}/sections/{onenoteSection%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOnenoteSectionItemRequestBuilder instantiates a new OnenoteSectionItemRequestBuilder and sets the default values.
func NewOnenoteSectionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OnenoteSectionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnenoteSectionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CopyToNotebook the copyToNotebook property
func (m *OnenoteSectionItemRequestBuilder) CopyToNotebook()(*i01d2ceb22ae4503a7efa4f3f42be4b20150f1563ff1abdfb3145e218691d452c.CopyToNotebookRequestBuilder) {
    return i01d2ceb22ae4503a7efa4f3f42be4b20150f1563ff1abdfb3145e218691d452c.NewCopyToNotebookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CopyToSectionGroup the copyToSectionGroup property
func (m *OnenoteSectionItemRequestBuilder) CopyToSectionGroup()(*i6e050e505bc8d2e93652aa5d0e458faf97a9419d3a49820c824d5e3e102309c2.CopyToSectionGroupRequestBuilder) {
    return i6e050e505bc8d2e93652aa5d0e458faf97a9419d3a49820c824d5e3e102309c2.NewCopyToSectionGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property sections for me
func (m *OnenoteSectionItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property sections for me
func (m *OnenoteSectionItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *OnenoteSectionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the sections in the section group. Read-only. Nullable.
func (m *OnenoteSectionItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the sections in the section group. Read-only. Nullable.
func (m *OnenoteSectionItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *OnenoteSectionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property sections in me
func (m *OnenoteSectionItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnenoteSectionable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property sections in me
func (m *OnenoteSectionItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnenoteSectionable, requestConfiguration *OnenoteSectionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property sections for me
func (m *OnenoteSectionItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property sections for me
func (m *OnenoteSectionItemRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *OnenoteSectionItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the sections in the section group. Read-only. Nullable.
func (m *OnenoteSectionItemRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnenoteSectionable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler the sections in the section group. Read-only. Nullable.
func (m *OnenoteSectionItemRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *OnenoteSectionItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnenoteSectionable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOnenoteSectionFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnenoteSectionable), nil
}
// Pages the pages property
func (m *OnenoteSectionItemRequestBuilder) Pages()(*id9ebf62531fc96b30c87bc4a394dde5231b93becab88dd2dd0edaacdd20e95a7.PagesRequestBuilder) {
    return id9ebf62531fc96b30c87bc4a394dde5231b93becab88dd2dd0edaacdd20e95a7.NewPagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PagesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.joinedGroups.item.onenote.sectionGroups.item.sections.item.pages.item collection
func (m *OnenoteSectionItemRequestBuilder) PagesById(id string)(*i6b77e4f5602b1fbbf8a15351ee900d3bfe6cab2432ca6d791653d2021c01975d.OnenotePageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenotePage%2Did"] = id
    }
    return i6b77e4f5602b1fbbf8a15351ee900d3bfe6cab2432ca6d791653d2021c01975d.NewOnenotePageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ParentNotebook the parentNotebook property
func (m *OnenoteSectionItemRequestBuilder) ParentNotebook()(*i3a84f7d928c1d37380c3fa16f26be7f80a66bb39cdb6a9fd2ce42f71b629fa16.ParentNotebookRequestBuilder) {
    return i3a84f7d928c1d37380c3fa16f26be7f80a66bb39cdb6a9fd2ce42f71b629fa16.NewParentNotebookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ParentSectionGroup the parentSectionGroup property
func (m *OnenoteSectionItemRequestBuilder) ParentSectionGroup()(*if80d49b49f94cd162dbf0da009c23761d509d92a5c9fbbefc81e95e47605d475.ParentSectionGroupRequestBuilder) {
    return if80d49b49f94cd162dbf0da009c23761d509d92a5c9fbbefc81e95e47605d475.NewParentSectionGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property sections in me
func (m *OnenoteSectionItemRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnenoteSectionable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property sections in me
func (m *OnenoteSectionItemRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnenoteSectionable, requestConfiguration *OnenoteSectionItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
