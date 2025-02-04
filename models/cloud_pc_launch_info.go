package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcLaunchInfo 
type CloudPcLaunchInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The unique identifier of the Cloud PC.
    cloudPcId *string
    // The connect URL of the Cloud PC.
    cloudPcLaunchUrl *string
    // The OdataType property
    odataType *string
}
// NewCloudPcLaunchInfo instantiates a new cloudPcLaunchInfo and sets the default values.
func NewCloudPcLaunchInfo()(*CloudPcLaunchInfo) {
    m := &CloudPcLaunchInfo{
    }
    m.SetAdditionalData(make(map[string]any));
    return m
}
// CreateCloudPcLaunchInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcLaunchInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcLaunchInfo(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CloudPcLaunchInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCloudPcId gets the cloudPcId property value. The unique identifier of the Cloud PC.
func (m *CloudPcLaunchInfo) GetCloudPcId()(*string) {
    return m.cloudPcId
}
// GetCloudPcLaunchUrl gets the cloudPcLaunchUrl property value. The connect URL of the Cloud PC.
func (m *CloudPcLaunchInfo) GetCloudPcLaunchUrl()(*string) {
    return m.cloudPcLaunchUrl
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcLaunchInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["cloudPcId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloudPcId(val)
        }
        return nil
    }
    res["cloudPcLaunchUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloudPcLaunchUrl(val)
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
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *CloudPcLaunchInfo) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *CloudPcLaunchInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("cloudPcId", m.GetCloudPcId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("cloudPcLaunchUrl", m.GetCloudPcLaunchUrl())
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
func (m *CloudPcLaunchInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCloudPcId sets the cloudPcId property value. The unique identifier of the Cloud PC.
func (m *CloudPcLaunchInfo) SetCloudPcId(value *string)() {
    m.cloudPcId = value
}
// SetCloudPcLaunchUrl sets the cloudPcLaunchUrl property value. The connect URL of the Cloud PC.
func (m *CloudPcLaunchInfo) SetCloudPcLaunchUrl(value *string)() {
    m.cloudPcLaunchUrl = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CloudPcLaunchInfo) SetOdataType(value *string)() {
    m.odataType = value
}
