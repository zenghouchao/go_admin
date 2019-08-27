package handler

import (
	"fmt"
	"net/http"
)

func LoginInterceptor(hFunc func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sess, _ := globalSessions.SessionStart(w, r)
		defer sess.SessionRelease(w)
		sessionInfo := sess.Get("userInfo")
		fmt.Println("session is :\n", sessionInfo)
		if sessionInfo == nil {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

		hFunc(w, r)
	})
}
