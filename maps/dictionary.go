package maps

type Dictionary map[string]string

//first look
// func  Search(dictionary map[string]string, word string) string {
// 	return dictionary[word]
// }

// var (
// 	ErrorNotFound = errors.New("could not find the word you were looking for")
// 	ErrWordExists = errors.New("cannot add word because it already exists")
// )

const (
	ErrNotFound     = DictionaryErr("could not find the word you were looking for")
	ErrWordExists   = DictionaryErr("cannot add word because it already exists")
	ErrWordNotExist = DictionaryErr("cannot update word because not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}
	return definition, nil

}

func (d Dictionary) Add(word, definition string) error {

	_, ok := d[word]
	if ok {
		return ErrWordExists
	} else {
		d[word] = definition
		return nil
	}
}

func (d Dictionary) Update(word, definition string) error {
	_, ok := d[word]
	if !ok {
		return ErrWordNotExist
	}
	d[word] = definition
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
