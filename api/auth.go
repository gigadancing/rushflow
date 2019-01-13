package main

import (
	"net/http"
	"rushflow/api/defs"
	"rushflow/api/sessions"
)

var (
	HEADER_FIELD_SESSION = "X-Session-Id"
	HEADER_FIELD_UNAME   = "X-User-Name"
)

//
func validateUserSession(r *http.Request) bool {
	sid := r.Header.Get(HEADER_FIELD_SESSION)
	if len(sid) == 0 {
		return false
	}

	username, ok := sessions.IsSessionExpired(sid)
	if ok {
		return false
	}
	r.Header.Add(HEADER_FIELD_UNAME, username)
	return true
}

//
func ValidateUser(w http.ResponseWriter, r *http.Request) bool {
	username := r.Header.Get(HEADER_FIELD_UNAME)
	if len(username) == 0 {
		sendErrorResponse(w, defs.ErrorNotAuthUser)
		return false
	}
	return true
}
