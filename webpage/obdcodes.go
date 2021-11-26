package webpage

import (
	"fmt"

	"github.com/anaskhan96/soup"
)

var m1 = map[string]string{
	"What does that mean?":             "",
	"Severity & Symptoms":              "",
	"Causes":                           "",
	"Diagnostic and Repair Procedures": "",
	"Related DTC Discussions":          "",
}
var str string

func Obdcodes() {
	resp, err := soup.Get("https://www.obd-codes.com/p00a9")
	if err != nil {
		fmt.Println(err)
	}
	links := soup.HTMLParse(resp)
	link := links.Find("div", "class", "main")
	res1 := link.Children()
	for index, i := range res1 {
		if _, ok := m1[i.Text()]; ok {
			for _, j := range res1[index+1:] {
				if _, ok := m1[j.Text()]; ok {
					str = " "
					break
				}
				if j.Text() != "" {
					str += j.Text()
				}

			}
		}
		m1[i.Text()] = str
		str = ""
	}
	fmt.Println(m1)

}
