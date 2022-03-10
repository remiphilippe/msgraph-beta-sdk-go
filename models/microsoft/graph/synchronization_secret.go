package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the collection of application entities.
type SynchronizationSecret int

const (
    NONE_SYNCHRONIZATIONSECRET SynchronizationSecret = iota
    USERNAME_SYNCHRONIZATIONSECRET
    PASSWORD_SYNCHRONIZATIONSECRET
    SECRETTOKEN_SYNCHRONIZATIONSECRET
    APPKEY_SYNCHRONIZATIONSECRET
    BASEADDRESS_SYNCHRONIZATIONSECRET
    CLIENTIDENTIFIER_SYNCHRONIZATIONSECRET
    CLIENTSECRET_SYNCHRONIZATIONSECRET
    SINGLESIGNONTYPE_SYNCHRONIZATIONSECRET
    SANDBOX_SYNCHRONIZATIONSECRET
    URL_SYNCHRONIZATIONSECRET
    DOMAIN_SYNCHRONIZATIONSECRET
    CONSUMERKEY_SYNCHRONIZATIONSECRET
    CONSUMERSECRET_SYNCHRONIZATIONSECRET
    TOKENKEY_SYNCHRONIZATIONSECRET
    TOKENEXPIRATION_SYNCHRONIZATIONSECRET
    OAUTH2ACCESSTOKEN_SYNCHRONIZATIONSECRET
    OAUTH2ACCESSTOKENCREATIONTIME_SYNCHRONIZATIONSECRET
    OAUTH2REFRESHTOKEN_SYNCHRONIZATIONSECRET
    SYNCALL_SYNCHRONIZATIONSECRET
    INSTANCENAME_SYNCHRONIZATIONSECRET
    OAUTH2CLIENTID_SYNCHRONIZATIONSECRET
    OAUTH2CLIENTSECRET_SYNCHRONIZATIONSECRET
    COMPANYID_SYNCHRONIZATIONSECRET
    UPDATEKEYONSOFTDELETE_SYNCHRONIZATIONSECRET
    SYNCHRONIZATIONSCHEDULE_SYNCHRONIZATIONSECRET
    SYSTEMOFRECORD_SYNCHRONIZATIONSECRET
    SANDBOXNAME_SYNCHRONIZATIONSECRET
    ENFORCEDOMAIN_SYNCHRONIZATIONSECRET
    SYNCNOTIFICATIONSETTINGS_SYNCHRONIZATIONSECRET
    SKIPOUTOFSCOPEDELETIONS_SYNCHRONIZATIONSECRET
    OAUTH2AUTHORIZATIONCODE_SYNCHRONIZATIONSECRET
    OAUTH2REDIRECTURI_SYNCHRONIZATIONSECRET
    APPLICATIONTEMPLATEIDENTIFIER_SYNCHRONIZATIONSECRET
    OAUTH2TOKENEXCHANGEURI_SYNCHRONIZATIONSECRET
    OAUTH2AUTHORIZATIONURI_SYNCHRONIZATIONSECRET
    AUTHENTICATIONTYPE_SYNCHRONIZATIONSECRET
    SERVER_SYNCHRONIZATIONSECRET
    PERFORMINBOUNDENTITLEMENTGRANTS_SYNCHRONIZATIONSECRET
    HARDDELETESENABLED_SYNCHRONIZATIONSECRET
    SYNCAGENTCOMPATIBILITYKEY_SYNCHRONIZATIONSECRET
    SYNCAGENTADCONTAINER_SYNCHRONIZATIONSECRET
    VALIDATEDOMAIN_SYNCHRONIZATIONSECRET
    TESTREFERENCES_SYNCHRONIZATIONSECRET
    CONNECTIONSTRING_SYNCHRONIZATIONSECRET
)

