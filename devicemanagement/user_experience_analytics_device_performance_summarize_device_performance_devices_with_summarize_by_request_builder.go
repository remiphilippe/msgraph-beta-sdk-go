package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UserExperienceAnalyticsDevicePerformanceSummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilder provides operations to call the summarizeDevicePerformanceDevices method.
type UserExperienceAnalyticsDevicePerformanceSummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UserExperienceAnalyticsDevicePerformanceSummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilderGetQueryParameters invoke function summarizeDevicePerformanceDevices
type UserExperienceAnalyticsDevicePerformanceSummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// UserExperienceAnalyticsDevicePerformanceSummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserExperienceAnalyticsDevicePerformanceSummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserExperienceAnalyticsDevicePerformanceSummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilderGetQueryParameters
}
// NewUserExperienceAnalyticsDevicePerformanceSummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilderInternal instantiates a new SummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsDevicePerformanceSummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, summarizeBy *string)(*UserExperienceAnalyticsDevicePerformanceSummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilder) {
    m := &UserExperienceAnalyticsDevicePerformanceSummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/userExperienceAnalyticsDevicePerformance/microsoft.graph.summarizeDevicePerformanceDevices(summarizeBy='{summarizeBy}'){?%24top,%24skip,%24search,%24filter,%24count,%24select,%24orderby}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if summarizeBy != nil {
        urlTplParams["summarizeBy"] = *summarizeBy
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUserExperienceAnalyticsDevicePerformanceSummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilder instantiates a new SummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsDevicePerformanceSummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserExperienceAnalyticsDevicePerformanceSummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserExperienceAnalyticsDevicePerformanceSummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function summarizeDevicePerformanceDevices
func (m *UserExperienceAnalyticsDevicePerformanceSummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilder) Get(ctx context.Context, requestConfiguration *UserExperienceAnalyticsDevicePerformanceSummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilderGetRequestConfiguration)(UserExperienceAnalyticsDevicePerformanceSummarizeDevicePerformanceDevicesWithSummarizeByResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateUserExperienceAnalyticsDevicePerformanceSummarizeDevicePerformanceDevicesWithSummarizeByResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(UserExperienceAnalyticsDevicePerformanceSummarizeDevicePerformanceDevicesWithSummarizeByResponseable), nil
}
// ToGetRequestInformation invoke function summarizeDevicePerformanceDevices
func (m *UserExperienceAnalyticsDevicePerformanceSummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserExperienceAnalyticsDevicePerformanceSummarizeDevicePerformanceDevicesWithSummarizeByRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
