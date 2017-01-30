package helpers

import (
	"github.com/steventux/obd2-data-api/models"
	"net/http"
)

func BuildObd2Data(r *http.Request) *models.Obd2Data {
	return &models.Obd2Data{
		Session:                   r.FormValue("session"),
		Timestamp:                 r.FormValue("time"),
		MassAirflowRate:           r.FormValue("k10"),
		ThrottlePosition:          r.FormValue("k11"),
		AirStatus:                 r.FormValue("k12"),
		FuelTrimBank1Sensor1:      r.FormValue("k14"),
		FuelTrimBank1Sensor2:      r.FormValue("k15"),
		FuelTrimBank1Sensor3:      r.FormValue("k16"),
		FuelTrimBank1Sensor4:      r.FormValue("k17"),
		FuelTrimBank2Sensor1:      r.FormValue("k18"),
		FuelTrimBank2Sensor2:      r.FormValue("k19"),
		FuelTrimBank2Sensor3:      r.FormValue("k1a"),
		FuelTrimBank2Sensor4:      r.FormValue("k1b"),
		RunTimeSinceEngineStart:   r.FormValue("k1f"),
		DistanceMILOn:             r.FormValue("k21"),
		RelativeFuelRailPressure:  r.FormValue("k22"),
		FuelRailPressure:          r.FormValue("k23"),
		O2Sensor1Equivalence:      r.FormValue("k24"),
		O2Sensor2Equivalence:      r.FormValue("k25"),
		O2Sensor3Equivalence:      r.FormValue("k26"),
		O2Sensor4Equivalence:      r.FormValue("k27"),
		O2Sensor5Equivalence:      r.FormValue("k28"),
		O2Sensor6Equivalence:      r.FormValue("k29"),
		O2Sensor7Equivalence:      r.FormValue("k2a"),
		O2Sensor8Equivalence:      r.FormValue("k2b"),
		EGRCommanded:              r.FormValue("k2c"),
		EGRError:                  r.FormValue("k2d"),
		FuelLevelECU:              r.FormValue("k2f"),
		FuelStatus:                r.FormValue("k3"),
		DistanceSinceCodesCleared: r.FormValue("k31"),
		EvapVapourPressure:        r.FormValue("k32"),
		BarometricPressure:        r.FormValue("k33"),
		O2Sensor1EquivalenceAlt:   r.FormValue("k34"),
		CatTempBank1Sensor1:       r.FormValue("k3c"),
		CatTempBank2Sensor1:       r.FormValue("k3d"),
		CatTempBank1Sensor2:       r.FormValue("k3e"),
		CatTempBank2Sensor2:       r.FormValue("k3f"),
		EngineLoad:                r.FormValue("k4"),
		Voltage:                   r.FormValue("k42"),
		AbsEngineLoad:             r.FormValue("k43"),
		CommandedEquivalenceRatio: r.FormValue("k44"),
		RelativeThrottlePosition:  r.FormValue("k45"),
		AmbientAirTemp:            r.FormValue("k46"),
		AbsThrottlePositionB:      r.FormValue("k47"),
		AccPedalPositionD:         r.FormValue("k49"),
		AccPedalPositionE:         r.FormValue("k4a"),
		AccPedalPositionF:         r.FormValue("k4b"),
		EngineCoolantTemp:         r.FormValue("k5"),
		EthanolFuelRatio:          r.FormValue("k52"),
		RelAccPedalPosition:       r.FormValue("k5a"),
		EngineOilTemp:             r.FormValue("k5c"),
		FuelTrimBank1ShortTerm:    r.FormValue("k6"),
		FuelTrimBank1LongTerm:     r.FormValue("k7"),
		ExhaustGasTemp1:           r.FormValue("k78"),
		ExhaustGasTemp2:           r.FormValue("k79"),
		FuelTrimBank2ShortTerm:    r.FormValue("k8"),
		FuelTrimBank2LongTerm:     r.FormValue("k9"),
		FuelPressure:              r.FormValue("ka"),
		IntakeManifoldPressure:    r.FormValue("kb"),
		TransmissionTempMethod2:   r.FormValue("kb4"),
		EngineRPM:                 r.FormValue("kc"),
		OBDSpeed:                  r.FormValue("kd"),
		TimingAdvance:             r.FormValue("ke"),
		IntakeAirTemp:             r.FormValue("kf"),
		TransmissionTempMethod1:   r.FormValue("kfe1805"),
		GPSSpeed:                  r.FormValue("kff1001"),
		GPSLong:                   r.FormValue("kff1005"),
		GPSLat:                    r.FormValue("kff1006"),
		GPSAltitude:               r.FormValue("kff1010"),
		MPGInstant:                r.FormValue("kff1201"),
		TurboBoost:                r.FormValue("kff1202"),
		KPLInstant:                r.FormValue("kff1203"),
		TripDistance:              r.FormValue("kff1204"),
		TripAverageMPG:            r.FormValue("kff1205"),
		TripAverageKPL:            r.FormValue("kff1206"),
		LitresPer100KMInstant:     r.FormValue("kff1207"),
		TripAverageLitresPer100KM: r.FormValue("kff1208"),
		TripDistanceProfile:       r.FormValue("kff120c"),
		O2VoltsBank1Sensor1:       r.FormValue("kff1214"),
		O2VoltsBank1Sensor2:       r.FormValue("kff1215"),
		O2VoltsBank1Sensor3:       r.FormValue("kff1216"),
		O2VoltsBank1Sensor4:       r.FormValue("kff1217"),
		O2VoltsBank2Sensor1:       r.FormValue("kff1218"),
		O2VoltsBank2Sensor2:       r.FormValue("kff1219"),
		O2VoltsBank2Sensor3:       r.FormValue("kff121a"),
		O2VoltsBank2Sensor4:       r.FormValue("kff121b"),
		AccelSensorX:              r.FormValue("kff1220"),
		AccelSensorY:              r.FormValue("kff1221"),
		AccelSensorZ:              r.FormValue("kff1222"),
		AccelSensorTotal:          r.FormValue("kff1223"),
		Torque:                    r.FormValue("kff1225"),
		Horsepower:                r.FormValue("kff1226"),
		Time0to60MPH:              r.FormValue("kff122d"),
		Time0to100KPH:             r.FormValue("kff122e"),
		QuarterMileTime:           r.FormValue("kff122f"),
		EighthMileTime:            r.FormValue("kff1230"),
		GPSvOBDSpeedDiff:          r.FormValue("kff1237"),
		VoltageOBDAdapter:         r.FormValue("kff1238"),
		GPSAccuracy:               r.FormValue("kff1239"),
		GPSSatellites:             r.FormValue("kff123a"),
		GPSBearing:                r.FormValue("kff123b"),
		O2Sensor1WideRangeVoltage: r.FormValue("kff1240"),
		O2Sensor2WideRangeVoltage: r.FormValue("kff1241"),
		O2Sensor3WideRangeVoltage: r.FormValue("kff1242"),
		O2Sensor4WideRangeVoltage: r.FormValue("kff1243"),
		O2Sensor5WideRangeVoltage: r.FormValue("kff1244"),
		O2Sensor6WideRangeVoltage: r.FormValue("kff1245"),
		O2Sensor7WideRangeVoltage: r.FormValue("kff1246"),
		O2Sensor8WideRangeVoltage: r.FormValue("kff1247"),
		AirFuelRatio:              r.FormValue("kff1249"),
		TiltX:                     r.FormValue("kff124a"),
		TiltY:                     r.FormValue("kff124b"),
		TiltZ:                     r.FormValue("kff124c"),
		AirFuelRatioCommanded:         r.FormValue("kff124d"),
		Time0to200KPH:                 r.FormValue("kff124f"),
		CO2inGperKMInstant:            r.FormValue("kff1257"),
		CO2inGperKMAverage:            r.FormValue("kff1258"),
		FuelFlowRatePerMinute:         r.FormValue("kff125a"),
		TripFuelCost:                  r.FormValue("kff125c"),
		FuelFlowRatePerHour:           r.FormValue("kff125d"),
		Time60to120MPH:                r.FormValue("kff125e"),
		Time60to80MPH:                 r.FormValue("kff125f"),
		Time40to60MPH:                 r.FormValue("kff1260"),
		Time80to100MPH:                r.FormValue("kff1261"),
		TripAverageSpeedMoving:        r.FormValue("kff1263"),
		Time100to0KPH:                 r.FormValue("kff1264"),
		Time60to0MPH:                  r.FormValue("kff1265"),
		TimeTripSinceStart:            r.FormValue("kff1266"),
		TimeTripStationary:            r.FormValue("kff1267"),
		TimeTripMoving:                r.FormValue("kff1268"),
		VolumetricEfficiency:          r.FormValue("kff1269"),
		DistanceToEmpty:               r.FormValue("kff126a"),
		FuelRemaining:                 r.FormValue("kff126b"),
		CostPerMileKMInstant:          r.FormValue("kff126d"),
		CostPerMileKMCalculated:       r.FormValue("kff126e"),
		BarometerOnDevice:             r.FormValue("kff1270"),
		TripFuelUsed:                  r.FormValue("kff1271"),
		TripAverageSpeedTotal:         r.FormValue("kff1272"),
		EngineKw:                      r.FormValue("kff1273"),
		Time80to120KPH:                r.FormValue("kff1275"),
		Time60to130MPH:                r.FormValue("kff1276"),
		Time0to30MPH:                  r.FormValue("kff1277"),
		MPGLongTermAverage:            r.FormValue("kff5201"),
		KPLLongTermAverage:            r.FormValue("kff5202"),
		LitresPer100KMLongTermAverage: r.FormValue("kff5203"),
	}
}
