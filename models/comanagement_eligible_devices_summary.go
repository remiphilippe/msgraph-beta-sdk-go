package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ComanagementEligibleDevicesSummary 
type ComanagementEligibleDevicesSummary struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Count of devices already Co-Managed
    comanagedCount *int32
    // Count of devices eligible for Co-Management but not yet joined to Azure Active Directory
    eligibleButNotAzureAdJoinedCount *int32
    // Count of devices fully eligible for Co-Management
    eligibleCount *int32
    // Count of devices ineligible for Co-Management
    ineligibleCount *int32
    // Count of devices that will be eligible for Co-Management after an OS update
    needsOsUpdateCount *int32
    // The OdataType property
    odataType *string
    // Count of devices scheduled for Co-Management enrollment. Valid values 0 to 9999999
    scheduledForEnrollmentCount *int32
}
// NewComanagementEligibleDevicesSummary instantiates a new comanagementEligibleDevicesSummary and sets the default values.
func NewComanagementEligibleDevicesSummary()(*ComanagementEligibleDevicesSummary) {
    m := &ComanagementEligibleDevicesSummary{
    }
    m.SetAdditionalData(make(map[string]any));
    return m
}
// CreateComanagementEligibleDevicesSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateComanagementEligibleDevicesSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewComanagementEligibleDevicesSummary(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ComanagementEligibleDevicesSummary) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetComanagedCount gets the comanagedCount property value. Count of devices already Co-Managed
func (m *ComanagementEligibleDevicesSummary) GetComanagedCount()(*int32) {
    return m.comanagedCount
}
// GetEligibleButNotAzureAdJoinedCount gets the eligibleButNotAzureAdJoinedCount property value. Count of devices eligible for Co-Management but not yet joined to Azure Active Directory
func (m *ComanagementEligibleDevicesSummary) GetEligibleButNotAzureAdJoinedCount()(*int32) {
    return m.eligibleButNotAzureAdJoinedCount
}
// GetEligibleCount gets the eligibleCount property value. Count of devices fully eligible for Co-Management
func (m *ComanagementEligibleDevicesSummary) GetEligibleCount()(*int32) {
    return m.eligibleCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ComanagementEligibleDevicesSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["comanagedCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComanagedCount(val)
        }
        return nil
    }
    res["eligibleButNotAzureAdJoinedCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEligibleButNotAzureAdJoinedCount(val)
        }
        return nil
    }
    res["eligibleCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEligibleCount(val)
        }
        return nil
    }
    res["ineligibleCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIneligibleCount(val)
        }
        return nil
    }
    res["needsOsUpdateCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNeedsOsUpdateCount(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["scheduledForEnrollmentCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScheduledForEnrollmentCount(val)
        }
        return nil
    }
    return res
}
// GetIneligibleCount gets the ineligibleCount property value. Count of devices ineligible for Co-Management
func (m *ComanagementEligibleDevicesSummary) GetIneligibleCount()(*int32) {
    return m.ineligibleCount
}
// GetNeedsOsUpdateCount gets the needsOsUpdateCount property value. Count of devices that will be eligible for Co-Management after an OS update
func (m *ComanagementEligibleDevicesSummary) GetNeedsOsUpdateCount()(*int32) {
    return m.needsOsUpdateCount
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ComanagementEligibleDevicesSummary) GetOdataType()(*string) {
    return m.odataType
}
// GetScheduledForEnrollmentCount gets the scheduledForEnrollmentCount property value. Count of devices scheduled for Co-Management enrollment. Valid values 0 to 9999999
func (m *ComanagementEligibleDevicesSummary) GetScheduledForEnrollmentCount()(*int32) {
    return m.scheduledForEnrollmentCount
}
// Serialize serializes information the current object
func (m *ComanagementEligibleDevicesSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("comanagedCount", m.GetComanagedCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("eligibleButNotAzureAdJoinedCount", m.GetEligibleButNotAzureAdJoinedCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("eligibleCount", m.GetEligibleCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("ineligibleCount", m.GetIneligibleCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("needsOsUpdateCount", m.GetNeedsOsUpdateCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("scheduledForEnrollmentCount", m.GetScheduledForEnrollmentCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ComanagementEligibleDevicesSummary) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetComanagedCount sets the comanagedCount property value. Count of devices already Co-Managed
func (m *ComanagementEligibleDevicesSummary) SetComanagedCount(value *int32)() {
    m.comanagedCount = value
}
// SetEligibleButNotAzureAdJoinedCount sets the eligibleButNotAzureAdJoinedCount property value. Count of devices eligible for Co-Management but not yet joined to Azure Active Directory
func (m *ComanagementEligibleDevicesSummary) SetEligibleButNotAzureAdJoinedCount(value *int32)() {
    m.eligibleButNotAzureAdJoinedCount = value
}
// SetEligibleCount sets the eligibleCount property value. Count of devices fully eligible for Co-Management
func (m *ComanagementEligibleDevicesSummary) SetEligibleCount(value *int32)() {
    m.eligibleCount = value
}
// SetIneligibleCount sets the ineligibleCount property value. Count of devices ineligible for Co-Management
func (m *ComanagementEligibleDevicesSummary) SetIneligibleCount(value *int32)() {
    m.ineligibleCount = value
}
// SetNeedsOsUpdateCount sets the needsOsUpdateCount property value. Count of devices that will be eligible for Co-Management after an OS update
func (m *ComanagementEligibleDevicesSummary) SetNeedsOsUpdateCount(value *int32)() {
    m.needsOsUpdateCount = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ComanagementEligibleDevicesSummary) SetOdataType(value *string)() {
    m.odataType = value
}
// SetScheduledForEnrollmentCount sets the scheduledForEnrollmentCount property value. Count of devices scheduled for Co-Management enrollment. Valid values 0 to 9999999
func (m *ComanagementEligibleDevicesSummary) SetScheduledForEnrollmentCount(value *int32)() {
    m.scheduledForEnrollmentCount = value
}
