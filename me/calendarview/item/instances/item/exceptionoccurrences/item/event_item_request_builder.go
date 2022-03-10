package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i152f667946903f3142a179ce6b8207b72be5d83dd570c317088c5f68e6426de1 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/instances/item/exceptionoccurrences/item/dismissreminder"
    i19008bdd393ac510836f89abdfc5961bbe94940d8ca6650e1dd92f85abdf1a89 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/instances/item/exceptionoccurrences/item/cancel"
    i1db01cd34af40fffde05b40f2ebd8f98ca26f2ed39afa357f54e38851ef29ddf "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/instances/item/exceptionoccurrences/item/snoozereminder"
    i1e5f1ccd25bca6841a67513eaf26575240a16533ebf30d66da8c1f9ab6317246 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/instances/item/exceptionoccurrences/item/decline"
    i20e9cdb6b9ecc1ab3d960602c9613673912c729196ffb07bd83dd5ddf4722522 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/instances/item/exceptionoccurrences/item/extensions"
    i21cb246107a672cced00b657124cc184c3bc3468f1b53f7ee8a6afb99469be41 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/instances/item/exceptionoccurrences/item/accept"
    i5b08890a5c58b7714fc9fef5bb602698db543a14ae9d904f36efa40e9ec4cca5 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties"
    i762fedaf6d516d048dc4cbb2c5bc4b5fb0b695ee49d85a441fa6ebbdc22996ff "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties"
    ia35c643ab7ea66cb5904e1e8413b9d9750d134553af0285e58ada35bf037c966 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/instances/item/exceptionoccurrences/item/attachments"
    ie536ed9529f4fe16b02d0801081385eb18aea8bbefc370965aa9b15dc9900045 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/instances/item/exceptionoccurrences/item/tentativelyaccept"
    ied2d4e78ae16ad635fa91183d46cfab11b73a2d44d9cdd29a311e5528614909d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/instances/item/exceptionoccurrences/item/calendar"
    ife7a07962e2bb49d0cdb81333af66776561905ba176340ecc5a9b2b49a09e8dc "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/instances/item/exceptionoccurrences/item/forward"
    i012601c386deb427e2ad383144b42a711462b3c5bb87560521efacd757400c56 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/instances/item/exceptionoccurrences/item/singlevalueextendedproperties/item"
    i266648a9362e43d946c4da032ba11dd1968a12c53707315c92a23dca9b936a7f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/instances/item/exceptionoccurrences/item/multivalueextendedproperties/item"
    i7e26973f0110103fc737f73342f2351a2be0efeff7b2f04e2e2c876d76e47478 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/instances/item/exceptionoccurrences/item/extensions/item"
    icaa62508cd21e490f961e360752b4fa560dea6978c255091494ded05f6ab9d6b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendarview/item/instances/item/exceptionoccurrences/item/attachments/item"
)

