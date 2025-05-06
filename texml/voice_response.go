package texml

func Voice(verbs []Element) (string, error) {
	doc, response := CreateDocument()
	if verbs != nil {
		AddAllVerbs(response, verbs)
	}
	return ToXML(doc)
}

// VoiceDial <Dial> TeXML Verb
//
// https://developers.telnyx.com/docs/voice/programmable-voice/texml-verbs/dial
//
// The <Dial> verb transfers an existing call to another destination. <Dial> will end this
// new call if: the called party does not answer, the number does not exist, or Telnyx
// receives a busy signal.
type VoiceDial struct {
	Number                        string
	Action                        string
	Method                        string
	CallerId                      string
	FromDisplayName               string
	HangupOnStar                  string
	Timeout                       string
	TimeLimit                     string
	Record                        string
	RecordingChannels             string
	RecordMaxLength               string
	RecordingStatusCallback       string
	RecordingStatusCallbackMethod string
	RecordingStatusCallbackEvent  string
	RingTone                      string
	InnerElements                 []Element
	OptionalAttributes            map[string]string
}

func (m VoiceDial) GetName() string {
	return "Dial"
}

func (m VoiceDial) GetText() string {
	return m.Number
}

func (m VoiceDial) GetAttr() (map[string]string, map[string]string) {
	paramsAttr := map[string]string{
		"Action":                        m.Action,
		"Method":                        m.Method,
		"CallerId":                      m.CallerId,
		"FromDisplayName":               m.FromDisplayName,
		"HangupOnStar":                  m.HangupOnStar,
		"Timeout":                       m.Timeout,
		"TimeLimit":                     m.TimeLimit,
		"Record":                        m.Record,
		"RecordingChannels":             m.RecordingChannels,
		"RecordMaxLength":               m.RecordMaxLength,
		"RecordingStatusCallback":       m.RecordingStatusCallback,
		"RecordingStatusCallbackMethod": m.RecordingStatusCallbackMethod,
		"RecordingStatusCallbackEvent":  m.RecordingStatusCallbackEvent,
		"RingTone":                      m.RingTone,
	}
	return m.OptionalAttributes, paramsAttr
}

func (m VoiceDial) GetInnerElements() []Element {
	return m.InnerElements
}

// VoiceNumber <Number> TeXML Child Tag
//
// Child tag within <Dial> verb.
//
// https://developers.telnyx.com/docs/voice/programmable-voice/texml-verbs/dial#number-attributes
type VoiceNumber struct {
	PhoneNumber             string
	StatusCallback          string
	StatusCallbackEvent     string
	StatusCallbackMethod    string
	Url                     string
	Method                  string
	SendDigits              string
	MachineDetection        string
	DetectionMode           string
	MachineDetectionTimeout string
	InnerElements           []Element
	OptionalAttributes      map[string]string
}

func (m VoiceNumber) GetName() string {
	return "Number"
}

func (m VoiceNumber) GetText() string {
	return m.PhoneNumber
}

func (m VoiceNumber) GetAttr() (map[string]string, map[string]string) {
	paramsAttr := map[string]string{
		"StatusCallback":          m.StatusCallback,
		"StatusCallbackEvent":     m.StatusCallbackEvent,
		"StatusCallbackMethod":    m.StatusCallbackMethod,
		"Url":                     m.Url,
		"Method":                  m.Method,
		"SendDigits":              m.SendDigits,
		"MachineDetection":        m.MachineDetection,
		"DetectionMode":           m.DetectionMode,
		"MachineDetectionTimeout": m.MachineDetectionTimeout,
	}
	return m.OptionalAttributes, paramsAttr
}

func (m VoiceNumber) GetInnerElements() []Element {
	return m.InnerElements
}

// VoiceSip <Sip> TeXML Child Tag
//
// Child tag within <Dial> verb.
//
// https://developers.telnyx.com/docs/voice/programmable-voice/texml-verbs/dial#sip-attributes
type VoiceSip struct {
	SipUrl                  string
	Username                string
	Password                string
	StatusCallback          string
	StatusCallbackEvent     string
	StatusCallbackMethod    string
	Url                     string
	Method                  string
	MachineDetection        string
	DetectionMode           string
	MachineDetectionTimeout string
	InnerElements           []Element
	OptionalAttributes      map[string]string
}

func (m VoiceSip) GetName() string {
	return "Sip"
}

func (m VoiceSip) GetText() string {
	return m.SipUrl
}

