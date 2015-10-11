// This file is auto-generated using the
// program found in 'cmd/fitgen/main.go'
// DO NOT EDIT.
// SDK Version: 16.10
// Generation time: Sun Oct 11 13:19:55 UTC 2015

package fit

import (
	"fmt"
	"reflect"
)

// field represents a fit message field in the profile field lookup table.
type field struct {
	sindex int
	array  uint8
	t      gotype
	num    byte
	btype  fitBaseType
}

func (f field) String() string {
	return fmt.Sprintf(
		"num: %d | btype: %v | sindex: %d | array: %d",
		f.num, f.btype, f.sindex, f.array,
	)
}

// gotype is used in the profile field lookup table to represent the data type
// (or type category) for a field when decoded into a Go message struct.
type gotype uint8

const (
	fit gotype = iota // Standard -> Fit base type/alias

	// Special (non-profile types)
	timeutc   // Time UTC 	-> time.Time
	timelocal // Time Local -> time.Time with Location
	lat       // Latitude 	-> fit.Latitude
	lng       // Longitude 	-> fit.Longitude
)

func (g gotype) String() string {
	if int(g) > len(gotypeString) {
		return fmt.Sprintf("gotype(%d)", g)
	}
	return gotypeString[g]
}

var gotypeString = [...]string{
	"fit",
	"timeutc",
	"timelocal",
	"lat",
	"lng",
}

var knownMsgNums = map[MesgNum]bool{
	MesgNumFileId:                  true,
	MesgNumCapabilities:            true,
	MesgNumDeviceSettings:          true,
	MesgNumUserProfile:             true,
	MesgNumHrmProfile:              true,
	MesgNumSdmProfile:              true,
	MesgNumBikeProfile:             true,
	MesgNumZonesTarget:             true,
	MesgNumHrZone:                  true,
	MesgNumPowerZone:               true,
	MesgNumMetZone:                 true,
	MesgNumSport:                   true,
	MesgNumGoal:                    true,
	MesgNumSession:                 true,
	MesgNumLap:                     true,
	MesgNumRecord:                  true,
	MesgNumEvent:                   true,
	MesgNumDeviceInfo:              true,
	MesgNumWorkout:                 true,
	MesgNumWorkoutStep:             true,
	MesgNumSchedule:                true,
	MesgNumWeightScale:             true,
	MesgNumCourse:                  true,
	MesgNumCoursePoint:             true,
	MesgNumTotals:                  true,
	MesgNumActivity:                true,
	MesgNumSoftware:                true,
	MesgNumFileCapabilities:        true,
	MesgNumMesgCapabilities:        true,
	MesgNumFieldCapabilities:       true,
	MesgNumFileCreator:             true,
	MesgNumBloodPressure:           true,
	MesgNumSpeedZone:               true,
	MesgNumMonitoring:              true,
	MesgNumTrainingFile:            true,
	MesgNumHrv:                     true,
	MesgNumLength:                  true,
	MesgNumMonitoringInfo:          true,
	MesgNumSlaveDevice:             true,
	MesgNumCadenceZone:             true,
	MesgNumSegmentLap:              true,
	MesgNumMemoGlob:                true,
	MesgNumSegmentId:               true,
	MesgNumSegmentLeaderboardEntry: true,
	MesgNumSegmentPoint:            true,
	MesgNumSegmentFile:             true,
	MesgNumCameraEvent:             true,
	MesgNumTimestampCorrelation:    true,
	MesgNumGyroscopeData:           true,
	MesgNumAccelerometerData:       true,
	MesgNumThreeDSensorCalibration: true,
	MesgNumVideoFrame:              true,
	MesgNumObdiiData:               true,
	MesgNumNmeaSentence:            true,
	MesgNumAviationAttitude:        true,
	MesgNumVideo:                   true,
	MesgNumVideoTitle:              true,
	MesgNumVideoDescription:        true,
	MesgNumVideoClip:               true,
}

