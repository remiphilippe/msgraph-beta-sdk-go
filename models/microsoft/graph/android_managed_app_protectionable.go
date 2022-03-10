package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AndroidManagedAppProtectionable 
type AndroidManagedAppProtectionable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    TargetedManagedAppProtectionable
    GetAllowedAndroidDeviceManufacturers()(*string)
    GetAllowedAndroidDeviceModels()([]string)
    GetAppActionIfAndroidDeviceManufacturerNotAllowed()(*ManagedAppRemediationAction)
    GetAppActionIfAndroidDeviceModelNotAllowed()(*ManagedAppRemediationAction)
    GetAppActionIfAndroidSafetyNetAppsVerificationFailed()(*ManagedAppRemediationAction)
    GetAppActionIfAndroidSafetyNetDeviceAttestationFailed()(*ManagedAppRemediationAction)
    GetAppActionIfDeviceLockNotSet()(*ManagedAppRemediationAction)
    GetApprovedKeyboards()([]KeyValuePairable)
    GetApps()([]ManagedMobileAppable)
    GetBiometricAuthenticationBlocked()(*bool)
    GetBlockAfterCompanyPortalUpdateDeferralInDays()(*int32)
    GetConnectToVpnOnLaunch()(*bool)
    GetCustomBrowserDisplayName()(*string)
    GetCustomBrowserPackageId()(*string)
    GetCustomDialerAppDisplayName()(*string)
    GetCustomDialerAppPackageId()(*string)
    GetDeployedAppCount()(*int32)
    GetDeploymentSummary()(ManagedAppPolicyDeploymentSummaryable)
    GetDeviceLockRequired()(*bool)
    GetDisableAppEncryptionIfDeviceEncryptionIsEnabled()(*bool)
    GetEncryptAppData()(*bool)
    GetExemptedAppPackages()([]KeyValuePairable)
    GetKeyboardsRestricted()(*bool)
    GetMinimumRequiredCompanyPortalVersion()(*string)
    GetMinimumRequiredPatchVersion()(*string)
    GetMinimumWarningCompanyPortalVersion()(*string)
    GetMinimumWarningPatchVersion()(*string)
    GetMinimumWipeCompanyPortalVersion()(*string)
    GetMinimumWipePatchVersion()(*string)
    GetRequiredAndroidSafetyNetAppsVerificationType()(*AndroidManagedAppSafetyNetAppsVerificationType)
    GetRequiredAndroidSafetyNetDeviceAttestationType()(*AndroidManagedAppSafetyNetDeviceAttestationType)
    GetRequiredAndroidSafetyNetEvaluationType()(*AndroidManagedAppSafetyNetEvaluationType)
    GetScreenCaptureBlocked()(*bool)
    GetWarnAfterCompanyPortalUpdateDeferralInDays()(*int32)
    GetWipeAfterCompanyPortalUpdateDeferralInDays()(*int32)
    SetAllowedAndroidDeviceManufacturers(value *string)()
    SetAllowedAndroidDeviceModels(value []string)()
    SetAppActionIfAndroidDeviceManufacturerNotAllowed(value *ManagedAppRemediationAction)()
    SetAppActionIfAndroidDeviceModelNotAllowed(value *ManagedAppRemediationAction)()
    SetAppActionIfAndroidSafetyNetAppsVerificationFailed(value *ManagedAppRemediationAction)()
    SetAppActionIfAndroidSafetyNetDeviceAttestationFailed(value *ManagedAppRemediationAction)()
    SetAppActionIfDeviceLockNotSet(value *ManagedAppRemediationAction)()
    SetApprovedKeyboards(value []KeyValuePairable)()
    SetApps(value []ManagedMobileAppable)()
    SetBiometricAuthenticationBlocked(value *bool)()
    SetBlockAfterCompanyPortalUpdateDeferralInDays(value *int32)()
    SetConnectToVpnOnLaunch(value *bool)()
    SetCustomBrowserDisplayName(value *string)()
    SetCustomBrowserPackageId(value *string)()
    SetCustomDialerAppDisplayName(value *string)()
    SetCustomDialerAppPackageId(value *string)()
    SetDeployedAppCount(value *int32)()
    SetDeploymentSummary(value ManagedAppPolicyDeploymentSummaryable)()
    SetDeviceLockRequired(value *bool)()
    SetDisableAppEncryptionIfDeviceEncryptionIsEnabled(value *bool)()
    SetEncryptAppData(value *bool)()
    SetExemptedAppPackages(value []KeyValuePairable)()
    SetKeyboardsRestricted(value *bool)()
    SetMinimumRequiredCompanyPortalVersion(value *string)()
    SetMinimumRequiredPatchVersion(value *string)()
    SetMinimumWarningCompanyPortalVersion(value *string)()
    SetMinimumWarningPatchVersion(value *string)()
    SetMinimumWipeCompanyPortalVersion(value *string)()
    SetMinimumWipePatchVersion(value *string)()
    SetRequiredAndroidSafetyNetAppsVerificationType(value *AndroidManagedAppSafetyNetAppsVerificationType)()
    SetRequiredAndroidSafetyNetDeviceAttestationType(value *AndroidManagedAppSafetyNetDeviceAttestationType)()
    SetRequiredAndroidSafetyNetEvaluationType(value *AndroidManagedAppSafetyNetEvaluationType)()
    SetScreenCaptureBlocked(value *bool)()
    SetWarnAfterCompanyPortalUpdateDeferralInDays(value *int32)()
    SetWipeAfterCompanyPortalUpdateDeferralInDays(value *int32)()
}
