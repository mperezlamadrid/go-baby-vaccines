package vaccinefy

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestHasVaccineApplication(t *testing.T) {
	type hasApplicationTestCase struct {
		dob         time.Time
		reqDate     time.Time
		response    bool
		description string
	}

	cases := []hasApplicationTestCase{
		hasApplicationTestCase{
			dob:         time.Date(2016, 11, 1, 0, 0, 0, 0, time.UTC),
			reqDate:     time.Date(2017, 1, 12, 0, 0, 0, 0, time.UTC),
			response:    true,
			description: "application of vaccines from 2 months",
		},
		hasApplicationTestCase{
			dob:         time.Date(2016, 11, 1, 0, 0, 0, 0, time.UTC),
			reqDate:     time.Date(2021, 11, 1, 0, 0, 0, 0, time.UTC),
			response:    true,
			description: "application of vaccines from 5 years",
		},
		hasApplicationTestCase{
			dob:         time.Date(2016, 11, 1, 0, 0, 0, 0, time.UTC),
			reqDate:     time.Date(2017, 2, 1, 0, 0, 0, 0, time.UTC),
			response:    false,
			description: "application of vaccines from 3 months",
		},
	}

	for _, HVACase := range cases {
		result := HasVaccineApplication(HVACase.dob, HVACase.reqDate)
		assert.Equal(t, HVACase.response, result, HVACase.description)
	}
}

func TestGetAllDatesToApplyVaccines(t *testing.T) {
	type getDatesTestCase struct {
		dob         time.Time
		response    Vaccines
		description string
	}

	cases := []getDatesTestCase{
		getDatesTestCase{
			dob: time.Date(2016, 11, 1, 0, 0, 0, 0, time.UTC),
			response: Vaccines{
				Date:       time.Date(2016, 11, 1, 0, 0, 0, 0, time.UTC),
				References: []vaccineReference{{"Hepatitis B", "0"}, {"Tuberculosis B.C.G", "unica"}},
			},
			description: "application of vaccines from 7 months",
		},
	}

	for _, getDateCase := range cases {
		result := GetAllDatesToApplyVaccines(getDateCase.dob)
		assert.Equal(t, getDateCase.response, result[0], getDateCase.description)
	}
}

func TestGetVaccinesReferenceByNumberOfMonth(t *testing.T) {
	type getGetVaccinesReferenceTestCase struct {
		dob         time.Time
		months      int
		response    Vaccines
		description string
	}

	cases := []getGetVaccinesReferenceTestCase{
		getGetVaccinesReferenceTestCase{
			dob:    time.Date(2016, 11, 1, 0, 0, 0, 0, time.UTC),
			months: 7,
			response: Vaccines{
				Date:       time.Date(2016, 6, 1, 0, 0, 0, 0, time.UTC),
				References: []vaccineReference{{"Influenza", "2"}},
			},
			description: "application of vaccines from 7 months",
		},
	}

	for _, getVaccinesRefCase := range cases {
		result := GetVaccinesReferenceByNumberOfMonth(getVaccinesRefCase.dob, getVaccinesRefCase.months)
		assert.Equal(t, getVaccinesRefCase.response.References, result.References, getVaccinesRefCase.description)
	}
}
