package advancedthreatprotectiononboardingstatesummary

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    ice19fdfd521d3baa084d788d924acd2488705b59b4281e0f60085cec8edcab8d "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/advancedthreatprotectiononboardingstatesummary/advancedthreatprotectiononboardingdevicesettingstates"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i8ac982e4babd498ccfdca8159c2c00708b65a8c1d00512bc9fe29edca2a507ca "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/advancedthreatprotectiononboardingstatesummary/advancedthreatprotectiononboardingdevicesettingstates/item"
)

// AdvancedThreatProtectionOnboardingStateSummaryRequestBuilder provides operations to manage the advancedThreatProtectionOnboardingStateSummary property of the microsoft.graph.deviceManagement entity.
type AdvancedThreatProtectionOnboardingStateSummaryRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// AdvancedThreatProtectionOnboardingStateSummaryRequestBuilderDeleteOptions options for Delete
type AdvancedThreatProtectionOnboardingStateSummaryRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AdvancedThreatProtectionOnboardingStateSummaryRequestBuilderGetOptions options for Get
type AdvancedThreatProtectionOnboardingStateSummaryRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *AdvancedThreatProtectionOnboardingStateSummaryRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AdvancedThreatProtectionOnboardingStateSummaryRequestBuilderGetQueryParameters the summary state of ATP onboarding state for this account.
type AdvancedThreatProtectionOnboardingStateSummaryRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// AdvancedThreatProtectionOnboardingStateSummaryRequestBuilderPatchOptions options for Patch
type AdvancedThreatProtectionOnboardingStateSummaryRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AdvancedThreatProtectionOnboardingStateSummaryable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *AdvancedThreatProtectionOnboardingStateSummaryRequestBuilder) AdvancedThreatProtectionOnboardingDeviceSettingStates()(*ice19fdfd521d3baa084d788d924acd2488705b59b4281e0f60085cec8edcab8d.AdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilder) {
    return ice19fdfd521d3baa084d788d924acd2488705b59b4281e0f60085cec8edcab8d.NewAdvancedThreatProtectionOnboardingDeviceSettingStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AdvancedThreatProtectionOnboardingDeviceSettingStatesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.advancedThreatProtectionOnboardingStateSummary.advancedThreatProtectionOnboardingDeviceSettingStates.item collection
func (m *AdvancedThreatProtectionOnboardingStateSummaryRequestBuilder) AdvancedThreatProtectionOnboardingDeviceSettingStatesById(id string)(*i8ac982e4babd498ccfdca8159c2c00708b65a8c1d00512bc9fe29edca2a507ca.AdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["advancedThreatProtectionOnboardingDeviceSettingState_id"] = id
    }
    return i8ac982e4babd498ccfdca8159c2c00708b65a8c1d00512bc9fe29edca2a507ca.NewAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewAdvancedThreatProtectionOnboardingStateSummaryRequestBuilderInternal instantiates a new AdvancedThreatProtectionOnboardingStateSummaryRequestBuilder and sets the default values.
func NewAdvancedThreatProtectionOnboardingStateSummaryRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AdvancedThreatProtectionOnboardingStateSummaryRequestBuilder) {
    m := &AdvancedThreatProtectionOnboardingStateSummaryRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/advancedThreatProtectionOnboardingStateSummary{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAdvancedThreatProtectionOnboardingStateSummaryRequestBuilder instantiates a new AdvancedThreatProtectionOnboardingStateSummaryRequestBuilder and sets the default values.
func NewAdvancedThreatProtectionOnboardingStateSummaryRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AdvancedThreatProtectionOnboardingStateSummaryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAdvancedThreatProtectionOnboardingStateSummaryRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property advancedThreatProtectionOnboardingStateSummary for deviceManagement
func (m *AdvancedThreatProtectionOnboardingStateSummaryRequestBuilder) CreateDeleteRequestInformation(options *AdvancedThreatProtectionOnboardingStateSummaryRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the summary state of ATP onboarding state for this account.
func (m *AdvancedThreatProtectionOnboardingStateSummaryRequestBuilder) CreateGetRequestInformation(options *AdvancedThreatProtectionOnboardingStateSummaryRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property advancedThreatProtectionOnboardingStateSummary in deviceManagement
func (m *AdvancedThreatProtectionOnboardingStateSummaryRequestBuilder) CreatePatchRequestInformation(options *AdvancedThreatProtectionOnboardingStateSummaryRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property advancedThreatProtectionOnboardingStateSummary for deviceManagement
func (m *AdvancedThreatProtectionOnboardingStateSummaryRequestBuilder) Delete(options *AdvancedThreatProtectionOnboardingStateSummaryRequestBuilderDeleteOptions)(error) {
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
// Get the summary state of ATP onboarding state for this account.
func (m *AdvancedThreatProtectionOnboardingStateSummaryRequestBuilder) Get(options *AdvancedThreatProtectionOnboardingStateSummaryRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AdvancedThreatProtectionOnboardingStateSummaryable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateAdvancedThreatProtectionOnboardingStateSummaryFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AdvancedThreatProtectionOnboardingStateSummaryable), nil
}
// Patch update the navigation property advancedThreatProtectionOnboardingStateSummary in deviceManagement
func (m *AdvancedThreatProtectionOnboardingStateSummaryRequestBuilder) Patch(options *AdvancedThreatProtectionOnboardingStateSummaryRequestBuilderPatchOptions)(error) {
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
