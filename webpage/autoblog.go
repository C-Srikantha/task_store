package webpage

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/anaskhan96/soup"
)

var url = map[string]string{"P0110": "https://www.autoblog.com/2016/03/11/p0110-obd-ii-trouble-code-intake-air-temperature-sensor-circui/",
	"P0010": "https://www.autoblog.com/2016/03/29/p0010-obd-ii-trouble-code-a-camshaft-position-actuator-circui/",
	"P0131": "https://www.autoblog.com/2016/03/23/p0131-obd-ii-trouble-code-o2-sensor-circuit-low-voltage-bank-1/",
	"P0171": "https://www.autoblog.com/2016/04/06/p0171-obd-ii-trouble-code-system-too-lean-bank-1/",
	"P0977": "https://www.autoblog.com/2016/09/14/p0977-obd-ii-trouble-code-shift-solenoid-b-control-circuit-high/", //
	"P0011": "https://www.autoblog.com/2016/03/25/p0011-obd-ii-trouble-code-camshaft-position-a-timing-over-adv/",
	"P0020": "https://www.autoblog.com/2016/03/18/p0020-obd-ii-trouble-code-camshaft-position-a-actuator-circui/",
	"P0013": "https://www.autoblog.com/2016/03/25/p0013-obd-ii-trouble-code-b-camshaft-position-open-or-short/",
	"P0012": "https://www.autoblog.com/2016/03/30/p0012-obd-ii-trouble-code-camshaft-position-a-timing-over-ret/",
	"P0016": "https://www.autoblog.com/2016/03/18/p0016-obd-ii-trouble-code-camshaft-position-a-camshaft-positi/",
	"P0084": "https://www.autoblog.com/2016/03/28/p0084-obd-ii-trouble-code-exhaust-valve-control-solenoid-circui/",
	"P0019": "https://www.autoblog.com/2016/03/28/p0019-obd-ii-trouble-codes-crankshaft-position-camshaft-posit/",
	"P0076": "https://www.autoblog.com/2016/03/28/p0076-obd-ii-trouble-code-intake-valve-control-solenoid-circuit/",
	"P0090": "https://www.autoblog.com/2016/03/28/p0090-obd-ii-trouble-code-fuel-pressure-regulator-1-control-cir/",
	"P0075": "https://www.autoblog.com/2016/03/24/p0075-obd-ii-trouble-code-intake-valve-control-solenoid-circuit/",
	"P0028": "https://www.autoblog.com/2016/03/18/p0028-obd-ii-trouble-code-intake-valve-control-solenoid-circuit/",
	"P0119": "https://www.autoblog.com/2016/03/15/p0119-obd-ii-trouble-code-ect-sensor-circuit-intermittent-malfu/",
}

var det []map[string]string
var str1 string = ""

func Autoblog() {
	for k, v := range url {
		var m1 = map[string]string{"code": "",
			"What the " + k + " code means":                                       "",
			"What does the " + k + " code mean?":                                  "",
			"What causes the " + k + " code?":                                     "",
			"What are the symptoms of the " + k + " code?":                        "",
			"How does a mechanic diagnose the " + k + " code?":                    "",
			"How serious is the " + k + " code?":                                  "",
			"What repairs can fix the " + k + " code?":                            "",
			"Common mistakes when diagnosing the " + k + " code.":                 "",
			"Common mistakes when diagnosing the " + k + " code":                  "",
			"Additional comments for consideration regarding the " + k + " code.": "",
			"Additional comments for consideration regarding the " + k + " code":  "",
			"What are the causes of the " + k + " code?":                          "",
			"Additional comments regarding the " + k + " code":                    "",
			"What are some of the symptoms of " + k + " code?":                    "",
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

					if _, ok1 := m1[i1.Text()]; ok1 {
						str1 = ""
						break

					}
					//rejecting empty strings
					if i1.Text() != "" {
						str1 += i1.FullText()
					}
					//parsing inside ul tags
					if i1.NodeValue == "ul" {
						res2 := i1.FindAll("li")
						for _, j := range res2 {

							str1 = str1 + " " + j.FullText()
							res3 := j.Find("p")
							if res3.Error == nil {
								str1 = str1 + " " + res3.FullText()
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
		var causes, meaning, symptoms string
		code := d["code"]
		fmt.Println(code)
		if d["What are the causes of the "+code+" code?"] != "" {
			causes = d["What are the causes of the "+code+" code?"]
		} else {
			causes = d["What causes the "+code+" code?"]
		}
		if d["What the "+code+" code means"] != "" {
			meaning = d["What the "+code+" code means"]
		} else {
			meaning = d["What does the "+code+" code mean?"]
		}
		if d["What are some of the symptoms of "+code+" code?"] != "" {
			symptoms = d["What are some of the symptoms of "+code+" code?"]
		} else {
			symptoms = d["What are the symptoms of the "+code+" code?"]
		}
		row := []string{code, meaning, causes, symptoms,
			d["How does a mechanic diagnose the "+code+" code?"], d["How serious is the "+code+" code?"], d["What repairs can fix the "+code+" code?"]}
		csvfile.Write(row)
	}
}
