package Errordict

import (
	"fmt"
)

var errorMap map[string]string

func init() {
	errorMap = make(map[string]string)
	errorMap["InvalidUser"] = "Please enter a valid username"
	errorMap["InvalidPassword"] = "Please enter a valid password"
	errorMap["UserNotFound"] = "User not found"
	errorMap["PermissionDenied"] = "Permission denied"
	errorMap["DatabaseError"] = "An error occurred while accessing the database"
	errorMap["InvalidInput"] = "Invalid input provided"
	errorMap["TimeoutError"] = "Operation timed out"
	errorMap["LoginFailed"] = "Login failed. Please try again."
	errorMap["SessionExpired"] = "Your session has expired. Please log in again."
	errorMap["InvalidToken"] = "Invalid token provided."
	errorMap["AccountLocked"] = "Your account has been locked. Please contact support."
	errorMap["InvalidEmail"] = "Please enter a valid email address."
	errorMap["EmailAlreadyExists"] = "This email address is already registered."
	errorMap["PasswordMismatch"] = "The passwords do not match."
	errorMap["InvalidOTP"] = "Invalid OTP. Please try again."
	errorMap["OTPExpired"] = "OTP has expired. Please request a new one."
	errorMap["UserSuspended"] = "Your account has been suspended."
	errorMap["AccountDeleted"] = "Your account has been deleted."
	errorMap["AccountInactive"] = "Your account is inactive. Please contact support."
	errorMap["InvalidRole"] = "Invalid user role."
	errorMap["RoleNotFound"] = "User role not found."
	errorMap["InvalidSession"] = "Invalid session. Please log in again."
	errorMap["SessionNotFound"] = "Session not found."
	errorMap["InvalidRequest"] = "Invalid request."
	errorMap["RequestTimeout"] = "Request timed out."
	errorMap["ServerUnavailable"] = "The server is currently unavailable. Please try again later."
	errorMap["InternalServerError"] = "Internal server error. Please try again later."
	errorMap["ServiceUnavailable"] = "Service is temporarily unavailable. Please try again later."
	errorMap["ResourceNotFound"] = "Requested resource not found."
	errorMap["ResourceConflict"] = "Resource conflict. Please resolve the conflict."
	errorMap["ResourceUnavailable"] = "Requested resource is currently unavailable."
	errorMap["InvalidCredentials"] = "Invalid credentials provided."
	errorMap["InvalidFormat"] = "Invalid format. Please provide the correct format."
	errorMap["InputValidationFailed"] = "Input validation failed. Please check your input."
	errorMap["InvalidState"] = "Invalid state. Please try again."
	errorMap["StateMismatch"] = "State mismatch. Please try again."
	errorMap["UnauthorizedAccess"] = "Unauthorized access. Please log in."
	errorMap["AccessDenied"] = "Access denied. You do not have permission to access this resource."
	errorMap["TooManyRequests"] = "Too many requests. Please try again later."
	errorMap["RateLimitExceeded"] = "Rate limit exceeded. Please try again later."
	errorMap["MaintenanceInProgress"] = "Maintenance is in progress. Please try again later."
	errorMap["ResourceLimitExceeded"] = "Resource limit exceeded."
	// Add more error codes and messages as needed
}

func GetErrorMessage(code string) (string, error) {
	errMsg, ok := errorMap[code]
	if !ok {
		return "", fmt.Errorf("error code %s not found", code)
	}
	return errMsg, nil
}

// GetErrorMessageByCode is a function to retrieve the error message based on the error code provided.
func GetErrorMessageByCode(code string) (string, error) {
	errMsg, err := GetErrorMessage(code)
	if err != nil {
		return "", err
	}
	return errMsg, nil
}

func Netxd() string {
	return "hi"
}
