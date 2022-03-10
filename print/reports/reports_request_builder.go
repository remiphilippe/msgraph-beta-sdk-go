package reports

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i254c294d6b358b3854c3ee45f61d16c1237e0477952e917b034bbb0d8a6802b1 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/usercredentialusagedetails"
    i27c74777b0efbee67b16cbd84335b5228b74a547d6ffe5e338e3c1c6cc55c61b "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/monthlyprintusagesummariesbyprinter"
    i330d7d86bcb93d59ef14fc1257795533ce121643f4e8becd8884b57bc63e7b4b "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/monthlyprintusagebyprinter"
    i3772e364c43656cba18386bc364e131c7943e3a0b0461e3a162fb5eae513ac0b "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/dailyprintusagesummariesbyuser"
    i40dc3bbc1c58ea07c4c8b16942aaeca4da00eb3514241eef66edcd213ba084f7 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/monthlyprintusagesummariesbyuser"
    i46fd699154cb21987f377181c37a09acabdc0e13ef2aec22fc2f1352022502ba "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/dailyprintusagebyuser"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i5ba6893546e6aaea0f4a67b2dd033135639babe968232c6fcc5163ab253c7847 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/dailyprintusagesummariesbyprinter"
    i6532f3214bfaa5d2aac58a9d9ffe42fb56e5e2242405e98165287c807365bcda "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/monthlyprintusagebyuser"
    ia74d4f7cd8fb70eb99b3c94cccde51ec24c099b4b64a46531b40bf30619e07c4 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/authenticationmethods"
    ibf41f160c65ef1194ee9e34f2d2b3bbf34a32e8ceb9026adff36a62f7067a2ac "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/dailyprintusagebyprinter"
    icddea0ed347960de17929a877880af4ac0bf296fb68f5c0e35eebd457fc0734f "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/applicationsignindetailedsummary"
    ice018542cd139d6164abff37b45f7b6a1d3055a473b3658086e14293b59be16e "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/credentialuserregistrationdetails"
    i1453ad3549b1a27e47913c2d4a29d5625c5b5e04deb3a2a6f264f5cd24a2e8da "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/dailyprintusagesummariesbyprinter/item"
    i1702ca3f0c998e29512738841125456c0b4e81357e5f7ca9f4176443002d28d5 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/monthlyprintusagebyuser/item"
    i2268eb94e47ea5d81b666c3310292b49733c18b0196efe2957b868503a9da7ff "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/monthlyprintusagesummariesbyuser/item"
    i22a3d34b6ea91c582f9826d88dc039aae3e5f8f62146be9df6713c324659e396 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/credentialuserregistrationdetails/item"
    i3808d92a196a8e86ccfe0455bb6aaf56bf5956b11460f112fbc1dbcd8375551d "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/dailyprintusagebyprinter/item"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i4ca23ea1fafc22df57a5d1150889ab294f6d1582a79d145b961145b60dbf4c5b "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/dailyprintusagebyuser/item"
    i4f1c6062b781770b431d75b41deb0a2ac227468563ebac0623bef14f3efe351d "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/usercredentialusagedetails/item"
    i55a12b578425084e14afca840f9f8ad56ad8296a76d84139beb67b9619a01ae2 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/monthlyprintusagebyprinter/item"
    i6581ea85ca9063d302a537e4fdb5af97c20fd90311aa357f5482935e84bcd517 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/applicationsignindetailedsummary/item"
    ic1e71dc755b571267687f63b13312b758bc574de61144bfa50cd81af3b6d3626 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/monthlyprintusagesummariesbyprinter/item"
    iffede79819a5c0828784daecd91c1db9829729eb545a0c3b1948096142482a96 "github.com/microsoftgraph/msgraph-beta-sdk-go/print/reports/dailyprintusagesummariesbyuser/item"
)

