package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessReviewStageSettings 
type AccessReviewStageSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Indicate which decisions will go to the next stage. Can be a sub-set of Approve, Deny, Recommendation, or NotReviewed. If not provided, all decisions will go to the next stage. Optional.
    decisionsThatWillMoveToNextStage []string
    // Defines the sequential or parallel order of the stages and depends on the stageId. Only sequential stages are currently supported. For example, if stageId is 2, then dependsOn must be 1. If stageId is 1, do not specify dependsOn. Required if stageId is not 1.
    dependsOn []string
    // The duration of the stage. Required.  NOTE: The cumulative value of this property across all stages  1. Will override the instanceDurationInDays setting on the accessReviewScheduleDefinition object. 2. Cannot exceed the length of one recurrence. That is, if the review recurs weekly, the cumulative durationInDays cannot exceed 7.
    durationInDays *int32
    // If provided, the fallback reviewers are asked to complete a review if the primary reviewers do not exist. For example, if managers are selected as reviewers and a principal under review does not have a manager in Azure AD, the fallback reviewers are asked to review that principal. NOTE: The value of this property will override the corresponding setting on the accessReviewScheduleDefinition object.
    fallbackReviewers []AccessReviewReviewerScopeable
    // The OdataType property
    odataType *string
    // The recommendationInsightSettings property
    recommendationInsightSettings []AccessReviewRecommendationInsightSettingable
    // Optional field. Indicates the time period of inactivity (with respect to the start date of the review instance) that recommendations will be configured from. The recommendation will be to deny if the user is inactive during the look back duration. For reviews of groups and Azure AD roles, any duration is accepted. For reviews of applications, 30 days is the maximum duration. If not specified, the duration is 30 days. NOTE: The value of this property will override the corresponding setting on the accessReviewScheduleDefinition object.
    recommendationLookBackDuration *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // Indicates whether showing recommendations to reviewers is enabled. Required. NOTE: The value of this property will override override the corresponding setting on the accessReviewScheduleDefinition object.
    recommendationsEnabled *bool
    // Defines who the reviewers are. If none are specified, the review is a self-review (users review their own access).  For examples of options for assigning reviewers, see Assign reviewers to your access review definition using the Microsoft Graph API. NOTE: The value of this property will override the corresponding setting on the accessReviewScheduleDefinition.
    reviewers []AccessReviewReviewerScopeable
    // Unique identifier of the accessReviewStageSettings. The stageId will be used in dependsOn property to indicate the stage relationship. Required.
    stageId *string
}
// NewAccessReviewStageSettings instantiates a new accessReviewStageSettings and sets the default values.
func NewAccessReviewStageSettings()(*AccessReviewStageSettings) {
    m := &AccessReviewStageSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAccessReviewStageSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessReviewStageSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessReviewStageSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AccessReviewStageSettings) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDecisionsThatWillMoveToNextStage gets the decisionsThatWillMoveToNextStage property value. Indicate which decisions will go to the next stage. Can be a sub-set of Approve, Deny, Recommendation, or NotReviewed. If not provided, all decisions will go to the next stage. Optional.
