package vaccinefy

import (
	"encoding/json"
	"io/ioutil"
	"sort"
	"strconv"
	"time"
)

const (
	shortForm = "2006-01-02"
)

type vaccineInfo struct {
	Name string `json:"name"`
	Dose string `json:"dose"`
}

// Vaccines is a struct that contain a date and a list of vaccines to apply.
type Vaccines struct {
	Date       time.Time
	References []vaccineInfo
}

// HasVaccineApplication receives date of birth and a request date and
// returns true or false about if that day there are vaccines to apply.
func HasVaccineApplication(dob, reqDate time.Time) bool {
	vaccinesPerMonth := getVaccinesData()
	monthsOld := strconv.Itoa(int(reqDate.Sub(dob).Hours() / 720))

	if len(vaccinesPerMonth[monthsOld]) > 0 {
		return true
	}
	return false
}

// GetAllDatesToApplyVaccines receives a date of birth and returns a
// list of dates and references about its vaccines.
func GetAllDatesToApplyVaccines(dob time.Time) []Vaccines {
	vaccinesPerMonth := getVaccinesData()
	var vaccinesDate []Vaccines

	for month, vaccines := range vaccinesPerMonth {
		numberOfMonth, _ := strconv.Atoi(month)
		new := Vaccines{
			Date:       dob.AddDate(0, numberOfMonth, 0),
			References: vaccines,
		}
		vaccinesDate = append(vaccinesDate, new)
	}

	sort.Sort(byDate(vaccinesDate))
	return vaccinesDate
}

func getVaccinesData() map[string][]vaccineInfo {
	file, _ := ioutil.ReadFile("data.json")

	var data map[string][]vaccineInfo
	json.Unmarshal(file, &data)

	return data
}

type byDate []Vaccines

func (d byDate) Len() int           { return len(d) }
func (d byDate) Less(i, j int) bool { return d[i].Date.Before(d[j].Date) }
func (d byDate) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }
