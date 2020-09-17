package main

import (
	"testing"
)

func TestActionExec(t *testing.T) {
	invalidPutActions := []EtcdActionPut{
		EtcdActionPut{EtcdActionBase{EtcdActPut}, "", ""},
		EtcdActionPut{EtcdActionBase{EtcdActPut}, "key1", ""},
		EtcdActionPut{EtcdActionBase{EtcdActPut}, "", "val1"},
	}
	for _, action := range invalidPutActions {
		_, err := action.Exec()
		if err == nil {
			t.Errorf("invalid put actions get no error")
		}
	}
	validPutActions := []EtcdActionPut{
		EtcdActionPut{EtcdActionBase{EtcdActPut}, "key1", "val1"},
	}
	for _, action := range validPutActions {
		_, err := action.Exec()
		if err != nil {
			t.Errorf("valid put actions get error")
		}
	}

	invalidGetActions := []EtcdActionGet{
		EtcdActionGet{EtcdActionBase{EtcdActGet}, "", ""},
	}
	for _, action := range invalidGetActions {
		_, err := action.Exec()
		if err == nil {
			t.Errorf("invalid get actions get no error")
		}
	}
	validGetActions := []EtcdActionGet{
		EtcdActionGet{EtcdActionBase{EtcdActGet}, "key1", ""},
		EtcdActionGet{EtcdActionBase{EtcdActGet}, "key1", "endRange"},
	}
	for _, action := range validGetActions {
		_, err := action.Exec()
		if err != nil {
			t.Errorf("valid get actions get error")
		}
	}

	invalidDeleteActions := []EtcdActionDelete{
		EtcdActionDelete{EtcdActionBase{EtcdActDelete}, "", ""},
	}
	for _, action := range invalidDeleteActions {
		_, err := action.Exec()
		if err == nil {
			t.Errorf("invalid get actions get no error")
		}
	}
	validDeleteActions := []EtcdActionDelete{
		EtcdActionDelete{EtcdActionBase{EtcdActDelete}, "key1", ""},
		EtcdActionDelete{EtcdActionBase{EtcdActDelete}, "key1", "endRange"},
	}
	for _, action := range validDeleteActions {
		_, err := action.Exec()
		if err != nil {
			t.Errorf("valid get actions get error")
		}
	}
}
