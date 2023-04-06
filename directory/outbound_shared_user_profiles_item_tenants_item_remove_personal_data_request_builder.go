package directory

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// OutboundSharedUserProfilesItemTenantsItemRemovePersonalDataRequestBuilder provides operations to call the removePersonalData method.
type OutboundSharedUserProfilesItemTenantsItemRemovePersonalDataRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// OutboundSharedUserProfilesItemTenantsItemRemovePersonalDataRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OutboundSharedUserProfilesItemTenantsItemRemovePersonalDataRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewOutboundSharedUserProfilesItemTenantsItemRemovePersonalDataRequestBuilderInternal instantiates a new RemovePersonalDataRequestBuilder and sets the default values.
func NewOutboundSharedUserProfilesItemTenantsItemRemovePersonalDataRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OutboundSharedUserProfilesItemTenantsItemRemovePersonalDataRequestBuilder) {
    m := &OutboundSharedUserProfilesItemTenantsItemRemovePersonalDataRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/directory/outboundSharedUserProfiles/{outboundSharedUserProfile%2DuserId}/tenants/{tenantReference%2DtenantId}/removePersonalData", pathParameters),
    }
    return m
}
// NewOutboundSharedUserProfilesItemTenantsItemRemovePersonalDataRequestBuilder instantiates a new RemovePersonalDataRequestBuilder and sets the default values.
func NewOutboundSharedUserProfilesItemTenantsItemRemovePersonalDataRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OutboundSharedUserProfilesItemTenantsItemRemovePersonalDataRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOutboundSharedUserProfilesItemTenantsItemRemovePersonalDataRequestBuilderInternal(urlParams, requestAdapter)
}
// Post create a request to remove the personal data for an outboundSharedUserProfile.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/tenantreference-removepersonaldata?view=graph-rest-1.0
func (m *OutboundSharedUserProfilesItemTenantsItemRemovePersonalDataRequestBuilder) Post(ctx context.Context, requestConfiguration *OutboundSharedUserProfilesItemTenantsItemRemovePersonalDataRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation create a request to remove the personal data for an outboundSharedUserProfile.
func (m *OutboundSharedUserProfilesItemTenantsItemRemovePersonalDataRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *OutboundSharedUserProfilesItemTenantsItemRemovePersonalDataRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
