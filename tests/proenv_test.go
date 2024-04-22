package tests

import (
	"os"
	"regexp"
	"testing"

	"github.com/irwinarruda/pro-cris-server/libs/proenv"
)

func TestLoadEnvNilPointerErr(t *testing.T) {
	var arg *interface{} = nil
	err := proenv.LoadEnv(arg)
	if err == nil {
		t.Fatal("it should throw an error with nil pointer arg")
	}
}

func TestLoadEnvNoStringErr(t *testing.T) {
	var arg = &struct {
		NoString int32
	}{}
	var want = regexp.MustCompile("type string")
	err := proenv.LoadEnv(arg)
	if err == nil || !want.MatchString(err.Error()) {
		t.Fatal("it should throw an error with no string struct params")
	}
}

func TestLoadEnvNoEnvTagErr(t *testing.T) {
	var arg = &struct {
		NoTag    string
		NoEnvTag string `other:"Other tag"`
	}{}
	var want = regexp.MustCompile("env tag")
	err := proenv.LoadEnv(arg)
	if err == nil || !want.MatchString(err.Error()) {
		t.Fatal("it should throw an error with no `env` tag")
	}
}

func TestLoadEnvFillEnvStruct(t *testing.T) {
	var arg = &struct {
		Env1 string `env:"ENV_NUM_1"`
		Env2 string `env:"ENV_NUM_2"`
	}{}
	want1 := "Num 1"
	want2 := "Num 2"
	os.Setenv("ENV_NUM_1", want1)
	os.Setenv("ENV_NUM_2", want2)
	err := proenv.LoadEnv(arg)
	if err != nil {
		t.Fatal("it should not throw an error with good args")
	}
	if arg.Env1 != want1 || arg.Env2 != want2 {
		t.Fatal("it should fill up a good arg")
	}
}
