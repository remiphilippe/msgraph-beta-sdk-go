package teamwork

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsValidatePropertiesRequestBuilder provides operations to call the validateProperties method.
type TeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsValidatePropertiesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsValidatePropertiesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsValidatePropertiesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsValidatePropertiesRequestBuilderInternal instantiates a new ValidatePropertiesRequestBuilder and sets the default values.
func NewTeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsValidatePropertiesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsValidatePropertiesRequestBuilder) {
    m := &TeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsValidatePropertiesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/teamwork/teamTemplates/{teamTemplate%2Did}/definitions/{teamTemplateDefinition%2Did}/teamDefinition/permissionGrants/microsoft.graph.validateProperties";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsValidatePropertiesRequestBuilder instantiates a new ValidatePropertiesRequestBuilder and sets the default values.
func NewTeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsValidatePropertiesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsValidatePropertiesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsValidatePropertiesRequestBuilderInternal(urlParams, requestAdapter)
}
// Post validate that a Microsoft 365 group's display name or mail nickname complies with naming policies.  Clients can use this API to determine whether a display name or mail nickname is valid before trying to **create** a Microsoft 365 group. For validating properties of an existing group, use the validateProperties function for groups. The following validations are performed for the display name and mail nickname properties: 1. Validate the prefix and suffix naming policy2. Validate the custom banned words policy3. Validate the mail nickname is unique This API returns with the first failure encountered. If one or more properties fail multiple validations, only the property with the first validation failure is returned. However, you can validate both the mail nickname and the display name and receive a collection of validation errors if you are only validating the prefix and suffix naming policy.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/directoryobject-validateproperties?view=graph-rest-1.0
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsValidatePropertiesRequestBuilder) Post(ctx context.Context, body TeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsValidatePropertiesPostRequestBodyable, requestConfiguration *TeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsValidatePropertiesRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation validate that a Microsoft 365 group's display name or mail nickname complies with naming policies.  Clients can use this API to determine whether a display name or mail nickname is valid before trying to **create** a Microsoft 365 group. For validating properties of an existing group, use the validateProperties function for groups. The following validations are performed for the display name and mail nickname properties: 1. Validate the prefix and suffix naming policy2. Validate the custom banned words policy3. Validate the mail nickname is unique This API returns with the first failure encountered. If one or more properties fail multiple validations, only the property with the first validation failure is returned. However, you can validate both the mail nickname and the display name and receive a collection of validation errors if you are only validating the prefix and suffix naming policy.
func (m *TeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsValidatePropertiesRequestBuilder) ToPostRequestInformation(ctx context.Context, body TeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsValidatePropertiesPostRequestBodyable, requestConfiguration *TeamTemplatesItemDefinitionsItemTeamDefinitionPermissionGrantsValidatePropertiesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
