package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceGetMemberGroupsRequestBuilder provides operations to call the getMemberGroups method.
type ItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceGetMemberGroupsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceGetMemberGroupsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceGetMemberGroupsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceGetMemberGroupsRequestBuilderInternal instantiates a new GetMemberGroupsRequestBuilder and sets the default values.
func NewItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceGetMemberGroupsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceGetMemberGroupsRequestBuilder) {
    m := &ItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceGetMemberGroupsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/authentication/passwordlessMicrosoftAuthenticatorMethods/{passwordlessMicrosoftAuthenticatorAuthenticationMethod%2Did}/device/microsoft.graph.getMemberGroups";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceGetMemberGroupsRequestBuilder instantiates a new GetMemberGroupsRequestBuilder and sets the default values.
func NewItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceGetMemberGroupsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceGetMemberGroupsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceGetMemberGroupsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post return all the group IDs for the groups that the specified user, group, service principal, organizational contact, device, or directory object is a member of. This function is transitive.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/directoryobject-getmembergroups?view=graph-rest-1.0
func (m *ItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceGetMemberGroupsRequestBuilder) Post(ctx context.Context, body ItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceGetMemberGroupsPostRequestBodyable, requestConfiguration *ItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceGetMemberGroupsRequestBuilderPostRequestConfiguration)(ItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceGetMemberGroupsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceGetMemberGroupsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceGetMemberGroupsResponseable), nil
}
// ToPostRequestInformation return all the group IDs for the groups that the specified user, group, service principal, organizational contact, device, or directory object is a member of. This function is transitive.
func (m *ItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceGetMemberGroupsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceGetMemberGroupsPostRequestBodyable, requestConfiguration *ItemAuthenticationPasswordlessMicrosoftAuthenticatorMethodsItemDeviceGetMemberGroupsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
