package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i3bd895fd978dad1672adc3f0311f66fdb03005344906384f3ab769063e50829f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/extensions"
    i3c3a9a8e0b281a1eed928493de2e7d54cad87b5cc7ce112c860a09f13fffdb4f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/multivalueextendedproperties"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i5072dcc71f81fbbc8e1cfabf9d699e903478b616baee2be98a65da249f1fa954 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/instances"
    i5f5f630d5ced06d8366f7b861f39f535ba283864a9c0b987a235807d16bcd2d9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/attachments"
    i62212df5f59e6343c09b1b76ea98ab81a010543533fdc29e488c313f56875602 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/accept"
    i7d64b836ebe68769a340fbfe80e29dc1e3d2d62c42cf1e7f6994fcbbc89fcb34 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/dismissreminder"
    i87f4103909d29d7e9abba659ab2e44d5ca7a48aaa0ab7a174cdffe70e67072f5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/exceptionoccurrences"
    i8fd84426a95c0c0cb3b5115bd277b9e826eab1ca15cbb4d40df967d7fd862d63 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/tentativelyaccept"
    ia46e16175ce133a4275988f6c7f673681cbd68546351722f3943a53f9f9cdcf4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/snoozereminder"
    ic00b975abeb3690b8a537e2dff81241198d1561dbf386671dc806f3228917476 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/singlevalueextendedproperties"
    iddaddb07fe634e5727af12a88fa5d19867a17bf51a0ca2e7db6fc889d2dd561b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/cancel"
    ie9f89f59a45fb8f2559548a58c139cde22ddbbdebe5d326c2d897ca32fb418e6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/decline"
    iecc8853e211513ba4b57daba10b08033aa72795dc491f48ddc834c2b85b35b95 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/forward"
    if0ee07e429daf61efbe896fc2c452f25eb4fb0f148f699739371a571c28387af "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/calendar"
    i4892f6b0282f166f307584c158a0b18f350bd5e674fd73a2f4a68c01793b5ddd "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/singlevalueextendedproperties/item"
    i8b61b6c9f100a3ca9dcf838c2b4ff31850b4d63e0a3898010d0740c45ec3fe53 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/instances/item"
    ic542911019ed2e6fb318ff87489a78092fc04495dfbe0c06e9f2f797b3f1457a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/extensions/item"
    ief7c421ee9694c46920200caadcf67aba8d4d0106872a7a7850705606effdc0f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/multivalueextendedproperties/item"
    if2a79ac6e19e9a10ee61e59ae32c00ed73834be0f7c7760ea2432801f5af3a7a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/attachments/item"
    if55b6d3ec50fcb8d34741fdcda5010c79e181eeee3537556b0cfaaa29dcef6cb "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/exceptionoccurrences/item"
)

