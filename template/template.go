package template

import (
	"bytes"
	"fmt"
	"gen/model"
	_ "gen/packrd"
	"github.com/gobuffalo/packr/v2"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"text/template"
	"unicode"
)

var entityTemplate = template.New("gen")

// 根据文件名获取静态文件内容
func Find(fileName string) (string, error) {
	box := packr.New("", "./tmpl")
	return box.FindString(fileName)
}

func ExecuteModelTemplate(output string, obj *model.Model, templateName string) error {
	filename := filepath.Join(output, strings.Join([]string{"gen", camel2sep(obj.ModelName, "."), "go"}, "."))

	path := path.Dir(filename)
	os.MkdirAll(path, 0755)

	fd, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer fd.Close()
	if err := entityTemplate.ExecuteTemplate(fd, templateName+".tmpl", obj); err != nil {
		return err
	}
	cmd := exec.Command("gofmt", "-w", filename)
	cmd.Run()
	return nil
}

func camel2sep(s string, sep string) string {
	nameBuf := bytes.NewBuffer(nil)
	for i := range s {
		n := rune(s[i]) // always ASCII?
		if unicode.IsUpper(n) {
			if i > 0 {
				nameBuf.WriteString(sep)
			}
			n = unicode.ToLower(n)
		}
		nameBuf.WriteRune(n)
	}
	return nameBuf.String()
}

// 初始化模板内容
func init() {
	funcMap := template.FuncMap{
		"ScoreToBigCamel":   model.ScoreToBigCamel,
		"ScoreToSmallCamel": model.ScoreToSmallCamel,
		"ToHTML":            model.ToHTML,
		"Contains":          strings.Contains,
	}
	tmplFiles := []string{
		"entity.tmpl",
		"mongo/repository.tmpl",
		"mongo/service.tmpl",
		"mongo/controller.tmpl",
	}
	box := packr.New("someBoxName", "./tmpl")
	for _, s := range tmplFiles {
		bytes, err := box.Find(s)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		names := strings.Split(s, "/")
		entityTemplate.New(names[len(names)-1]).Funcs(funcMap).Parse(string(bytes))
	}
}
