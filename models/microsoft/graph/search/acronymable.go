package search

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Acronymable 
type Acronymable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    SearchAnswerable
    GetStandsFor()(*string)
    GetState()(*AnswerState)
    SetStandsFor(value *string)()
    SetState(value *AnswerState)()
}