// Set length to 256, so that lookup for any
// field 255 (localMesgNumInvalid) will return nil.
var _fields = [...][256]*field{
	MesgNumFileId: {
		0: {0, 0, fit, 0, fitEnum},
		1: {1, 0, fit, 1, fitUint16},
		2: {2, 0, fit, 2, fitUint16},
		3: {3, 0, fit, 3, fitUint32z},
		4: {4, 0, timeutc, 4, fitUint32},
		5: {5, 0, fit, 5, fitUint16},
		8: {6, 0, fit, 8, fitString},
	},

	MesgNumFileCreator: {
		0: {0, 0, fit, 0, fitUint16},
		1: {1, 0, fit, 1, fitUint8},
	},

	MesgNumTimestampCorrelation: {},

	MesgNumSoftware: {
		254: {0, 0, fit, 254, fitUint16},
		3:   {1, 0, fit, 3, fitUint16},
		5:   {2, 0, fit, 5, fitString},
	},

	MesgNumSlaveDevice: {
		0: {0, 0, fit, 0, fitUint16},
		1: {1, 0, fit, 1, fitUint16},
	},

	MesgNumCapabilities: {
		0:  {0, 255, fit, 0, fitUint8z},
		1:  {1, 255, fit, 1, fitUint8z},
		21: {2, 0, fit, 21, fitUint32z},
		23: {3, 0, fit, 23, fitUint32z},
	},

	MesgNumFileCapabilities: {
		254: {0, 0, fit, 254, fitUint16},
		0:   {1, 0, fit, 0, fitEnum},
		1:   {2, 0, fit, 1, fitUint8z},
		2:   {3, 0, fit, 2, fitString},
		3:   {4, 0, fit, 3, fitUint16},
		4:   {5, 0, fit, 4, fitUint32},
	},

	MesgNumMesgCapabilities: {
		254: {0, 0, fit, 254, fitUint16},
		0:   {1, 0, fit, 0, fitEnum},
		1:   {2, 0, fit, 1, fitUint16},
		2:   {3, 0, fit, 2, fitEnum},
		3:   {4, 0, fit, 3, fitUint16},
	},

	MesgNumFieldCapabilities: {
		254: {0, 0, fit, 254, fitUint16},
		0:   {1, 0, fit, 0, fitEnum},
		1:   {2, 0, fit, 1, fitUint16},
		2:   {3, 0, fit, 2, fitUint8},
		3:   {4, 0, fit, 3, fitUint16},
	},

	MesgNumDeviceSettings: {
		0: {0, 0, fit, 0, fitUint8},
		1: {1, 0, fit, 1, fitUint32},
		5: {2, 255, fit, 5, fitSint8},
	},

	MesgNumUserProfile: {
		254: {0, 0, fit, 254, fitUint16},
		0:   {1, 0, fit, 0, fitString},
		1:   {2, 0, fit, 1, fitEnum},
		2:   {3, 0, fit, 2, fitUint8},
		3:   {4, 0, fit, 3, fitUint8},
		4:   {5, 0, fit, 4, fitUint16},
		5:   {6, 0, fit, 5, fitEnum},
		6:   {7, 0, fit, 6, fitEnum},
		7:   {8, 0, fit, 7, fitEnum},
		8:   {9, 0, fit, 8, fitUint8},
		9:   {10, 0, fit, 9, fitUint8},
		10:  {11, 0, fit, 10, fitUint8},
		11:  {12, 0, fit, 11, fitUint8},
		12:  {13, 0, fit, 12, fitEnum},
		13:  {14, 0, fit, 13, fitEnum},
		14:  {15, 0, fit, 14, fitEnum},
		16:  {16, 0, fit, 16, fitEnum},
		17:  {17, 0, fit, 17, fitEnum},
		18:  {18, 0, fit, 18, fitEnum},
		21:  {19, 0, fit, 21, fitEnum},
		22:  {20, 0, fit, 22, fitUint16},
		23:  {21, 6, fit, 23, fitByte},
		30:  {22, 0, fit, 30, fitEnum},
	},

	MesgNumHrmProfile: {
		254: {0, 0, fit, 254, fitUint16},
		0:   {1, 0, fit, 0, fitEnum},
		1:   {2, 0, fit, 1, fitUint16z},
		2:   {3, 0, fit, 2, fitEnum},
		3:   {4, 0, fit, 3, fitUint8z},
	},

	MesgNumSdmProfile: {
		254: {0, 0, fit, 254, fitUint16},
		0:   {1, 0, fit, 0, fitEnum},
		1:   {2, 0, fit, 1, fitUint16z},
		2:   {3, 0, fit, 2, fitUint16},
		3:   {4, 0, fit, 3, fitUint32},
		4:   {5, 0, fit, 4, fitEnum},
		5:   {6, 0, fit, 5, fitUint8z},
		7:   {7, 0, fit, 7, fitUint8},
	},

	MesgNumBikeProfile: {
		254: {0, 0, fit, 254, fitUint16},
		0:   {1, 0, fit, 0, fitString},
		1:   {2, 0, fit, 1, fitEnum},
		2:   {3, 0, fit, 2, fitEnum},
		3:   {4, 0, fit, 3, fitUint32},
		4:   {5, 0, fit, 4, fitUint16z},
		5:   {6, 0, fit, 5, fitUint16z},
		6:   {7, 0, fit, 6, fitUint16z},
		7:   {8, 0, fit, 7, fitUint16z},
		8:   {9, 0, fit, 8, fitUint16},
		9:   {10, 0, fit, 9, fitUint16},
		10:  {11, 0, fit, 10, fitUint16},
		11:  {12, 0, fit, 11, fitUint16},
		12:  {13, 0, fit, 12, fitEnum},
		13:  {14, 0, fit, 13, fitEnum},
		14:  {15, 0, fit, 14, fitUint8},
		15:  {16, 0, fit, 15, fitEnum},
		16:  {17, 0, fit, 16, fitEnum},
		17:  {18, 0, fit, 17, fitEnum},
		18:  {19, 0, fit, 18, fitEnum},
		19:  {20, 0, fit, 19, fitUint8},
		20:  {21, 0, fit, 20, fitEnum},
		21:  {22, 0, fit, 21, fitUint8z},
		22:  {23, 0, fit, 22, fitUint8z},
		23:  {24, 0, fit, 23, fitUint8z},
		24:  {25, 0, fit, 24, fitUint8z},
		37:  {26, 0, fit, 37, fitUint8},
		38:  {27, 0, fit, 38, fitUint8z},
		39:  {28, 255, fit, 39, fitUint8z},
		40:  {29, 0, fit, 40, fitUint8z},
		41:  {30, 255, fit, 41, fitUint8z},
		44:  {31, 0, fit, 44, fitEnum},
	},

	MesgNumZonesTarget: {
		1: {0, 0, fit, 1, fitUint8},
		2: {1, 0, fit, 2, fitUint8},
		3: {2, 0, fit, 3, fitUint16},
		5: {3, 0, fit, 5, fitEnum},
		7: {4, 0, fit, 7, fitEnum},
	},

	MesgNumSport: {
		0: {0, 0, fit, 0, fitEnum},
		1: {1, 0, fit, 1, fitEnum},
		3: {2, 0, fit, 3, fitString},
	},

	MesgNumHrZone: {
		254: {0, 0, fit, 254, fitUint16},
		1:   {1, 0, fit, 1, fitUint8},
		2:   {2, 0, fit, 2, fitString},
	},

	MesgNumSpeedZone: {
		254: {0, 0, fit, 254, fitUint16},
		0:   {1, 0, fit, 0, fitUint16},
		1:   {2, 0, fit, 1, fitString},
	},

	MesgNumCadenceZone: {
		254: {0, 0, fit, 254, fitUint16},
		0:   {1, 0, fit, 0, fitUint8},
		1:   {2, 0, fit, 1, fitString},
	},

	MesgNumPowerZone: {
		254: {0, 0, fit, 254, fitUint16},
		1:   {1, 0, fit, 1, fitUint16},
		2:   {2, 0, fit, 2, fitString},
	},

	MesgNumMetZone: {
		254: {0, 0, fit, 254, fitUint16},
		1:   {1, 0, fit, 1, fitUint8},
		2:   {2, 0, fit, 2, fitUint16},
		3:   {3, 0, fit, 3, fitUint8},
	},

	MesgNumGoal: {
		254: {0, 0, fit, 254, fitUint16},
		0:   {1, 0, fit, 0, fitEnum},
		1:   {2, 0, fit, 1, fitEnum},
		2:   {3, 0, timeutc, 2, fitUint32},
		3:   {4, 0, timeutc, 3, fitUint32},
		4:   {5, 0, fit, 4, fitEnum},
		5:   {6, 0, fit, 5, fitUint32},
		6:   {7, 0, fit, 6, fitEnum},
		7:   {8, 0, fit, 7, fitUint32},
		8:   {9, 0, fit, 8, fitEnum},
		9:   {10, 0, fit, 9, fitUint16},
		10:  {11, 0, fit, 10, fitEnum},
	},

	MesgNumActivity: {
		253: {0, 0, timeutc, 253, fitUint32},
		0:   {1, 0, fit, 0, fitUint32},
		1:   {2, 0, fit, 1, fitUint16},
		2:   {3, 0, fit, 2, fitEnum},
		3:   {4, 0, fit, 3, fitEnum},
		4:   {5, 0, fit, 4, fitEnum},
		5:   {6, 0, timelocal, 5, fitUint32},
		6:   {7, 0, fit, 6, fitUint8},
	},

	MesgNumSession: {
		254: {0, 0, fit, 254, fitUint16},
		253: {1, 0, timeutc, 253, fitUint32},
		0:   {2, 0, fit, 0, fitEnum},
		1:   {3, 0, fit, 1, fitEnum},
		2:   {4, 0, timeutc, 2, fitUint32},
		3:   {5, 0, lat, 3, fitSint32},
		4:   {6, 0, lng, 4, fitSint32},
		5:   {7, 0, fit, 5, fitEnum},
		6:   {8, 0, fit, 6, fitEnum},
		7:   {9, 0, fit, 7, fitUint32},
		8:   {10, 0, fit, 8, fitUint32},
		9:   {11, 0, fit, 9, fitUint32},
		10:  {12, 0, fit, 10, fitUint32},
		11:  {13, 0, fit, 11, fitUint16},
		13:  {14, 0, fit, 13, fitUint16},
		14:  {15, 0, fit, 14, fitUint16},
		15:  {16, 0, fit, 15, fitUint16},
		16:  {17, 0, fit, 16, fitUint8},
		17:  {18, 0, fit, 17, fitUint8},
		18:  {19, 0, fit, 18, fitUint8},
		19:  {20, 0, fit, 19, fitUint8},
		20:  {21, 0, fit, 20, fitUint16},
		21:  {22, 0, fit, 21, fitUint16},
		22:  {23, 0, fit, 22, fitUint16},
		23:  {24, 0, fit, 23, fitUint16},
		24:  {25, 0, fit, 24, fitUint8},
		25:  {26, 0, fit, 25, fitUint16},
		26:  {27, 0, fit, 26, fitUint16},
		27:  {28, 0, fit, 27, fitUint8},
		28:  {29, 0, fit, 28, fitEnum},
		29:  {30, 0, lat, 29, fitSint32},
		30:  {31, 0, lng, 30, fitSint32},
		31:  {32, 0, lat, 31, fitSint32},
		32:  {33, 0, lng, 32, fitSint32},
		34:  {34, 0, fit, 34, fitUint16},
		35:  {35, 0, fit, 35, fitUint16},
		36:  {36, 0, fit, 36, fitUint16},
		37:  {37, 0, fit, 37, fitUint16},
		41:  {38, 0, fit, 41, fitUint32},
		42:  {39, 0, fit, 42, fitUint16},
		43:  {40, 0, fit, 43, fitEnum},
		44:  {41, 0, fit, 44, fitUint16},
		45:  {42, 0, fit, 45, fitUint16},
		46:  {43, 0, fit, 46, fitEnum},
		47:  {44, 0, fit, 47, fitUint16},
		48:  {45, 0, fit, 48, fitUint32},
		49:  {46, 0, fit, 49, fitUint16},
		50:  {47, 0, fit, 50, fitUint16},
		51:  {48, 0, fit, 51, fitUint8},
		52:  {49, 0, fit, 52, fitSint16},
		53:  {50, 0, fit, 53, fitSint16},
		54:  {51, 0, fit, 54, fitSint16},
		55:  {52, 0, fit, 55, fitSint16},
		56:  {53, 0, fit, 56, fitSint16},
		57:  {54, 0, fit, 57, fitSint8},
		58:  {55, 0, fit, 58, fitSint8},
		59:  {56, 0, fit, 59, fitUint32},
		60:  {57, 0, fit, 60, fitSint16},
		61:  {58, 0, fit, 61, fitSint16},
		62:  {59, 0, fit, 62, fitSint16},
		63:  {60, 0, fit, 63, fitSint16},
		64:  {61, 0, fit, 64, fitUint8},
		65:  {62, 255, fit, 65, fitUint32},
		66:  {63, 255, fit, 66, fitUint32},
		67:  {64, 255, fit, 67, fitUint32},
		68:  {65, 255, fit, 68, fitUint32},
		69:  {66, 0, fit, 69, fitUint32},
		70:  {67, 0, fit, 70, fitUint16},
		71:  {68, 0, fit, 71, fitUint16},
		82:  {69, 0, fit, 82, fitUint16},
		83:  {70, 0, fit, 83, fitUint16},
		84:  {71, 0, fit, 84, fitString},
		85:  {72, 255, fit, 85, fitUint16},
		86:  {73, 255, fit, 86, fitUint16},
		87:  {74, 0, fit, 87, fitUint16},
		88:  {75, 0, fit, 88, fitUint16},
		89:  {76, 0, fit, 89, fitUint16},
		90:  {77, 0, fit, 90, fitUint16},
		91:  {78, 0, fit, 91, fitUint16},
		92:  {79, 0, fit, 92, fitUint8},
		93:  {80, 0, fit, 93, fitUint8},
		94:  {81, 0, fit, 94, fitUint8},
		111: {82, 0, fit, 111, fitUint8},
		124: {83, 0, fit, 124, fitUint32},
		125: {84, 0, fit, 125, fitUint32},
		126: {85, 0, fit, 126, fitUint32},
		127: {86, 0, fit, 127, fitUint32},
		128: {87, 0, fit, 128, fitUint32},
	},

	MesgNumLap: {
		254: {0, 0, fit, 254, fitUint16},
		253: {1, 0, timeutc, 253, fitUint32},
		0:   {2, 0, fit, 0, fitEnum},
		1:   {3, 0, fit, 1, fitEnum},
		2:   {4, 0, timeutc, 2, fitUint32},
		3:   {5, 0, lat, 3, fitSint32},
		4:   {6, 0, lng, 4, fitSint32},
		5:   {7, 0, lat, 5, fitSint32},
		6:   {8, 0, lng, 6, fitSint32},
		7:   {9, 0, fit, 7, fitUint32},
		8:   {10, 0, fit, 8, fitUint32},
		9:   {11, 0, fit, 9, fitUint32},
		10:  {12, 0, fit, 10, fitUint32},
		11:  {13, 0, fit, 11, fitUint16},
		12:  {14, 0, fit, 12, fitUint16},
		13:  {15, 0, fit, 13, fitUint16},
		14:  {16, 0, fit, 14, fitUint16},
		15:  {17, 0, fit, 15, fitUint8},
		16:  {18, 0, fit, 16, fitUint8},
		17:  {19, 0, fit, 17, fitUint8},
		18:  {20, 0, fit, 18, fitUint8},
		19:  {21, 0, fit, 19, fitUint16},
		20:  {22, 0, fit, 20, fitUint16},
		21:  {23, 0, fit, 21, fitUint16},
		22:  {24, 0, fit, 22, fitUint16},
		23:  {25, 0, fit, 23, fitEnum},
		24:  {26, 0, fit, 24, fitEnum},
		25:  {27, 0, fit, 25, fitEnum},
		26:  {28, 0, fit, 26, fitUint8},
		32:  {29, 0, fit, 32, fitUint16},
		33:  {30, 0, fit, 33, fitUint16},
		34:  {31, 0, fit, 34, fitUint16},
		35:  {32, 0, fit, 35, fitUint16},
		37:  {33, 0, fit, 37, fitUint16},
		38:  {34, 0, fit, 38, fitEnum},
		39:  {35, 0, fit, 39, fitEnum},
		40:  {36, 0, fit, 40, fitUint16},
		41:  {37, 0, fit, 41, fitUint32},
		42:  {38, 0, fit, 42, fitUint16},
		43:  {39, 0, fit, 43, fitUint16},
		44:  {40, 0, fit, 44, fitUint8},
		45:  {41, 0, fit, 45, fitSint16},
		46:  {42, 0, fit, 46, fitSint16},
		47:  {43, 0, fit, 47, fitSint16},
		48:  {44, 0, fit, 48, fitSint16},
		49:  {45, 0, fit, 49, fitSint16},
		50:  {46, 0, fit, 50, fitSint8},
		51:  {47, 0, fit, 51, fitSint8},
		52:  {48, 0, fit, 52, fitUint32},
		53:  {49, 0, fit, 53, fitSint16},
		54:  {50, 0, fit, 54, fitSint16},
		55:  {51, 0, fit, 55, fitSint16},
		56:  {52, 0, fit, 56, fitSint16},
		57:  {53, 255, fit, 57, fitUint32},
		58:  {54, 255, fit, 58, fitUint32},
		59:  {55, 255, fit, 59, fitUint32},
		60:  {56, 255, fit, 60, fitUint32},
		61:  {57, 0, fit, 61, fitUint16},
		62:  {58, 0, fit, 62, fitUint16},
		63:  {59, 0, fit, 63, fitUint8},
		71:  {60, 0, fit, 71, fitUint16},
		74:  {61, 0, fit, 74, fitUint16},
		75:  {62, 255, fit, 75, fitUint16},
		76:  {63, 255, fit, 76, fitUint16},
		77:  {64, 0, fit, 77, fitUint16},
		78:  {65, 0, fit, 78, fitUint16},
		79:  {66, 0, fit, 79, fitUint16},
		80:  {67, 0, fit, 80, fitUint8},
		81:  {68, 0, fit, 81, fitUint8},
		82:  {69, 0, fit, 82, fitUint8},
		83:  {70, 0, fit, 83, fitUint16},
		84:  {71, 255, fit, 84, fitUint16},
		85:  {72, 255, fit, 85, fitUint16},
		86:  {73, 255, fit, 86, fitUint16},
		87:  {74, 255, fit, 87, fitUint16},
		88:  {75, 255, fit, 88, fitUint16},
		89:  {76, 255, fit, 89, fitUint16},
		110: {77, 0, fit, 110, fitUint32},
		111: {78, 0, fit, 111, fitUint32},
		112: {79, 0, fit, 112, fitUint32},
		113: {80, 0, fit, 113, fitUint32},
		114: {81, 0, fit, 114, fitUint32},
	},

	MesgNumLength: {
		254: {0, 0, fit, 254, fitUint16},
		253: {1, 0, timeutc, 253, fitUint32},
		0:   {2, 0, fit, 0, fitEnum},
		1:   {3, 0, fit, 1, fitEnum},
		2:   {4, 0, timeutc, 2, fitUint32},
		3:   {5, 0, fit, 3, fitUint32},
		4:   {6, 0, fit, 4, fitUint32},
		5:   {7, 0, fit, 5, fitUint16},
		6:   {8, 0, fit, 6, fitUint16},
		7:   {9, 0, fit, 7, fitEnum},
		9:   {10, 0, fit, 9, fitUint8},
		10:  {11, 0, fit, 10, fitUint8},
		11:  {12, 0, fit, 11, fitUint16},
		12:  {13, 0, fit, 12, fitEnum},
		18:  {14, 0, fit, 18, fitUint16},
		19:  {15, 0, fit, 19, fitUint16},
		20:  {16, 255, fit, 20, fitUint16},
		21:  {17, 255, fit, 21, fitUint16},
	},

	MesgNumRecord: {
		253: {0, 0, timeutc, 253, fitUint32},
		0:   {1, 0, lat, 0, fitSint32},
		1:   {2, 0, lng, 1, fitSint32},
		2:   {3, 0, fit, 2, fitUint16},
		3:   {4, 0, fit, 3, fitUint8},
		4:   {5, 0, fit, 4, fitUint8},
		5:   {6, 0, fit, 5, fitUint32},
		6:   {7, 0, fit, 6, fitUint16},
		7:   {8, 0, fit, 7, fitUint16},
		8:   {9, 3, fit, 8, fitByte},
		9:   {10, 0, fit, 9, fitSint16},
		10:  {11, 0, fit, 10, fitUint8},
		11:  {12, 0, fit, 11, fitSint32},
		12:  {13, 0, fit, 12, fitUint8},
		13:  {14, 0, fit, 13, fitSint8},
		17:  {15, 255, fit, 17, fitUint8},
		18:  {16, 0, fit, 18, fitUint8},
		19:  {17, 0, fit, 19, fitUint32},
		28:  {18, 0, fit, 28, fitUint16},
		29:  {19, 0, fit, 29, fitUint32},
		30:  {20, 0, fit, 30, fitUint8},
		31:  {21, 0, fit, 31, fitUint8},
		32:  {22, 0, fit, 32, fitSint16},
		33:  {23, 0, fit, 33, fitUint16},
		39:  {24, 0, fit, 39, fitUint16},
		40:  {25, 0, fit, 40, fitUint16},
		41:  {26, 0, fit, 41, fitUint16},
		42:  {27, 0, fit, 42, fitEnum},
		43:  {28, 0, fit, 43, fitUint8},
		44:  {29, 0, fit, 44, fitUint8},
		45:  {30, 0, fit, 45, fitUint8},
		46:  {31, 0, fit, 46, fitUint8},
		47:  {32, 0, fit, 47, fitUint8},
		48:  {33, 0, fit, 48, fitUint8},
		49:  {34, 0, fit, 49, fitEnum},
		50:  {35, 0, fit, 50, fitUint8},
		51:  {36, 0, fit, 51, fitUint16},
		52:  {37, 0, fit, 52, fitUint16},
		53:  {38, 0, fit, 53, fitUint8},
		54:  {39, 0, fit, 54, fitUint16},
		55:  {40, 0, fit, 55, fitUint16},
		56:  {41, 0, fit, 56, fitUint16},
		57:  {42, 0, fit, 57, fitUint16},
		58:  {43, 0, fit, 58, fitUint16},
		59:  {44, 0, fit, 59, fitUint16},
		62:  {45, 0, fit, 62, fitUint8},
		73:  {46, 0, fit, 73, fitUint32},
		78:  {47, 0, fit, 78, fitUint32},
	},

	MesgNumEvent: {
		253: {0, 0, timeutc, 253, fitUint32},
		0:   {1, 0, fit, 0, fitEnum},
		1:   {2, 0, fit, 1, fitEnum},
		2:   {3, 0, fit, 2, fitUint16},
		3:   {4, 0, fit, 3, fitUint32},
		4:   {5, 0, fit, 4, fitUint8},
		7:   {6, 0, fit, 7, fitUint16},
		8:   {7, 0, fit, 8, fitUint16},
		9:   {8, 0, fit, 9, fitUint8z},
		10:  {9, 0, fit, 10, fitUint8z},
		11:  {10, 0, fit, 11, fitUint8z},
		12:  {11, 0, fit, 12, fitUint8z},
	},

	MesgNumDeviceInfo: {
		253: {0, 0, timeutc, 253, fitUint32},
		0:   {1, 0, fit, 0, fitUint8},
		1:   {2, 0, fit, 1, fitUint8},
		2:   {3, 0, fit, 2, fitUint16},
		3:   {4, 0, fit, 3, fitUint32z},
		4:   {5, 0, fit, 4, fitUint16},
		5:   {6, 0, fit, 5, fitUint16},
		6:   {7, 0, fit, 6, fitUint8},
		7:   {8, 0, fit, 7, fitUint32},
		10:  {9, 0, fit, 10, fitUint16},
		11:  {10, 0, fit, 11, fitUint8},
		18:  {11, 0, fit, 18, fitEnum},
		19:  {12, 0, fit, 19, fitString},
		20:  {13, 0, fit, 20, fitUint8z},
		21:  {14, 0, fit, 21, fitUint16z},
		22:  {15, 0, fit, 22, fitEnum},
		25:  {16, 0, fit, 25, fitEnum},
		27:  {17, 0, fit, 27, fitString},
	},

	MesgNumTrainingFile: {
		253: {0, 0, timeutc, 253, fitUint32},
		0:   {1, 0, fit, 0, fitEnum},
		1:   {2, 0, fit, 1, fitUint16},
		2:   {3, 0, fit, 2, fitUint16},
		3:   {4, 0, fit, 3, fitUint32z},
		4:   {5, 0, timeutc, 4, fitUint32},
	},

	MesgNumHrv: {
		0: {0, 255, fit, 0, fitUint16},
	},

	MesgNumCameraEvent: {},

	MesgNumGyroscopeData: {},

	MesgNumAccelerometerData: {},

	MesgNumThreeDSensorCalibration: {},

	MesgNumVideoFrame: {},

	MesgNumObdiiData: {},

	MesgNumNmeaSentence: {
		253: {0, 0, timeutc, 253, fitUint32},
		0:   {1, 0, fit, 0, fitUint16},
		1:   {2, 0, fit, 1, fitString},
	},

	MesgNumAviationAttitude: {
		253: {0, 0, timeutc, 253, fitUint32},
		0:   {1, 0, fit, 0, fitUint16},
		1:   {2, 255, fit, 1, fitUint32},
		2:   {3, 255, fit, 2, fitSint16},
		3:   {4, 255, fit, 3, fitSint16},
		4:   {5, 255, fit, 4, fitSint16},
		5:   {6, 255, fit, 5, fitSint16},
		6:   {7, 255, fit, 6, fitSint16},
		7:   {8, 255, fit, 7, fitEnum},
		8:   {9, 255, fit, 8, fitUint8},
		9:   {10, 255, fit, 9, fitUint16},
		10:  {11, 255, fit, 10, fitUint16},
	},

	MesgNumVideo: {},

	MesgNumVideoTitle: {
		254: {0, 0, fit, 254, fitUint16},
		0:   {1, 0, fit, 0, fitUint16},
		1:   {2, 0, fit, 1, fitString},
	},

	MesgNumVideoDescription: {
		254: {0, 0, fit, 254, fitUint16},
		0:   {1, 0, fit, 0, fitUint16},
		1:   {2, 0, fit, 1, fitString},
	},

	MesgNumVideoClip: {},

	MesgNumCourse: {
		4: {0, 0, fit, 4, fitEnum},
		5: {1, 0, fit, 5, fitString},
		6: {2, 0, fit, 6, fitUint32z},
	},

	MesgNumCoursePoint: {
		254: {0, 0, fit, 254, fitUint16},
		1:   {1, 0, timeutc, 1, fitUint32},
		2:   {2, 0, lat, 2, fitSint32},
		3:   {3, 0, lng, 3, fitSint32},
		4:   {4, 0, fit, 4, fitUint32},
		5:   {5, 0, fit, 5, fitEnum},
		6:   {6, 0, fit, 6, fitString},
		8:   {7, 0, fit, 8, fitEnum},
	},

	MesgNumSegmentId: {
		0: {0, 0, fit, 0, fitString},
		1: {1, 0, fit, 1, fitString},
		2: {2, 0, fit, 2, fitEnum},
		3: {3, 0, fit, 3, fitEnum},
		4: {4, 0, fit, 4, fitUint32},
		5: {5, 0, fit, 5, fitUint32},
		6: {6, 0, fit, 6, fitUint8},
		7: {7, 0, fit, 7, fitEnum},
		8: {8, 0, fit, 8, fitEnum},
	},

	MesgNumSegmentLeaderboardEntry: {
		254: {0, 0, fit, 254, fitUint16},
		0:   {1, 0, fit, 0, fitString},
		1:   {2, 0, fit, 1, fitEnum},
		2:   {3, 0, fit, 2, fitUint32},
		3:   {4, 0, fit, 3, fitUint32},
		4:   {5, 0, fit, 4, fitUint32},
	},

	MesgNumSegmentPoint: {
		254: {0, 0, fit, 254, fitUint16},
		1:   {1, 0, lat, 1, fitSint32},
		2:   {2, 0, lng, 2, fitSint32},
		3:   {3, 0, fit, 3, fitUint32},
		4:   {4, 0, fit, 4, fitUint16},
		5:   {5, 255, fit, 5, fitUint32},
	},

	MesgNumSegmentLap: {
		254: {0, 0, fit, 254, fitUint16},
		253: {1, 0, timeutc, 253, fitUint32},
		0:   {2, 0, fit, 0, fitEnum},
		1:   {3, 0, fit, 1, fitEnum},
		2:   {4, 0, timeutc, 2, fitUint32},
		3:   {5, 0, lat, 3, fitSint32},
		4:   {6, 0, lng, 4, fitSint32},
		5:   {7, 0, lat, 5, fitSint32},
		6:   {8, 0, lng, 6, fitSint32},
		7:   {9, 0, fit, 7, fitUint32},
		8:   {10, 0, fit, 8, fitUint32},
		9:   {11, 0, fit, 9, fitUint32},
		10:  {12, 0, fit, 10, fitUint32},
		11:  {13, 0, fit, 11, fitUint16},
		12:  {14, 0, fit, 12, fitUint16},
		13:  {15, 0, fit, 13, fitUint16},
		14:  {16, 0, fit, 14, fitUint16},
		15:  {17, 0, fit, 15, fitUint8},
		16:  {18, 0, fit, 16, fitUint8},
		17:  {19, 0, fit, 17, fitUint8},
		18:  {20, 0, fit, 18, fitUint8},
		19:  {21, 0, fit, 19, fitUint16},
		20:  {22, 0, fit, 20, fitUint16},
		21:  {23, 0, fit, 21, fitUint16},
		22:  {24, 0, fit, 22, fitUint16},
		23:  {25, 0, fit, 23, fitEnum},
		24:  {26, 0, fit, 24, fitUint8},
		25:  {27, 0, lat, 25, fitSint32},
		26:  {28, 0, lng, 26, fitSint32},
		27:  {29, 0, lat, 27, fitSint32},
		28:  {30, 0, lng, 28, fitSint32},
		29:  {31, 0, fit, 29, fitString},
		30:  {32, 0, fit, 30, fitUint16},
		31:  {33, 0, fit, 31, fitUint16},
		32:  {34, 0, fit, 32, fitEnum},
		33:  {35, 0, fit, 33, fitUint32},
		34:  {36, 0, fit, 34, fitUint16},
		35:  {37, 0, fit, 35, fitUint16},
		36:  {38, 0, fit, 36, fitUint8},
		37:  {39, 0, fit, 37, fitSint16},
		38:  {40, 0, fit, 38, fitSint16},
		39:  {41, 0, fit, 39, fitSint16},
		40:  {42, 0, fit, 40, fitSint16},
		41:  {43, 0, fit, 41, fitSint16},
		42:  {44, 0, fit, 42, fitSint8},
		43:  {45, 0, fit, 43, fitSint8},
		44:  {46, 0, fit, 44, fitUint32},
		45:  {47, 0, fit, 45, fitSint16},
		46:  {48, 0, fit, 46, fitSint16},
		47:  {49, 0, fit, 47, fitSint16},
		48:  {50, 0, fit, 48, fitSint16},
		49:  {51, 255, fit, 49, fitUint32},
		50:  {52, 255, fit, 50, fitUint32},
		51:  {53, 255, fit, 51, fitUint32},
		52:  {54, 255, fit, 52, fitUint32},
		53:  {55, 0, fit, 53, fitUint16},
		54:  {56, 0, fit, 54, fitUint16},
		55:  {57, 0, fit, 55, fitUint8},
		56:  {58, 0, fit, 56, fitUint32},
		57:  {59, 0, fit, 57, fitUint16},
		58:  {60, 0, fit, 58, fitEnum},
		59:  {61, 0, fit, 59, fitUint8},
		60:  {62, 0, fit, 60, fitUint8},
		61:  {63, 0, fit, 61, fitUint8},
		62:  {64, 0, fit, 62, fitUint8},
		63:  {65, 0, fit, 63, fitUint8},
		64:  {66, 0, fit, 64, fitEnum},
		65:  {67, 0, fit, 65, fitString},
		66:  {68, 0, fit, 66, fitUint8},
		67:  {69, 0, fit, 67, fitUint8},
		68:  {70, 0, fit, 68, fitUint8},
		69:  {71, 0, fit, 69, fitUint16},
		70:  {72, 0, fit, 70, fitUint16},
	},

	MesgNumSegmentFile: {
		254: {0, 0, fit, 254, fitUint16},
		1:   {1, 0, fit, 1, fitString},
		3:   {2, 0, fit, 3, fitEnum},
		4:   {3, 0, fit, 4, fitUint32},
		7:   {4, 255, fit, 7, fitEnum},
		8:   {5, 255, fit, 8, fitUint32},
		9:   {6, 255, fit, 9, fitUint32},
	},

	MesgNumWorkout: {
		4: {0, 0, fit, 4, fitEnum},
		5: {1, 0, fit, 5, fitUint32z},
		6: {2, 0, fit, 6, fitUint16},
		8: {3, 0, fit, 8, fitString},
	},

	MesgNumWorkoutStep: {
		254: {0, 0, fit, 254, fitUint16},
		0:   {1, 0, fit, 0, fitString},
		1:   {2, 0, fit, 1, fitEnum},
		2:   {3, 0, fit, 2, fitUint32},
		3:   {4, 0, fit, 3, fitEnum},
		4:   {5, 0, fit, 4, fitUint32},
		5:   {6, 0, fit, 5, fitUint32},
		6:   {7, 0, fit, 6, fitUint32},
		7:   {8, 0, fit, 7, fitEnum},
	},

	MesgNumSchedule: {
		0: {0, 0, fit, 0, fitUint16},
		1: {1, 0, fit, 1, fitUint16},
		2: {2, 0, fit, 2, fitUint32z},
		3: {3, 0, timeutc, 3, fitUint32},
		4: {4, 0, fit, 4, fitEnum},
		5: {5, 0, fit, 5, fitEnum},
		6: {6, 0, timelocal, 6, fitUint32},
	},

	MesgNumTotals: {
		254: {0, 0, fit, 254, fitUint16},
		253: {1, 0, timeutc, 253, fitUint32},
		0:   {2, 0, fit, 0, fitUint32},
		1:   {3, 0, fit, 1, fitUint32},
		2:   {4, 0, fit, 2, fitUint32},
		3:   {5, 0, fit, 3, fitEnum},
		4:   {6, 0, fit, 4, fitUint32},
		5:   {7, 0, fit, 5, fitUint16},
		6:   {8, 0, fit, 6, fitUint32},
	},

	MesgNumWeightScale: {
		253: {0, 0, timeutc, 253, fitUint32},
		0:   {1, 0, fit, 0, fitUint16},
		1:   {2, 0, fit, 1, fitUint16},
		2:   {3, 0, fit, 2, fitUint16},
		3:   {4, 0, fit, 3, fitUint16},
		4:   {5, 0, fit, 4, fitUint16},
		5:   {6, 0, fit, 5, fitUint16},
		7:   {7, 0, fit, 7, fitUint16},
		8:   {8, 0, fit, 8, fitUint8},
		9:   {9, 0, fit, 9, fitUint16},
		10:  {10, 0, fit, 10, fitUint8},
		11:  {11, 0, fit, 11, fitUint8},
		12:  {12, 0, fit, 12, fitUint16},
	},

	MesgNumBloodPressure: {
		253: {0, 0, timeutc, 253, fitUint32},
		0:   {1, 0, fit, 0, fitUint16},
		1:   {2, 0, fit, 1, fitUint16},
		2:   {3, 0, fit, 2, fitUint16},
		3:   {4, 0, fit, 3, fitUint16},
		4:   {5, 0, fit, 4, fitUint16},
		5:   {6, 0, fit, 5, fitUint16},
		6:   {7, 0, fit, 6, fitUint8},
		7:   {8, 0, fit, 7, fitEnum},
		8:   {9, 0, fit, 8, fitEnum},
		9:   {10, 0, fit, 9, fitUint16},
	},

	MesgNumMonitoringInfo: {
		253: {0, 0, timeutc, 253, fitUint32},
		0:   {1, 0, timelocal, 0, fitUint32},
	},

	MesgNumMonitoring: {
		253: {0, 0, timeutc, 253, fitUint32},
		0:   {1, 0, fit, 0, fitUint8},
		1:   {2, 0, fit, 1, fitUint16},
		2:   {3, 0, fit, 2, fitUint32},
		3:   {4, 0, fit, 3, fitUint32},
		4:   {5, 0, fit, 4, fitUint32},
		5:   {6, 0, fit, 5, fitEnum},
		6:   {7, 0, fit, 6, fitEnum},
		8:   {8, 0, fit, 8, fitUint16},
		9:   {9, 0, fit, 9, fitUint16},
		10:  {10, 0, fit, 10, fitUint16},
		11:  {11, 0, timelocal, 11, fitUint32},
	},

	MesgNumMemoGlob: {},
}

