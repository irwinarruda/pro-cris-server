package unit

import (
	"os"
	"regexp"
	"testing"

	"github.com/irwinarruda/pro-cris-server/libs/proenv"
	"github.com/stretchr/testify/assert"
)

func TestLoadEnv(t *testing.T) {
	var assert = assert.New(t)

	t.Run("Nil pointer error", func(t *testing.T) {
		var arg *interface{} = nil
		err := proenv.LoadEnv(arg)
		assert.Error(err, "it should throw an error with nil pointer arg")
	})

	t.Run("No string error", func(t *testing.T) {
		var arg = &struct {
			NoString int32
		}{}
		var want = regexp.MustCompile("type string")
		err := proenv.LoadEnv(arg)
		assert.Error(err, "it should throw an error with no string struct params")
		assert.Regexp(want, err.Error(), "it should throw an error with no string struct params")
	})

	t.Run("No env tag error", func(t *testing.T) {
		var arg = &struct {
			NoTag    string
			NoEnvTag string `other:"Other tag"`
		}{}
		var want = regexp.MustCompile("env tag")
		err := proenv.LoadEnv(arg)
		assert.Error(err, "it should throw an error with no `env` tag")
		assert.Regexp(want, err.Error(), "it should throw an error with no `env` tag")
	})

	t.Run("Fill env struct", func(t *testing.T) {
		var arg = &struct {
			Env1 string `env:"ENV_NUM_1"`
			Env2 string `env:"ENV_NUM_2"`
		}{}
		want1 := "Num 1"
		want2 := "Num 2"
		os.Setenv("ENV_NUM_1", want1)
		os.Setenv("ENV_NUM_2", want2)
		err := proenv.LoadEnv(arg)
		assert.NoError(err, "it should not throw an error with good args")
		assert.Equal(want1, arg.Env1, "it should fill up a good arg")
		assert.Equal(want2, arg.Env2, "it should fill up a good arg")
	})
}
