package options

// Opts opts.
type Opts struct {
	Signal        *Signal
	IOSPushSound  string
	IOSBadgeCount bool
	Ex            string
	Data          string
	IsGroup       bool
}

// Signal message id.
type Signal struct {
	ClientMsgID string
}
