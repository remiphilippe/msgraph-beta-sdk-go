package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i1bcb02c01f74f11e3aa24ae414b4c02fbc829ed1cb109d46b654b3389426216e "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/multivalueextendedproperties"
    i31a3b37c9dbd3939c543d5983439186da29622278bf71c6087acbd816f5c9239 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/dismissreminder"
    i3b113a779a2e1c26df62a5fc2779e15b76ed3e1090178152bfb73a05102fd7ea "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/extensions"
    i42ac3ded08c709492e9cc73bc81cf7f4f1449cd1dfae36ce915fb987d3483119 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/decline"
    i478bc546e67483a8a03a15808e207d1d211bbbfe1f2859bbad91fd68f0e9d11a "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/forward"
    i4c47294298a9edb565053aaf021f3f8a03f8dd2db6441a9594f3c81e71a84ca2 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/snoozereminder"
    i68776d8a0abd76d8055c38aaa88c6dfe00f5914bab2339676865bca7101b4c45 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/instances"
    i68a37fe1f53b4e518739b68cfbe9c468a06abeb302d9309775d0fb5864c0f17c "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/calendar"
    i7c71d80390df615835b393e48dd37596b7a2b2840beddb58f10ed47068f5db1a "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/tentativelyaccept"
    i7e83e7737d083b5a247d4b212e9838745f73ee0044cef7495db101a4d27d0473 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/cancel"
    i990077bc9239ae4abf654300201198e18c2649d863293f51f3f9660e7dd48615 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/singlevalueextendedproperties"
    ibeac8ac69135ca43836cbc7e86a3296df23382a77085df4916e1d8db90d1d1b2 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/exceptionoccurrences"
    id7b295f54e3c066e9aee590c7e1cf7df8f692457387f19fc085cb1eaf48eebe5 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/accept"
    ide5e10ee2a6673c0fe49d9fdac7898a497e239b5500fc38a3178a52ff622bf83 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/attachments"
    i0aff26eaed51b1fe4dcefaf7fdc9082ff77390914e54b94471bbabd9a4585874 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/instances/item"
    i0fdb512d350c6ee756976fb95fad12a002ad0d7b0d3fcf13590e89792367c733 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/multivalueextendedproperties/item"
    i53332243af1a23b575f7314ef56a041e8aad0a06d071d02238c1b61aeb8cefd9 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/exceptionoccurrences/item"
    iab96f1b678646231464bd3b7fb938638c604d6f548e5c414d63590125b8591b1 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/singlevalueextendedproperties/item"
    ic351771b8e3695ea47a713295cb3355c9bcc6d5ad38b45b39b527beb6475782b "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/attachments/item"
    ie035c10affe323f0959d54cad2fac3bd37b1a2e6c88bcaece5284fe0542447db "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/calendarview/item/extensions/item"
)

