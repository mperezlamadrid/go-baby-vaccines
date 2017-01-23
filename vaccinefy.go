package vaccinefy

import (
	"sort"
	"time"
)

type vaccineReference struct {
	Name string
	Dose string
}

// Vaccines is a struct that contain a date and a list of vaccines to apply.
type Vaccines struct {
	Date       time.Time
	References []vaccineReference
}

type byDate []Vaccines

var (
	shortForm = "2006-01-02"

	vaccinesPerMonth = map[int][]vaccineReference{
		0:  []vaccineReference{{"Hepatitis B", "0"}, {"Tuberculosis B.C.G", "unica"}},
		2:  []vaccineReference{{"Polio (Oral - IM)", "1"}, {"Hepatitis B", "1"}, {"Rotavirus", "1"}, {"Neumococo", "1"}},
		4:  []vaccineReference{{"Polio (Oral - IM)", "2"}, {"Hepatitis B", "2"}, {"Rotavirus", "2"}, {"Neumococo", "2"}},
		6:  []vaccineReference{{"Polio (Oral - IM)", "3"}, {"Hepatitis B", "3"}, {"Influenza", "1"}},
		7:  []vaccineReference{{"Influenza", "2"}},
		12: []vaccineReference{{"Sarampión Rubéola Paperas (SRP)", "1"}, {"Neumococo", "refuerzo"}, {"Influenza", "anual"}, {"Hepatitis B", "unica"}},
		18: []vaccineReference{{"Difteria - Tosferina - Tétano (DTP", "1 refuerzo"}, {"Polio (Oral - IM)", "1 refuerzo"}},
		60: []vaccineReference{{"Polio (Oral - IM)", "2 refuerzo"}, {"Difteria - Tosferina - Tétano (DTP)", "2 refuerzo"}, {"Sarampión Rubéola Paperas (SRP)", "refuerzo"}},
	}
)

// HasVaccineApplication receives date of birth and a request date and
// returns true or false about if that day there are vaccines to apply.
func HasVaccineApplication(dob, reqDate time.Time) bool {
	monthsOld := int(reqDate.Sub(dob).Hours() / 720)

	if len(vaccinesPerMonth[monthsOld]) > 0 {
		return true
	}
	return false
}

// GetAllDatesToApplyVaccines receives a date of birth and returns a
// list of dates and references about its vaccines.
func GetAllDatesToApplyVaccines(dob time.Time) []Vaccines {
	var vaccinesDate []Vaccines

	for months, vaccines := range vaccinesPerMonth {
		new := Vaccines{
			Date:       dob.AddDate(0, months, 0),
			References: vaccines,
		}
		vaccinesDate = append(vaccinesDate, new)
	}

	sort.Sort(byDate(vaccinesDate))
	return vaccinesDate
}

func GetVaccinesReferenceByNumberOfMonth(dob time.Time, months int) Vaccines {
	var vaccineDate Vaccines

	if len(vaccinesPerMonth[months]) > 0 {
		vaccineDate = Vaccines{
			Date:       dob.AddDate(0, months, 0),
			References: vaccinesPerMonth[months],
		}
	}
	return vaccineDate
}

func (d byDate) Len() int           { return len(d) }
func (d byDate) Less(i, j int) bool { return d[i].Date.Before(d[j].Date) }
func (d byDate) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }
