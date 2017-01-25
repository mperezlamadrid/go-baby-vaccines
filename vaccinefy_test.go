package vaccinefy

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestHasVaccinesToApply(t *testing.T) {
	type hasVaccinesTestCase struct {
		DOB         time.Time
		ReqDate     time.Time
		Result      bool
		Description string
	}

	cases := []hasVaccinesTestCase{
		hasVaccinesTestCase{
			DOB:         time.Date(2016, 11, 1, 0, 0, 0, 0, time.UTC),
			ReqDate:     time.Date(2017, 1, 12, 0, 0, 0, 0, time.UTC),
			Result:      true,
			Description: "application of vaccines from 2 months",
		},
		hasVaccinesTestCase{
			DOB:         time.Date(2016, 11, 1, 0, 0, 0, 0, time.UTC),
			ReqDate:     time.Date(2021, 11, 1, 0, 0, 0, 0, time.UTC),
			Result:      true,
			Description: "application of vaccines from 5 years",
		},
		hasVaccinesTestCase{
			DOB:         time.Date(2016, 11, 1, 0, 0, 0, 0, time.UTC),
			ReqDate:     time.Date(2017, 2, 1, 0, 0, 0, 0, time.UTC),
			Result:      false,
			Description: "application of vaccines from 3 months",
		},
	}

	for _, testCase := range cases {
		result := HasVaccinesToApply(testCase.DOB, testCase.ReqDate)
		assert.Equal(t, testCase.Result, result, testCase.Description)
	}
}

func TestGetDatesToApplyVaccines(t *testing.T) {
	type getDatesTestCase struct {
		DOB         time.Time
		Result      InfoVaccines
		Description string
	}

	cases := []getDatesTestCase{
		getDatesTestCase{
			DOB: time.Date(2016, 11, 1, 0, 0, 0, 0, time.UTC),
			Result: InfoVaccines{
				Date:       time.Date(2016, 11, 1, 0, 0, 0, 0, time.UTC),
				References: []vaccine{{"Hepatitis B", "0"}, {"Tuberculosis B.C.G", "unica"}},
			},
			Description: "application of vaccines from 7 months",
		},
	}

	for _, testCase := range cases {
		result := GetDatesToApplyVaccines(testCase.DOB)
		assert.Equal(t, testCase.Result, result[0], testCase.Description)
	}
}

func TestGetVaccinesReference(t *testing.T) {
	type getVaccinesReferenceTestCase struct {
		DOB         time.Time
		Months      int
		Result      InfoVaccines
		Description string
	}

	cases := []getVaccinesReferenceTestCase{
		getVaccinesReferenceTestCase{
			DOB:    time.Date(2016, 11, 1, 0, 0, 0, 0, time.UTC),
			Months: 7,
			Result: InfoVaccines{
				Date:       time.Date(2016, 6, 1, 0, 0, 0, 0, time.UTC),
				References: []vaccine{{"Influenza", "2"}},
			},
			Description: "application of vaccines from 7 months",
		},
		getVaccinesReferenceTestCase{
			DOB:         time.Date(2016, 11, 1, 0, 0, 0, 0, time.UTC),
			Months:      1,
			Result:      InfoVaccines{},
			Description: "application of vaccines from 1 month",
		},
	}

	for _, testCase := range cases {
		result := GetVaccinesReference(testCase.DOB, testCase.Months)
		assert.Equal(t, testCase.Result.References, result.References, testCase.Description)
	}
}
