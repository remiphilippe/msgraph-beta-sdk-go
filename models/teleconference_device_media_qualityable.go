package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeleconferenceDeviceMediaQualityable 
type TeleconferenceDeviceMediaQualityable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAverageInboundJitter()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetAverageInboundPacketLossRateInPercentage()(*float64)
    GetAverageInboundRoundTripDelay()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetAverageOutboundJitter()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetAverageOutboundPacketLossRateInPercentage()(*float64)
    GetAverageOutboundRoundTripDelay()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetChannelIndex()(*int32)
    GetInboundPackets()(*int64)
    GetLocalIPAddress()(*string)
    GetLocalPort()(*int32)
    GetMaximumInboundJitter()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetMaximumInboundPacketLossRateInPercentage()(*float64)
    GetMaximumInboundRoundTripDelay()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetMaximumOutboundJitter()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetMaximumOutboundPacketLossRateInPercentage()(*float64)
    GetMaximumOutboundRoundTripDelay()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetMediaDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetNetworkLinkSpeedInBytes()(*int64)
    GetOutboundPackets()(*int64)
    GetRemoteIPAddress()(*string)
    GetRemotePort()(*int32)
    SetAverageInboundJitter(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetAverageInboundPacketLossRateInPercentage(value *float64)()
    SetAverageInboundRoundTripDelay(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetAverageOutboundJitter(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetAverageOutboundPacketLossRateInPercentage(value *float64)()
    SetAverageOutboundRoundTripDelay(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetChannelIndex(value *int32)()
    SetInboundPackets(value *int64)()
    SetLocalIPAddress(value *string)()
    SetLocalPort(value *int32)()
    SetMaximumInboundJitter(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetMaximumInboundPacketLossRateInPercentage(value *float64)()
    SetMaximumInboundRoundTripDelay(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetMaximumOutboundJitter(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetMaximumOutboundPacketLossRateInPercentage(value *float64)()
    SetMaximumOutboundRoundTripDelay(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetMediaDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetNetworkLinkSpeedInBytes(value *int64)()
    SetOutboundPackets(value *int64)()
    SetRemoteIPAddress(value *string)()
    SetRemotePort(value *int32)()
}
