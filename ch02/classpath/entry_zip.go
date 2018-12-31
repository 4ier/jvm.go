package classpath

import "archive/zip"
import "errors"
import "io/ioutil"
import "path/filepath"

// ZipEntry 结构体，压缩文件类型
type ZipEntry struct {
	absPath string
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath}
}

func (me *ZipEntry) readClass(classname string) ([]byte, Entry, error) {
	r, err := zip.OpenReader(me.absPath)
	if err != nil {
		return nil, nil, err
	}

	defer r.Close()

	for _, f := range r.File {
		if f.Name == classname {
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}

			defer rc.Close()
			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}

			return data, me, err
		}
	}

	return nil, nil, errors.New("class not found: " + classname)
}

func (me *ZipEntry) String() string {
	return me.absPath
}
