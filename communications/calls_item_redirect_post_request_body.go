package communications

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// CallsItemRedirectPostRequestBody 
type CallsItemRedirectPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The callbackUri property
    callbackUri *string
    // The maskCallee property
    maskCallee *bool
    // The maskCaller property
    maskCaller *bool
    // The targetDisposition property
    targetDisposition *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CallDisposition
    // The targets property
    targets []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InvitationParticipantInfoable
    // The timeout property
    timeout *int32
}
// NewCallsItemRedirectPostRequestBody instantiates a new CallsItemRedirectPostRequestBody and sets the default values.
func NewCallsItemRedirectPostRequestBody()(*CallsItemRedirectPostRequestBody) {
    m := &CallsItemRedirectPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any));
    return m
}
// CreateCallsItemRedirectPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCallsItemRedirectPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCallsItemRedirectPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CallsItemRedirectPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCallbackUri gets the callbackUri property value. The callbackUri property
func (m *CallsItemRedirectPostRequestBody) GetCallbackUri()(*string) {
    return m.callbackUri
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CallsItemRedirectPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["callbackUri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCallbackUri(val)
        }
        return nil
    }
    res["maskCallee"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaskCallee(val)
        }
        return nil
    }
    res["maskCaller"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaskCaller(val)
        }
        return nil
    }
    res["targetDisposition"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ParseCallDisposition)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetDisposition(val.(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CallDisposition))
        }
        return nil
    }
    res["targets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateInvitationParticipantInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InvitationParticipantInfoable, len(val))
            for i, v := range val {
                res[i] = v.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InvitationParticipantInfoable)
            }
            m.SetTargets(res)
        }
        return nil
    }
    res["timeout"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeout(val)
        }
        return nil
    }
    return res
}
// GetMaskCallee gets the maskCallee property value. The maskCallee property
func (m *CallsItemRedirectPostRequestBody) GetMaskCallee()(*bool) {
    return m.maskCallee
}
// GetMaskCaller gets the maskCaller property value. The maskCaller property
func (m *CallsItemRedirectPostRequestBody) GetMaskCaller()(*bool) {
    return m.maskCaller
}
// GetTargetDisposition gets the targetDisposition property value. The targetDisposition property
func (m *CallsItemRedirectPostRequestBody) GetTargetDisposition()(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CallDisposition) {
    return m.targetDisposition
}
// GetTargets gets the targets property value. The targets property
func (m *CallsItemRedirectPostRequestBody) GetTargets()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InvitationParticipantInfoable) {
    return m.targets
}
// GetTimeout gets the timeout property value. The timeout property
func (m *CallsItemRedirectPostRequestBody) GetTimeout()(*int32) {
    return m.timeout
}
// Serialize serializes information the current object
func (m *CallsItemRedirectPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("callbackUri", m.GetCallbackUri())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("maskCallee", m.GetMaskCallee())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("maskCaller", m.GetMaskCaller())
        if err != nil {
            return err
        }
    }
    if m.GetTargetDisposition() != nil {
        cast := (*m.GetTargetDisposition()).String()
        err := writer.WriteStringValue("targetDisposition", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetTargets() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTargets()))
        for i, v := range m.GetTargets() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("targets", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("timeout", m.GetTimeout())
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
func (m *CallsItemRedirectPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCallbackUri sets the callbackUri property value. The callbackUri property
func (m *CallsItemRedirectPostRequestBody) SetCallbackUri(value *string)() {
    m.callbackUri = value
}
// SetMaskCallee sets the maskCallee property value. The maskCallee property
func (m *CallsItemRedirectPostRequestBody) SetMaskCallee(value *bool)() {
    m.maskCallee = value
}
// SetMaskCaller sets the maskCaller property value. The maskCaller property
func (m *CallsItemRedirectPostRequestBody) SetMaskCaller(value *bool)() {
    m.maskCaller = value
}
// SetTargetDisposition sets the targetDisposition property value. The targetDisposition property
func (m *CallsItemRedirectPostRequestBody) SetTargetDisposition(value *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CallDisposition)() {
    m.targetDisposition = value
}
// SetTargets sets the targets property value. The targets property
func (m *CallsItemRedirectPostRequestBody) SetTargets(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InvitationParticipantInfoable)() {
    m.targets = value
}
// SetTimeout sets the timeout property value. The timeout property
func (m *CallsItemRedirectPostRequestBody) SetTimeout(value *int32)() {
    m.timeout = value
}
