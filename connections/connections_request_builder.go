package connections

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    ifd0bef7cf3f9483c45767e76c4d696dee9baf1414a57fd07451c43ee66ea5501 "github.com/microsoftgraph/msgraph-beta-sdk-go/connections/count"
    i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/externalconnectors"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
)

// ConnectionsRequestBuilder provides operations to manage the collection of externalConnection entities.
type ConnectionsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ConnectionsRequestBuilderGetOptions options for Get
type ConnectionsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ConnectionsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ConnectionsRequestBuilderGetQueryParameters get entities from connections
type ConnectionsRequestBuilderGetQueryParameters struct {
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
// ConnectionsRequestBuilderPostOptions options for Post
type ConnectionsRequestBuilderPostOptions struct {
    // 
    Body i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367.ExternalConnectionable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewConnectionsRequestBuilderInternal instantiates a new ConnectionsRequestBuilder and sets the default values.
func NewConnectionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ConnectionsRequestBuilder) {
    m := &ConnectionsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/connections{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewConnectionsRequestBuilder instantiates a new ConnectionsRequestBuilder and sets the default values.
func NewConnectionsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ConnectionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConnectionsRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ConnectionsRequestBuilder) Count()(*ifd0bef7cf3f9483c45767e76c4d696dee9baf1414a57fd07451c43ee66ea5501.CountRequestBuilder) {
    return ifd0bef7cf3f9483c45767e76c4d696dee9baf1414a57fd07451c43ee66ea5501.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation get entities from connections
func (m *ConnectionsRequestBuilder) CreateGetRequestInformation(options *ConnectionsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePostRequestInformation add new entity to connections
func (m *ConnectionsRequestBuilder) CreatePostRequestInformation(options *ConnectionsRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get get entities from connections
func (m *ConnectionsRequestBuilder) Get(options *ConnectionsRequestBuilderGetOptions)(i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367.ExternalConnectionCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367.CreateExternalConnectionCollectionResponseFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367.ExternalConnectionCollectionResponseable), nil
}
// Post add new entity to connections
func (m *ConnectionsRequestBuilder) Post(options *ConnectionsRequestBuilderPostOptions)(i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367.ExternalConnectionable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367.CreateExternalConnectionFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367.ExternalConnectionable), nil
}
