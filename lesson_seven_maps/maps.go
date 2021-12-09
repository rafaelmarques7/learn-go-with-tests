package maps

// ======> Error logic
const (
	ErrCanNotUpdate = ErrDictionary("can not update word which does not exist")
	ErrWordNotFound = ErrDictionary("word not found")
	ErrWordExists   = ErrDictionary("word already exists and thus will not be modified")
)

type ErrDictionary string

func (e ErrDictionary) Error() string {
	return string(e)
}

// ======> Business logic

// Dictionary of words -> meaning (string -> string)
type Dictionary map[string]string

// Searches a dictionary for a word and
// returns its meaning or error if the word does not exist
func (d Dictionary) Search(word string) (string, error) {
	definition, wordFound := d[word]
	if !wordFound {
		return "", ErrWordNotFound
	}
	return definition, nil
}

// Adds word and meaning to dictionary,
// returns error if the word already exists or nil if the operation was successfull
func (d Dictionary) AddWord(word, meaning string) error {
	_, err := d.Search(word)

	switch err {
	case ErrWordNotFound:
		d[word] = meaning
		return nil
	case nil:
		return ErrWordExists
	default:
		// Handle unexpected errors by returning them
		return err
	}
}

// Updates a word, if it exists, or returns error if it does not
func (d Dictionary) UpdateWord(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrWordNotFound:
		return ErrCanNotUpdate
	case nil:
		d[word] = definition
		return nil
	default:
		return err
	}
}

// Deletes word, if it exists, or returns error if it does not
func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	switch err {
	case ErrWordNotFound:
		return ErrWordNotFound
	case nil:
		delete(d, word)
		return nil
	default:
		return err
	}
}