func (m VoiceSip) GetAttr() (map[string]string, map[string]string) {
	paramsAttr := map[string]string{
		"Username":                m.Username,
		"Password":                m.Password,
		"StatusCallback":          m.StatusCallback,
		"StatusCallbackEvent":     m.StatusCallbackEvent,
		"StatusCallbackMethod":    m.StatusCallbackMethod,
		"Url":                     m.Url,
		"Method":                  m.Method,
		"MachineDetection":        m.MachineDetection,
		"DetectionMode":           m.DetectionMode,
		"MachineDetectionTimeout": m.MachineDetectionTimeout,
	}
	return m.OptionalAttributes, paramsAttr
}

func (m VoiceSip) GetInnerElements() []Element {
	return m.InnerElements
}

// VoiceConference <Conference> TeXML Verb
//
// https://developers.telnyx.com/docs/voice/programmable-voice/texml-verbs/conference
//
// The <Dial> verb's <Conference> noun allows you to connect to a conference room. Much
// like how the <Number> noun allows you to connect to another phone number, the
// <Conference> noun allows you to connect to a named conference room and talk with the
// other callers who have also connected to that room. Conference is commonly used as a
// container for calls when implementing hold, transfer, and barge. If the specified
// conference name does not exist, a new conference will be created.
type VoiceConference struct {
	Name                          string
	Muted                         string
	StartConferenceOnEnter        string
	EndConferenceOnExit           string
	MaxParticipants               string
	Beep                          string
	Record                        string
	RecordBeep                    string
	RecordingStatusCallback       string
	RecordingStatusCallbackEvent  string
	RecordingStatusCallbackMethod string
	RecordingTimeout              string
	Trim                          string
	StatusCallback                string
	StatusCallbackEvent           string
	StatusCallbackMethod          string
	WaitUrl                       string
	WaitMethod                    string
	InnerElements                 []Element
	OptionalAttributes            map[string]string
}

func (m VoiceConference) GetName() string {
	return "Conference"
}

func (m VoiceConference) GetText() string {
	return m.Name
}

func (m VoiceConference) GetAttr() (map[string]string, map[string]string) {
	paramsAttr := map[string]string{
		"Muted":                         m.Muted,
		"StartConferenceOnEnter":        m.StartConferenceOnEnter,
		"EndConferenceOnExit":           m.EndConferenceOnExit,
		"MaxParticipants":               m.MaxParticipants,
		"Beep":                          m.Beep,
		"Record":                        m.Record,
		"RecordBeep":                    m.RecordBeep,
		"RecordingStatusCallback":       m.RecordingStatusCallback,
		"RecordingStatusCallbackEvent":  m.RecordingStatusCallbackEvent,
		"RecordingStatusCallbackMethod": m.RecordingStatusCallbackMethod,
		"RecordingTimeout":              m.RecordingTimeout,
		"Trim":                          m.Trim,
		"StatusCallback":                m.StatusCallback,
		"StatusCallbackEvent":           m.StatusCallbackEvent,
		"StatusCallbackMethod":          m.StatusCallbackMethod,
		"WaitUrl":                       m.WaitUrl,
		"WaitMethod":                    m.WaitMethod,
	}
	return m.OptionalAttributes, paramsAttr
}

func (m VoiceConference) GetInnerElements() []Element {
	return m.InnerElements
}

// VoiceEnqueue <Enqueue> TeXML Verb
//
// https://developers.telnyx.com/docs/voice/programmable-voice/texml-verbs/enqueue
//
// The <Enqueue> verb enqueues the current call in a call queue.
type VoiceEnqueue struct {
	Action             string
	Method             string
	WaitUrl            string
	WaitUrlMethod      string
	InnerElements      []Element
	OptionalAttributes map[string]string
}

func (m VoiceEnqueue) GetName() string {
	return "Enqueue"
}

func (m VoiceEnqueue) GetText() string {
	return ""
}

func (m VoiceEnqueue) GetAttr() (map[string]string, map[string]string) {
	paramsAttr := map[string]string{
		"Action":        m.Action,
		"Method":        m.Method,
		"WaitUrl":       m.WaitUrl,
		"WaitUrlMethod": m.WaitUrlMethod,
	}
	return m.OptionalAttributes, paramsAttr
}

func (m VoiceEnqueue) GetInnerElements() []Element {
	return m.InnerElements
}

