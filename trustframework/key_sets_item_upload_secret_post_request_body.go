package trustframework

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// KeySetsItemUploadSecretPostRequestBody 
type KeySetsItemUploadSecretPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The exp property
    exp *int64
    // The k property
    k *string
    // The nbf property
    nbf *int64
    // The use property
    use *string
}
// NewKeySetsItemUploadSecretPostRequestBody instantiates a new KeySetsItemUploadSecretPostRequestBody and sets the default values.
func NewKeySetsItemUploadSecretPostRequestBody()(*KeySetsItemUploadSecretPostRequestBody) {
    m := &KeySetsItemUploadSecretPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any));
    return m
}
// CreateKeySetsItemUploadSecretPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateKeySetsItemUploadSecretPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewKeySetsItemUploadSecretPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *KeySetsItemUploadSecretPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetExp gets the exp property value. The exp property
func (m *KeySetsItemUploadSecretPostRequestBody) GetExp()(*int64) {
    return m.exp
}
// GetFieldDeserializers the deserialization information for the current model
func (m *KeySetsItemUploadSecretPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["exp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExp(val)
        }
        return nil
    }
    res["k"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetK(val)
        }
        return nil
    }
    res["nbf"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNbf(val)
        }
        return nil
    }
    res["use"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUse(val)
        }
        return nil
    }
    return res
}
// GetK gets the k property value. The k property
func (m *KeySetsItemUploadSecretPostRequestBody) GetK()(*string) {
    return m.k
}
// GetNbf gets the nbf property value. The nbf property
func (m *KeySetsItemUploadSecretPostRequestBody) GetNbf()(*int64) {
    return m.nbf
}
// GetUse gets the use property value. The use property
func (m *KeySetsItemUploadSecretPostRequestBody) GetUse()(*string) {
    return m.use
}
// Serialize serializes information the current object
func (m *KeySetsItemUploadSecretPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("exp", m.GetExp())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("k", m.GetK())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("nbf", m.GetNbf())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("use", m.GetUse())
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
func (m *KeySetsItemUploadSecretPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetExp sets the exp property value. The exp property
func (m *KeySetsItemUploadSecretPostRequestBody) SetExp(value *int64)() {
    m.exp = value
}
// SetK sets the k property value. The k property
func (m *KeySetsItemUploadSecretPostRequestBody) SetK(value *string)() {
    m.k = value
}
// SetNbf sets the nbf property value. The nbf property
func (m *KeySetsItemUploadSecretPostRequestBody) SetNbf(value *int64)() {
    m.nbf = value
}
// SetUse sets the use property value. The use property
func (m *KeySetsItemUploadSecretPostRequestBody) SetUse(value *string)() {
    m.use = value
}
