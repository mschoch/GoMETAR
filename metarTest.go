package main
import(
"fmt"
"regexp"
"strings"
"time"
"encoding/json"
)
type Metar struct{
	Originalmetar string	`json:"-"`
	Metartype string	`json:"type,omitempty"`
	Mod string	`json:"mod,omitempty"`
	Station string	`json:"station"`
	Time time.Time	`json:"time"`
	Cycle int	`json:"cycle,omitempty"`
	WindDir int	`json:"windDir,omitempty"`
	WindSpeed int	`json:"windSpeed,omitempty"`
	WindGust int	`json:"windGust,omitempty"`
	WindDirFrom int	`json:"windDirFrom,omitempty"`
	WindDirTo int	`json:"windDirTo,omitempty"`
	Vis int	`json:"vis,omitempty"`
	VisDir int	`json:"visDir,omitempty"`
	MaxVis int	`json:"maxVis,omitempty"`
	MaxVisDir int	`json:"maxVisDir,omitempty"`
	Temp int	`json:"temp,omitempty"`
	Dewpt int	`json:"dewpt,omitempty"`
	Pressure int	`json:"press,omitempty"`
	Runway []int	`json:"runway,omitempty"`
	Weather []int	`json:"weather,omitempty"`
	Recent []int	`json:"recent,omitempty"`
	Sky []int	`json:"sky,omitempty"`
	Windshear int	`json:"windShear,omitempty"`
	WindSpeedPeak int	`json:"windSpeedPeak,omitempty"`
	WindDirPeak int	`json:"windDirPeak,omitempty"`
	PeakWindTime time.Time	`json:"-"`
	WindShiftTime time.Time	`json:"-"`
	MaxTemp6hr int	`json:"maxTemp6hr,omitempty"`
	MaxTemp24hr int	`json:"maxTemp24hr,omitempty"`
	MinTemp6hr int	`json:"minTemp6hr,omitempty"`
	MinTemp24hr int	`json:"minTemp24hr,omitempty"`
	PressureSeaLevel int	`json:"pressSeaLev,omitempty"`
	Precip1hr int	`json:"precip1hr,omitempty"`
	Precip3hr int	`json:"precip3hr,omitempty"`
	Precip6hr int	`json:"precip6hr,omitempty"`
	Precip24hr int	`json:"precip24hr,omitempty"`
	trend bool
	trendGroups []string
	remarks []string
	unparsedGroups []string
	unparsedRemarks []string
}
func main(){
	returnmetar := new(Metar)

	metar := strings.TrimSpace(originalmetar)
	match := type_re.FindString(metar)
	
	if(len(match)>1){
	returnmetar.Metartype = match
	metar = strings.TrimSpace(metar[len(match):])
	}

	match = station_re.FindString(metar)
		
	if(len(match)>1){
	returnmetar.Station = match
	metar = strings.TrimSpace(metar[len(match):])
	}
	
	match = time_re.FindString(metar)
	fmt.Print(returnmetar)
	
	
}

func (m *Metar) String()(string){
	return `Station:`+m.Station
}

func (m *Metar) JSON()(string){
	returnstring, _ := json.Marshal(m)
	return string(returnstring)
}

const(
originalmetar = `KDEN 300053Z 06008KT 10SM FEW080 SCT120 16/M01 A2997 RMK AO2 SLP115 T01561011`
)

