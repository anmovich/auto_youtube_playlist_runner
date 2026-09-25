package service

func (a *Auth) ClearLocked(username string) {
	a.repo.ClearLocked(username)
}
