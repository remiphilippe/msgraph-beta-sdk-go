package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PrinterShareable 
type PrinterShareable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    PrinterBaseable
    GetAllowAllUsers()(*bool)
    GetAllowedGroups()([]Groupable)
    GetAllowedUsers()([]Userable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPrinter()(Printerable)
    GetViewPoint()(PrinterShareViewpointable)
    SetAllowAllUsers(value *bool)()
    SetAllowedGroups(value []Groupable)()
    SetAllowedUsers(value []Userable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPrinter(value Printerable)()
    SetViewPoint(value PrinterShareViewpointable)()
}
