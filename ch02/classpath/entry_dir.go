package classpath

import "io/ioutil"
import "path/filepath"

// DirEntry 结构体，目录类型
type DirEntry struct {
	absDir string
}

func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir}
}

func (me *DirEntry) readClass(classname string) ([]byte, Entry, error) {
	fileName := filepath.Join(me.absDir, classname)
	data, err := ioutil.ReadFile(fileName)
	return data, me, err
}

func (me *DirEntry) String() string {
	return me.absDir
}
