package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type HardwareInformation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The number of charge cycles the device’s current battery has gone through. Valid values 0 to 2147483647
    batteryChargeCycles *int32;
    // The device’s current battery’s health percentage. Valid values 0 to 100
    batteryHealthPercentage *int32;
    // The serial number of the device’s current battery
    batterySerialNumber *string;
    // Cellular technology of the device
    cellularTechnology *string;
    // Returns the fully qualified domain name of the device (if any). If the device is not domain-joined, it returns an empty string.
    deviceFullQualifiedDomainName *string;
    // Local System Authority (LSA) credential guard status. . Possible values are: running, rebootRequired, notLicensed, notConfigured, virtualizationBasedSecurityNotRunning.
    deviceGuardLocalSystemAuthorityCredentialGuardState *DeviceGuardLocalSystemAuthorityCredentialGuardState;
    // Virtualization-based security hardware requirement status. Possible values are: meetHardwareRequirements, secureBootRequired, dmaProtectionRequired, hyperVNotSupportedForGuestVM, hyperVNotAvailable.
    deviceGuardVirtualizationBasedSecurityHardwareRequirementState *DeviceGuardVirtualizationBasedSecurityHardwareRequirementState;
    // Virtualization-based security status. . Possible values are: running, rebootRequired, require64BitArchitecture, notLicensed, notConfigured, doesNotMeetHardwareRequirements, other.
    deviceGuardVirtualizationBasedSecurityState *DeviceGuardVirtualizationBasedSecurityState;
    // eSIM identifier
    esimIdentifier *string;
    // Free storage space of the device.
    freeStorageSpace *int64;
    // IMEI
    imei *string;
    // IPAddressV4
    ipAddressV4 *string;
    // Encryption status of the device
    isEncrypted *bool;
    // Shared iPad
    isSharedDevice *bool;
    // Supervised mode of the device
    isSupervised *bool;
    // Manufacturer of the device
    manufacturer *string;
    // MEID
    meid *string;
    // Model of the device
    model *string;
    // String that specifies the OS edition.
    operatingSystemEdition *string;
    // Operating system language of the device
    operatingSystemLanguage *string;
    // Int that specifies the Windows Operating System ProductType. More details here https://go.microsoft.com/fwlink/?linkid=2126950. Valid values 0 to 2147483647
    operatingSystemProductType *int32;
    // Operating System Build Number on Android device
    osBuildNumber *string;
    // Phone number of the device
    phoneNumber *string;
    // Serial number.
    serialNumber *string;
    // All users on the shared Apple device
    sharedDeviceCachedUsers []SharedAppleDeviceUser;
    // SubnetAddress
    subnetAddress *string;
    // Subscriber carrier of the device
    subscriberCarrier *string;
    // Total storage space of the device.
    totalStorageSpace *int64;
    // String that specifies the specification version.
    tpmSpecificationVersion *string;
    // WiFi MAC address of the device
    wifiMac *string;
}
// Instantiates a new hardwareInformation and sets the default values.
func NewHardwareInformation()(*HardwareInformation) {
    m := &HardwareInformation{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *HardwareInformation) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the batteryChargeCycles property value. The number of charge cycles the device’s current battery has gone through. Valid values 0 to 2147483647
func (m *HardwareInformation) GetBatteryChargeCycles()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.batteryChargeCycles
    }
}
// Gets the batteryHealthPercentage property value. The device’s current battery’s health percentage. Valid values 0 to 100
func (m *HardwareInformation) GetBatteryHealthPercentage()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.batteryHealthPercentage
    }
}
// Gets the batterySerialNumber property value. The serial number of the device’s current battery
func (m *HardwareInformation) GetBatterySerialNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.batterySerialNumber
    }
}
// Gets the cellularTechnology property value. Cellular technology of the device
func (m *HardwareInformation) GetCellularTechnology()(*string) {
    if m == nil {
        return nil
    } else {
        return m.cellularTechnology
    }
}
// Gets the deviceFullQualifiedDomainName property value. Returns the fully qualified domain name of the device (if any). If the device is not domain-joined, it returns an empty string.
func (m *HardwareInformation) GetDeviceFullQualifiedDomainName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceFullQualifiedDomainName
    }
}
// Gets the deviceGuardLocalSystemAuthorityCredentialGuardState property value. Local System Authority (LSA) credential guard status. . Possible values are: running, rebootRequired, notLicensed, notConfigured, virtualizationBasedSecurityNotRunning.
func (m *HardwareInformation) GetDeviceGuardLocalSystemAuthorityCredentialGuardState()(*DeviceGuardLocalSystemAuthorityCredentialGuardState) {
    if m == nil {
        return nil
    } else {
        return m.deviceGuardLocalSystemAuthorityCredentialGuardState
    }
}
// Gets the deviceGuardVirtualizationBasedSecurityHardwareRequirementState property value. Virtualization-based security hardware requirement status. Possible values are: meetHardwareRequirements, secureBootRequired, dmaProtectionRequired, hyperVNotSupportedForGuestVM, hyperVNotAvailable.
func (m *HardwareInformation) GetDeviceGuardVirtualizationBasedSecurityHardwareRequirementState()(*DeviceGuardVirtualizationBasedSecurityHardwareRequirementState) {
    if m == nil {
        return nil
    } else {
        return m.deviceGuardVirtualizationBasedSecurityHardwareRequirementState
    }
}
// Gets the deviceGuardVirtualizationBasedSecurityState property value. Virtualization-based security status. . Possible values are: running, rebootRequired, require64BitArchitecture, notLicensed, notConfigured, doesNotMeetHardwareRequirements, other.
func (m *HardwareInformation) GetDeviceGuardVirtualizationBasedSecurityState()(*DeviceGuardVirtualizationBasedSecurityState) {
    if m == nil {
        return nil
    } else {
        return m.deviceGuardVirtualizationBasedSecurityState
    }
}
// Gets the esimIdentifier property value. eSIM identifier
func (m *HardwareInformation) GetEsimIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.esimIdentifier
    }
}
// Gets the freeStorageSpace property value. Free storage space of the device.
func (m *HardwareInformation) GetFreeStorageSpace()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.freeStorageSpace
    }
}
// Gets the imei property value. IMEI
func (m *HardwareInformation) GetImei()(*string) {
    if m == nil {
        return nil
    } else {
        return m.imei
    }
}
// Gets the ipAddressV4 property value. IPAddressV4
func (m *HardwareInformation) GetIpAddressV4()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ipAddressV4
    }
}
// Gets the isEncrypted property value. Encryption status of the device
func (m *HardwareInformation) GetIsEncrypted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEncrypted
    }
}
// Gets the isSharedDevice property value. Shared iPad
func (m *HardwareInformation) GetIsSharedDevice()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isSharedDevice
    }
}
// Gets the isSupervised property value. Supervised mode of the device
func (m *HardwareInformation) GetIsSupervised()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isSupervised
    }
}
// Gets the manufacturer property value. Manufacturer of the device
func (m *HardwareInformation) GetManufacturer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.manufacturer
    }
}
// Gets the meid property value. MEID
func (m *HardwareInformation) GetMeid()(*string) {
    if m == nil {
        return nil
    } else {
        return m.meid
    }
}
// Gets the model property value. Model of the device
func (m *HardwareInformation) GetModel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.model
    }
}
// Gets the operatingSystemEdition property value. String that specifies the OS edition.
func (m *HardwareInformation) GetOperatingSystemEdition()(*string) {
    if m == nil {
        return nil
    } else {
        return m.operatingSystemEdition
    }
}
// Gets the operatingSystemLanguage property value. Operating system language of the device
func (m *HardwareInformation) GetOperatingSystemLanguage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.operatingSystemLanguage
    }
}
// Gets the operatingSystemProductType property value. Int that specifies the Windows Operating System ProductType. More details here https://go.microsoft.com/fwlink/?linkid=2126950. Valid values 0 to 2147483647
func (m *HardwareInformation) GetOperatingSystemProductType()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.operatingSystemProductType
    }
}
// Gets the osBuildNumber property value. Operating System Build Number on Android device
func (m *HardwareInformation) GetOsBuildNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.osBuildNumber
    }
}
// Gets the phoneNumber property value. Phone number of the device
func (m *HardwareInformation) GetPhoneNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.phoneNumber
    }
}
// Gets the serialNumber property value. Serial number.
func (m *HardwareInformation) GetSerialNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.serialNumber
    }
}
// Gets the sharedDeviceCachedUsers property value. All users on the shared Apple device
func (m *HardwareInformation) GetSharedDeviceCachedUsers()([]SharedAppleDeviceUser) {
    if m == nil {
        return nil
    } else {
        return m.sharedDeviceCachedUsers
    }
}
// Gets the subnetAddress property value. SubnetAddress
func (m *HardwareInformation) GetSubnetAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subnetAddress
    }
}
// Gets the subscriberCarrier property value. Subscriber carrier of the device
func (m *HardwareInformation) GetSubscriberCarrier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subscriberCarrier
    }
}
// Gets the totalStorageSpace property value. Total storage space of the device.
func (m *HardwareInformation) GetTotalStorageSpace()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.totalStorageSpace
    }
}
// Gets the tpmSpecificationVersion property value. String that specifies the specification version.
func (m *HardwareInformation) GetTpmSpecificationVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tpmSpecificationVersion
    }
}
// Gets the wifiMac property value. WiFi MAC address of the device
func (m *HardwareInformation) GetWifiMac()(*string) {
    if m == nil {
        return nil
    } else {
        return m.wifiMac
    }
}
// The deserialization information for the current model
func (m *HardwareInformation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["batteryChargeCycles"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBatteryChargeCycles(val)
        }
        return nil
    }
    res["batteryHealthPercentage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBatteryHealthPercentage(val)
        }
        return nil
    }
    res["batterySerialNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBatterySerialNumber(val)
        }
        return nil
    }
    res["cellularTechnology"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCellularTechnology(val)
        }
        return nil
    }
    res["deviceFullQualifiedDomainName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceFullQualifiedDomainName(val)
        }
        return nil
    }
    res["deviceGuardLocalSystemAuthorityCredentialGuardState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceGuardLocalSystemAuthorityCredentialGuardState)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(DeviceGuardLocalSystemAuthorityCredentialGuardState)
            m.SetDeviceGuardLocalSystemAuthorityCredentialGuardState(&cast)
        }
        return nil
    }
    res["deviceGuardVirtualizationBasedSecurityHardwareRequirementState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceGuardVirtualizationBasedSecurityHardwareRequirementState)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(DeviceGuardVirtualizationBasedSecurityHardwareRequirementState)
            m.SetDeviceGuardVirtualizationBasedSecurityHardwareRequirementState(&cast)
        }
        return nil
    }
    res["deviceGuardVirtualizationBasedSecurityState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceGuardVirtualizationBasedSecurityState)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(DeviceGuardVirtualizationBasedSecurityState)
            m.SetDeviceGuardVirtualizationBasedSecurityState(&cast)
        }
        return nil
    }
    res["esimIdentifier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEsimIdentifier(val)
        }
        return nil
    }
    res["freeStorageSpace"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFreeStorageSpace(val)
        }
        return nil
    }
    res["imei"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImei(val)
        }
        return nil
    }
    res["ipAddressV4"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIpAddressV4(val)
        }
        return nil
    }
    res["isEncrypted"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEncrypted(val)
        }
        return nil
    }
    res["isSharedDevice"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSharedDevice(val)
        }
        return nil
    }
    res["isSupervised"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSupervised(val)
        }
        return nil
    }
    res["manufacturer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManufacturer(val)
        }
        return nil
    }
    res["meid"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMeid(val)
        }
        return nil
    }
    res["model"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModel(val)
        }
        return nil
    }
    res["operatingSystemEdition"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperatingSystemEdition(val)
        }
        return nil
    }
    res["operatingSystemLanguage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperatingSystemLanguage(val)
        }
        return nil
    }
    res["operatingSystemProductType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperatingSystemProductType(val)
        }
        return nil
    }
    res["osBuildNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsBuildNumber(val)
        }
        return nil
    }
    res["phoneNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhoneNumber(val)
        }
        return nil
    }
    res["serialNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSerialNumber(val)
        }
        return nil
    }
    res["sharedDeviceCachedUsers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSharedAppleDeviceUser() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SharedAppleDeviceUser, len(val))
            for i, v := range val {
                res[i] = *(v.(*SharedAppleDeviceUser))
            }
            m.SetSharedDeviceCachedUsers(res)
        }
        return nil
    }
    res["subnetAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubnetAddress(val)
        }
        return nil
    }
    res["subscriberCarrier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubscriberCarrier(val)
        }
        return nil
    }
    res["totalStorageSpace"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalStorageSpace(val)
        }
        return nil
    }
    res["tpmSpecificationVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTpmSpecificationVersion(val)
        }
        return nil
    }
    res["wifiMac"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWifiMac(val)
        }
        return nil
    }
    return res
}
func (m *HardwareInformation) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *HardwareInformation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("batteryChargeCycles", m.GetBatteryChargeCycles())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("batteryHealthPercentage", m.GetBatteryHealthPercentage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("batterySerialNumber", m.GetBatterySerialNumber())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("cellularTechnology", m.GetCellularTechnology())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("deviceFullQualifiedDomainName", m.GetDeviceFullQualifiedDomainName())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceGuardLocalSystemAuthorityCredentialGuardState() != nil {
        cast := m.GetDeviceGuardLocalSystemAuthorityCredentialGuardState().String()
        err := writer.WriteStringValue("deviceGuardLocalSystemAuthorityCredentialGuardState", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceGuardVirtualizationBasedSecurityHardwareRequirementState() != nil {
        cast := m.GetDeviceGuardVirtualizationBasedSecurityHardwareRequirementState().String()
        err := writer.WriteStringValue("deviceGuardVirtualizationBasedSecurityHardwareRequirementState", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceGuardVirtualizationBasedSecurityState() != nil {
        cast := m.GetDeviceGuardVirtualizationBasedSecurityState().String()
        err := writer.WriteStringValue("deviceGuardVirtualizationBasedSecurityState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("esimIdentifier", m.GetEsimIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("freeStorageSpace", m.GetFreeStorageSpace())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("imei", m.GetImei())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ipAddressV4", m.GetIpAddressV4())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isEncrypted", m.GetIsEncrypted())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isSharedDevice", m.GetIsSharedDevice())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isSupervised", m.GetIsSupervised())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("manufacturer", m.GetManufacturer())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("meid", m.GetMeid())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("model", m.GetModel())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("operatingSystemEdition", m.GetOperatingSystemEdition())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("operatingSystemLanguage", m.GetOperatingSystemLanguage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("operatingSystemProductType", m.GetOperatingSystemProductType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("osBuildNumber", m.GetOsBuildNumber())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("phoneNumber", m.GetPhoneNumber())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("serialNumber", m.GetSerialNumber())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSharedDeviceCachedUsers()))
        for i, v := range m.GetSharedDeviceCachedUsers() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("sharedDeviceCachedUsers", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("subnetAddress", m.GetSubnetAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("subscriberCarrier", m.GetSubscriberCarrier())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("totalStorageSpace", m.GetTotalStorageSpace())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("tpmSpecificationVersion", m.GetTpmSpecificationVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("wifiMac", m.GetWifiMac())
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *HardwareInformation) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the batteryChargeCycles property value. The number of charge cycles the device’s current battery has gone through. Valid values 0 to 2147483647
// Parameters:
//  - value : Value to set for the batteryChargeCycles property.
func (m *HardwareInformation) SetBatteryChargeCycles(value *int32)() {
    m.batteryChargeCycles = value
}
// Sets the batteryHealthPercentage property value. The device’s current battery’s health percentage. Valid values 0 to 100
// Parameters:
//  - value : Value to set for the batteryHealthPercentage property.
func (m *HardwareInformation) SetBatteryHealthPercentage(value *int32)() {
    m.batteryHealthPercentage = value
}
// Sets the batterySerialNumber property value. The serial number of the device’s current battery
// Parameters:
//  - value : Value to set for the batterySerialNumber property.
func (m *HardwareInformation) SetBatterySerialNumber(value *string)() {
    m.batterySerialNumber = value
}
// Sets the cellularTechnology property value. Cellular technology of the device
// Parameters:
//  - value : Value to set for the cellularTechnology property.
func (m *HardwareInformation) SetCellularTechnology(value *string)() {
    m.cellularTechnology = value
}
// Sets the deviceFullQualifiedDomainName property value. Returns the fully qualified domain name of the device (if any). If the device is not domain-joined, it returns an empty string.
// Parameters:
//  - value : Value to set for the deviceFullQualifiedDomainName property.
func (m *HardwareInformation) SetDeviceFullQualifiedDomainName(value *string)() {
    m.deviceFullQualifiedDomainName = value
}
// Sets the deviceGuardLocalSystemAuthorityCredentialGuardState property value. Local System Authority (LSA) credential guard status. . Possible values are: running, rebootRequired, notLicensed, notConfigured, virtualizationBasedSecurityNotRunning.
// Parameters:
//  - value : Value to set for the deviceGuardLocalSystemAuthorityCredentialGuardState property.
func (m *HardwareInformation) SetDeviceGuardLocalSystemAuthorityCredentialGuardState(value *DeviceGuardLocalSystemAuthorityCredentialGuardState)() {
    m.deviceGuardLocalSystemAuthorityCredentialGuardState = value
}
// Sets the deviceGuardVirtualizationBasedSecurityHardwareRequirementState property value. Virtualization-based security hardware requirement status. Possible values are: meetHardwareRequirements, secureBootRequired, dmaProtectionRequired, hyperVNotSupportedForGuestVM, hyperVNotAvailable.
// Parameters:
//  - value : Value to set for the deviceGuardVirtualizationBasedSecurityHardwareRequirementState property.
func (m *HardwareInformation) SetDeviceGuardVirtualizationBasedSecurityHardwareRequirementState(value *DeviceGuardVirtualizationBasedSecurityHardwareRequirementState)() {
    m.deviceGuardVirtualizationBasedSecurityHardwareRequirementState = value
}
// Sets the deviceGuardVirtualizationBasedSecurityState property value. Virtualization-based security status. . Possible values are: running, rebootRequired, require64BitArchitecture, notLicensed, notConfigured, doesNotMeetHardwareRequirements, other.
// Parameters:
//  - value : Value to set for the deviceGuardVirtualizationBasedSecurityState property.
func (m *HardwareInformation) SetDeviceGuardVirtualizationBasedSecurityState(value *DeviceGuardVirtualizationBasedSecurityState)() {
    m.deviceGuardVirtualizationBasedSecurityState = value
}
// Sets the esimIdentifier property value. eSIM identifier
// Parameters:
//  - value : Value to set for the esimIdentifier property.
func (m *HardwareInformation) SetEsimIdentifier(value *string)() {
    m.esimIdentifier = value
}
// Sets the freeStorageSpace property value. Free storage space of the device.
// Parameters:
//  - value : Value to set for the freeStorageSpace property.
func (m *HardwareInformation) SetFreeStorageSpace(value *int64)() {
    m.freeStorageSpace = value
}
// Sets the imei property value. IMEI
// Parameters:
//  - value : Value to set for the imei property.
func (m *HardwareInformation) SetImei(value *string)() {
    m.imei = value
}
// Sets the ipAddressV4 property value. IPAddressV4
// Parameters:
//  - value : Value to set for the ipAddressV4 property.
func (m *HardwareInformation) SetIpAddressV4(value *string)() {
    m.ipAddressV4 = value
}
// Sets the isEncrypted property value. Encryption status of the device
// Parameters:
//  - value : Value to set for the isEncrypted property.
func (m *HardwareInformation) SetIsEncrypted(value *bool)() {
    m.isEncrypted = value
}
// Sets the isSharedDevice property value. Shared iPad
// Parameters:
//  - value : Value to set for the isSharedDevice property.
func (m *HardwareInformation) SetIsSharedDevice(value *bool)() {
    m.isSharedDevice = value
}
// Sets the isSupervised property value. Supervised mode of the device
// Parameters:
//  - value : Value to set for the isSupervised property.
func (m *HardwareInformation) SetIsSupervised(value *bool)() {
    m.isSupervised = value
}
// Sets the manufacturer property value. Manufacturer of the device
// Parameters:
//  - value : Value to set for the manufacturer property.
func (m *HardwareInformation) SetManufacturer(value *string)() {
    m.manufacturer = value
}
// Sets the meid property value. MEID
// Parameters:
//  - value : Value to set for the meid property.
func (m *HardwareInformation) SetMeid(value *string)() {
    m.meid = value
}
// Sets the model property value. Model of the device
// Parameters:
//  - value : Value to set for the model property.
func (m *HardwareInformation) SetModel(value *string)() {
    m.model = value
}
// Sets the operatingSystemEdition property value. String that specifies the OS edition.
// Parameters:
//  - value : Value to set for the operatingSystemEdition property.
func (m *HardwareInformation) SetOperatingSystemEdition(value *string)() {
    m.operatingSystemEdition = value
}
// Sets the operatingSystemLanguage property value. Operating system language of the device
// Parameters:
//  - value : Value to set for the operatingSystemLanguage property.
func (m *HardwareInformation) SetOperatingSystemLanguage(value *string)() {
    m.operatingSystemLanguage = value
}
// Sets the operatingSystemProductType property value. Int that specifies the Windows Operating System ProductType. More details here https://go.microsoft.com/fwlink/?linkid=2126950. Valid values 0 to 2147483647
// Parameters:
//  - value : Value to set for the operatingSystemProductType property.
func (m *HardwareInformation) SetOperatingSystemProductType(value *int32)() {
    m.operatingSystemProductType = value
}
// Sets the osBuildNumber property value. Operating System Build Number on Android device
// Parameters:
//  - value : Value to set for the osBuildNumber property.
func (m *HardwareInformation) SetOsBuildNumber(value *string)() {
    m.osBuildNumber = value
}
// Sets the phoneNumber property value. Phone number of the device
// Parameters:
//  - value : Value to set for the phoneNumber property.
func (m *HardwareInformation) SetPhoneNumber(value *string)() {
    m.phoneNumber = value
}
// Sets the serialNumber property value. Serial number.
// Parameters:
//  - value : Value to set for the serialNumber property.
func (m *HardwareInformation) SetSerialNumber(value *string)() {
    m.serialNumber = value
}
// Sets the sharedDeviceCachedUsers property value. All users on the shared Apple device
// Parameters:
//  - value : Value to set for the sharedDeviceCachedUsers property.
func (m *HardwareInformation) SetSharedDeviceCachedUsers(value []SharedAppleDeviceUser)() {
    m.sharedDeviceCachedUsers = value
}
// Sets the subnetAddress property value. SubnetAddress
// Parameters:
//  - value : Value to set for the subnetAddress property.
func (m *HardwareInformation) SetSubnetAddress(value *string)() {
    m.subnetAddress = value
}
// Sets the subscriberCarrier property value. Subscriber carrier of the device
// Parameters:
//  - value : Value to set for the subscriberCarrier property.
func (m *HardwareInformation) SetSubscriberCarrier(value *string)() {
    m.subscriberCarrier = value
}
// Sets the totalStorageSpace property value. Total storage space of the device.
// Parameters:
//  - value : Value to set for the totalStorageSpace property.
func (m *HardwareInformation) SetTotalStorageSpace(value *int64)() {
    m.totalStorageSpace = value
}
// Sets the tpmSpecificationVersion property value. String that specifies the specification version.
// Parameters:
//  - value : Value to set for the tpmSpecificationVersion property.
func (m *HardwareInformation) SetTpmSpecificationVersion(value *string)() {
    m.tpmSpecificationVersion = value
}
// Sets the wifiMac property value. WiFi MAC address of the device
// Parameters:
//  - value : Value to set for the wifiMac property.
func (m *HardwareInformation) SetWifiMac(value *string)() {
    m.wifiMac = value
}
