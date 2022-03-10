package groups

import (
    i1bf6f351562689bc6b2c2007b10557f838392b03627c557ff2caa21135479027 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/validateproperties"
    i498ffd1b2f6044bdb24e5841ce70347da2962279c830c4f8e0c63e4185ecc573 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/evaluatedynamicmembership"
    i86542b1985ccd1d6e9cd7f0b9657c459f0a8d024d32198db68c5cfe8de1e0290 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/count"
    i9a486fb0e4b07920959c03436a2a48c8a425e7ee639e60bb6b65abd12235a827 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/getuserownedobjects"
    ib37f70f6afd357219dcd43a10bc09e3ce47b8bceed59b5516cc7481626e189c7 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/getbyids"
    ic020eb24df81d63ab670349e4923f4de66eebfcadab8e631fb8c7e67bea34738 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/delta"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
)

// GroupsRequestBuilder provides operations to manage the collection of group entities.
type GroupsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GroupsRequestBuilderGetOptions options for Get
type GroupsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *GroupsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// GroupsRequestBuilderGetQueryParameters get entities from groups
type GroupsRequestBuilderGetQueryParameters struct {
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
// GroupsRequestBuilderPostOptions options for Post
type GroupsRequestBuilderPostOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Groupable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewGroupsRequestBuilderInternal instantiates a new GroupsRequestBuilder and sets the default values.
func NewGroupsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GroupsRequestBuilder) {
    m := &GroupsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGroupsRequestBuilder instantiates a new GroupsRequestBuilder and sets the default values.
func NewGroupsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GroupsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupsRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *GroupsRequestBuilder) Count()(*i86542b1985ccd1d6e9cd7f0b9657c459f0a8d024d32198db68c5cfe8de1e0290.CountRequestBuilder) {
    return i86542b1985ccd1d6e9cd7f0b9657c459f0a8d024d32198db68c5cfe8de1e0290.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation get entities from groups
func (m *GroupsRequestBuilder) CreateGetRequestInformation(options *GroupsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePostRequestInformation add new entity to groups
func (m *GroupsRequestBuilder) CreatePostRequestInformation(options *GroupsRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.POST
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delta provides operations to call the delta method.
func (m *GroupsRequestBuilder) Delta()(*ic020eb24df81d63ab670349e4923f4de66eebfcadab8e631fb8c7e67bea34738.DeltaRequestBuilder) {
    return ic020eb24df81d63ab670349e4923f4de66eebfcadab8e631fb8c7e67bea34738.NewDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupsRequestBuilder) EvaluateDynamicMembership()(*i498ffd1b2f6044bdb24e5841ce70347da2962279c830c4f8e0c63e4185ecc573.EvaluateDynamicMembershipRequestBuilder) {
    return i498ffd1b2f6044bdb24e5841ce70347da2962279c830c4f8e0c63e4185ecc573.NewEvaluateDynamicMembershipRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get entities from groups
func (m *GroupsRequestBuilder) Get(options *GroupsRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GroupCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateGroupCollectionResponseFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GroupCollectionResponseable), nil
}
func (m *GroupsRequestBuilder) GetByIds()(*ib37f70f6afd357219dcd43a10bc09e3ce47b8bceed59b5516cc7481626e189c7.GetByIdsRequestBuilder) {
    return ib37f70f6afd357219dcd43a10bc09e3ce47b8bceed59b5516cc7481626e189c7.NewGetByIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupsRequestBuilder) GetUserOwnedObjects()(*i9a486fb0e4b07920959c03436a2a48c8a425e7ee639e60bb6b65abd12235a827.GetUserOwnedObjectsRequestBuilder) {
    return i9a486fb0e4b07920959c03436a2a48c8a425e7ee639e60bb6b65abd12235a827.NewGetUserOwnedObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Post add new entity to groups
func (m *GroupsRequestBuilder) Post(options *GroupsRequestBuilderPostOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Groupable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateGroupFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Groupable), nil
}
func (m *GroupsRequestBuilder) ValidateProperties()(*i1bf6f351562689bc6b2c2007b10557f838392b03627c557ff2caa21135479027.ValidatePropertiesRequestBuilder) {
    return i1bf6f351562689bc6b2c2007b10557f838392b03627c557ff2caa21135479027.NewValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
