package apply

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/managedtenants"
)

// ApplyRequestBuilder provides operations to call the apply method.
type ApplyRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ApplyRequestBuilderPostOptions options for Post
type ApplyRequestBuilderPostOptions struct {
    // 
    Body ApplyRequestBodyable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ApplyResponse union type wrapper for classes managementActionDeploymentStatus
type ApplyResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Union type representation for type managementActionDeploymentStatus
    managementActionDeploymentStatus i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementActionDeploymentStatusable;
}
// NewApplyResponse instantiates a new applyResponse and sets the default values.
func NewApplyResponse()(*ApplyResponse) {
    m := &ApplyResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func CreateApplyResponseFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewApplyResponse(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ApplyResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ApplyResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["managementActionDeploymentStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.CreateManagementActionDeploymentStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementActionDeploymentStatus(val.(i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementActionDeploymentStatusable))
        }
        return nil
    }
    return res
}
// GetManagementActionDeploymentStatus gets the managementActionDeploymentStatus property value. Union type representation for type managementActionDeploymentStatus
func (m *ApplyResponse) GetManagementActionDeploymentStatus()(i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementActionDeploymentStatusable) {
    if m == nil {
        return nil
    } else {
        return m.managementActionDeploymentStatus
    }
}
func (m *ApplyResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ApplyResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("managementActionDeploymentStatus", m.GetManagementActionDeploymentStatus())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ApplyResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetManagementActionDeploymentStatus sets the managementActionDeploymentStatus property value. Union type representation for type managementActionDeploymentStatus
func (m *ApplyResponse) SetManagementActionDeploymentStatus(value i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementActionDeploymentStatusable)() {
    if m != nil {
        m.managementActionDeploymentStatus = value
    }
}
// ApplyResponseable 
type ApplyResponseable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetManagementActionDeploymentStatus()(i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementActionDeploymentStatusable)
    SetManagementActionDeploymentStatus(value i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementActionDeploymentStatusable)()
}
// NewApplyRequestBuilderInternal instantiates a new ApplyRequestBuilder and sets the default values.
func NewApplyRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ApplyRequestBuilder) {
    m := &ApplyRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/tenantRelationships/managedTenants/managementActions/{managementAction_id}/microsoft.graph.managedTenants.apply";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewApplyRequestBuilder instantiates a new ApplyRequestBuilder and sets the default values.
func NewApplyRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ApplyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewApplyRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action apply
func (m *ApplyRequestBuilder) CreatePostRequestInformation(options *ApplyRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post invoke action apply
func (m *ApplyRequestBuilder) Post(options *ApplyRequestBuilderPostOptions)(ApplyResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateApplyResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(ApplyResponseable), nil
}
