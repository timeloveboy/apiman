package htmlpart

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

const (
	Regex_Loadtemplate = `<!--{{LoadTemplate(.{0,20})}}-->`
)

func Render(url, html string) string {
	reg := regexp.MustCompile(Regex_Loadtemplate)
	var mc []string
	mc = reg.FindAllString(html, -1)
	var result string = html
	for i := 0; i < len(mc); i++ {
		var tag string = mc[i]
		var str = tag[19 : len(tag)-6]
		var pagename = strings.TrimSpace(str)
		nodepath := ""

		if pagename[0] == '/' {
			root, _ := os.Getwd()
			nodepath = root + "/"
			nodeurl := nodepath + pagename
			value, err := ioutil.ReadFile(nodeurl)
			if err != nil {
				continue
			}
			v := Render(nodepath+pagename, string(value))
			result = strings.Replace(result, tag, v, -1)
		} else {
			root, _ := os.Getwd()
			inx := strings.LastIndex(url, "/")
			nodepath = root + url[:inx] + "/"
			fmt.Println(nodepath)
			nodeurl := nodepath + pagename
			value, err := ioutil.ReadFile(nodeurl)
			if err != nil {
				continue
			}
			v := Render(nodepath+pagename, string(value))
			result = strings.Replace(result, tag, v, -1)
		}
	}
	return result
}
