package vpnprofile

import (
	"encoding/base64"
	"strconv"
	"time"
)

// Interface for WindowsProfile profile and AppleProfile profile.
type Profile interface {
	TemplatePath() string
	Host() string
	PreSharedKey() string
	Username() string
	Password() string
}

// Windows profile implementation
type WindowsProfile struct {
	Connection
}

func (profile WindowsProfile) TemplatePath() string {
	// return "./vpnprofile/windows-vpn-settings.template"
	return WindowsTemplate
}

func (profile WindowsProfile) Host() string {
	return profile.Connection.Vpn.Host
}

func (profile WindowsProfile) PreSharedKey() string {
	return profile.Connection.Vpn.PreSharedKey
}

func (profile WindowsProfile) Username() string {
	return profile.Connection.User.Username
}

func (profile WindowsProfile) Password() string {
	return profile.Connection.User.Password
}

// See https://msdn.microsoft.com/en-us/library/ee791258.aspx
func (profile WindowsProfile) DialParamsUID() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}


// Apple profile implementation
type AppleProfile struct {
	Connection
	Id string
	Desc string
}

func (profile AppleProfile) TemplatePath() string {
	// return "./vpnprofile/apple-vpn-settings.template"
	return AppleTemplate
}

func (profile AppleProfile) Identifier() string {
	return profile.Id
}

func (profile AppleProfile) Host() string {
	return profile.Connection.Vpn.Host
}

func (profile AppleProfile) PreSharedKey() string {
	return base64.StdEncoding.EncodeToString([]byte(profile.Connection.Vpn.PreSharedKey))
}

func (profile AppleProfile) Username() string {
	return profile.Connection.User.Username
}

func (profile AppleProfile) Password() string {
	return profile.Connection.User.Password
}

func (profile AppleProfile) Description() string {
	return profile.Desc
}