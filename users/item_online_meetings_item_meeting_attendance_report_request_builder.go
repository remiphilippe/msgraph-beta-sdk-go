package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemOnlineMeetingsItemMeetingAttendanceReportRequestBuilder provides operations to manage the meetingAttendanceReport property of the microsoft.graph.onlineMeeting entity.
type ItemOnlineMeetingsItemMeetingAttendanceReportRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemOnlineMeetingsItemMeetingAttendanceReportRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOnlineMeetingsItemMeetingAttendanceReportRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemOnlineMeetingsItemMeetingAttendanceReportRequestBuilderGetQueryParameters get the meetingAttendanceReport for an onlineMeeting. Each time an online meeting ends, an attendance report will be generated for that session.
type ItemOnlineMeetingsItemMeetingAttendanceReportRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemOnlineMeetingsItemMeetingAttendanceReportRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOnlineMeetingsItemMeetingAttendanceReportRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemOnlineMeetingsItemMeetingAttendanceReportRequestBuilderGetQueryParameters
}
// ItemOnlineMeetingsItemMeetingAttendanceReportRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOnlineMeetingsItemMeetingAttendanceReportRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AttendanceRecords provides operations to manage the attendanceRecords property of the microsoft.graph.meetingAttendanceReport entity.
func (m *ItemOnlineMeetingsItemMeetingAttendanceReportRequestBuilder) AttendanceRecords()(*ItemOnlineMeetingsItemMeetingAttendanceReportAttendanceRecordsRequestBuilder) {
    return NewItemOnlineMeetingsItemMeetingAttendanceReportAttendanceRecordsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttendanceRecordsById provides operations to manage the attendanceRecords property of the microsoft.graph.meetingAttendanceReport entity.
func (m *ItemOnlineMeetingsItemMeetingAttendanceReportRequestBuilder) AttendanceRecordsById(id string)(*ItemOnlineMeetingsItemMeetingAttendanceReportAttendanceRecordsAttendanceRecordItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attendanceRecord%2Did"] = id
    }
    return NewItemOnlineMeetingsItemMeetingAttendanceReportAttendanceRecordsAttendanceRecordItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewItemOnlineMeetingsItemMeetingAttendanceReportRequestBuilderInternal instantiates a new MeetingAttendanceReportRequestBuilder and sets the default values.
func NewItemOnlineMeetingsItemMeetingAttendanceReportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnlineMeetingsItemMeetingAttendanceReportRequestBuilder) {
    m := &ItemOnlineMeetingsItemMeetingAttendanceReportRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/onlineMeetings/{onlineMeeting%2Did}/meetingAttendanceReport{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewItemOnlineMeetingsItemMeetingAttendanceReportRequestBuilder instantiates a new MeetingAttendanceReportRequestBuilder and sets the default values.
func NewItemOnlineMeetingsItemMeetingAttendanceReportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnlineMeetingsItemMeetingAttendanceReportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOnlineMeetingsItemMeetingAttendanceReportRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property meetingAttendanceReport for users
func (m *ItemOnlineMeetingsItemMeetingAttendanceReportRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemOnlineMeetingsItemMeetingAttendanceReportRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get the meetingAttendanceReport for an onlineMeeting. Each time an online meeting ends, an attendance report will be generated for that session.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/meetingattendancereport-get?view=graph-rest-1.0
func (m *ItemOnlineMeetingsItemMeetingAttendanceReportRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemOnlineMeetingsItemMeetingAttendanceReportRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MeetingAttendanceReportable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMeetingAttendanceReportFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MeetingAttendanceReportable), nil
}
// Patch update the navigation property meetingAttendanceReport in users
func (m *ItemOnlineMeetingsItemMeetingAttendanceReportRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MeetingAttendanceReportable, requestConfiguration *ItemOnlineMeetingsItemMeetingAttendanceReportRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MeetingAttendanceReportable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMeetingAttendanceReportFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MeetingAttendanceReportable), nil
}
// ToDeleteRequestInformation delete navigation property meetingAttendanceReport for users
func (m *ItemOnlineMeetingsItemMeetingAttendanceReportRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemOnlineMeetingsItemMeetingAttendanceReportRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation get the meetingAttendanceReport for an onlineMeeting. Each time an online meeting ends, an attendance report will be generated for that session.
func (m *ItemOnlineMeetingsItemMeetingAttendanceReportRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemOnlineMeetingsItemMeetingAttendanceReportRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property meetingAttendanceReport in users
func (m *ItemOnlineMeetingsItemMeetingAttendanceReportRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MeetingAttendanceReportable, requestConfiguration *ItemOnlineMeetingsItemMeetingAttendanceReportRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
