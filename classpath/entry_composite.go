package classpath

import (
	"errors"
	"strings"
)

// CompositeEntry 类型，由Entry 数组组成
type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
	compositeEntry := []Entry{}
	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}

	return compositeEntry
}

func (me CompositeEntry) readClass(classname string) ([]byte, Entry, error) {
	for _, entry := range me {
		data, from, err := entry.readClass(classname)
		if err == nil {
			return data, from, err
		}
	}

	return nil, nil, errors.New("class not found: " + classname)
}

func (me CompositeEntry) String() string {
	strs := make([]string, len(me))
	for i, entry := range me {
		strs[i] = entry.String()
	}
	return strings.Join(strs, pathListSeparator)
}
