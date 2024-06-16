package unit

import (
	"testing"

	"github.com/irwinarruda/pro-cris-server/shared/configs"
	"github.com/irwinarruda/pro-cris-server/shared/models"
)

func TestValidateWeekDay(t *testing.T) {
	validate := configs.GetValidate()
	correct := struct {
		Monday    string `validate:"weekday"`
		Tuesday   string `validate:"weekday"`
		Wednesday string `validate:"weekday"`
		Thursday  string `validate:"weekday"`
		Friday    string `validate:"weekday"`
		Saturday  string `validate:"weekday"`
		Sunday    string `validate:"weekday"`
	}{
		Monday:    models.Monday,
		Tuesday:   models.Tuesday,
		Wednesday: models.Wednesday,
		Thursday:  models.Thursday,
		Friday:    models.Friday,
		Saturday:  models.Saturday,
		Sunday:    models.Sunday,
	}
	err := validate.Struct(correct)
	if err != nil {
		t.Fatalf("it should not throw an error with correct weekdays")
	}

	wrong := struct {
		Other  string `validate:"weekday"`
		Other1 string `validate:"weekday"`
		Other2 int    `validate:"weekday"`
		Other3 bool   `validate:"weekday"`
	}{
		Other:  "other",
		Other1: "monday",
		Other2: 1,
		Other3: true,
	}
	err = validate.Struct(wrong)
	if err == nil {
		t.Fatalf("it should throw an error with wrong weekdays")
	}
}
