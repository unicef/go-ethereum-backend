package userservice

import "log"

//IsUsernameRegistered checks whether username is already registered or not
func (db *Service) IsUsernameRegistered(username string) bool {
	var count int32
	err := db.QueryRow(
		"SELECT count(*) as counter FROM user WHERE username=?",
		username,
	).Scan(&count)
	log.Println(count)
	return err == nil && count > 0
}
