package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the deviceManagement singleton.
type Windows10EditionType int

const (
    WINDOWS10ENTERPRISE_WINDOWS10EDITIONTYPE Windows10EditionType = iota
    WINDOWS10ENTERPRISEN_WINDOWS10EDITIONTYPE
    WINDOWS10EDUCATION_WINDOWS10EDITIONTYPE
    WINDOWS10EDUCATIONN_WINDOWS10EDITIONTYPE
    WINDOWS10MOBILEENTERPRISE_WINDOWS10EDITIONTYPE
    WINDOWS10HOLOGRAPHICENTERPRISE_WINDOWS10EDITIONTYPE
    WINDOWS10PROFESSIONAL_WINDOWS10EDITIONTYPE
    WINDOWS10PROFESSIONALN_WINDOWS10EDITIONTYPE
    WINDOWS10PROFESSIONALEDUCATION_WINDOWS10EDITIONTYPE
    WINDOWS10PROFESSIONALEDUCATIONN_WINDOWS10EDITIONTYPE
    WINDOWS10PROFESSIONALWORKSTATION_WINDOWS10EDITIONTYPE
    WINDOWS10PROFESSIONALWORKSTATIONN_WINDOWS10EDITIONTYPE
    NOTCONFIGURED_WINDOWS10EDITIONTYPE
    WINDOWS10HOME_WINDOWS10EDITIONTYPE
    WINDOWS10HOMECHINA_WINDOWS10EDITIONTYPE
    WINDOWS10HOMEN_WINDOWS10EDITIONTYPE
    WINDOWS10HOMESINGLELANGUAGE_WINDOWS10EDITIONTYPE
    WINDOWS10MOBILE_WINDOWS10EDITIONTYPE
    WINDOWS10IOTCORE_WINDOWS10EDITIONTYPE
    WINDOWS10IOTCORECOMMERCIAL_WINDOWS10EDITIONTYPE
)

func (i Windows10EditionType) String() string {
    return []string{"WINDOWS10ENTERPRISE", "WINDOWS10ENTERPRISEN", "WINDOWS10EDUCATION", "WINDOWS10EDUCATIONN", "WINDOWS10MOBILEENTERPRISE", "WINDOWS10HOLOGRAPHICENTERPRISE", "WINDOWS10PROFESSIONAL", "WINDOWS10PROFESSIONALN", "WINDOWS10PROFESSIONALEDUCATION", "WINDOWS10PROFESSIONALEDUCATIONN", "WINDOWS10PROFESSIONALWORKSTATION", "WINDOWS10PROFESSIONALWORKSTATIONN", "NOTCONFIGURED", "WINDOWS10HOME", "WINDOWS10HOMECHINA", "WINDOWS10HOMEN", "WINDOWS10HOMESINGLELANGUAGE", "WINDOWS10MOBILE", "WINDOWS10IOTCORE", "WINDOWS10IOTCORECOMMERCIAL"}[i]
}
func ParseWindows10EditionType(v string) (interface{}, error) {
    result := WINDOWS10ENTERPRISE_WINDOWS10EDITIONTYPE
    switch strings.ToUpper(v) {
        case "WINDOWS10ENTERPRISE":
            result = WINDOWS10ENTERPRISE_WINDOWS10EDITIONTYPE
        case "WINDOWS10ENTERPRISEN":
            result = WINDOWS10ENTERPRISEN_WINDOWS10EDITIONTYPE
        case "WINDOWS10EDUCATION":
            result = WINDOWS10EDUCATION_WINDOWS10EDITIONTYPE
        case "WINDOWS10EDUCATIONN":
            result = WINDOWS10EDUCATIONN_WINDOWS10EDITIONTYPE
        case "WINDOWS10MOBILEENTERPRISE":
            result = WINDOWS10MOBILEENTERPRISE_WINDOWS10EDITIONTYPE
        case "WINDOWS10HOLOGRAPHICENTERPRISE":
            result = WINDOWS10HOLOGRAPHICENTERPRISE_WINDOWS10EDITIONTYPE
        case "WINDOWS10PROFESSIONAL":
            result = WINDOWS10PROFESSIONAL_WINDOWS10EDITIONTYPE
        case "WINDOWS10PROFESSIONALN":
            result = WINDOWS10PROFESSIONALN_WINDOWS10EDITIONTYPE
        case "WINDOWS10PROFESSIONALEDUCATION":
            result = WINDOWS10PROFESSIONALEDUCATION_WINDOWS10EDITIONTYPE
        case "WINDOWS10PROFESSIONALEDUCATIONN":
            result = WINDOWS10PROFESSIONALEDUCATIONN_WINDOWS10EDITIONTYPE
        case "WINDOWS10PROFESSIONALWORKSTATION":
            result = WINDOWS10PROFESSIONALWORKSTATION_WINDOWS10EDITIONTYPE
        case "WINDOWS10PROFESSIONALWORKSTATIONN":
            result = WINDOWS10PROFESSIONALWORKSTATIONN_WINDOWS10EDITIONTYPE
        case "NOTCONFIGURED":
            result = NOTCONFIGURED_WINDOWS10EDITIONTYPE
        case "WINDOWS10HOME":
            result = WINDOWS10HOME_WINDOWS10EDITIONTYPE
        case "WINDOWS10HOMECHINA":
            result = WINDOWS10HOMECHINA_WINDOWS10EDITIONTYPE
        case "WINDOWS10HOMEN":
            result = WINDOWS10HOMEN_WINDOWS10EDITIONTYPE
        case "WINDOWS10HOMESINGLELANGUAGE":
            result = WINDOWS10HOMESINGLELANGUAGE_WINDOWS10EDITIONTYPE
        case "WINDOWS10MOBILE":
            result = WINDOWS10MOBILE_WINDOWS10EDITIONTYPE
        case "WINDOWS10IOTCORE":
            result = WINDOWS10IOTCORE_WINDOWS10EDITIONTYPE
        case "WINDOWS10IOTCORECOMMERCIAL":
            result = WINDOWS10IOTCORECOMMERCIAL_WINDOWS10EDITIONTYPE
        default:
            return 0, errors.New("Unknown Windows10EditionType value: " + v)
    }
    return &result, nil
}
func SerializeWindows10EditionType(values []Windows10EditionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
