// +build !appengine

package packages

import (
	"net"
)

func init() {
	Packages["net"] = map[string]interface{}{
		"CIDRMask":                   net.CIDRMask,
		"Dial":                       net.Dial,
		"DialIP":                     net.DialIP,
		"DialTCP":                    net.DialTCP,
		"DialTimeout":                net.DialTimeout,
		"DialUDP":                    net.DialUDP,
		"DialUnix":                   net.DialUnix,
		"ErrWriteToConnected":        net.ErrWriteToConnected,
		"FileConn":                   net.FileConn,
		"FlagBroadcast":              net.FlagBroadcast,
		"FlagLoopback":               net.FlagLoopback,
		"FlagMulticast":              net.FlagMulticast,
		"FlagPointToPoint":           net.FlagPointToPoint,
		"FlagUp":                     net.FlagUp,
		"IPv4":                       net.IPv4,
		"IPv4Mask":                   net.IPv4Mask,
		"IPv4allrouter":              net.IPv4allrouter,
		"IPv4allsys":                 net.IPv4allsys,
		"IPv4bcast":                  net.IPv4bcast,
		"IPv4len":                    net.IPv4len,
		"IPv4zero":                   net.IPv4zero,
		"IPv6interfacelocalallnodes": net.IPv6interfacelocalallnodes,
		"IPv6len":                    net.IPv6len,
		"IPv6linklocalallnodes":      net.IPv6linklocalallnodes,
		"IPv6linklocalallrouters":    net.IPv6linklocalallrouters,
		"IPv6loopback":               net.IPv6loopback,
		"IPv6unspecified":            net.IPv6unspecified,
		"IPv6zero":                   net.IPv6zero,
		"JoinHostPort":               net.JoinHostPort,
		"ParseCIDR":                  net.ParseCIDR,
		"ParseIP":                    net.ParseIP,
		"ParseMAC":                   net.ParseMAC,
		"ResolveIPAddr":              net.ResolveIPAddr,
		"ResolveTCPAddr":             net.ResolveTCPAddr,
		"ResolveUDPAddr":             net.ResolveUDPAddr,
		"ResolveUnixAddr":            net.ResolveUnixAddr,
		"SplitHostPort":              net.SplitHostPort,
	}
}
