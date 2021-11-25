package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/anaskhan96/soup"
)

var url = map[string]string{"P0110": "https://www.autoblog.com/2016/03/11/p0110-obd-ii-trouble-code-intake-air-temperature-sensor-circui/",
	"P0010": "https://www.autoblog.com/2016/03/29/p0010-obd-ii-trouble-code-a-camshaft-position-actuator-circui/"}

var det []map[string]string
var str1 string = ""

func main() {
	for k, v := range url {
		var m1 = map[string]string{"code": "",
			"What the " + k + " code means":                    "",
			"What causes the " + k + " code?":                  "",
			"What are the symptoms of the " + k + " code?":     "",
			"How does a mechanic diagnose the " + k + " code?": "",
			"How serious is the " + k + " code?":               "",
			"What repairs can fix the " + k + " code?":         "",
		}
		var m2 = map[string]string{"What the " + k + " code means": "",
			"What causes the " + k + " code?":                                     "",
			"What are the symptoms of the " + k + " code?":                        "",
			"How does a mechanic diagnose the " + k + " code?":                    "",
			"How serious is the " + k + " code?":                                  "",
			"What repairs can fix the " + k + " code?":                            "",
			"Common mistakes when diagnosing the " + k + " code.":                 "",
			"Additional comments for consideration regarding the " + k + " code.": "",
		}

		resp, err := soup.Get(v) //passing url and returns http response
		if err != nil {
			fmt.Println(err)
		}
		//parsing through html
		result := soup.HTMLParse(resp)
		m1["code"] = k
		links := result.Find("div", "class", "post-body")
		res1 := links.Children() //to get all the childrens from class "post-body"
		for index, i := range res1 {
			if _, ok := m1[i.Text()]; ok { //checking wheather its an header-"h3" tag
				res := res1[index+1:]
				//listing paragraph between h3 tags
				for _, i1 := range res {

					if _, ok1 := m2[i1.Text()]; ok1 {
						str1 = ""
						break

					}
					//rejecting empty strings
					if i1.Text() != "" {
						str1 = str1 + i1.Text()
						/*r := i1.Find("a")
						if r.Error == nil {
							str1 = str1 + r.Text()
							//fmt.Println(str1)
						}*/
					}
					//parsing inside ul tags
					if i1.NodeValue == "ul" {
						res2 := i1.FindAll("li")
						for _, j := range res2 {
							//str1 = str1 + j.Find("p").Text()
							str1 = str1 + j.Text()
							res3 := j.Find("p")
							if res3.Error == nil {
								str1 = str1 + res3.Text()
							}
						}
					}

					m1[i.Text()] = str1 //storing value for the keys

				}

			}
		}
		//fmt.Println(m1)
		det = append(det, m1) //storing maps into array
		str1 = ""             //clears all contents from variable
	}

	writefile()
}
func writefile() {
	file, err := os.Create("dct.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	csvfile := csv.NewWriter(file)
	defer csvfile.Flush()
	header := []string{"code", "General_meaning", "Causes", "Symptoms", "Mechanic_diagnosis", "Severity_level", "Suggested_repairs"}
	csvfile.Write(header) //writing header of file
	//storing all datas into file
	for _, d := range det {
		code := d["code"]
		fmt.Println(code)
		row := []string{code, d["What the "+code+" code means"], d["What causes the "+code+" code?"], d["What are the symptoms of the "+code+" code?"],
			d["How does a mechanic diagnose the "+code+" code?"], d["How serious is the "+code+" code?"], d["What repairs can fix the "+code+" code?"]}
		csvfile.Write(row)
	}
}
