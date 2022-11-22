package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MobileAppTroubleshootingEvent 
type MobileAppTroubleshootingEvent struct {
    DeviceManagementTroubleshootingEvent
    // Intune application identifier.
    applicationId *string
    // The collection property of AppLogUploadRequest.
    appLogCollectionRequests []AppLogCollectionRequestable
    // Intune Mobile Application Troubleshooting History Item
    history []MobileAppTroubleshootingHistoryItemable
    // Device identifier created or collected by Intune.
    managedDeviceIdentifier *string
    // Identifier for the user that tried to enroll the device.
    userId *string
}
// NewMobileAppTroubleshootingEvent instantiates a new MobileAppTroubleshootingEvent and sets the default values.
func NewMobileAppTroubleshootingEvent()(*MobileAppTroubleshootingEvent) {
    m := &MobileAppTroubleshootingEvent{
        DeviceManagementTroubleshootingEvent: *NewDeviceManagementTroubleshootingEvent(),
    }
    return m
}
// CreateMobileAppTroubleshootingEventFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMobileAppTroubleshootingEventFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMobileAppTroubleshootingEvent(), nil
}
// GetApplicationId gets the applicationId property value. Intune application identifier.
func (m *MobileAppTroubleshootingEvent) GetApplicationId()(*string) {
    return m.applicationId
}
// GetAppLogCollectionRequests gets the appLogCollectionRequests property value. The collection property of AppLogUploadRequest.
func (m *MobileAppTroubleshootingEvent) GetAppLogCollectionRequests()([]AppLogCollectionRequestable) {
    return m.appLogCollectionRequests
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MobileAppTroubleshootingEvent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementTroubleshootingEvent.GetFieldDeserializers()
    res["applicationId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetApplicationId)
    res["appLogCollectionRequests"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAppLogCollectionRequestFromDiscriminatorValue , m.SetAppLogCollectionRequests)
    res["history"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateMobileAppTroubleshootingHistoryItemFromDiscriminatorValue , m.SetHistory)
    res["managedDeviceIdentifier"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetManagedDeviceIdentifier)
    res["userId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUserId)
    return res
}
// GetHistory gets the history property value. Intune Mobile Application Troubleshooting History Item
func (m *MobileAppTroubleshootingEvent) GetHistory()([]MobileAppTroubleshootingHistoryItemable) {
    return m.history
}
// GetManagedDeviceIdentifier gets the managedDeviceIdentifier property value. Device identifier created or collected by Intune.
func (m *MobileAppTroubleshootingEvent) GetManagedDeviceIdentifier()(*string) {
    return m.managedDeviceIdentifier
}
// GetUserId gets the userId property value. Identifier for the user that tried to enroll the device.
func (m *MobileAppTroubleshootingEvent) GetUserId()(*string) {
    return m.userId
}
// Serialize serializes information the current object
func (m *MobileAppTroubleshootingEvent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementTroubleshootingEvent.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("applicationId", m.GetApplicationId())
        if err != nil {
            return err
        }
    }
    if m.GetAppLogCollectionRequests() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAppLogCollectionRequests())
        err = writer.WriteCollectionOfObjectValues("appLogCollectionRequests", cast)
        if err != nil {
            return err
        }
    }
    if m.GetHistory() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetHistory())
        err = writer.WriteCollectionOfObjectValues("history", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("managedDeviceIdentifier", m.GetManagedDeviceIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApplicationId sets the applicationId property value. Intune application identifier.
func (m *MobileAppTroubleshootingEvent) SetApplicationId(value *string)() {
    m.applicationId = value
}
// SetAppLogCollectionRequests sets the appLogCollectionRequests property value. The collection property of AppLogUploadRequest.
func (m *MobileAppTroubleshootingEvent) SetAppLogCollectionRequests(value []AppLogCollectionRequestable)() {
    m.appLogCollectionRequests = value
}
// SetHistory sets the history property value. Intune Mobile Application Troubleshooting History Item
func (m *MobileAppTroubleshootingEvent) SetHistory(value []MobileAppTroubleshootingHistoryItemable)() {
    m.history = value
}
// SetManagedDeviceIdentifier sets the managedDeviceIdentifier property value. Device identifier created or collected by Intune.
func (m *MobileAppTroubleshootingEvent) SetManagedDeviceIdentifier(value *string)() {
    m.managedDeviceIdentifier = value
}
// SetUserId sets the userId property value. Identifier for the user that tried to enroll the device.
func (m *MobileAppTroubleshootingEvent) SetUserId(value *string)() {
    m.userId = value
}
