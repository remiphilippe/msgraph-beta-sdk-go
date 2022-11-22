package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessReviewScheduleDefinition 
type AccessReviewScheduleDefinition struct {
    Entity
    // Defines the list of additional users or group members to be notified of the access review progress.
    additionalNotificationRecipients []AccessReviewNotificationRecipientItemable
    // The backupReviewers property
    backupReviewers []AccessReviewReviewerScopeable
    // User who created this review. Read-only.
    createdBy UserIdentityable
    // Timestamp when the access review series was created. Supports $select. Read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Description provided by review creators to provide more context of the review to admins. Supports $select.
    descriptionForAdmins *string
    // Description provided  by review creators to provide more context of the review to reviewers. Reviewers will see this description in the email sent to them requesting their review. Email notifications support up to 256 characters. Supports $select.
    descriptionForReviewers *string
    // Name of the access review series. Supports $select and $orderBy. Required on create.
    displayName *string
    // This collection of reviewer scopes is used to define the list of fallback reviewers. These fallback reviewers will be notified to take action if no users are found from the list of reviewers specified. This could occur when either the group owner is specified as the reviewer but the group owner does not exist, or manager is specified as reviewer but a user's manager does not exist. See accessReviewReviewerScope. Replaces backupReviewers. Supports $select. NOTE: The value of this property will be ignored if fallback reviewers are assigned through the stageSettings property.
    fallbackReviewers []AccessReviewReviewerScopeable
    // This property is required when scoping a review to guest users' access across all Microsoft 365 groups and determines which Microsoft 365 groups are reviewed. Each group will become a unique accessReviewInstance of the access review series.  For supported scopes, see accessReviewScope. Supports $select. For examples of options for configuring instanceEnumerationScope, see Configure the scope of your access review definition using the Microsoft Graph API.
    instanceEnumerationScope AccessReviewScopeable
    // Set of access reviews instances for this access review series. Access reviews that do not recur will only have one instance; otherwise, there is an instance for each recurrence.
    instances []AccessReviewInstanceable
    // Timestamp when the access review series was last modified. Supports $select. Read-only.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // This collection of access review scopes is used to define who are the reviewers. The reviewers property is only updatable if individual users are assigned as reviewers. Required on create. Supports $select. For examples of options for assigning reviewers, see Assign reviewers to your access review definition using the Microsoft Graph API. NOTE: The value of this property will be ignored if reviewers are assigned through the stageSettings property.
    reviewers []AccessReviewReviewerScopeable
    // Defines the entities whose access is reviewed. For supported scopes, see accessReviewScope. Required on create. Supports $select and $filter (contains only). For examples of options for configuring scope, see Configure the scope of your access review definition using the Microsoft Graph API.
    scope AccessReviewScopeable
    // The settings for an access review series, see type definition below. Supports $select. Required on create.
    settings AccessReviewScheduleSettingsable
    // Required only for a multi-stage access review to define the stages and their settings. You can break down each review instance into up to three sequential stages, where each stage can have a different set of reviewers, fallback reviewers, and settings. Stages will be created sequentially based on the dependsOn property. Optional.  When this property is defined, its settings are used instead of the corresponding settings in the accessReviewScheduleDefinition object and its settings, reviewers, and fallbackReviewers properties.
    stageSettings []AccessReviewStageSettingsable
    // This read-only field specifies the status of an access review. The typical states include Initializing, NotStarted, Starting, InProgress, Completing, Completed, AutoReviewing, and AutoReviewed.  Supports $select, $orderby, and $filter (eq only). Read-only.
    status *string
}
// NewAccessReviewScheduleDefinition instantiates a new accessReviewScheduleDefinition and sets the default values.
func NewAccessReviewScheduleDefinition()(*AccessReviewScheduleDefinition) {
    m := &AccessReviewScheduleDefinition{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAccessReviewScheduleDefinitionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessReviewScheduleDefinitionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessReviewScheduleDefinition(), nil
}
// GetAdditionalNotificationRecipients gets the additionalNotificationRecipients property value. Defines the list of additional users or group members to be notified of the access review progress.
func (m *AccessReviewScheduleDefinition) GetAdditionalNotificationRecipients()([]AccessReviewNotificationRecipientItemable) {
    return m.additionalNotificationRecipients
}
// GetBackupReviewers gets the backupReviewers property value. The backupReviewers property
func (m *AccessReviewScheduleDefinition) GetBackupReviewers()([]AccessReviewReviewerScopeable) {
    return m.backupReviewers
}
// GetCreatedBy gets the createdBy property value. User who created this review. Read-only.
func (m *AccessReviewScheduleDefinition) GetCreatedBy()(UserIdentityable) {
    return m.createdBy
}
// GetCreatedDateTime gets the createdDateTime property value. Timestamp when the access review series was created. Supports $select. Read-only.
func (m *AccessReviewScheduleDefinition) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDescriptionForAdmins gets the descriptionForAdmins property value. Description provided by review creators to provide more context of the review to admins. Supports $select.
func (m *AccessReviewScheduleDefinition) GetDescriptionForAdmins()(*string) {
    return m.descriptionForAdmins
}
// GetDescriptionForReviewers gets the descriptionForReviewers property value. Description provided  by review creators to provide more context of the review to reviewers. Reviewers will see this description in the email sent to them requesting their review. Email notifications support up to 256 characters. Supports $select.
func (m *AccessReviewScheduleDefinition) GetDescriptionForReviewers()(*string) {
    return m.descriptionForReviewers
}
// GetDisplayName gets the displayName property value. Name of the access review series. Supports $select and $orderBy. Required on create.
func (m *AccessReviewScheduleDefinition) GetDisplayName()(*string) {
    return m.displayName
}
// GetFallbackReviewers gets the fallbackReviewers property value. This collection of reviewer scopes is used to define the list of fallback reviewers. These fallback reviewers will be notified to take action if no users are found from the list of reviewers specified. This could occur when either the group owner is specified as the reviewer but the group owner does not exist, or manager is specified as reviewer but a user's manager does not exist. See accessReviewReviewerScope. Replaces backupReviewers. Supports $select. NOTE: The value of this property will be ignored if fallback reviewers are assigned through the stageSettings property.
func (m *AccessReviewScheduleDefinition) GetFallbackReviewers()([]AccessReviewReviewerScopeable) {
    return m.fallbackReviewers
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessReviewScheduleDefinition) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["additionalNotificationRecipients"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAccessReviewNotificationRecipientItemFromDiscriminatorValue , m.SetAdditionalNotificationRecipients)
    res["backupReviewers"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAccessReviewReviewerScopeFromDiscriminatorValue , m.SetBackupReviewers)
    res["createdBy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateUserIdentityFromDiscriminatorValue , m.SetCreatedBy)
    res["createdDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetCreatedDateTime)
    res["descriptionForAdmins"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescriptionForAdmins)
    res["descriptionForReviewers"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescriptionForReviewers)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["fallbackReviewers"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAccessReviewReviewerScopeFromDiscriminatorValue , m.SetFallbackReviewers)
    res["instanceEnumerationScope"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateAccessReviewScopeFromDiscriminatorValue , m.SetInstanceEnumerationScope)
    res["instances"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAccessReviewInstanceFromDiscriminatorValue , m.SetInstances)
    res["lastModifiedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastModifiedDateTime)
    res["reviewers"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAccessReviewReviewerScopeFromDiscriminatorValue , m.SetReviewers)
    res["scope"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateAccessReviewScopeFromDiscriminatorValue , m.SetScope)
    res["settings"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateAccessReviewScheduleSettingsFromDiscriminatorValue , m.SetSettings)
    res["stageSettings"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAccessReviewStageSettingsFromDiscriminatorValue , m.SetStageSettings)
    res["status"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetStatus)
    return res
}
// GetInstanceEnumerationScope gets the instanceEnumerationScope property value. This property is required when scoping a review to guest users' access across all Microsoft 365 groups and determines which Microsoft 365 groups are reviewed. Each group will become a unique accessReviewInstance of the access review series.  For supported scopes, see accessReviewScope. Supports $select. For examples of options for configuring instanceEnumerationScope, see Configure the scope of your access review definition using the Microsoft Graph API.
func (m *AccessReviewScheduleDefinition) GetInstanceEnumerationScope()(AccessReviewScopeable) {
    return m.instanceEnumerationScope
}
// GetInstances gets the instances property value. Set of access reviews instances for this access review series. Access reviews that do not recur will only have one instance; otherwise, there is an instance for each recurrence.
func (m *AccessReviewScheduleDefinition) GetInstances()([]AccessReviewInstanceable) {
    return m.instances
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Timestamp when the access review series was last modified. Supports $select. Read-only.
func (m *AccessReviewScheduleDefinition) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetReviewers gets the reviewers property value. This collection of access review scopes is used to define who are the reviewers. The reviewers property is only updatable if individual users are assigned as reviewers. Required on create. Supports $select. For examples of options for assigning reviewers, see Assign reviewers to your access review definition using the Microsoft Graph API. NOTE: The value of this property will be ignored if reviewers are assigned through the stageSettings property.
func (m *AccessReviewScheduleDefinition) GetReviewers()([]AccessReviewReviewerScopeable) {
    return m.reviewers
}
// GetScope gets the scope property value. Defines the entities whose access is reviewed. For supported scopes, see accessReviewScope. Required on create. Supports $select and $filter (contains only). For examples of options for configuring scope, see Configure the scope of your access review definition using the Microsoft Graph API.
func (m *AccessReviewScheduleDefinition) GetScope()(AccessReviewScopeable) {
    return m.scope
}
// GetSettings gets the settings property value. The settings for an access review series, see type definition below. Supports $select. Required on create.
func (m *AccessReviewScheduleDefinition) GetSettings()(AccessReviewScheduleSettingsable) {
    return m.settings
}
// GetStageSettings gets the stageSettings property value. Required only for a multi-stage access review to define the stages and their settings. You can break down each review instance into up to three sequential stages, where each stage can have a different set of reviewers, fallback reviewers, and settings. Stages will be created sequentially based on the dependsOn property. Optional.  When this property is defined, its settings are used instead of the corresponding settings in the accessReviewScheduleDefinition object and its settings, reviewers, and fallbackReviewers properties.
func (m *AccessReviewScheduleDefinition) GetStageSettings()([]AccessReviewStageSettingsable) {
    return m.stageSettings
}
// GetStatus gets the status property value. This read-only field specifies the status of an access review. The typical states include Initializing, NotStarted, Starting, InProgress, Completing, Completed, AutoReviewing, and AutoReviewed.  Supports $select, $orderby, and $filter (eq only). Read-only.
func (m *AccessReviewScheduleDefinition) GetStatus()(*string) {
    return m.status
}
// Serialize serializes information the current object
func (m *AccessReviewScheduleDefinition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAdditionalNotificationRecipients() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAdditionalNotificationRecipients())
        err = writer.WriteCollectionOfObjectValues("additionalNotificationRecipients", cast)
        if err != nil {
            return err
        }
    }
    if m.GetBackupReviewers() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetBackupReviewers())
        err = writer.WriteCollectionOfObjectValues("backupReviewers", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("descriptionForAdmins", m.GetDescriptionForAdmins())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("descriptionForReviewers", m.GetDescriptionForReviewers())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetFallbackReviewers() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetFallbackReviewers())
        err = writer.WriteCollectionOfObjectValues("fallbackReviewers", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("instanceEnumerationScope", m.GetInstanceEnumerationScope())
        if err != nil {
            return err
        }
    }
    if m.GetInstances() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetInstances())
        err = writer.WriteCollectionOfObjectValues("instances", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetReviewers() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetReviewers())
        err = writer.WriteCollectionOfObjectValues("reviewers", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("scope", m.GetScope())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("settings", m.GetSettings())
        if err != nil {
            return err
        }
    }
    if m.GetStageSettings() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetStageSettings())
        err = writer.WriteCollectionOfObjectValues("stageSettings", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalNotificationRecipients sets the additionalNotificationRecipients property value. Defines the list of additional users or group members to be notified of the access review progress.
func (m *AccessReviewScheduleDefinition) SetAdditionalNotificationRecipients(value []AccessReviewNotificationRecipientItemable)() {
    m.additionalNotificationRecipients = value
}
// SetBackupReviewers sets the backupReviewers property value. The backupReviewers property
func (m *AccessReviewScheduleDefinition) SetBackupReviewers(value []AccessReviewReviewerScopeable)() {
    m.backupReviewers = value
}
// SetCreatedBy sets the createdBy property value. User who created this review. Read-only.
func (m *AccessReviewScheduleDefinition) SetCreatedBy(value UserIdentityable)() {
    m.createdBy = value
}
// SetCreatedDateTime sets the createdDateTime property value. Timestamp when the access review series was created. Supports $select. Read-only.
func (m *AccessReviewScheduleDefinition) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDescriptionForAdmins sets the descriptionForAdmins property value. Description provided by review creators to provide more context of the review to admins. Supports $select.
func (m *AccessReviewScheduleDefinition) SetDescriptionForAdmins(value *string)() {
    m.descriptionForAdmins = value
}
// SetDescriptionForReviewers sets the descriptionForReviewers property value. Description provided  by review creators to provide more context of the review to reviewers. Reviewers will see this description in the email sent to them requesting their review. Email notifications support up to 256 characters. Supports $select.
func (m *AccessReviewScheduleDefinition) SetDescriptionForReviewers(value *string)() {
    m.descriptionForReviewers = value
}
// SetDisplayName sets the displayName property value. Name of the access review series. Supports $select and $orderBy. Required on create.
func (m *AccessReviewScheduleDefinition) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetFallbackReviewers sets the fallbackReviewers property value. This collection of reviewer scopes is used to define the list of fallback reviewers. These fallback reviewers will be notified to take action if no users are found from the list of reviewers specified. This could occur when either the group owner is specified as the reviewer but the group owner does not exist, or manager is specified as reviewer but a user's manager does not exist. See accessReviewReviewerScope. Replaces backupReviewers. Supports $select. NOTE: The value of this property will be ignored if fallback reviewers are assigned through the stageSettings property.
func (m *AccessReviewScheduleDefinition) SetFallbackReviewers(value []AccessReviewReviewerScopeable)() {
    m.fallbackReviewers = value
}
// SetInstanceEnumerationScope sets the instanceEnumerationScope property value. This property is required when scoping a review to guest users' access across all Microsoft 365 groups and determines which Microsoft 365 groups are reviewed. Each group will become a unique accessReviewInstance of the access review series.  For supported scopes, see accessReviewScope. Supports $select. For examples of options for configuring instanceEnumerationScope, see Configure the scope of your access review definition using the Microsoft Graph API.
func (m *AccessReviewScheduleDefinition) SetInstanceEnumerationScope(value AccessReviewScopeable)() {
    m.instanceEnumerationScope = value
}
// SetInstances sets the instances property value. Set of access reviews instances for this access review series. Access reviews that do not recur will only have one instance; otherwise, there is an instance for each recurrence.
func (m *AccessReviewScheduleDefinition) SetInstances(value []AccessReviewInstanceable)() {
    m.instances = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Timestamp when the access review series was last modified. Supports $select. Read-only.
func (m *AccessReviewScheduleDefinition) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetReviewers sets the reviewers property value. This collection of access review scopes is used to define who are the reviewers. The reviewers property is only updatable if individual users are assigned as reviewers. Required on create. Supports $select. For examples of options for assigning reviewers, see Assign reviewers to your access review definition using the Microsoft Graph API. NOTE: The value of this property will be ignored if reviewers are assigned through the stageSettings property.
func (m *AccessReviewScheduleDefinition) SetReviewers(value []AccessReviewReviewerScopeable)() {
    m.reviewers = value
}
// SetScope sets the scope property value. Defines the entities whose access is reviewed. For supported scopes, see accessReviewScope. Required on create. Supports $select and $filter (contains only). For examples of options for configuring scope, see Configure the scope of your access review definition using the Microsoft Graph API.
func (m *AccessReviewScheduleDefinition) SetScope(value AccessReviewScopeable)() {
    m.scope = value
}
// SetSettings sets the settings property value. The settings for an access review series, see type definition below. Supports $select. Required on create.
func (m *AccessReviewScheduleDefinition) SetSettings(value AccessReviewScheduleSettingsable)() {
    m.settings = value
}
// SetStageSettings sets the stageSettings property value. Required only for a multi-stage access review to define the stages and their settings. You can break down each review instance into up to three sequential stages, where each stage can have a different set of reviewers, fallback reviewers, and settings. Stages will be created sequentially based on the dependsOn property. Optional.  When this property is defined, its settings are used instead of the corresponding settings in the accessReviewScheduleDefinition object and its settings, reviewers, and fallbackReviewers properties.
func (m *AccessReviewScheduleDefinition) SetStageSettings(value []AccessReviewStageSettingsable)() {
    m.stageSettings = value
}
// SetStatus sets the status property value. This read-only field specifies the status of an access review. The typical states include Initializing, NotStarted, Starting, InProgress, Completing, Completed, AutoReviewing, and AutoReviewed.  Supports $select, $orderby, and $filter (eq only). Read-only.
func (m *AccessReviewScheduleDefinition) SetStatus(value *string)() {
    m.status = value
}
