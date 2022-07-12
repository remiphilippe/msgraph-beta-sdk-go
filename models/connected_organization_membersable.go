package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ConnectedOrganizationMembersable 
type ConnectedOrganizationMembersable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    UserSetable
    GetDescription()(*string)
    GetId()(*string)
    SetDescription(value *string)()
    SetId(value *string)()
}
