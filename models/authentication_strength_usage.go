package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuthenticationStrengthUsage 
type AuthenticationStrengthUsage struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The mfa property
    mfa []ConditionalAccessPolicyable
    // The none property
    none []ConditionalAccessPolicyable
    // The OdataType property
    odataType *string
}
// NewAuthenticationStrengthUsage instantiates a new authenticationStrengthUsage and sets the default values.
func NewAuthenticationStrengthUsage()(*AuthenticationStrengthUsage) {
    m := &AuthenticationStrengthUsage{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAuthenticationStrengthUsageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuthenticationStrengthUsageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthenticationStrengthUsage(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AuthenticationStrengthUsage) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuthenticationStrengthUsage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["mfa"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateConditionalAccessPolicyFromDiscriminatorValue , m.SetMfa)
    res["none"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateConditionalAccessPolicyFromDiscriminatorValue , m.SetNone)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    return res
}
// GetMfa gets the mfa property value. The mfa property
func (m *AuthenticationStrengthUsage) GetMfa()([]ConditionalAccessPolicyable) {
    return m.mfa
}
// GetNone gets the none property value. The none property
func (m *AuthenticationStrengthUsage) GetNone()([]ConditionalAccessPolicyable) {
    return m.none
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AuthenticationStrengthUsage) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *AuthenticationStrengthUsage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetMfa() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetMfa())
        err := writer.WriteCollectionOfObjectValues("mfa", cast)
        if err != nil {
            return err
        }
    }
    if m.GetNone() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetNone())
        err := writer.WriteCollectionOfObjectValues("none", cast)
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
func (m *AuthenticationStrengthUsage) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetMfa sets the mfa property value. The mfa property
func (m *AuthenticationStrengthUsage) SetMfa(value []ConditionalAccessPolicyable)() {
    m.mfa = value
}
// SetNone sets the none property value. The none property
func (m *AuthenticationStrengthUsage) SetNone(value []ConditionalAccessPolicyable)() {
    m.none = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AuthenticationStrengthUsage) SetOdataType(value *string)() {
    m.odataType = value
}
