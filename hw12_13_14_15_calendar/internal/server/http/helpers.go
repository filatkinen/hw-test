package internalhttp

import (
	"path"
	"strings"
)

func getUserID(path string) (userID string, err error) {
	_, tail := shiftPath(path)
	userID, _ = shiftPath(tail)
	if len(userID) == 0 {
		return "", ErrUserIDNotSet
	}
	if len(userID) != 36 {
		return "", ErrBadUserID
	}
	return userID, nil
}

func shiftPath(p string) (head, tail string) {
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	if i <= 0 {
		return p[1:], "/"
	}
	return p[1:i], p[i:]
}
