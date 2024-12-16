package domain

type Account struct {
	Id             int
	DocumentNumber string
	IsValid        bool
}

func (a *Account) IsDocumentValid() bool {
	a.IsValid = true // mocking it is to be true
	if len(a.DocumentNumber) >= 10 {
		if a.IsValid {
			return true
		}
	}

	return false
}
