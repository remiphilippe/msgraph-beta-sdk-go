package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder provides operations to manage the instances property of the microsoft.graph.event entity.
type ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilderGetQueryParameters the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
type ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilderGetQueryParameters
}
// Accept provides operations to call the accept method.
func (m *ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) Accept()(*ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemAcceptRequestBuilder) {
    return NewItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) Attachments()(*ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemAttachmentsRequestBuilder) {
    return NewItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) AttachmentsById(id string)(*ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemAttachmentsAttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return NewItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemAttachmentsAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
func (m *ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) Calendar()(*ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemCalendarRequestBuilder) {
    return NewItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel provides operations to call the cancel method.
func (m *ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) Cancel()(*ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemCancelRequestBuilder) {
    return NewItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) {
    m := &ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/calendars/{calendar%2Did}/calendarView/{event%2Did}/exceptionOccurrences/{event%2Did1}/instances/{event%2Did2}{?%24select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder instantiates a new EventItemRequestBuilder and sets the default values.
func NewItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Decline provides operations to call the decline method.
func (m *ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) Decline()(*ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemDeclineRequestBuilder) {
    return NewItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DismissReminder provides operations to call the dismissReminder method.
func (m *ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) DismissReminder()(*ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemDismissReminderRequestBuilder) {
    return NewItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) Extensions()(*ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemExtensionsRequestBuilder) {
    return NewItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) ExtensionsById(id string)(*ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemExtensionsExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return NewItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemExtensionsExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward provides operations to call the forward method.
func (m *ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) Forward()(*ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemForwardRequestBuilder) {
    return NewItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
func (m *ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEventFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable), nil
}
// MultiValueExtendedProperties provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) MultiValueExtendedProperties()(*ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemMultiValueExtendedPropertiesRequestBuilder) {
    return NewItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return NewItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SingleValueExtendedProperties provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) SingleValueExtendedProperties()(*ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemSingleValueExtendedPropertiesRequestBuilder) {
    return NewItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return NewItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder provides operations to call the snoozeReminder method.
func (m *ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) SnoozeReminder()(*ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemSnoozeReminderRequestBuilder) {
    return NewItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept provides operations to call the tentativelyAccept method.
func (m *ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) TentativelyAccept()(*ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemTentativelyAcceptRequestBuilder) {
    return NewItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ToGetRequestInformation the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
func (m *ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesEventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
