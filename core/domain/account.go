package domain

type Account struct {
	Id             int
	DocumentNumber string
	IsValid        bool
}

// IsDocumentValid
//
//	@Description: Check if document number is valid.
//	@return bool return true or false
func (a *Account) IsDocumentValid() bool {
	a.IsValid = true // mocking it is to be true
	if len(a.DocumentNumber) >= 10 {
		if a.IsValid {
			return true
		}
	}

	return false
}
