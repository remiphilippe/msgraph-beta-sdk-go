package models
import (
    "errors"
)
// Provides operations to manage the collection of accessReview entities.
type BrowserSharedCookieStatus int

const (
    // A sharedcookie that has been published
    PUBLISHED_BROWSERSHAREDCOOKIESTATUS BrowserSharedCookieStatus = iota
    // A sharedcookie that has been added pending publish
    PENDINGADD_BROWSERSHAREDCOOKIESTATUS
    // A sharedcookie that has been edited pending publish
    PENDINGEDIT_BROWSERSHAREDCOOKIESTATUS
    // A sharedcookie that has been deleted pending publish
    PENDINGDELETE_BROWSERSHAREDCOOKIESTATUS
    // Placeholder for evolvable enum, but this enum is never returned to the caller, so it shouldn't be necessary.
    UNKNOWNFUTUREVALUE_BROWSERSHAREDCOOKIESTATUS
)

func (i BrowserSharedCookieStatus) String() string {
    return []string{"published", "pendingAdd", "pendingEdit", "pendingDelete", "unknownFutureValue"}[i]
}
func ParseBrowserSharedCookieStatus(v string) (interface{}, error) {
    result := PUBLISHED_BROWSERSHAREDCOOKIESTATUS
    switch v {
        case "published":
            result = PUBLISHED_BROWSERSHAREDCOOKIESTATUS
        case "pendingAdd":
            result = PENDINGADD_BROWSERSHAREDCOOKIESTATUS
        case "pendingEdit":
            result = PENDINGEDIT_BROWSERSHAREDCOOKIESTATUS
        case "pendingDelete":
            result = PENDINGDELETE_BROWSERSHAREDCOOKIESTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_BROWSERSHAREDCOOKIESTATUS
        default:
            return 0, errors.New("Unknown BrowserSharedCookieStatus value: " + v)
    }
    return &result, nil
}
func SerializeBrowserSharedCookieStatus(values []BrowserSharedCookieStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
