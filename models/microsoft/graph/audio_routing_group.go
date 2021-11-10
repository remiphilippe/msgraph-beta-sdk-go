package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AudioRoutingGroup struct {
    Entity
    // List of receiving participant ids.
    receivers []string;
    // Routing group mode.  Possible values are: oneToOne, multicast.
    routingMode *RoutingMode;
    // List of source participant ids.
    sources []string;
}
// Instantiates a new audioRoutingGroup and sets the default values.
func NewAudioRoutingGroup()(*AudioRoutingGroup) {
    m := &AudioRoutingGroup{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the receivers property value. List of receiving participant ids.
func (m *AudioRoutingGroup) GetReceivers()([]string) {
    if m == nil {
        return nil
    } else {
        return m.receivers
    }
}
// Gets the routingMode property value. Routing group mode.  Possible values are: oneToOne, multicast.
func (m *AudioRoutingGroup) GetRoutingMode()(*RoutingMode) {
    if m == nil {
        return nil
    } else {
        return m.routingMode
    }
}
// Gets the sources property value. List of source participant ids.
func (m *AudioRoutingGroup) GetSources()([]string) {
    if m == nil {
        return nil
    } else {
        return m.sources
    }
}
// The deserialization information for the current model
func (m *AudioRoutingGroup) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["receivers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetReceivers(res)
        }
        return nil
    }
    res["routingMode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRoutingMode)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(RoutingMode)
            m.SetRoutingMode(&cast)
        }
        return nil
    }
    res["sources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSources(res)
        }
        return nil
    }
    return res
}
func (m *AudioRoutingGroup) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AudioRoutingGroup) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteCollectionOfStringValues("receivers", m.GetReceivers())
        if err != nil {
            return err
        }
    }
    if m.GetRoutingMode() != nil {
        cast := m.GetRoutingMode().String()
        err = writer.WriteStringValue("routingMode", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("sources", m.GetSources())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the receivers property value. List of receiving participant ids.
// Parameters:
//  - value : Value to set for the receivers property.
func (m *AudioRoutingGroup) SetReceivers(value []string)() {
    m.receivers = value
}
// Sets the routingMode property value. Routing group mode.  Possible values are: oneToOne, multicast.
// Parameters:
//  - value : Value to set for the routingMode property.
func (m *AudioRoutingGroup) SetRoutingMode(value *RoutingMode)() {
    m.routingMode = value
}
// Sets the sources property value. List of source participant ids.
// Parameters:
//  - value : Value to set for the sources property.
func (m *AudioRoutingGroup) SetSources(value []string)() {
    m.sources = value
}