// EventItemRequestBuilder provides operations to manage the calendarView property of the microsoft.graph.group entity.
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
// EventItemRequestBuilderGetQueryParameters the calendar view for the calendar. Read-only.
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
func (m *EventItemRequestBuilder) Accept()(*id7b295f54e3c066e9aee590c7e1cf7df8f692457387f19fc085cb1eaf48eebe5.AcceptRequestBuilder) {
    return id7b295f54e3c066e9aee590c7e1cf7df8f692457387f19fc085cb1eaf48eebe5.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Attachments()(*ide5e10ee2a6673c0fe49d9fdac7898a497e239b5500fc38a3178a52ff622bf83.AttachmentsRequestBuilder) {
    return ide5e10ee2a6673c0fe49d9fdac7898a497e239b5500fc38a3178a52ff622bf83.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendarView.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*ic351771b8e3695ea47a713295cb3355c9bcc6d5ad38b45b39b527beb6475782b.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return ic351771b8e3695ea47a713295cb3355c9bcc6d5ad38b45b39b527beb6475782b.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Calendar()(*i68a37fe1f53b4e518739b68cfbe9c468a06abeb302d9309775d0fb5864c0f17c.CalendarRequestBuilder) {
    return i68a37fe1f53b4e518739b68cfbe9c468a06abeb302d9309775d0fb5864c0f17c.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Cancel()(*i7e83e7737d083b5a247d4b212e9838745f73ee0044cef7495db101a4d27d0473.CancelRequestBuilder) {
    return i7e83e7737d083b5a247d4b212e9838745f73ee0044cef7495db101a4d27d0473.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/calendarView/{event_id}{?startDateTime,endDateTime,select}";
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
// CreateDeleteRequestInformation delete navigation property calendarView for groups
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
// CreateGetRequestInformation the calendar view for the calendar. Read-only.
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
// CreatePatchRequestInformation update the navigation property calendarView in groups
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
func (m *EventItemRequestBuilder) Decline()(*i42ac3ded08c709492e9cc73bc81cf7f4f1449cd1dfae36ce915fb987d3483119.DeclineRequestBuilder) {
    return i42ac3ded08c709492e9cc73bc81cf7f4f1449cd1dfae36ce915fb987d3483119.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property calendarView for groups
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
func (m *EventItemRequestBuilder) DismissReminder()(*i31a3b37c9dbd3939c543d5983439186da29622278bf71c6087acbd816f5c9239.DismissReminderRequestBuilder) {
    return i31a3b37c9dbd3939c543d5983439186da29622278bf71c6087acbd816f5c9239.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) ExceptionOccurrences()(*ibeac8ac69135ca43836cbc7e86a3296df23382a77085df4916e1d8db90d1d1b2.ExceptionOccurrencesRequestBuilder) {
    return ibeac8ac69135ca43836cbc7e86a3296df23382a77085df4916e1d8db90d1d1b2.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendarView.item.exceptionOccurrences.item collection
func (m *EventItemRequestBuilder) ExceptionOccurrencesById(id string)(*i53332243af1a23b575f7314ef56a041e8aad0a06d071d02238c1b61aeb8cefd9.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return i53332243af1a23b575f7314ef56a041e8aad0a06d071d02238c1b61aeb8cefd9.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Extensions()(*i3b113a779a2e1c26df62a5fc2779e15b76ed3e1090178152bfb73a05102fd7ea.ExtensionsRequestBuilder) {
    return i3b113a779a2e1c26df62a5fc2779e15b76ed3e1090178152bfb73a05102fd7ea.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendarView.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*ie035c10affe323f0959d54cad2fac3bd37b1a2e6c88bcaece5284fe0542447db.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return ie035c10affe323f0959d54cad2fac3bd37b1a2e6c88bcaece5284fe0542447db.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Forward()(*i478bc546e67483a8a03a15808e207d1d211bbbfe1f2859bbad91fd68f0e9d11a.ForwardRequestBuilder) {
    return i478bc546e67483a8a03a15808e207d1d211bbbfe1f2859bbad91fd68f0e9d11a.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the calendar view for the calendar. Read-only.
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
func (m *EventItemRequestBuilder) Instances()(*i68776d8a0abd76d8055c38aaa88c6dfe00f5914bab2339676865bca7101b4c45.InstancesRequestBuilder) {
    return i68776d8a0abd76d8055c38aaa88c6dfe00f5914bab2339676865bca7101b4c45.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendarView.item.instances.item collection
func (m *EventItemRequestBuilder) InstancesById(id string)(*i0aff26eaed51b1fe4dcefaf7fdc9082ff77390914e54b94471bbabd9a4585874.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return i0aff26eaed51b1fe4dcefaf7fdc9082ff77390914e54b94471bbabd9a4585874.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i1bcb02c01f74f11e3aa24ae414b4c02fbc829ed1cb109d46b654b3389426216e.MultiValueExtendedPropertiesRequestBuilder) {
    return i1bcb02c01f74f11e3aa24ae414b4c02fbc829ed1cb109d46b654b3389426216e.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendarView.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i0fdb512d350c6ee756976fb95fad12a002ad0d7b0d3fcf13590e89792367c733.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i0fdb512d350c6ee756976fb95fad12a002ad0d7b0d3fcf13590e89792367c733.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property calendarView in groups
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i990077bc9239ae4abf654300201198e18c2649d863293f51f3f9660e7dd48615.SingleValueExtendedPropertiesRequestBuilder) {
    return i990077bc9239ae4abf654300201198e18c2649d863293f51f3f9660e7dd48615.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.calendarView.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*iab96f1b678646231464bd3b7fb938638c604d6f548e5c414d63590125b8591b1.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return iab96f1b678646231464bd3b7fb938638c604d6f548e5c414d63590125b8591b1.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) SnoozeReminder()(*i4c47294298a9edb565053aaf021f3f8a03f8dd2db6441a9594f3c81e71a84ca2.SnoozeReminderRequestBuilder) {
    return i4c47294298a9edb565053aaf021f3f8a03f8dd2db6441a9594f3c81e71a84ca2.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) TentativelyAccept()(*i7c71d80390df615835b393e48dd37596b7a2b2840beddb58f10ed47068f5db1a.TentativelyAcceptRequestBuilder) {
    return i7c71d80390df615835b393e48dd37596b7a2b2840beddb58f10ed47068f5db1a.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
