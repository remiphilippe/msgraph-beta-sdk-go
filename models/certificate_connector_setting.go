package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CertificateConnectorSetting certificate connector settings.
type CertificateConnectorSetting struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Certificate expire time
    certExpiryTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Version of certificate connector
    connectorVersion *string
    // Certificate connector enrollment error
    enrollmentError *string
    // Last time certificate connector connected
    lastConnectorConnectionTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Version of last uploaded certificate connector
    lastUploadVersion *int64
    // The OdataType property
    odataType *string
    // Certificate connector status
    status *int32
}
// NewCertificateConnectorSetting instantiates a new certificateConnectorSetting and sets the default values.
func NewCertificateConnectorSetting()(*CertificateConnectorSetting) {
    m := &CertificateConnectorSetting{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCertificateConnectorSettingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCertificateConnectorSettingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCertificateConnectorSetting(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CertificateConnectorSetting) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCertExpiryTime gets the certExpiryTime property value. Certificate expire time
func (m *CertificateConnectorSetting) GetCertExpiryTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.certExpiryTime
}
// GetConnectorVersion gets the connectorVersion property value. Version of certificate connector
func (m *CertificateConnectorSetting) GetConnectorVersion()(*string) {
    return m.connectorVersion
}
// GetEnrollmentError gets the enrollmentError property value. Certificate connector enrollment error
func (m *CertificateConnectorSetting) GetEnrollmentError()(*string) {
    return m.enrollmentError
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CertificateConnectorSetting) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["certExpiryTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetCertExpiryTime)
    res["connectorVersion"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetConnectorVersion)
    res["enrollmentError"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetEnrollmentError)
    res["lastConnectorConnectionTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastConnectorConnectionTime)
    res["lastUploadVersion"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt64Value(m.SetLastUploadVersion)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["status"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetStatus)
    return res
}
// GetLastConnectorConnectionTime gets the lastConnectorConnectionTime property value. Last time certificate connector connected
func (m *CertificateConnectorSetting) GetLastConnectorConnectionTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastConnectorConnectionTime
}
// GetLastUploadVersion gets the lastUploadVersion property value. Version of last uploaded certificate connector
func (m *CertificateConnectorSetting) GetLastUploadVersion()(*int64) {
    return m.lastUploadVersion
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *CertificateConnectorSetting) GetOdataType()(*string) {
    return m.odataType
}
// GetStatus gets the status property value. Certificate connector status
func (m *CertificateConnectorSetting) GetStatus()(*int32) {
    return m.status
}
// Serialize serializes information the current object
func (m *CertificateConnectorSetting) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("certExpiryTime", m.GetCertExpiryTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("connectorVersion", m.GetConnectorVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("enrollmentError", m.GetEnrollmentError())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("lastConnectorConnectionTime", m.GetLastConnectorConnectionTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("lastUploadVersion", m.GetLastUploadVersion())
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
        err := writer.WriteInt32Value("status", m.GetStatus())
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
func (m *CertificateConnectorSetting) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCertExpiryTime sets the certExpiryTime property value. Certificate expire time
func (m *CertificateConnectorSetting) SetCertExpiryTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.certExpiryTime = value
}
// SetConnectorVersion sets the connectorVersion property value. Version of certificate connector
func (m *CertificateConnectorSetting) SetConnectorVersion(value *string)() {
    m.connectorVersion = value
}
// SetEnrollmentError sets the enrollmentError property value. Certificate connector enrollment error
func (m *CertificateConnectorSetting) SetEnrollmentError(value *string)() {
    m.enrollmentError = value
}
// SetLastConnectorConnectionTime sets the lastConnectorConnectionTime property value. Last time certificate connector connected
func (m *CertificateConnectorSetting) SetLastConnectorConnectionTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastConnectorConnectionTime = value
}
// SetLastUploadVersion sets the lastUploadVersion property value. Version of last uploaded certificate connector
func (m *CertificateConnectorSetting) SetLastUploadVersion(value *int64)() {
    m.lastUploadVersion = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CertificateConnectorSetting) SetOdataType(value *string)() {
    m.odataType = value
}
// SetStatus sets the status property value. Certificate connector status
func (m *CertificateConnectorSetting) SetStatus(value *int32)() {
    m.status = value
}
