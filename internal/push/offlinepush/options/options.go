package options

// Opts opts.
type Opts struct {
	Signal        *Signal
	IOSPushSound  string
	IOSBadgeCount bool
	Ex            string
	Data          string
	IsGroup       bool
	IsVoip        bool
	VoipData      VoipData
}
type VoipData struct {
	CustomType int `json:"customType"`
}

// Signal message id.
type Signal struct {
	ClientMsgID string
}
