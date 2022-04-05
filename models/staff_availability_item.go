package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// StaffAvailabilityItem 
type StaffAvailabilityItem struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The availabilityItems property
    availabilityItems []AvailabilityItemable;
    // The staffId property
    staffId *string;
}
// NewStaffAvailabilityItem instantiates a new staffAvailabilityItem and sets the default values.
func NewStaffAvailabilityItem()(*StaffAvailabilityItem) {
    m := &StaffAvailabilityItem{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateStaffAvailabilityItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateStaffAvailabilityItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewStaffAvailabilityItem(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *StaffAvailabilityItem) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAvailabilityItems gets the availabilityItems property value. The availabilityItems property
func (m *StaffAvailabilityItem) GetAvailabilityItems()([]AvailabilityItemable) {
    if m == nil {
        return nil
    } else {
        return m.availabilityItems
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *StaffAvailabilityItem) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["availabilityItems"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAvailabilityItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AvailabilityItemable, len(val))
            for i, v := range val {
                res[i] = v.(AvailabilityItemable)
            }
            m.SetAvailabilityItems(res)
        }
        return nil
    }
    res["staffId"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStaffId(val)
        }
        return nil
    }
    return res
}
// GetStaffId gets the staffId property value. The staffId property
func (m *StaffAvailabilityItem) GetStaffId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.staffId
    }
}
// Serialize serializes information the current object
func (m *StaffAvailabilityItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAvailabilityItems() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAvailabilityItems()))
        for i, v := range m.GetAvailabilityItems() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("availabilityItems", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("staffId", m.GetStaffId())
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
func (m *StaffAvailabilityItem) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAvailabilityItems sets the availabilityItems property value. The availabilityItems property
func (m *StaffAvailabilityItem) SetAvailabilityItems(value []AvailabilityItemable)() {
    if m != nil {
        m.availabilityItems = value
    }
}
// SetStaffId sets the staffId property value. The staffId property
func (m *StaffAvailabilityItem) SetStaffId(value *string)() {
    if m != nil {
        m.staffId = value
    }
}
