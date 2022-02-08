package getcompliancesettingdetailsreport

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// GetComplianceSettingDetailsReportRequestBuilder builds and executes requests for operations under \deviceManagement\reports\microsoft.graph.getComplianceSettingDetailsReport
type GetComplianceSettingDetailsReportRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GetComplianceSettingDetailsReportRequestBuilderPostOptions options for Post
type GetComplianceSettingDetailsReportRequestBuilderPostOptions struct {
    // 
    Body *GetComplianceSettingDetailsReportRequestBody;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewGetComplianceSettingDetailsReportRequestBuilderInternal instantiates a new GetComplianceSettingDetailsReportRequestBuilder and sets the default values.
func NewGetComplianceSettingDetailsReportRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetComplianceSettingDetailsReportRequestBuilder) {
    m := &GetComplianceSettingDetailsReportRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/reports/microsoft.graph.getComplianceSettingDetailsReport";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetComplianceSettingDetailsReportRequestBuilder instantiates a new GetComplianceSettingDetailsReportRequestBuilder and sets the default values.
func NewGetComplianceSettingDetailsReportRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetComplianceSettingDetailsReportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetComplianceSettingDetailsReportRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action getComplianceSettingDetailsReport
func (m *GetComplianceSettingDetailsReportRequestBuilder) CreatePostRequestInformation(options *GetComplianceSettingDetailsReportRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.POST
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
// Post invoke action getComplianceSettingDetailsReport
func (m *GetComplianceSettingDetailsReportRequestBuilder) Post(options *GetComplianceSettingDetailsReportRequestBuilderPostOptions)([]byte, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendPrimitiveAsync(*requestInfo, "byte", nil)
    if err != nil {
        return nil, err
    }
    return res.([]byte), nil
}
