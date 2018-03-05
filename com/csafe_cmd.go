package com

import (
	"fmt"
	"log"
)

func _Int2bytes(numbytes uint, num uint) (result []uint) {
	for k := uint(0); k < numbytes; k++ {
		result = append(result, (num>>uint(8*k))&0xFF)
	}
	return result
}

func _Bytes2int(rawBytes []uint) uint {
	integer := uint(0)
	for k, v := range rawBytes {
		integer = (v << uint(8*k)) | integer
	}
	return integer

}

func _Bytes2ascii(rawBytes []byte) string {
	return string(rawBytes)
}

func sum(values []uint) (result uint) {
	for _, v := range values {
		result += v
	}
	return
}

func abs(value uint) uint {
	if value > 0 {
		return value
	}
	return -value
}

// for sending
func write(arguments ...uint) (message []uint) {
	var wrapper uint
	var wrapped []uint
	maxresponse := uint(3) //start & stop flag & status
	i := 0
	for i < len(arguments) {
		argument := arguments[i]
		cmdprop := Cmds[argument]
		//fmt.Printf("%#v, %v\n", argument, cmdprop)
		var command []uint
		// load variables if command is a Long Command
		if cmdprop.Variable != nil {
			for _, varbyte := range *cmdprop.Variable {
				i++
				intvalue := arguments[i]
				value := _Int2bytes(varbyte, intvalue)
				command = append(command, value...)
			}
			// data byte count
			cmdbytes := uint(len(command))
			command = append([]uint{cmdbytes}, command...)
		}
		//add command id
		command = append([]uint{cmdprop.ID}, command...)

		// closes wrapper if required
		if len(wrapped) > 0 && (cmdprop.Data != nil || (*cmdprop.Data)[0] != wrapper) {
			wrapped = append([]uint{uint(len(wrapped))}, wrapped...) //data byte count for wrapper
			wrapped = append([]uint{wrapper}, wrapped...)            //wrapper command id
			message = append(message, wrapped...)                    // adds wrapper to message
			wrapped = nil
			wrapper = 0
		}

		// create or extend wrapper
		if cmdprop.Data != nil { // checks if command needs a wrapper
			if wrapper == (*cmdprop.Data)[0] { // checks if currently in the same wrapper
				wrapped = append(wrapped, command...)
			} else { // creating a new wrapper
				wrapped = command
				wrapper = (*cmdprop.Data)[0]
				maxresponse += 2
			}
			command = nil //clear command to prevent it from getting into message
		}

		// max message length
		cmdid := cmdprop.ID | (wrapper << 8)
		// double return to account for stuffing
		resp := Resps[cmdid]
		maxresponse += abs(sum(*resp.Bytes))*2 + 1

		//add completed command to final message
		message = append(message, command...)
		i++
	}

	// closes wrapper if message ended on it
	if len(wrapped) > 0 {
		wrapped = append([]uint{uint(len(wrapped))}, wrapped...) //data byte count for wrapper
		wrapped = append([]uint{wrapper}, wrapped...)            //wrapper command id
		message = append(message, wrapped...)                    // adds wrapper to message
	}

	// prime variables
	checksum := uint(0x0)
	j := 0

	// checksum and byte stuffing
	for j < len(message) {
		//calculate checksum
		checksum = checksum ^ message[j]

		//byte stuffing
		if 0xF0 <= message[j] && message[j] <= 0xF3 {
			message = append(message, 0)
			copy(message[j+1:], message[j:])
			message[j] = ByteStuffingFlag
			j++
			message[j] = message[j] & 0x3
		}
		j++
	}

	// add checksum to end of message
	message = append(message, checksum)

	// start & stop frames
	message = append([]uint{StandardFrameStartFlag}, message...)
	message = append(message, StopFrameFlag)

	// check for frame size (96 bytes)
	if len(message) > 96 {
		log.Printf("Message is too long: %v", len(message))
	}

	// report IDs
	maxmessage := uint(len(message) + 1)
	if maxresponse > maxmessage {
		maxmessage = maxresponse
	}

	if maxmessage <= 21 {
		message = append([]uint{0x01}, message...)
		message = append(message, make([]uint, 21-len(message))...)
	} else if maxmessage <= 63 {
		message = append([]uint{0x04}, message...)
		message = append(message, make([]uint, 63-len(message))...)
	} else if (len(message) + 1) <= 121 {
		message = append([]uint{0x02}, message...)
		message = append(message, make([]uint, 121-len(message))...)
		if maxresponse > 121 {
			log.Printf("Response may be too long to recieve.  Max possible length %v", maxresponse)
		}
	} else {
		log.Printf("Message too long.  Message length %v", len(message))
		message = nil
	}
	fmt.Println(message)
	return
}

