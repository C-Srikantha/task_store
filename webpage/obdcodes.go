package webpage

import (
	"encoding/csv"
	"fmt"
	"os"
	"regexp"

	"github.com/anaskhan96/soup"
)

/*var m1 = map[string]string{
	"What does that mean?":             "",
	"Severity & Symptoms":              "",
	"Causes":                           "",
	"Diagnostic and Repair Procedures": "",
	"Related DTC Discussions":          "",
}*/
var previoush2 string
var str string = ""
var urls = map[string]string{
	"P00A9": "https://www.obd-codes.com/p00a9",
	"P00A5": "https://www.obd-codes.com/p00a5",
	"P00A6": "https://www.obd-codes.com/p00a6",
	"P00A7": "https://www.obd-codes.com/p00a7",
	"P00A8": "https://www.obd-codes.com/p00a8",
	"P0000": "https://www.obd-codes.com/p0000",
	"P0001": "https://www.obd-codes.com/p0001",
	"P0002": "https://www.obd-codes.com/p0002",
}
var flag int = 1
var det1 []map[string]string

func Obdcodes() {
	for key, val := range urls {
		var m1 = map[string]string{
			"code123":   "",
			"mean":      "",
			"Causes":    "",
			"Symptoms":  "",
			"Repair":    "",
			"Solutions": "",
		}
		resp, err := soup.Get(val)
		if err != nil {
			fmt.Println(err)
		}
		links := soup.HTMLParse(resp)
		m1["code123"] = key
		//
		link := links.Find("div", "class", "main")
		res1 := link.Children() //gets all childrens from class main
		for _, i := range res1 {
			if i.NodeValue == "h2" { //checks tag is h2 or not
				for k := range m1 {
					re := regexp.MustCompile(k)
					if re.FindString(i.FullText()) == k { //h2 string matches with key or not
						previoush2 = k //assinging h2 to variable
						flag = 0
						break
					} else {
						flag = 1
					}
				}

				str = ""
			}
			if flag != 1 {
				if i.NodeValue == "p" {
					check := i.Find("ins", "class", "adsbygoogle") //rejects reading add parts inside para
					if check.Error != nil {
						str += i.FullText() + "\n"
					}

				}
				if i.NodeValue == "ul" { //parsing inside ul tag
					res := i.FindAll("li")
					for _, j := range res {
						str += j.FullText() + "\n"
						res1 := j.Find("p")
						if res1.Error == nil {
							str += res1.FullText() + "\n"
						}

					}
				}
				m1[previoush2] = str //assiging val to h2
			}

		}
		det1 = append(det1, m1)
		str = ""
	}
	writefile1()
}
func writefile1() {
	var repair string
	file, err := os.Create("dct1.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	csvfile := csv.NewWriter(file)
	defer csvfile.Flush()
	header := []string{"code", "General_meaning", "Causes", "Symptoms", "Mechanic_diagnosis", "Severity_level", "Suggested_repairs"}
	csvfile.Write(header)
	for _, m := range det1 {
		if m["Repair"] != "" {
			repair = m["Repair"]
		} else {
			repair = m["Solutions"]
		}
		row := []string{m["code123"], m["mean"], m["Causes"], m["Symptoms"], repair, m["Symptoms"], repair}
		csvfile.Write(row)
	}
	fmt.Println("success")
}
