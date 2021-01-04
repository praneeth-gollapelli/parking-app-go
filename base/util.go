package base

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

//PareBody
func PareBody(r *http.Request) map[string]interface{} {
	body := map[string]interface{}{}
	r.ParseForm()
	for k := range r.Form {
		body[k] = r.FormValue(k)
	}
	return body
}

//StrToInt
func StrToInt(i interface{}) int {
	if i == nil || i == "" {
		return 0
	}
	intValue, _ := strconv.Atoi(i.(string))
	return intValue
}

//StructToMap
func StructToMap(input interface{}) map[string]interface{} {
	inputMap := map[string]interface{}{}
	by, _ := json.Marshal(input)
	json.Unmarshal(by, &inputMap)
	return inputMap
}

//bind template And data
func BindTemplate(htmlTemplate string, data interface{}) string {
	htmlbytes, _ := ioutil.ReadFile(htmlTemplate)
	t, _ := template.New("mail").Parse(string(htmlbytes))
	buf := new(bytes.Buffer)
	err := t.Execute(buf, data)
	if err != nil {
		log.Println(err)
		return ""
	}
	return buf.String()
}
