package users

import (
	i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
	ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ItemAssignLicensePostRequestBody
type ItemAssignLicensePostRequestBody struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The addLicenses property
	addLicenses []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AssignedLicenseable
	// The removeLicenses property
	removeLicenses []i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
}

// NewItemAssignLicensePostRequestBody instantiates a new ItemAssignLicensePostRequestBody and sets the default values.
func NewItemAssignLicensePostRequestBody() *ItemAssignLicensePostRequestBody {
	m := &ItemAssignLicensePostRequestBody{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateItemAssignLicensePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemAssignLicensePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewItemAssignLicensePostRequestBody(), nil
}

// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemAssignLicensePostRequestBody) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetAddLicenses gets the addLicenses property value. The addLicenses property
func (m *ItemAssignLicensePostRequestBody) GetAddLicenses() []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AssignedLicenseable {
	return m.addLicenses
}

// GetFieldDeserializers the deserialization information for the current model
func (m *ItemAssignLicensePostRequestBody) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["addLicenses"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAssignedLicenseFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AssignedLicenseable, len(val))
			for i, v := range val {
				res[i] = v.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AssignedLicenseable)
			}
			m.SetAddLicenses(res)
		}
		return nil
	}
	res["removeLicenses"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfPrimitiveValues("uuid")
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID, len(val))
			for i, v := range val {
				res[i] = *(v.(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID))
			}
			m.SetRemoveLicenses(res)
		}
		return nil
	}
	return res
}

// GetRemoveLicenses gets the removeLicenses property value. The removeLicenses property
func (m *ItemAssignLicensePostRequestBody) GetRemoveLicenses() []i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID {
	return m.removeLicenses
}

// Serialize serializes information the current object
func (m *ItemAssignLicensePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	if m.GetAddLicenses() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAddLicenses()))
		for i, v := range m.GetAddLicenses() {
			cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
		}
		err := writer.WriteCollectionOfObjectValues("addLicenses", cast)
		if err != nil {
			return err
		}
	}
	if m.GetRemoveLicenses() != nil {
		err := writer.WriteCollectionOfUUIDValues("removeLicenses", m.GetRemoveLicenses())
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
func (m *ItemAssignLicensePostRequestBody) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetAddLicenses sets the addLicenses property value. The addLicenses property
func (m *ItemAssignLicensePostRequestBody) SetAddLicenses(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AssignedLicenseable) {
	m.addLicenses = value
}

// SetRemoveLicenses sets the removeLicenses property value. The removeLicenses property
func (m *ItemAssignLicensePostRequestBody) SetRemoveLicenses(value []i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
	m.removeLicenses = value
}
