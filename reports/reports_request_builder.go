package reports

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ReportsRequestBuilder provides operations to manage the reportRoot singleton.
type ReportsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ReportsRequestBuilderGetQueryParameters get reports
type ReportsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ReportsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReportsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ReportsRequestBuilderGetQueryParameters
}
// ReportsRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReportsRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ApplicationSignInDetailedSummary provides operations to manage the applicationSignInDetailedSummary property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) ApplicationSignInDetailedSummary()(*ApplicationSignInDetailedSummaryRequestBuilder) {
    return NewApplicationSignInDetailedSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ApplicationSignInDetailedSummaryById provides operations to manage the applicationSignInDetailedSummary property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) ApplicationSignInDetailedSummaryById(id string)(*ApplicationSignInDetailedSummaryApplicationSignInDetailedSummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["applicationSignInDetailedSummary%2Did"] = id
    }
    return NewApplicationSignInDetailedSummaryApplicationSignInDetailedSummaryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AuthenticationMethods provides operations to manage the authenticationMethods property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) AuthenticationMethods()(*AuthenticationMethodsRequestBuilder) {
    return NewAuthenticationMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewReportsRequestBuilderInternal instantiates a new ReportsRequestBuilder and sets the default values.
func NewReportsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsRequestBuilder) {
    m := &ReportsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/reports{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewReportsRequestBuilder instantiates a new ReportsRequestBuilder and sets the default values.
func NewReportsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReportsRequestBuilderInternal(urlParams, requestAdapter)
}
// CredentialUserRegistrationDetails provides operations to manage the credentialUserRegistrationDetails property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) CredentialUserRegistrationDetails()(*CredentialUserRegistrationDetailsRequestBuilder) {
    return NewCredentialUserRegistrationDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CredentialUserRegistrationDetailsById provides operations to manage the credentialUserRegistrationDetails property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) CredentialUserRegistrationDetailsById(id string)(*CredentialUserRegistrationDetailsCredentialUserRegistrationDetailsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["credentialUserRegistrationDetails%2Did"] = id
    }
    return NewCredentialUserRegistrationDetailsCredentialUserRegistrationDetailsItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DailyPrintUsage provides operations to manage the dailyPrintUsage property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) DailyPrintUsage()(*DailyPrintUsageRequestBuilder) {
    return NewDailyPrintUsageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DailyPrintUsageById provides operations to manage the dailyPrintUsage property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) DailyPrintUsageById(id string)(*DailyPrintUsagePrintUsageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsage%2Did"] = id
    }
    return NewDailyPrintUsagePrintUsageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DailyPrintUsageByPrinter provides operations to manage the dailyPrintUsageByPrinter property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) DailyPrintUsageByPrinter()(*DailyPrintUsageByPrinterRequestBuilder) {
    return NewDailyPrintUsageByPrinterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DailyPrintUsageByPrinterById provides operations to manage the dailyPrintUsageByPrinter property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) DailyPrintUsageByPrinterById(id string)(*DailyPrintUsageByPrinterPrintUsageByPrinterItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByPrinter%2Did"] = id
    }
    return NewDailyPrintUsageByPrinterPrintUsageByPrinterItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DailyPrintUsageByUser provides operations to manage the dailyPrintUsageByUser property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) DailyPrintUsageByUser()(*DailyPrintUsageByUserRequestBuilder) {
    return NewDailyPrintUsageByUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DailyPrintUsageByUserById provides operations to manage the dailyPrintUsageByUser property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) DailyPrintUsageByUserById(id string)(*DailyPrintUsageByUserPrintUsageByUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByUser%2Did"] = id
    }
    return NewDailyPrintUsageByUserPrintUsageByUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DailyPrintUsageSummariesByPrinter provides operations to manage the dailyPrintUsageSummariesByPrinter property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) DailyPrintUsageSummariesByPrinter()(*DailyPrintUsageSummariesByPrinterRequestBuilder) {
    return NewDailyPrintUsageSummariesByPrinterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DailyPrintUsageSummariesByPrinterById provides operations to manage the dailyPrintUsageSummariesByPrinter property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) DailyPrintUsageSummariesByPrinterById(id string)(*DailyPrintUsageSummariesByPrinterPrintUsageByPrinterItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByPrinter%2Did"] = id
    }
    return NewDailyPrintUsageSummariesByPrinterPrintUsageByPrinterItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DailyPrintUsageSummariesByUser provides operations to manage the dailyPrintUsageSummariesByUser property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) DailyPrintUsageSummariesByUser()(*DailyPrintUsageSummariesByUserRequestBuilder) {
    return NewDailyPrintUsageSummariesByUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DailyPrintUsageSummariesByUserById provides operations to manage the dailyPrintUsageSummariesByUser property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) DailyPrintUsageSummariesByUserById(id string)(*DailyPrintUsageSummariesByUserPrintUsageByUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByUser%2Did"] = id
    }
    return NewDailyPrintUsageSummariesByUserPrintUsageByUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DeviceConfigurationDeviceActivity provides operations to call the deviceConfigurationDeviceActivity method.
