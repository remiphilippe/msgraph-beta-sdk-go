package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RemoteAssistancePartner 
type RemoteAssistancePartner struct {
    Entity
    // Display name of the partner.
    displayName *string
    // Timestamp of the last request sent to Intune by the TEM partner.
    lastConnectionDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // When the OnboardingStatus is Onboarding, This is the date time when the onboarding request expires.
    onboardingRequestExpiryDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The current TeamViewer connector status
    onboardingStatus *RemoteAssistanceOnboardingStatus
    // URL of the partner's onboarding portal, where an administrator can configure their Remote Assistance service.
    onboardingUrl *string
}
// NewRemoteAssistancePartner instantiates a new RemoteAssistancePartner and sets the default values.
func NewRemoteAssistancePartner()(*RemoteAssistancePartner) {
    m := &RemoteAssistancePartner{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.remoteAssistancePartner";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateRemoteAssistancePartnerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRemoteAssistancePartnerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRemoteAssistancePartner(), nil
}
// GetDisplayName gets the displayName property value. Display name of the partner.
func (m *RemoteAssistancePartner) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RemoteAssistancePartner) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["lastConnectionDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastConnectionDateTime(val)
        }
        return nil
    }
    res["onboardingRequestExpiryDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnboardingRequestExpiryDateTime(val)
        }
        return nil
    }
    res["onboardingStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRemoteAssistanceOnboardingStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnboardingStatus(val.(*RemoteAssistanceOnboardingStatus))
        }
        return nil
    }
    res["onboardingUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnboardingUrl(val)
        }
        return nil
    }
    return res
}
// GetLastConnectionDateTime gets the lastConnectionDateTime property value. Timestamp of the last request sent to Intune by the TEM partner.
func (m *RemoteAssistancePartner) GetLastConnectionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastConnectionDateTime
    }
}
// GetOnboardingRequestExpiryDateTime gets the onboardingRequestExpiryDateTime property value. When the OnboardingStatus is Onboarding, This is the date time when the onboarding request expires.
func (m *RemoteAssistancePartner) GetOnboardingRequestExpiryDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.onboardingRequestExpiryDateTime
    }
}
// GetOnboardingStatus gets the onboardingStatus property value. The current TeamViewer connector status
func (m *RemoteAssistancePartner) GetOnboardingStatus()(*RemoteAssistanceOnboardingStatus) {
    if m == nil {
        return nil
    } else {
        return m.onboardingStatus
    }
}
// GetOnboardingUrl gets the onboardingUrl property value. URL of the partner's onboarding portal, where an administrator can configure their Remote Assistance service.
func (m *RemoteAssistancePartner) GetOnboardingUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.onboardingUrl
    }
}
// Serialize serializes information the current object
func (m *RemoteAssistancePartner) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastConnectionDateTime", m.GetLastConnectionDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("onboardingRequestExpiryDateTime", m.GetOnboardingRequestExpiryDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetOnboardingStatus() != nil {
        cast := (*m.GetOnboardingStatus()).String()
        err = writer.WriteStringValue("onboardingStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("onboardingUrl", m.GetOnboardingUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. Display name of the partner.
func (m *RemoteAssistancePartner) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetLastConnectionDateTime sets the lastConnectionDateTime property value. Timestamp of the last request sent to Intune by the TEM partner.
func (m *RemoteAssistancePartner) SetLastConnectionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastConnectionDateTime = value
    }
}
// SetOnboardingRequestExpiryDateTime sets the onboardingRequestExpiryDateTime property value. When the OnboardingStatus is Onboarding, This is the date time when the onboarding request expires.
func (m *RemoteAssistancePartner) SetOnboardingRequestExpiryDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.onboardingRequestExpiryDateTime = value
    }
}
// SetOnboardingStatus sets the onboardingStatus property value. The current TeamViewer connector status
func (m *RemoteAssistancePartner) SetOnboardingStatus(value *RemoteAssistanceOnboardingStatus)() {
    if m != nil {
        m.onboardingStatus = value
    }
}
// SetOnboardingUrl sets the onboardingUrl property value. URL of the partner's onboarding portal, where an administrator can configure their Remote Assistance service.
func (m *RemoteAssistancePartner) SetOnboardingUrl(value *string)() {
    if m != nil {
        m.onboardingUrl = value
    }
}
