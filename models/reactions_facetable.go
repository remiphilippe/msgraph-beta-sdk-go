package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ReactionsFacetable 
type ReactionsFacetable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCommentCount()(*int32)
    GetLikeCount()(*int32)
    GetOdataType()(*string)
    GetShareCount()(*int32)
    SetCommentCount(value *int32)()
    SetLikeCount(value *int32)()
    SetOdataType(value *string)()
    SetShareCount(value *int32)()
}
