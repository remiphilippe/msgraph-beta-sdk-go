package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DefaultManagedAppProtectionable 
type DefaultManagedAppProtectionable interface {
    ManagedAppProtectionable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAllowedAndroidDeviceManufacturers()(*string)
    GetAllowedAndroidDeviceModels()([]string)
    GetAllowedIosDeviceModels()(*string)
    GetAppActionIfAndroidDeviceManufacturerNotAllowed()(*ManagedAppRemediationAction)
    GetAppActionIfAndroidDeviceModelNotAllowed()(*ManagedAppRemediationAction)
    GetAppActionIfAndroidSafetyNetAppsVerificationFailed()(*ManagedAppRemediationAction)
    GetAppActionIfAndroidSafetyNetDeviceAttestationFailed()(*ManagedAppRemediationAction)
    GetAppActionIfDeviceLockNotSet()(*ManagedAppRemediationAction)
    GetAppActionIfIosDeviceModelNotAllowed()(*ManagedAppRemediationAction)
    GetAppDataEncryptionType()(*ManagedAppDataEncryptionType)
    GetApps()([]ManagedMobileAppable)
    GetBiometricAuthenticationBlocked()(*bool)
    GetBlockAfterCompanyPortalUpdateDeferralInDays()(*int32)
    GetConnectToVpnOnLaunch()(*bool)
    GetCustomBrowserDisplayName()(*string)
    GetCustomBrowserPackageId()(*string)
    GetCustomBrowserProtocol()(*string)
    GetCustomDialerAppDisplayName()(*string)
    GetCustomDialerAppPackageId()(*string)
    GetCustomDialerAppProtocol()(*string)
    GetCustomSettings()([]KeyValuePairable)
    GetDeployedAppCount()(*int32)
    GetDeploymentSummary()(ManagedAppPolicyDeploymentSummaryable)
    GetDeviceLockRequired()(*bool)
    GetDisableAppEncryptionIfDeviceEncryptionIsEnabled()(*bool)
    GetDisableProtectionOfManagedOutboundOpenInData()(*bool)
    GetEncryptAppData()(*bool)
    GetExemptedAppPackages()([]KeyValuePairable)
    GetExemptedAppProtocols()([]KeyValuePairable)
    GetFaceIdBlocked()(*bool)
    GetFilterOpenInToOnlyManagedApps()(*bool)
    GetMinimumRequiredCompanyPortalVersion()(*string)
    GetMinimumRequiredPatchVersion()(*string)
    GetMinimumRequiredSdkVersion()(*string)
    GetMinimumWarningCompanyPortalVersion()(*string)
    GetMinimumWarningPatchVersion()(*string)
    GetMinimumWipeCompanyPortalVersion()(*string)
    GetMinimumWipePatchVersion()(*string)
    GetMinimumWipeSdkVersion()(*string)
    GetProtectInboundDataFromUnknownSources()(*bool)
    GetRequiredAndroidSafetyNetAppsVerificationType()(*AndroidManagedAppSafetyNetAppsVerificationType)
    GetRequiredAndroidSafetyNetDeviceAttestationType()(*AndroidManagedAppSafetyNetDeviceAttestationType)
    GetRequiredAndroidSafetyNetEvaluationType()(*AndroidManagedAppSafetyNetEvaluationType)
    GetScreenCaptureBlocked()(*bool)
    GetThirdPartyKeyboardsBlocked()(*bool)
    GetWarnAfterCompanyPortalUpdateDeferralInDays()(*int32)
    GetWipeAfterCompanyPortalUpdateDeferralInDays()(*int32)
    SetAllowedAndroidDeviceManufacturers(value *string)()
    SetAllowedAndroidDeviceModels(value []string)()
    SetAllowedIosDeviceModels(value *string)()
    SetAppActionIfAndroidDeviceManufacturerNotAllowed(value *ManagedAppRemediationAction)()
    SetAppActionIfAndroidDeviceModelNotAllowed(value *ManagedAppRemediationAction)()
    SetAppActionIfAndroidSafetyNetAppsVerificationFailed(value *ManagedAppRemediationAction)()
    SetAppActionIfAndroidSafetyNetDeviceAttestationFailed(value *ManagedAppRemediationAction)()
    SetAppActionIfDeviceLockNotSet(value *ManagedAppRemediationAction)()
    SetAppActionIfIosDeviceModelNotAllowed(value *ManagedAppRemediationAction)()
    SetAppDataEncryptionType(value *ManagedAppDataEncryptionType)()
    SetApps(value []ManagedMobileAppable)()
    SetBiometricAuthenticationBlocked(value *bool)()
    SetBlockAfterCompanyPortalUpdateDeferralInDays(value *int32)()
    SetConnectToVpnOnLaunch(value *bool)()
    SetCustomBrowserDisplayName(value *string)()
    SetCustomBrowserPackageId(value *string)()
    SetCustomBrowserProtocol(value *string)()
    SetCustomDialerAppDisplayName(value *string)()
    SetCustomDialerAppPackageId(value *string)()
    SetCustomDialerAppProtocol(value *string)()
    SetCustomSettings(value []KeyValuePairable)()
    SetDeployedAppCount(value *int32)()
    SetDeploymentSummary(value ManagedAppPolicyDeploymentSummaryable)()
    SetDeviceLockRequired(value *bool)()
    SetDisableAppEncryptionIfDeviceEncryptionIsEnabled(value *bool)()
    SetDisableProtectionOfManagedOutboundOpenInData(value *bool)()
    SetEncryptAppData(value *bool)()
    SetExemptedAppPackages(value []KeyValuePairable)()
    SetExemptedAppProtocols(value []KeyValuePairable)()
    SetFaceIdBlocked(value *bool)()
    SetFilterOpenInToOnlyManagedApps(value *bool)()
    SetMinimumRequiredCompanyPortalVersion(value *string)()
    SetMinimumRequiredPatchVersion(value *string)()
    SetMinimumRequiredSdkVersion(value *string)()
    SetMinimumWarningCompanyPortalVersion(value *string)()
    SetMinimumWarningPatchVersion(value *string)()
    SetMinimumWipeCompanyPortalVersion(value *string)()
    SetMinimumWipePatchVersion(value *string)()
    SetMinimumWipeSdkVersion(value *string)()
    SetProtectInboundDataFromUnknownSources(value *bool)()
    SetRequiredAndroidSafetyNetAppsVerificationType(value *AndroidManagedAppSafetyNetAppsVerificationType)()
    SetRequiredAndroidSafetyNetDeviceAttestationType(value *AndroidManagedAppSafetyNetDeviceAttestationType)()
    SetRequiredAndroidSafetyNetEvaluationType(value *AndroidManagedAppSafetyNetEvaluationType)()
    SetScreenCaptureBlocked(value *bool)()
    SetThirdPartyKeyboardsBlocked(value *bool)()
    SetWarnAfterCompanyPortalUpdateDeferralInDays(value *int32)()
    SetWipeAfterCompanyPortalUpdateDeferralInDays(value *int32)()
}
