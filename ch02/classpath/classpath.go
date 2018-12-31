package classpath

import (
	"os"
	"path/filepath"
)

// Classpath 结构体
type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

// Parse 方法，用于解析用户输入的Classpath
func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

// ReadClass 方法，用于查找Class，查找优先级bootClasspath > extClasspath > userClasspath
func (me *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if data, entry, err := me.bootClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	if data, entry, err := me.extClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	return me.userClasspath.readClass(className)
}

func (me *Classpath) String() string {
	return me.userClasspath.String()
}

func (me *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)

	jreLibPath := filepath.Join(jreDir, "lib", "*")
	me.bootClasspath = newWildcardEntry(jreLibPath)

	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	me.extClasspath = newWildcardEntry(jreExtPath)
}

func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	if exists("./jre") {
		return "./jre"
	}
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}
	panic("Can not find jre folder!")
}

func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func (me *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	me.userClasspath = newEntry(cpOption)
}
