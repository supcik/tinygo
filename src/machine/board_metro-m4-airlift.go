// +build sam,atsamd51,metro_m4_airlift

package machine

import "device/sam"

// used to reset into bootloader
const RESET_MAGIC_VALUE = 0xf01669ef

// GPIO Pins
const (
	D0 = PA23 // UART0 RX/PWM available
	D1 = PA22 // UART0 TX/PWM available
	D2 = PB17 // PWM available
	D3 = PB16 // PWM available
	D4 = PB13 // PWM available
	D5 = PB14 // PWM available
	D6 = PB15 // PWM available
	D7 = PB12 // PWM available

	D8  = PA21 // PWM available
	D9  = PA20 // PWM available
	D10 = PA18 // can be used for PWM or UART1 TX
	D11 = PA19 // can be used for PWM or UART1 RX
	D12 = PA17 // PWM available
	D13 = PA16 // PWM available

	D40 = PB22 // built-in neopixel
)

// Analog pins
const (
	A0 = PA02 // ADC/AIN[0]
	A1 = PA05 // ADC/AIN[2]
	A2 = PB06 // ADC/AIN[3]
	A3 = PB00 // ADC/AIN[4] // NOTE: different between "airlift" and non-airlift versions
	A4 = PB08 // ADC/AIN[5]
	A5 = PB09 // ADC/AIN[10]
)

const (
	LED = D13
)

// UART0 aka USBCDC pins
const (
	USBCDC_DM_PIN = PA24
	USBCDC_DP_PIN = PA25
)

const (
	UART_TX_PIN = D1
	UART_RX_PIN = D0
)

// Note: UART1 is on SERCOM3, defined in machine_atsamd51.go

const (
	NINA_CS     = PA15
	NINA_ACK    = PB04
	NINA_GPIO0  = PB01
	NINA_RESETN = PB05

	NINA_TX  = PA04
	NINA_RX  = PA07
	NINA_RTS = PB23
)

// UART2 on the Metro M4 Airlift Lite connects to the onboard ESP32-WROOM chip.
var (
	UART2 = UART{
		Buffer: NewRingBuffer(),
		Bus:    sam.SERCOM0_USART_INT,
		Mode:   PinSERCOMAlt,
	}
)

//go:export SERCOM0_IRQHandler
func handleUART2() {
	UART2.Receive(byte((UART2.Bus.DATA.Get() & 0xFF)))
	UART2.Bus.INTFLAG.SetBits(sam.SERCOM_USART_INT_INTFLAG_RXC)
}

// I2C pins
const (
	SDA_PIN = PB02 // SDA: SERCOM5/PAD[0]
	SCL_PIN = PB03 // SCL: SERCOM5/PAD[1]
)

// I2C on the Metro M4.
var (
	I2C0 = I2C{Bus: sam.SERCOM5_I2CM,
		SDA:     SDA_PIN,
		SCL:     SCL_PIN,
		PinMode: PinSERCOMAlt}
)

// SPI pins
const (
	SPI0_SCK_PIN  = PA13 // SCK:  SERCOM2/PAD[1]
	SPI0_MOSI_PIN = PA12 // MOSI: SERCOM2/PAD[0]
	SPI0_MISO_PIN = PA14 // MISO: SERCOM2/PAD[2]
)

// SPI on the Metro M4.
var (
	SPI0 = SPI{
		Bus:         sam.SERCOM2_SPIM,
		SCK:         SPI0_SCK_PIN,
		MOSI:        SPI0_MOSI_PIN,
		MISO:        SPI0_MISO_PIN,
		DOpad:       spiTXPad0SCK1,
		DIpad:       sercomRXPad2,
		MISOPinMode: PinSERCOM,
		MOSIPinMode: PinSERCOM,
		SCKPinMode:  PinSERCOM,
	}
)

const (
	SPI1_SCK_PIN  = D12 // MISO: SERCOM1/PAD[1]
	SPI1_MOSI_PIN = D11 // MOSI: SERCOM1/PAD[3]
	SPI1_MISO_PIN = D13 // SCK:  SERCOM1/PAD[0]
)

// SPI1 on the Metro M4 on pins 11,12,13
var (
	SPI1 = SPI{
		Bus:         sam.SERCOM1_SPIM,
		SCK:         SPI1_SCK_PIN,
		MOSI:        SPI1_MOSI_PIN,
		MISO:        SPI1_MISO_PIN,
		DOpad:       spiTXPad3SCK1,
		DIpad:       sercomRXPad0,
		MISOPinMode: PinSERCOM,
		MOSIPinMode: PinSERCOM,
		SCKPinMode:  PinSERCOM,
	}
)