package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
)

// UserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilder provides operations to manage the metricDevices property of the microsoft.graph.userExperienceAnalyticsWorkFromAnywhereMetric entity.
type UserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// UserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilderDeleteOptions options for Delete
type UserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilderGetOptions options for Get
type UserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *UserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// UserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilderGetQueryParameters the work from anywhere metric devices.
type UserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// UserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilderPatchOptions options for Patch
type UserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserExperienceAnalyticsWorkFromAnywhereDeviceable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilderInternal instantiates a new UserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilder) {
    m := &UserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/userExperienceAnalyticsWorkFromAnywhereMetrics/{userExperienceAnalyticsWorkFromAnywhereMetric_id}/metricDevices/{userExperienceAnalyticsWorkFromAnywhereDevice_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilder instantiates a new UserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*UserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property metricDevices for deviceManagement
func (m *UserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilder) CreateDeleteRequestInformation(options *UserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the work from anywhere metric devices.
func (m *UserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilder) CreateGetRequestInformation(options *UserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property metricDevices in deviceManagement
func (m *UserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilder) CreatePatchRequestInformation(options *UserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property metricDevices for deviceManagement
func (m *UserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilder) Delete(options *UserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilderDeleteOptions)(error) {
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
// Get the work from anywhere metric devices.
func (m *UserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilder) Get(options *UserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserExperienceAnalyticsWorkFromAnywhereDeviceable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateUserExperienceAnalyticsWorkFromAnywhereDeviceFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UserExperienceAnalyticsWorkFromAnywhereDeviceable), nil
}
// Patch update the navigation property metricDevices in deviceManagement
func (m *UserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilder) Patch(options *UserExperienceAnalyticsWorkFromAnywhereDeviceItemRequestBuilderPatchOptions)(error) {
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
