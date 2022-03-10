package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Userable 
type Userable interface {
    DirectoryObjectable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAboutMe()(*string)
    GetAccountEnabled()(*bool)
    GetActivities()([]UserActivityable)
    GetAgeGroup()(*string)
    GetAgreementAcceptances()([]AgreementAcceptanceable)
    GetAnalytics()(UserAnalyticsable)
    GetAppConsentRequestsForApproval()([]AppConsentRequestable)
    GetAppRoleAssignments()([]AppRoleAssignmentable)
    GetApprovals()([]Approvalable)
    GetAssignedLicenses()([]AssignedLicenseable)
    GetAssignedPlans()([]AssignedPlanable)
    GetAuthentication()(Authenticationable)
    GetBirthday()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetBusinessPhones()([]string)
    GetCalendar()(Calendarable)
    GetCalendarGroups()([]CalendarGroupable)
    GetCalendars()([]Calendarable)
    GetCalendarView()([]Eventable)
    GetChats()([]Chatable)
    GetCity()(*string)
    GetCompanyName()(*string)
    GetConsentProvidedForMinor()(*string)
    GetContactFolders()([]ContactFolderable)
    GetContacts()([]Contactable)
    GetCountry()(*string)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCreatedObjects()([]DirectoryObjectable)
    GetCreationType()(*string)
    GetCustomSecurityAttributes()(CustomSecurityAttributeValueable)
    GetDepartment()(*string)
    GetDeviceEnrollmentConfigurations()([]DeviceEnrollmentConfigurationable)
    GetDeviceEnrollmentLimit()(*int32)
    GetDeviceKeys()([]DeviceKeyable)
    GetDeviceManagementTroubleshootingEvents()([]DeviceManagementTroubleshootingEventable)
    GetDevices()([]Deviceable)
    GetDirectReports()([]DirectoryObjectable)
    GetDisplayName()(*string)
    GetDrive()(Driveable)
    GetDrives()([]Driveable)
    GetEmployeeHireDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetEmployeeId()(*string)
    GetEmployeeOrgData()(EmployeeOrgDataable)
    GetEmployeeType()(*string)
    GetEvents()([]Eventable)
    GetExtensions()([]Extensionable)
    GetExternalUserState()(*string)
    GetExternalUserStateChangeDateTime()(*string)
    GetFaxNumber()(*string)
    GetFollowedSites()([]Siteable)
    GetGivenName()(*string)
    GetHireDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetIdentities()([]ObjectIdentityable)
    GetImAddresses()([]string)
    GetInferenceClassification()(InferenceClassificationable)
    GetInfoCatalogs()([]string)
    GetInformationProtection()(InformationProtectionable)
    GetInsights()(ItemInsightsable)
    GetInterests()([]string)
    GetIsResourceAccount()(*bool)
    GetJobTitle()(*string)
    GetJoinedGroups()([]Groupable)
    GetJoinedTeams()([]Teamable)
    GetLastPasswordChangeDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLegalAgeGroupClassification()(*string)
    GetLicenseAssignmentStates()([]LicenseAssignmentStateable)
    GetLicenseDetails()([]LicenseDetailsable)
    GetMail()(*string)
    GetMailboxSettings()(MailboxSettingsable)
    GetMailFolders()([]MailFolderable)
    GetMailNickname()(*string)
    GetManagedAppRegistrations()([]ManagedAppRegistrationable)
    GetManagedDevices()([]ManagedDeviceable)
    GetManager()(DirectoryObjectable)
    GetMemberOf()([]DirectoryObjectable)
    GetMessages()([]Messageable)
    GetMobileAppIntentAndStates()([]MobileAppIntentAndStateable)
    GetMobileAppTroubleshootingEvents()([]MobileAppTroubleshootingEventable)
    GetMobilePhone()(*string)
    GetMySite()(*string)
    GetNotifications()([]Notificationable)
    GetOauth2PermissionGrants()([]OAuth2PermissionGrantable)
    GetOfficeLocation()(*string)
    GetOnenote()(Onenoteable)
    GetOnlineMeetings()([]OnlineMeetingable)
    GetOnPremisesDistinguishedName()(*string)
    GetOnPremisesDomainName()(*string)
    GetOnPremisesExtensionAttributes()(OnPremisesExtensionAttributesable)
    GetOnPremisesImmutableId()(*string)
    GetOnPremisesLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetOnPremisesProvisioningErrors()([]OnPremisesProvisioningErrorable)
    GetOnPremisesSamAccountName()(*string)
    GetOnPremisesSecurityIdentifier()(*string)
    GetOnPremisesSyncEnabled()(*bool)
    GetOnPremisesUserPrincipalName()(*string)
    GetOtherMails()([]string)
    GetOutlook()(OutlookUserable)
    GetOwnedDevices()([]DirectoryObjectable)
    GetOwnedObjects()([]DirectoryObjectable)
    GetPasswordPolicies()(*string)
    GetPasswordProfile()(PasswordProfileable)
    GetPastProjects()([]string)
    GetPendingAccessReviewInstances()([]AccessReviewInstanceable)
    GetPeople()([]Personable)
    GetPhoto()(ProfilePhotoable)
    GetPhotos()([]ProfilePhotoable)
    GetPlanner()(PlannerUserable)
    GetPostalCode()(*string)
    GetPreferredDataLocation()(*string)
    GetPreferredLanguage()(*string)
    GetPreferredName()(*string)
    GetPresence()(Presenceable)
    GetProfile()(Profileable)
    GetProvisionedPlans()([]ProvisionedPlanable)
    GetProxyAddresses()([]string)
    GetRefreshTokensValidFromDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRegisteredDevices()([]DirectoryObjectable)
    GetResponsibilities()([]string)
    GetSchools()([]string)
    GetScopedRoleMemberOf()([]ScopedRoleMembershipable)
    GetSettings()(UserSettingsable)
    GetShowInAddressList()(*bool)
    GetSignInActivity()(SignInActivityable)
    GetSignInSessionsValidFromDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSkills()([]string)
    GetState()(*string)
    GetStreetAddress()(*string)
    GetSurname()(*string)
    GetTasks()(Tasksable)
    GetTeamwork()(UserTeamworkable)
    GetTodo()(Todoable)
    GetTransitiveMemberOf()([]DirectoryObjectable)
    GetTransitiveReports()([]DirectoryObjectable)
    GetUsageLocation()(*string)
    GetUsageRights()([]UsageRightable)
    GetUserPrincipalName()(*string)
    GetUserType()(*string)
    GetWindowsInformationProtectionDeviceRegistrations()([]WindowsInformationProtectionDeviceRegistrationable)
    SetAboutMe(value *string)()
    SetAccountEnabled(value *bool)()
    SetActivities(value []UserActivityable)()
    SetAgeGroup(value *string)()
    SetAgreementAcceptances(value []AgreementAcceptanceable)()
    SetAnalytics(value UserAnalyticsable)()
    SetAppConsentRequestsForApproval(value []AppConsentRequestable)()
    SetAppRoleAssignments(value []AppRoleAssignmentable)()
    SetApprovals(value []Approvalable)()
    SetAssignedLicenses(value []AssignedLicenseable)()
    SetAssignedPlans(value []AssignedPlanable)()
    SetAuthentication(value Authenticationable)()
    SetBirthday(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetBusinessPhones(value []string)()
    SetCalendar(value Calendarable)()
    SetCalendarGroups(value []CalendarGroupable)()
    SetCalendars(value []Calendarable)()
    SetCalendarView(value []Eventable)()
    SetChats(value []Chatable)()
    SetCity(value *string)()
    SetCompanyName(value *string)()
    SetConsentProvidedForMinor(value *string)()
    SetContactFolders(value []ContactFolderable)()
    SetContacts(value []Contactable)()
    SetCountry(value *string)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCreatedObjects(value []DirectoryObjectable)()
    SetCreationType(value *string)()
    SetCustomSecurityAttributes(value CustomSecurityAttributeValueable)()
    SetDepartment(value *string)()
    SetDeviceEnrollmentConfigurations(value []DeviceEnrollmentConfigurationable)()
    SetDeviceEnrollmentLimit(value *int32)()
    SetDeviceKeys(value []DeviceKeyable)()
    SetDeviceManagementTroubleshootingEvents(value []DeviceManagementTroubleshootingEventable)()
    SetDevices(value []Deviceable)()
    SetDirectReports(value []DirectoryObjectable)()
    SetDisplayName(value *string)()
    SetDrive(value Driveable)()
    SetDrives(value []Driveable)()
    SetEmployeeHireDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetEmployeeId(value *string)()
    SetEmployeeOrgData(value EmployeeOrgDataable)()
    SetEmployeeType(value *string)()
    SetEvents(value []Eventable)()
    SetExtensions(value []Extensionable)()
    SetExternalUserState(value *string)()
    SetExternalUserStateChangeDateTime(value *string)()
    SetFaxNumber(value *string)()
    SetFollowedSites(value []Siteable)()
    SetGivenName(value *string)()
    SetHireDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetIdentities(value []ObjectIdentityable)()
    SetImAddresses(value []string)()
    SetInferenceClassification(value InferenceClassificationable)()
    SetInfoCatalogs(value []string)()
    SetInformationProtection(value InformationProtectionable)()
    SetInsights(value ItemInsightsable)()
    SetInterests(value []string)()
    SetIsResourceAccount(value *bool)()
    SetJobTitle(value *string)()
    SetJoinedGroups(value []Groupable)()
    SetJoinedTeams(value []Teamable)()
    SetLastPasswordChangeDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLegalAgeGroupClassification(value *string)()
    SetLicenseAssignmentStates(value []LicenseAssignmentStateable)()
    SetLicenseDetails(value []LicenseDetailsable)()
    SetMail(value *string)()
    SetMailboxSettings(value MailboxSettingsable)()
    SetMailFolders(value []MailFolderable)()
    SetMailNickname(value *string)()
    SetManagedAppRegistrations(value []ManagedAppRegistrationable)()
    SetManagedDevices(value []ManagedDeviceable)()
    SetManager(value DirectoryObjectable)()
    SetMemberOf(value []DirectoryObjectable)()
    SetMessages(value []Messageable)()
    SetMobileAppIntentAndStates(value []MobileAppIntentAndStateable)()
    SetMobileAppTroubleshootingEvents(value []MobileAppTroubleshootingEventable)()
    SetMobilePhone(value *string)()
    SetMySite(value *string)()
    SetNotifications(value []Notificationable)()
    SetOauth2PermissionGrants(value []OAuth2PermissionGrantable)()
    SetOfficeLocation(value *string)()
    SetOnenote(value Onenoteable)()
    SetOnlineMeetings(value []OnlineMeetingable)()
    SetOnPremisesDistinguishedName(value *string)()
    SetOnPremisesDomainName(value *string)()
    SetOnPremisesExtensionAttributes(value OnPremisesExtensionAttributesable)()
    SetOnPremisesImmutableId(value *string)()
    SetOnPremisesLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetOnPremisesProvisioningErrors(value []OnPremisesProvisioningErrorable)()
    SetOnPremisesSamAccountName(value *string)()
    SetOnPremisesSecurityIdentifier(value *string)()
    SetOnPremisesSyncEnabled(value *bool)()
    SetOnPremisesUserPrincipalName(value *string)()
    SetOtherMails(value []string)()
    SetOutlook(value OutlookUserable)()
    SetOwnedDevices(value []DirectoryObjectable)()
    SetOwnedObjects(value []DirectoryObjectable)()
    SetPasswordPolicies(value *string)()
    SetPasswordProfile(value PasswordProfileable)()
    SetPastProjects(value []string)()
    SetPendingAccessReviewInstances(value []AccessReviewInstanceable)()
    SetPeople(value []Personable)()
    SetPhoto(value ProfilePhotoable)()
    SetPhotos(value []ProfilePhotoable)()
    SetPlanner(value PlannerUserable)()
    SetPostalCode(value *string)()
    SetPreferredDataLocation(value *string)()
    SetPreferredLanguage(value *string)()
    SetPreferredName(value *string)()
    SetPresence(value Presenceable)()
    SetProfile(value Profileable)()
    SetProvisionedPlans(value []ProvisionedPlanable)()
    SetProxyAddresses(value []string)()
    SetRefreshTokensValidFromDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRegisteredDevices(value []DirectoryObjectable)()
    SetResponsibilities(value []string)()
    SetSchools(value []string)()
    SetScopedRoleMemberOf(value []ScopedRoleMembershipable)()
    SetSettings(value UserSettingsable)()
    SetShowInAddressList(value *bool)()
    SetSignInActivity(value SignInActivityable)()
    SetSignInSessionsValidFromDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSkills(value []string)()
    SetState(value *string)()
    SetStreetAddress(value *string)()
    SetSurname(value *string)()
    SetTasks(value Tasksable)()
    SetTeamwork(value UserTeamworkable)()
    SetTodo(value Todoable)()
    SetTransitiveMemberOf(value []DirectoryObjectable)()
    SetTransitiveReports(value []DirectoryObjectable)()
    SetUsageLocation(value *string)()
    SetUsageRights(value []UsageRightable)()
    SetUserPrincipalName(value *string)()
    SetUserType(value *string)()
    SetWindowsInformationProtectionDeviceRegistrations(value []WindowsInformationProtectionDeviceRegistrationable)()
}
