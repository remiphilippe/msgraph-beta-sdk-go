package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Teamworkable 
type Teamworkable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeletedTeams()([]DeletedTeamable)
    GetDevices()([]TeamworkDeviceable)
    GetTeamsAppSettings()(TeamsAppSettingsable)
    GetTeamTemplates()([]TeamTemplateable)
    GetWorkforceIntegrations()([]WorkforceIntegrationable)
    SetDeletedTeams(value []DeletedTeamable)()
    SetDevices(value []TeamworkDeviceable)()
    SetTeamsAppSettings(value TeamsAppSettingsable)()
    SetTeamTemplates(value []TeamTemplateable)()
    SetWorkforceIntegrations(value []WorkforceIntegrationable)()
}
