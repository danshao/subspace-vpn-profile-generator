package vpnprofile

import "encoding/base64"


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
  return "./vpnprofile/windows-vpn-settings.template"
}

func (profile WindowsProfile) Host() string {
  return profile.Connection.Vpn.Host
}

func (profile WindowsProfile) PreSharedKey() string {
  return profile.Connection.Vpn.PreSharedKey
}

func (profile WindowsProfile) Username() string {
  return profile.Connection.User.UserName
}

func (profile WindowsProfile) Password() string {
  return profile.Connection.User.Password
}


// Apple profile implementation
type AppleProfile struct {
  Connection
}

func (profile AppleProfile) TemplatePath() string {
  return "./vpnprofile/apple-vpn-settings.template"
}

func (profile AppleProfile) Host() string {
  return profile.Connection.Vpn.Host
}

func (profile AppleProfile) PreSharedKey() string {
  return base64.StdEncoding.EncodeToString([]byte(profile.Connection.Vpn.PreSharedKey))
}

func (profile AppleProfile) Username() string {
  return profile.Connection.User.UserName
}

func (profile AppleProfile) Password() string {
  return profile.Connection.User.Password
}