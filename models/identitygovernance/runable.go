package identitygovernance

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// Runable 
type Runable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCompletedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetFailedTasksCount()(*int32)
    GetFailedUsersCount()(*int32)
    GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetProcessingStatus()(*LifecycleWorkflowProcessingStatus)
    GetScheduledDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetStartedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSuccessfulUsersCount()(*int32)
    GetTaskProcessingResults()([]TaskProcessingResultable)
    GetTotalTasksCount()(*int32)
    GetTotalUnprocessedTasksCount()(*int32)
    GetTotalUsersCount()(*int32)
    GetUserProcessingResults()([]UserProcessingResultable)
    GetWorkflowExecutionType()(*WorkflowExecutionType)
    SetCompletedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetFailedTasksCount(value *int32)()
    SetFailedUsersCount(value *int32)()
    SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetProcessingStatus(value *LifecycleWorkflowProcessingStatus)()
    SetScheduledDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetStartedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSuccessfulUsersCount(value *int32)()
    SetTaskProcessingResults(value []TaskProcessingResultable)()
    SetTotalTasksCount(value *int32)()
    SetTotalUnprocessedTasksCount(value *int32)()
    SetTotalUsersCount(value *int32)()
    SetUserProcessingResults(value []UserProcessingResultable)()
    SetWorkflowExecutionType(value *WorkflowExecutionType)()
}
