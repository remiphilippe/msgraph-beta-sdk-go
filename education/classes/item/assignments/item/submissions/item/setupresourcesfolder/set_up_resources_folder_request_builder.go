package setupresourcesfolder

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// SetUpResourcesFolderRequestBuilder provides operations to call the setUpResourcesFolder method.
type SetUpResourcesFolderRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// SetUpResourcesFolderRequestBuilderPostOptions options for Post
type SetUpResourcesFolderRequestBuilderPostOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// SetUpResourcesFolderResponse union type wrapper for classes educationSubmission
type SetUpResourcesFolderResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Union type representation for type educationSubmission
    educationSubmission i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EducationSubmissionable;
}
// NewSetUpResourcesFolderResponse instantiates a new setUpResourcesFolderResponse and sets the default values.
func NewSetUpResourcesFolderResponse()(*SetUpResourcesFolderResponse) {
    m := &SetUpResourcesFolderResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func CreateSetUpResourcesFolderResponseFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewSetUpResourcesFolderResponse(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SetUpResourcesFolderResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetEducationSubmission gets the educationSubmission property value. Union type representation for type educationSubmission
func (m *SetUpResourcesFolderResponse) GetEducationSubmission()(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EducationSubmissionable) {
    if m == nil {
        return nil
    } else {
        return m.educationSubmission
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SetUpResourcesFolderResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["educationSubmission"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateEducationSubmissionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEducationSubmission(val.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EducationSubmissionable))
        }
        return nil
    }
    return res
}
func (m *SetUpResourcesFolderResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *SetUpResourcesFolderResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("educationSubmission", m.GetEducationSubmission())
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
func (m *SetUpResourcesFolderResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetEducationSubmission sets the educationSubmission property value. Union type representation for type educationSubmission
func (m *SetUpResourcesFolderResponse) SetEducationSubmission(value i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EducationSubmissionable)() {
    if m != nil {
        m.educationSubmission = value
    }
}
// SetUpResourcesFolderResponseable 
type SetUpResourcesFolderResponseable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetEducationSubmission()(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EducationSubmissionable)
    SetEducationSubmission(value i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EducationSubmissionable)()
}
// NewSetUpResourcesFolderRequestBuilderInternal instantiates a new SetUpResourcesFolderRequestBuilder and sets the default values.
func NewSetUpResourcesFolderRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SetUpResourcesFolderRequestBuilder) {
    m := &SetUpResourcesFolderRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/education/classes/{educationClass_id}/assignments/{educationAssignment_id}/submissions/{educationSubmission_id}/microsoft.graph.setUpResourcesFolder";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSetUpResourcesFolderRequestBuilder instantiates a new SetUpResourcesFolderRequestBuilder and sets the default values.
func NewSetUpResourcesFolderRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SetUpResourcesFolderRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSetUpResourcesFolderRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action setUpResourcesFolder
func (m *SetUpResourcesFolderRequestBuilder) CreatePostRequestInformation(options *SetUpResourcesFolderRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Post invoke action setUpResourcesFolder
func (m *SetUpResourcesFolderRequestBuilder) Post(options *SetUpResourcesFolderRequestBuilderPostOptions)(SetUpResourcesFolderResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateSetUpResourcesFolderResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(SetUpResourcesFolderResponseable), nil
}
