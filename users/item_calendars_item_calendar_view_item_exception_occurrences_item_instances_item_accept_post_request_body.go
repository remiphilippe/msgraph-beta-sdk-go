package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemAcceptPostRequestBody 
type ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemAcceptPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The Comment property
    comment *string
    // The SendResponse property
    sendResponse *bool
}
// NewItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemAcceptPostRequestBody instantiates a new ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemAcceptPostRequestBody and sets the default values.
func NewItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemAcceptPostRequestBody()(*ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemAcceptPostRequestBody) {
    m := &ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemAcceptPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any));
    return m
}
// CreateItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemAcceptPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemAcceptPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemAcceptPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemAcceptPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetComment gets the comment property value. The Comment property
func (m *ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemAcceptPostRequestBody) GetComment()(*string) {
    return m.comment
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemAcceptPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["comment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComment(val)
        }
        return nil
    }
    res["sendResponse"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSendResponse(val)
        }
        return nil
    }
    return res
}
// GetSendResponse gets the sendResponse property value. The SendResponse property
func (m *ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemAcceptPostRequestBody) GetSendResponse()(*bool) {
    return m.sendResponse
}
// Serialize serializes information the current object
func (m *ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemAcceptPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("comment", m.GetComment())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("sendResponse", m.GetSendResponse())
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
func (m *ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemAcceptPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetComment sets the comment property value. The Comment property
func (m *ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemAcceptPostRequestBody) SetComment(value *string)() {
    m.comment = value
}
// SetSendResponse sets the sendResponse property value. The SendResponse property
func (m *ItemCalendarsItemCalendarViewItemExceptionOccurrencesItemInstancesItemAcceptPostRequestBody) SetSendResponse(value *bool)() {
    m.sendResponse = value
}
