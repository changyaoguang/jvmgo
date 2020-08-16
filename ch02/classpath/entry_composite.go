package classpath

import (
	"errors"
	"strings"
)

type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
	compositeEntry := []Entry{}
	for _, p := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(p)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}

func (self CompositeEntry) ReadClass(className string) ([]byte, Entry, error) {
	for _, entry := range self {
		data, en, err := entry.ReadClass(className)
		if err == nil {
			return data, en, err
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}

func (self CompositeEntry) String() string {
	s := make([]string, len(self))
	for i, entry := range self {
		s[i] = entry.String()
	}
	return strings.Join(s, pathListSeparator)
}
