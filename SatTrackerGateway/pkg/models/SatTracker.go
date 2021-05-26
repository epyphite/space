package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SatTrackBasic struct {
	ID                     primitive.ObjectID `bson:"_id,omitempty"`
	Country                string             `json:"COUNTRY`
	SpadocCd               string             `json:"SPADOC_CD"`
	OrbitalTba             string             `json:"ORBITAL_TBA"`
	OrbitalPayloadCount    string             `json:"ORBITAL_PAYLOAD_COUNT"`
	OrbitalRocketBodyCount string             `json:"ORBITAL_ROCKET_BODY_COUNT"`
	OrbitalDebrisCount     string             `json:"ORBITAL_DEBRIS_COUNT"`
	OrbitalTotalCount      string             `json:"ORBITAL_TOTAL_COUNT"`
	DecayedPayloadCount    string             `json:"DECAYED_PAYLOAD_COUNT"`
	DecayedRocketBodyCount string             `json:"DECAYED_ROCKET_BODY_COUNT"`
	DecayedDebrisCount     string             `json:"DECAYED_DEBRIS_COUNT"`
	DecayedTotalCount      string             `json:"DECAYED_TOTAL_COUNT"`
	CountryTotal           string             `json:"COUNTRY_TOTAL"`
}

type SatTrackStandard struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Intldes      string             `json:"INTLDES"`
	Noradcatid   string             `json:"NORAD_CAT_ID"`
	Objecttype   string             `json:"OBJECT_TYPE"`
	Satname      string             `json:"SATNAME"`
	Country      string             `json:"COUNTRY"`
	Launch       string             `json:"LAUNCH"`
	Site         string             `json:"SITE"`
	Decay        interface{}        `json:"DECAY"`
	Period       string             `json:"PERIOD"`
	Inclination  string             `json:"INCLINATION"`
	Apogee       string             `json:"APOGEE"`
	Perigee      string             `json:"PERIGEE"`
	Comment      interface{}        `json:"COMMENT"`
	Commentcode  string             `json:"COMMENTCODE"`
	Rcsvalue     string             `json:"RCSVALUE"`
	Rcssize      interface{}        `json:"RCS_SIZE"`
	File         string             `json:"FILE"`
	Launchyear   string             `json:"LAUNCH_YEAR"`
	Launchnum    string             `json:"LAUNCH_NUM"`
	Launchpiece  string             `json:"LAUNCH_PIECE"`
	Current      string             `json:"CURRENT"`
	Objectname   string             `json:"OBJECT_NAME"`
	Objectid     string             `json:"OBJECT_ID"`
	Objectnumber string             `json:"OBJECT_NUMBER"`
}
