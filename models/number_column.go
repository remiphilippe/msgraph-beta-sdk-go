package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// NumberColumn 
type NumberColumn struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // How many decimal places to display. See below for information about the possible values.
    decimalPlaces *string
    // How the value should be presented in the UX. Must be one of number or percentage. If unspecified, treated as number.
    displayAs *string
    // The maximum permitted value.
    maximum *float64
    // The minimum permitted value.
    minimum *float64
    // The OdataType property
    odataType *string
}
// NewNumberColumn instantiates a new numberColumn and sets the default values.
func NewNumberColumn()(*NumberColumn) {
    m := &NumberColumn{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateNumberColumnFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateNumberColumnFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNumberColumn(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *NumberColumn) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDecimalPlaces gets the decimalPlaces property value. How many decimal places to display. See below for information about the possible values.
func (m *NumberColumn) GetDecimalPlaces()(*string) {
    return m.decimalPlaces
}
// GetDisplayAs gets the displayAs property value. How the value should be presented in the UX. Must be one of number or percentage. If unspecified, treated as number.
func (m *NumberColumn) GetDisplayAs()(*string) {
    return m.displayAs
}
// GetFieldDeserializers the deserialization information for the current model
func (m *NumberColumn) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["decimalPlaces"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDecimalPlaces)
    res["displayAs"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayAs)
    res["maximum"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetMaximum)
    res["minimum"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetMinimum)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    return res
}
// GetMaximum gets the maximum property value. The maximum permitted value.
func (m *NumberColumn) GetMaximum()(*float64) {
    return m.maximum
}
// GetMinimum gets the minimum property value. The minimum permitted value.
func (m *NumberColumn) GetMinimum()(*float64) {
    return m.minimum
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *NumberColumn) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *NumberColumn) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("decimalPlaces", m.GetDecimalPlaces())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayAs", m.GetDisplayAs())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("maximum", m.GetMaximum())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("minimum", m.GetMinimum())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *NumberColumn) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDecimalPlaces sets the decimalPlaces property value. How many decimal places to display. See below for information about the possible values.
func (m *NumberColumn) SetDecimalPlaces(value *string)() {
    m.decimalPlaces = value
}
// SetDisplayAs sets the displayAs property value. How the value should be presented in the UX. Must be one of number or percentage. If unspecified, treated as number.
func (m *NumberColumn) SetDisplayAs(value *string)() {
    m.displayAs = value
}
// SetMaximum sets the maximum property value. The maximum permitted value.
func (m *NumberColumn) SetMaximum(value *float64)() {
    m.maximum = value
}
// SetMinimum sets the minimum property value. The minimum permitted value.
func (m *NumberColumn) SetMinimum(value *float64)() {
    m.minimum = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *NumberColumn) SetOdataType(value *string)() {
    m.odataType = value
}
