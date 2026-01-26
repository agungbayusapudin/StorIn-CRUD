package response

// Daftar Response Code Standar Aplikasi
const (
	// Success
	CodeSuccess = "0000"

	// Auth Errors (AUTH)
	CodeAuthInvalid   = "AUTH-001"
	CodeAuthExpired   = "AUTH-002"
	CodeAuthForbidden = "AUTH-003"

	// Validation Errors (VAL)
	CodeBadRequest    = "VAL-001"
	CodeInvalidFormat = "VAL-002"

	// Database Errors (DB)
	CodeNotFound = "DB-001"
	CodeConflict = "DB-002"

	CodeMethodNotAllowed = "SYC-003"

	// Server Errors (SYS)
	CodeInternalError = "9999"
)

// Mapping Code ke Pesan (Default Message)
var messageMap = map[string]string{
	CodeSuccess:          "Success",
	CodeAuthInvalid:      "Invalid or missing token",
	CodeAuthExpired:      "Token has expired",
	CodeAuthForbidden:    "You don't have permission to access this resource",
	CodeBadRequest:       "Invalid request body or missing required fields",
	CodeInvalidFormat:    "Data format is not supported",
	CodeNotFound:         "Data not found",
	CodeConflict:         "Data already exists",
	CodeMethodNotAllowed: "Method not allowed",
	CodeInternalError:    "Internal server error, please try again later",
}

// GetMessage mengembalikan pesan berdasarkan kode
func GetMessage(code string) string {
	msg, tx := messageMap[code]
	if !tx {
		return "Unknown Status"
	}
	return msg
}
