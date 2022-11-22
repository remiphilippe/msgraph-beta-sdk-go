package security

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// Security 
type Security struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // The informationProtection property
    informationProtection InformationProtectionable
}
// NewSecurity instantiates a new security and sets the default values.
func NewSecurity()(*Security) {
    m := &Security{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateSecurityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSecurityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecurity(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Security) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["informationProtection"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateInformationProtectionFromDiscriminatorValue , m.SetInformationProtection)
    return res
}
// GetInformationProtection gets the informationProtection property value. The informationProtection property
func (m *Security) GetInformationProtection()(InformationProtectionable) {
    return m.informationProtection
}
// Serialize serializes information the current object
func (m *Security) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("informationProtection", m.GetInformationProtection())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetInformationProtection sets the informationProtection property value. The informationProtection property
func (m *Security) SetInformationProtection(value InformationProtectionable)() {
    m.informationProtection = value
}
