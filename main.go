package main

import (
	"fmt"
	"log"

	"github.com/karalabe/hid"
	"github.com/smaragden/gorow/com"
)

const C2_VENDOR_ID = 0x17a4
const MIN_FRAME_GAP = .050 //in seconds
const INTERFACE = 0

func main() {
	/*
		// Initialize a new Context.
		ctx := gousb.NewContext()
		defer ctx.Close()


		// Iterate through available Devices, finding all that match a known VID/PID.
		vid := gousb.ID(C2_VENDOR_ID)
		devs, err := ctx.OpenDevices(func(desc *gousb.DeviceDesc) bool {
			// this function is called for every device present.
			// Returning true means the device should be opened.
			return desc.Vendor == vid
		})

		// All returned devices are now open and will need to be closed.
		for _, d := range devs {
			defer d.Close()
		}

		if err != nil {
			log.Fatalf("OpenDevices(): %v", err)
		}

		dev := devs[0]

		// Claim the default interface using a convenience function.
		// The default interface is always #0 alt #0 in the currently active
		// config.
		intf, done, err := dev.DefaultInterface()
		if err != nil {
			log.Fatalf("%s.DefaultInterface(): %v", dev, err)
		}
		defer done()

		// Open an OUT endpoint.
		iep, err := intf.OutEndpoint(7)
		if err != nil {
			log.Fatalf("%s.OutEndpoint(7): %v", intf, err)
		}
	*/
	for _, deviceInfo := range hid.Enumerate(C2_VENDOR_ID, 0) {
		fmt.Printf("Path: %s\n", deviceInfo.Path)                 // Platform-specific deviceInfo path
		fmt.Printf("VendorID: %x\n", deviceInfo.VendorID)         // deviceInfo Vendor ID
		fmt.Printf("ProductID: %x\n", deviceInfo.ProductID)       // deviceInfo Product ID
		fmt.Printf("Release: %d\n", deviceInfo.Release)           // deviceInfo Release Number in binary-coded decimal, also known as deviceInfo Version Number
		fmt.Printf("Serial: %s\n", deviceInfo.Serial)             // Serial Number
		fmt.Printf("Manufacturer: %s\n", deviceInfo.Manufacturer) // Manufacturer String
		fmt.Printf("Product: %s\n", deviceInfo.Product)           // Product string
		fmt.Printf("UsagePage: %d\n", deviceInfo.UsagePage)       // Usage Page for this deviceInfo/Interface (Windows/Mac only)
		fmt.Printf("Usage: %d\n\n", deviceInfo.Usage)             // Usage for this Device/Interface (Windows/Mac only)

		device, err := deviceInfo.Open()
		if err != nil {
			log.Fatalf("failed to open device: %v\n", err)
		}
		defer device.Close()

		command := []byte{com.CSAFE_GETVERSION_CMD, com.CSAFE_GETSERIAL_CMD, com.CSAFE_GETCAPS_CMD, 0x00}
		results := send(device, command)
		fmt.Printf("SUCCESS: %#v\n", results)
		break
	}

}

func send(device *hid.Device, message []byte) []byte {
	/*
	   Converts and sends message to erg; receives, converts, and returns ergs response
	*/

	//Checks that enough time has passed since the last message was sent,
	//if not program sleeps till time has passed
	//now = datetime.datetime.now()
	//delta = now - self.__lastsend
	//deltaraw = delta.seconds + delta.microseconds/1000000.
	//if deltaraw < MIN_FRAME_GAP:
	//    time.sleep(MIN_FRAME_GAP - deltaraw)

	//convert message to byte array
	csafe := com.Write(message...)
	//sends message to erg and records length of message
	length, err := device.Write(csafe)
	if err != nil {
		log.Fatalf("Failed to write message: %v\n", err)
	}
	fmt.Println("Read: ", length, " of data")
	//records time when message was sent
	//self.__lastsend = datetime.datetime.now()

	var response []byte
	length, err = device.Read(response)
	if err != nil {
		log.Fatalf("Failed to read response: %v\n", err)
	}

	return response
}
