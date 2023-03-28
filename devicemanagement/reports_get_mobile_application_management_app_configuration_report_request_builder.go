package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ReportsGetMobileApplicationManagementAppConfigurationReportRequestBuilder provides operations to call the getMobileApplicationManagementAppConfigurationReport method.
type ReportsGetMobileApplicationManagementAppConfigurationReportRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ReportsGetMobileApplicationManagementAppConfigurationReportRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReportsGetMobileApplicationManagementAppConfigurationReportRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewReportsGetMobileApplicationManagementAppConfigurationReportRequestBuilderInternal instantiates a new GetMobileApplicationManagementAppConfigurationReportRequestBuilder and sets the default values.
func NewReportsGetMobileApplicationManagementAppConfigurationReportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsGetMobileApplicationManagementAppConfigurationReportRequestBuilder) {
    m := &ReportsGetMobileApplicationManagementAppConfigurationReportRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/reports/getMobileApplicationManagementAppConfigurationReport", pathParameters),
    }
    return m
}
// NewReportsGetMobileApplicationManagementAppConfigurationReportRequestBuilder instantiates a new GetMobileApplicationManagementAppConfigurationReportRequestBuilder and sets the default values.
func NewReportsGetMobileApplicationManagementAppConfigurationReportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsGetMobileApplicationManagementAppConfigurationReportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReportsGetMobileApplicationManagementAppConfigurationReportRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action getMobileApplicationManagementAppConfigurationReport
func (m *ReportsGetMobileApplicationManagementAppConfigurationReportRequestBuilder) Post(ctx context.Context, body ReportsGetMobileApplicationManagementAppConfigurationReportPostRequestBodyable, requestConfiguration *ReportsGetMobileApplicationManagementAppConfigurationReportRequestBuilderPostRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToPostRequestInformation invoke action getMobileApplicationManagementAppConfigurationReport
func (m *ReportsGetMobileApplicationManagementAppConfigurationReportRequestBuilder) ToPostRequestInformation(ctx context.Context, body ReportsGetMobileApplicationManagementAppConfigurationReportPostRequestBodyable, requestConfiguration *ReportsGetMobileApplicationManagementAppConfigurationReportRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
