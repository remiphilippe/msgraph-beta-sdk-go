package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i367a45f04f002f42b81ce6533585aeeba4211ee6e776fb0ddb4cedf64aa10d6f "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/runsummary"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i494c760a2f26845bb6a83e1609802c363f94efb574cc06272e4db09ffdd5950e "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/groupassignments"
    i5860a2467c9825730f7e5fe8b8587f837d06a6bce92dbb17df1197b8c3e63d0c "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/assign"
    i9caa21903d914933c04aab2668b67aaf07ff8c88530bbc0f796266bff7b47bb2 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/assignments"
    ia6cd2ea0f3d71e372a080d28d01095fd6857f4d1b6250e215940436d6784cbba "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/userrunstates"
    ib4c60324ec4c204995c1740d3c32783fe43242863d0216859b5414b0a63bd4a6 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/devicerunstates"
    i076188fa8b5444fe1eee63d44522c5efaa3b78288e012e2ab0e0bf58bc49a392 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/groupassignments/item"
    i2d5a54257a3a154494e70602191b25e474f0eba5ac59afc7bbe7d9947af31558 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/userrunstates/item"
    i868d4ff15f075e109c23674f69f3c9fe368b53fdc85733d5cad6217558ffb536 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/devicerunstates/item"
    ife64bae8f85f9d9ffa4f9b0e1f47986455765630e26bfebb51ad049410aab2c6 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/devicecustomattributeshellscripts/item/assignments/item"
)

