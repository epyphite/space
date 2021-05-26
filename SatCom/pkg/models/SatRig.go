package models

//Communicate and interface with Rigs HamLib

import (
	"fmt"

	"github.com/dh1tw/goHamlib"
)

//SatComRig basic satcom rig struct to access hamlib
type SatComRig struct {
	Rig *goHamlib.Rig
}

//NewSatComRig Creates a new object of type SatComRig
func NewSatComRig(model string, path string) (SatComRig, error) {
	var satcom SatComRig
	var rig goHamlib.Rig

	var err error
	rig = goHamlib.Rig{}
	rig.Init(-1)
	if model == "" {
		fmt.Println(goHamlib.ListModels())
	}
	return satcom, err
}

//Open HamLib device
func (st *SatComRig) Open() error {
	var err error
	err = st.Rig.Open()
	return err
}

//Close HamLib Device
func (st *SatComRig) Close() error {
	var err error
	err = st.Close()
	return err
}
