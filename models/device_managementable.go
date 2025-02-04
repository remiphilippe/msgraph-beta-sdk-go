package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementable 
type DeviceManagementable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccountMoveCompletionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetAdminConsent()(AdminConsentable)
    GetAdvancedThreatProtectionOnboardingStateSummary()(AdvancedThreatProtectionOnboardingStateSummaryable)
    GetAndroidDeviceOwnerEnrollmentProfiles()([]AndroidDeviceOwnerEnrollmentProfileable)
    GetAndroidForWorkAppConfigurationSchemas()([]AndroidForWorkAppConfigurationSchemaable)
    GetAndroidForWorkEnrollmentProfiles()([]AndroidForWorkEnrollmentProfileable)
    GetAndroidForWorkSettings()(AndroidForWorkSettingsable)
    GetAndroidManagedStoreAccountEnterpriseSettings()(AndroidManagedStoreAccountEnterpriseSettingsable)
    GetAndroidManagedStoreAppConfigurationSchemas()([]AndroidManagedStoreAppConfigurationSchemaable)
    GetApplePushNotificationCertificate()(ApplePushNotificationCertificateable)
    GetAppleUserInitiatedEnrollmentProfiles()([]AppleUserInitiatedEnrollmentProfileable)
    GetAssignmentFilters()([]DeviceAndAppManagementAssignmentFilterable)
    GetAuditEvents()([]AuditEventable)
    GetAutopilotEvents()([]DeviceManagementAutopilotEventable)
    GetCartToClassAssociations()([]CartToClassAssociationable)
    GetCategories()([]DeviceManagementSettingCategoryable)
    GetCertificateConnectorDetails()([]CertificateConnectorDetailsable)
    GetChromeOSOnboardingSettings()([]ChromeOSOnboardingSettingsable)
    GetCloudPCConnectivityIssues()([]CloudPCConnectivityIssueable)
    GetComanagedDevices()([]ManagedDeviceable)
    GetComanagementEligibleDevices()([]ComanagementEligibleDeviceable)
    GetComplianceCategories()([]DeviceManagementConfigurationCategoryable)
    GetComplianceManagementPartners()([]ComplianceManagementPartnerable)
    GetCompliancePolicies()([]DeviceManagementCompliancePolicyable)
    GetComplianceSettings()([]DeviceManagementConfigurationSettingDefinitionable)
    GetConditionalAccessSettings()(OnPremisesConditionalAccessSettingsable)
    GetConfigManagerCollections()([]ConfigManagerCollectionable)
    GetConfigurationCategories()([]DeviceManagementConfigurationCategoryable)
    GetConfigurationPolicies()([]DeviceManagementConfigurationPolicyable)
    GetConfigurationPolicyTemplates()([]DeviceManagementConfigurationPolicyTemplateable)
    GetConfigurationSettings()([]DeviceManagementConfigurationSettingDefinitionable)
    GetDataProcessorServiceForWindowsFeaturesOnboarding()(DataProcessorServiceForWindowsFeaturesOnboardingable)
    GetDataSharingConsents()([]DataSharingConsentable)
    GetDepOnboardingSettings()([]DepOnboardingSettingable)
    GetDerivedCredentials()([]DeviceManagementDerivedCredentialSettingsable)
    GetDetectedApps()([]DetectedAppable)
    GetDeviceCategories()([]DeviceCategoryable)
    GetDeviceCompliancePolicies()([]DeviceCompliancePolicyable)
    GetDeviceCompliancePolicyDeviceStateSummary()(DeviceCompliancePolicyDeviceStateSummaryable)
    GetDeviceCompliancePolicySettingStateSummaries()([]DeviceCompliancePolicySettingStateSummaryable)
    GetDeviceComplianceReportSummarizationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDeviceComplianceScripts()([]DeviceComplianceScriptable)
    GetDeviceConfigurationConflictSummary()([]DeviceConfigurationConflictSummaryable)
    GetDeviceConfigurationDeviceStateSummaries()(DeviceConfigurationDeviceStateSummaryable)
    GetDeviceConfigurationRestrictedAppsViolations()([]RestrictedAppsViolationable)
    GetDeviceConfigurations()([]DeviceConfigurationable)
    GetDeviceConfigurationsAllManagedDeviceCertificateStates()([]ManagedAllDeviceCertificateStateable)
    GetDeviceConfigurationUserStateSummaries()(DeviceConfigurationUserStateSummaryable)
    GetDeviceCustomAttributeShellScripts()([]DeviceCustomAttributeShellScriptable)
    GetDeviceEnrollmentConfigurations()([]DeviceEnrollmentConfigurationable)
    GetDeviceHealthScripts()([]DeviceHealthScriptable)
    GetDeviceManagementPartners()([]DeviceManagementPartnerable)
    GetDeviceManagementScripts()([]DeviceManagementScriptable)
    GetDeviceProtectionOverview()(DeviceProtectionOverviewable)
    GetDeviceShellScripts()([]DeviceShellScriptable)
    GetDomainJoinConnectors()([]DeviceManagementDomainJoinConnectorable)
    GetEmbeddedSIMActivationCodePools()([]EmbeddedSIMActivationCodePoolable)
    GetExchangeConnectors()([]DeviceManagementExchangeConnectorable)
    GetExchangeOnPremisesPolicies()([]DeviceManagementExchangeOnPremisesPolicyable)
    GetExchangeOnPremisesPolicy()(DeviceManagementExchangeOnPremisesPolicyable)
    GetGroupPolicyCategories()([]GroupPolicyCategoryable)
    GetGroupPolicyConfigurations()([]GroupPolicyConfigurationable)
    GetGroupPolicyDefinitionFiles()([]GroupPolicyDefinitionFileable)
    GetGroupPolicyDefinitions()([]GroupPolicyDefinitionable)
    GetGroupPolicyMigrationReports()([]GroupPolicyMigrationReportable)
    GetGroupPolicyObjectFiles()([]GroupPolicyObjectFileable)
    GetGroupPolicyUploadedDefinitionFiles()([]GroupPolicyUploadedDefinitionFileable)
    GetImportedDeviceIdentities()([]ImportedDeviceIdentityable)
    GetImportedWindowsAutopilotDeviceIdentities()([]ImportedWindowsAutopilotDeviceIdentityable)
    GetIntents()([]DeviceManagementIntentable)
    GetIntuneAccountId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetIntuneBrand()(IntuneBrandable)
    GetIntuneBrandingProfiles()([]IntuneBrandingProfileable)
    GetIosUpdateStatuses()([]IosUpdateDeviceStatusable)
    GetLastReportAggregationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLegacyPcManangementEnabled()(*bool)
    GetMacOSSoftwareUpdateAccountSummaries()([]MacOSSoftwareUpdateAccountSummaryable)
    GetManagedDeviceCleanupSettings()(ManagedDeviceCleanupSettingsable)
    GetManagedDeviceEncryptionStates()([]ManagedDeviceEncryptionStateable)
    GetManagedDeviceOverview()(ManagedDeviceOverviewable)
    GetManagedDevices()([]ManagedDeviceable)
    GetMaximumDepTokens()(*int32)
    GetMicrosoftTunnelConfigurations()([]MicrosoftTunnelConfigurationable)
    GetMicrosoftTunnelHealthThresholds()([]MicrosoftTunnelHealthThresholdable)
    GetMicrosoftTunnelServerLogCollectionResponses()([]MicrosoftTunnelServerLogCollectionResponseable)
    GetMicrosoftTunnelSites()([]MicrosoftTunnelSiteable)
    GetMobileAppTroubleshootingEvents()([]MobileAppTroubleshootingEventable)
    GetMobileThreatDefenseConnectors()([]MobileThreatDefenseConnectorable)
    GetNdesConnectors()([]NdesConnectorable)
    GetNotificationMessageTemplates()([]NotificationMessageTemplateable)
    GetOemWarrantyInformationOnboarding()([]OemWarrantyInformationOnboardingable)
    GetRemoteActionAudits()([]RemoteActionAuditable)
    GetRemoteAssistancePartners()([]RemoteAssistancePartnerable)
    GetRemoteAssistanceSettings()(RemoteAssistanceSettingsable)
    GetReports()(DeviceManagementReportsable)
    GetResourceAccessProfiles()([]DeviceManagementResourceAccessProfileBaseable)
    GetResourceOperations()([]ResourceOperationable)
    GetReusablePolicySettings()([]DeviceManagementReusablePolicySettingable)
    GetReusableSettings()([]DeviceManagementConfigurationSettingDefinitionable)
    GetRoleAssignments()([]DeviceAndAppManagementRoleAssignmentable)
    GetRoleDefinitions()([]RoleDefinitionable)
    GetRoleScopeTags()([]RoleScopeTagable)
    GetSettingDefinitions()([]DeviceManagementSettingDefinitionable)
    GetSettings()(DeviceManagementSettingsable)
    GetSoftwareUpdateStatusSummary()(SoftwareUpdateStatusSummaryable)
    GetSubscriptions()(*DeviceManagementSubscriptions)
    GetSubscriptionState()(*DeviceManagementSubscriptionState)
    GetTelecomExpenseManagementPartners()([]TelecomExpenseManagementPartnerable)
    GetTemplates()([]DeviceManagementTemplateable)
    GetTemplateSettings()([]DeviceManagementConfigurationSettingTemplateable)
    GetTenantAttachRBAC()(TenantAttachRBACable)
    GetTermsAndConditions()([]TermsAndConditionsable)
    GetTroubleshootingEvents()([]DeviceManagementTroubleshootingEventable)
    GetUnlicensedAdminstratorsEnabled()(*bool)
    GetUserExperienceAnalyticsAnomaly()([]UserExperienceAnalyticsAnomalyable)
    GetUserExperienceAnalyticsAnomalyDevice()([]UserExperienceAnalyticsAnomalyDeviceable)
    GetUserExperienceAnalyticsAnomalySeverityOverview()(UserExperienceAnalyticsAnomalySeverityOverviewable)
    GetUserExperienceAnalyticsAppHealthApplicationPerformance()([]UserExperienceAnalyticsAppHealthApplicationPerformanceable)
    GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion()([]UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionable)
    GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails()([]UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsable)
    GetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId()([]UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdable)
    GetUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion()([]UserExperienceAnalyticsAppHealthAppPerformanceByOSVersionable)
    GetUserExperienceAnalyticsAppHealthDeviceModelPerformance()([]UserExperienceAnalyticsAppHealthDeviceModelPerformanceable)
    GetUserExperienceAnalyticsAppHealthDevicePerformance()([]UserExperienceAnalyticsAppHealthDevicePerformanceable)
    GetUserExperienceAnalyticsAppHealthDevicePerformanceDetails()([]UserExperienceAnalyticsAppHealthDevicePerformanceDetailsable)
    GetUserExperienceAnalyticsAppHealthOSVersionPerformance()([]UserExperienceAnalyticsAppHealthOSVersionPerformanceable)
    GetUserExperienceAnalyticsAppHealthOverview()(UserExperienceAnalyticsCategoryable)
    GetUserExperienceAnalyticsBaselines()([]UserExperienceAnalyticsBaselineable)
    GetUserExperienceAnalyticsBatteryHealthAppImpact()([]UserExperienceAnalyticsBatteryHealthAppImpactable)
    GetUserExperienceAnalyticsBatteryHealthCapacityDetails()(UserExperienceAnalyticsBatteryHealthCapacityDetailsable)
    GetUserExperienceAnalyticsBatteryHealthDeviceAppImpact()([]UserExperienceAnalyticsBatteryHealthDeviceAppImpactable)
    GetUserExperienceAnalyticsBatteryHealthDevicePerformance()([]UserExperienceAnalyticsBatteryHealthDevicePerformanceable)
    GetUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory()([]UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryable)
    GetUserExperienceAnalyticsBatteryHealthModelPerformance()([]UserExperienceAnalyticsBatteryHealthModelPerformanceable)
    GetUserExperienceAnalyticsBatteryHealthOsPerformance()([]UserExperienceAnalyticsBatteryHealthOsPerformanceable)
    GetUserExperienceAnalyticsBatteryHealthRuntimeDetails()(UserExperienceAnalyticsBatteryHealthRuntimeDetailsable)
    GetUserExperienceAnalyticsCategories()([]UserExperienceAnalyticsCategoryable)
    GetUserExperienceAnalyticsDeviceMetricHistory()([]UserExperienceAnalyticsMetricHistoryable)
    GetUserExperienceAnalyticsDevicePerformance()([]UserExperienceAnalyticsDevicePerformanceable)
    GetUserExperienceAnalyticsDeviceScope()(UserExperienceAnalyticsDeviceScopeable)
    GetUserExperienceAnalyticsDeviceScopes()([]UserExperienceAnalyticsDeviceScopeable)
    GetUserExperienceAnalyticsDeviceScores()([]UserExperienceAnalyticsDeviceScoresable)
    GetUserExperienceAnalyticsDeviceStartupHistory()([]UserExperienceAnalyticsDeviceStartupHistoryable)
    GetUserExperienceAnalyticsDeviceStartupProcesses()([]UserExperienceAnalyticsDeviceStartupProcessable)
    GetUserExperienceAnalyticsDeviceStartupProcessPerformance()([]UserExperienceAnalyticsDeviceStartupProcessPerformanceable)
    GetUserExperienceAnalyticsDevicesWithoutCloudIdentity()([]UserExperienceAnalyticsDeviceWithoutCloudIdentityable)
    GetUserExperienceAnalyticsDeviceTimelineEvents()([]UserExperienceAnalyticsDeviceTimelineEventsable)
    GetUserExperienceAnalyticsImpactingProcess()([]UserExperienceAnalyticsImpactingProcessable)
    GetUserExperienceAnalyticsMetricHistory()([]UserExperienceAnalyticsMetricHistoryable)
    GetUserExperienceAnalyticsModelScores()([]UserExperienceAnalyticsModelScoresable)
    GetUserExperienceAnalyticsNotAutopilotReadyDevice()([]UserExperienceAnalyticsNotAutopilotReadyDeviceable)
    GetUserExperienceAnalyticsOverview()(UserExperienceAnalyticsOverviewable)
    GetUserExperienceAnalyticsRemoteConnection()([]UserExperienceAnalyticsRemoteConnectionable)
    GetUserExperienceAnalyticsResourcePerformance()([]UserExperienceAnalyticsResourcePerformanceable)
    GetUserExperienceAnalyticsScoreHistory()([]UserExperienceAnalyticsScoreHistoryable)
    GetUserExperienceAnalyticsSettings()(UserExperienceAnalyticsSettingsable)
    GetUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric()(UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricable)
    GetUserExperienceAnalyticsWorkFromAnywhereMetrics()([]UserExperienceAnalyticsWorkFromAnywhereMetricable)
    GetUserExperienceAnalyticsWorkFromAnywhereModelPerformance()([]UserExperienceAnalyticsWorkFromAnywhereModelPerformanceable)
    GetUserPfxCertificates()([]UserPFXCertificateable)
    GetVirtualEndpoint()(VirtualEndpointable)
    GetWindowsAutopilotDeploymentProfiles()([]WindowsAutopilotDeploymentProfileable)
    GetWindowsAutopilotDeviceIdentities()([]WindowsAutopilotDeviceIdentityable)
    GetWindowsAutopilotSettings()(WindowsAutopilotSettingsable)
    GetWindowsDriverUpdateProfiles()([]WindowsDriverUpdateProfileable)
    GetWindowsFeatureUpdateProfiles()([]WindowsFeatureUpdateProfileable)
    GetWindowsInformationProtectionAppLearningSummaries()([]WindowsInformationProtectionAppLearningSummaryable)
    GetWindowsInformationProtectionNetworkLearningSummaries()([]WindowsInformationProtectionNetworkLearningSummaryable)
    GetWindowsMalwareInformation()([]WindowsMalwareInformationable)
    GetWindowsMalwareOverview()(WindowsMalwareOverviewable)
    GetWindowsQualityUpdateProfiles()([]WindowsQualityUpdateProfileable)
    GetWindowsUpdateCatalogItems()([]WindowsUpdateCatalogItemable)
    GetZebraFotaArtifacts()([]ZebraFotaArtifactable)
    GetZebraFotaConnector()(ZebraFotaConnectorable)
    GetZebraFotaDeployments()([]ZebraFotaDeploymentable)
    SetAccountMoveCompletionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetAdminConsent(value AdminConsentable)()
    SetAdvancedThreatProtectionOnboardingStateSummary(value AdvancedThreatProtectionOnboardingStateSummaryable)()
    SetAndroidDeviceOwnerEnrollmentProfiles(value []AndroidDeviceOwnerEnrollmentProfileable)()
    SetAndroidForWorkAppConfigurationSchemas(value []AndroidForWorkAppConfigurationSchemaable)()
    SetAndroidForWorkEnrollmentProfiles(value []AndroidForWorkEnrollmentProfileable)()
    SetAndroidForWorkSettings(value AndroidForWorkSettingsable)()
    SetAndroidManagedStoreAccountEnterpriseSettings(value AndroidManagedStoreAccountEnterpriseSettingsable)()
    SetAndroidManagedStoreAppConfigurationSchemas(value []AndroidManagedStoreAppConfigurationSchemaable)()
    SetApplePushNotificationCertificate(value ApplePushNotificationCertificateable)()
    SetAppleUserInitiatedEnrollmentProfiles(value []AppleUserInitiatedEnrollmentProfileable)()
    SetAssignmentFilters(value []DeviceAndAppManagementAssignmentFilterable)()
    SetAuditEvents(value []AuditEventable)()
    SetAutopilotEvents(value []DeviceManagementAutopilotEventable)()
    SetCartToClassAssociations(value []CartToClassAssociationable)()
    SetCategories(value []DeviceManagementSettingCategoryable)()
    SetCertificateConnectorDetails(value []CertificateConnectorDetailsable)()
    SetChromeOSOnboardingSettings(value []ChromeOSOnboardingSettingsable)()
    SetCloudPCConnectivityIssues(value []CloudPCConnectivityIssueable)()
    SetComanagedDevices(value []ManagedDeviceable)()
    SetComanagementEligibleDevices(value []ComanagementEligibleDeviceable)()
    SetComplianceCategories(value []DeviceManagementConfigurationCategoryable)()
    SetComplianceManagementPartners(value []ComplianceManagementPartnerable)()
    SetCompliancePolicies(value []DeviceManagementCompliancePolicyable)()
    SetComplianceSettings(value []DeviceManagementConfigurationSettingDefinitionable)()
    SetConditionalAccessSettings(value OnPremisesConditionalAccessSettingsable)()
    SetConfigManagerCollections(value []ConfigManagerCollectionable)()
    SetConfigurationCategories(value []DeviceManagementConfigurationCategoryable)()
    SetConfigurationPolicies(value []DeviceManagementConfigurationPolicyable)()
    SetConfigurationPolicyTemplates(value []DeviceManagementConfigurationPolicyTemplateable)()
    SetConfigurationSettings(value []DeviceManagementConfigurationSettingDefinitionable)()
    SetDataProcessorServiceForWindowsFeaturesOnboarding(value DataProcessorServiceForWindowsFeaturesOnboardingable)()
    SetDataSharingConsents(value []DataSharingConsentable)()
    SetDepOnboardingSettings(value []DepOnboardingSettingable)()
    SetDerivedCredentials(value []DeviceManagementDerivedCredentialSettingsable)()
    SetDetectedApps(value []DetectedAppable)()
    SetDeviceCategories(value []DeviceCategoryable)()
    SetDeviceCompliancePolicies(value []DeviceCompliancePolicyable)()
    SetDeviceCompliancePolicyDeviceStateSummary(value DeviceCompliancePolicyDeviceStateSummaryable)()
    SetDeviceCompliancePolicySettingStateSummaries(value []DeviceCompliancePolicySettingStateSummaryable)()
    SetDeviceComplianceReportSummarizationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDeviceComplianceScripts(value []DeviceComplianceScriptable)()
    SetDeviceConfigurationConflictSummary(value []DeviceConfigurationConflictSummaryable)()
    SetDeviceConfigurationDeviceStateSummaries(value DeviceConfigurationDeviceStateSummaryable)()
    SetDeviceConfigurationRestrictedAppsViolations(value []RestrictedAppsViolationable)()
    SetDeviceConfigurations(value []DeviceConfigurationable)()
    SetDeviceConfigurationsAllManagedDeviceCertificateStates(value []ManagedAllDeviceCertificateStateable)()
    SetDeviceConfigurationUserStateSummaries(value DeviceConfigurationUserStateSummaryable)()
    SetDeviceCustomAttributeShellScripts(value []DeviceCustomAttributeShellScriptable)()
    SetDeviceEnrollmentConfigurations(value []DeviceEnrollmentConfigurationable)()
    SetDeviceHealthScripts(value []DeviceHealthScriptable)()
    SetDeviceManagementPartners(value []DeviceManagementPartnerable)()
    SetDeviceManagementScripts(value []DeviceManagementScriptable)()
    SetDeviceProtectionOverview(value DeviceProtectionOverviewable)()
    SetDeviceShellScripts(value []DeviceShellScriptable)()
    SetDomainJoinConnectors(value []DeviceManagementDomainJoinConnectorable)()
    SetEmbeddedSIMActivationCodePools(value []EmbeddedSIMActivationCodePoolable)()
    SetExchangeConnectors(value []DeviceManagementExchangeConnectorable)()
    SetExchangeOnPremisesPolicies(value []DeviceManagementExchangeOnPremisesPolicyable)()
    SetExchangeOnPremisesPolicy(value DeviceManagementExchangeOnPremisesPolicyable)()
    SetGroupPolicyCategories(value []GroupPolicyCategoryable)()
    SetGroupPolicyConfigurations(value []GroupPolicyConfigurationable)()
    SetGroupPolicyDefinitionFiles(value []GroupPolicyDefinitionFileable)()
    SetGroupPolicyDefinitions(value []GroupPolicyDefinitionable)()
    SetGroupPolicyMigrationReports(value []GroupPolicyMigrationReportable)()
    SetGroupPolicyObjectFiles(value []GroupPolicyObjectFileable)()
    SetGroupPolicyUploadedDefinitionFiles(value []GroupPolicyUploadedDefinitionFileable)()
    SetImportedDeviceIdentities(value []ImportedDeviceIdentityable)()
    SetImportedWindowsAutopilotDeviceIdentities(value []ImportedWindowsAutopilotDeviceIdentityable)()
    SetIntents(value []DeviceManagementIntentable)()
    SetIntuneAccountId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetIntuneBrand(value IntuneBrandable)()
    SetIntuneBrandingProfiles(value []IntuneBrandingProfileable)()
    SetIosUpdateStatuses(value []IosUpdateDeviceStatusable)()
    SetLastReportAggregationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLegacyPcManangementEnabled(value *bool)()
    SetMacOSSoftwareUpdateAccountSummaries(value []MacOSSoftwareUpdateAccountSummaryable)()
    SetManagedDeviceCleanupSettings(value ManagedDeviceCleanupSettingsable)()
    SetManagedDeviceEncryptionStates(value []ManagedDeviceEncryptionStateable)()
    SetManagedDeviceOverview(value ManagedDeviceOverviewable)()
    SetManagedDevices(value []ManagedDeviceable)()
    SetMaximumDepTokens(value *int32)()
    SetMicrosoftTunnelConfigurations(value []MicrosoftTunnelConfigurationable)()
    SetMicrosoftTunnelHealthThresholds(value []MicrosoftTunnelHealthThresholdable)()
    SetMicrosoftTunnelServerLogCollectionResponses(value []MicrosoftTunnelServerLogCollectionResponseable)()
    SetMicrosoftTunnelSites(value []MicrosoftTunnelSiteable)()
    SetMobileAppTroubleshootingEvents(value []MobileAppTroubleshootingEventable)()
    SetMobileThreatDefenseConnectors(value []MobileThreatDefenseConnectorable)()
    SetNdesConnectors(value []NdesConnectorable)()
    SetNotificationMessageTemplates(value []NotificationMessageTemplateable)()
    SetOemWarrantyInformationOnboarding(value []OemWarrantyInformationOnboardingable)()
    SetRemoteActionAudits(value []RemoteActionAuditable)()
    SetRemoteAssistancePartners(value []RemoteAssistancePartnerable)()
    SetRemoteAssistanceSettings(value RemoteAssistanceSettingsable)()
    SetReports(value DeviceManagementReportsable)()
    SetResourceAccessProfiles(value []DeviceManagementResourceAccessProfileBaseable)()
    SetResourceOperations(value []ResourceOperationable)()
    SetReusablePolicySettings(value []DeviceManagementReusablePolicySettingable)()
    SetReusableSettings(value []DeviceManagementConfigurationSettingDefinitionable)()
    SetRoleAssignments(value []DeviceAndAppManagementRoleAssignmentable)()
    SetRoleDefinitions(value []RoleDefinitionable)()
    SetRoleScopeTags(value []RoleScopeTagable)()
    SetSettingDefinitions(value []DeviceManagementSettingDefinitionable)()
    SetSettings(value DeviceManagementSettingsable)()
    SetSoftwareUpdateStatusSummary(value SoftwareUpdateStatusSummaryable)()
    SetSubscriptions(value *DeviceManagementSubscriptions)()
    SetSubscriptionState(value *DeviceManagementSubscriptionState)()
    SetTelecomExpenseManagementPartners(value []TelecomExpenseManagementPartnerable)()
    SetTemplates(value []DeviceManagementTemplateable)()
    SetTemplateSettings(value []DeviceManagementConfigurationSettingTemplateable)()
    SetTenantAttachRBAC(value TenantAttachRBACable)()
    SetTermsAndConditions(value []TermsAndConditionsable)()
    SetTroubleshootingEvents(value []DeviceManagementTroubleshootingEventable)()
    SetUnlicensedAdminstratorsEnabled(value *bool)()
    SetUserExperienceAnalyticsAnomaly(value []UserExperienceAnalyticsAnomalyable)()
    SetUserExperienceAnalyticsAnomalyDevice(value []UserExperienceAnalyticsAnomalyDeviceable)()
    SetUserExperienceAnalyticsAnomalySeverityOverview(value UserExperienceAnalyticsAnomalySeverityOverviewable)()
    SetUserExperienceAnalyticsAppHealthApplicationPerformance(value []UserExperienceAnalyticsAppHealthApplicationPerformanceable)()
    SetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion(value []UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionable)()
    SetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails(value []UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsable)()
    SetUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId(value []UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceIdable)()
    SetUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion(value []UserExperienceAnalyticsAppHealthAppPerformanceByOSVersionable)()
    SetUserExperienceAnalyticsAppHealthDeviceModelPerformance(value []UserExperienceAnalyticsAppHealthDeviceModelPerformanceable)()
    SetUserExperienceAnalyticsAppHealthDevicePerformance(value []UserExperienceAnalyticsAppHealthDevicePerformanceable)()
    SetUserExperienceAnalyticsAppHealthDevicePerformanceDetails(value []UserExperienceAnalyticsAppHealthDevicePerformanceDetailsable)()
    SetUserExperienceAnalyticsAppHealthOSVersionPerformance(value []UserExperienceAnalyticsAppHealthOSVersionPerformanceable)()
    SetUserExperienceAnalyticsAppHealthOverview(value UserExperienceAnalyticsCategoryable)()
    SetUserExperienceAnalyticsBaselines(value []UserExperienceAnalyticsBaselineable)()
    SetUserExperienceAnalyticsBatteryHealthAppImpact(value []UserExperienceAnalyticsBatteryHealthAppImpactable)()
    SetUserExperienceAnalyticsBatteryHealthCapacityDetails(value UserExperienceAnalyticsBatteryHealthCapacityDetailsable)()
    SetUserExperienceAnalyticsBatteryHealthDeviceAppImpact(value []UserExperienceAnalyticsBatteryHealthDeviceAppImpactable)()
    SetUserExperienceAnalyticsBatteryHealthDevicePerformance(value []UserExperienceAnalyticsBatteryHealthDevicePerformanceable)()
    SetUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory(value []UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryable)()
    SetUserExperienceAnalyticsBatteryHealthModelPerformance(value []UserExperienceAnalyticsBatteryHealthModelPerformanceable)()
    SetUserExperienceAnalyticsBatteryHealthOsPerformance(value []UserExperienceAnalyticsBatteryHealthOsPerformanceable)()
    SetUserExperienceAnalyticsBatteryHealthRuntimeDetails(value UserExperienceAnalyticsBatteryHealthRuntimeDetailsable)()
    SetUserExperienceAnalyticsCategories(value []UserExperienceAnalyticsCategoryable)()
    SetUserExperienceAnalyticsDeviceMetricHistory(value []UserExperienceAnalyticsMetricHistoryable)()
    SetUserExperienceAnalyticsDevicePerformance(value []UserExperienceAnalyticsDevicePerformanceable)()
    SetUserExperienceAnalyticsDeviceScope(value UserExperienceAnalyticsDeviceScopeable)()
    SetUserExperienceAnalyticsDeviceScopes(value []UserExperienceAnalyticsDeviceScopeable)()
    SetUserExperienceAnalyticsDeviceScores(value []UserExperienceAnalyticsDeviceScoresable)()
    SetUserExperienceAnalyticsDeviceStartupHistory(value []UserExperienceAnalyticsDeviceStartupHistoryable)()
    SetUserExperienceAnalyticsDeviceStartupProcesses(value []UserExperienceAnalyticsDeviceStartupProcessable)()
    SetUserExperienceAnalyticsDeviceStartupProcessPerformance(value []UserExperienceAnalyticsDeviceStartupProcessPerformanceable)()
    SetUserExperienceAnalyticsDevicesWithoutCloudIdentity(value []UserExperienceAnalyticsDeviceWithoutCloudIdentityable)()
    SetUserExperienceAnalyticsDeviceTimelineEvents(value []UserExperienceAnalyticsDeviceTimelineEventsable)()
    SetUserExperienceAnalyticsImpactingProcess(value []UserExperienceAnalyticsImpactingProcessable)()
    SetUserExperienceAnalyticsMetricHistory(value []UserExperienceAnalyticsMetricHistoryable)()
    SetUserExperienceAnalyticsModelScores(value []UserExperienceAnalyticsModelScoresable)()
    SetUserExperienceAnalyticsNotAutopilotReadyDevice(value []UserExperienceAnalyticsNotAutopilotReadyDeviceable)()
    SetUserExperienceAnalyticsOverview(value UserExperienceAnalyticsOverviewable)()
    SetUserExperienceAnalyticsRemoteConnection(value []UserExperienceAnalyticsRemoteConnectionable)()
    SetUserExperienceAnalyticsResourcePerformance(value []UserExperienceAnalyticsResourcePerformanceable)()
    SetUserExperienceAnalyticsScoreHistory(value []UserExperienceAnalyticsScoreHistoryable)()
    SetUserExperienceAnalyticsSettings(value UserExperienceAnalyticsSettingsable)()
    SetUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric(value UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricable)()
    SetUserExperienceAnalyticsWorkFromAnywhereMetrics(value []UserExperienceAnalyticsWorkFromAnywhereMetricable)()
    SetUserExperienceAnalyticsWorkFromAnywhereModelPerformance(value []UserExperienceAnalyticsWorkFromAnywhereModelPerformanceable)()
    SetUserPfxCertificates(value []UserPFXCertificateable)()
    SetVirtualEndpoint(value VirtualEndpointable)()
    SetWindowsAutopilotDeploymentProfiles(value []WindowsAutopilotDeploymentProfileable)()
    SetWindowsAutopilotDeviceIdentities(value []WindowsAutopilotDeviceIdentityable)()
    SetWindowsAutopilotSettings(value WindowsAutopilotSettingsable)()
    SetWindowsDriverUpdateProfiles(value []WindowsDriverUpdateProfileable)()
    SetWindowsFeatureUpdateProfiles(value []WindowsFeatureUpdateProfileable)()
    SetWindowsInformationProtectionAppLearningSummaries(value []WindowsInformationProtectionAppLearningSummaryable)()
    SetWindowsInformationProtectionNetworkLearningSummaries(value []WindowsInformationProtectionNetworkLearningSummaryable)()
    SetWindowsMalwareInformation(value []WindowsMalwareInformationable)()
    SetWindowsMalwareOverview(value WindowsMalwareOverviewable)()
    SetWindowsQualityUpdateProfiles(value []WindowsQualityUpdateProfileable)()
    SetWindowsUpdateCatalogItems(value []WindowsUpdateCatalogItemable)()
    SetZebraFotaArtifacts(value []ZebraFotaArtifactable)()
    SetZebraFotaConnector(value ZebraFotaConnectorable)()
    SetZebraFotaDeployments(value []ZebraFotaDeploymentable)()
}
