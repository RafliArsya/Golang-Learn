package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	kubus              Kubus   = Kubus{4}
	volumeSeharusnya   float64 = 64
	luasSeharusnya     float64 = 96
	kelilingSeharusnya float64 = 48
)

// go test
func TestHitung(t *testing.T) {
	assert.Equal(
		t,
		kubus.Volume(),      //unit to test
		volumeSeharusnya,    //wanted value
		"Perhitungan salah", //error value
	)
}
