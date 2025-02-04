package models
import (
    "errors"
)
// 
type EntityType int

const (
    EVENT_ENTITYTYPE EntityType = iota
    MESSAGE_ENTITYTYPE
    DRIVEITEM_ENTITYTYPE
    EXTERNALITEM_ENTITYTYPE
    SITE_ENTITYTYPE
    LIST_ENTITYTYPE
    LISTITEM_ENTITYTYPE
    DRIVE_ENTITYTYPE
    UNKNOWNFUTUREVALUE_ENTITYTYPE
    ACRONYM_ENTITYTYPE
    BOOKMARK_ENTITYTYPE
    CHATMESSAGE_ENTITYTYPE
)

func (i EntityType) String() string {
    return []string{"event", "message", "driveItem", "externalItem", "site", "list", "listItem", "drive", "unknownFutureValue", "acronym", "bookmark", "chatMessage"}[i]
}
func ParseEntityType(v string) (any, error) {
    result := EVENT_ENTITYTYPE
    switch v {
        case "event":
            result = EVENT_ENTITYTYPE
        case "message":
            result = MESSAGE_ENTITYTYPE
        case "driveItem":
            result = DRIVEITEM_ENTITYTYPE
        case "externalItem":
            result = EXTERNALITEM_ENTITYTYPE
        case "site":
            result = SITE_ENTITYTYPE
        case "list":
            result = LIST_ENTITYTYPE
        case "listItem":
            result = LISTITEM_ENTITYTYPE
        case "drive":
            result = DRIVE_ENTITYTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ENTITYTYPE
        case "acronym":
            result = ACRONYM_ENTITYTYPE
        case "bookmark":
            result = BOOKMARK_ENTITYTYPE
        case "chatMessage":
            result = CHATMESSAGE_ENTITYTYPE
        default:
            return 0, errors.New("Unknown EntityType value: " + v)
    }
    return &result, nil
}
func SerializeEntityType(values []EntityType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
