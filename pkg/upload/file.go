package upload

import (
	"code.coolops.cn/blog_services/global"
	"code.coolops.cn/blog_services/pkg/util"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

// 上传
type FileType int

const TypeImage FileType = iota + 1

// 获取文件名
func GetFileName(name string) string {
	ext := GetFileExt(name)
	fileName := strings.TrimPrefix(name, ext)
	// MD5转码
	fileName = util.EncodeMD5(fileName)
	return fileName
}

// 获取文件后缀
func GetFileExt(name string) string {
	return path.Ext(name)
}

// 保存文件
func GetSavePath() string {
	return global.AppSetting.UploadSavePath
}

// 检查文件保存路径
func CheckSavePath(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsNotExist(err)
}

// 检查文件后缀
func CheckContainExt(t FileType, name string) bool {
	// 获取后缀
	ext := GetFileExt(name)
	ext = strings.ToUpper(ext)
	switch t {
	case TypeImage:
		for _, allowExt := range global.AppSetting.UploadImageAllowExt {
			if strings.ToUpper(allowExt) == ext {
				return true
			}
		}
	}
	return false
}

// 检查文件大小
func CheckMaxSize(t FileType, f multipart.File) bool {
	// 获取文件大小
	content, _ := ioutil.ReadAll(f)
	size := len(content)
	switch t {
	case TypeImage:
		if size >= global.AppSetting.UploadImageMaxSize*1024*1024 {
			return true
		}
	}
	return false
}

// 检查目录权限
func CheckPermission(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsPermission(err)
}

// 创建保存目录
func CreateSavePath(dst string, perm os.FileMode) error {
	if err := os.MkdirAll(dst, perm); err != nil {
		return err
	}
	return nil
}

// 保存文件
func SaveFile(file multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}
