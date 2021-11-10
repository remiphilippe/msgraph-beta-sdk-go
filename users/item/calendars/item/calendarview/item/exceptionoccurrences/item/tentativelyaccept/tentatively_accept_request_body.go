package tentativelyaccept

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// 
type TentativelyAcceptRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    comment *string;
    // 
    proposedNewTime *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TimeSlot;
    // 
    sendResponse *bool;
}
// Instantiates a new tentativelyAcceptRequestBody and sets the default values.
func NewTentativelyAcceptRequestBody()(*TentativelyAcceptRequestBody) {
    m := &TentativelyAcceptRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TentativelyAcceptRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the Comment property value. 
func (m *TentativelyAcceptRequestBody) GetComment()(*string) {
    if m == nil {
        return nil
    } else {
        return m.comment
    }
}
// Gets the ProposedNewTime property value. 
func (m *TentativelyAcceptRequestBody) GetProposedNewTime()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TimeSlot) {
    if m == nil {
        return nil
    } else {
        return m.proposedNewTime
    }
}
// Gets the SendResponse property value. 
func (m *TentativelyAcceptRequestBody) GetSendResponse()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.sendResponse
    }
}
// The deserialization information for the current model
func (m *TentativelyAcceptRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["comment"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComment(val)
        }
        return nil
    }
    res["proposedNewTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewTimeSlot() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProposedNewTime(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TimeSlot))
        }
        return nil
    }
    res["sendResponse"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSendResponse(val)
        }
        return nil
    }
    return res
}
func (m *TentativelyAcceptRequestBody) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *TentativelyAcceptRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("comment", m.GetComment())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("proposedNewTime", m.GetProposedNewTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("sendResponse", m.GetSendResponse())
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *TentativelyAcceptRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the Comment property value. 
// Parameters:
//  - value : Value to set for the Comment property.
func (m *TentativelyAcceptRequestBody) SetComment(value *string)() {
    m.comment = value
}
// Sets the ProposedNewTime property value. 
// Parameters:
//  - value : Value to set for the ProposedNewTime property.
func (m *TentativelyAcceptRequestBody) SetProposedNewTime(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TimeSlot)() {
    m.proposedNewTime = value
}
// Sets the SendResponse property value. 
// Parameters:
//  - value : Value to set for the SendResponse property.
func (m *TentativelyAcceptRequestBody) SetSendResponse(value *bool)() {
    m.sendResponse = value
}
