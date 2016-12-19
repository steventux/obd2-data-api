// handlers_test.go
package helpers

import (
	"bytes"
	"github.com/steventux/obd2-data-api/helpers"
	"net/http"
	"net/url"
	"testing"
)

func TestBuildObd2Data(t *testing.T) {
	data := url.Values{}
	data.Set("k42", "14")
	data.Add("kc", "2550")
	data.Add("kff1225", "199")
	data.Add("kff5201", "36")

	req, err := http.NewRequest("POST", "/", bytes.NewBufferString(data.Encode()))

	if err != nil {
		t.Fatal(err)
	}

	obd2Data := helpers.BuildObd2Data(req)

	if obd2Data.Voltage != 14 {
		t.Errorf("Voltage not correct: got %v want %v", obd2Data.Voltage, 14)
	}

	if obd2Data.EngineRPM != 2550 {
		t.Errorf("EngineRPM not correct: got %v want %v", obd2Data.EngineRPM, 2550)
	}

	if obd2Data.Torque != 199 {
		t.Errorf("Torque not correct: got %v want %v", obd2Data.Torque, 199)
	}

	if obd2Data.MPGLongTermAverage != 36 {
		t.Errorf("MPGLongTermAverage not correct: got %v want %v", obd2Data.MPGLongTermAverage, 36)
	}
}
