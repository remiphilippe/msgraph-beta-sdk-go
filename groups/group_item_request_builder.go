package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GroupItemRequestBuilder provides operations to manage the collection of group entities.
type GroupItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GroupItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// GroupItemRequestBuilderGetQueryParameters get the properties and relationships of a group object. This operation returns by default only a subset of all the available properties, as noted in the Properties section. To get properties that are _not_ returned by default, specify them in a `$select` OData query option. The **hasMembersWithLicenseErrors** and **isArchived** properties are an exception and are not returned in the `$select` query. Because the **group** resource supports extensions, you can also use the `GET` operation to get custom properties and extension data in a **group** instance.
type GroupItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// GroupItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GroupItemRequestBuilderGetQueryParameters
}
// GroupItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AcceptedSenders provides operations to manage the acceptedSenders property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) AcceptedSenders()(*ItemAcceptedSendersRequestBuilder) {
    return NewItemAcceptedSendersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AcceptedSendersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.acceptedSenders.item collection
func (m *GroupItemRequestBuilder) AcceptedSendersById(id string)(*ItemAcceptedSendersDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewItemAcceptedSendersDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AddFavorite provides operations to call the addFavorite method.
func (m *GroupItemRequestBuilder) AddFavorite()(*ItemAddFavoriteRequestBuilder) {
    return NewItemAddFavoriteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppRoleAssignments provides operations to manage the appRoleAssignments property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) AppRoleAssignments()(*ItemAppRoleAssignmentsRequestBuilder) {
    return NewItemAppRoleAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AppRoleAssignmentsById provides operations to manage the appRoleAssignments property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) AppRoleAssignmentsById(id string)(*ItemAppRoleAssignmentsAppRoleAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["appRoleAssignment%2Did"] = id
    }
    return NewItemAppRoleAssignmentsAppRoleAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AssignLicense provides operations to call the assignLicense method.
func (m *GroupItemRequestBuilder) AssignLicense()(*ItemAssignLicenseRequestBuilder) {
    return NewItemAssignLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) Calendar()(*ItemCalendarRequestBuilder) {
    return NewItemCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarView provides operations to manage the calendarView property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) CalendarView()(*ItemCalendarViewRequestBuilder) {
    return NewItemCalendarViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarViewById provides operations to manage the calendarView property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) CalendarViewById(id string)(*ItemCalendarViewEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did"] = id
    }
    return NewItemCalendarViewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CheckGrantedPermissionsForApp provides operations to call the checkGrantedPermissionsForApp method.
func (m *GroupItemRequestBuilder) CheckGrantedPermissionsForApp()(*ItemCheckGrantedPermissionsForAppRequestBuilder) {
    return NewItemCheckGrantedPermissionsForAppRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberGroups provides operations to call the checkMemberGroups method.
func (m *GroupItemRequestBuilder) CheckMemberGroups()(*ItemCheckMemberGroupsRequestBuilder) {
    return NewItemCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects provides operations to call the checkMemberObjects method.
func (m *GroupItemRequestBuilder) CheckMemberObjects()(*ItemCheckMemberObjectsRequestBuilder) {
    return NewItemCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewGroupItemRequestBuilderInternal instantiates a new GroupItemRequestBuilder and sets the default values.
func NewGroupItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupItemRequestBuilder) {
    m := &GroupItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGroupItemRequestBuilder instantiates a new GroupItemRequestBuilder and sets the default values.
func NewGroupItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Conversations provides operations to manage the conversations property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) Conversations()(*ItemConversationsRequestBuilder) {
    return NewItemConversationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ConversationsById provides operations to manage the conversations property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) ConversationsById(id string)(*ItemConversationsConversationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversation%2Did"] = id
    }
    return NewItemConversationsConversationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreatedOnBehalfOf provides operations to manage the createdOnBehalfOf property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) CreatedOnBehalfOf()(*ItemCreatedOnBehalfOfRequestBuilder) {
    return NewItemCreatedOnBehalfOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete deletes a group. When deleted, Microsoft 365 groups are moved to a temporary container and can be restored within 30 days. After that time, they are permanently deleted. This isn't applicable to Security groups and Distribution groups which are permanently deleted immediately. To learn more, see deletedItems.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/group-delete?view=graph-rest-1.0
func (m *GroupItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *GroupItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
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
// Drive provides operations to manage the drive property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) Drive()(*ItemDriveRequestBuilder) {
    return NewItemDriveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Drives provides operations to manage the drives property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) Drives()(*ItemDrivesRequestBuilder) {
    return NewItemDrivesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DrivesById provides operations to manage the drives property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) DrivesById(id string)(*ItemDrivesDriveItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["drive%2Did"] = id
    }
    return NewItemDrivesDriveItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Endpoints provides operations to manage the endpoints property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) Endpoints()(*ItemEndpointsRequestBuilder) {
    return NewItemEndpointsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EndpointsById provides operations to manage the endpoints property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) EndpointsById(id string)(*ItemEndpointsEndpointItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["endpoint%2Did"] = id
    }
    return NewItemEndpointsEndpointItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// EvaluateDynamicMembership provides operations to call the evaluateDynamicMembership method.
func (m *GroupItemRequestBuilder) EvaluateDynamicMembership()(*ItemEvaluateDynamicMembershipRequestBuilder) {
    return NewItemEvaluateDynamicMembershipRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Events provides operations to manage the events property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) Events()(*ItemEventsRequestBuilder) {
    return NewItemEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EventsById provides operations to manage the events property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) EventsById(id string)(*ItemEventsEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did"] = id
    }
    return NewItemEventsEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) Extensions()(*ItemExtensionsRequestBuilder) {
    return NewItemExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) ExtensionsById(id string)(*ItemExtensionsExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return NewItemExtensionsExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get the properties and relationships of a group object. This operation returns by default only a subset of all the available properties, as noted in the Properties section. To get properties that are _not_ returned by default, specify them in a `$select` OData query option. The **hasMembersWithLicenseErrors** and **isArchived** properties are an exception and are not returned in the `$select` query. Because the **group** resource supports extensions, you can also use the `GET` operation to get custom properties and extension data in a **group** instance.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/group-get?view=graph-rest-1.0
func (m *GroupItemRequestBuilder) Get(ctx context.Context, requestConfiguration *GroupItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Groupable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Groupable), nil
}
// GetMemberGroups provides operations to call the getMemberGroups method.
func (m *GroupItemRequestBuilder) GetMemberGroups()(*ItemGetMemberGroupsRequestBuilder) {
    return NewItemGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects provides operations to call the getMemberObjects method.
func (m *GroupItemRequestBuilder) GetMemberObjects()(*ItemGetMemberObjectsRequestBuilder) {
    return NewItemGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GroupLifecyclePolicies provides operations to manage the groupLifecyclePolicies property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) GroupLifecyclePolicies()(*ItemGroupLifecyclePoliciesRequestBuilder) {
    return NewItemGroupLifecyclePoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GroupLifecyclePoliciesById provides operations to manage the groupLifecyclePolicies property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) GroupLifecyclePoliciesById(id string)(*ItemGroupLifecyclePoliciesGroupLifecyclePolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["groupLifecyclePolicy%2Did"] = id
    }
    return NewItemGroupLifecyclePoliciesGroupLifecyclePolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MemberOf provides operations to manage the memberOf property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) MemberOf()(*ItemMemberOfRequestBuilder) {
    return NewItemMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MemberOfById provides operations to manage the memberOf property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) MemberOfById(id string)(*ItemMemberOfDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewItemMemberOfDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Members provides operations to manage the members property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) Members()(*ItemMembersRequestBuilder) {
    return NewItemMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.members.item collection
func (m *GroupItemRequestBuilder) MembersById(id string)(*ItemMembersDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewItemMembersDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MembersWithLicenseErrors provides operations to manage the membersWithLicenseErrors property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) MembersWithLicenseErrors()(*ItemMembersWithLicenseErrorsRequestBuilder) {
    return NewItemMembersWithLicenseErrorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersWithLicenseErrorsById provides operations to manage the membersWithLicenseErrors property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) MembersWithLicenseErrorsById(id string)(*ItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewItemMembersWithLicenseErrorsDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Onenote provides operations to manage the onenote property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) Onenote()(*ItemOnenoteRequestBuilder) {
    return NewItemOnenoteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Owners provides operations to manage the owners property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) Owners()(*ItemOwnersRequestBuilder) {
    return NewItemOwnersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OwnersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.owners.item collection
func (m *GroupItemRequestBuilder) OwnersById(id string)(*ItemOwnersDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewItemOwnersDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the properties of a group object.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/group-update?view=graph-rest-1.0
func (m *GroupItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Groupable, requestConfiguration *GroupItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Groupable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Groupable), nil
}
// PermissionGrants provides operations to manage the permissionGrants property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) PermissionGrants()(*ItemPermissionGrantsRequestBuilder) {
    return NewItemPermissionGrantsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionGrantsById provides operations to manage the permissionGrants property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) PermissionGrantsById(id string)(*ItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["resourceSpecificPermissionGrant%2Did"] = id
    }
    return NewItemPermissionGrantsResourceSpecificPermissionGrantItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Photo provides operations to manage the photo property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) Photo()(*ItemPhotoRequestBuilder) {
    return NewItemPhotoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Photos provides operations to manage the photos property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) Photos()(*ItemPhotosRequestBuilder) {
    return NewItemPhotosRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PhotosById provides operations to manage the photos property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) PhotosById(id string)(*ItemPhotosProfilePhotoItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["profilePhoto%2Did"] = id
    }
    return NewItemPhotosProfilePhotoItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Planner provides operations to manage the planner property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) Planner()(*ItemPlannerRequestBuilder) {
    return NewItemPlannerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RejectedSenders provides operations to manage the rejectedSenders property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) RejectedSenders()(*ItemRejectedSendersRequestBuilder) {
    return NewItemRejectedSendersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RejectedSendersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.rejectedSenders.item collection
func (m *GroupItemRequestBuilder) RejectedSendersById(id string)(*ItemRejectedSendersDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewItemRejectedSendersDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RemoveFavorite provides operations to call the removeFavorite method.
func (m *GroupItemRequestBuilder) RemoveFavorite()(*ItemRemoveFavoriteRequestBuilder) {
    return NewItemRemoveFavoriteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Renew provides operations to call the renew method.
func (m *GroupItemRequestBuilder) Renew()(*ItemRenewRequestBuilder) {
    return NewItemRenewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResetUnseenCount provides operations to call the resetUnseenCount method.
func (m *GroupItemRequestBuilder) ResetUnseenCount()(*ItemResetUnseenCountRequestBuilder) {
    return NewItemResetUnseenCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore provides operations to call the restore method.
func (m *GroupItemRequestBuilder) Restore()(*ItemRestoreRequestBuilder) {
    return NewItemRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Settings provides operations to manage the settings property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) Settings()(*ItemSettingsRequestBuilder) {
    return NewItemSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SettingsById provides operations to manage the settings property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) SettingsById(id string)(*ItemSettingsDirectorySettingItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directorySetting%2Did"] = id
    }
    return NewItemSettingsDirectorySettingItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Sites provides operations to manage the sites property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) Sites()(*ItemSitesRequestBuilder) {
    return NewItemSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SitesById provides operations to manage the sites property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) SitesById(id string)(*ItemSitesSiteItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["site%2Did"] = id
    }
    return NewItemSitesSiteItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SubscribeByMail provides operations to call the subscribeByMail method.
func (m *GroupItemRequestBuilder) SubscribeByMail()(*ItemSubscribeByMailRequestBuilder) {
    return NewItemSubscribeByMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Team provides operations to manage the team property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) Team()(*ItemTeamRequestBuilder) {
    return NewItemTeamRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Threads provides operations to manage the threads property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) Threads()(*ItemThreadsRequestBuilder) {
    return NewItemThreadsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThreadsById provides operations to manage the threads property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) ThreadsById(id string)(*ItemThreadsConversationThreadItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationThread%2Did"] = id
    }
    return NewItemThreadsConversationThreadItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ToDeleteRequestInformation deletes a group. When deleted, Microsoft 365 groups are moved to a temporary container and can be restored within 30 days. After that time, they are permanently deleted. This isn't applicable to Security groups and Distribution groups which are permanently deleted immediately. To learn more, see deletedItems.
func (m *GroupItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *GroupItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation get the properties and relationships of a group object. This operation returns by default only a subset of all the available properties, as noted in the Properties section. To get properties that are _not_ returned by default, specify them in a `$select` OData query option. The **hasMembersWithLicenseErrors** and **isArchived** properties are an exception and are not returned in the `$select` query. Because the **group** resource supports extensions, you can also use the `GET` operation to get custom properties and extension data in a **group** instance.
func (m *GroupItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GroupItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a group object.
func (m *GroupItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Groupable, requestConfiguration *GroupItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// TransitiveMemberOf provides operations to manage the transitiveMemberOf property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) TransitiveMemberOf()(*ItemTransitiveMemberOfRequestBuilder) {
    return NewItemTransitiveMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TransitiveMemberOfById provides operations to manage the transitiveMemberOf property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) TransitiveMemberOfById(id string)(*ItemTransitiveMemberOfDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewItemTransitiveMemberOfDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TransitiveMembers provides operations to manage the transitiveMembers property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) TransitiveMembers()(*ItemTransitiveMembersRequestBuilder) {
    return NewItemTransitiveMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TransitiveMembersById provides operations to manage the transitiveMembers property of the microsoft.graph.group entity.
func (m *GroupItemRequestBuilder) TransitiveMembersById(id string)(*ItemTransitiveMembersDirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject%2Did"] = id
    }
    return NewItemTransitiveMembersDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UnsubscribeByMail provides operations to call the unsubscribeByMail method.
func (m *GroupItemRequestBuilder) UnsubscribeByMail()(*ItemUnsubscribeByMailRequestBuilder) {
    return NewItemUnsubscribeByMailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ValidateProperties provides operations to call the validateProperties method.
func (m *GroupItemRequestBuilder) ValidateProperties()(*ItemValidatePropertiesRequestBuilder) {
    return NewItemValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
