package template

import (
	"fmt"
	"gen/model"
	mapset "github.com/deckarep/golang-set"
	"strings"
	"testing"
)

func TestExecuteModelTemplate(t *testing.T) {
	mapset.NewSet()
	filePath := "/Users/lony/work/orm-generator/config/user.yaml"
	outPath := "/Users/lony/work/orm-generator/out"
	model, err := model.Read(filePath)
	t.Log(model)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	err = ExecuteModelTemplate(outPath, model, "entity.tmpl")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
}
func TestExecuteRepositoryTemplate(t *testing.T) {
	mapset.NewSet()
	filePath := "/Users/lony/work/orm-generator/config/sys_button.yaml"
	outPath := "/Users/lony/work/orm-generator/out"
	model, err := model.Read(filePath)
	t.Log(model)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	fmt.Println(strings.Contains("ab", "a"))
	err = ExecuteModelTemplate(outPath+"/entity", model, "entity")
	err = ExecuteModelTemplate(outPath+"/repository", model, "repository")
	err = ExecuteModelTemplate(outPath+"/service", model, "service")
	err = ExecuteModelTemplate(outPath+"/controller", model, "controller")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
}
