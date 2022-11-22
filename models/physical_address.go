package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PhysicalAddress 
type PhysicalAddress struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The city.
    city *string
    // The country or region. It's a free-format string value, for example, 'United States'.
    countryOrRegion *string
    // The OdataType property
    odataType *string
    // The postal code.
    postalCode *string
    // The post office box number.
    postOfficeBox *string
    // The state.
    state *string
    // The street.
    street *string
    // The type of address. Possible values are: unknown, home, business, other.
    type_escaped *PhysicalAddressType
}
// NewPhysicalAddress instantiates a new physicalAddress and sets the default values.
func NewPhysicalAddress()(*PhysicalAddress) {
    m := &PhysicalAddress{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreatePhysicalAddressFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePhysicalAddressFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPhysicalAddress(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PhysicalAddress) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCity gets the city property value. The city.
func (m *PhysicalAddress) GetCity()(*string) {
    return m.city
}
// GetCountryOrRegion gets the countryOrRegion property value. The country or region. It's a free-format string value, for example, 'United States'.
func (m *PhysicalAddress) GetCountryOrRegion()(*string) {
    return m.countryOrRegion
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PhysicalAddress) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["city"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCity)
    res["countryOrRegion"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCountryOrRegion)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["postalCode"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPostalCode)
    res["postOfficeBox"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPostOfficeBox)
    res["state"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetState)
    res["street"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetStreet)
    res["type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParsePhysicalAddressType , m.SetType)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PhysicalAddress) GetOdataType()(*string) {
    return m.odataType
}
// GetPostalCode gets the postalCode property value. The postal code.
func (m *PhysicalAddress) GetPostalCode()(*string) {
    return m.postalCode
}
// GetPostOfficeBox gets the postOfficeBox property value. The post office box number.
func (m *PhysicalAddress) GetPostOfficeBox()(*string) {
    return m.postOfficeBox
}
// GetState gets the state property value. The state.
func (m *PhysicalAddress) GetState()(*string) {
    return m.state
}
// GetStreet gets the street property value. The street.
func (m *PhysicalAddress) GetStreet()(*string) {
    return m.street
}
// GetType gets the type property value. The type of address. Possible values are: unknown, home, business, other.
func (m *PhysicalAddress) GetType()(*PhysicalAddressType) {
    return m.type_escaped
}
// Serialize serializes information the current object
func (m *PhysicalAddress) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("city", m.GetCity())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("countryOrRegion", m.GetCountryOrRegion())
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
        err := writer.WriteStringValue("postalCode", m.GetPostalCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("postOfficeBox", m.GetPostOfficeBox())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("state", m.GetState())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("street", m.GetStreet())
        if err != nil {
            return err
        }
    }
    if m.GetType() != nil {
        cast := (*m.GetType()).String()
        err := writer.WriteStringValue("type", &cast)
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
func (m *PhysicalAddress) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCity sets the city property value. The city.
func (m *PhysicalAddress) SetCity(value *string)() {
    m.city = value
}
// SetCountryOrRegion sets the countryOrRegion property value. The country or region. It's a free-format string value, for example, 'United States'.
func (m *PhysicalAddress) SetCountryOrRegion(value *string)() {
    m.countryOrRegion = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PhysicalAddress) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPostalCode sets the postalCode property value. The postal code.
func (m *PhysicalAddress) SetPostalCode(value *string)() {
    m.postalCode = value
}
// SetPostOfficeBox sets the postOfficeBox property value. The post office box number.
func (m *PhysicalAddress) SetPostOfficeBox(value *string)() {
    m.postOfficeBox = value
}
// SetState sets the state property value. The state.
func (m *PhysicalAddress) SetState(value *string)() {
    m.state = value
}
// SetStreet sets the street property value. The street.
func (m *PhysicalAddress) SetStreet(value *string)() {
    m.street = value
}
// SetType sets the type property value. The type of address. Possible values are: unknown, home, business, other.
func (m *PhysicalAddress) SetType(value *PhysicalAddressType)() {
    m.type_escaped = value
}
