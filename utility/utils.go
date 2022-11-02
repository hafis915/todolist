package utility

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"todoList/pkg/setting"
)


func UploadFile(file *multipart.FileHeader) (*string,error) {
	src, err := file.Open()
	if err != nil {
		return  nil,err
	}

	defer src.Close()

	extensions := filepath.Ext(file.Filename)

	if extensions != ".txt" && extensions != ".pdf" {
		return nil, errors.New("Unexpected extensions file")
	}

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