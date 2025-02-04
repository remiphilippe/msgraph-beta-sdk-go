package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CustomSubjectAlternativeName custom Subject Alternative Name definition
type CustomSubjectAlternativeName struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Custom SAN Name
    name *string
    // The OdataType property
    odataType *string
    // Subject Alternative Name Options.
    sanType *SubjectAlternativeNameType
}
// NewCustomSubjectAlternativeName instantiates a new customSubjectAlternativeName and sets the default values.
func NewCustomSubjectAlternativeName()(*CustomSubjectAlternativeName) {
    m := &CustomSubjectAlternativeName{
    }
    m.SetAdditionalData(make(map[string]any));
    return m
}
// CreateCustomSubjectAlternativeNameFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCustomSubjectAlternativeNameFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCustomSubjectAlternativeName(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CustomSubjectAlternativeName) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CustomSubjectAlternativeName) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["sanType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSubjectAlternativeNameType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSanType(val.(*SubjectAlternativeNameType))
        }
        return nil
    }
    return res
}
// GetName gets the name property value. Custom SAN Name
func (m *CustomSubjectAlternativeName) GetName()(*string) {
    return m.name
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *CustomSubjectAlternativeName) GetOdataType()(*string) {
    return m.odataType
}
// GetSanType gets the sanType property value. Subject Alternative Name Options.
func (m *CustomSubjectAlternativeName) GetSanType()(*SubjectAlternativeNameType) {
    return m.sanType
}
// Serialize serializes information the current object
func (m *CustomSubjectAlternativeName) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetSanType() != nil {
        cast := (*m.GetSanType()).String()
        err := writer.WriteStringValue("sanType", &cast)
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
func (m *CustomSubjectAlternativeName) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetName sets the name property value. Custom SAN Name
func (m *CustomSubjectAlternativeName) SetName(value *string)() {
    m.name = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CustomSubjectAlternativeName) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSanType sets the sanType property value. Subject Alternative Name Options.
func (m *CustomSubjectAlternativeName) SetSanType(value *SubjectAlternativeNameType)() {
    m.sanType = value
}
