package cerrors

import "net/http"

/**
 * Write your custom errors here
 */

var (
	// Using this kind of msg lets you i18n in the frontend
	EmailAlreadyUsed = New("email_already_used", 99, http.StatusConflict)
)
