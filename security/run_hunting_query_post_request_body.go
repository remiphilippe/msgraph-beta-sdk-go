package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RunHuntingQueryPostRequestBody 
type RunHuntingQueryPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The query property
    query *string
}
// NewRunHuntingQueryPostRequestBody instantiates a new RunHuntingQueryPostRequestBody and sets the default values.
func NewRunHuntingQueryPostRequestBody()(*RunHuntingQueryPostRequestBody) {
    m := &RunHuntingQueryPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any));
    return m
}
// CreateRunHuntingQueryPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRunHuntingQueryPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRunHuntingQueryPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RunHuntingQueryPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RunHuntingQueryPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["query"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQuery(val)
        }
        return nil
    }
    return res
}
// GetQuery gets the query property value. The query property
func (m *RunHuntingQueryPostRequestBody) GetQuery()(*string) {
    return m.query
}
// Serialize serializes information the current object
func (m *RunHuntingQueryPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("query", m.GetQuery())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RunHuntingQueryPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetQuery sets the query property value. The query property
func (m *RunHuntingQueryPostRequestBody) SetQuery(value *string)() {
    m.query = value
}
