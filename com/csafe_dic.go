package com

/*
#Copyright (c) 2011, Sam Gambrell
#Licensed under the Simplified BSD License.
#NOTE: This code has not been thoroughly tested and may not function as advertised.
# Please report and findings to the author so that they may be addressed in a stable release.
*/

// Unique Frame Flags
const ExtendedFrameStartFlag = byte(0xF0)
const StandardFrameStartFlag = byte(0xF1)
const StopFrameFlag = byte(0xF2)
const ByteStuffingFlag = byte(0xF3)

type Cmd struct {
	ID       byte
	Variable *[]byte
	Data     *[]byte
}

// Command IDs
const (
	CSAFE_GETSTATUS_CMD byte = iota
	CSAFE_RESET_CMD
	CSAFE_GOIDLE_CMD
	CSAFE_GOHAVEID_CMD
	CSAFE_GOINUSE_CMD
	CSAFE_GOFINISHED_CMD
	CSAFE_GOREADY_CMD
	CSAFE_BADID_CMD
	CSAFE_GETVERSION_CMD
	CSAFE_GETID_CMD
	CSAFE_GETUNITS_CMD
	CSAFE_GETSERIAL_CMD
	CSAFE_GETODOMETER_CMD
	CSAFE_GETERRORCODE_CMD
	CSAFE_GETTWORK_CMD
	CSAFE_GETHORIZONTAL_CMD
	CSAFE_GETCALORIES_CMD
	CSAFE_GETPROGRAM_CMD
	CSAFE_GETPACE_CMD
	CSAFE_GETCADENCE_CMD
	CSAFE_GETUSERINFO_CMD
	CSAFE_GETHRCUR_CMD
	CSAFE_GETPOWER_CMD
	CSAFE_AUTOUPLOAD_CMD
	CSAFE_IDDIGITS_CMD
	CSAFE_SETTIME_CMD
	CSAFE_SETDATE_CMD
	CSAFE_SETTIMEOUT_CMD
	CSAFE_SETUSERCFG1_CMD
	CSAFE_SETTWORK_CMD
	CSAFE_SETHORIZONTAL_CMD
	CSAFE_SETCALORIES_CMD
	CSAFE_SETPROGRAM_CMD
	CSAFE_SETPOWER_CMD
	CSAFE_GETCAPS_CMD
	CSAFE_PM_GET_WORKOUTTYPE
	CSAFE_PM_GET_DRAGFACTOR
	CSAFE_PM_GET_STROKESTATE
	CSAFE_PM_GET_WORKTIME
	CSAFE_PM_GET_WORKDISTANCE
	CSAFE_PM_GET_ERRORVALUE
	CSAFE_PM_GET_WORKOUTSTATE
	CSAFE_PM_GET_WORKOUTINTERVALCOUNT
	CSAFE_PM_GET_INTERVALTYPE
	CSAFE_PM_GET_RESTTIME
	CSAFE_PM_SET_SPLITDURATION
	CSAFE_PM_GET_FORCEPLOTDATA
	CSAFE_PM_SET_SCREENERRORMODE
	CSAFE_PM_GET_HEARTBEATDATA
)

