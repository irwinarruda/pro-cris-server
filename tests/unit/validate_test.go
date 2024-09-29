package unit

import (
	"testing"

	"github.com/irwinarruda/pro-cris-server/shared/configs"
	"github.com/irwinarruda/pro-cris-server/shared/models"
	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	var assert = assert.New(t)
	var validate = configs.GetValidate(map[string][]string{})

	t.Run("WeekDay Happy Path", func(t *testing.T) {
		correct := struct {
			Monday    models.WeekDay `validate:"weekday"`
			Tuesday   models.WeekDay `validate:"required,weekday"`
			Wednesday models.WeekDay `validate:"required,weekday"`
			Thursday  string         `validate:"weekday"`
			Friday    string         `validate:"weekday"`
			Saturday  string         `validate:"weekday"`
			Sunday    string         `validate:"weekday"`
		}{
			Monday:    models.Monday,
			Tuesday:   models.Tuesday,
			Wednesday: models.Wednesday,
			Thursday:  models.Thursday.String(),
			Friday:    models.Friday.String(),
			Saturday:  models.Saturday.String(),
			Sunday:    models.Sunday.String(),
		}
		err := validate.Struct(correct)
		assert.NoError(err, "it should not throw an error with correct weekdays")
	})

	t.Run("WeekDay Error Path", func(t *testing.T) {
		wrong := struct {
			Other  models.WeekDay `validate:"weekday"`
			Other1 string         `validate:"weekday"`
			Other2 int            `validate:"weekday"`
			Other3 bool           `validate:"weekday"`
		}{
			Other:  "other",
			Other1: "monday",
			Other2: 1,
			Other3: true,
		}
		err := validate.Struct(wrong)
		assert.Error(err, "it should throw an error with wrong weekdays")
	})

	t.Run("Gender Happy Path", func(t *testing.T) {
		correct := struct {
			Male   models.Gender `validate:"gender"`
			Female string        `validate:"gender"`
		}{
			Male:   models.Male,
			Female: models.Female.String(),
		}
		err := validate.Struct(correct)
		assert.NoError(err, "it should not throw an error with correct gender")
	})

	t.Run("Gender Error Path", func(t *testing.T) {
		wrong := struct {
			Other  models.Gender `validate:"gender"`
			Other1 string        `validate:"gender"`
			Other2 int           `validate:"gender"`
			Other3 bool          `validate:"gender"`
		}{
			Other:  "other",
			Other1: "male",
			Other2: 1,
			Other3: true,
		}
		err := validate.Struct(wrong)
		assert.Error(err, "it should throw an error with wrong gender")
	})
}
