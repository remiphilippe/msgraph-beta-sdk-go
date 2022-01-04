package getyammerdeviceusageuserdetailwithdate

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// GetYammerDeviceUsageUserDetailWithDateRequestBuilder builds and executes requests for operations under \reports\microsoft.graph.getYammerDeviceUsageUserDetail(date={date})
type GetYammerDeviceUsageUserDetailWithDateRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GetYammerDeviceUsageUserDetailWithDateRequestBuilderGetOptions options for Get
type GetYammerDeviceUsageUserDetailWithDateRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewGetYammerDeviceUsageUserDetailWithDateRequestBuilderInternal instantiates a new GetYammerDeviceUsageUserDetailWithDateRequestBuilder and sets the default values.
func NewGetYammerDeviceUsageUserDetailWithDateRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter, date *string)(*GetYammerDeviceUsageUserDetailWithDateRequestBuilder) {
    m := &GetYammerDeviceUsageUserDetailWithDateRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/reports/microsoft.graph.getYammerDeviceUsageUserDetail(date={date})";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if date != nil {
        urlTplParams["date"] = *date
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetYammerDeviceUsageUserDetailWithDateRequestBuilder instantiates a new GetYammerDeviceUsageUserDetailWithDateRequestBuilder and sets the default values.
func NewGetYammerDeviceUsageUserDetailWithDateRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetYammerDeviceUsageUserDetailWithDateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetYammerDeviceUsageUserDetailWithDateRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CreateGetRequestInformation invoke function getYammerDeviceUsageUserDetail
func (m *GetYammerDeviceUsageUserDetailWithDateRequestBuilder) CreateGetRequestInformation(options *GetYammerDeviceUsageUserDetailWithDateRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
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
// Get invoke function getYammerDeviceUsageUserDetail
func (m *GetYammerDeviceUsageUserDetailWithDateRequestBuilder) Get(options *GetYammerDeviceUsageUserDetailWithDateRequestBuilderGetOptions)([]byte, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendPrimitiveAsync(*requestInfo, "byte", nil)
    if err != nil {
        return nil, err
    }
    return res.([]byte), nil
}