// ReportsRequestBuilder provides operations to manage the reports property of the microsoft.graph.print entity.
type ReportsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ReportsRequestBuilderDeleteOptions options for Delete
type ReportsRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ReportsRequestBuilderGetOptions options for Get
type ReportsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ReportsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ReportsRequestBuilderGetQueryParameters get reports from print
type ReportsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ReportsRequestBuilderPatchOptions options for Patch
type ReportsRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ReportRootable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ReportsRequestBuilder) ApplicationSignInDetailedSummary()(*icddea0ed347960de17929a877880af4ac0bf296fb68f5c0e35eebd457fc0734f.ApplicationSignInDetailedSummaryRequestBuilder) {
    return icddea0ed347960de17929a877880af4ac0bf296fb68f5c0e35eebd457fc0734f.NewApplicationSignInDetailedSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ApplicationSignInDetailedSummaryById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.print.reports.applicationSignInDetailedSummary.item collection
func (m *ReportsRequestBuilder) ApplicationSignInDetailedSummaryById(id string)(*i6581ea85ca9063d302a537e4fdb5af97c20fd90311aa357f5482935e84bcd517.ApplicationSignInDetailedSummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["applicationSignInDetailedSummary_id"] = id
    }
    return i6581ea85ca9063d302a537e4fdb5af97c20fd90311aa357f5482935e84bcd517.NewApplicationSignInDetailedSummaryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ReportsRequestBuilder) AuthenticationMethods()(*ia74d4f7cd8fb70eb99b3c94cccde51ec24c099b4b64a46531b40bf30619e07c4.AuthenticationMethodsRequestBuilder) {
    return ia74d4f7cd8fb70eb99b3c94cccde51ec24c099b4b64a46531b40bf30619e07c4.NewAuthenticationMethodsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewReportsRequestBuilderInternal instantiates a new ReportsRequestBuilder and sets the default values.
func NewReportsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ReportsRequestBuilder) {
    m := &ReportsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/print/reports{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewReportsRequestBuilder instantiates a new ReportsRequestBuilder and sets the default values.
func NewReportsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ReportsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReportsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property reports for print
func (m *ReportsRequestBuilder) CreateDeleteRequestInformation(options *ReportsRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get reports from print
func (m *ReportsRequestBuilder) CreateGetRequestInformation(options *ReportsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property reports in print
func (m *ReportsRequestBuilder) CreatePatchRequestInformation(options *ReportsRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *ReportsRequestBuilder) CredentialUserRegistrationDetails()(*ice018542cd139d6164abff37b45f7b6a1d3055a473b3658086e14293b59be16e.CredentialUserRegistrationDetailsRequestBuilder) {
    return ice018542cd139d6164abff37b45f7b6a1d3055a473b3658086e14293b59be16e.NewCredentialUserRegistrationDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CredentialUserRegistrationDetailsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.print.reports.credentialUserRegistrationDetails.item collection
func (m *ReportsRequestBuilder) CredentialUserRegistrationDetailsById(id string)(*i22a3d34b6ea91c582f9826d88dc039aae3e5f8f62146be9df6713c324659e396.CredentialUserRegistrationDetailsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["credentialUserRegistrationDetails_id"] = id
    }
    return i22a3d34b6ea91c582f9826d88dc039aae3e5f8f62146be9df6713c324659e396.NewCredentialUserRegistrationDetailsItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ReportsRequestBuilder) DailyPrintUsageByPrinter()(*ibf41f160c65ef1194ee9e34f2d2b3bbf34a32e8ceb9026adff36a62f7067a2ac.DailyPrintUsageByPrinterRequestBuilder) {
    return ibf41f160c65ef1194ee9e34f2d2b3bbf34a32e8ceb9026adff36a62f7067a2ac.NewDailyPrintUsageByPrinterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DailyPrintUsageByPrinterById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.print.reports.dailyPrintUsageByPrinter.item collection
func (m *ReportsRequestBuilder) DailyPrintUsageByPrinterById(id string)(*i3808d92a196a8e86ccfe0455bb6aaf56bf5956b11460f112fbc1dbcd8375551d.PrintUsageByPrinterItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByPrinter_id"] = id
    }
    return i3808d92a196a8e86ccfe0455bb6aaf56bf5956b11460f112fbc1dbcd8375551d.NewPrintUsageByPrinterItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ReportsRequestBuilder) DailyPrintUsageByUser()(*i46fd699154cb21987f377181c37a09acabdc0e13ef2aec22fc2f1352022502ba.DailyPrintUsageByUserRequestBuilder) {
    return i46fd699154cb21987f377181c37a09acabdc0e13ef2aec22fc2f1352022502ba.NewDailyPrintUsageByUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DailyPrintUsageByUserById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.print.reports.dailyPrintUsageByUser.item collection
func (m *ReportsRequestBuilder) DailyPrintUsageByUserById(id string)(*i4ca23ea1fafc22df57a5d1150889ab294f6d1582a79d145b961145b60dbf4c5b.PrintUsageByUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByUser_id"] = id
    }
    return i4ca23ea1fafc22df57a5d1150889ab294f6d1582a79d145b961145b60dbf4c5b.NewPrintUsageByUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ReportsRequestBuilder) DailyPrintUsageSummariesByPrinter()(*i5ba6893546e6aaea0f4a67b2dd033135639babe968232c6fcc5163ab253c7847.DailyPrintUsageSummariesByPrinterRequestBuilder) {
    return i5ba6893546e6aaea0f4a67b2dd033135639babe968232c6fcc5163ab253c7847.NewDailyPrintUsageSummariesByPrinterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DailyPrintUsageSummariesByPrinterById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.print.reports.dailyPrintUsageSummariesByPrinter.item collection
func (m *ReportsRequestBuilder) DailyPrintUsageSummariesByPrinterById(id string)(*i1453ad3549b1a27e47913c2d4a29d5625c5b5e04deb3a2a6f264f5cd24a2e8da.PrintUsageByPrinterItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByPrinter_id"] = id
    }
    return i1453ad3549b1a27e47913c2d4a29d5625c5b5e04deb3a2a6f264f5cd24a2e8da.NewPrintUsageByPrinterItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ReportsRequestBuilder) DailyPrintUsageSummariesByUser()(*i3772e364c43656cba18386bc364e131c7943e3a0b0461e3a162fb5eae513ac0b.DailyPrintUsageSummariesByUserRequestBuilder) {
    return i3772e364c43656cba18386bc364e131c7943e3a0b0461e3a162fb5eae513ac0b.NewDailyPrintUsageSummariesByUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DailyPrintUsageSummariesByUserById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.print.reports.dailyPrintUsageSummariesByUser.item collection
func (m *ReportsRequestBuilder) DailyPrintUsageSummariesByUserById(id string)(*iffede79819a5c0828784daecd91c1db9829729eb545a0c3b1948096142482a96.PrintUsageByUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByUser_id"] = id
    }
    return iffede79819a5c0828784daecd91c1db9829729eb545a0c3b1948096142482a96.NewPrintUsageByUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Delete delete navigation property reports for print
func (m *ReportsRequestBuilder) Delete(options *ReportsRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get reports from print
func (m *ReportsRequestBuilder) Get(options *ReportsRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ReportRootable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateReportRootFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ReportRootable), nil
}
func (m *ReportsRequestBuilder) MonthlyPrintUsageByPrinter()(*i330d7d86bcb93d59ef14fc1257795533ce121643f4e8becd8884b57bc63e7b4b.MonthlyPrintUsageByPrinterRequestBuilder) {
    return i330d7d86bcb93d59ef14fc1257795533ce121643f4e8becd8884b57bc63e7b4b.NewMonthlyPrintUsageByPrinterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MonthlyPrintUsageByPrinterById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.print.reports.monthlyPrintUsageByPrinter.item collection
func (m *ReportsRequestBuilder) MonthlyPrintUsageByPrinterById(id string)(*i55a12b578425084e14afca840f9f8ad56ad8296a76d84139beb67b9619a01ae2.PrintUsageByPrinterItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByPrinter_id"] = id
    }
    return i55a12b578425084e14afca840f9f8ad56ad8296a76d84139beb67b9619a01ae2.NewPrintUsageByPrinterItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ReportsRequestBuilder) MonthlyPrintUsageByUser()(*i6532f3214bfaa5d2aac58a9d9ffe42fb56e5e2242405e98165287c807365bcda.MonthlyPrintUsageByUserRequestBuilder) {
    return i6532f3214bfaa5d2aac58a9d9ffe42fb56e5e2242405e98165287c807365bcda.NewMonthlyPrintUsageByUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MonthlyPrintUsageByUserById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.print.reports.monthlyPrintUsageByUser.item collection
func (m *ReportsRequestBuilder) MonthlyPrintUsageByUserById(id string)(*i1702ca3f0c998e29512738841125456c0b4e81357e5f7ca9f4176443002d28d5.PrintUsageByUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByUser_id"] = id
    }
    return i1702ca3f0c998e29512738841125456c0b4e81357e5f7ca9f4176443002d28d5.NewPrintUsageByUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ReportsRequestBuilder) MonthlyPrintUsageSummariesByPrinter()(*i27c74777b0efbee67b16cbd84335b5228b74a547d6ffe5e338e3c1c6cc55c61b.MonthlyPrintUsageSummariesByPrinterRequestBuilder) {
    return i27c74777b0efbee67b16cbd84335b5228b74a547d6ffe5e338e3c1c6cc55c61b.NewMonthlyPrintUsageSummariesByPrinterRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MonthlyPrintUsageSummariesByPrinterById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.print.reports.monthlyPrintUsageSummariesByPrinter.item collection
func (m *ReportsRequestBuilder) MonthlyPrintUsageSummariesByPrinterById(id string)(*ic1e71dc755b571267687f63b13312b758bc574de61144bfa50cd81af3b6d3626.PrintUsageByPrinterItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByPrinter_id"] = id
    }
    return ic1e71dc755b571267687f63b13312b758bc574de61144bfa50cd81af3b6d3626.NewPrintUsageByPrinterItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ReportsRequestBuilder) MonthlyPrintUsageSummariesByUser()(*i40dc3bbc1c58ea07c4c8b16942aaeca4da00eb3514241eef66edcd213ba084f7.MonthlyPrintUsageSummariesByUserRequestBuilder) {
    return i40dc3bbc1c58ea07c4c8b16942aaeca4da00eb3514241eef66edcd213ba084f7.NewMonthlyPrintUsageSummariesByUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MonthlyPrintUsageSummariesByUserById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.print.reports.monthlyPrintUsageSummariesByUser.item collection
func (m *ReportsRequestBuilder) MonthlyPrintUsageSummariesByUserById(id string)(*i2268eb94e47ea5d81b666c3310292b49733c18b0196efe2957b868503a9da7ff.PrintUsageByUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["printUsageByUser_id"] = id
    }
    return i2268eb94e47ea5d81b666c3310292b49733c18b0196efe2957b868503a9da7ff.NewPrintUsageByUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property reports in print
func (m *ReportsRequestBuilder) Patch(options *ReportsRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *ReportsRequestBuilder) UserCredentialUsageDetails()(*i254c294d6b358b3854c3ee45f61d16c1237e0477952e917b034bbb0d8a6802b1.UserCredentialUsageDetailsRequestBuilder) {
    return i254c294d6b358b3854c3ee45f61d16c1237e0477952e917b034bbb0d8a6802b1.NewUserCredentialUsageDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserCredentialUsageDetailsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.print.reports.userCredentialUsageDetails.item collection
func (m *ReportsRequestBuilder) UserCredentialUsageDetailsById(id string)(*i4f1c6062b781770b431d75b41deb0a2ac227468563ebac0623bef14f3efe351d.UserCredentialUsageDetailsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userCredentialUsageDetails_id"] = id
    }
    return i4f1c6062b781770b431d75b41deb0a2ac227468563ebac0623bef14f3efe351d.NewUserCredentialUsageDetailsItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
