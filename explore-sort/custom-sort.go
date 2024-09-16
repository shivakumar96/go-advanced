package customsort

import "strings"

const (
	SortByHumidity   = iota
	SortByTemperatue = iota
	SortByDay        = iota
)

type SortCategory uint

type DayWeather struct {
	Humidity    float32
	Temperature float32
	Day         string
}

type DayWeatherSliceStruct struct {
	DayWeatherSlice []DayWeather
	SortBy          SortCategory
}

func NewDayWeatherSliceStruct(sortBy SortCategory, size int) *DayWeatherSliceStruct {
	d := &DayWeatherSliceStruct{DayWeatherSlice: make([]DayWeather, 0, size), SortBy: sortBy}
	return d
}

func (d *DayWeatherSliceStruct) AddWeather(humidity float32, temperature float32, day string) {
	d.DayWeatherSlice = append(d.DayWeatherSlice, DayWeather{Humidity: humidity, Temperature: temperature, Day: day})
}

func (d *DayWeatherSliceStruct) Len() int {
	return len(d.DayWeatherSlice)
}
func (d *DayWeatherSliceStruct) Swap(i, j int) {
	d.DayWeatherSlice[i], d.DayWeatherSlice[j] = d.DayWeatherSlice[j], d.DayWeatherSlice[i]
}

func (d *DayWeatherSliceStruct) Less(i, j int) bool {
	switch d.SortBy {
	case SortByHumidity:
		return d.DayWeatherSlice[i].Humidity < d.DayWeatherSlice[j].Humidity
	case SortByTemperatue:
		return d.DayWeatherSlice[i].Temperature < d.DayWeatherSlice[j].Temperature
	case SortByDay:
		res := strings.Compare(d.DayWeatherSlice[i].Day, d.DayWeatherSlice[j].Day)
		if res <= 0 {
			return true
		} else {
			return false
		}
	default:
		return d.DayWeatherSlice[i].Temperature < d.DayWeatherSlice[j].Temperature
	}
}
