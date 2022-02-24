package data

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image/color"
	"image/jpeg"

	sm "github.com/flopp/go-staticmaps"
	"github.com/golang/geo/s2"
	log "github.com/sirupsen/logrus"
)

// GetMapImageForCoordinates gets a map image for the given lat, long and zoom level
// returns the map image as a base64 encoded jpeg
func GetMapImageForCoordinates(lat, long float64, zoom int) (MapImageResponse, error) {

	//	Set a default zoom
	if zoom == 0 {
		zoom = 3
	}

	retval := MapImageResponse{
		Lat:  lat,
		Long: long,
		Zoom: zoom,
	}

	//	Get the map image
	ctx := sm.NewContext()
	ctx.SetSize(600, 360)
	ctx.SetZoom(zoom)
	ctx.SetTileProvider(sm.NewTileProviderWikimedia())
	ctx.AddObject(
		sm.NewMarker(
			s2.LatLngFromDegrees(lat, long),
			color.RGBA{0xff, 0, 0, 0xff},
			16.0,
		),
	)

	img, err := ctx.Render()
	if err != nil {
		log.WithError(err).WithFields(log.Fields{
			"lat":  lat,
			"long": long,
			"zoom": zoom,
		}).Error("error rendering map image")

		return retval, fmt.Errorf("error rendering image: %v", err)
	}

	//	Encode to jpg
	buffer := new(bytes.Buffer)
	jpeg.Encode(buffer, img, nil)

	//	Encode the jpeg to base64
	retval.Image = base64.StdEncoding.EncodeToString(buffer.Bytes())

	//	Return what we have
	return retval, nil
}