var Cmds = map[byte]Cmd{
	// Short Commands
	CSAFE_GETSTATUS_CMD:     Cmd{0x80, nil, nil},
	CSAFE_RESET_CMD:         Cmd{0x81, nil, nil},
	CSAFE_GOIDLE_CMD:        Cmd{0x82, nil, nil},
	CSAFE_GOHAVEID_CMD:      Cmd{0x83, nil, nil},
	CSAFE_GOINUSE_CMD:       Cmd{0x85, nil, nil},
	CSAFE_GOFINISHED_CMD:    Cmd{0x86, nil, nil},
	CSAFE_GOREADY_CMD:       Cmd{0x87, nil, nil},
	CSAFE_BADID_CMD:         Cmd{0x88, nil, nil},
	CSAFE_GETVERSION_CMD:    Cmd{0x91, nil, nil},
	CSAFE_GETID_CMD:         Cmd{0x92, nil, nil},
	CSAFE_GETUNITS_CMD:      Cmd{0x93, nil, nil},
	CSAFE_GETSERIAL_CMD:     Cmd{0x94, nil, nil},
	CSAFE_GETODOMETER_CMD:   Cmd{0x9B, nil, nil},
	CSAFE_GETERRORCODE_CMD:  Cmd{0x9C, nil, nil},
	CSAFE_GETTWORK_CMD:      Cmd{0xA0, nil, nil},
	CSAFE_GETHORIZONTAL_CMD: Cmd{0xA1, nil, nil},
	CSAFE_GETCALORIES_CMD:   Cmd{0xA3, nil, nil},
	CSAFE_GETPROGRAM_CMD:    Cmd{0xA4, nil, nil},
	CSAFE_GETPACE_CMD:       Cmd{0xA6, nil, nil},
	CSAFE_GETCADENCE_CMD:    Cmd{0xA7, nil, nil},
	CSAFE_GETUSERINFO_CMD:   Cmd{0xAB, nil, nil},
	CSAFE_GETHRCUR_CMD:      Cmd{0xB0, nil, nil},
	CSAFE_GETPOWER_CMD:      Cmd{0xB4, nil, nil},

	// Long Commands
	CSAFE_AUTOUPLOAD_CMD:    Cmd{0x01, &[]byte{1}, nil},       // Configuration (no affect)
	CSAFE_IDDIGITS_CMD:      Cmd{0x10, &[]byte{1}, nil},       // Number of Digits
	CSAFE_SETTIME_CMD:       Cmd{0x11, &[]byte{1, 1, 1}, nil}, // Hour, Minute, Seconds
	CSAFE_SETDATE_CMD:       Cmd{0x12, &[]byte{1, 1, 1}, nil}, // Year, Month, Day
	CSAFE_SETTIMEOUT_CMD:    Cmd{0x13, &[]byte{1}, nil},       // State Timeout
	CSAFE_SETUSERCFG1_CMD:   Cmd{0x1A, &[]byte{0}, nil},       // PM3 Specific Command (length computed)
	CSAFE_SETTWORK_CMD:      Cmd{0x20, &[]byte{1, 1, 1}, nil}, // Hour, Minute, Seconds
	CSAFE_SETHORIZONTAL_CMD: Cmd{0x21, &[]byte{2, 1}, nil},    // Distance, Units
	CSAFE_SETCALORIES_CMD:   Cmd{0x23, &[]byte{2}, nil},       // Total Calories
	CSAFE_SETPROGRAM_CMD:    Cmd{0x24, &[]byte{1, 1}, nil},    // Workout ID, N/A
	CSAFE_SETPOWER_CMD:      Cmd{0x34, &[]byte{2, 1}, nil},    // Stroke Watts, Units
	CSAFE_GETCAPS_CMD:       Cmd{0x70, &[]byte{1}, nil},       // Capability Code

	// PM3 Specific Short Commands
	CSAFE_PM_GET_WORKOUTTYPE:          Cmd{0x89, nil, &[]byte{0x1A}},
	CSAFE_PM_GET_DRAGFACTOR:           Cmd{0xC1, nil, &[]byte{0x1A}},
	CSAFE_PM_GET_STROKESTATE:          Cmd{0xBF, nil, &[]byte{0x1A}},
	CSAFE_PM_GET_WORKTIME:             Cmd{0xA0, nil, &[]byte{0x1A}},
	CSAFE_PM_GET_WORKDISTANCE:         Cmd{0xA3, nil, &[]byte{0x1A}},
	CSAFE_PM_GET_ERRORVALUE:           Cmd{0xC9, nil, &[]byte{0x1A}},
	CSAFE_PM_GET_WORKOUTSTATE:         Cmd{0x8D, nil, &[]byte{0x1A}},
	CSAFE_PM_GET_WORKOUTINTERVALCOUNT: Cmd{0x9F, nil, &[]byte{0x1A}},
	CSAFE_PM_GET_INTERVALTYPE:         Cmd{0x8E, nil, &[]byte{0x1A}},
	CSAFE_PM_GET_RESTTIME:             Cmd{0xCF, nil, &[]byte{0x1A}},

	// PM3 Specific Long Commands
	CSAFE_PM_SET_SPLITDURATION:   Cmd{0x05, &[]byte{1, 4}, &[]byte{0x1A}}, //Time(0)/Distance(128), Duration
	CSAFE_PM_GET_FORCEPLOTDATA:   Cmd{0x6B, &[]byte{1}, &[]byte{0x1A}},    //Block Length
	CSAFE_PM_SET_SCREENERRORMODE: Cmd{0x27, &[]byte{1}, &[]byte{0x1A}},    //Disable(0)/Enable(1)
	CSAFE_PM_GET_HEARTBEATDATA:   Cmd{0x6C, &[]byte{1}, &[]byte{0x1A}},    //Block Length
}

