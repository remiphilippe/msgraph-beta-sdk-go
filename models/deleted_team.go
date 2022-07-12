package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeletedTeam 
type DeletedTeam struct {
    Entity
    // The channels those are either shared with this deleted team or created in this deleted team.
    channels []Channelable
}
// NewDeletedTeam instantiates a new DeletedTeam and sets the default values.
func NewDeletedTeam()(*DeletedTeam) {
    m := &DeletedTeam{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeletedTeamFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeletedTeamFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeletedTeam(), nil
}
// GetChannels gets the channels property value. The channels those are either shared with this deleted team or created in this deleted team.
func (m *DeletedTeam) GetChannels()([]Channelable) {
    if m == nil {
        return nil
    } else {
        return m.channels
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeletedTeam) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["channels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateChannelFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Channelable, len(val))
            for i, v := range val {
                res[i] = v.(Channelable)
            }
            m.SetChannels(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *DeletedTeam) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetChannels() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetChannels()))
        for i, v := range m.GetChannels() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("channels", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetChannels sets the channels property value. The channels those are either shared with this deleted team or created in this deleted team.
func (m *DeletedTeam) SetChannels(value []Channelable)() {
    if m != nil {
        m.channels = value
    }
}
