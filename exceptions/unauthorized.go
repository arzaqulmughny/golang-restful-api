package exceptions

type UnauthorizedException struct {
	Error string
}

func NewUnauthorizedException(error string) UnauthorizedException {
	return UnauthorizedException{Error: error}
}