// EventItemRequestBuilder provides operations to manage the exceptionOccurrences property of the microsoft.graph.event entity.
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
// EventItemRequestBuilderGetQueryParameters get exceptionOccurrences from me
type EventItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
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
func (m *EventItemRequestBuilder) Accept()(*i21cb246107a672cced00b657124cc184c3bc3468f1b53f7ee8a6afb99469be41.AcceptRequestBuilder) {
    return i21cb246107a672cced00b657124cc184c3bc3468f1b53f7ee8a6afb99469be41.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Attachments()(*ia35c643ab7ea66cb5904e1e8413b9d9750d134553af0285e58ada35bf037c966.AttachmentsRequestBuilder) {
    return ia35c643ab7ea66cb5904e1e8413b9d9750d134553af0285e58ada35bf037c966.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarView.item.instances.item.exceptionOccurrences.item.attachments.item collection
func (m *EventItemRequestBuilder) AttachmentsById(id string)(*icaa62508cd21e490f961e360752b4fa560dea6978c255091494ded05f6ab9d6b.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return icaa62508cd21e490f961e360752b4fa560dea6978c255091494ded05f6ab9d6b.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Calendar()(*ied2d4e78ae16ad635fa91183d46cfab11b73a2d44d9cdd29a311e5528614909d.CalendarRequestBuilder) {
    return ied2d4e78ae16ad635fa91183d46cfab11b73a2d44d9cdd29a311e5528614909d.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Cancel()(*i19008bdd393ac510836f89abdfc5961bbe94940d8ca6650e1dd92f85abdf1a89.CancelRequestBuilder) {
    return i19008bdd393ac510836f89abdfc5961bbe94940d8ca6650e1dd92f85abdf1a89.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventItemRequestBuilder) {
    m := &EventItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/calendarView/{event_id}/instances/{event_id1}/exceptionOccurrences/{event_id2}{?select,expand}";
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
// CreateDeleteRequestInformation delete navigation property exceptionOccurrences for me
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
// CreateGetRequestInformation get exceptionOccurrences from me
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
// CreatePatchRequestInformation update the navigation property exceptionOccurrences in me
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
func (m *EventItemRequestBuilder) Decline()(*i1e5f1ccd25bca6841a67513eaf26575240a16533ebf30d66da8c1f9ab6317246.DeclineRequestBuilder) {
    return i1e5f1ccd25bca6841a67513eaf26575240a16533ebf30d66da8c1f9ab6317246.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property exceptionOccurrences for me
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
func (m *EventItemRequestBuilder) DismissReminder()(*i152f667946903f3142a179ce6b8207b72be5d83dd570c317088c5f68e6426de1.DismissReminderRequestBuilder) {
    return i152f667946903f3142a179ce6b8207b72be5d83dd570c317088c5f68e6426de1.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Extensions()(*i20e9cdb6b9ecc1ab3d960602c9613673912c729196ffb07bd83dd5ddf4722522.ExtensionsRequestBuilder) {
    return i20e9cdb6b9ecc1ab3d960602c9613673912c729196ffb07bd83dd5ddf4722522.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarView.item.instances.item.exceptionOccurrences.item.extensions.item collection
func (m *EventItemRequestBuilder) ExtensionsById(id string)(*i7e26973f0110103fc737f73342f2351a2be0efeff7b2f04e2e2c876d76e47478.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return i7e26973f0110103fc737f73342f2351a2be0efeff7b2f04e2e2c876d76e47478.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) Forward()(*ife7a07962e2bb49d0cdb81333af66776561905ba176340ecc5a9b2b49a09e8dc.ForwardRequestBuilder) {
    return ife7a07962e2bb49d0cdb81333af66776561905ba176340ecc5a9b2b49a09e8dc.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get exceptionOccurrences from me
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
func (m *EventItemRequestBuilder) MultiValueExtendedProperties()(*i5b08890a5c58b7714fc9fef5bb602698db543a14ae9d904f36efa40e9ec4cca5.MultiValueExtendedPropertiesRequestBuilder) {
    return i5b08890a5c58b7714fc9fef5bb602698db543a14ae9d904f36efa40e9ec4cca5.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarView.item.instances.item.exceptionOccurrences.item.multiValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i266648a9362e43d946c4da032ba11dd1968a12c53707315c92a23dca9b936a7f.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i266648a9362e43d946c4da032ba11dd1968a12c53707315c92a23dca9b936a7f.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property exceptionOccurrences in me
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
func (m *EventItemRequestBuilder) SingleValueExtendedProperties()(*i762fedaf6d516d048dc4cbb2c5bc4b5fb0b695ee49d85a441fa6ebbdc22996ff.SingleValueExtendedPropertiesRequestBuilder) {
    return i762fedaf6d516d048dc4cbb2c5bc4b5fb0b695ee49d85a441fa6ebbdc22996ff.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.calendarView.item.instances.item.exceptionOccurrences.item.singleValueExtendedProperties.item collection
func (m *EventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i012601c386deb427e2ad383144b42a711462b3c5bb87560521efacd757400c56.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i012601c386deb427e2ad383144b42a711462b3c5bb87560521efacd757400c56.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventItemRequestBuilder) SnoozeReminder()(*i1db01cd34af40fffde05b40f2ebd8f98ca26f2ed39afa357f54e38851ef29ddf.SnoozeReminderRequestBuilder) {
    return i1db01cd34af40fffde05b40f2ebd8f98ca26f2ed39afa357f54e38851ef29ddf.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventItemRequestBuilder) TentativelyAccept()(*ie536ed9529f4fe16b02d0801081385eb18aea8bbefc370965aa9b15dc9900045.TentativelyAcceptRequestBuilder) {
    return ie536ed9529f4fe16b02d0801081385eb18aea8bbefc370965aa9b15dc9900045.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
