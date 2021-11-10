package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type InformationProtectionPolicy struct {
    Entity
    // 
    labels []InformationProtectionLabel;
}
// Instantiates a new informationProtectionPolicy and sets the default values.
func NewInformationProtectionPolicy()(*InformationProtectionPolicy) {
    m := &InformationProtectionPolicy{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the labels property value. 
func (m *InformationProtectionPolicy) GetLabels()([]InformationProtectionLabel) {
    if m == nil {
        return nil
    } else {
        return m.labels
    }
}
// The deserialization information for the current model
func (m *InformationProtectionPolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["labels"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewInformationProtectionLabel() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]InformationProtectionLabel, len(val))
            for i, v := range val {
                res[i] = *(v.(*InformationProtectionLabel))
            }
            m.SetLabels(res)
        }
        return nil
    }
    return res
}
func (m *InformationProtectionPolicy) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *InformationProtectionPolicy) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetLabels()))
        for i, v := range m.GetLabels() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("labels", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the labels property value. 
// Parameters:
//  - value : Value to set for the labels property.
func (m *InformationProtectionPolicy) SetLabels(value []InformationProtectionLabel)() {
    m.labels = value
}
