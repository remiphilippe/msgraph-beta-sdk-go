package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DataSharingConsentsItemConsentToDataSharingRequestBuilder provides operations to call the consentToDataSharing method.
type DataSharingConsentsItemConsentToDataSharingRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DataSharingConsentsItemConsentToDataSharingRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DataSharingConsentsItemConsentToDataSharingRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDataSharingConsentsItemConsentToDataSharingRequestBuilderInternal instantiates a new ConsentToDataSharingRequestBuilder and sets the default values.
func NewDataSharingConsentsItemConsentToDataSharingRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DataSharingConsentsItemConsentToDataSharingRequestBuilder) {
    m := &DataSharingConsentsItemConsentToDataSharingRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/dataSharingConsents/{dataSharingConsent%2Did}/consentToDataSharing", pathParameters),
    }
    return m
}
// NewDataSharingConsentsItemConsentToDataSharingRequestBuilder instantiates a new ConsentToDataSharingRequestBuilder and sets the default values.
func NewDataSharingConsentsItemConsentToDataSharingRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DataSharingConsentsItemConsentToDataSharingRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDataSharingConsentsItemConsentToDataSharingRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action consentToDataSharing
func (m *DataSharingConsentsItemConsentToDataSharingRequestBuilder) Post(ctx context.Context, requestConfiguration *DataSharingConsentsItemConsentToDataSharingRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DataSharingConsentable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDataSharingConsentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DataSharingConsentable), nil
}
// ToPostRequestInformation invoke action consentToDataSharing
func (m *DataSharingConsentsItemConsentToDataSharingRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *DataSharingConsentsItemConsentToDataSharingRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
