package data_test

import (
	"testing"

	"github.com/danesparza/mapimage/data"
)

func TestMapImage_GetMapImageForCoordinates_ReturnsValidData(t *testing.T) {
	//	Arrange
	lat := 33.97561
	long := -83.883747
	zoom := 11

	//	Act
	response, err := data.GetMapImageForCoordinates(lat, long, zoom)

	//	Assert
	if err != nil {
		t.Errorf("Error calling GetMapImageForCoordinates: %v", err)
	}

	t.Logf("Returned data: %+v", response)

}
