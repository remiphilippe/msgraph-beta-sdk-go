package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i16903a7f82aef1b9172b084c2f688434ec7df2fe7382cd313d3acfe5d99975b1 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/assignments"
    i1c6baf2c64fdbaacd3e4b6e2fc919aa7f1ee2113534c7b651c6f577f13cdcbd9 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/groupassignments"
    i7f208efc5d45cbfede2e60d9bbf7b141234bf2bb3b1aa6abb8ccde6a36185ad5 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/userrunstates"
    i801a613b15d2cbc0073244db83bf53fd7b41404a5087e44c919ad86255b3948b "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/runsummary"
    ibc0bdc7d62407627f799f1887a2a83100a7295907f2203c9b7be0ceb05727fa8 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/assign"
    ic9701e6869c3e935d0c0f54c628c153ce833fd193972ac4ccd97a4090f51262c "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/devicerunstates"
    i3afdf7675877765716ac7dd1dadedd54dc760d42e9a7d4d3f623971af61f9ddd "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/userrunstates/item"
    i4e99eee29f4d27b3c3ed9b39a732efaaadc4e736ad51030857642293b548d3cb "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/devicerunstates/item"
    i60b1d872fe3c79385687ac815dc9a345c65cf3a1d7475fcd16834aa195d0ef49 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/assignments/item"
    ib0522360ed077a6c16fe77b5b9aca7adca496e97d4d37f10bf42394220f94da8 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceshellscripts/item/groupassignments/item"
)

