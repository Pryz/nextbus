package nextbus

import (
    "encoding/xml"
    "io/ioutil"
    "net/http"
)

type AgencyResponse struct {
    XMLName    xml.Name `xml:"body"`
    AgencyList []Agency `xml:"agency"`
}

type Agency struct {
    XMLName     xml.Name `xml:"agency"`
    Tag         string   `xml:"tag,attr"`
    Title       string   `xml:"title,attr"`
    RegionTitle string   `xml:"regionTitle,attr"`
}

func GetAgencyList() ([]Agency, error) {
    resp, httpErr := http.Get("http://webservices.nextbus.com/service/publicXMLFeed?command=agencyList")
    if httpErr != nil {
        return nil, httpErr
    }
    defer resp.Body.Close()

    body, readErr := ioutil.ReadAll(resp.Body)
    if readErr != nil {
        return nil, readErr
    }

    var a AgencyResponse
    xmlErr := xml.Unmarshal(body, &a)
    if xmlErr != nil {
        return nil, xmlErr
    }
    return a.AgencyList, nil
}

type RouteResponse struct {
    XMLName   xml.Name `xml:"body"`
    RouteList []Route  `xml:"route"`
}

type Route struct {
    XMLName xml.Name `xml:"route"`
    Tag     string   `xml:"tag,attr"`
    Title   string   `xml:"title,attr"`
}

func GetRouteList(agencyTag string) ([]Route, error) {
    resp, httpErr := http.Get("http://webservices.nextbus.com/service/publicXMLFeed?command=routeList&a=" + agencyTag)
    if httpErr != nil {
        return nil, httpErr
    }
    defer resp.Body.Close()

    body, readErr := ioutil.ReadAll(resp.Body)
    if readErr != nil {
        return nil, readErr
    }

    var a RouteResponse
    xmlErr := xml.Unmarshal(body, &a)
    if xmlErr != nil {
        return nil, xmlErr
    }
    return a.RouteList, nil
}

type RouteConfigResponse struct {
    XMLName   xml.Name      `xml:"body"`
    RouteList []RouteConfig `xml:"route"`
}

type RouteConfig struct {
    XMLName       xml.Name  `xml:"route"`
    StopList      []Stop    `xml:"stop"`
    Tag           string    `xml:"tag,attr"`
    Title         string    `xml:"title,attr"`
    Color         string    `xml:"color,attr"`
    OppositeColor string    `xml:"oppositeColor,attr"`
    LatMin        string    `xml:"latMin,attr"`
    LatMax        string    `xml:"latMax,attr"`
    LonMin        string    `xml:"lonMin,attr"`
    LonMax        string    `xml:"lonMax,attr"`
    Dir           Direction `xml:"direction"`
    PathList      []Path    `xml:"path"`
}

type Stop struct {
    XMLName xml.Name `xml:"stop"`
    Tag     string   `xml:"tag,attr"`
    Title   string   `xml:"title,attr"`
    Lat     string   `xml:"lat,attr"`
    Lon     string   `xml:"lon,attr"`
    StopId  string   `xml:"stopId,attr"`
}

type Direction struct {
    XMLName        xml.Name     `xml:"direction"`
    Tag            string       `xml:"tag,attr"`
    Title          string       `xml:"title,attr"`
    Name           string       `xml:"name,attr"`
    UseForUI       string       `xml:"useForUI,attr"`
    StopMarkerList []StopMarker `xml:"stop"`
}

type StopMarker struct {
    XMLName xml.Name `xml:"stop"`
    Tag     string   `xml:"tag,attr"`
}

type Path struct {
    XMLName   xml.Name `xml:"path"`
    PointList []Point  `xml:"point"`
}

type Point struct {
    XMLName xml.Name `xml:"point"`
    Lat     string   `xml:"lat,attr"`
    Lon     string   `xml:"lon,attr"`
}

func GetRouteConfig(agencyTag string, routeTag string) ([]RouteConfig, error) {
    resp, httpErr := http.Get("http://webservices.nextbus.com/service/publicXMLFeed?command=routeConfig&a=" + agencyTag + "&r=" + routeTag)
    if httpErr != nil {
        return nil, httpErr
    }
    defer resp.Body.Close()

    body, readErr := ioutil.ReadAll(resp.Body)
    if readErr != nil {
        return nil, readErr
    }

    var a RouteConfigResponse
    xmlErr := xml.Unmarshal(body, &a)
    if xmlErr != nil {
        return nil, xmlErr
    }
    return a.RouteList, nil
}

