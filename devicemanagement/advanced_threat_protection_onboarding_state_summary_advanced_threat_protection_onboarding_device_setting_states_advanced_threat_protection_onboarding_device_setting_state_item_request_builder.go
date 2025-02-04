package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilder provides operations to manage the advancedThreatProtectionOnboardingDeviceSettingStates property of the microsoft.graph.advancedThreatProtectionOnboardingStateSummary entity.
type AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilderGetQueryParameters not yet documented
type AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilderGetQueryParameters
}
// AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilderInternal instantiates a new AdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilder and sets the default values.
func NewAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilder) {
    m := &AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/advancedThreatProtectionOnboardingStateSummary/advancedThreatProtectionOnboardingDeviceSettingStates/{advancedThreatProtectionOnboardingDeviceSettingState%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilder instantiates a new AdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilder and sets the default values.
func NewAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property advancedThreatProtectionOnboardingDeviceSettingStates for deviceManagement
func (m *AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get not yet documented
func (m *AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdvancedThreatProtectionOnboardingDeviceSettingStateable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAdvancedThreatProtectionOnboardingDeviceSettingStateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdvancedThreatProtectionOnboardingDeviceSettingStateable), nil
}
// Patch update the navigation property advancedThreatProtectionOnboardingDeviceSettingStates in deviceManagement
func (m *AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdvancedThreatProtectionOnboardingDeviceSettingStateable, requestConfiguration *AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdvancedThreatProtectionOnboardingDeviceSettingStateable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAdvancedThreatProtectionOnboardingDeviceSettingStateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdvancedThreatProtectionOnboardingDeviceSettingStateable), nil
}
// ToDeleteRequestInformation delete navigation property advancedThreatProtectionOnboardingDeviceSettingStates for deviceManagement
func (m *AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation not yet documented
func (m *AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property advancedThreatProtectionOnboardingDeviceSettingStates in deviceManagement
func (m *AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AdvancedThreatProtectionOnboardingDeviceSettingStateable, requestConfiguration *AdvancedThreatProtectionOnboardingStateSummaryAdvancedThreatProtectionOnboardingDeviceSettingStatesAdvancedThreatProtectionOnboardingDeviceSettingStateItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
