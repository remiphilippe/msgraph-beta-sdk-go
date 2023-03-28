package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemDevicesDeviceItemRequestBuilder provides operations to manage the devices property of the microsoft.graph.user entity.
type ItemDevicesDeviceItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemDevicesDeviceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemDevicesDeviceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemDevicesDeviceItemRequestBuilderGetQueryParameters get devices from users
type ItemDevicesDeviceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemDevicesDeviceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemDevicesDeviceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemDevicesDeviceItemRequestBuilderGetQueryParameters
}
// ItemDevicesDeviceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemDevicesDeviceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CheckMemberGroups provides operations to call the checkMemberGroups method.
func (m *ItemDevicesDeviceItemRequestBuilder) CheckMemberGroups()(*ItemDevicesItemCheckMemberGroupsRequestBuilder) {
    return NewItemDevicesItemCheckMemberGroupsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CheckMemberObjects provides operations to call the checkMemberObjects method.
func (m *ItemDevicesDeviceItemRequestBuilder) CheckMemberObjects()(*ItemDevicesItemCheckMemberObjectsRequestBuilder) {
    return NewItemDevicesItemCheckMemberObjectsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Commands provides operations to manage the commands property of the microsoft.graph.device entity.
func (m *ItemDevicesDeviceItemRequestBuilder) Commands()(*ItemDevicesItemCommandsRequestBuilder) {
    return NewItemDevicesItemCommandsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CommandsById provides operations to manage the commands property of the microsoft.graph.device entity.
func (m *ItemDevicesDeviceItemRequestBuilder) CommandsById(id string)(*ItemDevicesItemCommandsCommandItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["command%2Did"] = id
    }
    return NewItemDevicesItemCommandsCommandItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemDevicesDeviceItemRequestBuilderInternal instantiates a new DeviceItemRequestBuilder and sets the default values.
func NewItemDevicesDeviceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDevicesDeviceItemRequestBuilder) {
    m := &ItemDevicesDeviceItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/devices/{device%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewItemDevicesDeviceItemRequestBuilder instantiates a new DeviceItemRequestBuilder and sets the default values.
func NewItemDevicesDeviceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDevicesDeviceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemDevicesDeviceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property devices for users
func (m *ItemDevicesDeviceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemDevicesDeviceItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
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
// Extensions provides operations to manage the extensions property of the microsoft.graph.device entity.
func (m *ItemDevicesDeviceItemRequestBuilder) Extensions()(*ItemDevicesItemExtensionsRequestBuilder) {
    return NewItemDevicesItemExtensionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.device entity.
func (m *ItemDevicesDeviceItemRequestBuilder) ExtensionsById(id string)(*ItemDevicesItemExtensionsExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return NewItemDevicesItemExtensionsExtensionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// Get get devices from users
func (m *ItemDevicesDeviceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemDevicesDeviceItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Deviceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Deviceable), nil
}
// GetMemberGroups provides operations to call the getMemberGroups method.
func (m *ItemDevicesDeviceItemRequestBuilder) GetMemberGroups()(*ItemDevicesItemGetMemberGroupsRequestBuilder) {
    return NewItemDevicesItemGetMemberGroupsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetMemberObjects provides operations to call the getMemberObjects method.
func (m *ItemDevicesDeviceItemRequestBuilder) GetMemberObjects()(*ItemDevicesItemGetMemberObjectsRequestBuilder) {
    return NewItemDevicesItemGetMemberObjectsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MemberOf provides operations to manage the memberOf property of the microsoft.graph.device entity.
func (m *ItemDevicesDeviceItemRequestBuilder) MemberOf()(*ItemDevicesItemMemberOfRequestBuilder) {
    return NewItemDevicesItemMemberOfRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MemberOfById provides operations to manage the memberOf property of the microsoft.graph.device entity.
func (m *ItemDevicesDeviceItemRequestBuilder) MemberOfById(id string)(*ItemDevicesItemMemberOfDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewItemDevicesItemMemberOfDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property devices in users
func (m *ItemDevicesDeviceItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Deviceable, requestConfiguration *ItemDevicesDeviceItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Deviceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Deviceable), nil
}
// RegisteredOwners provides operations to manage the registeredOwners property of the microsoft.graph.device entity.
func (m *ItemDevicesDeviceItemRequestBuilder) RegisteredOwners()(*ItemDevicesItemRegisteredOwnersRequestBuilder) {
    return NewItemDevicesItemRegisteredOwnersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RegisteredOwnersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.devices.item.registeredOwners.item collection
func (m *ItemDevicesDeviceItemRequestBuilder) RegisteredOwnersById(id string)(*ItemDevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewItemDevicesItemRegisteredOwnersDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// RegisteredUsers provides operations to manage the registeredUsers property of the microsoft.graph.device entity.
func (m *ItemDevicesDeviceItemRequestBuilder) RegisteredUsers()(*ItemDevicesItemRegisteredUsersRequestBuilder) {
    return NewItemDevicesItemRegisteredUsersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RegisteredUsersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.devices.item.registeredUsers.item collection
func (m *ItemDevicesDeviceItemRequestBuilder) RegisteredUsersById(id string)(*ItemDevicesItemRegisteredUsersDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewItemDevicesItemRegisteredUsersDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// Restore provides operations to call the restore method.
func (m *ItemDevicesDeviceItemRequestBuilder) Restore()(*ItemDevicesItemRestoreRequestBuilder) {
    return NewItemDevicesItemRestoreRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property devices for users
func (m *ItemDevicesDeviceItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemDevicesDeviceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation get devices from users
func (m *ItemDevicesDeviceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemDevicesDeviceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property devices in users
func (m *ItemDevicesDeviceItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Deviceable, requestConfiguration *ItemDevicesDeviceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
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
// TransitiveMemberOf provides operations to manage the transitiveMemberOf property of the microsoft.graph.device entity.
func (m *ItemDevicesDeviceItemRequestBuilder) TransitiveMemberOf()(*ItemDevicesItemTransitiveMemberOfRequestBuilder) {
    return NewItemDevicesItemTransitiveMemberOfRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TransitiveMemberOfById provides operations to manage the transitiveMemberOf property of the microsoft.graph.device entity.
func (m *ItemDevicesDeviceItemRequestBuilder) TransitiveMemberOfById(id string)(*ItemDevicesItemTransitiveMemberOfDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewItemDevicesItemTransitiveMemberOfDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// UsageRights provides operations to manage the usageRights property of the microsoft.graph.device entity.
func (m *ItemDevicesDeviceItemRequestBuilder) UsageRights()(*ItemDevicesItemUsageRightsRequestBuilder) {
    return NewItemDevicesItemUsageRightsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UsageRightsById provides operations to manage the usageRights property of the microsoft.graph.device entity.
func (m *ItemDevicesDeviceItemRequestBuilder) UsageRightsById(id string)(*ItemDevicesItemUsageRightsUsageRightItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["usageRight%2Did"] = id
    }
    return NewItemDevicesItemUsageRightsUsageRightItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
