package main

import (
	"os"
	"fmt"
	"periph.io/x/periph/host"
	"periph.io/x/periph/conn/physic"
	"periph.io/x/periph/conn/spi"
	"periph.io/x/periph/conn/spi/spireg"
)

func readValue() (uint16, error) {
	var err error
	var val uint16
	spiMode := spi.Mode(0)
	spidev, err := spireg.Open("SPI1.0")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to open spi: %s.\n", err)
		return val, err
	}
	defer spidev.Close()
	con, err := spidev.Connect(physic.MegaHertz, spiMode, 8)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect devices: %s.\n", err)
		return val, err
	}
	inBuf := make([]byte, 2)
	outBuf := make([]byte, 2)
	if err = con.Tx(outBuf, inBuf); err == nil {
		val = (uint16(inBuf[0] & 0x1F) << 7) | uint16(inBuf[1] >> 1)
	}
	return val, err
}

func main() {
	host.Init()
	val, err := readValue()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s.\n", err)
		os.Exit(1)
	}
	fmt.Printf("Current value: %d\n", val)
	// R2 = (ADC*R1)/(4096-ADC)
}
