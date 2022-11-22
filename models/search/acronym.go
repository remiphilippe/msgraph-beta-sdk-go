package search

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Acronym 
type Acronym struct {
    SearchAnswer
    // What the acronym stands for.
    standsFor *string
    // The state property
    state *AnswerState
}
// NewAcronym instantiates a new Acronym and sets the default values.
func NewAcronym()(*Acronym) {
    m := &Acronym{
        SearchAnswer: *NewSearchAnswer(),
    }
    return m
}
// CreateAcronymFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAcronymFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAcronym(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Acronym) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.SearchAnswer.GetFieldDeserializers()
    res["standsFor"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetStandsFor)
    res["state"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseAnswerState , m.SetState)
    return res
}
// GetStandsFor gets the standsFor property value. What the acronym stands for.
func (m *Acronym) GetStandsFor()(*string) {
    return m.standsFor
}
// GetState gets the state property value. The state property
func (m *Acronym) GetState()(*AnswerState) {
    return m.state
}
// Serialize serializes information the current object
func (m *Acronym) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.SearchAnswer.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("standsFor", m.GetStandsFor())
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err = writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetStandsFor sets the standsFor property value. What the acronym stands for.
func (m *Acronym) SetStandsFor(value *string)() {
    m.standsFor = value
}
// SetState sets the state property value. The state property
func (m *Acronym) SetState(value *AnswerState)() {
    m.state = value
}