// VoiceGather <Gather> TeXML Verb
//
// https://developers.telnyx.com/docs/voice/programmable-voice/texml-verbs/gather
//
// The <Gather> verb collects DTMF tones during a call. <Say> can be nested within
// <Gather> to create an interactive IVR with text-to-speech.
type VoiceGather struct {
	Action              string
	Timeout             string
	FinishOnKey         string
	NumDigits           string
	Language            string
	ValidDigits         string
	InvalidDigitsAction string
	MinDigits           string
	MaxDigits           string
	InnerElements       []Element
	OptionalAttributes  map[string]string
}

func (m VoiceGather) GetName() string {
	return "Gather"
}

func (m VoiceGather) GetText() string {
	return ""
}

func (m VoiceGather) GetAttr() (map[string]string, map[string]string) {
	paramsAttr := map[string]string{
		"Action":              m.Action,
		"Timeout":             m.Timeout,
		"FinishOnKey":         m.FinishOnKey,
		"NumDigits":           m.NumDigits,
		"Language":            m.Language,
		"ValidDigits":         m.ValidDigits,
		"InvalidDigitsAction": m.InvalidDigitsAction,
		"MinDigits":           m.MinDigits,
		"MaxDigits":           m.MaxDigits,
	}
	return m.OptionalAttributes, paramsAttr
}

func (m VoiceGather) GetInnerElements() []Element {
	return m.InnerElements
}

// VoiceHangup <Hangup> TeXML Verb
//
// https://developers.telnyx.com/docs/voice/programmable-voice/texml-verbs/hangup
//
// Ends the call
type VoiceHangup struct {
	InnerElements      []Element
	OptionalAttributes map[string]string
}

func (m VoiceHangup) GetName() string {
	return "Hangup"
}

func (m VoiceHangup) GetText() string {
	return ""
}

func (m VoiceHangup) GetAttr() (map[string]string, map[string]string) {
	return m.OptionalAttributes, nil
}

func (m VoiceHangup) GetInnerElements() []Element {
	return m.InnerElements
}

// VoiceLeave <Leave> TeXML Verb
//
// https://developers.telnyx.com/docs/voice/programmable-voice/texml-verbs/leave
//
// The <Leave> verb removes a call from the queue and continues with the next verb after
// the original <Enqueue>. The <Leave> verb doesn't support any attributes.
type VoiceLeave struct {
	InnerElements      []Element
	OptionalAttributes map[string]string
}

func (m VoiceLeave) GetName() string {
	return "Leave"
}

func (m VoiceLeave) GetText() string {
	return ""
}

func (m VoiceLeave) GetAttr() (map[string]string, map[string]string) {
	return m.OptionalAttributes, nil
}

func (m VoiceLeave) GetInnerElements() []Element {
	return m.InnerElements
}

// VoicePause <Pause> TeXML Verb
//
// https://developers.telnyx.com/docs/voice/programmable-voice/texml-verbs/pause
//
// The <Pause> verb waits silently for a specified number of seconds or one second by
// default. No nouns can be nested within <Pause>, and a self-closing tag must be used.
type VoicePause struct {
	Length             string
	InnerElements      []Element
	OptionalAttributes map[string]string
}

func (m VoicePause) GetName() string {
	return "Pause"
}

func (m VoicePause) GetText() string {
	return ""
}

func (m VoicePause) GetAttr() (map[string]string, map[string]string) {
	paramsAttr := map[string]string{
		"Length": m.Length,
	}
	return m.OptionalAttributes, paramsAttr
}

func (m VoicePause) GetInnerElements() []Element {
	return m.InnerElements
}

// VoicePlay <Play> TeXML Verb
//
// https://developers.telnyx.com/docs/voice/programmable-voice/texml-verbs/play
//
// The <Play> verb plays an MP3 or WAV audio file, which Telnyx fetches back to the caller
// from the URL you configure. Alternatively, specify mediaStorage="true" to fetch a file
// you previously uploaded to Telnyx using media storage APIs. When mediaStorage="true" is
// used, the verb expects a media_name instead of a URL. <Play> can be used independently
// as a verb or nested within <Gather> as a noun to play an audio file while waiting for
// DTMF tones.
type VoicePlay struct {
	Url                string
	Loop               string
	MediaStorage       string
	InnerElements      []Element
	OptionalAttributes map[string]string
}

func (m VoicePlay) GetName() string {
	return "Play"
}

func (m VoicePlay) GetText() string {
	return m.Url
}

