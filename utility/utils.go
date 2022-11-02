package utility

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"todoList/pkg/setting"
)


func UploadFile(file *multipart.FileHeader) (*string,error) {
	src, err := file.Open()
	if err != nil {
		return  nil,err
	}

	defer src.Close()

	dst, err := os.Create(fmt.Sprintf("%s/%s", setting.FileHandlerSetting.Filepath, file.Filename))
	if err != nil {
		return nil,err
	}
	defer dst.Close()

	if _, err := io.Copy(dst,src) ; err != nil {
		return nil, err
	}

	return &file.Filename, nil
}