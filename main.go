package main

import (
  "./vpnprofile"
  "fmt"
)

func main() {
  server := vpnprofile.Server{"my.vpn.host", "ecowork"}
  user := vpnprofile.User{"carter", "1234"}

  windowsProfile := server.GenerateProfile(vpnprofile.WINDOWS, user)
  fmt.Println(windowsProfile)

  appleProfile := server.GenerateProfile(vpnprofile.APPLE, user)
  fmt.Println(appleProfile)
}
