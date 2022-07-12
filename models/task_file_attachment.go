package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TaskFileAttachment 
type TaskFileAttachment struct {
    AttachmentBase
    // The base64-encoded contents of the file.
    contentBytes []byte
}
// NewTaskFileAttachment instantiates a new TaskFileAttachment and sets the default values.
func NewTaskFileAttachment()(*TaskFileAttachment) {
    m := &TaskFileAttachment{
        AttachmentBase: *NewAttachmentBase(),
    }
    return m
}
// CreateTaskFileAttachmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTaskFileAttachmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTaskFileAttachment(), nil
}
// GetContentBytes gets the contentBytes property value. The base64-encoded contents of the file.
func (m *TaskFileAttachment) GetContentBytes()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.contentBytes
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TaskFileAttachment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AttachmentBase.GetFieldDeserializers()
    res["contentBytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentBytes(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *TaskFileAttachment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AttachmentBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteByteArrayValue("contentBytes", m.GetContentBytes())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContentBytes sets the contentBytes property value. The base64-encoded contents of the file.
func (m *TaskFileAttachment) SetContentBytes(value []byte)() {
    if m != nil {
        m.contentBytes = value
    }
}
