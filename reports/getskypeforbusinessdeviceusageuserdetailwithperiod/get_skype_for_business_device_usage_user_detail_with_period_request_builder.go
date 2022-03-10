package getskypeforbusinessdeviceusageuserdetailwithperiod

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// GetSkypeForBusinessDeviceUsageUserDetailWithPeriodRequestBuilder provides operations to call the getSkypeForBusinessDeviceUsageUserDetail method.
type GetSkypeForBusinessDeviceUsageUserDetailWithPeriodRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GetSkypeForBusinessDeviceUsageUserDetailWithPeriodRequestBuilderGetOptions options for Get
type GetSkypeForBusinessDeviceUsageUserDetailWithPeriodRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewGetSkypeForBusinessDeviceUsageUserDetailWithPeriodRequestBuilderInternal instantiates a new GetSkypeForBusinessDeviceUsageUserDetailWithPeriodRequestBuilder and sets the default values.
func NewGetSkypeForBusinessDeviceUsageUserDetailWithPeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter, period *string)(*GetSkypeForBusinessDeviceUsageUserDetailWithPeriodRequestBuilder) {
    m := &GetSkypeForBusinessDeviceUsageUserDetailWithPeriodRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/reports/microsoft.graph.getSkypeForBusinessDeviceUsageUserDetail(period='{period}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if period != nil {
        urlTplParams["period"] = *period
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetSkypeForBusinessDeviceUsageUserDetailWithPeriodRequestBuilder instantiates a new GetSkypeForBusinessDeviceUsageUserDetailWithPeriodRequestBuilder and sets the default values.
func NewGetSkypeForBusinessDeviceUsageUserDetailWithPeriodRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetSkypeForBusinessDeviceUsageUserDetailWithPeriodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetSkypeForBusinessDeviceUsageUserDetailWithPeriodRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CreateGetRequestInformation invoke function getSkypeForBusinessDeviceUsageUserDetail
func (m *GetSkypeForBusinessDeviceUsageUserDetailWithPeriodRequestBuilder) CreateGetRequestInformation(options *GetSkypeForBusinessDeviceUsageUserDetailWithPeriodRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get invoke function getSkypeForBusinessDeviceUsageUserDetail
func (m *GetSkypeForBusinessDeviceUsageUserDetailWithPeriodRequestBuilder) Get(options *GetSkypeForBusinessDeviceUsageUserDetailWithPeriodRequestBuilderGetOptions)(GetSkypeForBusinessDeviceUsageUserDetailWithPeriodResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetSkypeForBusinessDeviceUsageUserDetailWithPeriodResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetSkypeForBusinessDeviceUsageUserDetailWithPeriodResponseable), nil
}
