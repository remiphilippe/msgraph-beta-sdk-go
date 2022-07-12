package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// Alertable 
type Alertable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActorDisplayName()(*string)
    GetAlertWebUrl()(*string)
    GetAssignedTo()(*string)
    GetCategory()(*string)
    GetClassification()(*AlertClassification)
    GetComments()([]AlertCommentable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetDetectionSource()(*DetectionSource)
    GetDetectorId()(*string)
    GetDetermination()(*AlertDetermination)
    GetEvidence()([]AlertEvidenceable)
    GetFirstActivityDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetIncidentId()(*string)
    GetIncidentWebUrl()(*string)
    GetLastActivityDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastUpdateDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMitreTechniques()([]string)
    GetProviderAlertId()(*string)
    GetRecommendedActions()(*string)
    GetResolvedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetServiceSource()(*ServiceSource)
    GetSeverity()(*AlertSeverity)
    GetStatus()(*AlertStatus)
    GetTenantId()(*string)
    GetThreatDisplayName()(*string)
    GetThreatFamilyName()(*string)
    GetTitle()(*string)
    SetActorDisplayName(value *string)()
    SetAlertWebUrl(value *string)()
    SetAssignedTo(value *string)()
    SetCategory(value *string)()
    SetClassification(value *AlertClassification)()
    SetComments(value []AlertCommentable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetDetectionSource(value *DetectionSource)()
    SetDetectorId(value *string)()
    SetDetermination(value *AlertDetermination)()
    SetEvidence(value []AlertEvidenceable)()
    SetFirstActivityDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetIncidentId(value *string)()
    SetIncidentWebUrl(value *string)()
    SetLastActivityDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastUpdateDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMitreTechniques(value []string)()
    SetProviderAlertId(value *string)()
    SetRecommendedActions(value *string)()
    SetResolvedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetServiceSource(value *ServiceSource)()
    SetSeverity(value *AlertSeverity)()
    SetStatus(value *AlertStatus)()
    SetTenantId(value *string)()
    SetThreatDisplayName(value *string)()
    SetThreatFamilyName(value *string)()
    SetTitle(value *string)()
}
