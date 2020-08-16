package classpath

import (
	"os"
	"strings"
)

//存放路径分隔符
const pathListSeparator = string(os.PathListSeparator)

//类路径
type Entry interface {
	//寻找和加载class 文件
	// className class文件的相对路径，路径之间用斜 线（/）分隔，文件名有.class后缀。比如要读取java.lang.Object类，传 入的参数应该是java/lang/Object.class。
	// []byte 读取到的字节数据
	// Entry 最终定位到class文件的Entry
	// error 错误信息
	ReadClass(className string) ([]byte, Entry, error)

	//相当于Java中的toString（），用于返回变量 的字符串表示
	String() string
}

// 根据参数创建不同类型的Entry实例
func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		// 组合多个 ：分割
		return newCompositeEntry(path)
	} else if strings.HasSuffix(path, "*") {
		// 通配符*
		return newWildcardEntry(path)
	} else if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		// jar zip 包
		return newZipEntry(path)
	} else {
		// 目录
		return newDirEntry(path)
	}
}








