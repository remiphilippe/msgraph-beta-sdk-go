package deviceconfigurationsallmanageddevicecertificatestates

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    iedf847a249af5d381be78d32a395069f479c7729ddced45f9b5c72ba105f91b0 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceconfigurationsallmanageddevicecertificatestates/count"
)

// DeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilder provides operations to manage the deviceConfigurationsAllManagedDeviceCertificateStates property of the microsoft.graph.deviceManagement entity.
type DeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilderGetOptions options for Get
type DeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *DeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// DeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilderGetQueryParameters summary of all certificates for all devices.
type DeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool;
    // Expand related entities
    Expand []string;
    // Filter items by property values
    Filter *string;
    // Order items by property values
    Orderby []string;
    // Search items by search phrases
    Search *string;
    // Select properties to be returned
    Select []string;
    // Skip the first n items
    Skip *int32;
    // Show only the first n items
    Top *int32;
}
// DeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilderPostOptions options for Post
type DeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilderPostOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedAllDeviceCertificateStateable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// NewDeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilderInternal instantiates a new DeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilder and sets the default values.
func NewDeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilder) {
    m := &DeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceConfigurationsAllManagedDeviceCertificateStates{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilder instantiates a new DeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilder and sets the default values.
func NewDeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the count property
func (m *DeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilder) Count()(*iedf847a249af5d381be78d32a395069f479c7729ddced45f9b5c72ba105f91b0.CountRequestBuilder) {
    return iedf847a249af5d381be78d32a395069f479c7729ddced45f9b5c72ba105f91b0.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation summary of all certificates for all devices.
func (m *DeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilder) CreateGetRequestInformation(options *DeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to deviceConfigurationsAllManagedDeviceCertificateStates for deviceManagement
func (m *DeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilder) CreatePostRequestInformation(options *DeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilderPostOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
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
// Get summary of all certificates for all devices.
func (m *DeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilder) Get(options *DeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedAllDeviceCertificateStateCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateManagedAllDeviceCertificateStateCollectionResponseFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedAllDeviceCertificateStateCollectionResponseable), nil
}
// Post create new navigation property to deviceConfigurationsAllManagedDeviceCertificateStates for deviceManagement
func (m *DeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilder) Post(options *DeviceConfigurationsAllManagedDeviceCertificateStatesRequestBuilderPostOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedAllDeviceCertificateStateable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateManagedAllDeviceCertificateStateFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedAllDeviceCertificateStateable), nil
}
