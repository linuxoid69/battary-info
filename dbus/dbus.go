package dbus

import (
	"fmt"
	"github.com/godbus/dbus"
)

// DConnect connect to dbus
func SystemBusConnect() (conn *dbus.Conn, err error) {
	conn, err = dbus.SystemBus()
	return
}

func SessionBusConnect() (conn *dbus.Conn, err error) {
	conn, err = dbus.SessionBus()
	return
}

// GetPercentage get percentage from dbus
func GetPercentage(conn *dbus.Conn) (p float64, err error) {
	err = conn.Object("org.freedesktop.UPower", "/org/freedesktop/UPower/devices/battery_BAT0").
		Call("org.freedesktop.DBus.Properties.Get", 0, "org.freedesktop.UPower.Device", "Percentage").
		Store(&p)
	return
}

//BattNotification send notification
func BattNotification(conn *dbus.Conn, p int) (err error) {
	sp := fmt.Sprintf("Battery percentage - %d", p)
	obj := conn.Object("org.freedesktop.Notifications", "/org/freedesktop/Notifications")
	call := obj.Call("org.freedesktop.Notifications.Notify", 0, "",
		uint32(0), "battery-low", "Battery:", sp, []string{}, map[string]dbus.Variant{}, int32(5000))
	if call.Err != nil {
		return call.Err
	}
	return
}
