package applications

import (
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemRemoveKeyPostRequestBody 
type ItemRemoveKeyPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The keyId property
    keyId *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // The proof property
    proof *string
}
// NewItemRemoveKeyPostRequestBody instantiates a new ItemRemoveKeyPostRequestBody and sets the default values.
func NewItemRemoveKeyPostRequestBody()(*ItemRemoveKeyPostRequestBody) {
    m := &ItemRemoveKeyPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any));
    return m
}
// CreateItemRemoveKeyPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemRemoveKeyPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemRemoveKeyPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemRemoveKeyPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemRemoveKeyPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["keyId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyId(val)
        }
        return nil
    }
    res["proof"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProof(val)
        }
        return nil
    }
    return res
}
// GetKeyId gets the keyId property value. The keyId property
func (m *ItemRemoveKeyPostRequestBody) GetKeyId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.keyId
}
// GetProof gets the proof property value. The proof property
func (m *ItemRemoveKeyPostRequestBody) GetProof()(*string) {
    return m.proof
}
// Serialize serializes information the current object
func (m *ItemRemoveKeyPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteUUIDValue("keyId", m.GetKeyId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("proof", m.GetProof())
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
func (m *ItemRemoveKeyPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetKeyId sets the keyId property value. The keyId property
func (m *ItemRemoveKeyPostRequestBody) SetKeyId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.keyId = value
}
// SetProof sets the proof property value. The proof property
func (m *ItemRemoveKeyPostRequestBody) SetProof(value *string)() {
    m.proof = value
}