/*
resp[0xCmd_Id] = [COMMAND_NAME, [Bytes, ...]]
negative number for ASCII
use absolute max number for variable, (getid & getcaps)
*/

type Resp struct {
	ID    byte
	Bytes *[]byte
	ASCII bool
}

var Resps = map[uint]Resp{
	// Response Data to Short Commands
	0x80: Resp{CSAFE_GETSTATUS_CMD, &[]byte{0}, false}, //Status
	0x81: Resp{CSAFE_RESET_CMD, &[]byte{0}, false},
	0x82: Resp{CSAFE_GOIDLE_CMD, &[]byte{0}, false},
	0x83: Resp{CSAFE_GOHAVEID_CMD, &[]byte{0}, false},
	0x85: Resp{CSAFE_GOINUSE_CMD, &[]byte{0}, false},
	0x86: Resp{CSAFE_GOFINISHED_CMD, &[]byte{0}, false},
	0x87: Resp{CSAFE_GOREADY_CMD, &[]byte{0}, false},
	0x88: Resp{CSAFE_BADID_CMD, &[]byte{0}, false},
	0x91: Resp{CSAFE_GETVERSION_CMD, &[]byte{1, 1, 1, 2, 2}, false}, //Mfg ID, CID, Model, HW Version, SW Version
	0x92: Resp{CSAFE_GETID_CMD, &[]byte{5}, true},                   //ASCII Digit (variable)
	0x93: Resp{CSAFE_GETUNITS_CMD, &[]byte{1}, false},               //Units Type
	0x94: Resp{CSAFE_GETSERIAL_CMD, &[]byte{9}, true},               //ASCII Serial Number
	0x9B: Resp{CSAFE_GETODOMETER_CMD, &[]byte{4, 1}, false},         //Distance, Units Specifier
	0x9C: Resp{CSAFE_GETERRORCODE_CMD, &[]byte{3}, false},           //Error Code
	0xA0: Resp{CSAFE_GETTWORK_CMD, &[]byte{1, 1, 1}, false},         //Hours, Minutes, Seconds
	0xA1: Resp{CSAFE_GETHORIZONTAL_CMD, &[]byte{2, 1}, false},       //Distance, Units Specifier
	0xA3: Resp{CSAFE_GETCALORIES_CMD, &[]byte{2}, false},            //Total Calories
	0xA4: Resp{CSAFE_GETPROGRAM_CMD, &[]byte{1}, false},             //Program Number
	0xA6: Resp{CSAFE_GETPACE_CMD, &[]byte{2, 1}, false},             //Stroke Pace, Units Specifier
	0xA7: Resp{CSAFE_GETCADENCE_CMD, &[]byte{2, 1}, false},          //Stroke Rate, Units Specifier
	0xAB: Resp{CSAFE_GETUSERINFO_CMD, &[]byte{2, 1, 1, 1}, false},   //Weight, Units Specifier, Age, Gender
	0xB0: Resp{CSAFE_GETHRCUR_CMD, &[]byte{1}, false},               //Beats/Min
	0xB4: Resp{CSAFE_GETPOWER_CMD, &[]byte{2, 1}, false},            //Stroke Watts

	// Response Data to Long Commands
	0x01: Resp{CSAFE_AUTOUPLOAD_CMD, &[]byte{0}, false},
	0x10: Resp{CSAFE_IDDIGITS_CMD, &[]byte{0}, false},
	0x11: Resp{CSAFE_SETTIME_CMD, &[]byte{0}, false},
	0x12: Resp{CSAFE_SETDATE_CMD, &[]byte{0}, false},
	0x13: Resp{CSAFE_SETTIMEOUT_CMD, &[]byte{0}, false},
	0x1A: Resp{CSAFE_SETUSERCFG1_CMD, &[]byte{0}, false}, //PM3 Specific Command ID
	0x20: Resp{CSAFE_SETTWORK_CMD, &[]byte{0}, false},
	0x21: Resp{CSAFE_SETHORIZONTAL_CMD, &[]byte{0}, false},
	0x23: Resp{CSAFE_SETCALORIES_CMD, &[]byte{0}, false},
	0x24: Resp{CSAFE_SETPROGRAM_CMD, &[]byte{0}, false},
	0x34: Resp{CSAFE_SETPOWER_CMD, &[]byte{0}, false},
	0x70: Resp{CSAFE_GETCAPS_CMD, &[]byte{11}, false}, //Depended on Capability Code (variable)

	//Response Data to PM3 Specific Short Commands
	0x1A89: Resp{CSAFE_PM_GET_WORKOUTTYPE, &[]byte{1}, false}, //Workout Type
	0x1AC1: Resp{CSAFE_PM_GET_DRAGFACTOR, &[]byte{1}, false},  //Drag Factor
	0x1ABF: Resp{CSAFE_PM_GET_STROKESTATE, &[]byte{1}, false}, //Stroke State
	//Work Time (seconds * 100), Fractional Work Time (1/100)
	0x1AA0: Resp{CSAFE_PM_GET_WORKTIME, &[]byte{4, 1}, false},
	//Work Distance (meters * 10), Fractional Work Distance (1/10)
	0x1AA3: Resp{CSAFE_PM_GET_WORKDISTANCE, &[]byte{4, 1}, false},
	0x1AC9: Resp{CSAFE_PM_GET_ERRORVALUE, &[]byte{2}, false},           //Error Value
	0x1A8D: Resp{CSAFE_PM_GET_WORKOUTSTATE, &[]byte{1}, false},         //Workout State
	0x1A9F: Resp{CSAFE_PM_GET_WORKOUTINTERVALCOUNT, &[]byte{1}, false}, //Workout Interval Count
	0x1A8E: Resp{CSAFE_PM_GET_INTERVALTYPE, &[]byte{1}, false},         //Interval Type
	0x1ACF: Resp{CSAFE_PM_GET_RESTTIME, &[]byte{2}, false},             //Rest Time

	// Response Data to PM3 Specific Long Commands
	0x1A05: Resp{CSAFE_PM_SET_SPLITDURATION, &[]byte{0}, false},                                                 //No variables returned !! double check
	0x1A6B: Resp{CSAFE_PM_GET_FORCEPLOTDATA, &[]byte{1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}, false}, //Bytes read, data ...
	0x1A27: Resp{CSAFE_PM_SET_SCREENERRORMODE, &[]byte{0}, false},                                               //No variables returned !! double check
	0x1A6C: Resp{CSAFE_PM_GET_HEARTBEATDATA, &[]byte{1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}, false}, //Bytes read, data ...
}