func (m VoicePlay) GetAttr() (map[string]string, map[string]string) {
	paramsAttr := map[string]string{
		"Loop":         m.Loop,
		"MediaStorage": m.MediaStorage,
	}
	return m.OptionalAttributes, paramsAttr
}

func (m VoicePlay) GetInnerElements() []Element {
	return m.InnerElements
}

// VoiceRecord <Record> TeXML Verb
//
// https://developers.telnyx.com/docs/voice/programmable-voice/texml-verbs/record
//
// The <Record> verb creates an audio file with the call audio. If a
// recordingStatusCallback is provided, Telnyx will deliver the URL for the recording to
// that address once the call has ended. Recording URLs are valid for 10 minutes after the
// call has ended. All recordings are also available via the Telnyx Mission Control Portal
type VoiceRecord struct {
	Action                        string
	Method                        string
	FinishOnKey                   string
	Timeout                       string
	MaxLength                     string
	PlayBeep                      string
	Trim                          string
	Channels                      string
	RecordingStatusCallback       string
	RecordingStatusCallbackMethod string
	InnerElements                 []Element
	OptionalAttributes            map[string]string
}

func (m VoiceRecord) GetName() string {
	return "Record"
}

func (m VoiceRecord) GetText() string {
	return ""
}

func (m VoiceRecord) GetAttr() (map[string]string, map[string]string) {
	paramsAttr := map[string]string{
		"Action":                        m.Action,
		"Method":                        m.Method,
		"FinishOnKey":                   m.FinishOnKey,
		"Timeout":                       m.Timeout,
		"MaxLength":                     m.MaxLength,
		"PlayBeep":                      m.PlayBeep,
		"Trim":                          m.Trim,
		"Channels":                      m.Channels,
		"RecordingStatusCallback":       m.RecordingStatusCallback,
		"RecordingStatusCallbackMethod": m.RecordingStatusCallbackMethod,
	}
	return m.OptionalAttributes, paramsAttr
}

func (m VoiceRecord) GetInnerElements() []Element {
	return m.InnerElements
}

// VoiceRedirect <Redirect> TeXML Verb
//
// https://developers.telnyx.com/docs/voice/programmable-voice/texml-verbs/redirect
//
// The <Redirect> verb transfers control of a call to the TeXML document to another TeXML application. This is useful to create a tree structure of TeXML files for different applications. No nouns can be nested within <Redirect>
type VoiceRedirect struct {
	Url                string
	Method             string
	InnerElements      []Element
	OptionalAttributes map[string]string
}

func (m VoiceRedirect) GetName() string {
	return "Redirect"
}

func (m VoiceRedirect) GetText() string {
	return m.Url
}

func (m VoiceRedirect) GetAttr() (map[string]string, map[string]string) {
	paramsAttr := map[string]string{
		"Method": m.Method,
	}
	return m.OptionalAttributes, paramsAttr
}

func (m VoiceRedirect) GetInnerElements() []Element {
	return m.InnerElements
}

// VoiceRefer <Refer> TeXML Verb
//
// https://developers.telnyx.com/docs/voice/programmable-voice/texml-verbs/refer
//
// The <Refer> verb in Telnyx allows you to transfer a phone call to another SIP
// infrastructure during a TeXML call. You can initiate it at any point during the call.
// When you use the Refer verb, Telnyx will replace the original call with a new call to
// the external system you specify, effectively transferring the call to that system.
type VoiceRefer struct {
	Action             string
	Method             string
	InnerElements      []Element
	OptionalAttributes map[string]string
}

func (m VoiceRefer) GetName() string {
	return "Refer"
}

func (m VoiceRefer) GetText() string {
	return ""
}

func (m VoiceRefer) GetAttr() (map[string]string, map[string]string) {
	paramsAttr := map[string]string{
		"Action": m.Action,
		"Method": m.Method,
	}
	return m.OptionalAttributes, paramsAttr
}

func (m VoiceRefer) GetInnerElements() []Element {
	return m.InnerElements
}

// VoiceReferSip <Sip> TeXML Noun used in <Refer>
//
// https://developers.telnyx.com/docs/voice/programmable-voice/texml-verbs/refer#examples
type VoiceReferSip struct {
	SipUrl             string
	InnerElements      []Element
	OptionalAttributes map[string]string
}

func (m VoiceReferSip) GetName() string {
	return "Sip"
}

func (m VoiceReferSip) GetText() string {
	return m.SipUrl
}

func (m VoiceReferSip) GetAttr() (map[string]string, map[string]string) {
	return m.OptionalAttributes, nil
}

