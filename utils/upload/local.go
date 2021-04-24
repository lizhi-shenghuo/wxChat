package upload

import (
	"errors"
	"fmt"
	"getaway/utils"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"time"
)

type Local struct{}

//@object: *Local
//@function: UploadFile
//@description: 上传文件

func (*Local) UploadFile(file *multipart.FileHeader) (string, string, error) {
	// 读取文件后缀
	ext := path.Ext(file.Filename)
	// 读取文件名并加密
	name := strings.TrimSuffix(file.Filename, ext)
	name = utils.MD5V([]byte(name))
	// 拼接新文件名
	filename := name + "_" + time.Now().Format("20060102150405") + ext

	// 拼接路径和文件名
	p := "images" + "/" + filename

	f, openError := file.Open() // 读取文件
	if openError != nil {
		fmt.Printf("function os.Create() Filed, error(%v)", openError.Error())
		return "", "", errors.New("function file.Open() Filed, err:" + openError.Error())
	}
	defer f.Close() // 创建文件 defer 关闭

	out, createErr := os.Create(p)
	if createErr != nil {
		fmt.Printf("function os.Create() Filed, error(%v)", createErr.Error())
		return "", "", errors.New("function os.Create() Filed, err:" + createErr.Error())
	}
	defer out.Close() // 创建文件 defer 关闭

	_, copyErr := io.Copy(out, f) // 传输（拷贝）文件
	if copyErr != nil {
		return "", "", errors.New("function io.Copy() Filed, err:" + copyErr.Error())
	}
	return p, filename, nil
}

//@object: *Local
//@function: DeleteFile
//@description: 删除文件
func (*Local) DeleteFile(key string) error {
	p := "images" + "/" + key
	if strings.Contains(p, "images"+"/"+key) {
		if err := os.Remove(p); err != nil {
			return errors.New("本地文件删除失败, err:" + err.Error())
		}
	}
	return nil
}
