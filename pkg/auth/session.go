package auth

// func GetSession() string{
// 	return "loggeding"
// }

//can also do like this
func extractSession() string{
	return "loggiding"
}

func GetSession() string{
	return extractSession()
}