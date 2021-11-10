package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type PreAuthorizedApplication struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The unique identifier for the application.
    appId *string;
    // The unique identifier for the oauth2PermissionScopes the application requires.
    permissionIds []string;
}
// Instantiates a new preAuthorizedApplication and sets the default values.
func NewPreAuthorizedApplication()(*PreAuthorizedApplication) {
    m := &PreAuthorizedApplication{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PreAuthorizedApplication) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the appId property value. The unique identifier for the application.
func (m *PreAuthorizedApplication) GetAppId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appId
    }
}
// Gets the permissionIds property value. The unique identifier for the oauth2PermissionScopes the application requires.
func (m *PreAuthorizedApplication) GetPermissionIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.permissionIds
    }
}
// The deserialization information for the current model
func (m *PreAuthorizedApplication) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["appId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppId(val)
        }
        return nil
    }
    res["permissionIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetPermissionIds(res)
        }
        return nil
    }
    return res
}
func (m *PreAuthorizedApplication) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *PreAuthorizedApplication) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("appId", m.GetAppId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("permissionIds", m.GetPermissionIds())
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *PreAuthorizedApplication) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the appId property value. The unique identifier for the application.
// Parameters:
//  - value : Value to set for the appId property.
func (m *PreAuthorizedApplication) SetAppId(value *string)() {
    m.appId = value
}
// Sets the permissionIds property value. The unique identifier for the oauth2PermissionScopes the application requires.
// Parameters:
//  - value : Value to set for the permissionIds property.
func (m *PreAuthorizedApplication) SetPermissionIds(value []string)() {
    m.permissionIds = value
}
