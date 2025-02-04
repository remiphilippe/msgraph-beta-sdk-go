package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkflowVersion 
type WorkflowVersion struct {
    WorkflowBase
    // The version of the workflow.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
    versionNumber *int32
}
// NewWorkflowVersion instantiates a new workflowVersion and sets the default values.
func NewWorkflowVersion()(*WorkflowVersion) {
    m := &WorkflowVersion{
        WorkflowBase: *NewWorkflowBase(),
    }
    odataTypeValue := "#microsoft.graph.identityGovernance.workflowVersion";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWorkflowVersionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkflowVersionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkflowVersion(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkflowVersion) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.WorkflowBase.GetFieldDeserializers()
    res["versionNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersionNumber(val)
        }
        return nil
    }
    return res
}
// GetVersionNumber gets the versionNumber property value. The version of the workflow.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
func (m *WorkflowVersion) GetVersionNumber()(*int32) {
    return m.versionNumber
}
// Serialize serializes information the current object
func (m *WorkflowVersion) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.WorkflowBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("versionNumber", m.GetVersionNumber())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetVersionNumber sets the versionNumber property value. The version of the workflow.Supports $filter(lt, le, gt, ge, eq, ne) and $orderby.
func (m *WorkflowVersion) SetVersionNumber(value *int32)() {
    m.versionNumber = value
}
