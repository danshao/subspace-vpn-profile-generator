package main

import (
	"fmt"
	"gitlab.ecoworkinc.com/Subspace/vpn-profile-generator/vpnprofile"
)

func main() {
	id := vpnprofile.FormatMobileConfigIdentifier(
		"d3f8cc82-1d5c-4b5f-9629-6cbb2e577f65",
		"subspace",
		1,
		1,
	)

	descApple := "Subspace VPN"

	metadata := vpnprofile.Metadata{id, descApple}

	server := vpnprofile.Server{"54.204.175.254", "subspace"}
	user := vpnprofile.User{"1_1493786663608955659", "PwmWUdwdRN"}

	windowsProfile := server.GenerateProfile(vpnprofile.WINDOWS, user, metadata)
	fmt.Println(windowsProfile)

	appleProfile := server.GenerateProfile(vpnprofile.APPLE, user, metadata)
	fmt.Println(appleProfile)
}
