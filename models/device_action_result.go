package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceActionResult device action result
type DeviceActionResult struct {
    // Action name
    actionName *string
    // The actionState property
    actionState *ActionState
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Time the action state was last updated
    lastUpdatedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The OdataType property
    odataType *string
    // Time the action was initiated
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewDeviceActionResult instantiates a new deviceActionResult and sets the default values.
func NewDeviceActionResult()(*DeviceActionResult) {
    m := &DeviceActionResult{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceActionResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceActionResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.activateDeviceEsimActionResult":
                        return NewActivateDeviceEsimActionResult(), nil
                    case "#microsoft.graph.configurationManagerActionResult":
                        return NewConfigurationManagerActionResult(), nil
                    case "#microsoft.graph.deleteUserFromSharedAppleDeviceActionResult":
                        return NewDeleteUserFromSharedAppleDeviceActionResult(), nil
                    case "#microsoft.graph.locateDeviceActionResult":
                        return NewLocateDeviceActionResult(), nil
                    case "#microsoft.graph.remoteLockActionResult":
                        return NewRemoteLockActionResult(), nil
                    case "#microsoft.graph.resetPasscodeActionResult":
                        return NewResetPasscodeActionResult(), nil
                    case "#microsoft.graph.revokeAppleVppLicensesActionResult":
                        return NewRevokeAppleVppLicensesActionResult(), nil
                    case "#microsoft.graph.rotateBitLockerKeysDeviceActionResult":
                        return NewRotateBitLockerKeysDeviceActionResult(), nil
                    case "#microsoft.graph.windowsDefenderScanActionResult":
                        return NewWindowsDefenderScanActionResult(), nil
                }
            }
        }
    }
    return NewDeviceActionResult(), nil
}
// GetActionName gets the actionName property value. Action name
func (m *DeviceActionResult) GetActionName()(*string) {
    return m.actionName
}
// GetActionState gets the actionState property value. The actionState property
func (m *DeviceActionResult) GetActionState()(*ActionState) {
    return m.actionState
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceActionResult) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceActionResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["actionName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetActionName)
    res["actionState"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseActionState , m.SetActionState)
    res["lastUpdatedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastUpdatedDateTime)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["startDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetStartDateTime)
    return res
}
// GetLastUpdatedDateTime gets the lastUpdatedDateTime property value. Time the action state was last updated
func (m *DeviceActionResult) GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastUpdatedDateTime
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceActionResult) GetOdataType()(*string) {
    return m.odataType
}
// GetStartDateTime gets the startDateTime property value. Time the action was initiated
func (m *DeviceActionResult) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.startDateTime
}
// Serialize serializes information the current object
func (m *DeviceActionResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("actionName", m.GetActionName())
        if err != nil {
            return err
        }
    }
    if m.GetActionState() != nil {
        cast := (*m.GetActionState()).String()
        err := writer.WriteStringValue("actionState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("lastUpdatedDateTime", m.GetLastUpdatedDateTime())
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
        err := writer.WriteTimeValue("startDateTime", m.GetStartDateTime())
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
// SetActionName sets the actionName property value. Action name
func (m *DeviceActionResult) SetActionName(value *string)() {
    m.actionName = value
}
// SetActionState sets the actionState property value. The actionState property
func (m *DeviceActionResult) SetActionState(value *ActionState)() {
    m.actionState = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceActionResult) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetLastUpdatedDateTime sets the lastUpdatedDateTime property value. Time the action state was last updated
func (m *DeviceActionResult) SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastUpdatedDateTime = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceActionResult) SetOdataType(value *string)() {
    m.odataType = value
}
// SetStartDateTime sets the startDateTime property value. Time the action was initiated
func (m *DeviceActionResult) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startDateTime = value
}
