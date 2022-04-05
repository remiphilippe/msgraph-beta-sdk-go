package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i072f978c5e76c8f6f7f146e9adbc77e18bc617a145954334b9d760baaf124b24 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/synchronization/jobs/item/schema"
    i1037a2ef85bc7c104d93f4c5a1e5a28ff559376fe6cb1bb568e1d1c84273c5d7 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/synchronization/jobs/item/stop"
    i17256af27a6172f385302adc8f2924bddbaa8027a02a4e613619e8304ec94fe5 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/synchronization/jobs/item/provisionondemand"
    i37054552034248ca68fec46390fca0923f7a2e6c10539e7a67204b742ef8e9fb "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/synchronization/jobs/item/restart"
    ia80ae8b8dfc2c4ddfe68e3d7f5fb7158905c34e938b0cc1afbae2798e432005b "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/synchronization/jobs/item/validatecredentials"
    id1feb9558d42229603da95cf0e6bf17dfe11c8a034277b3ec93c1b87e7a947a7 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/synchronization/jobs/item/pause"
    ie7dd06601f7be4a75e499cc220f5c87d45dca4695a9b092b00f512bf0edb0643 "github.com/microsoftgraph/msgraph-beta-sdk-go/serviceprincipals/item/synchronization/jobs/item/start"
)

// SynchronizationJobItemRequestBuilder provides operations to manage the jobs property of the microsoft.graph.synchronization entity.
type SynchronizationJobItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// SynchronizationJobItemRequestBuilderDeleteOptions options for Delete
type SynchronizationJobItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// SynchronizationJobItemRequestBuilderGetOptions options for Get
type SynchronizationJobItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *SynchronizationJobItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// SynchronizationJobItemRequestBuilderGetQueryParameters performs synchronization by periodically running in the background, polling for changes in one directory, and pushing them to another directory.
type SynchronizationJobItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// SynchronizationJobItemRequestBuilderPatchOptions options for Patch
type SynchronizationJobItemRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SynchronizationJobable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// NewSynchronizationJobItemRequestBuilderInternal instantiates a new SynchronizationJobItemRequestBuilder and sets the default values.
func NewSynchronizationJobItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SynchronizationJobItemRequestBuilder) {
    m := &SynchronizationJobItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/servicePrincipals/{servicePrincipal_id}/synchronization/jobs/{synchronizationJob_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSynchronizationJobItemRequestBuilder instantiates a new SynchronizationJobItemRequestBuilder and sets the default values.
func NewSynchronizationJobItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SynchronizationJobItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSynchronizationJobItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property jobs for servicePrincipals
func (m *SynchronizationJobItemRequestBuilder) CreateDeleteRequestInformation(options *SynchronizationJobItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation performs synchronization by periodically running in the background, polling for changes in one directory, and pushing them to another directory.
func (m *SynchronizationJobItemRequestBuilder) CreateGetRequestInformation(options *SynchronizationJobItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property jobs in servicePrincipals
func (m *SynchronizationJobItemRequestBuilder) CreatePatchRequestInformation(options *SynchronizationJobItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property jobs for servicePrincipals
func (m *SynchronizationJobItemRequestBuilder) Delete(options *SynchronizationJobItemRequestBuilderDeleteOptions)(error) {
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
// Get performs synchronization by periodically running in the background, polling for changes in one directory, and pushing them to another directory.
func (m *SynchronizationJobItemRequestBuilder) Get(options *SynchronizationJobItemRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SynchronizationJobable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSynchronizationJobFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SynchronizationJobable), nil
}
// Patch update the navigation property jobs in servicePrincipals
func (m *SynchronizationJobItemRequestBuilder) Patch(options *SynchronizationJobItemRequestBuilderPatchOptions)(error) {
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
// Pause the pause property
func (m *SynchronizationJobItemRequestBuilder) Pause()(*id1feb9558d42229603da95cf0e6bf17dfe11c8a034277b3ec93c1b87e7a947a7.PauseRequestBuilder) {
    return id1feb9558d42229603da95cf0e6bf17dfe11c8a034277b3ec93c1b87e7a947a7.NewPauseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ProvisionOnDemand the provisionOnDemand property
func (m *SynchronizationJobItemRequestBuilder) ProvisionOnDemand()(*i17256af27a6172f385302adc8f2924bddbaa8027a02a4e613619e8304ec94fe5.ProvisionOnDemandRequestBuilder) {
    return i17256af27a6172f385302adc8f2924bddbaa8027a02a4e613619e8304ec94fe5.NewProvisionOnDemandRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restart the restart property
func (m *SynchronizationJobItemRequestBuilder) Restart()(*i37054552034248ca68fec46390fca0923f7a2e6c10539e7a67204b742ef8e9fb.RestartRequestBuilder) {
    return i37054552034248ca68fec46390fca0923f7a2e6c10539e7a67204b742ef8e9fb.NewRestartRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Schema the schema property
func (m *SynchronizationJobItemRequestBuilder) Schema()(*i072f978c5e76c8f6f7f146e9adbc77e18bc617a145954334b9d760baaf124b24.SchemaRequestBuilder) {
    return i072f978c5e76c8f6f7f146e9adbc77e18bc617a145954334b9d760baaf124b24.NewSchemaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Start the start property
func (m *SynchronizationJobItemRequestBuilder) Start()(*ie7dd06601f7be4a75e499cc220f5c87d45dca4695a9b092b00f512bf0edb0643.StartRequestBuilder) {
    return ie7dd06601f7be4a75e499cc220f5c87d45dca4695a9b092b00f512bf0edb0643.NewStartRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Stop the stop property
func (m *SynchronizationJobItemRequestBuilder) Stop()(*i1037a2ef85bc7c104d93f4c5a1e5a28ff559376fe6cb1bb568e1d1c84273c5d7.StopRequestBuilder) {
    return i1037a2ef85bc7c104d93f4c5a1e5a28ff559376fe6cb1bb568e1d1c84273c5d7.NewStopRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ValidateCredentials the validateCredentials property
func (m *SynchronizationJobItemRequestBuilder) ValidateCredentials()(*ia80ae8b8dfc2c4ddfe68e3d7f5fb7158905c34e938b0cc1afbae2798e432005b.ValidateCredentialsRequestBuilder) {
    return ia80ae8b8dfc2c4ddfe68e3d7f5fb7158905c34e938b0cc1afbae2798e432005b.NewValidateCredentialsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
