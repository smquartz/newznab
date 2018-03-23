package newznab

import "fmt"

// NError describes an error defined by the newznab specification
type NError struct {
	Code        int
	Description string
}

// Error implements the error interface for the NError type
func (n NError) Error() string {
	return n.Description
}

// errors defined by the newznab spec
var (
	ErrIncorrectUserCredentials            = NError{Code: 100, Description: "Incorrect user credentials"}
	ErrAccountSuspended                    = NError{Code: 101, Description: "Account suspended"}
	ErrInsufficientPrivilegesNotAuthorized = NError{Code: 102, Description: "Insufficient privileges/not authorized"}
	ErrRegistrationDenied                  = NError{Code: 103, Description: "Registration denied"}
	ErrRegistrationsClosed                 = NError{Code: 104, Description: "Registrations are closed"}
	ErrInvalidRegistrationEmailTaken       = NError{Code: 105, Description: "Invalid registration (Email Address Taken)"}
	ErrInvalidRegistrationBadEmail         = NError{Code: 106, Description: "Invalid registration (Email Address Bad Format)"}
	ErrRegistrationFailedDataError         = NError{Code: 107, Description: "Registration Failed (Data error)"}
	ErrMissingParameter                    = NError{Code: 200, Description: "Missing parameter"}
	ErrIncorrectParameter                  = NError{Code: 201, Description: "Incorrect parameter"}
	ErrNoSuchFunction                      = NError{Code: 202, Description: "No such function. (Function not defined in this specification)"}
	ErrFunctionNotAvailable                = NError{Code: 203, Description: "Function not available. (Optional function is not implemented)"}
	ErrNoSuchItem                          = NError{Code: 300, Description: "No such item"}
	ErrItemAlreadyExists                   = NError{Code: 300, Description: "Item already exists"}
	ErrUnknownError                        = NError{Code: 900, Description: "Unknown error"}
	ErrAPIDisabled                         = NError{Code: 910, Description: "API Disabled"}
)

var nerrors = map[int]NError{
	100: ErrIncorrectUserCredentials,
	101: ErrAccountSuspended,
	102: ErrInsufficientPrivilegesNotAuthorized,
	103: ErrRegistrationDenied,
	104: ErrRegistrationsClosed,
	105: ErrInvalidRegistrationEmailTaken,
	106: ErrInvalidRegistrationBadEmail,
	107: ErrRegistrationFailedDataError,
	200: ErrMissingParameter,
	201: ErrIncorrectParameter,
	202: ErrNoSuchFunction,
	203: ErrFunctionNotAvailable,
	300: ErrNoSuchItem,
	900: ErrUnknownError,
	910: ErrAPIDisabled,
}

// NErrorRange describes an error range defined by the newznab spec
// it is used for when we do not know the meaning of the specific error code,
// but can work out its general category from the range it falls into
type NErrorRange struct {
	// the lower bound of the error range
	Min int
	// the upper bound of the error range
	Max int
	// the description of the error range
	Description string
	// the code of the actual unknown error
	Code int
}

// Error implements the error interface for the NErorrRange type
func (n NErrorRange) Error() string {
	if n.Code != 0 {
		if !n.codeWithin(n.Code) {
			return fmt.Sprintf("unknown newznab error")
		}
		return fmt.Sprintf("unknown newznab error %d: %s", n.Code, n.Description)
	}
	return fmt.Sprintf("unknown newznab error: %s", n.Description)
}

// withCode returns a copy of the called-on NErrorRange, with its Code field
// set to the provided value
func (n NErrorRange) withCode(code int) NErrorRange {
	n.Code = code
	return n
}

// codeWithin returns whether an arbitrary error code is within the range
// defined in a NErrorRange
func (n NErrorRange) codeWithin(code int) bool {
	return n.Code >= n.Min && n.Code <= n.Max
}

// error ranges defined within the newznab spec
var (
	ErrUnspecifiedAccountOrUserCredentials = NErrorRange{Min: 100, Max: 199, Description: "issue with user credentials or account"}
	ErrUnspecifiedAPICall                  = NErrorRange{Min: 200, Max: 299, Description: "API-call specific error"}
	ErrUnspecifiedContent                  = NErrorRange{Min: 300, Max: 399, Description: "issue with content"}
	ErrUnspecifiedOther                    = NErrorRange{Min: 900, Max: 999, Description: "other error"}
)

// nerrorRangeFromCode takes the code of an unknown error, and returns the
// corresponding NErrorRange
func nerrorRangeFromCode(code int) NErrorRange {
	ranges := []NErrorRange{ErrUnspecifiedAccountOrUserCredentials, ErrUnspecifiedAPICall, ErrUnspecifiedContent, ErrUnspecifiedOther}
	for _, ner := range ranges {
		if ner.codeWithin(code) {
			return ner
		}
	}
	return ErrUnspecifiedOther
}
