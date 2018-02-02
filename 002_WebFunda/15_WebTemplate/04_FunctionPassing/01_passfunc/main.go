package main

import (
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

var fnmap = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	// Look at the Special Initialization Needed
	tpl = template.Must(template.New("").Funcs(fnmap).ParseGlob("*.html"))
	// Actually the part `template.New("")` returns the `*Template` and
	//	again the Function Map needs to be assigned before parsing of the
	//  other templates.
	//  So initially a Blank Template is used to get the FuncMap ready.
	//  Then we use the `ParseGlob` to process the HTML files in the directory.
	// Another Way to do this is:
	// tpl = template.New("").Funcs(fnmap)
	// tpl = template.Must(tpl.ParseGlob("*.html"))
}

// firstThree obtains the first 3 characters in a given string.
// It also removed spaces if any in the input string passed to the function.
func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

type people struct {
	Name  string
	Quote string
}

func main() {
	p1 := people{
		Name:  "Krishna",
		Quote: "Submit on to me",
	}
	p2 := people{
		Name:  "Swami Vivekananda",
		Quote: "Arise Awake , stop not till the Goal is reached",
	}
	p3 := people{
		Name:  "Lala Lajpath Rai",
		Quote: "Swaraj is my Birthright",
	}
	pSlice := []people{
		p1,
		p2,
		p3,
	}
	tpl.ExecuteTemplate(os.Stdout, "tpl.html", pSlice)
}
