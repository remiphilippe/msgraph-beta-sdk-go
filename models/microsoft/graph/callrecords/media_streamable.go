package callrecords

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MediaStreamable 
type MediaStreamable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAverageAudioDegradation()(*float32)
    GetAverageAudioNetworkJitter()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)
    GetAverageBandwidthEstimate()(*int64)
    GetAverageJitter()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)
    GetAveragePacketLossRate()(*float32)
    GetAverageRatioOfConcealedSamples()(*float32)
    GetAverageReceivedFrameRate()(*float32)
    GetAverageRoundTripTime()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)
    GetAverageVideoFrameLossPercentage()(*float32)
    GetAverageVideoFrameRate()(*float32)
    GetAverageVideoPacketLossRate()(*float32)
    GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLowFrameRateRatio()(*float32)
    GetLowVideoProcessingCapabilityRatio()(*float32)
    GetMaxAudioNetworkJitter()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)
    GetMaxJitter()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)
    GetMaxPacketLossRate()(*float32)
    GetMaxRatioOfConcealedSamples()(*float32)
    GetMaxRoundTripTime()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)
    GetPacketUtilization()(*int64)
    GetPostForwardErrorCorrectionPacketLossRate()(*float32)
    GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetStreamDirection()(*MediaStreamDirection)
    GetStreamId()(*string)
    GetWasMediaBypassed()(*bool)
    SetAverageAudioDegradation(value *float32)()
    SetAverageAudioNetworkJitter(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)()
    SetAverageBandwidthEstimate(value *int64)()
    SetAverageJitter(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)()
    SetAveragePacketLossRate(value *float32)()
    SetAverageRatioOfConcealedSamples(value *float32)()
    SetAverageReceivedFrameRate(value *float32)()
    SetAverageRoundTripTime(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)()
    SetAverageVideoFrameLossPercentage(value *float32)()
    SetAverageVideoFrameRate(value *float32)()
    SetAverageVideoPacketLossRate(value *float32)()
    SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLowFrameRateRatio(value *float32)()
    SetLowVideoProcessingCapabilityRatio(value *float32)()
    SetMaxAudioNetworkJitter(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)()
    SetMaxJitter(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)()
    SetMaxPacketLossRate(value *float32)()
    SetMaxRatioOfConcealedSamples(value *float32)()
    SetMaxRoundTripTime(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)()
    SetPacketUtilization(value *int64)()
    SetPostForwardErrorCorrectionPacketLossRate(value *float32)()
    SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetStreamDirection(value *MediaStreamDirection)()
    SetStreamId(value *string)()
    SetWasMediaBypassed(value *bool)()
}
