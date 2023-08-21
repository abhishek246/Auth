package login

import (
	"auth/internal/db"
	"auth/internal/models"
)

func AuthenticateLogin(email_id, password, clientId string) (string, bool) {

	var credentials models.Credientials

	subquery := db.DB.Select("id").Where("email_id = ? AND tenant_id = ?", email_id, clientId).Limit(1).Table("users")

	// TODO: Hash the passwprd here to make match this the logic
	result := db.DB.Where("user_id = (?) AND password = ? AND deleted_at IS NULL", subquery, password).
		First(&credentials)
	if result.Error == nil {
		return credentials.UserId, true
	}
	return "", false
}
