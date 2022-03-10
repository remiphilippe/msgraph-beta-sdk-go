package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceAppManagementable 
type DeviceAppManagementable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAndroidManagedAppProtections()([]AndroidManagedAppProtectionable)
    GetDefaultManagedAppProtections()([]DefaultManagedAppProtectionable)
    GetDeviceAppManagementTasks()([]DeviceAppManagementTaskable)
    GetEnterpriseCodeSigningCertificates()([]EnterpriseCodeSigningCertificateable)
    GetIosLobAppProvisioningConfigurations()([]IosLobAppProvisioningConfigurationable)
    GetIosManagedAppProtections()([]IosManagedAppProtectionable)
    GetIsEnabledForMicrosoftStoreForBusiness()(*bool)
    GetManagedAppPolicies()([]ManagedAppPolicyable)
    GetManagedAppRegistrations()([]ManagedAppRegistrationable)
    GetManagedAppStatuses()([]ManagedAppStatusable)
    GetManagedEBookCategories()([]ManagedEBookCategoryable)
    GetManagedEBooks()([]ManagedEBookable)
    GetMdmWindowsInformationProtectionPolicies()([]MdmWindowsInformationProtectionPolicyable)
    GetMicrosoftStoreForBusinessLanguage()(*string)
    GetMicrosoftStoreForBusinessLastCompletedApplicationSyncTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMicrosoftStoreForBusinessLastSuccessfulSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMicrosoftStoreForBusinessPortalSelection()(*MicrosoftStoreForBusinessPortalSelectionOptions)
    GetMobileAppCategories()([]MobileAppCategoryable)
    GetMobileAppConfigurations()([]ManagedDeviceMobileAppConfigurationable)
    GetMobileApps()([]MobileAppable)
    GetPolicySets()([]PolicySetable)
    GetSideLoadingKeys()([]SideLoadingKeyable)
    GetSymantecCodeSigningCertificate()(SymantecCodeSigningCertificateable)
    GetTargetedManagedAppConfigurations()([]TargetedManagedAppConfigurationable)
    GetVppTokens()([]VppTokenable)
    GetWdacSupplementalPolicies()([]WindowsDefenderApplicationControlSupplementalPolicyable)
    GetWindowsInformationProtectionDeviceRegistrations()([]WindowsInformationProtectionDeviceRegistrationable)
    GetWindowsInformationProtectionPolicies()([]WindowsInformationProtectionPolicyable)
    GetWindowsInformationProtectionWipeActions()([]WindowsInformationProtectionWipeActionable)
    GetWindowsManagementApp()(WindowsManagementAppable)
    SetAndroidManagedAppProtections(value []AndroidManagedAppProtectionable)()
    SetDefaultManagedAppProtections(value []DefaultManagedAppProtectionable)()
    SetDeviceAppManagementTasks(value []DeviceAppManagementTaskable)()
    SetEnterpriseCodeSigningCertificates(value []EnterpriseCodeSigningCertificateable)()
    SetIosLobAppProvisioningConfigurations(value []IosLobAppProvisioningConfigurationable)()
    SetIosManagedAppProtections(value []IosManagedAppProtectionable)()
    SetIsEnabledForMicrosoftStoreForBusiness(value *bool)()
    SetManagedAppPolicies(value []ManagedAppPolicyable)()
    SetManagedAppRegistrations(value []ManagedAppRegistrationable)()
    SetManagedAppStatuses(value []ManagedAppStatusable)()
    SetManagedEBookCategories(value []ManagedEBookCategoryable)()
    SetManagedEBooks(value []ManagedEBookable)()
    SetMdmWindowsInformationProtectionPolicies(value []MdmWindowsInformationProtectionPolicyable)()
    SetMicrosoftStoreForBusinessLanguage(value *string)()
    SetMicrosoftStoreForBusinessLastCompletedApplicationSyncTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMicrosoftStoreForBusinessLastSuccessfulSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMicrosoftStoreForBusinessPortalSelection(value *MicrosoftStoreForBusinessPortalSelectionOptions)()
    SetMobileAppCategories(value []MobileAppCategoryable)()
    SetMobileAppConfigurations(value []ManagedDeviceMobileAppConfigurationable)()
    SetMobileApps(value []MobileAppable)()
    SetPolicySets(value []PolicySetable)()
    SetSideLoadingKeys(value []SideLoadingKeyable)()
    SetSymantecCodeSigningCertificate(value SymantecCodeSigningCertificateable)()
    SetTargetedManagedAppConfigurations(value []TargetedManagedAppConfigurationable)()
    SetVppTokens(value []VppTokenable)()
    SetWdacSupplementalPolicies(value []WindowsDefenderApplicationControlSupplementalPolicyable)()
    SetWindowsInformationProtectionDeviceRegistrations(value []WindowsInformationProtectionDeviceRegistrationable)()
    SetWindowsInformationProtectionPolicies(value []WindowsInformationProtectionPolicyable)()
    SetWindowsInformationProtectionWipeActions(value []WindowsInformationProtectionWipeActionable)()
    SetWindowsManagementApp(value WindowsManagementAppable)()
}
