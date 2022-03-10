package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Journalable 
type Journalable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAccount()(Accountable)
    GetBalancingAccountId()(*string)
    GetBalancingAccountNumber()(*string)
    GetCode()(*string)
    GetDisplayName()(*string)
    GetJournalLines()([]JournalLineable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetAccount(value Accountable)()
    SetBalancingAccountId(value *string)()
    SetBalancingAccountNumber(value *string)()
    SetCode(value *string)()
    SetDisplayName(value *string)()
    SetJournalLines(value []JournalLineable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
