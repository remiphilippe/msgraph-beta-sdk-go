package getskypeforbusinessparticipantactivityusercountswithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// 
type GetSkypeForBusinessParticipantActivityUserCountsWithPeriod struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    // 
    appSharing *int64;
    // 
    audioVideo *int64;
    // 
    dialInOut3rdParty *int64;
    // 
    im *int64;
    // 
    reportDate *string;
    // 
    reportPeriod *string;
    // 
    reportRefreshDate *string;
    // 
    web *int64;
}
// Instantiates a new getSkypeForBusinessParticipantActivityUserCountsWithPeriod and sets the default values.
func NewGetSkypeForBusinessParticipantActivityUserCountsWithPeriod()(*GetSkypeForBusinessParticipantActivityUserCountsWithPeriod) {
    m := &GetSkypeForBusinessParticipantActivityUserCountsWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
// Gets the appSharing property value. 
func (m *GetSkypeForBusinessParticipantActivityUserCountsWithPeriod) GetAppSharing()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.appSharing
    }
}
// Gets the audioVideo property value. 
func (m *GetSkypeForBusinessParticipantActivityUserCountsWithPeriod) GetAudioVideo()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.audioVideo
    }
}
// Gets the dialInOut3rdParty property value. 
func (m *GetSkypeForBusinessParticipantActivityUserCountsWithPeriod) GetDialInOut3rdParty()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.dialInOut3rdParty
    }
}
// Gets the im property value. 
func (m *GetSkypeForBusinessParticipantActivityUserCountsWithPeriod) GetIm()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.im
    }
}
// Gets the reportDate property value. 
func (m *GetSkypeForBusinessParticipantActivityUserCountsWithPeriod) GetReportDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportDate
    }
}
// Gets the reportPeriod property value. 
func (m *GetSkypeForBusinessParticipantActivityUserCountsWithPeriod) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
// Gets the reportRefreshDate property value. 
func (m *GetSkypeForBusinessParticipantActivityUserCountsWithPeriod) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
// Gets the web property value. 
func (m *GetSkypeForBusinessParticipantActivityUserCountsWithPeriod) GetWeb()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.web
    }
}
// The deserialization information for the current model
func (m *GetSkypeForBusinessParticipantActivityUserCountsWithPeriod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appSharing"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppSharing(val)
        }
        return nil
    }
    res["audioVideo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAudioVideo(val)
        }
        return nil
    }
    res["dialInOut3rdParty"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDialInOut3rdParty(val)
        }
        return nil
    }
    res["im"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIm(val)
        }
        return nil
    }
    res["reportDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReportDate(val)
        }
        return nil
    }
    res["reportPeriod"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReportPeriod(val)
        }
        return nil
    }
    res["reportRefreshDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReportRefreshDate(val)
        }
        return nil
    }
    res["web"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWeb(val)
        }
        return nil
    }
    return res
}
func (m *GetSkypeForBusinessParticipantActivityUserCountsWithPeriod) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *GetSkypeForBusinessParticipantActivityUserCountsWithPeriod) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("appSharing", m.GetAppSharing())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("audioVideo", m.GetAudioVideo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("dialInOut3rdParty", m.GetDialInOut3rdParty())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("im", m.GetIm())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("reportDate", m.GetReportDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("reportPeriod", m.GetReportPeriod())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("reportRefreshDate", m.GetReportRefreshDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("web", m.GetWeb())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the appSharing property value. 
// Parameters:
//  - value : Value to set for the appSharing property.
func (m *GetSkypeForBusinessParticipantActivityUserCountsWithPeriod) SetAppSharing(value *int64)() {
    m.appSharing = value
}
// Sets the audioVideo property value. 
// Parameters:
//  - value : Value to set for the audioVideo property.
func (m *GetSkypeForBusinessParticipantActivityUserCountsWithPeriod) SetAudioVideo(value *int64)() {
    m.audioVideo = value
}
// Sets the dialInOut3rdParty property value. 
// Parameters:
//  - value : Value to set for the dialInOut3rdParty property.
func (m *GetSkypeForBusinessParticipantActivityUserCountsWithPeriod) SetDialInOut3rdParty(value *int64)() {
    m.dialInOut3rdParty = value
}
// Sets the im property value. 
// Parameters:
//  - value : Value to set for the im property.
func (m *GetSkypeForBusinessParticipantActivityUserCountsWithPeriod) SetIm(value *int64)() {
    m.im = value
}
// Sets the reportDate property value. 
// Parameters:
//  - value : Value to set for the reportDate property.
func (m *GetSkypeForBusinessParticipantActivityUserCountsWithPeriod) SetReportDate(value *string)() {
    m.reportDate = value
}
// Sets the reportPeriod property value. 
// Parameters:
//  - value : Value to set for the reportPeriod property.
func (m *GetSkypeForBusinessParticipantActivityUserCountsWithPeriod) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
// Sets the reportRefreshDate property value. 
// Parameters:
//  - value : Value to set for the reportRefreshDate property.
func (m *GetSkypeForBusinessParticipantActivityUserCountsWithPeriod) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}
// Sets the web property value. 
// Parameters:
//  - value : Value to set for the web property.
func (m *GetSkypeForBusinessParticipantActivityUserCountsWithPeriod) SetWeb(value *int64)() {
    m.web = value
}
