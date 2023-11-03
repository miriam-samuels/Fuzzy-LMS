package helper

func FindUserByToken(token string) string {
	var userId string
	claims, valid := VerifyJWT(token)
	if valid {
		userId = claims.UserId
		return userId
	} else {
		return ""
	}
}
