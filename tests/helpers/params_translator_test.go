// handlers_test.go
package helpers

import (
	"github.com/steventux/obd2-data-api/helpers"
	"net/http"
	"testing"
)

func TestBuildObd2Data(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)

	query := req.URL.Query()
	query.Set("k42", "14.4")
	query.Add("kc", "2550")
	query.Add("kff1225", "199")
	query.Add("kff5201", "36")

	req.URL.RawQuery = query.Encode()

	obd2Data := helpers.BuildObd2Data(req)

	if obd2Data.Voltage != "14.4" {
		t.Errorf("Voltage not correct: got %v want %v", obd2Data.Voltage, "14.4")
	}

	if obd2Data.EngineRPM != "2550" {
		t.Errorf("EngineRPM not correct: got %v want %v", obd2Data.EngineRPM, "2550")
	}

	if obd2Data.Torque != "199" {
		t.Errorf("Torque not correct: got %v want %v", obd2Data.Torque, "199")
	}

	if obd2Data.MPGLongTermAverage != "36" {
		t.Errorf("MPGLongTermAverage not correct: got %v want %v", obd2Data.MPGLongTermAverage, "36")
	}
}
