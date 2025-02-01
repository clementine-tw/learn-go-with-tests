package maps

type Dictionary map[string]string

type DictionaryErr string

const (
	ErrNotFound          = DictionaryErr("could not find the key you were looking for")
	ErrWordExists        = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExists = DictionaryErr("cannot perform operation on word because it does not exists")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	item, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return item, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		d[word] = definition
		return nil
	case nil:
		return ErrWordExists
	default:
		return err
	}
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExists
	case nil:
		d[word] = definition
		return nil
	default:
		return err
	}
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExists
	case nil:
		delete(d, word)
		return nil
	default:
		return err
	}
}
