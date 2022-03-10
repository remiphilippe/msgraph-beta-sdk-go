package sendcustomnotificationtocompanyportal

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SendCustomNotificationToCompanyPortalRequestBodyable 
type SendCustomNotificationToCompanyPortalRequestBodyable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetGroupsToNotify()([]string)
    GetNotificationBody()(*string)
    GetNotificationTitle()(*string)
    SetGroupsToNotify(value []string)()
    SetNotificationBody(value *string)()
    SetNotificationTitle(value *string)()
}
