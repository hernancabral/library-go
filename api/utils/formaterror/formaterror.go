package formaterror

import "strings"

var errorMessages = make(map[string]string)

func FormatError(errString string) map[string]string {

	if strings.Contains(errString, "isbn") {
		errorMessages["Taken_isbn"] = "There is already a book with that isbn, remember it is unique"
	}

	if strings.Contains(errString, "title") {
		errorMessages["Taken_isbn"] = "There is already a book with that title"
	}

	if len(errorMessages) > 0 {
		return errorMessages
	}

	if len(errorMessages) == 0 {
		errorMessages["Incorrect_details"] = "Incorrect Details"
		return errorMessages
	}

	return nil
}
