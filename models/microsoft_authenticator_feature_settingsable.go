package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MicrosoftAuthenticatorFeatureSettingsable 
type MicrosoftAuthenticatorFeatureSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayAppInformationRequiredState()(AuthenticationMethodFeatureConfigurationable)
    GetDisplayLocationInformationRequiredState()(AuthenticationMethodFeatureConfigurationable)
    GetNumberMatchingRequiredState()(AuthenticationMethodFeatureConfigurationable)
    GetOdataType()(*string)
    SetDisplayAppInformationRequiredState(value AuthenticationMethodFeatureConfigurationable)()
    SetDisplayLocationInformationRequiredState(value AuthenticationMethodFeatureConfigurationable)()
    SetNumberMatchingRequiredState(value AuthenticationMethodFeatureConfigurationable)()
    SetOdataType(value *string)()
}
