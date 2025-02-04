package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PowerliftIncidentMetadataable 
type PowerliftIncidentMetadataable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApplication()(*string)
    GetClientVersion()(*string)
    GetCreatedAtDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetEasyId()(*string)
    GetFileNames()([]string)
    GetLocale()(*string)
    GetOdataType()(*string)
    GetPlatform()(*string)
    GetPowerliftId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    SetApplication(value *string)()
    SetClientVersion(value *string)()
    SetCreatedAtDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetEasyId(value *string)()
    SetFileNames(value []string)()
    SetLocale(value *string)()
    SetOdataType(value *string)()
    SetPlatform(value *string)()
    SetPowerliftId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
}
