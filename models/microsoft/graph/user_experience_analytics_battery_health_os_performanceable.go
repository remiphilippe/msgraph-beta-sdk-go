package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserExperienceAnalyticsBatteryHealthOsPerformanceable 
type UserExperienceAnalyticsBatteryHealthOsPerformanceable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetActiveDevices()(*int32)
    GetAverageBatteryAgeInDays()(*int32)
    GetAverageEstimatedRuntimeInMinutes()(*int32)
    GetAverageMaxCapacityPercentage()(*int32)
    GetOsBuildNumber()(*string)
    GetOsVersion()(*string)
    SetActiveDevices(value *int32)()
    SetAverageBatteryAgeInDays(value *int32)()
    SetAverageEstimatedRuntimeInMinutes(value *int32)()
    SetAverageMaxCapacityPercentage(value *int32)()
    SetOsBuildNumber(value *string)()
    SetOsVersion(value *string)()
}
