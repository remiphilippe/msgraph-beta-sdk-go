package compliance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EdiscoveryCasesItemReviewSetsItemEdiscoveryExportRequestBuilder provides operations to call the export method.
type EdiscoveryCasesItemReviewSetsItemEdiscoveryExportRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EdiscoveryCasesItemReviewSetsItemEdiscoveryExportRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryCasesItemReviewSetsItemEdiscoveryExportRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEdiscoveryCasesItemReviewSetsItemEdiscoveryExportRequestBuilderInternal instantiates a new EdiscoveryExportRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemReviewSetsItemEdiscoveryExportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemReviewSetsItemEdiscoveryExportRequestBuilder) {
    m := &EdiscoveryCasesItemReviewSetsItemEdiscoveryExportRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/compliance/ediscovery/cases/{case%2Did}/reviewSets/{reviewSet%2Did}/ediscovery.export", pathParameters),
    }
    return m
}
// NewEdiscoveryCasesItemReviewSetsItemEdiscoveryExportRequestBuilder instantiates a new EdiscoveryExportRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemReviewSetsItemEdiscoveryExportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemReviewSetsItemEdiscoveryExportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEdiscoveryCasesItemReviewSetsItemEdiscoveryExportRequestBuilderInternal(urlParams, requestAdapter)
}
// Post initiate an export from a **reviewSet**.  For details, see Export documents from a review set in Advanced eDiscovery.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/ediscovery-reviewset-export?view=graph-rest-1.0
func (m *EdiscoveryCasesItemReviewSetsItemEdiscoveryExportRequestBuilder) Post(ctx context.Context, body EdiscoveryCasesItemReviewSetsItemEdiscoveryExportExportPostRequestBodyable, requestConfiguration *EdiscoveryCasesItemReviewSetsItemEdiscoveryExportRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation initiate an export from a **reviewSet**.  For details, see Export documents from a review set in Advanced eDiscovery.
func (m *EdiscoveryCasesItemReviewSetsItemEdiscoveryExportRequestBuilder) ToPostRequestInformation(ctx context.Context, body EdiscoveryCasesItemReviewSetsItemEdiscoveryExportExportPostRequestBodyable, requestConfiguration *EdiscoveryCasesItemReviewSetsItemEdiscoveryExportRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
