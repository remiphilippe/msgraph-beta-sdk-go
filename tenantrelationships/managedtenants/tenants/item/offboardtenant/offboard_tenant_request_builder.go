package offboardtenant

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/managedtenants"
)

// OffboardTenantRequestBuilder provides operations to call the offboardTenant method.
type OffboardTenantRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// OffboardTenantRequestBuilderPostOptions options for Post
type OffboardTenantRequestBuilderPostOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// OffboardTenantResponse union type wrapper for classes tenant
type OffboardTenantResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Union type representation for type tenant
    tenant i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.Tenantable;
}
// NewOffboardTenantResponse instantiates a new offboardTenantResponse and sets the default values.
func NewOffboardTenantResponse()(*OffboardTenantResponse) {
    m := &OffboardTenantResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func CreateOffboardTenantResponseFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewOffboardTenantResponse(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OffboardTenantResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OffboardTenantResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["tenant"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.CreateTenantFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenant(val.(i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.Tenantable))
        }
        return nil
    }
    return res
}
// GetTenant gets the tenant property value. Union type representation for type tenant
func (m *OffboardTenantResponse) GetTenant()(i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.Tenantable) {
    if m == nil {
        return nil
    } else {
        return m.tenant
    }
}
func (m *OffboardTenantResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *OffboardTenantResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("tenant", m.GetTenant())
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
func (m *OffboardTenantResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetTenant sets the tenant property value. Union type representation for type tenant
func (m *OffboardTenantResponse) SetTenant(value i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.Tenantable)() {
    if m != nil {
        m.tenant = value
    }
}
// OffboardTenantResponseable 
type OffboardTenantResponseable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetTenant()(i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.Tenantable)
    SetTenant(value i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.Tenantable)()
}
// NewOffboardTenantRequestBuilderInternal instantiates a new OffboardTenantRequestBuilder and sets the default values.
func NewOffboardTenantRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OffboardTenantRequestBuilder) {
    m := &OffboardTenantRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/tenantRelationships/managedTenants/tenants/{tenant_id}/microsoft.graph.managedTenants.offboardTenant";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOffboardTenantRequestBuilder instantiates a new OffboardTenantRequestBuilder and sets the default values.
func NewOffboardTenantRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OffboardTenantRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOffboardTenantRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action offboardTenant
func (m *OffboardTenantRequestBuilder) CreatePostRequestInformation(options *OffboardTenantRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.POST
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
// Post invoke action offboardTenant
func (m *OffboardTenantRequestBuilder) Post(options *OffboardTenantRequestBuilderPostOptions)(OffboardTenantResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateOffboardTenantResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(OffboardTenantResponseable), nil
}
