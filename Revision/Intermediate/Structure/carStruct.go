// Create a Car struct with fields like Brand, Model, and Year. Write a method to print car details.

package structure

import "strconv"

type CarStruct struct {
	Brand string `json:"brand"`
	Model string `json:"model"`
	Year  int    `json:"year"`
}

func NewCarStruct(brand, model string, year int) CarStruct {
	return CarStruct{
		Brand: brand,
		Model: model,
		Year:  year,
	}
}

func (c CarStruct) PrintCarDetails() string {
	return "Car Details: " + c.Brand + " " + c.Model + " year " + strconv.Itoa(c.Year)
}
