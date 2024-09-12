package utils

func CheckPasswordSimple(reqPassword string, userPassword string) bool {
	return reqPassword == userPassword
}
