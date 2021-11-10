package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type Picture struct {
    Entity
    // 
    content []byte;
    // 
    contentType *string;
    // 
    height *int32;
    // 
    width *int32;
}
// Instantiates a new picture and sets the default values.
func NewPicture()(*Picture) {
    m := &Picture{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the content property value. 
func (m *Picture) GetContent()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.content
    }
}
// Gets the contentType property value. 
func (m *Picture) GetContentType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contentType
    }
}
// Gets the height property value. 
func (m *Picture) GetHeight()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.height
    }
}
// Gets the width property value. 
func (m *Picture) GetWidth()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.width
    }
}
// The deserialization information for the current model
func (m *Picture) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["content"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContent(val)
        }
        return nil
    }
    res["contentType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentType(val)
        }
        return nil
    }
    res["height"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHeight(val)
        }
        return nil
    }
    res["width"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWidth(val)
        }
        return nil
    }
    return res
}
func (m *Picture) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *Picture) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteByteArrayValue("content", m.GetContent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("contentType", m.GetContentType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("height", m.GetHeight())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("width", m.GetWidth())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the content property value. 
// Parameters:
//  - value : Value to set for the content property.
func (m *Picture) SetContent(value []byte)() {
    m.content = value
}
// Sets the contentType property value. 
// Parameters:
//  - value : Value to set for the contentType property.
func (m *Picture) SetContentType(value *string)() {
    m.contentType = value
}
// Sets the height property value. 
// Parameters:
//  - value : Value to set for the height property.
func (m *Picture) SetHeight(value *int32)() {
    m.height = value
}
// Sets the width property value. 
// Parameters:
//  - value : Value to set for the width property.
func (m *Picture) SetWidth(value *int32)() {
    m.width = value
}
