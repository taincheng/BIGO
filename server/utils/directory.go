package utils

import (
	"errors"
	"os"
)

// @function: PathExists
// @description: 判断所给路径文件夹是否存在
// @param: path string
// @return: bool, error
func PathExists(path string) (bool, error) {
	stat, err := os.Stat(path)
	if err == nil {
		if stat.IsDir() {
			return true, nil
		}
		return false, errors.New(path + " : 是文件")
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