func CheckMessage(input []uint) (message []uint) {
	//prime variables
	message = input
	i := 0
	checksum := uint(0)

	//checksum and unstuff
	for i < len(message) {
		//byte unstuffing
		if message[i] == ByteStuffingFlag {
			stuffvalue := message[i+1]
			message = append(message[:i+1], message[i+2:]...)
			message[i] = 0xF0 | stuffvalue
		}
		fmt.Println(checksum)
		//calculate checksum
		checksum = checksum ^ message[i]

		i++
	}

	//checks checksum
	if checksum != 0 {
		log.Println("Checksum error")
		return []uint{}
	}

	//remove checksum from  end of message
	message = message[:len(message)-1]

	return
}

func makeArray(length, value uint) []uint {
	result := make([]uint, length)
	for i := range result {
		result[i] = value
	}
	return result
}

// for recieving!!
func read(transmission ...uint) (respone map[uint]*[]uint) {
	//prime variables
	var message []uint
	stopfound := false
	//reportid = transmission[0]
	startflag := transmission[1]
	var j int

	if startflag == ExtendedFrameStartFlag {
		//destination = transmission[2]
		//source = transmission[3]
		j = 4
	} else if startflag == StandardFrameStartFlag {
		j = 2
	} else {
		log.Printf("No Start Flag found")
		return
	}

	for ; j < len(transmission); j++ {
		t := transmission[j]
		if t == StopFrameFlag {
			stopfound = true
			break
		}
		message = append(message, t)
	}

	if !stopfound {
		log.Println("No Stop Flag found.")
		return
	}

	message = CheckMessage(message)
	status, message := message[0], message[1:]

	//prime variables
	var response = make(map[uint]*[]uint)
	response[CSAFE_GETSTATUS_CMD] = &[]uint{status}
	var k int
	wrapend := -1
	wrapper := uint(0x0)

	for k < len(message) {
		var result *[]uint

		//get command name
		msgcmd := message[k]
		if k <= wrapend {
			msgcmd = wrapper | msgcmd //check if still in wrapper
		}
		msgprop := Resps[msgcmd]
		k++

		//get data byte count
		bytecount := message[k]
		k++

		//if wrapper command then gets command in wrapper
		if msgprop.ID == CSAFE_SETUSERCFG1_CMD {
			wrapper = message[k-2] << 8
			wrapend = k + int(bytecount) - 1
			if bytecount != 0 { //If wrapper length != 0
				msgcmd = wrapper | message[k]
				msgprop = Resps[msgcmd]
				k++
				bytecount = message[k]
				k++
			}
		}

		//special case for capability code, response lengths differ based off capability code
		if msgprop.ID == CSAFE_GETCAPS_CMD {
			*msgprop.Bytes = makeArray(1, bytecount)
		}

		//special case for get id, response length is variable
		if msgprop.ID == CSAFE_GETID_CMD {
			msgprop.Bytes = nil
		}

		// checking that the recieved data byte is the expected length, sanity check
		if abs(sum(*msgprop.Bytes)) != 0 && bytecount != abs(sum(*msgprop.Bytes)) {
			log.Println("Warning: bytecount is an unexpected length")
		}

		//extract values
		for _, numbytes := range *msgprop.Bytes {
			rawBytes := message[uint(k) : uint(k)+abs(numbytes)]
			var value = rawBytes
			if numbytes >= 0 {
				value = []uint{_Bytes2int(rawBytes)}
			}
			*result = append(*result, value...)
			k = k + int(abs(numbytes))
		}
		response[msgprop.ID] = result
	}
	return
}
