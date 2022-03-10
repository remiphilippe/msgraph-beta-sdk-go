package extractcontentlabel

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/security"
)

// ExtractContentLabelRequestBodyable 
type ExtractContentLabelRequestBodyable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetContentInfo()(i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.ContentInfoable)
    SetContentInfo(value i2263de81f518180fb490a1c688534af1ccfbd4dae2a6d9830596b78378fe7849.ContentInfoable)()
}
