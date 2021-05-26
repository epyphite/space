package satcom


import (
	"errors"
	v1 "epyphite/space/SatCom/pkg/api/v1"
)

func (r *SatCom) GetCaps () (v1.Capabilities, error){
	return r.caps.
} 