// DeviceCustomAttributeShellScriptItemRequestBuilder provides operations to manage the deviceCustomAttributeShellScripts property of the microsoft.graph.deviceManagement entity.
type DeviceCustomAttributeShellScriptItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DeviceCustomAttributeShellScriptItemRequestBuilderDeleteOptions options for Delete
type DeviceCustomAttributeShellScriptItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceCustomAttributeShellScriptItemRequestBuilderGetOptions options for Get
type DeviceCustomAttributeShellScriptItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DeviceCustomAttributeShellScriptItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceCustomAttributeShellScriptItemRequestBuilderGetQueryParameters the list of device custom attribute shell scripts associated with the tenant.
type DeviceCustomAttributeShellScriptItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DeviceCustomAttributeShellScriptItemRequestBuilderPatchOptions options for Patch
type DeviceCustomAttributeShellScriptItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceCustomAttributeShellScriptable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *DeviceCustomAttributeShellScriptItemRequestBuilder) Assign()(*i5860a2467c9825730f7e5fe8b8587f837d06a6bce92dbb17df1197b8c3e63d0c.AssignRequestBuilder) {
    return i5860a2467c9825730f7e5fe8b8587f837d06a6bce92dbb17df1197b8c3e63d0c.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceCustomAttributeShellScriptItemRequestBuilder) Assignments()(*i9caa21903d914933c04aab2668b67aaf07ff8c88530bbc0f796266bff7b47bb2.AssignmentsRequestBuilder) {
    return i9caa21903d914933c04aab2668b67aaf07ff8c88530bbc0f796266bff7b47bb2.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.deviceCustomAttributeShellScripts.item.assignments.item collection
func (m *DeviceCustomAttributeShellScriptItemRequestBuilder) AssignmentsById(id string)(*ife64bae8f85f9d9ffa4f9b0e1f47986455765630e26bfebb51ad049410aab2c6.DeviceManagementScriptAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementScriptAssignment_id"] = id
    }
    return ife64bae8f85f9d9ffa4f9b0e1f47986455765630e26bfebb51ad049410aab2c6.NewDeviceManagementScriptAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceCustomAttributeShellScriptItemRequestBuilderInternal instantiates a new DeviceCustomAttributeShellScriptItemRequestBuilder and sets the default values.
func NewDeviceCustomAttributeShellScriptItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceCustomAttributeShellScriptItemRequestBuilder) {
    m := &DeviceCustomAttributeShellScriptItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceCustomAttributeShellScripts/{deviceCustomAttributeShellScript_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceCustomAttributeShellScriptItemRequestBuilder instantiates a new DeviceCustomAttributeShellScriptItemRequestBuilder and sets the default values.
func NewDeviceCustomAttributeShellScriptItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceCustomAttributeShellScriptItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceCustomAttributeShellScriptItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property deviceCustomAttributeShellScripts for deviceManagement
func (m *DeviceCustomAttributeShellScriptItemRequestBuilder) CreateDeleteRequestInformation(options *DeviceCustomAttributeShellScriptItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the list of device custom attribute shell scripts associated with the tenant.
func (m *DeviceCustomAttributeShellScriptItemRequestBuilder) CreateGetRequestInformation(options *DeviceCustomAttributeShellScriptItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property deviceCustomAttributeShellScripts in deviceManagement
func (m *DeviceCustomAttributeShellScriptItemRequestBuilder) CreatePatchRequestInformation(options *DeviceCustomAttributeShellScriptItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property deviceCustomAttributeShellScripts for deviceManagement
func (m *DeviceCustomAttributeShellScriptItemRequestBuilder) Delete(options *DeviceCustomAttributeShellScriptItemRequestBuilderDeleteOptions)(error) {
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
func (m *DeviceCustomAttributeShellScriptItemRequestBuilder) DeviceRunStates()(*ib4c60324ec4c204995c1740d3c32783fe43242863d0216859b5414b0a63bd4a6.DeviceRunStatesRequestBuilder) {
    return ib4c60324ec4c204995c1740d3c32783fe43242863d0216859b5414b0a63bd4a6.NewDeviceRunStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceRunStatesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.deviceCustomAttributeShellScripts.item.deviceRunStates.item collection
func (m *DeviceCustomAttributeShellScriptItemRequestBuilder) DeviceRunStatesById(id string)(*i868d4ff15f075e109c23674f69f3c9fe368b53fdc85733d5cad6217558ffb536.DeviceManagementScriptDeviceStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementScriptDeviceState_id"] = id
    }
    return i868d4ff15f075e109c23674f69f3c9fe368b53fdc85733d5cad6217558ffb536.NewDeviceManagementScriptDeviceStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the list of device custom attribute shell scripts associated with the tenant.
func (m *DeviceCustomAttributeShellScriptItemRequestBuilder) Get(options *DeviceCustomAttributeShellScriptItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceCustomAttributeShellScriptable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateDeviceCustomAttributeShellScriptFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceCustomAttributeShellScriptable), nil
}
func (m *DeviceCustomAttributeShellScriptItemRequestBuilder) GroupAssignments()(*i494c760a2f26845bb6a83e1609802c363f94efb574cc06272e4db09ffdd5950e.GroupAssignmentsRequestBuilder) {
    return i494c760a2f26845bb6a83e1609802c363f94efb574cc06272e4db09ffdd5950e.NewGroupAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GroupAssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.deviceCustomAttributeShellScripts.item.groupAssignments.item collection
func (m *DeviceCustomAttributeShellScriptItemRequestBuilder) GroupAssignmentsById(id string)(*i076188fa8b5444fe1eee63d44522c5efaa3b78288e012e2ab0e0bf58bc49a392.DeviceManagementScriptGroupAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementScriptGroupAssignment_id"] = id
    }
    return i076188fa8b5444fe1eee63d44522c5efaa3b78288e012e2ab0e0bf58bc49a392.NewDeviceManagementScriptGroupAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property deviceCustomAttributeShellScripts in deviceManagement
func (m *DeviceCustomAttributeShellScriptItemRequestBuilder) Patch(options *DeviceCustomAttributeShellScriptItemRequestBuilderPatchOptions)(error) {
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
func (m *DeviceCustomAttributeShellScriptItemRequestBuilder) RunSummary()(*i367a45f04f002f42b81ce6533585aeeba4211ee6e776fb0ddb4cedf64aa10d6f.RunSummaryRequestBuilder) {
    return i367a45f04f002f42b81ce6533585aeeba4211ee6e776fb0ddb4cedf64aa10d6f.NewRunSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceCustomAttributeShellScriptItemRequestBuilder) UserRunStates()(*ia6cd2ea0f3d71e372a080d28d01095fd6857f4d1b6250e215940436d6784cbba.UserRunStatesRequestBuilder) {
    return ia6cd2ea0f3d71e372a080d28d01095fd6857f4d1b6250e215940436d6784cbba.NewUserRunStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserRunStatesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.deviceCustomAttributeShellScripts.item.userRunStates.item collection
func (m *DeviceCustomAttributeShellScriptItemRequestBuilder) UserRunStatesById(id string)(*i2d5a54257a3a154494e70602191b25e474f0eba5ac59afc7bbe7d9947af31558.DeviceManagementScriptUserStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementScriptUserState_id"] = id
    }
    return i2d5a54257a3a154494e70602191b25e474f0eba5ac59afc7bbe7d9947af31558.NewDeviceManagementScriptUserStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
