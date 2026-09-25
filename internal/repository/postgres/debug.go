package postgres

func (a *Auth) ClearLocked(username string) {
	a.clearLoginAttemps(username)
	a.setNULLLockedUntil(username)
}
