package models
import (
    "errors"
)
// 
type MeetingMessageType int

const (
    NONE_MEETINGMESSAGETYPE MeetingMessageType = iota
    MEETINGREQUEST_MEETINGMESSAGETYPE
    MEETINGCANCELLED_MEETINGMESSAGETYPE
    MEETINGACCEPTED_MEETINGMESSAGETYPE
    MEETINGTENTATIVELYACCEPTED_MEETINGMESSAGETYPE
    MEETINGDECLINED_MEETINGMESSAGETYPE
)

func (i MeetingMessageType) String() string {
    return []string{"none", "meetingRequest", "meetingCancelled", "meetingAccepted", "meetingTentativelyAccepted", "meetingDeclined"}[i]
}
func ParseMeetingMessageType(v string) (any, error) {
    result := NONE_MEETINGMESSAGETYPE
    switch v {
        case "none":
            result = NONE_MEETINGMESSAGETYPE
        case "meetingRequest":
            result = MEETINGREQUEST_MEETINGMESSAGETYPE
        case "meetingCancelled":
            result = MEETINGCANCELLED_MEETINGMESSAGETYPE
        case "meetingAccepted":
            result = MEETINGACCEPTED_MEETINGMESSAGETYPE
        case "meetingTentativelyAccepted":
            result = MEETINGTENTATIVELYACCEPTED_MEETINGMESSAGETYPE
        case "meetingDeclined":
            result = MEETINGDECLINED_MEETINGMESSAGETYPE
        default:
            return 0, errors.New("Unknown MeetingMessageType value: " + v)
    }
    return &result, nil
}
func SerializeMeetingMessageType(values []MeetingMessageType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
