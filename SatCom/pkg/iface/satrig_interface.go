package iface

import v1 "epyphite/space/SatCom/pkg/api/v1"

type SatRigI interface {
	GetCaps() (v1.Capabilities, error)
	GetState() (v1.State, error)
	GetFrequency() (float64, error)
	SetFrequency(freq float64) error
	GetMode() (string, int, error)
	SetMode(mode string, pbWidth int) error
	GetVfo() (string, error)
	SetVfo(string) error
	GetRit() (int, error)
	SetRit(rit int) error
	GetXit() (int, error)
	SetXit(xit int) error
	GetAntenna() (int, error)
	SetAntenna(int) error
	GetPtt() (bool, error)
	SetPtt(bool) error
	ExecVfoOps([]string) error
	GetTuningStep() (int, error)
	SetTuningStep(int) error
	GetPowerstat() (bool, error)
	SetPowerstat(bool) error
	GetSplitVfo() (string, bool, error)
	SetSplitVfo(string, bool) error
	GetSplitFrequency() (float64, error)
	SetSplitFrequency(float64) error
	GetSplitMode() (string, int, error)
	SetSplitMode(string, int) error
	GetSplitPbWidth() (int, error)
	SetSplitPbWidth(int) error
	SetSplitFrequencyMode(float64, string, int) error
	GetSplitFrequencyMode() (float64, string, int, error)
	GetFunction(string) (bool, error)
	SetFunction(string, bool) error
	GetLevel(string) (float32, error)
	SetLevel(string, float32) error
	GetParameter(string) (float32, error)
	SetParameter(string, float32) error
}