func (m VoiceReferSip) GetInnerElements() []Element {
	return m.InnerElements
}

// VoiceReject <Reject> TeXML Verb
//
// https://developers.telnyx.com/docs/voice/programmable-voice/texml-verbs/reject
//
// The <Reject> verb rejects a call to your Telnyx number. It is effectively an exit
// statement from the current document, as there is no way to return to any instructions
// listed after the <Reject> verb. If placed as the very first verb in an incoming call,
// <Reject> will prevent the call from being answered and will incur no cost. If placed
// elsewhere in the call, the call will hang up but will be charged up to that point.
type VoiceReject struct {
	Reason             string
	InnerElements      []Element
	OptionalAttributes map[string]string
}

func (m VoiceReject) GetName() string {
	return "Reject"
}

func (m VoiceReject) GetText() string {
	return ""
}

func (m VoiceReject) GetAttr() (map[string]string, map[string]string) {
	paramsAttr := map[string]string{
		"Reason": m.Reason,
	}
	return m.OptionalAttributes, paramsAttr
}

func (m VoiceReject) GetInnerElements() []Element {
	return m.InnerElements
}

// VoiceSay <Say> TeXML Verb
//
// https://developers.telnyx.com/docs/voice/programmable-voice/texml-verbs/say
//
// The <Say> verb speaks the text specified back to the caller, enabling text-to-speech
// for any application.
type VoiceSay struct {
	Message            string
	Voice              string
	Language           string
	Loop               string
	InnerElements      []Element
	OptionalAttributes map[string]string
}

func (m VoiceSay) GetName() string {
	return "Say"
}

func (m VoiceSay) GetText() string {
	return m.Message
}

func (m VoiceSay) GetAttr() (map[string]string, map[string]string) {
	paramsAttr := map[string]string{
		"Voice":    m.Voice,
		"Language": m.Language,
		"Loop":     m.Loop,
	}
	return m.OptionalAttributes, paramsAttr
}

func (m VoiceSay) GetInnerElements() []Element {
	return m.InnerElements
}

// VoiceStop <Stop> TeXML Verb
//
// https://developers.telnyx.com/docs/voice/programmable-voice/texml-verbs/stop
//
// The <Stop> verb stops the instruction specified by noun on a call.
type VoiceStop struct {
	InnerElements      []Element
	OptionalAttributes map[string]string
}

func (m VoiceStop) GetName() string {
	return "Stop"
}

func (m VoiceStop) GetText() string {
	return ""
}

func (m VoiceStop) GetAttr() (map[string]string, map[string]string) {
	return m.OptionalAttributes, nil
}

func (m VoiceStop) GetInnerElements() []Element {
	return m.InnerElements
}

// VoiceStream <Stream> TeXML Noun
//
// https://developers.telnyx.com/docs/voice/programmable-voice/texml-verbs/stream
//
// The <Stream> instruction starts streaming the media from a call to a specific WebSocket
// address in near-real-time. Audio will be delivered as base64-encoded RTP payloads
// (no headers), wrapped in JSON payloads.
type VoiceStream struct {
	Url                string
	Track              string
	Name               string
	Codec              string
	BidirectionalMode  string
	BidirectionalCodec string
	InnerElements      []Element
	OptionalAttributes map[string]string
}

func (m VoiceStream) GetName() string {
	return "Stream"
}

func (m VoiceStream) GetText() string {
	return ""
}

func (m VoiceStream) GetAttr() (map[string]string, map[string]string) {
	paramsAttr := map[string]string{
		"Url":                m.Url,
		"Track":              m.Track,
		"Name":               m.Name,
		"Codec":              m.Codec,
		"BidirectionalMode":  m.BidirectionalMode,
		"BidirectionalCodec": m.BidirectionalCodec,
	}
	return m.OptionalAttributes, paramsAttr
}

func (m VoiceStream) GetInnerElements() []Element {
	return m.InnerElements
}

// VoiceStart <Start> TeXML Verb
//
// https://developers.telnyx.com/docs/voice/programmable-voice/texml-verbs/stream
type VoiceStart struct {
	InnerElements      []Element
	OptionalAttributes map[string]string
}

func (m VoiceStart) GetName() string {
	return "Start"
}

func (m VoiceStart) GetText() string {
	return ""
}

func (m VoiceStart) GetAttr() (map[string]string, map[string]string) {
	return m.OptionalAttributes, nil
}

func (m VoiceStart) GetInnerElements() []Element {
	return m.InnerElements
}
