package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RiskServicePrincipalActivity 
type RiskServicePrincipalActivity struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Details of the detected risk. Note: Details for this property are only available for Workload Identities Premium customers. Events in tenants without that license will be returned hidden. The possible values are: none, hidden, adminConfirmedServicePrincipalCompromised, adminDismissedAllRiskForServicePrincipal. Note that you must use the Prefer: include-unknown-enum-members request header to get the following value(s) in this evolvable enum: adminConfirmedServicePrincipalCompromised , adminDismissedAllRiskForServicePrincipal.
    detail *RiskDetail
    // The OdataType property
    odataType *string
    // The riskEventTypes property
    riskEventTypes []string
}
// NewRiskServicePrincipalActivity instantiates a new riskServicePrincipalActivity and sets the default values.
func NewRiskServicePrincipalActivity()(*RiskServicePrincipalActivity) {
    m := &RiskServicePrincipalActivity{
    }
    m.SetAdditionalData(make(map[string]any));
    return m
}
// CreateRiskServicePrincipalActivityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRiskServicePrincipalActivityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRiskServicePrincipalActivity(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RiskServicePrincipalActivity) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDetail gets the detail property value. Details of the detected risk. Note: Details for this property are only available for Workload Identities Premium customers. Events in tenants without that license will be returned hidden. The possible values are: none, hidden, adminConfirmedServicePrincipalCompromised, adminDismissedAllRiskForServicePrincipal. Note that you must use the Prefer: include-unknown-enum-members request header to get the following value(s) in this evolvable enum: adminConfirmedServicePrincipalCompromised , adminDismissedAllRiskForServicePrincipal.
func (m *RiskServicePrincipalActivity) GetDetail()(*RiskDetail) {
    return m.detail
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RiskServicePrincipalActivity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["detail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRiskDetail)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetail(val.(*RiskDetail))
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
    res["riskEventTypes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRiskEventTypes(res)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *RiskServicePrincipalActivity) GetOdataType()(*string) {
    return m.odataType
}
// GetRiskEventTypes gets the riskEventTypes property value. The riskEventTypes property
func (m *RiskServicePrincipalActivity) GetRiskEventTypes()([]string) {
    return m.riskEventTypes
}
// Serialize serializes information the current object
func (m *RiskServicePrincipalActivity) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDetail() != nil {
        cast := (*m.GetDetail()).String()
        err := writer.WriteStringValue("detail", &cast)
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
    if m.GetRiskEventTypes() != nil {
        err := writer.WriteCollectionOfStringValues("riskEventTypes", m.GetRiskEventTypes())
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
func (m *RiskServicePrincipalActivity) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDetail sets the detail property value. Details of the detected risk. Note: Details for this property are only available for Workload Identities Premium customers. Events in tenants without that license will be returned hidden. The possible values are: none, hidden, adminConfirmedServicePrincipalCompromised, adminDismissedAllRiskForServicePrincipal. Note that you must use the Prefer: include-unknown-enum-members request header to get the following value(s) in this evolvable enum: adminConfirmedServicePrincipalCompromised , adminDismissedAllRiskForServicePrincipal.
func (m *RiskServicePrincipalActivity) SetDetail(value *RiskDetail)() {
    m.detail = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *RiskServicePrincipalActivity) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRiskEventTypes sets the riskEventTypes property value. The riskEventTypes property
func (m *RiskServicePrincipalActivity) SetRiskEventTypes(value []string)() {
    m.riskEventTypes = value
}
