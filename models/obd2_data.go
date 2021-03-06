package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Obd2Data struct {
	ID                            bson.ObjectId `bson:"_id,omitempty"`
	Session                       string
	Timestamp                     string
	CreatedAt                     time.Time
	MassAirflowRate               string
	ThrottlePosition              string
	AirStatus                     string
	FuelTrimBank1Sensor1          string
	FuelTrimBank1Sensor2          string
	FuelTrimBank1Sensor3          string
	FuelTrimBank1Sensor4          string
	FuelTrimBank2Sensor1          string
	FuelTrimBank2Sensor2          string
	FuelTrimBank2Sensor3          string
	FuelTrimBank2Sensor4          string
	RunTimeSinceEngineStart       string
	DistanceMILOn                 string
	RelativeFuelRailPressure      string
	FuelRailPressure              string
	O2Sensor1Equivalence          string
	O2Sensor2Equivalence          string
	O2Sensor3Equivalence          string
	O2Sensor4Equivalence          string
	O2Sensor5Equivalence          string
	O2Sensor6Equivalence          string
	O2Sensor7Equivalence          string
	O2Sensor8Equivalence          string
	EGRCommanded                  string
	EGRError                      string
	FuelLevelECU                  string
	FuelStatus                    string
	DistanceSinceCodesCleared     string
	EvapVapourPressure            string
	BarometricPressure            string
	O2Sensor1EquivalenceAlt       string
	CatTempBank1Sensor1           string
	CatTempBank2Sensor1           string
	CatTempBank1Sensor2           string
	CatTempBank2Sensor2           string
	EngineLoad                    string
	Voltage                       string
	AbsEngineLoad                 string
	CommandedEquivalenceRatio     string
	RelativeThrottlePosition      string
	AmbientAirTemp                string
	AbsThrottlePositionB          string
	AccPedalPositionD             string
	AccPedalPositionE             string
	AccPedalPositionF             string
	EngineCoolantTemp             string
	EthanolFuelRatio              string
	RelAccPedalPosition           string
	EngineOilTemp                 string
	FuelTrimBank1ShortTerm        string
	FuelTrimBank1LongTerm         string
	ExhaustGasTemp1               string
	ExhaustGasTemp2               string
	FuelTrimBank2ShortTerm        string
	FuelTrimBank2LongTerm         string
	FuelPressure                  string
	IntakeManifoldPressure        string
	TransmissionTempMethod2       string
	EngineRPM                     string
	OBDSpeed                      string
	TimingAdvance                 string
	IntakeAirTemp                 string
	TransmissionTempMethod1       string
	GPSSpeed                      string
	GPSLong                       string
	GPSLat                        string
	GPSAltitude                   string
	MPGInstant                    string
	TurboBoost                    string
	KPLInstant                    string
	TripDistance                  string
	TripAverageMPG                string
	TripAverageKPL                string
	LitresPer100KMInstant         string
	TripAverageLitresPer100KM     string
	TripDistanceProfile           string
	O2VoltsBank1Sensor1           string
	O2VoltsBank1Sensor2           string
	O2VoltsBank1Sensor3           string
	O2VoltsBank1Sensor4           string
	O2VoltsBank2Sensor1           string
	O2VoltsBank2Sensor2           string
	O2VoltsBank2Sensor3           string
	O2VoltsBank2Sensor4           string
	AccelSensorX                  string
	AccelSensorY                  string
	AccelSensorZ                  string
	AccelSensorTotal              string
	Torque                        string
	Horsepower                    string
	Time0to60MPH                  string
	Time0to100KPH                 string
	QuarterMileTime               string
	EighthMileTime                string
	GPSvOBDSpeedDiff              string
	VoltageOBDAdapter             string
	GPSAccuracy                   string
	GPSSatellites                 string
	GPSBearing                    string
	O2Sensor1WideRangeVoltage     string
	O2Sensor2WideRangeVoltage     string
	O2Sensor3WideRangeVoltage     string
	O2Sensor4WideRangeVoltage     string
	O2Sensor5WideRangeVoltage     string
	O2Sensor6WideRangeVoltage     string
	O2Sensor7WideRangeVoltage     string
	O2Sensor8WideRangeVoltage     string
	AirFuelRatio                  string
	TiltX                         string
	TiltY                         string
	TiltZ                         string
	AirFuelRatioCommanded         string
	Time0to200KPH                 string
	CO2inGperKMInstant            string
	CO2inGperKMAverage            string
	FuelFlowRatePerMinute         string
	TripFuelCost                  string
	FuelFlowRatePerHour           string
	Time60to120MPH                string
	Time60to80MPH                 string
	Time40to60MPH                 string
	Time80to100MPH                string
	TripAverageSpeedMoving        string
	Time100to0KPH                 string
	Time60to0MPH                  string
	TimeTripSinceStart            string
	TimeTripStationary            string
	TimeTripMoving                string
	VolumetricEfficiency          string
	DistanceToEmpty               string
	FuelRemaining                 string
	CostPerMileKMInstant          string
	CostPerMileKMCalculated       string
	BarometerOnDevice             string
	TripFuelUsed                  string
	TripAverageSpeedTotal         string
	EngineKw                      string
	Time80to120KPH                string
	Time60to130MPH                string
	Time0to30MPH                  string
	MPGLongTermAverage            string
	KPLLongTermAverage            string
	LitresPer100KMLongTermAverage string
}
