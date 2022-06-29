package system

import "os"

type AttachService struct{}

// 判断目录是否存在
func (attachService *AttachService) PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}