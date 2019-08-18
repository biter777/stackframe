package stackframe

// A Printer formats error messages.
//
// The most common implementation of Printer is the one provided by package fmt
// during Printf (as of Go 1.13). Localization packages such as golang.org/x/text/message
// typically provide their own implementations.
type Printer interface {
	// Print appends args to the message output.
	Print(args ...interface{})

	// Printf writes a formatted string.
	Printf(format string, args ...interface{})

	// Detail reports whether error detail is requested.
	// After the first call to Detail, all text written to the Printer
	// is formatted as additional detail, or ignored when
	// detail has not been requested.
	// If Detail returns false, the caller can avoid printing the detail at all.
	Detail() bool
}
