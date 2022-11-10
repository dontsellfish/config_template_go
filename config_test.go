package config_go

import (
	"fmt"
	"testing"
)

func TestConfigUtil_Update_Rollback(t *testing.T) {
	mod, err := NewConfigUtil("data/test_config.json")
	if err != nil {
		t.Error(err)
	}
	_ = mod.Rollback()
	err = mod.Update("example", "updated")
	if err != nil {
		t.Error(err)
	}

	if mod.Data.Example != "updated" {
		t.Error(fmt.Sprintf("modifyed wrong: %s != %s", mod.Data.Example, "updated"))
	}

	err = mod.Rollback()
	if err != nil {
		t.Error(err)
	}
	if mod.Data.Example != "hello world" {
		t.Error(fmt.Sprintf("modifyed wrong: %s != %s", mod.Data.Example, "hello world"))
	}

	err = mod.Update("sub", "sub_data", 314)
	if err != nil {
		t.Error(err)
	}
	if mod.Data.Sub.SubData != 314 {
		t.Error(fmt.Sprintf("modifyed wrong: %d != %d", mod.Data.Sub.SubData, 314))
	}

	err = mod.Rollback()
	if err != nil {
		t.Error(err)
	}
	err = mod.Update("sub", "sub_data", 314)
	if err != nil {
		t.Error(err)
	}
	if mod.Data.Sub.SubData != 314 {
		t.Error(fmt.Sprintf("modifyed wrong: %d != %d", mod.Data.Sub.SubData, 314))
	}

	err = mod.Rollback()
	if err != nil {
		t.Error(err)
	}
	err = mod.Update("missing", "whaaaaat")
	if err == nil {
		t.Error("updated, tho hadn't")
	}

}
