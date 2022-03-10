package callrecords

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// NetworkInfoable 
type NetworkInfoable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetBandwidthLowEventRatio()(*float32)
    GetBasicServiceSetIdentifier()(*string)
    GetConnectionType()(*NetworkConnectionType)
    GetDelayEventRatio()(*float32)
    GetDnsSuffix()(*string)
    GetIpAddress()(*string)
    GetLinkSpeed()(*int64)
    GetMacAddress()(*string)
    GetPort()(*int32)
    GetReceivedQualityEventRatio()(*float32)
    GetReflexiveIPAddress()(*string)
    GetRelayIPAddress()(*string)
    GetRelayPort()(*int32)
    GetSentQualityEventRatio()(*float32)
    GetSubnet()(*string)
    GetWifiBand()(*WifiBand)
    GetWifiBatteryCharge()(*int32)
    GetWifiChannel()(*int32)
    GetWifiMicrosoftDriver()(*string)
    GetWifiMicrosoftDriverVersion()(*string)
    GetWifiRadioType()(*WifiRadioType)
    GetWifiSignalStrength()(*int32)
    GetWifiVendorDriver()(*string)
    GetWifiVendorDriverVersion()(*string)
    SetBandwidthLowEventRatio(value *float32)()
    SetBasicServiceSetIdentifier(value *string)()
    SetConnectionType(value *NetworkConnectionType)()
    SetDelayEventRatio(value *float32)()
    SetDnsSuffix(value *string)()
    SetIpAddress(value *string)()
    SetLinkSpeed(value *int64)()
    SetMacAddress(value *string)()
    SetPort(value *int32)()
    SetReceivedQualityEventRatio(value *float32)()
    SetReflexiveIPAddress(value *string)()
    SetRelayIPAddress(value *string)()
    SetRelayPort(value *int32)()
    SetSentQualityEventRatio(value *float32)()
    SetSubnet(value *string)()
    SetWifiBand(value *WifiBand)()
    SetWifiBatteryCharge(value *int32)()
    SetWifiChannel(value *int32)()
    SetWifiMicrosoftDriver(value *string)()
    SetWifiMicrosoftDriverVersion(value *string)()
    SetWifiRadioType(value *WifiRadioType)()
    SetWifiSignalStrength(value *int32)()
    SetWifiVendorDriver(value *string)()
    SetWifiVendorDriverVersion(value *string)()
}
