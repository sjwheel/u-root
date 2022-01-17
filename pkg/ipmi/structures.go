// Copyright 2022 the u-root Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ipmi

import (
	"os"
	"unsafe"
)

// IPMI represents access to the IPMI interface.
type IPMI struct {
	*os.File
}

// Command is the command code for a given message.
type Command byte

// NetFn is the network function of the class of message being sent.
type NetFn byte

// Msg is the full IPMI message to be sent.
type Msg struct {
	Netfn   NetFn
	Cmd     Command
	DataLen uint16
	Data    unsafe.Pointer
}

type req struct {
	addr    *systemInterfaceAddr
	addrLen uint32
	msgid   int64 //nolint:structcheck
	msg     Msg
}

type recv struct {
	recvType int32 //nolint:structcheck
	addr     *systemInterfaceAddr
	addrLen  uint32
	msgid    int64 //nolint:structcheck
	msg      Msg
}

type systemInterfaceAddr struct {
	addrType int32
	channel  int16
	lun      byte //nolint:unused
}

// StandardEvent is a standard systemevent.
//
// The data in this event should follow IPMI spec
type StandardEvent struct {
	Timestamp    uint32
	GenID        uint16
	EvMRev       uint8
	SensorType   uint8
	SensorNum    uint8
	EventTypeDir uint8
	EventData    [3]uint8
}

// OEMTsEvent is a timestamped OEM-custom event.
//
// It holds 6 bytes of OEM-defined arbitrary data.
type OEMTsEvent struct {
	Timestamp        uint32
	ManfID           [3]uint8
	OEMTsDefinedData [6]uint8
}

// OEMNonTsEvent is a non-timestamped OEM-custom event.
//
// It holds 13 bytes of OEM-defined arbitrary data.
type OEMNontsEvent struct {
	OEMNontsDefinedData [13]uint8
}

// Event is included three kinds of events, Standard, OEM timestamped and OEM non-timestamped
//
// The record type decides which event should be used
type Event struct {
	RecordID   uint16
	RecordType uint8
	StandardEvent
	OEMTsEvent
	OEMNontsEvent
}

type setSystemInfoReq struct {
	paramSelector byte
	setSelector   byte
	strData       [_SYSTEM_INFO_BLK_SZ]byte
}

type DevID struct {
	DeviceID          byte
	DeviceRevision    byte
	FwRev1            byte
	FwRev2            byte
	IpmiVersion       byte
	AdtlDeviceSupport byte
	ManufacturerID    [3]byte
	ProductID         [2]byte
	AuxFwRev          [4]byte
}

type ChassisStatus struct {
	CurrentPowerState byte
	LastPowerEvent    byte
	MiscChassisState  byte
	FrontPanelButton  byte
}

type SELInfo struct {
	Version     byte
	Entries     uint16
	FreeSpace   uint16
	LastAddTime uint32
	LastDelTime uint32
	OpSupport   byte
}