// EventItemRequestBuilder provides operations to manage the calendarView property of the microsoft.graph.user entity.
type EventItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// EventItemRequestBuilderDeleteOptions options for Delete
type EventItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// EventItemRequestBuilderGetOptions options for Get
type EventItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *EventItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// EventItemRequestBuilderGetQueryParameters the calendar view for the calendar. Read-only. Nullable.
type EventItemRequestBuilderGetQueryParameters struct {
    // The end date and time of the time range, represented in ISO 8601 format. For example, 2019-11-08T20:00:00-08:00
    EndDateTime *string;
    // Select properties to be returned
    Select []string;
    // The start date and time of the time range, represented in ISO 8601 format. For example, 2019-11-08T19:00:00-08:00
    StartDateTime *string;
}
// EventItemRequestBuilderPatchOptions options for Patch
type EventItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Eventable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *EventItemRequestBuilder) Accept()(*i62212df5f59e6343c09b1b76ea98ab81a010543533fdc29e488c313f56875602.AcceptRequestBuilder) {
    return i62212df5f59e6343c09b1b76ea98ab81a010543533fdc29e488c313f56875602.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Attachments()(*i5f5f630d5ced06d8366f7b861f39f535ba283864a9c0b987a235807d16bcd2d9.AttachmentsRequestBuilder) {
    return i5f5f630d5ced06d8366f7b861f39f535ba283864a9c0b987a235807d16bcd2d9.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarView.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*if2a79ac6e19e9a10ee61e59ae32c00ed73834be0f7c7760ea2432801f5af3a7a.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return if2a79ac6e19e9a10ee61e59ae32c00ed73834be0f7c7760ea2432801f5af3a7a.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Calendar()(*if0ee07e429daf61efbe896fc2c452f25eb4fb0f148f699739371a571c28387af.CalendarRequestBuilder) {
    return if0ee07e429daf61efbe896fc2c452f25eb4fb0f148f699739371a571c28387af.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Cancel()(*iddaddb07fe634e5727af12a88fa5d19867a17bf51a0ca2e7db6fc889d2dd561b.CancelRequestBuilder) {
    return iddaddb07fe634e5727af12a88fa5d19867a17bf51a0ca2e7db6fc889d2dd561b.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendarView/{event_id}{?startDateTime,endDateTime,select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEventItemRequestBuilder instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property calendarView for me
func (m *EventItemRequestBuilder) CreateDeleteRequestInformation(options *EventItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the calendar view for the calendar. Read-only. Nullable.
func (m *EventItemRequestBuilder) CreateGetRequestInformation(options *EventItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property calendarView in me
func (m *EventItemRequestBuilder) CreatePatchRequestInformation(options *EventItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *EventItemRequestBuilder) Decline()(*ie9f89f59a45fb8f2559548a58c139cde22ddbbdebe5d326c2d897ca32fb418e6.DeclineRequestBuilder) {
    return ie9f89f59a45fb8f2559548a58c139cde22ddbbdebe5d326c2d897ca32fb418e6.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property calendarView for me
func (m *EventItemRequestBuilder) Delete(options *EventItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *EventItemRequestBuilder) DismissReminder()(*i7d64b836ebe68769a340fbfe80e29dc1e3d2d62c42cf1e7f6994fcbbc89fcb34.DismissReminderRequestBuilder) {
    return i7d64b836ebe68769a340fbfe80e29dc1e3d2d62c42cf1e7f6994fcbbc89fcb34.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) ExceptionOccurrences()(*i87f4103909d29d7e9abba659ab2e44d5ca7a48aaa0ab7a174cdffe70e67072f5.ExceptionOccurrencesRequestBuilder) {
    return i87f4103909d29d7e9abba659ab2e44d5ca7a48aaa0ab7a174cdffe70e67072f5.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarView.item.exceptionOccurrences.item collection
func (m *EventItemRequestBuilder) ExceptionOccurrencesById(id string)(*if55b6d3ec50fcb8d34741fdcda5010c79e181eeee3537556b0cfaaa29dcef6cb.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return if55b6d3ec50fcb8d34741fdcda5010c79e181eeee3537556b0cfaaa29dcef6cb.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Extensions()(*i3bd895fd978dad1672adc3f0311f66fdb03005344906384f3ab769063e50829f.ExtensionsRequestBuilder) {
    return i3bd895fd978dad1672adc3f0311f66fdb03005344906384f3ab769063e50829f.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarView.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*ic542911019ed2e6fb318ff87489a78092fc04495dfbe0c06e9f2f797b3f1457a.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return ic542911019ed2e6fb318ff87489a78092fc04495dfbe0c06e9f2f797b3f1457a.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Forward()(*iecc8853e211513ba4b57daba10b08033aa72795dc491f48ddc834c2b85b35b95.ForwardRequestBuilder) {
    return iecc8853e211513ba4b57daba10b08033aa72795dc491f48ddc834c2b85b35b95.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the calendar view for the calendar. Read-only. Nullable.
func (m *EventItemRequestBuilder) Get(options *EventItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Eventable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateEventFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Eventable), nil
}
func (m *EventItemRequestBuilder) Instances()(*i5072dcc71f81fbbc8e1cfabf9d699e903478b616baee2be98a65da249f1fa954.InstancesRequestBuilder) {
    return i5072dcc71f81fbbc8e1cfabf9d699e903478b616baee2be98a65da249f1fa954.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarView.item.instances.item collection
func (m *EventItemRequestBuilder) InstancesById(id string)(*i8b61b6c9f100a3ca9dcf838c2b4ff31850b4d63e0a3898010d0740c45ec3fe53.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return i8b61b6c9f100a3ca9dcf838c2b4ff31850b4d63e0a3898010d0740c45ec3fe53.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i3c3a9a8e0b281a1eed928493de2e7d54cad87b5cc7ce112c860a09f13fffdb4f.MultiValueExtendedPropertiesRequestBuilder) {
    return i3c3a9a8e0b281a1eed928493de2e7d54cad87b5cc7ce112c860a09f13fffdb4f.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarView.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ief7c421ee9694c46920200caadcf67aba8d4d0106872a7a7850705606effdc0f.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return ief7c421ee9694c46920200caadcf67aba8d4d0106872a7a7850705606effdc0f.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property calendarView in me
func (m *EventItemRequestBuilder) Patch(options *EventItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*ic00b975abeb3690b8a537e2dff81241198d1561dbf386671dc806f3228917476.SingleValueExtendedPropertiesRequestBuilder) {
    return ic00b975abeb3690b8a537e2dff81241198d1561dbf386671dc806f3228917476.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarView.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i4892f6b0282f166f307584c158a0b18f350bd5e674fd73a2f4a68c01793b5ddd.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i4892f6b0282f166f307584c158a0b18f350bd5e674fd73a2f4a68c01793b5ddd.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) SnoozeReminder()(*ia46e16175ce133a4275988f6c7f673681cbd68546351722f3943a53f9f9cdcf4.SnoozeReminderRequestBuilder) {
    return ia46e16175ce133a4275988f6c7f673681cbd68546351722f3943a53f9f9cdcf4.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) TentativelyAccept()(*i8fd84426a95c0c0cb3b5115bd277b9e826eab1ca15cbb4d40df967d7fd862d63.TentativelyAcceptRequestBuilder) {
    return i8fd84426a95c0c0cb3b5115bd277b9e826eab1ca15cbb4d40df967d7fd862d63.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