type PredictionResponse struct {
    XMLName            xml.Name         `xml:"body"`
    PredictionDataList []PredictionData `xml:"predictions"`
}

type PredictionData struct {
    XMLName                 xml.Name              `xml:"predictions"`
    PredictionDirectionList []PredictionDirection `xml:"direction"`
    MessageList             []Message             `xml:"message"`
    AgencyTitle             string                `xml:"agencyTitle,attr"`
    RouteTitle              string                `xml:"routeTitle,attr"`
    RouteTag                string                `xml:"routeTag,attr"`
    StopTitle               string                `xml:"stopTitle,attr"`
    StopTag                 string                `xml:"stopTag,attr"`
}

type PredictionDirection struct {
    XMLName        xml.Name     `xml:"direction"`
    PredictionList []Prediction `xml:"prediction"`
    Title          string       `xml:"title,attr"`
}

type Prediction struct {
    XMLName           xml.Name `xml:"prediction"`
    EpochTime         string   `xml:"epochTime,attr"`
    Seconds           string   `xml:"seconds,attr"`
    Minutes           string   `xml:"minutes,attr"`
    IsDeparture       string   `xml:"isDeparture,attr"`
    DirTag            string   `xml:"dirTag,attr"`
    Vehicle           string   `xml:"vehicle,attr"`
    VehiclesInConsist string   `xml:"vehiclesInConsist,attr"`
    Block             string   `xml:"block,attr"`
    TripTag           string   `xml:"tripTag,attr"`
}

type Message struct {
    XMLName  xml.Name `xml:"message"`
    Text     string   `xml:"text,attr"`
    Priority string   `xml:"priority,attr"`
}

func GetPredictions(agencyTag string, routeTag string, stopTag string) ([]PredictionData, error) {
    resp, httpErr := http.Get("http://webservices.nextbus.com/service/publicXMLFeed?command=predictions&a=" + agencyTag + "&r=" + routeTag + "&s=" + stopTag)
    if httpErr != nil {
        return nil, httpErr
    }
    defer resp.Body.Close()

    body, readErr := ioutil.ReadAll(resp.Body)
    if readErr != nil {
        return nil, readErr
    }

    var a PredictionResponse
    xmlErr := xml.Unmarshal(body, &a)
    if xmlErr != nil {
        return nil, xmlErr
    }
    return a.PredictionDataList, nil
}


type ScheduleResponse struct {
    XMLName            xml.Name       `xml:"body"`
    ScheduleRouteList []ScheduleRoute `xml:"route"`
}

type ScheduleRoute struct {
    XMLName       xml.Name				`xml:"route"`
		Header				ScheduleHeader	`xml:"header"`
    Tr						[]ScheduleTr		`xml:"tr"`
    RouteTitle    string					`xml:"title,attr"`
		RouteTag      string					`xml:"tag,attr"`
    ScheduleClass string					`xml:"scheduleClass,attr"`
    ServiceClass  string					`xml:"serviceClass,attr"`
    Direction     string					`xml:"direction,attr"`
}

type ScheduleHeader struct {
	XMLName	 xml.Name				`xml:"header"`
	StopList []ScheduleHeaderStop `xml:"stop"`
}

type ScheduleHeaderStop struct {
	XMLName xml.Name	`xml:"stop"`
	Tag     string		`xml:"tag,attr"`
	Title		string		`xml:",chardata"`
}

type ScheduleTr struct {
	XMLName		xml.Name					`xml:"tr"`
	StopList	[]ScheduleTrStop	`xml:"stop"`
	BlockID	  int								`xml:"blockID,attr"`
}

type ScheduleTrStop struct {
	XMLName		xml.Name	`xml:"stop"`
	Tag   	  string		`xml:"tag,attr"`
	EpochTime	int				`xml:"epochTime,attr"`
	Time			string		`xml:",chardata"`
}

func GetSchedule(agencyTag string, routeTag string) ([]ScheduleRoute, error) {
    resp, httpErr := http.Get("http://webservices.nextbus.com/service/publicXMLFeed?command=schedule&a=" + agencyTag + "&r=" + routeTag)
    if httpErr != nil {
        return nil, httpErr
    }
    defer resp.Body.Close()

    body, readErr := ioutil.ReadAll(resp.Body)
    if readErr != nil {
        return nil, readErr
    }

    var a ScheduleResponse
    xmlErr := xml.Unmarshal(body, &a)
    if xmlErr != nil {
        return nil, xmlErr
    }
    return a.ScheduleRouteList, nil
}
