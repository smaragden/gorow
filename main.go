package main

import (
	"log"

	"github.com/google/gousb"
)

const C2_VENDOR_ID = 0x17a4
const MIN_FRAME_GAP = .050 //in seconds
const INTERFACE = 0

func main() {
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

}
