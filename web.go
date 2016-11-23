package skt

import (
	"encoding/json"
	"html/template"
	"io"
	"net/http"
	"regexp"
)

import "path/filepath"

// Config 系统配置
type Config struct {
	TempDir string
}

// Static 设置web静态文件目录，第一个参数是在url中的路径，第二个参数是真实文件路径
func Static(urlPath string, sysPath string) {
	http.Handle(urlPath, http.StripPrefix(urlPath, http.FileServer(http.Dir(sysPath))))
}

// Render 渲染同目录下的所有文件
func Render(w http.ResponseWriter, partten string, data interface{}) error {
	tempDir := "./templates/"
	dir, file := filepath.Split(partten)
	tb := template.New("skt").Delims("{!", "!}")
	t, err := tb.ParseGlob(tempDir + dir + "*.*")
	t.Delims("{!", "!}")
	if err != nil {
		return err
	}
	return t.ExecuteTemplate(w, filepath.Base(file), data)
}

// SendJSON 返回json到客户端
func SendJSON(w http.ResponseWriter, i interface{}) error {
	return json.NewEncoder(w).Encode(i)
}

// RemoteIP 只取IP地址，去除remoteAddr的端口号
func RemoteIP(s string) string {
	re := regexp.MustCompile("\\[(.*)\\]")
	return re.FindString(s)
}

// DEJSON 解码json
func DEJSON(r io.Reader) (map[string]interface{}, error) {
	var m map[string]interface{}
	err := json.NewDecoder(r).Decode(&m)
	return m, err
}

// ENJSON 编码json
func ENJSON(i interface{}) ([]byte, error) {
	return json.Marshal(i)
}
