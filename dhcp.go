package fwsmConfig

import (
	"io"
	"net"
)

type DHCPCommon struct {
	NSs NSs
}

type DHCP struct {
	DHCPCommon

	RangeStart net.IP
	RangeEnd   net.IP
}

func (dhcp DHCPCommon) WriteTo(writer io.Writer) error {
	return nil
}

func (dhcp DHCP) WriteTo(writer io.Writer) error {
	return nil
}
