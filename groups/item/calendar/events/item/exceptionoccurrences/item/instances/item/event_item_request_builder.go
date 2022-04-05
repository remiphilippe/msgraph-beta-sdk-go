package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0cbade8d505fcd4065ba6a046cf1b6476fb2d5cb7375c1ea884408e9dda2ef57 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/exceptionoccurrences/item/instances/item/multivalueextendedproperties"
    i202cba09bdb9c3764e9a526a214855d2b0393590b36a3c55d488b8617b8db393 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/exceptionoccurrences/item/instances/item/forward"
    i3c47cad1f06ec4a4667820f777cc04b6da035e30654d220d07e9651aab26aabb "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/exceptionoccurrences/item/instances/item/dismissreminder"
    i4f734e5521df49df51c59633017a73962da3b01bde123a92ccfb42762761608f "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/exceptionoccurrences/item/instances/item/cancel"
    i6ee8cbed7f0d4c7688b40e825fa73ffc9521582fa2f33304664e6502927f7b14 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/exceptionoccurrences/item/instances/item/extensions"
    i8ca30338be8787775f26e845c0d441a4591da72458f5a5f8be5a9933b509484b "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/exceptionoccurrences/item/instances/item/accept"
    i9bffa218bef582a222f2a91d6c58187964eca0c1658277ce2aef71686127bf8f "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/exceptionoccurrences/item/instances/item/singlevalueextendedproperties"
    ia78ef70881c99feb64d1f3b5592c2d71b7046d3ee4c3d62344ab456627ef6bb3 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/exceptionoccurrences/item/instances/item/tentativelyaccept"
    ie85b5fd67d71db43375562635fa14404e14dc5b53600225061a3f2444269d265 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/exceptionoccurrences/item/instances/item/attachments"
    if1c28d97bf1e7d848e37036c69dc75bc0b075a75122303c335309ac286013dc4 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/exceptionoccurrences/item/instances/item/calendar"
    if3b1ecf1dd30c4a74da56c153da625e287deb06b2bd52d271de6ac5bbc8885a2 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/exceptionoccurrences/item/instances/item/decline"
    ifcf948c331468453725109b0a1949727bd49fe4a3fe82a8719617e6f3928f6c4 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/exceptionoccurrences/item/instances/item/snoozereminder"
    i2d6ba016bd07e41612078e8cc706dda2fe1ad2e3d321eed55fbf3139d7a19687 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/exceptionoccurrences/item/instances/item/singlevalueextendedproperties/item"
    i53b59241aac3d74cf5307e1b098072f38731a365f8d71e860513d8099d38e913 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/exceptionoccurrences/item/instances/item/attachments/item"
    i82a996a8eaa0804f1737c0542a5ef5bf2422f20fa884eafb042c213ede5378b5 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/exceptionoccurrences/item/instances/item/extensions/item"
    if19dbb9863c664af3e5f7d418f4a9f25cd10eaba0a5c372206f65f8f609dcb61 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendar/events/item/exceptionoccurrences/item/instances/item/multivalueextendedproperties/item"
)

