package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
)

// IosUpdateDeviceStatusItemRequestBuilder provides operations to manage the iosUpdateStatuses property of the microsoft.graph.deviceManagement entity.
type IosUpdateDeviceStatusItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// IosUpdateDeviceStatusItemRequestBuilderDeleteOptions options for Delete
type IosUpdateDeviceStatusItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// IosUpdateDeviceStatusItemRequestBuilderGetOptions options for Get
type IosUpdateDeviceStatusItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *IosUpdateDeviceStatusItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// IosUpdateDeviceStatusItemRequestBuilderGetQueryParameters the IOS software update installation statuses for this account.
type IosUpdateDeviceStatusItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// IosUpdateDeviceStatusItemRequestBuilderPatchOptions options for Patch
type IosUpdateDeviceStatusItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.IosUpdateDeviceStatusable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewIosUpdateDeviceStatusItemRequestBuilderInternal instantiates a new IosUpdateDeviceStatusItemRequestBuilder and sets the default values.
func NewIosUpdateDeviceStatusItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*IosUpdateDeviceStatusItemRequestBuilder) {
    m := &IosUpdateDeviceStatusItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/iosUpdateStatuses/{iosUpdateDeviceStatus_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIosUpdateDeviceStatusItemRequestBuilder instantiates a new IosUpdateDeviceStatusItemRequestBuilder and sets the default values.
func NewIosUpdateDeviceStatusItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*IosUpdateDeviceStatusItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIosUpdateDeviceStatusItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property iosUpdateStatuses for deviceManagement
func (m *IosUpdateDeviceStatusItemRequestBuilder) CreateDeleteRequestInformation(options *IosUpdateDeviceStatusItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the IOS software update installation statuses for this account.
func (m *IosUpdateDeviceStatusItemRequestBuilder) CreateGetRequestInformation(options *IosUpdateDeviceStatusItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property iosUpdateStatuses in deviceManagement
func (m *IosUpdateDeviceStatusItemRequestBuilder) CreatePatchRequestInformation(options *IosUpdateDeviceStatusItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property iosUpdateStatuses for deviceManagement
func (m *IosUpdateDeviceStatusItemRequestBuilder) Delete(options *IosUpdateDeviceStatusItemRequestBuilderDeleteOptions)(error) {
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
// Get the IOS software update installation statuses for this account.
func (m *IosUpdateDeviceStatusItemRequestBuilder) Get(options *IosUpdateDeviceStatusItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.IosUpdateDeviceStatusable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateIosUpdateDeviceStatusFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.IosUpdateDeviceStatusable), nil
}
// Patch update the navigation property iosUpdateStatuses in deviceManagement
func (m *IosUpdateDeviceStatusItemRequestBuilder) Patch(options *IosUpdateDeviceStatusItemRequestBuilderPatchOptions)(error) {
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
