package security

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// InformationProtectionPolicySettingable 
type InformationProtectionPolicySettingable interface {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDefaultLabelId()(*string)
    GetIsDowngradeJustificationRequired()(*bool)
    GetIsMandatory()(*bool)
    GetMoreInfoUrl()(*string)
    SetDefaultLabelId(value *string)()
    SetIsDowngradeJustificationRequired(value *bool)()
    SetIsMandatory(value *bool)()
    SetMoreInfoUrl(value *string)()
}
