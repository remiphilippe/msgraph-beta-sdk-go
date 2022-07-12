package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CommunicationsUserIdentity 
type CommunicationsUserIdentity struct {
    Identity
    // The tenantId property
    tenantId *string
}
// NewCommunicationsUserIdentity instantiates a new CommunicationsUserIdentity and sets the default values.
func NewCommunicationsUserIdentity()(*CommunicationsUserIdentity) {
    m := &CommunicationsUserIdentity{
        Identity: *NewIdentity(),
    }
    return m
}
// CreateCommunicationsUserIdentityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCommunicationsUserIdentityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCommunicationsUserIdentity(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CommunicationsUserIdentity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Identity.GetFieldDeserializers()
    res["tenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantId(val)
        }
        return nil
    }
    return res
}
// GetTenantId gets the tenantId property value. The tenantId property
func (m *CommunicationsUserIdentity) GetTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantId
    }
}
// Serialize serializes information the current object
func (m *CommunicationsUserIdentity) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Identity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("tenantId", m.GetTenantId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetTenantId sets the tenantId property value. The tenantId property
func (m *CommunicationsUserIdentity) SetTenantId(value *string)() {
    if m != nil {
        m.tenantId = value
    }
}
