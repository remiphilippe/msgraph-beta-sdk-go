package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PlannerTaskDetails 
type PlannerTaskDetails struct {
    PlannerDelta
    // The collection of checklist items on the task.
    checklist PlannerChecklistItemsable
    // Description of the task.
    description *string
    // Rich text description of the task. To be used by HTML-aware clients. For backwards compatibility, a plain-text version of the HTML description will be synced to the 'description' field. If this field has not previously been set but 'description' has been, the existing description will be synchronized to 'notes' with minimal whitespace-preserving HTML markup. Setting both 'description' and 'notes' is an error and will result in an exception.
    notes ItemBodyable
    // This sets the type of preview that shows up on the task. Possible values are: automatic, noPreview, checklist, description, reference. When set to automatic the displayed preview is chosen by the app viewing the task.
    previewType *PlannerPreviewType
    // The collection of references on the task.
    references PlannerExternalReferencesable
}
// NewPlannerTaskDetails instantiates a new plannerTaskDetails and sets the default values.
func NewPlannerTaskDetails()(*PlannerTaskDetails) {
    m := &PlannerTaskDetails{
        PlannerDelta: *NewPlannerDelta(),
    }
    return m
}
// CreatePlannerTaskDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlannerTaskDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPlannerTaskDetails(), nil
}
// GetChecklist gets the checklist property value. The collection of checklist items on the task.
func (m *PlannerTaskDetails) GetChecklist()(PlannerChecklistItemsable) {
    return m.checklist
}
// GetDescription gets the description property value. Description of the task.
func (m *PlannerTaskDetails) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlannerTaskDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PlannerDelta.GetFieldDeserializers()
    res["checklist"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePlannerChecklistItemsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChecklist(val.(PlannerChecklistItemsable))
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["notes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemBodyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotes(val.(ItemBodyable))
        }
        return nil
    }
    res["previewType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePlannerPreviewType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreviewType(val.(*PlannerPreviewType))
        }
        return nil
    }
    res["references"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePlannerExternalReferencesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReferences(val.(PlannerExternalReferencesable))
        }
        return nil
    }
    return res
}
// GetNotes gets the notes property value. Rich text description of the task. To be used by HTML-aware clients. For backwards compatibility, a plain-text version of the HTML description will be synced to the 'description' field. If this field has not previously been set but 'description' has been, the existing description will be synchronized to 'notes' with minimal whitespace-preserving HTML markup. Setting both 'description' and 'notes' is an error and will result in an exception.
func (m *PlannerTaskDetails) GetNotes()(ItemBodyable) {
    return m.notes
}
// GetPreviewType gets the previewType property value. This sets the type of preview that shows up on the task. Possible values are: automatic, noPreview, checklist, description, reference. When set to automatic the displayed preview is chosen by the app viewing the task.
func (m *PlannerTaskDetails) GetPreviewType()(*PlannerPreviewType) {
    return m.previewType
}
// GetReferences gets the references property value. The collection of references on the task.
func (m *PlannerTaskDetails) GetReferences()(PlannerExternalReferencesable) {
    return m.references
}
// Serialize serializes information the current object
func (m *PlannerTaskDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PlannerDelta.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("checklist", m.GetChecklist())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("notes", m.GetNotes())
        if err != nil {
            return err
        }
    }
    if m.GetPreviewType() != nil {
        cast := (*m.GetPreviewType()).String()
        err = writer.WriteStringValue("previewType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("references", m.GetReferences())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetChecklist sets the checklist property value. The collection of checklist items on the task.
func (m *PlannerTaskDetails) SetChecklist(value PlannerChecklistItemsable)() {
    m.checklist = value
}
// SetDescription sets the description property value. Description of the task.
func (m *PlannerTaskDetails) SetDescription(value *string)() {
    m.description = value
}
// SetNotes sets the notes property value. Rich text description of the task. To be used by HTML-aware clients. For backwards compatibility, a plain-text version of the HTML description will be synced to the 'description' field. If this field has not previously been set but 'description' has been, the existing description will be synchronized to 'notes' with minimal whitespace-preserving HTML markup. Setting both 'description' and 'notes' is an error and will result in an exception.
func (m *PlannerTaskDetails) SetNotes(value ItemBodyable)() {
    m.notes = value
}
// SetPreviewType sets the previewType property value. This sets the type of preview that shows up on the task. Possible values are: automatic, noPreview, checklist, description, reference. When set to automatic the displayed preview is chosen by the app viewing the task.
func (m *PlannerTaskDetails) SetPreviewType(value *PlannerPreviewType)() {
    m.previewType = value
}
// SetReferences sets the references property value. The collection of references on the task.
func (m *PlannerTaskDetails) SetReferences(value PlannerExternalReferencesable)() {
    m.references = value
}
