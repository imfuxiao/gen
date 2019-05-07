package model

import (
	"fmt"
	"testing"
)

func TestModelRead(t *testing.T) {
	filePath := "/Users/lony/work/orm-generator/config/user.yaml"
	model, err := Read(filePath)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	fmt.Println(*model)
}

func TestCamel2Name(t *testing.T) {
	name := Camel2Name("ormTeam")
	fmt.Println(name)
	name = ScoreToBigCamel("orm_team")
	fmt.Println(name)
	name = ScoreToSmallCamel("orm_team")
	fmt.Println(name)
}
