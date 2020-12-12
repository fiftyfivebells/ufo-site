package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"runtime/debug"
	"time"

	"github.com/joho/godotenv"
)

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

func (app *application) addDefaultData(td *templateData, r *http.Request) *templateData {
	if td == nil {
		td = &templateData{}
	}
	td.CurrentYear = time.Now().Year()
	td.Flash = app.session.PopString(r, "flash")

	return td
}

func (app *application) renderTemplate(w http.ResponseWriter, r *http.Request, name string, td *templateData) {
	ts, ok := app.templateCache[name]

	if !ok {
		app.serverError(w, fmt.Errorf("The template %s does not exist", name))
	}

	buf := new(bytes.Buffer)

	// Write template to buffer to see if it works
	err := ts.Execute(buf, app.addDefaultData(td, r))
	if err != nil {
		app.serverError(w, err)
	}

	// If no error, write the buffer to the response writer
	buf.WriteTo(w)
}

func (app *application) convertTimeToSeason(time time.Time) string {
	m := time.Month()
	var season string

	if m >= 3 && m < 6 {
		season = "spring"
	} else if m >= 6 && m < 9 {
		season = "summer"
	} else if m >= 9 && m < 12 {
		season = "fall"
	} else if m == 1 || m == 2 || m == 12 {
		season = "winter"
	}

	return season
}

func (app *application) getLatAndLong(city, state string) (float64, float64) {

	type geo struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}

	type geoList struct {
		Data []geo `json:"data"`
	}

	err := godotenv.Load(".env")
	if err != nil {
		return 0, 0
	}

	key := os.Getenv("POSITIONSTACK_API")
	query := fmt.Sprintf("http://api.positionstack.com/v1/forward?access_key=%s&query=%s&region=%s", key, city, state)

	resp, err := http.Get(query)
	if err != nil {
		app.errorLog.Println(err)
		return -1, -1
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		app.errorLog.Println(err)
		return -1, -1
	}

	var gl geoList

	err = json.Unmarshal(bodyBytes, &gl)
	if err != nil {
		app.errorLog.Println(err)
		return 0, 0
	}

	return gl.Data[0].Latitude, gl.Data[0].Longitude
}

func (app *application) getStateAbbrev(state string) string {
	abbrevs := make(map[string]string)
	abbrevs["alabama"] = "al"
	abbrevs["alaska"] = "ak"
	abbrevs["arizona"] = "az"
	abbrevs["arkansas"] = "ar"
	abbrevs["california"] = "ca"
	abbrevs["colorado"] = "co"
	abbrevs["connecticut"] = "ct"
	abbrevs["delaware"] = "de"
	abbrevs["dc"] = "dc"
	abbrevs["florida"] = "fl"
	abbrevs["georgia"] = "ga"
	abbrevs["hawaii"] = "hi"
	abbrevs["idaho"] = "id"
	abbrevs["illinois"] = "il"
	abbrevs["indiana"] = "in"
	abbrevs["iowa"] = "ia"
	abbrevs["kansas"] = "ks"
	abbrevs["kentucky"] = "ky"
	abbrevs["louisiana"] = "la"
	abbrevs["maine"] = "me"
	abbrevs["maryland"] = "md"
	abbrevs["massachusetts"] = "ma"
	abbrevs["michigan"] = "mi"
	abbrevs["minnesota"] = "mn"
	abbrevs["mississippi"] = "ms"
	abbrevs["missouri"] = "mo"
	abbrevs["montana"] = "mt"
	abbrevs["nebraska"] = "ne"
	abbrevs["nevada"] = "nv"
	abbrevs["new hampshire"] = "nh"
	abbrevs["new jersey"] = "nj"
	abbrevs["new mexico"] = "nm"
	abbrevs["new york"] = "ny"
	abbrevs["north carolina"] = "nc"
	abbrevs["north dakota"] = "nd"
	abbrevs["ohio"] = "oh"
	abbrevs["oklahoma"] = "ok"
	abbrevs["oregon"] = "or"
	abbrevs["pennsylvania"] = "pa"
	abbrevs["rhode island"] = "ri"
	abbrevs["south carolina"] = "sc"
	abbrevs["south dakota"] = "sd"
	abbrevs["tennessee"] = "tn"
	abbrevs["texas"] = "tx"
	abbrevs["utah"] = "ut"
	abbrevs["vermont"] = "vt"
	abbrevs["virginia"] = "va"
	abbrevs["washington"] = "wa"
	abbrevs["west virginia"] = "wv"
	abbrevs["wisconsin"] = "wi"
	abbrevs["wyoming"] = "wy"

	return abbrevs[state]
}
