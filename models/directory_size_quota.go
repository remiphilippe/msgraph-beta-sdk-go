package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DirectorySizeQuota 
type DirectorySizeQuota struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The OdataType property
    odataType *string
    // Total amount of the directory quota.
    total *int32
    // Used amount of the directory quota.
    used *int32
}
// NewDirectorySizeQuota instantiates a new directorySizeQuota and sets the default values.
func NewDirectorySizeQuota()(*DirectorySizeQuota) {
    m := &DirectorySizeQuota{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDirectorySizeQuotaFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDirectorySizeQuotaFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDirectorySizeQuota(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DirectorySizeQuota) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DirectorySizeQuota) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["total"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetTotal)
    res["used"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetUsed)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DirectorySizeQuota) GetOdataType()(*string) {
    return m.odataType
}
// GetTotal gets the total property value. Total amount of the directory quota.
func (m *DirectorySizeQuota) GetTotal()(*int32) {
    return m.total
}
// GetUsed gets the used property value. Used amount of the directory quota.
func (m *DirectorySizeQuota) GetUsed()(*int32) {
    return m.used
}
// Serialize serializes information the current object
func (m *DirectorySizeQuota) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total", m.GetTotal())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("used", m.GetUsed())
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
func (m *DirectorySizeQuota) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DirectorySizeQuota) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTotal sets the total property value. Total amount of the directory quota.
func (m *DirectorySizeQuota) SetTotal(value *int32)() {
    m.total = value
}
// SetUsed sets the used property value. Used amount of the directory quota.
func (m *DirectorySizeQuota) SetUsed(value *int32)() {
    m.used = value
}
