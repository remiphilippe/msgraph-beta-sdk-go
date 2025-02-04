package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PlannerTaskDetailsable 
type PlannerTaskDetailsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    PlannerDeltaable
    GetChecklist()(PlannerChecklistItemsable)
    GetDescription()(*string)
    GetNotes()(ItemBodyable)
    GetPreviewType()(*PlannerPreviewType)
    GetReferences()(PlannerExternalReferencesable)
    SetChecklist(value PlannerChecklistItemsable)()
    SetDescription(value *string)()
    SetNotes(value ItemBodyable)()
    SetPreviewType(value *PlannerPreviewType)()
    SetReferences(value PlannerExternalReferencesable)()
}
