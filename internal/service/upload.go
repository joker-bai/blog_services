package service

import (
	"code.coolops.cn/blog_services/global"
	"code.coolops.cn/blog_services/pkg/upload"
	"errors"
	"mime/multipart"
	"os"
)

// 上传文件
type FileInfo struct {
	Name      string
	AccessUrl string
}

func (s *Service) UploadFile(fileType upload.FileType, file multipart.File, fileHeader multipart.FileHeader) (*FileInfo, error) {
	// 获取文件名
	fileName := upload.GetFileName(fileHeader.Filename)
	// 获取上传路径
	uploadSavePath := upload.GetSavePath()
	// 目标目录
	dst := uploadSavePath + "/" + fileName
	// 检查后缀名
	if !upload.CheckContainExt(fileType, fileName) {
		return nil, errors.New("file suffix is not allowed")
	}
	// 检查上传路径，如果不存在，就创建
	if upload.CheckSavePath(uploadSavePath) {
		if err := upload.CreateSavePath(uploadSavePath, os.ModePerm); err != nil {
			return nil, errors.New("failed to create save path")
		}
	}
	// 检查文件大小
	if upload.CheckMaxSize(fileType, file) {
		return nil, errors.New("exceeded maximum file limit")
	}
	// 检查权限
	if upload.CheckPermission(uploadSavePath) {
		return nil, errors.New("insufficient file permissions")
	}
	// 保存文件
	if err := upload.SaveFile(fileHeader, dst); err != nil {
		return nil, err
	}
	accessUrl := global.AppSetting.UploadServerUri + fileName
	return &FileInfo{
		Name:      fileName,
		AccessUrl: accessUrl,
	}, nil
}
