package getbrowseruserdetailwithperiod

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// GetBrowserUserDetailWithPeriodRequestBuilder provides operations to call the getBrowserUserDetail method.
type GetBrowserUserDetailWithPeriodRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GetBrowserUserDetailWithPeriodRequestBuilderGetOptions options for Get
type GetBrowserUserDetailWithPeriodRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewGetBrowserUserDetailWithPeriodRequestBuilderInternal instantiates a new GetBrowserUserDetailWithPeriodRequestBuilder and sets the default values.
func NewGetBrowserUserDetailWithPeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter, period *string)(*GetBrowserUserDetailWithPeriodRequestBuilder) {
    m := &GetBrowserUserDetailWithPeriodRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/reports/microsoft.graph.getBrowserUserDetail(period='{period}')";
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
// NewGetBrowserUserDetailWithPeriodRequestBuilder instantiates a new GetBrowserUserDetailWithPeriodRequestBuilder and sets the default values.
func NewGetBrowserUserDetailWithPeriodRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetBrowserUserDetailWithPeriodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetBrowserUserDetailWithPeriodRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CreateGetRequestInformation invoke function getBrowserUserDetail
func (m *GetBrowserUserDetailWithPeriodRequestBuilder) CreateGetRequestInformation(options *GetBrowserUserDetailWithPeriodRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get invoke function getBrowserUserDetail
func (m *GetBrowserUserDetailWithPeriodRequestBuilder) Get(options *GetBrowserUserDetailWithPeriodRequestBuilderGetOptions)(GetBrowserUserDetailWithPeriodResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetBrowserUserDetailWithPeriodResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetBrowserUserDetailWithPeriodResponseable), nil
}