func (m *AccessReviewStageSettings) GetDecisionsThatWillMoveToNextStage()([]string) {
    return m.decisionsThatWillMoveToNextStage
}
// GetDependsOn gets the dependsOn property value. Defines the sequential or parallel order of the stages and depends on the stageId. Only sequential stages are currently supported. For example, if stageId is 2, then dependsOn must be 1. If stageId is 1, do not specify dependsOn. Required if stageId is not 1.
func (m *AccessReviewStageSettings) GetDependsOn()([]string) {
    return m.dependsOn
}
// GetDurationInDays gets the durationInDays property value. The duration of the stage. Required.  NOTE: The cumulative value of this property across all stages  1. Will override the instanceDurationInDays setting on the accessReviewScheduleDefinition object. 2. Cannot exceed the length of one recurrence. That is, if the review recurs weekly, the cumulative durationInDays cannot exceed 7.
func (m *AccessReviewStageSettings) GetDurationInDays()(*int32) {
    return m.durationInDays
}
// GetFallbackReviewers gets the fallbackReviewers property value. If provided, the fallback reviewers are asked to complete a review if the primary reviewers do not exist. For example, if managers are selected as reviewers and a principal under review does not have a manager in Azure AD, the fallback reviewers are asked to review that principal. NOTE: The value of this property will override the corresponding setting on the accessReviewScheduleDefinition object.
func (m *AccessReviewStageSettings) GetFallbackReviewers()([]AccessReviewReviewerScopeable) {
    return m.fallbackReviewers
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessReviewStageSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["decisionsThatWillMoveToNextStage"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetDecisionsThatWillMoveToNextStage)
    res["dependsOn"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetDependsOn)
    res["durationInDays"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetDurationInDays)
    res["fallbackReviewers"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAccessReviewReviewerScopeFromDiscriminatorValue , m.SetFallbackReviewers)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["recommendationInsightSettings"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAccessReviewRecommendationInsightSettingFromDiscriminatorValue , m.SetRecommendationInsightSettings)
    res["recommendationLookBackDuration"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetISODurationValue(m.SetRecommendationLookBackDuration)
    res["recommendationsEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetRecommendationsEnabled)
    res["reviewers"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAccessReviewReviewerScopeFromDiscriminatorValue , m.SetReviewers)
    res["stageId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetStageId)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AccessReviewStageSettings) GetOdataType()(*string) {
    return m.odataType
}
// GetRecommendationInsightSettings gets the recommendationInsightSettings property value. The recommendationInsightSettings property
func (m *AccessReviewStageSettings) GetRecommendationInsightSettings()([]AccessReviewRecommendationInsightSettingable) {
    return m.recommendationInsightSettings
}
// GetRecommendationLookBackDuration gets the recommendationLookBackDuration property value. Optional field. Indicates the time period of inactivity (with respect to the start date of the review instance) that recommendations will be configured from. The recommendation will be to deny if the user is inactive during the look back duration. For reviews of groups and Azure AD roles, any duration is accepted. For reviews of applications, 30 days is the maximum duration. If not specified, the duration is 30 days. NOTE: The value of this property will override the corresponding setting on the accessReviewScheduleDefinition object.
func (m *AccessReviewStageSettings) GetRecommendationLookBackDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.recommendationLookBackDuration
}
// GetRecommendationsEnabled gets the recommendationsEnabled property value. Indicates whether showing recommendations to reviewers is enabled. Required. NOTE: The value of this property will override override the corresponding setting on the accessReviewScheduleDefinition object.
func (m *AccessReviewStageSettings) GetRecommendationsEnabled()(*bool) {
    return m.recommendationsEnabled
}
// GetReviewers gets the reviewers property value. Defines who the reviewers are. If none are specified, the review is a self-review (users review their own access).  For examples of options for assigning reviewers, see Assign reviewers to your access review definition using the Microsoft Graph API. NOTE: The value of this property will override the corresponding setting on the accessReviewScheduleDefinition.
func (m *AccessReviewStageSettings) GetReviewers()([]AccessReviewReviewerScopeable) {
    return m.reviewers
}
// GetStageId gets the stageId property value. Unique identifier of the accessReviewStageSettings. The stageId will be used in dependsOn property to indicate the stage relationship. Required.
func (m *AccessReviewStageSettings) GetStageId()(*string) {
    return m.stageId
}
// Serialize serializes information the current object
func (m *AccessReviewStageSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDecisionsThatWillMoveToNextStage() != nil {
        err := writer.WriteCollectionOfStringValues("decisionsThatWillMoveToNextStage", m.GetDecisionsThatWillMoveToNextStage())
        if err != nil {
            return err
        }
    }
    if m.GetDependsOn() != nil {
        err := writer.WriteCollectionOfStringValues("dependsOn", m.GetDependsOn())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("durationInDays", m.GetDurationInDays())
        if err != nil {
            return err
        }
    }
    if m.GetFallbackReviewers() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetFallbackReviewers())
        err := writer.WriteCollectionOfObjectValues("fallbackReviewers", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetRecommendationInsightSettings() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetRecommendationInsightSettings())
        err := writer.WriteCollectionOfObjectValues("recommendationInsightSettings", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteISODurationValue("recommendationLookBackDuration", m.GetRecommendationLookBackDuration())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("recommendationsEnabled", m.GetRecommendationsEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetReviewers() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetReviewers())
        err := writer.WriteCollectionOfObjectValues("reviewers", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("stageId", m.GetStageId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AccessReviewStageSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDecisionsThatWillMoveToNextStage sets the decisionsThatWillMoveToNextStage property value. Indicate which decisions will go to the next stage. Can be a sub-set of Approve, Deny, Recommendation, or NotReviewed. If not provided, all decisions will go to the next stage. Optional.
func (m *AccessReviewStageSettings) SetDecisionsThatWillMoveToNextStage(value []string)() {
    m.decisionsThatWillMoveToNextStage = value
}
// SetDependsOn sets the dependsOn property value. Defines the sequential or parallel order of the stages and depends on the stageId. Only sequential stages are currently supported. For example, if stageId is 2, then dependsOn must be 1. If stageId is 1, do not specify dependsOn. Required if stageId is not 1.
func (m *AccessReviewStageSettings) SetDependsOn(value []string)() {
    m.dependsOn = value
}
// SetDurationInDays sets the durationInDays property value. The duration of the stage. Required.  NOTE: The cumulative value of this property across all stages  1. Will override the instanceDurationInDays setting on the accessReviewScheduleDefinition object. 2. Cannot exceed the length of one recurrence. That is, if the review recurs weekly, the cumulative durationInDays cannot exceed 7.
func (m *AccessReviewStageSettings) SetDurationInDays(value *int32)() {
    m.durationInDays = value
}
// SetFallbackReviewers sets the fallbackReviewers property value. If provided, the fallback reviewers are asked to complete a review if the primary reviewers do not exist. For example, if managers are selected as reviewers and a principal under review does not have a manager in Azure AD, the fallback reviewers are asked to review that principal. NOTE: The value of this property will override the corresponding setting on the accessReviewScheduleDefinition object.
func (m *AccessReviewStageSettings) SetFallbackReviewers(value []AccessReviewReviewerScopeable)() {
    m.fallbackReviewers = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AccessReviewStageSettings) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRecommendationInsightSettings sets the recommendationInsightSettings property value. The recommendationInsightSettings property
func (m *AccessReviewStageSettings) SetRecommendationInsightSettings(value []AccessReviewRecommendationInsightSettingable)() {
    m.recommendationInsightSettings = value
}
// SetRecommendationLookBackDuration sets the recommendationLookBackDuration property value. Optional field. Indicates the time period of inactivity (with respect to the start date of the review instance) that recommendations will be configured from. The recommendation will be to deny if the user is inactive during the look back duration. For reviews of groups and Azure AD roles, any duration is accepted. For reviews of applications, 30 days is the maximum duration. If not specified, the duration is 30 days. NOTE: The value of this property will override the corresponding setting on the accessReviewScheduleDefinition object.
func (m *AccessReviewStageSettings) SetRecommendationLookBackDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.recommendationLookBackDuration = value
}
// SetRecommendationsEnabled sets the recommendationsEnabled property value. Indicates whether showing recommendations to reviewers is enabled. Required. NOTE: The value of this property will override override the corresponding setting on the accessReviewScheduleDefinition object.
func (m *AccessReviewStageSettings) SetRecommendationsEnabled(value *bool)() {
    m.recommendationsEnabled = value
}
// SetReviewers sets the reviewers property value. Defines who the reviewers are. If none are specified, the review is a self-review (users review their own access).  For examples of options for assigning reviewers, see Assign reviewers to your access review definition using the Microsoft Graph API. NOTE: The value of this property will override the corresponding setting on the accessReviewScheduleDefinition.
func (m *AccessReviewStageSettings) SetReviewers(value []AccessReviewReviewerScopeable)() {
    m.reviewers = value
}
// SetStageId sets the stageId property value. Unique identifier of the accessReviewStageSettings. The stageId will be used in dependsOn property to indicate the stage relationship. Required.
func (m *AccessReviewStageSettings) SetStageId(value *string)() {
    m.stageId = value
}
