package main

import "errors"

var (
	known         string = "this is just a test"
	unknown       string = "could not find the word you were looking for"
	ErrNotFound          = errors.New(unknown)
	ErrWordExists        = errors.New("cannot add word because it already exists")
)

func Search(dict map[string]string, word string) string {
	return dict[word]
}

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err

	}
	return nil
}
