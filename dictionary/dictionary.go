package main

import "errors"

var known string = "this is just a test"
var unknown string = "could not find the word you were looking for"
var ErrNotFound = errors.New(unknown)

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
