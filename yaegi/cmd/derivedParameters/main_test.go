package main

import (
	"github.com/containous/yaegi/interp"
	"github.com/containous/yaegi/stdlib"
	"testing"
)

func BenchmarkCalcWindDir(b *testing.B) {
	ugrdData, err := getData("./data/ugrd.bin")
	if err != nil {
		b.Fatal("Error opening grid data:", err)
	}
	vgrdData, err := getData("./data/vgrd.bin")
	if err != nil {
		b.Fatal("Error opening grid data:", err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		calcWindDir(ugrdData[1:len(ugrdData)-1], vgrdData[1:len(vgrdData)-1])
	}
}

func BenchmarkCalcWindSpd(b *testing.B) {
	ugrdData, err := getData("./data/ugrd.bin")
	if err != nil {
		b.Fatal("Error opening grid data:", err)
	}
	vgrdData, err := getData("./data/vgrd.bin")
	if err != nil {
		b.Fatal("Error opening grid data:", err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		calcWindSpd(ugrdData[1:len(ugrdData)-1], vgrdData[1:len(vgrdData)-1])
	}
}

func BenchmarkCalcWindDirInterp(b *testing.B) {
	ugrdData, err := getData("./data/ugrd.bin")
	if err != nil {
		b.Fatal("Error opening grid data:", err)
	}
	vgrdData, err := getData("./data/vgrd.bin")
	if err != nil {
		b.Fatal("Error opening grid data:", err)
	}
	i := interp.New(interp.Options{})
	i.Use(stdlib.Symbols)
	_, err = i.Eval(windScript)
	if err != nil {
		b.Fatal("Error evaluating the script:", err)
	}
	v, err := i.Eval("foo.CalcWindDir")
	if err != nil {
		b.Fatal("Error evaluating the script:", err)
	}
	myfunc := v.Interface().(func([]float32, []float32) []float64)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		myfunc(ugrdData[1:len(ugrdData)-1], vgrdData[1:len(vgrdData)-1])
	}
}

func BenchmarkCalcWindSpdInterp(b *testing.B) {
	ugrdData, err := getData("./data/ugrd.bin")
	if err != nil {
		b.Fatal("Error opening grid data:", err)
	}
	vgrdData, err := getData("./data/vgrd.bin")
	if err != nil {
		b.Fatal("Error opening grid data:", err)
	}
	i := interp.New(interp.Options{})
	i.Use(stdlib.Symbols)
	_, err = i.Eval(windScript)
	if err != nil {
		b.Fatal("Error evaluating the script:", err)
	}
	v, err := i.Eval("foo.CalcWindSpd")
	if err != nil {
		b.Fatal("Error evaluating the script:", err)
	}
	myfunc := v.Interface().(func([]float32, []float32) []float64)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		myfunc(ugrdData[1:len(ugrdData)-1], vgrdData[1:len(vgrdData)-1])
	}
}
