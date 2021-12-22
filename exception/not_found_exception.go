package exception

type NotFoundException struct {
	Error string
}

func NewNotFoundException(error string) NotFoundException {
	return NotFoundException{Error: error}
}


