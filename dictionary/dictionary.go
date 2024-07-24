package dictionary

import (
	"fmt"
)

type Dictionary map[string]string

const (
	ErrWordNotFound     = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("this word already exists")
	ErrWordDoesNotExist = DictionaryErr("this word does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(s string) (string, error) {
	definition, ok := d[s]
	if !ok {
		return "", ErrWordNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(key, value string) (string, error) {
	_, ok := d[key]
	if ok {
		return "", ErrWordExists
	}
	d[key] = value
	return fmt.Sprintf("key %q, value: %q added successfully", key, value), nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrWordNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}
