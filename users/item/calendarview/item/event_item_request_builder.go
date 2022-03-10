package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i1bb040a6172da0f3751a7ceda22fda1a4e82f431532706fa23be932b15ce4da7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/accept"
    i49b3399d3f94916e2dcd2961f6cc29b41beca49a880c4d215fb7071373610633 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/multivalueextendedproperties"
    i5273a0eb53fd57878ca0ece6fb7c5aa1c6b53a2713f8fb3cfda89307d1ebbb18 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/extensions"
    i59caea15222b4e5de6959c735972cffb8c4402d6efdf6aaaf8c7c239a8d57d39 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/singlevalueextendedproperties"
    i6229e0dcc67d4b2439cec0174ff358ae49a1d7c4ba1b18919ea2f4ae0ade1775 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/tentativelyaccept"
    i84de11bd1dac1de0c3b8e7b1613efa73af6db81a30531e727741b72681135b66 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/attachments"
    i892c6aafc6d41ce4059451ac393b4a64e888cc586703223009a6984d4044678c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/dismissreminder"
    ia8dfdfa9be3f0f71b1319d005e93ad8157e2145d947f2e0d992b9c8e2aa1edb4 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/forward"
    ib714f5d74222e2dc5b36353082172e5399fd0349109b9d44f53ac74ae21d9436 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/snoozereminder"
    ic6ce6a431d330b0c7010d3239fc771d8757d7df7f7020314f17078102f22f534 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/instances"
    id0d2b6dbed65b3df6307ebd38618fdef6a17c12e18ec709bdb429080d1148b9e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/calendar"
    id1287a1e327e26acfdcd299e4732689834b3895fb4db5fae88ccf989bd24a46d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/exceptionoccurrences"
    id822bec9817ce7b201d36abc0694c419616196694c625ded53a514e4b502ad26 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/cancel"
    ie4c764f177dd767bef6e8417b11477361fb6c82afd5afecf584ba40ba776a5bd "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/decline"
    i236482608f31b0fda79b0cbbe06fe1399a6be22f9b861ae3e2b2dc11325de539 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/attachments/item"
    i848daab0afa54e7503df8363a3b46fa1b112c1ddc7abb259a29792f4d55276c2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/instances/item"
    i85ba546a2606ea05abdcd903570e68f1f0153bb14bedc01ef322f040d645e05e "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/multivalueextendedproperties/item"
    ib521e23bb7525673c729e3af728ba92bac1d9ff85f2443be14ac282fba44de57 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/singlevalueextendedproperties/item"
    icdda30c8f71f344f627842a9ad82860cf8d02d3a7f04be4929db3d975aa70ee8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/extensions/item"
    iee92f89d05dcbe8d2cd9b4a2f793efd2c5a69122576cc5ea9173539a69c33ce8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendarview/item/exceptionoccurrences/item"
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
func (m *EventItemRequestBuilder) Accept()(*i1bb040a6172da0f3751a7ceda22fda1a4e82f431532706fa23be932b15ce4da7.AcceptRequestBuilder) {
    return i1bb040a6172da0f3751a7ceda22fda1a4e82f431532706fa23be932b15ce4da7.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Attachments()(*i84de11bd1dac1de0c3b8e7b1613efa73af6db81a30531e727741b72681135b66.AttachmentsRequestBuilder) {
    return i84de11bd1dac1de0c3b8e7b1613efa73af6db81a30531e727741b72681135b66.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarView.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*i236482608f31b0fda79b0cbbe06fe1399a6be22f9b861ae3e2b2dc11325de539.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i236482608f31b0fda79b0cbbe06fe1399a6be22f9b861ae3e2b2dc11325de539.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Calendar()(*id0d2b6dbed65b3df6307ebd38618fdef6a17c12e18ec709bdb429080d1148b9e.CalendarRequestBuilder) {
    return id0d2b6dbed65b3df6307ebd38618fdef6a17c12e18ec709bdb429080d1148b9e.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Cancel()(*id822bec9817ce7b201d36abc0694c419616196694c625ded53a514e4b502ad26.CancelRequestBuilder) {
    return id822bec9817ce7b201d36abc0694c419616196694c625ded53a514e4b502ad26.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/calendarView/{event_id}{?startDateTime,endDateTime,select}";
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
// CreateDeleteRequestInformation delete navigation property calendarView for users
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
// CreatePatchRequestInformation update the navigation property calendarView in users
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
func (m *EventItemRequestBuilder) Decline()(*ie4c764f177dd767bef6e8417b11477361fb6c82afd5afecf584ba40ba776a5bd.DeclineRequestBuilder) {
    return ie4c764f177dd767bef6e8417b11477361fb6c82afd5afecf584ba40ba776a5bd.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property calendarView for users
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
func (m *EventItemRequestBuilder) DismissReminder()(*i892c6aafc6d41ce4059451ac393b4a64e888cc586703223009a6984d4044678c.DismissReminderRequestBuilder) {
    return i892c6aafc6d41ce4059451ac393b4a64e888cc586703223009a6984d4044678c.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) ExceptionOccurrences()(*id1287a1e327e26acfdcd299e4732689834b3895fb4db5fae88ccf989bd24a46d.ExceptionOccurrencesRequestBuilder) {
    return id1287a1e327e26acfdcd299e4732689834b3895fb4db5fae88ccf989bd24a46d.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExceptionOccurrencesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarView.item.exceptionOccurrences.item collection
func (m *EventItemRequestBuilder) ExceptionOccurrencesById(id string)(*iee92f89d05dcbe8d2cd9b4a2f793efd2c5a69122576cc5ea9173539a69c33ce8.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return iee92f89d05dcbe8d2cd9b4a2f793efd2c5a69122576cc5ea9173539a69c33ce8.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Extensions()(*i5273a0eb53fd57878ca0ece6fb7c5aa1c6b53a2713f8fb3cfda89307d1ebbb18.ExtensionsRequestBuilder) {
    return i5273a0eb53fd57878ca0ece6fb7c5aa1c6b53a2713f8fb3cfda89307d1ebbb18.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarView.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*icdda30c8f71f344f627842a9ad82860cf8d02d3a7f04be4929db3d975aa70ee8.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return icdda30c8f71f344f627842a9ad82860cf8d02d3a7f04be4929db3d975aa70ee8.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Forward()(*ia8dfdfa9be3f0f71b1319d005e93ad8157e2145d947f2e0d992b9c8e2aa1edb4.ForwardRequestBuilder) {
    return ia8dfdfa9be3f0f71b1319d005e93ad8157e2145d947f2e0d992b9c8e2aa1edb4.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventItemRequestBuilder) Instances()(*ic6ce6a431d330b0c7010d3239fc771d8757d7df7f7020314f17078102f22f534.InstancesRequestBuilder) {
    return ic6ce6a431d330b0c7010d3239fc771d8757d7df7f7020314f17078102f22f534.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstancesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarView.item.instances.item collection
func (m *EventItemRequestBuilder) InstancesById(id string)(*i848daab0afa54e7503df8363a3b46fa1b112c1ddc7abb259a29792f4d55276c2.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return i848daab0afa54e7503df8363a3b46fa1b112c1ddc7abb259a29792f4d55276c2.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i49b3399d3f94916e2dcd2961f6cc29b41beca49a880c4d215fb7071373610633.MultiValueExtendedPropertiesRequestBuilder) {
    return i49b3399d3f94916e2dcd2961f6cc29b41beca49a880c4d215fb7071373610633.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarView.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i85ba546a2606ea05abdcd903570e68f1f0153bb14bedc01ef322f040d645e05e.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i85ba546a2606ea05abdcd903570e68f1f0153bb14bedc01ef322f040d645e05e.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property calendarView in users
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i59caea15222b4e5de6959c735972cffb8c4402d6efdf6aaaf8c7c239a8d57d39.SingleValueExtendedPropertiesRequestBuilder) {
    return i59caea15222b4e5de6959c735972cffb8c4402d6efdf6aaaf8c7c239a8d57d39.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.calendarView.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ib521e23bb7525673c729e3af728ba92bac1d9ff85f2443be14ac282fba44de57.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return ib521e23bb7525673c729e3af728ba92bac1d9ff85f2443be14ac282fba44de57.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) SnoozeReminder()(*ib714f5d74222e2dc5b36353082172e5399fd0349109b9d44f53ac74ae21d9436.SnoozeReminderRequestBuilder) {
    return ib714f5d74222e2dc5b36353082172e5399fd0349109b9d44f53ac74ae21d9436.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) TentativelyAccept()(*i6229e0dcc67d4b2439cec0174ff358ae49a1d7c4ba1b18919ea2f4ae0ade1775.TentativelyAcceptRequestBuilder) {
    return i6229e0dcc67d4b2439cec0174ff358ae49a1d7c4ba1b18919ea2f4ae0ade1775.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
