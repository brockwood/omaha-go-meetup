package main

import (
	"encoding/binary"
	"fmt"
	"math"
	"os"

	"github.com/containous/yaegi/interp"
	"github.com/containous/yaegi/stdlib"
)

var (
	myCounter  = 0
	windScript = `package foo

import "math"

const x = math.Pi/180

func Rad(d float64) float64 {
	return d*x
}
func Deg(r float64) float64 {
	return r/x
}

func CalcWindDir(windu, windv []float32) []float64 {
	ws := make([]float64, len(windu))
	var windu64 float64
	var windv64 float64
	for component := range windu {
		windu64 = float64(windu[component])
		windv64 = float64(windv[component])
		ws[component] = 270.0 - Deg(math.Atan2(windv64, windu64))
	}
	return ws
}

func CalcWindSpd(windu, windv []float32) []float64 {
	wd := make([]float64, len(windu))
	for component := range windu {
		wu := float64(windu[component])
		wv := float64(windv[component])
		wd[component] = math.Sqrt(wu * wu + wv * wv)
	}
	return wd
}`
)

const x = math.Pi / 180

func Rad(d float64) float64 {
	return d * x
}
func Deg(r float64) float64 {
	return r / x
}

func main() {
	ugrdData, err := getData("./data/ugrd.bin")
	if err != nil {
		fmt.Println("Error opening grid data:", err)
		return
	}
	vgrdData, err := getData("./data/vgrd.bin")
	if err != nil {
		fmt.Println("Error opening grid data:", err)
		return
	}
	wdir := calcWindDir(ugrdData[1:len(ugrdData)-1], vgrdData[1:len(vgrdData)-1])
	wspd := calcWindSpd(ugrdData[1:len(ugrdData)-1], vgrdData[1:len(vgrdData)-1])
	i := interp.New(interp.Options{})
	i.Use(stdlib.Symbols)
	_, err = i.Eval(windScript)
	if err != nil {
		panic(err)
	}
	windDir, err := i.Eval("foo.CalcWindDir")
	if err != nil {
		panic(err)
	}
	myWindDirScript := windDir.Interface().(func([]float32, []float32) []float64)
	wdirS := myWindDirScript(ugrdData[1:len(ugrdData)-1], vgrdData[1:len(vgrdData)-1])
	windSpd, err := i.Eval("foo.CalcWindSpd")
	if err != nil {
		panic(err)
	}
	myWindSpdScript := windSpd.Interface().(func([]float32, []float32) []float64)
	wspdS := myWindSpdScript(ugrdData[1:len(ugrdData)-1], vgrdData[1:len(vgrdData)-1])
	fmt.Println(wdir[0], wdirS[0], wdir[1], wdirS[1])
	fmt.Println(wspd[0], wspdS[0], wspd[1], wspdS[1])
}

func calcWindDir(windu, windv []float32) []float64 {
	ws := make([]float64, len(windu))
	var windu64 float64
	var windv64 float64
	for component := range windu {
		windu64 = float64(windu[component])
		windv64 = float64(windv[component])
		ws[component] = 270.0 - Deg(math.Atan2(windv64, windu64))
	}
	return ws
}

func calcWindSpd(windu, windv []float32) []float64 {
	wd := make([]float64, len(windu))
	for component := range windu {
		wu := float64(windu[component])
		wv := float64(windv[component])
		wd[component] = math.Sqrt(wu*wu + wv*wv)
	}
	return wd
}

func getData(file string) ([]float32, error) {
	fileHandle, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	fi, err := fileHandle.Stat()
	if err != nil {
		return nil, err
	}
	data := make([]float32, fi.Size()/4)
	err = binary.Read(fileHandle, binary.LittleEndian, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
