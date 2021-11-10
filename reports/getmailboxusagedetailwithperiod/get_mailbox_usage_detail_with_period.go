package getmailboxusagedetailwithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// 
type GetMailboxUsageDetailWithPeriod struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    // 
    createdDate *string;
    // 
    deletedDate *string;
    // 
    deletedItemCount *int64;
    // 
    deletedItemSizeInBytes *int64;
    // 
    displayName *string;
    // 
    isDeleted *bool;
    // 
    issueWarningQuotaInBytes *int64;
    // 
    itemCount *int64;
    // 
    lastActivityDate *string;
    // 
    prohibitSendQuotaInBytes *int64;
    // 
    prohibitSendReceiveQuotaInBytes *int64;
    // 
    reportPeriod *string;
    // 
    reportRefreshDate *string;
    // 
    storageUsedInBytes *int64;
    // 
    userPrincipalName *string;
}
// Instantiates a new getMailboxUsageDetailWithPeriod and sets the default values.
func NewGetMailboxUsageDetailWithPeriod()(*GetMailboxUsageDetailWithPeriod) {
    m := &GetMailboxUsageDetailWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
// Gets the createdDate property value. 
func (m *GetMailboxUsageDetailWithPeriod) GetCreatedDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.createdDate
    }
}
// Gets the deletedDate property value. 
func (m *GetMailboxUsageDetailWithPeriod) GetDeletedDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deletedDate
    }
}
// Gets the deletedItemCount property value. 
func (m *GetMailboxUsageDetailWithPeriod) GetDeletedItemCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.deletedItemCount
    }
}
// Gets the deletedItemSizeInBytes property value. 
func (m *GetMailboxUsageDetailWithPeriod) GetDeletedItemSizeInBytes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.deletedItemSizeInBytes
    }
}
// Gets the displayName property value. 
func (m *GetMailboxUsageDetailWithPeriod) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the isDeleted property value. 
func (m *GetMailboxUsageDetailWithPeriod) GetIsDeleted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDeleted
    }
}
// Gets the issueWarningQuotaInBytes property value. 
func (m *GetMailboxUsageDetailWithPeriod) GetIssueWarningQuotaInBytes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.issueWarningQuotaInBytes
    }
}
// Gets the itemCount property value. 
func (m *GetMailboxUsageDetailWithPeriod) GetItemCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.itemCount
    }
}
// Gets the lastActivityDate property value. 
func (m *GetMailboxUsageDetailWithPeriod) GetLastActivityDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lastActivityDate
    }
}
// Gets the prohibitSendQuotaInBytes property value. 
func (m *GetMailboxUsageDetailWithPeriod) GetProhibitSendQuotaInBytes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.prohibitSendQuotaInBytes
    }
}
// Gets the prohibitSendReceiveQuotaInBytes property value. 
func (m *GetMailboxUsageDetailWithPeriod) GetProhibitSendReceiveQuotaInBytes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.prohibitSendReceiveQuotaInBytes
    }
}
// Gets the reportPeriod property value. 
func (m *GetMailboxUsageDetailWithPeriod) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
// Gets the reportRefreshDate property value. 
func (m *GetMailboxUsageDetailWithPeriod) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
// Gets the storageUsedInBytes property value. 
func (m *GetMailboxUsageDetailWithPeriod) GetStorageUsedInBytes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.storageUsedInBytes
    }
}
// Gets the userPrincipalName property value. 
func (m *GetMailboxUsageDetailWithPeriod) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// The deserialization information for the current model
func (m *GetMailboxUsageDetailWithPeriod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["createdDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDate(val)
        }
        return nil
    }
    res["deletedDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeletedDate(val)
        }
        return nil
    }
    res["deletedItemCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeletedItemCount(val)
        }
        return nil
    }
    res["deletedItemSizeInBytes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeletedItemSizeInBytes(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["isDeleted"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDeleted(val)
        }
        return nil
    }
    res["issueWarningQuotaInBytes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssueWarningQuotaInBytes(val)
        }
        return nil
    }
    res["itemCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetItemCount(val)
        }
        return nil
    }
    res["lastActivityDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastActivityDate(val)
        }
        return nil
    }
    res["prohibitSendQuotaInBytes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProhibitSendQuotaInBytes(val)
        }
        return nil
    }
    res["prohibitSendReceiveQuotaInBytes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProhibitSendReceiveQuotaInBytes(val)
        }
        return nil
    }
    res["reportPeriod"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReportPeriod(val)
        }
        return nil
    }
    res["reportRefreshDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReportRefreshDate(val)
        }
        return nil
    }
    res["storageUsedInBytes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorageUsedInBytes(val)
        }
        return nil
    }
    res["userPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    return res
}
func (m *GetMailboxUsageDetailWithPeriod) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *GetMailboxUsageDetailWithPeriod) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("createdDate", m.GetCreatedDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deletedDate", m.GetDeletedDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("deletedItemCount", m.GetDeletedItemCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("deletedItemSizeInBytes", m.GetDeletedItemSizeInBytes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDeleted", m.GetIsDeleted())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("issueWarningQuotaInBytes", m.GetIssueWarningQuotaInBytes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("itemCount", m.GetItemCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("lastActivityDate", m.GetLastActivityDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("prohibitSendQuotaInBytes", m.GetProhibitSendQuotaInBytes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("prohibitSendReceiveQuotaInBytes", m.GetProhibitSendReceiveQuotaInBytes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("reportPeriod", m.GetReportPeriod())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("reportRefreshDate", m.GetReportRefreshDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("storageUsedInBytes", m.GetStorageUsedInBytes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the createdDate property value. 
// Parameters:
//  - value : Value to set for the createdDate property.
func (m *GetMailboxUsageDetailWithPeriod) SetCreatedDate(value *string)() {
    m.createdDate = value
}
// Sets the deletedDate property value. 
// Parameters:
//  - value : Value to set for the deletedDate property.
func (m *GetMailboxUsageDetailWithPeriod) SetDeletedDate(value *string)() {
    m.deletedDate = value
}
// Sets the deletedItemCount property value. 
// Parameters:
//  - value : Value to set for the deletedItemCount property.
func (m *GetMailboxUsageDetailWithPeriod) SetDeletedItemCount(value *int64)() {
    m.deletedItemCount = value
}
// Sets the deletedItemSizeInBytes property value. 
// Parameters:
//  - value : Value to set for the deletedItemSizeInBytes property.
func (m *GetMailboxUsageDetailWithPeriod) SetDeletedItemSizeInBytes(value *int64)() {
    m.deletedItemSizeInBytes = value
}
// Sets the displayName property value. 
// Parameters:
//  - value : Value to set for the displayName property.
func (m *GetMailboxUsageDetailWithPeriod) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the isDeleted property value. 
// Parameters:
//  - value : Value to set for the isDeleted property.
func (m *GetMailboxUsageDetailWithPeriod) SetIsDeleted(value *bool)() {
    m.isDeleted = value
}
// Sets the issueWarningQuotaInBytes property value. 
// Parameters:
//  - value : Value to set for the issueWarningQuotaInBytes property.
func (m *GetMailboxUsageDetailWithPeriod) SetIssueWarningQuotaInBytes(value *int64)() {
    m.issueWarningQuotaInBytes = value
}
// Sets the itemCount property value. 
// Parameters:
//  - value : Value to set for the itemCount property.
func (m *GetMailboxUsageDetailWithPeriod) SetItemCount(value *int64)() {
    m.itemCount = value
}
// Sets the lastActivityDate property value. 
// Parameters:
//  - value : Value to set for the lastActivityDate property.
func (m *GetMailboxUsageDetailWithPeriod) SetLastActivityDate(value *string)() {
    m.lastActivityDate = value
}
// Sets the prohibitSendQuotaInBytes property value. 
// Parameters:
//  - value : Value to set for the prohibitSendQuotaInBytes property.
func (m *GetMailboxUsageDetailWithPeriod) SetProhibitSendQuotaInBytes(value *int64)() {
    m.prohibitSendQuotaInBytes = value
}
// Sets the prohibitSendReceiveQuotaInBytes property value. 
// Parameters:
//  - value : Value to set for the prohibitSendReceiveQuotaInBytes property.
func (m *GetMailboxUsageDetailWithPeriod) SetProhibitSendReceiveQuotaInBytes(value *int64)() {
    m.prohibitSendReceiveQuotaInBytes = value
}
// Sets the reportPeriod property value. 
// Parameters:
//  - value : Value to set for the reportPeriod property.
func (m *GetMailboxUsageDetailWithPeriod) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
// Sets the reportRefreshDate property value. 
// Parameters:
//  - value : Value to set for the reportRefreshDate property.
func (m *GetMailboxUsageDetailWithPeriod) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}
// Sets the storageUsedInBytes property value. 
// Parameters:
//  - value : Value to set for the storageUsedInBytes property.
func (m *GetMailboxUsageDetailWithPeriod) SetStorageUsedInBytes(value *int64)() {
    m.storageUsedInBytes = value
}
// Sets the userPrincipalName property value. 
// Parameters:
//  - value : Value to set for the userPrincipalName property.
func (m *GetMailboxUsageDetailWithPeriod) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