func getField(gmn MesgNum, fdn byte) (*field, bool) {
	if int(gmn) >= len(_fields) {
		return nil, false
	}
	f := _fields[gmn][fdn]
	if f == nil {
		return nil, false
	}
	return f, true
}

var msgsAllInvalid = [...]reflect.Value{
	MesgNumFileId: reflect.ValueOf(&FileIdMsg{
		0xFF,
		0xFFFF,
		0xFFFF,
		0xFFFFFFFF,
		timeBase,
		0xFFFF,
		"",
	}),
	MesgNumFileCreator: reflect.ValueOf(&FileCreatorMsg{
		0xFFFF,
		0xFF,
	}),
	MesgNumTimestampCorrelation: reflect.ValueOf(&TimestampCorrelationMsg{}),
	MesgNumSoftware: reflect.ValueOf(&SoftwareMsg{
		0xFFFF,
		0xFFFF,
		"",
	}),
	MesgNumSlaveDevice: reflect.ValueOf(&SlaveDeviceMsg{
		0xFFFF,
		0xFFFF,
	}),
	MesgNumCapabilities: reflect.ValueOf(&CapabilitiesMsg{
		nil,
		nil,
		0x00000000,
		0x00000000,
	}),
	MesgNumFileCapabilities: reflect.ValueOf(&FileCapabilitiesMsg{
		0xFFFF,
		0xFF,
		0x00,
		"",
		0xFFFF,
		0xFFFFFFFF,
	}),
	MesgNumMesgCapabilities: reflect.ValueOf(&MesgCapabilitiesMsg{
		0xFFFF,
		0xFF,
		0xFFFF,
		0xFF,
		0xFFFF,
	}),
	MesgNumFieldCapabilities: reflect.ValueOf(&FieldCapabilitiesMsg{
		0xFFFF,
		0xFF,
		0xFFFF,
		0xFF,
		0xFFFF,
	}),
	MesgNumDeviceSettings: reflect.ValueOf(&DeviceSettingsMsg{
		0xFF,
		0xFFFFFFFF,
		nil,
	}),
	MesgNumUserProfile: reflect.ValueOf(&UserProfileMsg{
		0xFFFF,
		"",
		0xFF,
		0xFF,
		0xFF,
		0xFFFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		0xFFFF,
		nil,
		0xFF,
	}),
	MesgNumHrmProfile: reflect.ValueOf(&HrmProfileMsg{
		0xFFFF,
		0xFF,
		0xFFFF,
		0xFF,
		0xFF,
	}),
	MesgNumSdmProfile: reflect.ValueOf(&SdmProfileMsg{
		0xFFFF,
		0xFF,
		0xFFFF,
		0xFFFF,
		0xFFFFFFFF,
		0xFF,
		0xFF,
		0xFF,
	}),
	MesgNumBikeProfile: reflect.ValueOf(&BikeProfileMsg{
		0xFFFF,
		"",
		0xFF,
		0xFF,
		0xFFFFFFFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		nil,
		0xFF,
		nil,
		0xFF,
	}),
	MesgNumZonesTarget: reflect.ValueOf(&ZonesTargetMsg{
		0xFF,
		0xFF,
		0xFFFF,
		0xFF,
		0xFF,
	}),
	MesgNumSport: reflect.ValueOf(&SportMsg{
		0xFF,
		0xFF,
		"",
	}),
	MesgNumHrZone: reflect.ValueOf(&HrZoneMsg{
		0xFFFF,
		0xFF,
		"",
	}),
	MesgNumSpeedZone: reflect.ValueOf(&SpeedZoneMsg{
		0xFFFF,
		0xFFFF,
		"",
	}),
	MesgNumCadenceZone: reflect.ValueOf(&CadenceZoneMsg{
		0xFFFF,
		0xFF,
		"",
	}),
	MesgNumPowerZone: reflect.ValueOf(&PowerZoneMsg{
		0xFFFF,
		0xFFFF,
		"",
	}),
	MesgNumMetZone: reflect.ValueOf(&MetZoneMsg{
		0xFFFF,
		0xFF,
		0xFFFF,
		0xFF,
	}),
	MesgNumGoal: reflect.ValueOf(&GoalMsg{
		0xFFFF,
		0xFF,
		0xFF,
		timeBase,
		timeBase,
		0xFF,
		0xFFFFFFFF,
		0xFF,
		0xFFFFFFFF,
		0xFF,
		0xFFFF,
		0xFF,
	}),
	MesgNumActivity: reflect.ValueOf(&ActivityMsg{
		timeBase,
		0xFFFFFFFF,
		0xFFFF,
		0xFF,
		0xFF,
		0xFF,
		timeBase,
		0xFF,
	}),
	MesgNumSession: reflect.ValueOf(&SessionMsg{
		0xFFFF,
		timeBase,
		0xFF,
		0xFF,
		timeBase,
		0x7FFFFFFF,
		0x7FFFFFFF,
		0xFF,
		0xFF,
		0xFFFFFFFF,
		0xFFFFFFFF,
		0xFFFFFFFF,
		0xFFFFFFFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFF,
		0xFFFF,
		0xFFFF,
		0xFF,
		0xFF,
		0x7FFFFFFF,
		0x7FFFFFFF,
		0x7FFFFFFF,
		0x7FFFFFFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFFFFFFFF,
		0xFFFF,
		0xFF,
		0xFFFF,
		0xFFFF,
		0xFF,
		0xFFFF,
		0xFFFFFFFF,
		0xFFFF,
		0xFFFF,
		0xFF,
		0x7FFF,
		0x7FFF,
		0x7FFF,
		0x7FFF,
		0x7FFF,
		0x7F,
		0x7F,
		0xFFFFFFFF,
		0x7FFF,
		0x7FFF,
		0x7FFF,
		0x7FFF,
		0xFF,
		nil,
		nil,
		nil,
		nil,
		0xFFFFFFFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		"",
		nil,
		nil,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		0xFFFFFFFF,
		0xFFFFFFFF,
		0xFFFFFFFF,
		0xFFFFFFFF,
		0xFFFFFFFF,
	}),
	MesgNumLap: reflect.ValueOf(&LapMsg{
		0xFFFF,
		timeBase,
		0xFF,
		0xFF,
		timeBase,
		0x7FFFFFFF,
		0x7FFFFFFF,
		0x7FFFFFFF,
		0x7FFFFFFF,
		0xFFFFFFFF,
		0xFFFFFFFF,
		0xFFFFFFFF,
		0xFFFFFFFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFF,
		0xFF,
		0xFFFF,
		0xFFFFFFFF,
		0xFFFF,
		0xFFFF,
		0xFF,
		0x7FFF,
		0x7FFF,
		0x7FFF,
		0x7FFF,
		0x7FFF,
		0x7F,
		0x7F,
		0xFFFFFFFF,
		0x7FFF,
		0x7FFF,
		0x7FFF,
		0x7FFF,
		nil,
		nil,
		nil,
		nil,
		0xFFFF,
		0xFFFF,
		0xFF,
		0xFFFF,
		0xFFFF,
		nil,
		nil,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFF,
		0xFF,
		0xFF,
		0xFFFF,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		0xFFFFFFFF,
		0xFFFFFFFF,
		0xFFFFFFFF,
		0xFFFFFFFF,
		0xFFFFFFFF,
	}),
	MesgNumLength: reflect.ValueOf(&LengthMsg{
		0xFFFF,
		timeBase,
		0xFF,
		0xFF,
		timeBase,
		0xFFFFFFFF,
		0xFFFFFFFF,
		0xFFFF,
		0xFFFF,
		0xFF,
		0xFF,
		0xFF,
		0xFFFF,
		0xFF,
		0xFFFF,
		0xFFFF,
		nil,
		nil,
	}),
	MesgNumRecord: reflect.ValueOf(&RecordMsg{
		timeBase,
		0x7FFFFFFF,
		0x7FFFFFFF,
		0xFFFF,
		0xFF,
		0xFF,
		0xFFFFFFFF,
		0xFFFF,
		0xFFFF,
		nil,
		0x7FFF,
		0xFF,
		0x7FFFFFFF,
		0xFF,
		0x7F,
		nil,
		0xFF,
		0xFFFFFFFF,
		0xFFFF,
		0xFFFFFFFF,
		0xFF,
		0xFF,
		0x7FFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		0xFFFF,
		0xFFFF,
		0xFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFF,
		0xFFFFFFFF,
		0xFFFFFFFF,
	}),
	MesgNumEvent: reflect.ValueOf(&EventMsg{
		timeBase,
		0xFF,
		0xFF,
		0xFFFF,
		0xFFFFFFFF,
		0xFF,
		0xFFFF,
		0xFFFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
	}),
	MesgNumDeviceInfo: reflect.ValueOf(&DeviceInfoMsg{
		timeBase,
		0xFF,
		0xFF,
		0xFFFF,
		0xFFFFFFFF,
		0xFFFF,
		0xFFFF,
		0xFF,
		0xFFFFFFFF,
		0xFFFF,
		0xFF,
		0xFF,
		"",
		0xFF,
		0xFFFF,
		0xFF,
		0xFF,
		"",
	}),
	MesgNumTrainingFile: reflect.ValueOf(&TrainingFileMsg{
		timeBase,
		0xFF,
		0xFFFF,
		0xFFFF,
		0xFFFFFFFF,
		timeBase,
	}),
	MesgNumHrv: reflect.ValueOf(&HrvMsg{
		nil,
	}),
	MesgNumCameraEvent:             reflect.ValueOf(&CameraEventMsg{}),
	MesgNumGyroscopeData:           reflect.ValueOf(&GyroscopeDataMsg{}),
	MesgNumAccelerometerData:       reflect.ValueOf(&AccelerometerDataMsg{}),
	MesgNumThreeDSensorCalibration: reflect.ValueOf(&ThreeDSensorCalibrationMsg{}),
	MesgNumVideoFrame:              reflect.ValueOf(&VideoFrameMsg{}),
	MesgNumObdiiData:               reflect.ValueOf(&ObdiiDataMsg{}),
	MesgNumNmeaSentence: reflect.ValueOf(&NmeaSentenceMsg{
		timeBase,
		0xFFFF,
		"",
	}),
	MesgNumAviationAttitude: reflect.ValueOf(&AviationAttitudeMsg{
		timeBase,
		0xFFFF,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
	}),
	MesgNumVideo: reflect.ValueOf(&VideoMsg{}),
	MesgNumVideoTitle: reflect.ValueOf(&VideoTitleMsg{
		0xFFFF,
		0xFFFF,
		"",
	}),
	MesgNumVideoDescription: reflect.ValueOf(&VideoDescriptionMsg{
		0xFFFF,
		0xFFFF,
		"",
	}),
	MesgNumVideoClip: reflect.ValueOf(&VideoClipMsg{}),
	MesgNumCourse: reflect.ValueOf(&CourseMsg{
		0xFF,
		"",
		0x00000000,
	}),
	MesgNumCoursePoint: reflect.ValueOf(&CoursePointMsg{
		0xFFFF,
		timeBase,
		0x7FFFFFFF,
		0x7FFFFFFF,
		0xFFFFFFFF,
		0xFF,
		"",
		0xFF,
	}),
	MesgNumSegmentId: reflect.ValueOf(&SegmentIdMsg{
		"",
		"",
		0xFF,
		0xFF,
		0xFFFFFFFF,
		0xFFFFFFFF,
		0xFF,
		0xFF,
		0xFF,
	}),
	MesgNumSegmentLeaderboardEntry: reflect.ValueOf(&SegmentLeaderboardEntryMsg{
		0xFFFF,
		"",
		0xFF,
		0xFFFFFFFF,
		0xFFFFFFFF,
		0xFFFFFFFF,
	}),
	MesgNumSegmentPoint: reflect.ValueOf(&SegmentPointMsg{
		0xFFFF,
		0x7FFFFFFF,
		0x7FFFFFFF,
		0xFFFFFFFF,
		0xFFFF,
		nil,
	}),
	MesgNumSegmentLap: reflect.ValueOf(&SegmentLapMsg{
		0xFFFF,
		timeBase,
		0xFF,
		0xFF,
		timeBase,
		0x7FFFFFFF,
		0x7FFFFFFF,
		0x7FFFFFFF,
		0x7FFFFFFF,
		0xFFFFFFFF,
		0xFFFFFFFF,
		0xFFFFFFFF,
		0xFFFFFFFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFF,
		0xFF,
		0x7FFFFFFF,
		0x7FFFFFFF,
		0x7FFFFFFF,
		0x7FFFFFFF,
		"",
		0xFFFF,
		0xFFFF,
		0xFF,
		0xFFFFFFFF,
		0xFFFF,
		0xFFFF,
		0xFF,
		0x7FFF,
		0x7FFF,
		0x7FFF,
		0x7FFF,
		0x7FFF,
		0x7F,
		0x7F,
		0xFFFFFFFF,
		0x7FFF,
		0x7FFF,
		0x7FFF,
		0x7FFF,
		nil,
		nil,
		nil,
		nil,
		0xFFFF,
		0xFFFF,
		0xFF,
		0xFFFFFFFF,
		0xFFFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		"",
		0xFF,
		0xFF,
		0xFF,
		0xFFFF,
		0xFFFF,
	}),
	MesgNumSegmentFile: reflect.ValueOf(&SegmentFileMsg{
		0xFFFF,
		"",
		0xFF,
		0xFFFFFFFF,
		nil,
		nil,
		nil,
	}),
	MesgNumWorkout: reflect.ValueOf(&WorkoutMsg{
		0xFF,
		0x00000000,
		0xFFFF,
		"",
	}),
	MesgNumWorkoutStep: reflect.ValueOf(&WorkoutStepMsg{
		0xFFFF,
		"",
		0xFF,
		0xFFFFFFFF,
		0xFF,
		0xFFFFFFFF,
		0xFFFFFFFF,
		0xFFFFFFFF,
		0xFF,
	}),
	MesgNumSchedule: reflect.ValueOf(&ScheduleMsg{
		0xFFFF,
		0xFFFF,
		0xFFFFFFFF,
		timeBase,
		0xFF,
		0xFF,
		timeBase,
	}),
	MesgNumTotals: reflect.ValueOf(&TotalsMsg{
		0xFFFF,
		timeBase,
		0xFFFFFFFF,
		0xFFFFFFFF,
		0xFFFFFFFF,
		0xFF,
		0xFFFFFFFF,
		0xFFFF,
		0xFFFFFFFF,
	}),
	MesgNumWeightScale: reflect.ValueOf(&WeightScaleMsg{
		timeBase,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFF,
		0xFFFF,
		0xFF,
		0xFF,
		0xFFFF,
	}),
	MesgNumBloodPressure: reflect.ValueOf(&BloodPressureMsg{
		timeBase,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFF,
		0xFF,
		0xFF,
		0xFFFF,
	}),
	MesgNumMonitoringInfo: reflect.ValueOf(&MonitoringInfoMsg{
		timeBase,
		timeBase,
	}),
	MesgNumMonitoring: reflect.ValueOf(&MonitoringMsg{
		timeBase,
		0xFF,
		0xFFFF,
		0xFFFFFFFF,
		0xFFFFFFFF,
		0xFFFFFFFF,
		0xFF,
		0xFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		timeBase,
	}),
	MesgNumMemoGlob: reflect.ValueOf(&MemoGlobMsg{}),
}

func getMesgAllInvalid(mn MesgNum) reflect.Value {
	return reflect.ValueOf(msgsAllInvalid[mn].Interface()).Elem()
}
