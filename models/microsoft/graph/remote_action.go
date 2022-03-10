package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the deviceManagement singleton.
type RemoteAction int

const (
    UNKNOWN_REMOTEACTION RemoteAction = iota
    FACTORYRESET_REMOTEACTION
    REMOVECOMPANYDATA_REMOTEACTION
    RESETPASSCODE_REMOTEACTION
    REMOTELOCK_REMOTEACTION
    ENABLELOSTMODE_REMOTEACTION
    DISABLELOSTMODE_REMOTEACTION
    LOCATEDEVICE_REMOTEACTION
    REBOOTNOW_REMOTEACTION
    RECOVERPASSCODE_REMOTEACTION
    CLEANWINDOWSDEVICE_REMOTEACTION
    LOGOUTSHAREDAPPLEDEVICEACTIVEUSER_REMOTEACTION
    QUICKSCAN_REMOTEACTION
    FULLSCAN_REMOTEACTION
    WINDOWSDEFENDERUPDATESIGNATURES_REMOTEACTION
    FACTORYRESETKEEPENROLLMENTDATA_REMOTEACTION
    UPDATEDEVICEACCOUNT_REMOTEACTION
    AUTOMATICREDEPLOYMENT_REMOTEACTION
    SHUTDOWN_REMOTEACTION
    ROTATEBITLOCKERKEYS_REMOTEACTION
    ROTATEFILEVAULTKEY_REMOTEACTION
    GETFILEVAULTKEY_REMOTEACTION
    SETDEVICENAME_REMOTEACTION
    ACTIVATEDEVICEESIM_REMOTEACTION
)

func (i RemoteAction) String() string {
    return []string{"UNKNOWN", "FACTORYRESET", "REMOVECOMPANYDATA", "RESETPASSCODE", "REMOTELOCK", "ENABLELOSTMODE", "DISABLELOSTMODE", "LOCATEDEVICE", "REBOOTNOW", "RECOVERPASSCODE", "CLEANWINDOWSDEVICE", "LOGOUTSHAREDAPPLEDEVICEACTIVEUSER", "QUICKSCAN", "FULLSCAN", "WINDOWSDEFENDERUPDATESIGNATURES", "FACTORYRESETKEEPENROLLMENTDATA", "UPDATEDEVICEACCOUNT", "AUTOMATICREDEPLOYMENT", "SHUTDOWN", "ROTATEBITLOCKERKEYS", "ROTATEFILEVAULTKEY", "GETFILEVAULTKEY", "SETDEVICENAME", "ACTIVATEDEVICEESIM"}[i]
}
func ParseRemoteAction(v string) (interface{}, error) {
    result := UNKNOWN_REMOTEACTION
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_REMOTEACTION
        case "FACTORYRESET":
            result = FACTORYRESET_REMOTEACTION
        case "REMOVECOMPANYDATA":
            result = REMOVECOMPANYDATA_REMOTEACTION
        case "RESETPASSCODE":
            result = RESETPASSCODE_REMOTEACTION
        case "REMOTELOCK":
            result = REMOTELOCK_REMOTEACTION
        case "ENABLELOSTMODE":
            result = ENABLELOSTMODE_REMOTEACTION
        case "DISABLELOSTMODE":
            result = DISABLELOSTMODE_REMOTEACTION
        case "LOCATEDEVICE":
            result = LOCATEDEVICE_REMOTEACTION
        case "REBOOTNOW":
            result = REBOOTNOW_REMOTEACTION
        case "RECOVERPASSCODE":
            result = RECOVERPASSCODE_REMOTEACTION
        case "CLEANWINDOWSDEVICE":
            result = CLEANWINDOWSDEVICE_REMOTEACTION
        case "LOGOUTSHAREDAPPLEDEVICEACTIVEUSER":
            result = LOGOUTSHAREDAPPLEDEVICEACTIVEUSER_REMOTEACTION
        case "QUICKSCAN":
            result = QUICKSCAN_REMOTEACTION
        case "FULLSCAN":
            result = FULLSCAN_REMOTEACTION
        case "WINDOWSDEFENDERUPDATESIGNATURES":
            result = WINDOWSDEFENDERUPDATESIGNATURES_REMOTEACTION
        case "FACTORYRESETKEEPENROLLMENTDATA":
            result = FACTORYRESETKEEPENROLLMENTDATA_REMOTEACTION
        case "UPDATEDEVICEACCOUNT":
            result = UPDATEDEVICEACCOUNT_REMOTEACTION
        case "AUTOMATICREDEPLOYMENT":
            result = AUTOMATICREDEPLOYMENT_REMOTEACTION
        case "SHUTDOWN":
            result = SHUTDOWN_REMOTEACTION
        case "ROTATEBITLOCKERKEYS":
            result = ROTATEBITLOCKERKEYS_REMOTEACTION
        case "ROTATEFILEVAULTKEY":
            result = ROTATEFILEVAULTKEY_REMOTEACTION
        case "GETFILEVAULTKEY":
            result = GETFILEVAULTKEY_REMOTEACTION
        case "SETDEVICENAME":
            result = SETDEVICENAME_REMOTEACTION
        case "ACTIVATEDEVICEESIM":
            result = ACTIVATEDEVICEESIM_REMOTEACTION
        default:
            return 0, errors.New("Unknown RemoteAction value: " + v)
    }
    return &result, nil
}
func SerializeRemoteAction(values []RemoteAction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
