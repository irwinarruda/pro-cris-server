package unit

import (
	"encoding/json"
	"testing"

	"github.com/irwinarruda/pro-cris-server/shared/models"
	"github.com/stretchr/testify/assert"
)

func TestUnmarshalWeekDay(t *testing.T) {
	var assert = assert.New(t)

	t.Run("Normal JSON", func(t *testing.T) {
		str := `{
      "weekDay": "Monday",
      "random": true
    }`
		strStruct := struct {
			WeekDay models.WeekDay `json:"weekDay"`
		}{}
		err := json.Unmarshal([]byte(str), &strStruct)
		assert.NoError(err, "Unmarshal should not throw an error")
		assert.Equal(models.Monday, strStruct.WeekDay, "WeekDay should be Monday")
	})

	t.Run("Nullish JSON", func(t *testing.T) {
		str := `{
      "random": true,
      "weekDay": null
    }`
		strStruct := struct {
			WeekDay models.WeekDay `json:"weekDay"`
		}{}
		err := json.Unmarshal([]byte(str), &strStruct)
		assert.Error(err, "Unmarshal should throw an error")
		assert.Equal(models.WeekDay(""), strStruct.WeekDay, "Unmarshal should not change WeekDay")

		strStruct1 := struct {
			WeekDay *models.WeekDay `json:"weekDay"`
		}{}
		err = json.Unmarshal([]byte(str), &strStruct1)
		assert.NoError(err, "Unmarshal should not throw an error")
		assert.Nil(strStruct1.WeekDay, "WeekDay should be nil")
	})
}

func TestUnmarshalGender(t *testing.T) {
	var assert = assert.New(t)

	t.Run("Normal JSON", func(t *testing.T) {
		str := `{
      "gender": "Female",
      "random": true
    }`
		strStruct := struct {
			Gender models.Gender `json:"gender"`
		}{}
		err := json.Unmarshal([]byte(str), &strStruct)
		assert.NoError(err, "Unmarshal should not throw an error")
		assert.Equal(models.Female, strStruct.Gender, "Gender should be Monday")
	})

	t.Run("Nullish JSON", func(t *testing.T) {
		str := `{
      "random": true,
      "gender": null
    }`
		strStruct := struct {
			Gender models.Gender `json:"gender"`
		}{}
		err := json.Unmarshal([]byte(str), &strStruct)
		assert.Error(err, "Unmarshal should throw an error")
		assert.Equal(models.Gender(""), strStruct.Gender, "Unmarshal should not change Gender")

		strStruct1 := struct {
			Gender *models.Gender `json:"gender"`
		}{}
		err = json.Unmarshal([]byte(str), &strStruct1)
		assert.NoError(err, "Unmarshal should not throw an error")
		assert.Nil(strStruct1.Gender, "Gender should be nil")
	})
}
