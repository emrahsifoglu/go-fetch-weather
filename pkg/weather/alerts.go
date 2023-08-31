package weather

import (
	"log"
	"io"
	"net/http"
	"github.com/paulmach/orb/geojson"
)

const (
	alerts = "/alerts/active"
)

func (c *Client) GetAlerts(area string) (fc *geojson.FeatureCollection, err error) {
	params := make(map[string]string)
	params["area"] = area

	response, err := c.do(http.MethodGet, alerts, params)

	if err != nil {
		log.Panicf("Error occurred in getting response. Err: %s", err)
		return
	}

	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)

	if err != nil {
		log.Panicf("Error occurred in reading data. Err: %s", err)
		return
	}

	if fc, err = geojson.UnmarshalFeatureCollection(data); err != nil {
		return fc, err
	}
 
	return
}