var (
type_re =  regexp.MustCompile(`METAR|SPECI`)
station_re =  regexp.MustCompile(`([A-Z][A-Z0-9]{3})`)
time_re = regexp.MustCompile(`\d\d\d\d\d\dZ`)
modifier_re = regexp.MustCompile(`(AUTO|FINO|NIL|TEST|CORR?|RTD|CC[A-G])`)
wind_re = regexp.MustCompile(`(?P<dir>[\dO]{3}|[0O]|///|MMM|VRB)
                          (?P<speed>P?[\dO]{2,3}|[0O]+|[/M]{2,3})
                        (G(?P<gust>P?(\d{1,3}|[/M]{1,3})))?
                          (?P<units>KTS?|LT|K|T|KMH|MPS)?
                      (\s+(?P<varfrom>\d\d\d)V
                          (?P<varto>\d\d\d))?\s`)
visibility_re = regexp.MustCompile(`(?P<vis>(?P<dist>\d\d\d\d|////)
                                   (?P<dir>[NSEW][EW]? | NDV)? |
                                   (?P<distu>M?(\d+|\d\d?/\d\d?|\d+\s+\d/\d))
                                   (?P<units>SM|KM|M|U) | 
                                   CAVOK )`)
runway_re = regexp.MustCompile(`(RVRNO | 
                             R(?P<name>\d\d(RR?|LL?|C)?)/
                              (?P<low>(M|P)?\d\d\d\d)
                            (V(?P<high>(M|P)?\d\d\d\d))?
                              (?P<unit>FT)?[/NDU]*)`)
weather_re = regexp.MustCompile(`(?P<int>(-|\+|VC)*)
                             (?P<desc>(MI|PR|BC|DR|BL|SH|TS|FZ)+)?
                             (?P<prec>(DZ|RA|SN|SG|IC|PL|GR|GS|UP|/)*)
                             (?P<obsc>BR|FG|FU|VA|DU|SA|HZ|PY)?
                             (?P<other>PO|SQ|FC|SS|DS|NSW|/+)?
                             (?P<int2>[-+])?`)
sky_re = regexp.MustCompile(`(?P<cover>VV|CLR|SKC|SCK|NSC|NCD|BKN|SCT|FEW|[O0]VC|///)
                        (?P<height>[\dO]{2,4}|///)?
                        (?P<cloud>([A-Z][A-Z]+|///))?`)
temperature_re = regexp.MustCompile(`(?P<temp>(M|-)?\d+|//|XX|MM)/
                          (?P<dewpt>(M|-)?\d+|//|XX|MM)?`)
pressure_re = regexp.MustCompile(`(?P<unit>A|Q|QNH|SLP)?
                           (?P<press>[\dO]{3,4}|////)
                           (?P<unit2>INS)?`)
recent_re = regexp.MustCompile(`RE(?P<desc>MI|PR|BC|DR|BL|SH|TS|FZ)?
                              (?P<prec>(DZ|RA|SN|SG|IC|PL|GR|GS|UP)*)?
                              (?P<obsc>BR|FG|FU|VA|DU|SA|HZ|PY)?
                              (?P<other>PO|SQ|FC|SS|DS)?`)
windshear_re = regexp.MustCompile(`(WS\s+)?(ALL\s+RWY|RWY(?P<name>\d\d(RR?|L?|C)?))`)
color_re = regexp.MustCompile(`(BLACK)?(BLU|GRN|WHT|RED)\+?
                        (/?(BLACK)?(BLU|GRN|WHT|RED)\+?)*`)
runwaystate_re = regexp.MustCompile(`((?P<name>\d\d) |
                                 R(?P<namenew>\d\d)(RR?|LL?|C)?/?)
                                (?P<deposit>(\d|/))
                                (?P<extent>(\d|/))
                                (?P<depth>(\d\d|//))
                                (?P<friction>(\d\d|//))`)
trend_re = regexp.MustCompile(`(?P<trend>TEMPO|BECMG|FCST|NOSIG)`)
trendtime_re = regexp.MustCompile(`(?P<when>(FM|TL|AT))(?P<hour>\d\d)(?P<min>\d\d)`)
remark_re = regexp.MustCompile(`(RMKS?|NOSPECI|NOSIG)`)
auto_re = regexp.MustCompile(`AO(?P<type>\d)`)
sealvl_pressure_re = regexp.MustCompile(`SLP(?P<press>\d\d\d)`)
peak_wind_re = regexp.MustCompile(`P[A-Z]\s+WND\s+
                               (?P<dir>\d\d\d)
                               (?P<speed>P?\d\d\d?)/
                               (?P<hour>\d\d)?
                               (?P<min>\d\d)`)
wind_shift_re = regexp.MustCompile(`WSHFT\s+
                                (?P<hour>\d\d)?
                                (?P<min>\d\d)
                                (\s+(?P<front>FROPA))?`)
precip_1hr_re = regexp.MustCompile(`P(?P<precip>\d\d\d\d)`)
precip_24hr_re = regexp.MustCompile(`(?P<type>6|7)
                                 (?P<precip>\d\d\d\d)`)
press_3hr_re = regexp.MustCompile(`5(?P<tend>[0-8])
                                (?P<press>\d\d\d)`)
temp_1hr_re = regexp.MustCompile(`T(?P<tsign>0|1)
                               (?P<temp>\d\d\d)
                               ((?P<dsign>0|1)
                               (?P<dewpt>\d\d\d))?`)
temp_6hr_re = regexp.MustCompile(`(?P<type>1|2)
                              (?P<sign>0|1)
                              (?P<temp>\d\d\d)`)
temp_24hr_re = regexp.MustCompile(`4(?P<smaxt>0|1)
                                (?P<maxt>\d\d\d)
                                (?P<smint>0|1)
                                (?P<mint>\d\d\d)`)
unparsed_re = regexp.MustCompile(`(?P<group>\S+)`)
lightning_re = regexp.MustCompile(`((?P<freq>OCNL|FRQ|CONS)\s+)?
                             LTG(?P<type>(IC|CC|CG|CA)*)
                                ( \s+(?P<loc>( OHD | VC | DSNT\s+ | \s+AND\s+ | 
                                 [NSEW][EW]? (-[NSEW][EW]?)* )+) )?`)                                       
ts_loc_re = regexp.MustCompile(`TS(\s+(?P<loc>( OHD | VC | DSNT\s+ | \s+AND\s+ | 
                                           [NSEW][EW]? (-[NSEW][EW]?)* )+))?
                                          ( \s+MOV\s+(?P<dir>[NSEW][EW]?) )?`)
loc_terms = map[string]string{"OHD": "overhead",
			"DSNT": "distant",
			"AND": "and",
			"VC": "nearby"}
sky_cover = map[string]string{"SKC":"clear",
			"CLR":"clear",
			"NSC":"clear",
			"NCD":"clear",
			"FEW":"a few ",
			"SCT":"scattered ",
			"BKN":"broken ",
			"OVC":"overcast",
			"///":"",
			"VV":"indefinite ceiling" }
cloud_type = map[string]string{ "TCU":"towering cumulus",
               "CU":"cumulus",
               "CB":"cumulonimbus",
               "SC":"stratocumulus",
               "CBMAM":"cumulonimbus mammatus",
               "ACC":"altocumulus castellanus",
               "SCSL":"standing lenticular stratocumulus",
               "CCSL":"standing lenticular cirrocumulus",
               "ACSL":"standing lenticular altocumulus" }
weather_int = map[string]string{ "-":"light", 
                "+":"heavy", 
                "-VC":"nearby light", 
                "+VC":"nearby heavy", 
                "VC":"nearby" }
weather_desc = map[string]string{ "MI":"shallow",
                 "PR":"partial",
                 "BC":"patches of", 
                 "DR":"low drifting", 
                 "BL":"blowing",
                 "SH":"showers",
                 "TS":"thunderstorm",
                 "FZ":"freezing" }
weather_prec = map[string]string{ "DZ":"drizzle",
                 "RA":"rain",
                 "SN":"snow",
                 "SG":"snow grains",
                 "IC":"ice crystals",
                 "PL":"ice pellets",
                 "GR":"hail",
                 "GS":"snow pellets",
                 "UP":"unknown precipitation",
                 "//":"" }
weather_obsc = map[string]string{ "BR":"mist",
                 "FG":"fog",
                 "FU":"smoke",
                 "VA":"volcanic ash",
                 "DU":"dust",
                 "SA":"sand",
                 "HZ":"haze",
                 "PY":"spray" }
weather_other = map[string]string{ "PO":"sand whirls",
                  "SQ":"squalls",
                  "FC":"funnel cloud",
                  "SS":"sandstorm",
                  "DS":"dust storm" }
weather_special = map[string]string{ "+FC":"tornado" }
color = map[string]string{ "BLU":"blue",
          "GRN":"green",
          "WHT":"white" }
pressure_tendency = map[string]string{ "0":"increasing, then decreasing",
                      "1":"increasing more slowly",
                      "2":"increasing",        
                      "3":"increasing more quickly",
                      "4":"steady",
                      "5":"decreasing, then increasing",
                      "6":"decreasing more slowly",
                      "7":"decreasing",
                      "8":"decreasing more quickly" }
lightning_frequency = map[string]string{ "OCNL":"occasional",
                        "FRQ":"frequent",
                        "CONS":"constant" }
lightning_type = map[string]string{ "IC":"intracloud",
                   "CC":"cloud-to-cloud",
                   "CG":"cloud-to-ground",
                   "CA":"cloud-to-air" }
report_type = map[string]string{ "METAR":"routine report",
                "SPECI":"special report",
                "AUTO":"automatic report",
                "COR":"manually corrected report" }			   
)