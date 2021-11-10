package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type PlannerAssignedToTaskBoardTaskFormat struct {
    PlannerDelta
    // Dictionary of hints used to order tasks on the AssignedTo view of the Task Board. The key of each entry is one of the users the task is assigned to and the value is the order hint. The format of each value is defined as outlined here.
    orderHintsByAssignee *PlannerOrderHintsByAssignee;
    // Hint value used to order the task on the AssignedTo view of the Task Board when the task is not assigned to anyone, or if the orderHintsByAssignee dictionary does not provide an order hint for the user the task is assigned to. The format is defined as outlined here.
    unassignedOrderHint *string;
}
// Instantiates a new plannerAssignedToTaskBoardTaskFormat and sets the default values.
func NewPlannerAssignedToTaskBoardTaskFormat()(*PlannerAssignedToTaskBoardTaskFormat) {
    m := &PlannerAssignedToTaskBoardTaskFormat{
        PlannerDelta: *NewPlannerDelta(),
    }
    return m
}
// Gets the orderHintsByAssignee property value. Dictionary of hints used to order tasks on the AssignedTo view of the Task Board. The key of each entry is one of the users the task is assigned to and the value is the order hint. The format of each value is defined as outlined here.
func (m *PlannerAssignedToTaskBoardTaskFormat) GetOrderHintsByAssignee()(*PlannerOrderHintsByAssignee) {
    if m == nil {
        return nil
    } else {
        return m.orderHintsByAssignee
    }
}
// Gets the unassignedOrderHint property value. Hint value used to order the task on the AssignedTo view of the Task Board when the task is not assigned to anyone, or if the orderHintsByAssignee dictionary does not provide an order hint for the user the task is assigned to. The format is defined as outlined here.
func (m *PlannerAssignedToTaskBoardTaskFormat) GetUnassignedOrderHint()(*string) {
    if m == nil {
        return nil
    } else {
        return m.unassignedOrderHint
    }
}
// The deserialization information for the current model
func (m *PlannerAssignedToTaskBoardTaskFormat) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.PlannerDelta.GetFieldDeserializers()
    res["orderHintsByAssignee"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPlannerOrderHintsByAssignee() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrderHintsByAssignee(val.(*PlannerOrderHintsByAssignee))
        }
        return nil
    }
    res["unassignedOrderHint"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnassignedOrderHint(val)
        }
        return nil
    }
    return res
}
func (m *PlannerAssignedToTaskBoardTaskFormat) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *PlannerAssignedToTaskBoardTaskFormat) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.PlannerDelta.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("orderHintsByAssignee", m.GetOrderHintsByAssignee())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("unassignedOrderHint", m.GetUnassignedOrderHint())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the orderHintsByAssignee property value. Dictionary of hints used to order tasks on the AssignedTo view of the Task Board. The key of each entry is one of the users the task is assigned to and the value is the order hint. The format of each value is defined as outlined here.
// Parameters:
//  - value : Value to set for the orderHintsByAssignee property.
func (m *PlannerAssignedToTaskBoardTaskFormat) SetOrderHintsByAssignee(value *PlannerOrderHintsByAssignee)() {
    m.orderHintsByAssignee = value
}
// Sets the unassignedOrderHint property value. Hint value used to order the task on the AssignedTo view of the Task Board when the task is not assigned to anyone, or if the orderHintsByAssignee dictionary does not provide an order hint for the user the task is assigned to. The format is defined as outlined here.
// Parameters:
//  - value : Value to set for the unassignedOrderHint property.
func (m *PlannerAssignedToTaskBoardTaskFormat) SetUnassignedOrderHint(value *string)() {
    m.unassignedOrderHint = value
}
