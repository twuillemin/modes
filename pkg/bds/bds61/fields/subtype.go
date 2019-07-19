package fields

// Subtype is Subtype of the message
//
// Specified in Doc 9871 / Table A-2-9
type Subtype interface {
	// GetSubtype returns the subtype itself
	GetSubtype() Subtype
}
