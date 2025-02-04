package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EntitlementManagementAccessPackageAssignmentsItemAccessPackageMoveToCatalogPostRequestBody 
type EntitlementManagementAccessPackageAssignmentsItemAccessPackageMoveToCatalogPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The catalogId property
    catalogId *string
}
// NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageMoveToCatalogPostRequestBody instantiates a new EntitlementManagementAccessPackageAssignmentsItemAccessPackageMoveToCatalogPostRequestBody and sets the default values.
func NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageMoveToCatalogPostRequestBody()(*EntitlementManagementAccessPackageAssignmentsItemAccessPackageMoveToCatalogPostRequestBody) {
    m := &EntitlementManagementAccessPackageAssignmentsItemAccessPackageMoveToCatalogPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any));
    return m
}
// CreateEntitlementManagementAccessPackageAssignmentsItemAccessPackageMoveToCatalogPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEntitlementManagementAccessPackageAssignmentsItemAccessPackageMoveToCatalogPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEntitlementManagementAccessPackageAssignmentsItemAccessPackageMoveToCatalogPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageMoveToCatalogPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCatalogId gets the catalogId property value. The catalogId property
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageMoveToCatalogPostRequestBody) GetCatalogId()(*string) {
    return m.catalogId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageMoveToCatalogPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["catalogId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCatalogId(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageMoveToCatalogPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("catalogId", m.GetCatalogId())
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
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageMoveToCatalogPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCatalogId sets the catalogId property value. The catalogId property
func (m *EntitlementManagementAccessPackageAssignmentsItemAccessPackageMoveToCatalogPostRequestBody) SetCatalogId(value *string)() {
    m.catalogId = value
}