func (m *ReportsRequestBuilder) DeviceConfigurationDeviceActivity()(*DeviceConfigurationDeviceActivityRequestBuilder) {
    return NewDeviceConfigurationDeviceActivityRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceConfigurationUserActivity provides operations to call the deviceConfigurationUserActivity method.
func (m *ReportsRequestBuilder) DeviceConfigurationUserActivity()(*DeviceConfigurationUserActivityRequestBuilder) {
    return NewDeviceConfigurationUserActivityRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get reports
func (m *ReportsRequestBuilder) Get(ctx context.Context, requestConfiguration *ReportsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReportRootable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateReportRootFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReportRootable), nil
}
// GetAttackSimulationRepeatOffenders provides operations to call the getAttackSimulationRepeatOffenders method.
func (m *ReportsRequestBuilder) GetAttackSimulationRepeatOffenders()(*GetAttackSimulationRepeatOffendersRequestBuilder) {
    return NewGetAttackSimulationRepeatOffendersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetAttackSimulationSimulationUserCoverage provides operations to call the getAttackSimulationSimulationUserCoverage method.
func (m *ReportsRequestBuilder) GetAttackSimulationSimulationUserCoverage()(*GetAttackSimulationSimulationUserCoverageRequestBuilder) {
    return NewGetAttackSimulationSimulationUserCoverageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetAttackSimulationTrainingUserCoverage provides operations to call the getAttackSimulationTrainingUserCoverage method.
func (m *ReportsRequestBuilder) GetAttackSimulationTrainingUserCoverage()(*GetAttackSimulationTrainingUserCoverageRequestBuilder) {
    return NewGetAttackSimulationTrainingUserCoverageRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetAzureADApplicationSignInSummaryWithPeriod provides operations to call the getAzureADApplicationSignInSummary method.
func (m *ReportsRequestBuilder) GetAzureADApplicationSignInSummaryWithPeriod(period *string)(*GetAzureADApplicationSignInSummaryWithPeriodRequestBuilder) {
    return NewGetAzureADApplicationSignInSummaryWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetBrowserDistributionUserCountsWithPeriod provides operations to call the getBrowserDistributionUserCounts method.
func (m *ReportsRequestBuilder) GetBrowserDistributionUserCountsWithPeriod(period *string)(*GetBrowserDistributionUserCountsWithPeriodRequestBuilder) {
    return NewGetBrowserDistributionUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetBrowserUserCountsWithPeriod provides operations to call the getBrowserUserCounts method.
func (m *ReportsRequestBuilder) GetBrowserUserCountsWithPeriod(period *string)(*GetBrowserUserCountsWithPeriodRequestBuilder) {
    return NewGetBrowserUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetBrowserUserDetailWithPeriod provides operations to call the getBrowserUserDetail method.
func (m *ReportsRequestBuilder) GetBrowserUserDetailWithPeriod(period *string)(*GetBrowserUserDetailWithPeriodRequestBuilder) {
    return NewGetBrowserUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetCredentialUsageSummaryWithPeriod provides operations to call the getCredentialUsageSummary method.
func (m *ReportsRequestBuilder) GetCredentialUsageSummaryWithPeriod(period *string)(*GetCredentialUsageSummaryWithPeriodRequestBuilder) {
    return NewGetCredentialUsageSummaryWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetCredentialUserRegistrationCount provides operations to call the getCredentialUserRegistrationCount method.
func (m *ReportsRequestBuilder) GetCredentialUserRegistrationCount()(*GetCredentialUserRegistrationCountRequestBuilder) {
    return NewGetCredentialUserRegistrationCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetEmailActivityCountsWithPeriod provides operations to call the getEmailActivityCounts method.
func (m *ReportsRequestBuilder) GetEmailActivityCountsWithPeriod(period *string)(*GetEmailActivityCountsWithPeriodRequestBuilder) {
    return NewGetEmailActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetEmailActivityUserCountsWithPeriod provides operations to call the getEmailActivityUserCounts method.
func (m *ReportsRequestBuilder) GetEmailActivityUserCountsWithPeriod(period *string)(*GetEmailActivityUserCountsWithPeriodRequestBuilder) {
    return NewGetEmailActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetEmailActivityUserDetailWithDate provides operations to call the getEmailActivityUserDetail method.
func (m *ReportsRequestBuilder) GetEmailActivityUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*GetEmailActivityUserDetailWithDateRequestBuilder) {
    return NewGetEmailActivityUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetEmailActivityUserDetailWithPeriod provides operations to call the getEmailActivityUserDetail method.
func (m *ReportsRequestBuilder) GetEmailActivityUserDetailWithPeriod(period *string)(*GetEmailActivityUserDetailWithPeriodRequestBuilder) {
    return NewGetEmailActivityUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetEmailAppUsageAppsUserCountsWithPeriod provides operations to call the getEmailAppUsageAppsUserCounts method.
func (m *ReportsRequestBuilder) GetEmailAppUsageAppsUserCountsWithPeriod(period *string)(*GetEmailAppUsageAppsUserCountsWithPeriodRequestBuilder) {
    return NewGetEmailAppUsageAppsUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetEmailAppUsageUserCountsWithPeriod provides operations to call the getEmailAppUsageUserCounts method.
func (m *ReportsRequestBuilder) GetEmailAppUsageUserCountsWithPeriod(period *string)(*GetEmailAppUsageUserCountsWithPeriodRequestBuilder) {
    return NewGetEmailAppUsageUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetEmailAppUsageUserDetailWithDate provides operations to call the getEmailAppUsageUserDetail method.
func (m *ReportsRequestBuilder) GetEmailAppUsageUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*GetEmailAppUsageUserDetailWithDateRequestBuilder) {
    return NewGetEmailAppUsageUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetEmailAppUsageUserDetailWithPeriod provides operations to call the getEmailAppUsageUserDetail method.
func (m *ReportsRequestBuilder) GetEmailAppUsageUserDetailWithPeriod(period *string)(*GetEmailAppUsageUserDetailWithPeriodRequestBuilder) {
    return NewGetEmailAppUsageUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetEmailAppUsageVersionsUserCountsWithPeriod provides operations to call the getEmailAppUsageVersionsUserCounts method.
func (m *ReportsRequestBuilder) GetEmailAppUsageVersionsUserCountsWithPeriod(period *string)(*GetEmailAppUsageVersionsUserCountsWithPeriodRequestBuilder) {
    return NewGetEmailAppUsageVersionsUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTime provides operations to call the getGroupArchivedPrintJobs method.
func (m *ReportsRequestBuilder) GetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTime(endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, groupId *string, startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)(*GetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return NewGetGroupArchivedPrintJobsWithGroupIdWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, groupId, startDateTime);
}
// GetM365AppPlatformUserCountsWithPeriod provides operations to call the getM365AppPlatformUserCounts method.
func (m *ReportsRequestBuilder) GetM365AppPlatformUserCountsWithPeriod(period *string)(*GetM365AppPlatformUserCountsWithPeriodRequestBuilder) {
    return NewGetM365AppPlatformUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetM365AppUserCountsWithPeriod provides operations to call the getM365AppUserCounts method.
func (m *ReportsRequestBuilder) GetM365AppUserCountsWithPeriod(period *string)(*GetM365AppUserCountsWithPeriodRequestBuilder) {
    return NewGetM365AppUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetM365AppUserDetailWithDate provides operations to call the getM365AppUserDetail method.
func (m *ReportsRequestBuilder) GetM365AppUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*GetM365AppUserDetailWithDateRequestBuilder) {
    return NewGetM365AppUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetM365AppUserDetailWithPeriod provides operations to call the getM365AppUserDetail method.
func (m *ReportsRequestBuilder) GetM365AppUserDetailWithPeriod(period *string)(*GetM365AppUserDetailWithPeriodRequestBuilder) {
    return NewGetM365AppUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetMailboxUsageDetailWithPeriod provides operations to call the getMailboxUsageDetail method.
func (m *ReportsRequestBuilder) GetMailboxUsageDetailWithPeriod(period *string)(*GetMailboxUsageDetailWithPeriodRequestBuilder) {
    return NewGetMailboxUsageDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetMailboxUsageMailboxCountsWithPeriod provides operations to call the getMailboxUsageMailboxCounts method.
func (m *ReportsRequestBuilder) GetMailboxUsageMailboxCountsWithPeriod(period *string)(*GetMailboxUsageMailboxCountsWithPeriodRequestBuilder) {
    return NewGetMailboxUsageMailboxCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetMailboxUsageQuotaStatusMailboxCountsWithPeriod provides operations to call the getMailboxUsageQuotaStatusMailboxCounts method.
func (m *ReportsRequestBuilder) GetMailboxUsageQuotaStatusMailboxCountsWithPeriod(period *string)(*GetMailboxUsageQuotaStatusMailboxCountsWithPeriodRequestBuilder) {
    return NewGetMailboxUsageQuotaStatusMailboxCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetMailboxUsageStorageWithPeriod provides operations to call the getMailboxUsageStorage method.
func (m *ReportsRequestBuilder) GetMailboxUsageStorageWithPeriod(period *string)(*GetMailboxUsageStorageWithPeriodRequestBuilder) {
    return NewGetMailboxUsageStorageWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOffice365ActivationCounts provides operations to call the getOffice365ActivationCounts method.
func (m *ReportsRequestBuilder) GetOffice365ActivationCounts()(*GetOffice365ActivationCountsRequestBuilder) {
    return NewGetOffice365ActivationCountsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetOffice365ActivationsUserCounts provides operations to call the getOffice365ActivationsUserCounts method.
func (m *ReportsRequestBuilder) GetOffice365ActivationsUserCounts()(*GetOffice365ActivationsUserCountsRequestBuilder) {
    return NewGetOffice365ActivationsUserCountsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetOffice365ActivationsUserDetail provides operations to call the getOffice365ActivationsUserDetail method.
func (m *ReportsRequestBuilder) GetOffice365ActivationsUserDetail()(*GetOffice365ActivationsUserDetailRequestBuilder) {
    return NewGetOffice365ActivationsUserDetailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetOffice365ActiveUserCountsWithPeriod provides operations to call the getOffice365ActiveUserCounts method.
func (m *ReportsRequestBuilder) GetOffice365ActiveUserCountsWithPeriod(period *string)(*GetOffice365ActiveUserCountsWithPeriodRequestBuilder) {
    return NewGetOffice365ActiveUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOffice365ActiveUserDetailWithDate provides operations to call the getOffice365ActiveUserDetail method.
func (m *ReportsRequestBuilder) GetOffice365ActiveUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*GetOffice365ActiveUserDetailWithDateRequestBuilder) {
    return NewGetOffice365ActiveUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetOffice365ActiveUserDetailWithPeriod provides operations to call the getOffice365ActiveUserDetail method.
func (m *ReportsRequestBuilder) GetOffice365ActiveUserDetailWithPeriod(period *string)(*GetOffice365ActiveUserDetailWithPeriodRequestBuilder) {
    return NewGetOffice365ActiveUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOffice365GroupsActivityCountsWithPeriod provides operations to call the getOffice365GroupsActivityCounts method.
func (m *ReportsRequestBuilder) GetOffice365GroupsActivityCountsWithPeriod(period *string)(*GetOffice365GroupsActivityCountsWithPeriodRequestBuilder) {
    return NewGetOffice365GroupsActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOffice365GroupsActivityDetailWithDate provides operations to call the getOffice365GroupsActivityDetail method.
func (m *ReportsRequestBuilder) GetOffice365GroupsActivityDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*GetOffice365GroupsActivityDetailWithDateRequestBuilder) {
    return NewGetOffice365GroupsActivityDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetOffice365GroupsActivityDetailWithPeriod provides operations to call the getOffice365GroupsActivityDetail method.
func (m *ReportsRequestBuilder) GetOffice365GroupsActivityDetailWithPeriod(period *string)(*GetOffice365GroupsActivityDetailWithPeriodRequestBuilder) {
    return NewGetOffice365GroupsActivityDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOffice365GroupsActivityFileCountsWithPeriod provides operations to call the getOffice365GroupsActivityFileCounts method.
func (m *ReportsRequestBuilder) GetOffice365GroupsActivityFileCountsWithPeriod(period *string)(*GetOffice365GroupsActivityFileCountsWithPeriodRequestBuilder) {
    return NewGetOffice365GroupsActivityFileCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOffice365GroupsActivityGroupCountsWithPeriod provides operations to call the getOffice365GroupsActivityGroupCounts method.
func (m *ReportsRequestBuilder) GetOffice365GroupsActivityGroupCountsWithPeriod(period *string)(*GetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilder) {
    return NewGetOffice365GroupsActivityGroupCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOffice365GroupsActivityStorageWithPeriod provides operations to call the getOffice365GroupsActivityStorage method.
func (m *ReportsRequestBuilder) GetOffice365GroupsActivityStorageWithPeriod(period *string)(*GetOffice365GroupsActivityStorageWithPeriodRequestBuilder) {
    return NewGetOffice365GroupsActivityStorageWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOffice365ServicesUserCountsWithPeriod provides operations to call the getOffice365ServicesUserCounts method.
func (m *ReportsRequestBuilder) GetOffice365ServicesUserCountsWithPeriod(period *string)(*GetOffice365ServicesUserCountsWithPeriodRequestBuilder) {
    return NewGetOffice365ServicesUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOneDriveActivityFileCountsWithPeriod provides operations to call the getOneDriveActivityFileCounts method.
func (m *ReportsRequestBuilder) GetOneDriveActivityFileCountsWithPeriod(period *string)(*GetOneDriveActivityFileCountsWithPeriodRequestBuilder) {
    return NewGetOneDriveActivityFileCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOneDriveActivityUserCountsWithPeriod provides operations to call the getOneDriveActivityUserCounts method.
func (m *ReportsRequestBuilder) GetOneDriveActivityUserCountsWithPeriod(period *string)(*GetOneDriveActivityUserCountsWithPeriodRequestBuilder) {
    return NewGetOneDriveActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOneDriveActivityUserDetailWithDate provides operations to call the getOneDriveActivityUserDetail method.
func (m *ReportsRequestBuilder) GetOneDriveActivityUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*GetOneDriveActivityUserDetailWithDateRequestBuilder) {
    return NewGetOneDriveActivityUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetOneDriveActivityUserDetailWithPeriod provides operations to call the getOneDriveActivityUserDetail method.
func (m *ReportsRequestBuilder) GetOneDriveActivityUserDetailWithPeriod(period *string)(*GetOneDriveActivityUserDetailWithPeriodRequestBuilder) {
    return NewGetOneDriveActivityUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOneDriveUsageAccountCountsWithPeriod provides operations to call the getOneDriveUsageAccountCounts method.
func (m *ReportsRequestBuilder) GetOneDriveUsageAccountCountsWithPeriod(period *string)(*GetOneDriveUsageAccountCountsWithPeriodRequestBuilder) {
    return NewGetOneDriveUsageAccountCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOneDriveUsageAccountDetailWithDate provides operations to call the getOneDriveUsageAccountDetail method.
func (m *ReportsRequestBuilder) GetOneDriveUsageAccountDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*GetOneDriveUsageAccountDetailWithDateRequestBuilder) {
    return NewGetOneDriveUsageAccountDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetOneDriveUsageAccountDetailWithPeriod provides operations to call the getOneDriveUsageAccountDetail method.
func (m *ReportsRequestBuilder) GetOneDriveUsageAccountDetailWithPeriod(period *string)(*GetOneDriveUsageAccountDetailWithPeriodRequestBuilder) {
    return NewGetOneDriveUsageAccountDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOneDriveUsageFileCountsWithPeriod provides operations to call the getOneDriveUsageFileCounts method.
func (m *ReportsRequestBuilder) GetOneDriveUsageFileCountsWithPeriod(period *string)(*GetOneDriveUsageFileCountsWithPeriodRequestBuilder) {
    return NewGetOneDriveUsageFileCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetOneDriveUsageStorageWithPeriod provides operations to call the getOneDriveUsageStorage method.
func (m *ReportsRequestBuilder) GetOneDriveUsageStorageWithPeriod(period *string)(*GetOneDriveUsageStorageWithPeriodRequestBuilder) {
    return NewGetOneDriveUsageStorageWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTime provides operations to call the getPrinterArchivedPrintJobs method.
func (m *ReportsRequestBuilder) GetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTime(endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, printerId *string, startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)(*GetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return NewGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, printerId, startDateTime);
}
// GetRelyingPartyDetailedSummaryWithPeriod provides operations to call the getRelyingPartyDetailedSummary method.
func (m *ReportsRequestBuilder) GetRelyingPartyDetailedSummaryWithPeriod(period *string)(*GetRelyingPartyDetailedSummaryWithPeriodRequestBuilder) {
    return NewGetRelyingPartyDetailedSummaryWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSharePointActivityFileCountsWithPeriod provides operations to call the getSharePointActivityFileCounts method.
func (m *ReportsRequestBuilder) GetSharePointActivityFileCountsWithPeriod(period *string)(*GetSharePointActivityFileCountsWithPeriodRequestBuilder) {
    return NewGetSharePointActivityFileCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSharePointActivityPagesWithPeriod provides operations to call the getSharePointActivityPages method.
func (m *ReportsRequestBuilder) GetSharePointActivityPagesWithPeriod(period *string)(*GetSharePointActivityPagesWithPeriodRequestBuilder) {
    return NewGetSharePointActivityPagesWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSharePointActivityUserCountsWithPeriod provides operations to call the getSharePointActivityUserCounts method.
func (m *ReportsRequestBuilder) GetSharePointActivityUserCountsWithPeriod(period *string)(*GetSharePointActivityUserCountsWithPeriodRequestBuilder) {
    return NewGetSharePointActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSharePointActivityUserDetailWithDate provides operations to call the getSharePointActivityUserDetail method.
func (m *ReportsRequestBuilder) GetSharePointActivityUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*GetSharePointActivityUserDetailWithDateRequestBuilder) {
    return NewGetSharePointActivityUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetSharePointActivityUserDetailWithPeriod provides operations to call the getSharePointActivityUserDetail method.
func (m *ReportsRequestBuilder) GetSharePointActivityUserDetailWithPeriod(period *string)(*GetSharePointActivityUserDetailWithPeriodRequestBuilder) {
    return NewGetSharePointActivityUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSharePointSiteUsageDetailWithDate provides operations to call the getSharePointSiteUsageDetail method.
func (m *ReportsRequestBuilder) GetSharePointSiteUsageDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*GetSharePointSiteUsageDetailWithDateRequestBuilder) {
    return NewGetSharePointSiteUsageDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetSharePointSiteUsageDetailWithPeriod provides operations to call the getSharePointSiteUsageDetail method.
func (m *ReportsRequestBuilder) GetSharePointSiteUsageDetailWithPeriod(period *string)(*GetSharePointSiteUsageDetailWithPeriodRequestBuilder) {
    return NewGetSharePointSiteUsageDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSharePointSiteUsageFileCountsWithPeriod provides operations to call the getSharePointSiteUsageFileCounts method.
func (m *ReportsRequestBuilder) GetSharePointSiteUsageFileCountsWithPeriod(period *string)(*GetSharePointSiteUsageFileCountsWithPeriodRequestBuilder) {
    return NewGetSharePointSiteUsageFileCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSharePointSiteUsagePagesWithPeriod provides operations to call the getSharePointSiteUsagePages method.
func (m *ReportsRequestBuilder) GetSharePointSiteUsagePagesWithPeriod(period *string)(*GetSharePointSiteUsagePagesWithPeriodRequestBuilder) {
    return NewGetSharePointSiteUsagePagesWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSharePointSiteUsageSiteCountsWithPeriod provides operations to call the getSharePointSiteUsageSiteCounts method.
func (m *ReportsRequestBuilder) GetSharePointSiteUsageSiteCountsWithPeriod(period *string)(*GetSharePointSiteUsageSiteCountsWithPeriodRequestBuilder) {
    return NewGetSharePointSiteUsageSiteCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSharePointSiteUsageStorageWithPeriod provides operations to call the getSharePointSiteUsageStorage method.
func (m *ReportsRequestBuilder) GetSharePointSiteUsageStorageWithPeriod(period *string)(*GetSharePointSiteUsageStorageWithPeriodRequestBuilder) {
    return NewGetSharePointSiteUsageStorageWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessActivityCountsWithPeriod provides operations to call the getSkypeForBusinessActivityCounts method.
func (m *ReportsRequestBuilder) GetSkypeForBusinessActivityCountsWithPeriod(period *string)(*GetSkypeForBusinessActivityCountsWithPeriodRequestBuilder) {
    return NewGetSkypeForBusinessActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessActivityUserCountsWithPeriod provides operations to call the getSkypeForBusinessActivityUserCounts method.
func (m *ReportsRequestBuilder) GetSkypeForBusinessActivityUserCountsWithPeriod(period *string)(*GetSkypeForBusinessActivityUserCountsWithPeriodRequestBuilder) {
    return NewGetSkypeForBusinessActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessActivityUserDetailWithDate provides operations to call the getSkypeForBusinessActivityUserDetail method.
func (m *ReportsRequestBuilder) GetSkypeForBusinessActivityUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*GetSkypeForBusinessActivityUserDetailWithDateRequestBuilder) {
    return NewGetSkypeForBusinessActivityUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetSkypeForBusinessActivityUserDetailWithPeriod provides operations to call the getSkypeForBusinessActivityUserDetail method.
func (m *ReportsRequestBuilder) GetSkypeForBusinessActivityUserDetailWithPeriod(period *string)(*GetSkypeForBusinessActivityUserDetailWithPeriodRequestBuilder) {
    return NewGetSkypeForBusinessActivityUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriod provides operations to call the getSkypeForBusinessDeviceUsageDistributionUserCounts method.
func (m *ReportsRequestBuilder) GetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriod(period *string)(*GetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilder) {
    return NewGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessDeviceUsageUserCountsWithPeriod provides operations to call the getSkypeForBusinessDeviceUsageUserCounts method.
func (m *ReportsRequestBuilder) GetSkypeForBusinessDeviceUsageUserCountsWithPeriod(period *string)(*GetSkypeForBusinessDeviceUsageUserCountsWithPeriodRequestBuilder) {
    return NewGetSkypeForBusinessDeviceUsageUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessDeviceUsageUserDetailWithDate provides operations to call the getSkypeForBusinessDeviceUsageUserDetail method.
func (m *ReportsRequestBuilder) GetSkypeForBusinessDeviceUsageUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*GetSkypeForBusinessDeviceUsageUserDetailWithDateRequestBuilder) {
    return NewGetSkypeForBusinessDeviceUsageUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetSkypeForBusinessDeviceUsageUserDetailWithPeriod provides operations to call the getSkypeForBusinessDeviceUsageUserDetail method.
func (m *ReportsRequestBuilder) GetSkypeForBusinessDeviceUsageUserDetailWithPeriod(period *string)(*GetSkypeForBusinessDeviceUsageUserDetailWithPeriodRequestBuilder) {
    return NewGetSkypeForBusinessDeviceUsageUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessOrganizerActivityCountsWithPeriod provides operations to call the getSkypeForBusinessOrganizerActivityCounts method.
func (m *ReportsRequestBuilder) GetSkypeForBusinessOrganizerActivityCountsWithPeriod(period *string)(*GetSkypeForBusinessOrganizerActivityCountsWithPeriodRequestBuilder) {
    return NewGetSkypeForBusinessOrganizerActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessOrganizerActivityMinuteCountsWithPeriod provides operations to call the getSkypeForBusinessOrganizerActivityMinuteCounts method.
func (m *ReportsRequestBuilder) GetSkypeForBusinessOrganizerActivityMinuteCountsWithPeriod(period *string)(*GetSkypeForBusinessOrganizerActivityMinuteCountsWithPeriodRequestBuilder) {
    return NewGetSkypeForBusinessOrganizerActivityMinuteCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessOrganizerActivityUserCountsWithPeriod provides operations to call the getSkypeForBusinessOrganizerActivityUserCounts method.
func (m *ReportsRequestBuilder) GetSkypeForBusinessOrganizerActivityUserCountsWithPeriod(period *string)(*GetSkypeForBusinessOrganizerActivityUserCountsWithPeriodRequestBuilder) {
    return NewGetSkypeForBusinessOrganizerActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessParticipantActivityCountsWithPeriod provides operations to call the getSkypeForBusinessParticipantActivityCounts method.
func (m *ReportsRequestBuilder) GetSkypeForBusinessParticipantActivityCountsWithPeriod(period *string)(*GetSkypeForBusinessParticipantActivityCountsWithPeriodRequestBuilder) {
    return NewGetSkypeForBusinessParticipantActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessParticipantActivityMinuteCountsWithPeriod provides operations to call the getSkypeForBusinessParticipantActivityMinuteCounts method.
func (m *ReportsRequestBuilder) GetSkypeForBusinessParticipantActivityMinuteCountsWithPeriod(period *string)(*GetSkypeForBusinessParticipantActivityMinuteCountsWithPeriodRequestBuilder) {
    return NewGetSkypeForBusinessParticipantActivityMinuteCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessParticipantActivityUserCountsWithPeriod provides operations to call the getSkypeForBusinessParticipantActivityUserCounts method.
func (m *ReportsRequestBuilder) GetSkypeForBusinessParticipantActivityUserCountsWithPeriod(period *string)(*GetSkypeForBusinessParticipantActivityUserCountsWithPeriodRequestBuilder) {
    return NewGetSkypeForBusinessParticipantActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessPeerToPeerActivityCountsWithPeriod provides operations to call the getSkypeForBusinessPeerToPeerActivityCounts method.
func (m *ReportsRequestBuilder) GetSkypeForBusinessPeerToPeerActivityCountsWithPeriod(period *string)(*GetSkypeForBusinessPeerToPeerActivityCountsWithPeriodRequestBuilder) {
    return NewGetSkypeForBusinessPeerToPeerActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessPeerToPeerActivityMinuteCountsWithPeriod provides operations to call the getSkypeForBusinessPeerToPeerActivityMinuteCounts method.
func (m *ReportsRequestBuilder) GetSkypeForBusinessPeerToPeerActivityMinuteCountsWithPeriod(period *string)(*GetSkypeForBusinessPeerToPeerActivityMinuteCountsWithPeriodRequestBuilder) {
    return NewGetSkypeForBusinessPeerToPeerActivityMinuteCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetSkypeForBusinessPeerToPeerActivityUserCountsWithPeriod provides operations to call the getSkypeForBusinessPeerToPeerActivityUserCounts method.
func (m *ReportsRequestBuilder) GetSkypeForBusinessPeerToPeerActivityUserCountsWithPeriod(period *string)(*GetSkypeForBusinessPeerToPeerActivityUserCountsWithPeriodRequestBuilder) {
    return NewGetSkypeForBusinessPeerToPeerActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriod provides operations to call the getTeamsDeviceUsageDistributionTotalUserCounts method.
func (m *ReportsRequestBuilder) GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriod(period *string)(*GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriodRequestBuilder) {
    return NewGetTeamsDeviceUsageDistributionTotalUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsDeviceUsageDistributionUserCountsWithPeriod provides operations to call the getTeamsDeviceUsageDistributionUserCounts method.
func (m *ReportsRequestBuilder) GetTeamsDeviceUsageDistributionUserCountsWithPeriod(period *string)(*GetTeamsDeviceUsageDistributionUserCountsWithPeriodRequestBuilder) {
    return NewGetTeamsDeviceUsageDistributionUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsDeviceUsageTotalUserCountsWithPeriod provides operations to call the getTeamsDeviceUsageTotalUserCounts method.
func (m *ReportsRequestBuilder) GetTeamsDeviceUsageTotalUserCountsWithPeriod(period *string)(*GetTeamsDeviceUsageTotalUserCountsWithPeriodRequestBuilder) {
    return NewGetTeamsDeviceUsageTotalUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsDeviceUsageUserCountsWithPeriod provides operations to call the getTeamsDeviceUsageUserCounts method.
func (m *ReportsRequestBuilder) GetTeamsDeviceUsageUserCountsWithPeriod(period *string)(*GetTeamsDeviceUsageUserCountsWithPeriodRequestBuilder) {
    return NewGetTeamsDeviceUsageUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsDeviceUsageUserDetailWithDate provides operations to call the getTeamsDeviceUsageUserDetail method.
func (m *ReportsRequestBuilder) GetTeamsDeviceUsageUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*GetTeamsDeviceUsageUserDetailWithDateRequestBuilder) {
    return NewGetTeamsDeviceUsageUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetTeamsDeviceUsageUserDetailWithPeriod provides operations to call the getTeamsDeviceUsageUserDetail method.
func (m *ReportsRequestBuilder) GetTeamsDeviceUsageUserDetailWithPeriod(period *string)(*GetTeamsDeviceUsageUserDetailWithPeriodRequestBuilder) {
    return NewGetTeamsDeviceUsageUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsTeamActivityCountsWithPeriod provides operations to call the getTeamsTeamActivityCounts method.
func (m *ReportsRequestBuilder) GetTeamsTeamActivityCountsWithPeriod(period *string)(*GetTeamsTeamActivityCountsWithPeriodRequestBuilder) {
    return NewGetTeamsTeamActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsTeamActivityDetailWithDate provides operations to call the getTeamsTeamActivityDetail method.
func (m *ReportsRequestBuilder) GetTeamsTeamActivityDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*GetTeamsTeamActivityDetailWithDateRequestBuilder) {
    return NewGetTeamsTeamActivityDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetTeamsTeamActivityDetailWithPeriod provides operations to call the getTeamsTeamActivityDetail method.
func (m *ReportsRequestBuilder) GetTeamsTeamActivityDetailWithPeriod(period *string)(*GetTeamsTeamActivityDetailWithPeriodRequestBuilder) {
    return NewGetTeamsTeamActivityDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsTeamActivityDistributionCountsWithPeriod provides operations to call the getTeamsTeamActivityDistributionCounts method.
func (m *ReportsRequestBuilder) GetTeamsTeamActivityDistributionCountsWithPeriod(period *string)(*GetTeamsTeamActivityDistributionCountsWithPeriodRequestBuilder) {
    return NewGetTeamsTeamActivityDistributionCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsTeamCountsWithPeriod provides operations to call the getTeamsTeamCounts method.
func (m *ReportsRequestBuilder) GetTeamsTeamCountsWithPeriod(period *string)(*GetTeamsTeamCountsWithPeriodRequestBuilder) {
    return NewGetTeamsTeamCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsUserActivityCountsWithPeriod provides operations to call the getTeamsUserActivityCounts method.
func (m *ReportsRequestBuilder) GetTeamsUserActivityCountsWithPeriod(period *string)(*GetTeamsUserActivityCountsWithPeriodRequestBuilder) {
    return NewGetTeamsUserActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsUserActivityDistributionTotalUserCountsWithPeriod provides operations to call the getTeamsUserActivityDistributionTotalUserCounts method.
func (m *ReportsRequestBuilder) GetTeamsUserActivityDistributionTotalUserCountsWithPeriod(period *string)(*GetTeamsUserActivityDistributionTotalUserCountsWithPeriodRequestBuilder) {
    return NewGetTeamsUserActivityDistributionTotalUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsUserActivityDistributionUserCountsWithPeriod provides operations to call the getTeamsUserActivityDistributionUserCounts method.
func (m *ReportsRequestBuilder) GetTeamsUserActivityDistributionUserCountsWithPeriod(period *string)(*GetTeamsUserActivityDistributionUserCountsWithPeriodRequestBuilder) {
    return NewGetTeamsUserActivityDistributionUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsUserActivityTotalCountsWithPeriod provides operations to call the getTeamsUserActivityTotalCounts method.
func (m *ReportsRequestBuilder) GetTeamsUserActivityTotalCountsWithPeriod(period *string)(*GetTeamsUserActivityTotalCountsWithPeriodRequestBuilder) {
    return NewGetTeamsUserActivityTotalCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsUserActivityTotalDistributionCountsWithPeriod provides operations to call the getTeamsUserActivityTotalDistributionCounts method.
func (m *ReportsRequestBuilder) GetTeamsUserActivityTotalDistributionCountsWithPeriod(period *string)(*GetTeamsUserActivityTotalDistributionCountsWithPeriodRequestBuilder) {
    return NewGetTeamsUserActivityTotalDistributionCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsUserActivityTotalUserCountsWithPeriod provides operations to call the getTeamsUserActivityTotalUserCounts method.
func (m *ReportsRequestBuilder) GetTeamsUserActivityTotalUserCountsWithPeriod(period *string)(*GetTeamsUserActivityTotalUserCountsWithPeriodRequestBuilder) {
    return NewGetTeamsUserActivityTotalUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsUserActivityUserCountsWithPeriod provides operations to call the getTeamsUserActivityUserCounts method.
func (m *ReportsRequestBuilder) GetTeamsUserActivityUserCountsWithPeriod(period *string)(*GetTeamsUserActivityUserCountsWithPeriodRequestBuilder) {
    return NewGetTeamsUserActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetTeamsUserActivityUserDetailWithDate provides operations to call the getTeamsUserActivityUserDetail method.
func (m *ReportsRequestBuilder) GetTeamsUserActivityUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*GetTeamsUserActivityUserDetailWithDateRequestBuilder) {
    return NewGetTeamsUserActivityUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetTeamsUserActivityUserDetailWithPeriod provides operations to call the getTeamsUserActivityUserDetail method.
func (m *ReportsRequestBuilder) GetTeamsUserActivityUserDetailWithPeriod(period *string)(*GetTeamsUserActivityUserDetailWithPeriodRequestBuilder) {
    return NewGetTeamsUserActivityUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTime provides operations to call the getUserArchivedPrintJobs method.
func (m *ReportsRequestBuilder) GetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTime(endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time, userId *string)(*GetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilder) {
    return NewGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTimeRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, startDateTime, userId);
}
// GetYammerActivityCountsWithPeriod provides operations to call the getYammerActivityCounts method.
func (m *ReportsRequestBuilder) GetYammerActivityCountsWithPeriod(period *string)(*GetYammerActivityCountsWithPeriodRequestBuilder) {
    return NewGetYammerActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetYammerActivityUserCountsWithPeriod provides operations to call the getYammerActivityUserCounts method.
func (m *ReportsRequestBuilder) GetYammerActivityUserCountsWithPeriod(period *string)(*GetYammerActivityUserCountsWithPeriodRequestBuilder) {
    return NewGetYammerActivityUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetYammerActivityUserDetailWithDate provides operations to call the getYammerActivityUserDetail method.
func (m *ReportsRequestBuilder) GetYammerActivityUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*GetYammerActivityUserDetailWithDateRequestBuilder) {
    return NewGetYammerActivityUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetYammerActivityUserDetailWithPeriod provides operations to call the getYammerActivityUserDetail method.
func (m *ReportsRequestBuilder) GetYammerActivityUserDetailWithPeriod(period *string)(*GetYammerActivityUserDetailWithPeriodRequestBuilder) {
    return NewGetYammerActivityUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetYammerDeviceUsageDistributionUserCountsWithPeriod provides operations to call the getYammerDeviceUsageDistributionUserCounts method.
func (m *ReportsRequestBuilder) GetYammerDeviceUsageDistributionUserCountsWithPeriod(period *string)(*GetYammerDeviceUsageDistributionUserCountsWithPeriodRequestBuilder) {
    return NewGetYammerDeviceUsageDistributionUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetYammerDeviceUsageUserCountsWithPeriod provides operations to call the getYammerDeviceUsageUserCounts method.
func (m *ReportsRequestBuilder) GetYammerDeviceUsageUserCountsWithPeriod(period *string)(*GetYammerDeviceUsageUserCountsWithPeriodRequestBuilder) {
    return NewGetYammerDeviceUsageUserCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetYammerDeviceUsageUserDetailWithDate provides operations to call the getYammerDeviceUsageUserDetail method.
func (m *ReportsRequestBuilder) GetYammerDeviceUsageUserDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*GetYammerDeviceUsageUserDetailWithDateRequestBuilder) {
    return NewGetYammerDeviceUsageUserDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetYammerDeviceUsageUserDetailWithPeriod provides operations to call the getYammerDeviceUsageUserDetail method.
func (m *ReportsRequestBuilder) GetYammerDeviceUsageUserDetailWithPeriod(period *string)(*GetYammerDeviceUsageUserDetailWithPeriodRequestBuilder) {
    return NewGetYammerDeviceUsageUserDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetYammerGroupsActivityCountsWithPeriod provides operations to call the getYammerGroupsActivityCounts method.
func (m *ReportsRequestBuilder) GetYammerGroupsActivityCountsWithPeriod(period *string)(*GetYammerGroupsActivityCountsWithPeriodRequestBuilder) {
    return NewGetYammerGroupsActivityCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetYammerGroupsActivityDetailWithDate provides operations to call the getYammerGroupsActivityDetail method.
func (m *ReportsRequestBuilder) GetYammerGroupsActivityDetailWithDate(date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*GetYammerGroupsActivityDetailWithDateRequestBuilder) {
    return NewGetYammerGroupsActivityDetailWithDateRequestBuilderInternal(m.pathParameters, m.requestAdapter, date);
}
// GetYammerGroupsActivityDetailWithPeriod provides operations to call the getYammerGroupsActivityDetail method.
func (m *ReportsRequestBuilder) GetYammerGroupsActivityDetailWithPeriod(period *string)(*GetYammerGroupsActivityDetailWithPeriodRequestBuilder) {
    return NewGetYammerGroupsActivityDetailWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// GetYammerGroupsActivityGroupCountsWithPeriod provides operations to call the getYammerGroupsActivityGroupCounts method.
func (m *ReportsRequestBuilder) GetYammerGroupsActivityGroupCountsWithPeriod(period *string)(*GetYammerGroupsActivityGroupCountsWithPeriodRequestBuilder) {
    return NewGetYammerGroupsActivityGroupCountsWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// ManagedDeviceEnrollmentAbandonmentDetailsWithSkipWithTopWithFilterWithSkipToken provides operations to call the managedDeviceEnrollmentAbandonmentDetails method.
func (m *ReportsRequestBuilder) ManagedDeviceEnrollmentAbandonmentDetailsWithSkipWithTopWithFilterWithSkipToken(filter *string, skip *int32, skipToken *string, top *int32)(*ManagedDeviceEnrollmentAbandonmentDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilder) {
    return NewManagedDeviceEnrollmentAbandonmentDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter, filter, skip, skipToken, top);
}
// ManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipToken provides operations to call the managedDeviceEnrollmentAbandonmentSummary method.
func (m *ReportsRequestBuilder) ManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipToken(filter *string, skip *int32, skipToken *string, top *int32)(*ManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipTokenRequestBuilder) {
    return NewManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter, filter, skip, skipToken, top);
}
// ManagedDeviceEnrollmentFailureDetails provides operations to call the managedDeviceEnrollmentFailureDetails method.
func (m *ReportsRequestBuilder) ManagedDeviceEnrollmentFailureDetails()(*ManagedDeviceEnrollmentFailureDetailsRequestBuilder) {
    return NewManagedDeviceEnrollmentFailureDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedDeviceEnrollmentFailureDetailsWithSkipWithTopWithFilterWithSkipToken provides operations to call the managedDeviceEnrollmentFailureDetails method.
func (m *ReportsRequestBuilder) ManagedDeviceEnrollmentFailureDetailsWithSkipWithTopWithFilterWithSkipToken(filter *string, skip *int32, skipToken *string, top *int32)(*ManagedDeviceEnrollmentFailureDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilder) {
    return NewManagedDeviceEnrollmentFailureDetailsWithSkipWithTopWithFilterWithSkipTokenRequestBuilderInternal(m.pathParameters, m.requestAdapter, filter, skip, skipToken, top);
}
// ManagedDeviceEnrollmentFailureTrends provides operations to call the managedDeviceEnrollmentFailureTrends method.
func (m *ReportsRequestBuilder) ManagedDeviceEnrollmentFailureTrends()(*ManagedDeviceEnrollmentFailureTrendsRequestBuilder) {
    return NewManagedDeviceEnrollmentFailureTrendsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedDeviceEnrollmentTopFailures provides operations to call the managedDeviceEnrollmentTopFailures method.
func (m *ReportsRequestBuilder) ManagedDeviceEnrollmentTopFailures()(*ManagedDeviceEnrollmentTopFailuresRequestBuilder) {
    return NewManagedDeviceEnrollmentTopFailuresRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedDeviceEnrollmentTopFailuresWithPeriod provides operations to call the managedDeviceEnrollmentTopFailures method.
func (m *ReportsRequestBuilder) ManagedDeviceEnrollmentTopFailuresWithPeriod(period *string)(*ManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilder) {
    return NewManagedDeviceEnrollmentTopFailuresWithPeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter, period);
}
// MonthlyPrintUsageByPrinter provides operations to manage the monthlyPrintUsageByPrinter property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) MonthlyPrintUsageByPrinter()(*MonthlyPrintUsageByPrinterRequestBuilder) {
    return NewMonthlyPrintUsageByPrinterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MonthlyPrintUsageByPrinterById provides operations to manage the monthlyPrintUsageByPrinter property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) MonthlyPrintUsageByPrinterById(id string)(*MonthlyPrintUsageByPrinterPrintUsageByPrinterItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByPrinter%2Did"] = id
    }
    return NewMonthlyPrintUsageByPrinterPrintUsageByPrinterItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MonthlyPrintUsageByUser provides operations to manage the monthlyPrintUsageByUser property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) MonthlyPrintUsageByUser()(*MonthlyPrintUsageByUserRequestBuilder) {
    return NewMonthlyPrintUsageByUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MonthlyPrintUsageByUserById provides operations to manage the monthlyPrintUsageByUser property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) MonthlyPrintUsageByUserById(id string)(*MonthlyPrintUsageByUserPrintUsageByUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByUser%2Did"] = id
    }
    return NewMonthlyPrintUsageByUserPrintUsageByUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MonthlyPrintUsageSummariesByPrinter provides operations to manage the monthlyPrintUsageSummariesByPrinter property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) MonthlyPrintUsageSummariesByPrinter()(*MonthlyPrintUsageSummariesByPrinterRequestBuilder) {
    return NewMonthlyPrintUsageSummariesByPrinterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MonthlyPrintUsageSummariesByPrinterById provides operations to manage the monthlyPrintUsageSummariesByPrinter property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) MonthlyPrintUsageSummariesByPrinterById(id string)(*MonthlyPrintUsageSummariesByPrinterPrintUsageByPrinterItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByPrinter%2Did"] = id
    }
    return NewMonthlyPrintUsageSummariesByPrinterPrintUsageByPrinterItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MonthlyPrintUsageSummariesByUser provides operations to manage the monthlyPrintUsageSummariesByUser property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) MonthlyPrintUsageSummariesByUser()(*MonthlyPrintUsageSummariesByUserRequestBuilder) {
    return NewMonthlyPrintUsageSummariesByUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MonthlyPrintUsageSummariesByUserById provides operations to manage the monthlyPrintUsageSummariesByUser property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) MonthlyPrintUsageSummariesByUserById(id string)(*MonthlyPrintUsageSummariesByUserPrintUsageByUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByUser%2Did"] = id
    }
    return NewMonthlyPrintUsageSummariesByUserPrintUsageByUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update reports
func (m *ReportsRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReportRootable, requestConfiguration *ReportsRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReportRootable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateReportRootFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReportRootable), nil
}
// Security provides operations to manage the security property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) Security()(*SecurityRequestBuilder) {
    return NewSecurityRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ToGetRequestInformation get reports
func (m *ReportsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ReportsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPatchRequestInformation update reports
func (m *ReportsRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ReportRootable, requestConfiguration *ReportsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// UserCredentialUsageDetails provides operations to manage the userCredentialUsageDetails property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) UserCredentialUsageDetails()(*UserCredentialUsageDetailsRequestBuilder) {
    return NewUserCredentialUsageDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserCredentialUsageDetailsById provides operations to manage the userCredentialUsageDetails property of the microsoft.graph.reportRoot entity.
func (m *ReportsRequestBuilder) UserCredentialUsageDetailsById(id string)(*UserCredentialUsageDetailsUserCredentialUsageDetailsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userCredentialUsageDetails%2Did"] = id
    }
    return NewUserCredentialUsageDetailsUserCredentialUsageDetailsItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