func (i SynchronizationSecret) String() string {
    return []string{"NONE", "USERNAME", "PASSWORD", "SECRETTOKEN", "APPKEY", "BASEADDRESS", "CLIENTIDENTIFIER", "CLIENTSECRET", "SINGLESIGNONTYPE", "SANDBOX", "URL", "DOMAIN", "CONSUMERKEY", "CONSUMERSECRET", "TOKENKEY", "TOKENEXPIRATION", "OAUTH2ACCESSTOKEN", "OAUTH2ACCESSTOKENCREATIONTIME", "OAUTH2REFRESHTOKEN", "SYNCALL", "INSTANCENAME", "OAUTH2CLIENTID", "OAUTH2CLIENTSECRET", "COMPANYID", "UPDATEKEYONSOFTDELETE", "SYNCHRONIZATIONSCHEDULE", "SYSTEMOFRECORD", "SANDBOXNAME", "ENFORCEDOMAIN", "SYNCNOTIFICATIONSETTINGS", "SKIPOUTOFSCOPEDELETIONS", "OAUTH2AUTHORIZATIONCODE", "OAUTH2REDIRECTURI", "APPLICATIONTEMPLATEIDENTIFIER", "OAUTH2TOKENEXCHANGEURI", "OAUTH2AUTHORIZATIONURI", "AUTHENTICATIONTYPE", "SERVER", "PERFORMINBOUNDENTITLEMENTGRANTS", "HARDDELETESENABLED", "SYNCAGENTCOMPATIBILITYKEY", "SYNCAGENTADCONTAINER", "VALIDATEDOMAIN", "TESTREFERENCES", "CONNECTIONSTRING"}[i]
}
func ParseSynchronizationSecret(v string) (interface{}, error) {
    result := NONE_SYNCHRONIZATIONSECRET
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_SYNCHRONIZATIONSECRET
        case "USERNAME":
            result = USERNAME_SYNCHRONIZATIONSECRET
        case "PASSWORD":
            result = PASSWORD_SYNCHRONIZATIONSECRET
        case "SECRETTOKEN":
            result = SECRETTOKEN_SYNCHRONIZATIONSECRET
        case "APPKEY":
            result = APPKEY_SYNCHRONIZATIONSECRET
        case "BASEADDRESS":
            result = BASEADDRESS_SYNCHRONIZATIONSECRET
        case "CLIENTIDENTIFIER":
            result = CLIENTIDENTIFIER_SYNCHRONIZATIONSECRET
        case "CLIENTSECRET":
            result = CLIENTSECRET_SYNCHRONIZATIONSECRET
        case "SINGLESIGNONTYPE":
            result = SINGLESIGNONTYPE_SYNCHRONIZATIONSECRET
        case "SANDBOX":
            result = SANDBOX_SYNCHRONIZATIONSECRET
        case "URL":
            result = URL_SYNCHRONIZATIONSECRET
        case "DOMAIN":
            result = DOMAIN_SYNCHRONIZATIONSECRET
        case "CONSUMERKEY":
            result = CONSUMERKEY_SYNCHRONIZATIONSECRET
        case "CONSUMERSECRET":
            result = CONSUMERSECRET_SYNCHRONIZATIONSECRET
        case "TOKENKEY":
            result = TOKENKEY_SYNCHRONIZATIONSECRET
        case "TOKENEXPIRATION":
            result = TOKENEXPIRATION_SYNCHRONIZATIONSECRET
        case "OAUTH2ACCESSTOKEN":
            result = OAUTH2ACCESSTOKEN_SYNCHRONIZATIONSECRET
        case "OAUTH2ACCESSTOKENCREATIONTIME":
            result = OAUTH2ACCESSTOKENCREATIONTIME_SYNCHRONIZATIONSECRET
        case "OAUTH2REFRESHTOKEN":
            result = OAUTH2REFRESHTOKEN_SYNCHRONIZATIONSECRET
        case "SYNCALL":
            result = SYNCALL_SYNCHRONIZATIONSECRET
        case "INSTANCENAME":
            result = INSTANCENAME_SYNCHRONIZATIONSECRET
        case "OAUTH2CLIENTID":
            result = OAUTH2CLIENTID_SYNCHRONIZATIONSECRET
        case "OAUTH2CLIENTSECRET":
            result = OAUTH2CLIENTSECRET_SYNCHRONIZATIONSECRET
        case "COMPANYID":
            result = COMPANYID_SYNCHRONIZATIONSECRET
        case "UPDATEKEYONSOFTDELETE":
            result = UPDATEKEYONSOFTDELETE_SYNCHRONIZATIONSECRET
        case "SYNCHRONIZATIONSCHEDULE":
            result = SYNCHRONIZATIONSCHEDULE_SYNCHRONIZATIONSECRET
        case "SYSTEMOFRECORD":
            result = SYSTEMOFRECORD_SYNCHRONIZATIONSECRET
        case "SANDBOXNAME":
            result = SANDBOXNAME_SYNCHRONIZATIONSECRET
        case "ENFORCEDOMAIN":
            result = ENFORCEDOMAIN_SYNCHRONIZATIONSECRET
        case "SYNCNOTIFICATIONSETTINGS":
            result = SYNCNOTIFICATIONSETTINGS_SYNCHRONIZATIONSECRET
        case "SKIPOUTOFSCOPEDELETIONS":
            result = SKIPOUTOFSCOPEDELETIONS_SYNCHRONIZATIONSECRET
        case "OAUTH2AUTHORIZATIONCODE":
            result = OAUTH2AUTHORIZATIONCODE_SYNCHRONIZATIONSECRET
        case "OAUTH2REDIRECTURI":
            result = OAUTH2REDIRECTURI_SYNCHRONIZATIONSECRET
        case "APPLICATIONTEMPLATEIDENTIFIER":
            result = APPLICATIONTEMPLATEIDENTIFIER_SYNCHRONIZATIONSECRET
        case "OAUTH2TOKENEXCHANGEURI":
            result = OAUTH2TOKENEXCHANGEURI_SYNCHRONIZATIONSECRET
        case "OAUTH2AUTHORIZATIONURI":
            result = OAUTH2AUTHORIZATIONURI_SYNCHRONIZATIONSECRET
        case "AUTHENTICATIONTYPE":
            result = AUTHENTICATIONTYPE_SYNCHRONIZATIONSECRET
        case "SERVER":
            result = SERVER_SYNCHRONIZATIONSECRET
        case "PERFORMINBOUNDENTITLEMENTGRANTS":
            result = PERFORMINBOUNDENTITLEMENTGRANTS_SYNCHRONIZATIONSECRET
        case "HARDDELETESENABLED":
            result = HARDDELETESENABLED_SYNCHRONIZATIONSECRET
        case "SYNCAGENTCOMPATIBILITYKEY":
            result = SYNCAGENTCOMPATIBILITYKEY_SYNCHRONIZATIONSECRET
        case "SYNCAGENTADCONTAINER":
            result = SYNCAGENTADCONTAINER_SYNCHRONIZATIONSECRET
        case "VALIDATEDOMAIN":
            result = VALIDATEDOMAIN_SYNCHRONIZATIONSECRET
        case "TESTREFERENCES":
            result = TESTREFERENCES_SYNCHRONIZATIONSECRET
        case "CONNECTIONSTRING":
            result = CONNECTIONSTRING_SYNCHRONIZATIONSECRET
        default:
            return 0, errors.New("Unknown SynchronizationSecret value: " + v)
    }
    return &result, nil
}
func SerializeSynchronizationSecret(values []SynchronizationSecret) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