// EventItemRequestBuilder provides operations to manage the instances property of the microsoft.graph.event entity.
type EventItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// EventItemRequestBuilderDeleteOptions options for Delete
type EventItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// EventItemRequestBuilderGetOptions options for Get
type EventItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *EventItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// EventItemRequestBuilderGetQueryParameters the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
type EventItemRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string;
}
// EventItemRequestBuilderPatchOptions options for Patch
type EventItemRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// Accept the accept property
func (m *EventItemRequestBuilder) Accept()(*i8ca30338be8787775f26e845c0d441a4591da72458f5a5f8be5a9933b509484b.AcceptRequestBuilder) {
    return i8ca30338be8787775f26e845c0d441a4591da72458f5a5f8be5a9933b509484b.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Attachments the attachments property
func (m *EventItemRequestBuilder) Attachments()(*ie85b5fd67d71db43375562635fa14404e14dc5b53600225061a3f2444269d265.AttachmentsRequestBuilder) {
    return ie85b5fd67d71db43375562635fa14404e14dc5b53600225061a3f2444269d265.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.events.item.exceptionOccurrences.item.instances.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i53b59241aac3d74cf5307e1b098072f38731a365f8d71e860513d8099d38e913.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i53b59241aac3d74cf5307e1b098072f38731a365f8d71e860513d8099d38e913.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Calendar the calendar property
func (m *EventItemRequestBuilder) Calendar()(*if1c28d97bf1e7d848e37036c69dc75bc0b075a75122303c335309ac286013dc4.CalendarRequestBuilder) {
    return if1c28d97bf1e7d848e37036c69dc75bc0b075a75122303c335309ac286013dc4.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Cancel the cancel property
func (m *EventItemRequestBuilder) Cancel()(*i4f734e5521df49df51c59633017a73962da3b01bde123a92ccfb42762761608f.CancelRequestBuilder) {
    return i4f734e5521df49df51c59633017a73962da3b01bde123a92ccfb42762761608f.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/calendar/events/{event_id}/exceptionOccurrences/{event_id1}/instances/{event_id2}{?select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEventItemRequestBuilder instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property instances for groups
func (m *EventItemRequestBuilder) CreateDeleteRequestInformation(options *EventItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
func (m *EventItemRequestBuilder) CreateGetRequestInformation(options *EventItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property instances in groups
func (m *EventItemRequestBuilder) CreatePatchRequestInformation(options *EventItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Decline the decline property
func (m *EventItemRequestBuilder) Decline()(*if3b1ecf1dd30c4a74da56c153da625e287deb06b2bd52d271de6ac5bbc8885a2.DeclineRequestBuilder) {
    return if3b1ecf1dd30c4a74da56c153da625e287deb06b2bd52d271de6ac5bbc8885a2.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property instances for groups
func (m *EventItemRequestBuilder) Delete(options *EventItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DismissReminder the dismissReminder property
func (m *EventItemRequestBuilder) DismissReminder()(*i3c47cad1f06ec4a4667820f777cc04b6da035e30654d220d07e9651aab26aabb.DismissReminderRequestBuilder) {
    return i3c47cad1f06ec4a4667820f777cc04b6da035e30654d220d07e9651aab26aabb.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Extensions the extensions property
func (m *EventItemRequestBuilder) Extensions()(*i6ee8cbed7f0d4c7688b40e825fa73ffc9521582fa2f33304664e6502927f7b14.ExtensionsRequestBuilder) {
    return i6ee8cbed7f0d4c7688b40e825fa73ffc9521582fa2f33304664e6502927f7b14.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.events.item.exceptionOccurrences.item.instances.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i82a996a8eaa0804f1737c0542a5ef5bf2422f20fa884eafb042c213ede5378b5.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i82a996a8eaa0804f1737c0542a5ef5bf2422f20fa884eafb042c213ede5378b5.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Forward the forward property
func (m *EventItemRequestBuilder) Forward()(*i202cba09bdb9c3764e9a526a214855d2b0393590b36a3c55d488b8617b8db393.ForwardRequestBuilder) {
    return i202cba09bdb9c3764e9a526a214855d2b0393590b36a3c55d488b8617b8db393.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
func (m *EventItemRequestBuilder) Get(options *EventItemRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEventFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable), nil
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i0cbade8d505fcd4065ba6a046cf1b6476fb2d5cb7375c1ea884408e9dda2ef57.MultiValueExtendedPropertiesRequestBuilder) {
    return i0cbade8d505fcd4065ba6a046cf1b6476fb2d5cb7375c1ea884408e9dda2ef57.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.events.item.exceptionOccurrences.item.instances.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*if19dbb9863c664af3e5f7d418f4a9f25cd10eaba0a5c372206f65f8f609dcb61.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return if19dbb9863c664af3e5f7d418f4a9f25cd10eaba0a5c372206f65f8f609dcb61.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property instances in groups
func (m *EventItemRequestBuilder) Patch(options *EventItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i9bffa218bef582a222f2a91d6c58187964eca0c1658277ce2aef71686127bf8f.SingleValueExtendedPropertiesRequestBuilder) {
    return i9bffa218bef582a222f2a91d6c58187964eca0c1658277ce2aef71686127bf8f.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendar.events.item.exceptionOccurrences.item.instances.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i2d6ba016bd07e41612078e8cc706dda2fe1ad2e3d321eed55fbf3139d7a19687.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i2d6ba016bd07e41612078e8cc706dda2fe1ad2e3d321eed55fbf3139d7a19687.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SnoozeReminder the snoozeReminder property
func (m *EventItemRequestBuilder) SnoozeReminder()(*ifcf948c331468453725109b0a1949727bd49fe4a3fe82a8719617e6f3928f6c4.SnoozeReminderRequestBuilder) {
    return ifcf948c331468453725109b0a1949727bd49fe4a3fe82a8719617e6f3928f6c4.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventItemRequestBuilder) TentativelyAccept()(*ia78ef70881c99feb64d1f3b5592c2d71b7046d3ee4c3d62344ab456627ef6bb3.TentativelyAcceptRequestBuilder) {
    return ia78ef70881c99feb64d1f3b5592c2d71b7046d3ee4c3d62344ab456627ef6bb3.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
