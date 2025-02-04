package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IntegerRange 
type IntegerRange struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The inclusive upper bound of the integer range.
    end *int64
    // The maximum property
    maximum *int64
    // The minimum property
    minimum *int64
    // The OdataType property
    odataType *string
    // The inclusive lower bound of the integer range.
    start *int64
}
// NewIntegerRange instantiates a new integerRange and sets the default values.
func NewIntegerRange()(*IntegerRange) {
    m := &IntegerRange{
    }
    m.SetAdditionalData(make(map[string]any));
    return m
}
// CreateIntegerRangeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIntegerRangeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIntegerRange(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IntegerRange) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEnd gets the end property value. The inclusive upper bound of the integer range.
func (m *IntegerRange) GetEnd()(*int64) {
    return m.end
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IntegerRange) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["end"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnd(val)
        }
        return nil
    }
    res["maximum"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximum(val)
        }
        return nil
    }
    res["minimum"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimum(val)
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
    res["start"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStart(val)
        }
        return nil
    }
    return res
}
// GetMaximum gets the maximum property value. The maximum property
func (m *IntegerRange) GetMaximum()(*int64) {
    return m.maximum
}
// GetMinimum gets the minimum property value. The minimum property
func (m *IntegerRange) GetMinimum()(*int64) {
    return m.minimum
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *IntegerRange) GetOdataType()(*string) {
    return m.odataType
}
// GetStart gets the start property value. The inclusive lower bound of the integer range.
func (m *IntegerRange) GetStart()(*int64) {
    return m.start
}
// Serialize serializes information the current object
func (m *IntegerRange) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("end", m.GetEnd())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("maximum", m.GetMaximum())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("minimum", m.GetMinimum())
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
    {
        err := writer.WriteInt64Value("start", m.GetStart())
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
func (m *IntegerRange) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEnd sets the end property value. The inclusive upper bound of the integer range.
func (m *IntegerRange) SetEnd(value *int64)() {
    m.end = value
}
// SetMaximum sets the maximum property value. The maximum property
func (m *IntegerRange) SetMaximum(value *int64)() {
    m.maximum = value
}
// SetMinimum sets the minimum property value. The minimum property
func (m *IntegerRange) SetMinimum(value *int64)() {
    m.minimum = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *IntegerRange) SetOdataType(value *string)() {
    m.odataType = value
}
// SetStart sets the start property value. The inclusive lower bound of the integer range.
func (m *IntegerRange) SetStart(value *int64)() {
    m.start = value
}
