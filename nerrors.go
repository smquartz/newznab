package newznab

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
