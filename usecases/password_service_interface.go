package usecases

type PasswordServiceInterface interface {
	HashPassword(password string) (string, error)
	ComparePassword(existingPassword string, userPassword string) bool
}