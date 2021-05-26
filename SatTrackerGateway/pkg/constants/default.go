package constants

var (
	BASEURI              = "https://www.space-track.org"
	LOGINURI             = "/ajaxauth/login"
	REQUESACTIONURI      = "/basicspacedata/query"
	RequestFindStarlinks = "/class/tle_latest/NORAD_CAT_ID/>40000/ORDINAL/1/OBJECT_NAME/STARLINK~~/format/json/orderby/NORAD_CAT_ID%20asc"
	RequestOMMStarlink1  = "/class/omm/NORAD_CAT_ID/"
	BoxScore             = "/basicspacedata/query/class/boxscore/format/json"
	CurrentLEO           = "/basicspacedata/query/class/satcat/PERIOD/<128/DECAY/null-val/CURRENT/Y/"
	GPHistory            = "/basicspacedata/query/class/gp_history/format/json"
)
