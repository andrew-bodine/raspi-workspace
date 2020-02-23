package vibration_monitor

type VibrationMonitorState string

const (
	VibrationMonitorStateStill         VibrationMonitorState = "Still"
	VibrationMonitorStateHalfVibrating VibrationMonitorState = "HalfVibrating"
	VibrationMonitorStateVibrating     VibrationMonitorState = "Vibrating"
)
