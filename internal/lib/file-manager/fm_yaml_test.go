package file_manager

import (
	"testing"
)

type TestStruct struct {
	FieldOne string `yaml:"field_one"`
	FieldTwo string `yaml:"field_two"`
}

func TestYamlManager_Read(t *testing.T) {
	fm := NewYamlManager()
	ts := TestStruct{}
	err := fm.Read(`test.yaml`, &ts)
	if err != nil {
		t.Error(err)
	}
	t.Log(ts)
}

func TestYamlManager_Write(t *testing.T) {
	fm := NewYamlManager()
	ts := TestStruct{
		FieldOne: "Sample_1",
		FieldTwo: "Sample_2",
	}
	err := fm.Write(`test.yaml`, &ts)
	if err != nil {
		t.Error(err)
	}
	t.Log(ts)
	tsr := TestStruct{}
	err = fm.Read(`test.yaml`, &tsr)
	if err != nil {
		t.Error(err)
	}
	t.Log(tsr)
}
