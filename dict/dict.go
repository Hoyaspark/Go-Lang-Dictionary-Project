package dict

import "errors"

// Dictionary type
type Dictionary map[string]string

var (
	errNotFound      error = errors.New("not Found")
	errDuplicateWord error = errors.New("duplicate Word")
)

func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]

	if exists {
		return value, nil
	}

	return "", errNotFound
}

func (d Dictionary) Add(key, value string) error {
	_, exists := d[key]

	if exists {
		return errDuplicateWord
	}

	d[key] = value

	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, exists := d[word]

	if !exists {
		return errNotFound
	}

	d[word] = definition

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
