package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DrivesItemRootRequestBuilder provides operations to manage the root property of the microsoft.graph.drive entity.
type DrivesItemRootRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DrivesItemRootRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DrivesItemRootRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DrivesItemRootRequestBuilderGetQueryParameters retrieve the metadata for a driveItem in a drive by file system path or ID.
type DrivesItemRootRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DrivesItemRootRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DrivesItemRootRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DrivesItemRootRequestBuilderGetQueryParameters
}
// DrivesItemRootRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DrivesItemRootRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Activities provides operations to manage the activities property of the microsoft.graph.driveItem entity.
func (m *DrivesItemRootRequestBuilder) Activities()(*DrivesItemRootActivitiesRequestBuilder) {
    return NewDrivesItemRootActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById provides operations to manage the activities property of the microsoft.graph.driveItem entity.
func (m *DrivesItemRootRequestBuilder) ActivitiesById(id string)(*DrivesItemRootActivitiesItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD%2Did"] = id
    }
    return NewDrivesItemRootActivitiesItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Analytics provides operations to manage the analytics property of the microsoft.graph.driveItem entity.
func (m *DrivesItemRootRequestBuilder) Analytics()(*DrivesItemRootAnalyticsRequestBuilder) {
    return NewDrivesItemRootAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignSensitivityLabel provides operations to call the assignSensitivityLabel method.
func (m *DrivesItemRootRequestBuilder) AssignSensitivityLabel()(*DrivesItemRootAssignSensitivityLabelRequestBuilder) {
    return NewDrivesItemRootAssignSensitivityLabelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Checkin provides operations to call the checkin method.
func (m *DrivesItemRootRequestBuilder) Checkin()(*DrivesItemRootCheckinRequestBuilder) {
    return NewDrivesItemRootCheckinRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Checkout provides operations to call the checkout method.
func (m *DrivesItemRootRequestBuilder) Checkout()(*DrivesItemRootCheckoutRequestBuilder) {
    return NewDrivesItemRootCheckoutRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Children provides operations to manage the children property of the microsoft.graph.driveItem entity.
func (m *DrivesItemRootRequestBuilder) Children()(*DrivesItemRootChildrenRequestBuilder) {
    return NewDrivesItemRootChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById provides operations to manage the children property of the microsoft.graph.driveItem entity.
func (m *DrivesItemRootRequestBuilder) ChildrenById(id string)(*DrivesItemRootChildrenDriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did"] = id
    }
    return NewDrivesItemRootChildrenDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDrivesItemRootRequestBuilderInternal instantiates a new RootRequestBuilder and sets the default values.
func NewDrivesItemRootRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DrivesItemRootRequestBuilder) {
    m := &DrivesItemRootRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/drives/{drive%2Did}/root{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDrivesItemRootRequestBuilder instantiates a new RootRequestBuilder and sets the default values.
func NewDrivesItemRootRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DrivesItemRootRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDrivesItemRootRequestBuilderInternal(urlParams, requestAdapter)
}
// Content provides operations to manage the media for the user entity.
func (m *DrivesItemRootRequestBuilder) Content()(*DrivesItemRootContentRequestBuilder) {
    return NewDrivesItemRootContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Copy provides operations to call the copy method.
func (m *DrivesItemRootRequestBuilder) Copy()(*DrivesItemRootCopyRequestBuilder) {
    return NewDrivesItemRootCopyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateLink provides operations to call the createLink method.
func (m *DrivesItemRootRequestBuilder) CreateLink()(*DrivesItemRootCreateLinkRequestBuilder) {
    return NewDrivesItemRootCreateLinkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateUploadSession provides operations to call the createUploadSession method.
func (m *DrivesItemRootRequestBuilder) CreateUploadSession()(*DrivesItemRootCreateUploadSessionRequestBuilder) {
    return NewDrivesItemRootCreateUploadSessionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property root for me
func (m *DrivesItemRootRequestBuilder) Delete(ctx context.Context, requestConfiguration *DrivesItemRootRequestBuilderDeleteRequestConfiguration)(error) {
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
// Delta provides operations to call the delta method.
func (m *DrivesItemRootRequestBuilder) Delta()(*DrivesItemRootDeltaRequestBuilder) {
    return NewDrivesItemRootDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeltaWithToken provides operations to call the delta method.
func (m *DrivesItemRootRequestBuilder) DeltaWithToken(token *string)(*DrivesItemRootDeltaWithTokenRequestBuilder) {
    return NewDrivesItemRootDeltaWithTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter, token);
}
// ExtractSensitivityLabels provides operations to call the extractSensitivityLabels method.
func (m *DrivesItemRootRequestBuilder) ExtractSensitivityLabels()(*DrivesItemRootExtractSensitivityLabelsRequestBuilder) {
    return NewDrivesItemRootExtractSensitivityLabelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Follow provides operations to call the follow method.
func (m *DrivesItemRootRequestBuilder) Follow()(*DrivesItemRootFollowRequestBuilder) {
    return NewDrivesItemRootFollowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get retrieve the metadata for a driveItem in a drive by file system path or ID.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/driveitem-get?view=graph-rest-1.0
func (m *DrivesItemRootRequestBuilder) Get(ctx context.Context, requestConfiguration *DrivesItemRootRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDriveItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable), nil
}
// GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval provides operations to call the getActivitiesByInterval method.
func (m *DrivesItemRootRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(endDateTime *string, interval *string, startDateTime *string)(*DrivesItemRootGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return NewDrivesItemRootGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, interval, startDateTime);
}
// Invite provides operations to call the invite method.
func (m *DrivesItemRootRequestBuilder) Invite()(*DrivesItemRootInviteRequestBuilder) {
    return NewDrivesItemRootInviteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ListItem provides operations to manage the listItem property of the microsoft.graph.driveItem entity.
func (m *DrivesItemRootRequestBuilder) ListItem()(*DrivesItemRootListItemRequestBuilder) {
    return NewDrivesItemRootListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property root in me
func (m *DrivesItemRootRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, requestConfiguration *DrivesItemRootRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDriveItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable), nil
}
// Permissions provides operations to manage the permissions property of the microsoft.graph.driveItem entity.
func (m *DrivesItemRootRequestBuilder) Permissions()(*DrivesItemRootPermissionsRequestBuilder) {
    return NewDrivesItemRootPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById provides operations to manage the permissions property of the microsoft.graph.driveItem entity.
func (m *DrivesItemRootRequestBuilder) PermissionsById(id string)(*DrivesItemRootPermissionsPermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission%2Did"] = id
    }
    return NewDrivesItemRootPermissionsPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Preview provides operations to call the preview method.
func (m *DrivesItemRootRequestBuilder) Preview()(*DrivesItemRootPreviewRequestBuilder) {
    return NewDrivesItemRootPreviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore provides operations to call the restore method.
func (m *DrivesItemRootRequestBuilder) Restore()(*DrivesItemRootRestoreRequestBuilder) {
    return NewDrivesItemRootRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SearchWithQ provides operations to call the search method.
func (m *DrivesItemRootRequestBuilder) SearchWithQ(q *string)(*DrivesItemRootSearchWithQRequestBuilder) {
    return NewDrivesItemRootSearchWithQRequestBuilderInternal(m.pathParameters, m.requestAdapter, q);
}
// Subscriptions provides operations to manage the subscriptions property of the microsoft.graph.driveItem entity.
func (m *DrivesItemRootRequestBuilder) Subscriptions()(*DrivesItemRootSubscriptionsRequestBuilder) {
    return NewDrivesItemRootSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById provides operations to manage the subscriptions property of the microsoft.graph.driveItem entity.
func (m *DrivesItemRootRequestBuilder) SubscriptionsById(id string)(*DrivesItemRootSubscriptionsSubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription%2Did"] = id
    }
    return NewDrivesItemRootSubscriptionsSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Thumbnails provides operations to manage the thumbnails property of the microsoft.graph.driveItem entity.
func (m *DrivesItemRootRequestBuilder) Thumbnails()(*DrivesItemRootThumbnailsRequestBuilder) {
    return NewDrivesItemRootThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThumbnailsById provides operations to manage the thumbnails property of the microsoft.graph.driveItem entity.
func (m *DrivesItemRootRequestBuilder) ThumbnailsById(id string)(*DrivesItemRootThumbnailsThumbnailSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet%2Did"] = id
    }
    return NewDrivesItemRootThumbnailsThumbnailSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ToDeleteRequestInformation delete navigation property root for me
func (m *DrivesItemRootRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DrivesItemRootRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation retrieve the metadata for a driveItem in a drive by file system path or ID.
func (m *DrivesItemRootRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DrivesItemRootRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property root in me
func (m *DrivesItemRootRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, requestConfiguration *DrivesItemRootRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Unfollow provides operations to call the unfollow method.
func (m *DrivesItemRootRequestBuilder) Unfollow()(*DrivesItemRootUnfollowRequestBuilder) {
    return NewDrivesItemRootUnfollowRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ValidatePermission provides operations to call the validatePermission method.
func (m *DrivesItemRootRequestBuilder) ValidatePermission()(*DrivesItemRootValidatePermissionRequestBuilder) {
    return NewDrivesItemRootValidatePermissionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Versions provides operations to manage the versions property of the microsoft.graph.driveItem entity.
func (m *DrivesItemRootRequestBuilder) Versions()(*DrivesItemRootVersionsRequestBuilder) {
    return NewDrivesItemRootVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById provides operations to manage the versions property of the microsoft.graph.driveItem entity.
func (m *DrivesItemRootRequestBuilder) VersionsById(id string)(*DrivesItemRootVersionsDriveItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion%2Did"] = id
    }
    return NewDrivesItemRootVersionsDriveItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
