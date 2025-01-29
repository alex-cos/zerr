package zerr

// Severity defines the severity enum.
type Severity int

// Severity Constantes.
const (
	Debug Severity = iota
	Info
	Notice
	Warning
	Error
	Critical
	Fatal
)

// String returns a nice string that represents the owned entity.
func (value Severity) String() string {
	names := [...]string{
		"Debug",
		"Info",
		"Notice",
		"Warning",
		"Error",
		"Critical",
		"Fatal",
	}
	if value < Debug || value > Fatal {
		return names[0]
	}
	return names[value]
}
