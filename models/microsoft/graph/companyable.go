package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Companyable 
type Companyable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAccounts()([]Accountable)
    GetAgedAccountsPayable()([]AgedAccountsPayableable)
    GetAgedAccountsReceivable()([]AgedAccountsReceivableable)
    GetBusinessProfileId()(*string)
    GetCompanyInformation()([]CompanyInformationable)
    GetCountriesRegions()([]CountryRegionable)
    GetCurrencies()([]Currencyable)
    GetCustomerPaymentJournals()([]CustomerPaymentJournalable)
    GetCustomerPayments()([]CustomerPaymentable)
    GetCustomers()([]Customerable)
    GetDimensions()([]Dimensionable)
    GetDimensionValues()([]DimensionValueable)
    GetDisplayName()(*string)
    GetEmployees()([]Employeeable)
    GetGeneralLedgerEntries()([]GeneralLedgerEntryable)
    GetItemCategories()([]ItemCategoryable)
    GetItems()([]Itemable)
    GetJournalLines()([]JournalLineable)
    GetJournals()([]Journalable)
    GetName()(*string)
    GetPaymentMethods()([]PaymentMethodable)
    GetPaymentTerms()([]PaymentTermable)
    GetPicture()([]Pictureable)
    GetPurchaseInvoiceLines()([]PurchaseInvoiceLineable)
    GetPurchaseInvoices()([]PurchaseInvoiceable)
    GetSalesCreditMemoLines()([]SalesCreditMemoLineable)
    GetSalesCreditMemos()([]SalesCreditMemoable)
    GetSalesInvoiceLines()([]SalesInvoiceLineable)
    GetSalesInvoices()([]SalesInvoiceable)
    GetSalesOrderLines()([]SalesOrderLineable)
    GetSalesOrders()([]SalesOrderable)
    GetSalesQuoteLines()([]SalesQuoteLineable)
    GetSalesQuotes()([]SalesQuoteable)
    GetShipmentMethods()([]ShipmentMethodable)
    GetSystemVersion()(*string)
    GetTaxAreas()([]TaxAreaable)
    GetTaxGroups()([]TaxGroupable)
    GetUnitsOfMeasure()([]UnitOfMeasureable)
    GetVendors()([]Vendor_escapedable)
    SetAccounts(value []Accountable)()
    SetAgedAccountsPayable(value []AgedAccountsPayableable)()
    SetAgedAccountsReceivable(value []AgedAccountsReceivableable)()
    SetBusinessProfileId(value *string)()
    SetCompanyInformation(value []CompanyInformationable)()
    SetCountriesRegions(value []CountryRegionable)()
    SetCurrencies(value []Currencyable)()
    SetCustomerPaymentJournals(value []CustomerPaymentJournalable)()
    SetCustomerPayments(value []CustomerPaymentable)()
    SetCustomers(value []Customerable)()
    SetDimensions(value []Dimensionable)()
    SetDimensionValues(value []DimensionValueable)()
    SetDisplayName(value *string)()
    SetEmployees(value []Employeeable)()
    SetGeneralLedgerEntries(value []GeneralLedgerEntryable)()
    SetItemCategories(value []ItemCategoryable)()
    SetItems(value []Itemable)()
    SetJournalLines(value []JournalLineable)()
    SetJournals(value []Journalable)()
    SetName(value *string)()
    SetPaymentMethods(value []PaymentMethodable)()
    SetPaymentTerms(value []PaymentTermable)()
    SetPicture(value []Pictureable)()
    SetPurchaseInvoiceLines(value []PurchaseInvoiceLineable)()
    SetPurchaseInvoices(value []PurchaseInvoiceable)()
    SetSalesCreditMemoLines(value []SalesCreditMemoLineable)()
    SetSalesCreditMemos(value []SalesCreditMemoable)()
    SetSalesInvoiceLines(value []SalesInvoiceLineable)()
    SetSalesInvoices(value []SalesInvoiceable)()
    SetSalesOrderLines(value []SalesOrderLineable)()
    SetSalesOrders(value []SalesOrderable)()
    SetSalesQuoteLines(value []SalesQuoteLineable)()
    SetSalesQuotes(value []SalesQuoteable)()
    SetShipmentMethods(value []ShipmentMethodable)()
    SetSystemVersion(value *string)()
    SetTaxAreas(value []TaxAreaable)()
    SetTaxGroups(value []TaxGroupable)()
    SetUnitsOfMeasure(value []UnitOfMeasureable)()
    SetVendors(value []Vendor_escapedable)()
}