// DeviceShellScriptItemRequestBuilder provides operations to manage the deviceShellScripts property of the microsoft.graph.deviceManagement entity.
type DeviceShellScriptItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DeviceShellScriptItemRequestBuilderDeleteOptions options for Delete
type DeviceShellScriptItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// DeviceShellScriptItemRequestBuilderGetOptions options for Get
type DeviceShellScriptItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *DeviceShellScriptItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// DeviceShellScriptItemRequestBuilderGetQueryParameters the list of device shell scripts associated with the tenant.
type DeviceShellScriptItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DeviceShellScriptItemRequestBuilderPatchOptions options for Patch
type DeviceShellScriptItemRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceShellScriptable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// Assign the assign property
func (m *DeviceShellScriptItemRequestBuilder) Assign()(*ibc0bdc7d62407627f799f1887a2a83100a7295907f2203c9b7be0ceb05727fa8.AssignRequestBuilder) {
    return ibc0bdc7d62407627f799f1887a2a83100a7295907f2203c9b7be0ceb05727fa8.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Assignments the assignments property
func (m *DeviceShellScriptItemRequestBuilder) Assignments()(*i16903a7f82aef1b9172b084c2f688434ec7df2fe7382cd313d3acfe5d99975b1.AssignmentsRequestBuilder) {
    return i16903a7f82aef1b9172b084c2f688434ec7df2fe7382cd313d3acfe5d99975b1.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.deviceShellScripts.item.assignments.item collection
func (m *DeviceShellScriptItemRequestBuilder) AssignmentsById(id string)(*i60b1d872fe3c79385687ac815dc9a345c65cf3a1d7475fcd16834aa195d0ef49.DeviceManagementScriptAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementScriptAssignment_id"] = id
    }
    return i60b1d872fe3c79385687ac815dc9a345c65cf3a1d7475fcd16834aa195d0ef49.NewDeviceManagementScriptAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceShellScriptItemRequestBuilderInternal instantiates a new DeviceShellScriptItemRequestBuilder and sets the default values.
func NewDeviceShellScriptItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceShellScriptItemRequestBuilder) {
    m := &DeviceShellScriptItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceShellScripts/{deviceShellScript_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceShellScriptItemRequestBuilder instantiates a new DeviceShellScriptItemRequestBuilder and sets the default values.
func NewDeviceShellScriptItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceShellScriptItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceShellScriptItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property deviceShellScripts for deviceManagement
func (m *DeviceShellScriptItemRequestBuilder) CreateDeleteRequestInformation(options *DeviceShellScriptItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the list of device shell scripts associated with the tenant.
func (m *DeviceShellScriptItemRequestBuilder) CreateGetRequestInformation(options *DeviceShellScriptItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property deviceShellScripts in deviceManagement
func (m *DeviceShellScriptItemRequestBuilder) CreatePatchRequestInformation(options *DeviceShellScriptItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property deviceShellScripts for deviceManagement
func (m *DeviceShellScriptItemRequestBuilder) Delete(options *DeviceShellScriptItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DeviceRunStates the deviceRunStates property
func (m *DeviceShellScriptItemRequestBuilder) DeviceRunStates()(*ic9701e6869c3e935d0c0f54c628c153ce833fd193972ac4ccd97a4090f51262c.DeviceRunStatesRequestBuilder) {
    return ic9701e6869c3e935d0c0f54c628c153ce833fd193972ac4ccd97a4090f51262c.NewDeviceRunStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceRunStatesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.deviceShellScripts.item.deviceRunStates.item collection
func (m *DeviceShellScriptItemRequestBuilder) DeviceRunStatesById(id string)(*i4e99eee29f4d27b3c3ed9b39a732efaaadc4e736ad51030857642293b548d3cb.DeviceManagementScriptDeviceStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementScriptDeviceState_id"] = id
    }
    return i4e99eee29f4d27b3c3ed9b39a732efaaadc4e736ad51030857642293b548d3cb.NewDeviceManagementScriptDeviceStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the list of device shell scripts associated with the tenant.
func (m *DeviceShellScriptItemRequestBuilder) Get(options *DeviceShellScriptItemRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceShellScriptable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceShellScriptFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceShellScriptable), nil
}
// GroupAssignments the groupAssignments property
func (m *DeviceShellScriptItemRequestBuilder) GroupAssignments()(*i1c6baf2c64fdbaacd3e4b6e2fc919aa7f1ee2113534c7b651c6f577f13cdcbd9.GroupAssignmentsRequestBuilder) {
    return i1c6baf2c64fdbaacd3e4b6e2fc919aa7f1ee2113534c7b651c6f577f13cdcbd9.NewGroupAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GroupAssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.deviceShellScripts.item.groupAssignments.item collection
func (m *DeviceShellScriptItemRequestBuilder) GroupAssignmentsById(id string)(*ib0522360ed077a6c16fe77b5b9aca7adca496e97d4d37f10bf42394220f94da8.DeviceManagementScriptGroupAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementScriptGroupAssignment_id"] = id
    }
    return ib0522360ed077a6c16fe77b5b9aca7adca496e97d4d37f10bf42394220f94da8.NewDeviceManagementScriptGroupAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property deviceShellScripts in deviceManagement
func (m *DeviceShellScriptItemRequestBuilder) Patch(options *DeviceShellScriptItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// RunSummary the runSummary property
func (m *DeviceShellScriptItemRequestBuilder) RunSummary()(*i801a613b15d2cbc0073244db83bf53fd7b41404a5087e44c919ad86255b3948b.RunSummaryRequestBuilder) {
    return i801a613b15d2cbc0073244db83bf53fd7b41404a5087e44c919ad86255b3948b.NewRunSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserRunStates the userRunStates property
func (m *DeviceShellScriptItemRequestBuilder) UserRunStates()(*i7f208efc5d45cbfede2e60d9bbf7b141234bf2bb3b1aa6abb8ccde6a36185ad5.UserRunStatesRequestBuilder) {
    return i7f208efc5d45cbfede2e60d9bbf7b141234bf2bb3b1aa6abb8ccde6a36185ad5.NewUserRunStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserRunStatesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.deviceShellScripts.item.userRunStates.item collection
func (m *DeviceShellScriptItemRequestBuilder) UserRunStatesById(id string)(*i3afdf7675877765716ac7dd1dadedd54dc760d42e9a7d4d3f623971af61f9ddd.DeviceManagementScriptUserStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementScriptUserState_id"] = id
    }
    return i3afdf7675877765716ac7dd1dadedd54dc760d42e9a7d4d3f623971af61f9ddd.NewDeviceManagementScriptUserStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
