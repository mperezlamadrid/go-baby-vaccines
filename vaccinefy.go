package vaccinefy

import (
	"sort"
	"time"
)

type vaccine struct {
	Name string
	Dose string
}

// InfoVaccines is a struct that contain a date and a list of vaccines to apply.
type InfoVaccines struct {
	Date       time.Time
	References []vaccine
}

type byDate []InfoVaccines

var (
	shortForm = "2006-01-02"

	hoursPerMonth = float64(720)

	vaccinesPerMonth = map[int][]vaccine{
		0:  []vaccine{{"Hepatitis B", "0"}, {"Tuberculosis B.C.G", "unica"}},
		2:  []vaccine{{"Polio (Oral - IM)", "1"}, {"Hepatitis B", "1"}, {"Rotavirus", "1"}, {"Neumococo", "1"}},
		4:  []vaccine{{"Polio (Oral - IM)", "2"}, {"Hepatitis B", "2"}, {"Rotavirus", "2"}, {"Neumococo", "2"}},
		6:  []vaccine{{"Polio (Oral - IM)", "3"}, {"Hepatitis B", "3"}, {"Influenza", "1"}},
		7:  []vaccine{{"Influenza", "2"}},
		12: []vaccine{{"Sarampión Rubéola Paperas (SRP)", "1"}, {"Neumococo", "refuerzo"}, {"Influenza", "anual"}, {"Hepatitis B", "unica"}},
		18: []vaccine{{"Difteria - Tosferina - Tétano (DTP", "1 refuerzo"}, {"Polio (Oral - IM)", "1 refuerzo"}},
		60: []vaccine{{"Polio (Oral - IM)", "2 refuerzo"}, {"Difteria - Tosferina - Tétano (DTP)", "2 refuerzo"}, {"Sarampión Rubéola Paperas (SRP)", "refuerzo"}},
	}
)

// HasVaccinesToApply receives date of birth and a request date and
// returns true or false about if that day there are vaccines to apply.
func HasVaccinesToApply(DOB, reqDate time.Time) bool {
	monthsDif := int(reqDate.Sub(DOB).Hours() / hoursPerMonth)

	if len(vaccinesPerMonth[monthsDif]) > 0 {
		return true
	}
	return false
}

// GetDatesToApplyVaccines receives a date of birth and returns a
// list of dates and references about its vaccines.
func GetDatesToApplyVaccines(DOB time.Time) []InfoVaccines {
	var vaccinesDate []InfoVaccines

	for months, vaccines := range vaccinesPerMonth {
		new := InfoVaccines{
			Date:       DOB.AddDate(0, months, 0),
			References: vaccines,
		}
		vaccinesDate = append(vaccinesDate, new)
	}

	sort.Sort(byDate(vaccinesDate))
	return vaccinesDate
}

// GetVaccinesReference receives a date of birth and number of months,
// it returns a vaccines reference if there
func GetVaccinesReference(DOB time.Time, months int) InfoVaccines {
	var vaccineDate InfoVaccines

	if len(vaccinesPerMonth[months]) == 0 {
		return vaccineDate
	}

	vaccineDate = InfoVaccines{
		Date:       DOB.AddDate(0, months, 0),
		References: vaccinesPerMonth[months],
	}

	return vaccineDate
}

func (d byDate) Len() int           { return len(d) }
func (d byDate) Less(i, j int) bool { return d[i].Date.Before(d[j].Date) }
func (d byDate) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }
