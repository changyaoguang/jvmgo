package classpath

import (
	"os"
	"path/filepath"
)

type ClassPath struct {
	bootClassPath Entry
	extClassPath  Entry
	userClassPath Entry
}

func Parse(jreOption, cpOptrion string) *ClassPath {
	cp := &ClassPath{}
	cp.parseBootAndExtClassPath(jreOption)
	cp.parseUserClassPath(cpOptrion)
	return cp

}

func (self *ClassPath) ReadClass(path string) ([]byte, Entry, error) {
	className := path + ".class"
	if data, entry, err := self.bootClassPath.ReadClass(className); err == nil {
		return data, entry, err
	}
	if data, entry, err := self.extClassPath.ReadClass(className); err == nil {
		return data, entry, err
	}
	return self.userClassPath.ReadClass(className)

}

func (self *ClassPath) String() string {
	return self.userClassPath.String()
}

func (self *ClassPath) parseBootAndExtClassPath(jreOption string) {
	jreDir := getJreDir(jreOption)
	// jre/lib/*
	jreLibPath := filepath.Join(jreDir,"lib","*")
	self.bootClassPath = newWildcardEntry(jreLibPath)
	// jre/lib/ext/*
	jreLibExtPath := filepath.Join(jreDir,"lib","ext","*")
	self.extClassPath = newWildcardEntry(jreLibExtPath)
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
	panic("can not find jre folder")

}

func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func (self *ClassPath) parseUserClassPath(cpOptrion string) {
	if cpOptrion == "" {
		cpOptrion = "."
	}
	self.userClassPath = newEntry(cpOptrion)
}
