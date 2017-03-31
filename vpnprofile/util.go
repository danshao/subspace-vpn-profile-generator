package vpnprofile

import (
	"bytes"
	"text/template"
)

// Enum for platform
type Platform uint8

const (
	WINDOWS Platform = iota + 1
	APPLE
)

type Server struct {
	Host         string
	PreSharedKey string
}

type User struct {
	Username string
	Password string
}

type Connection struct {
	Vpn  Server
	User User
}

func (vpn Server) GenerateProfile(platform Platform, user User) string {
	connection := Connection{vpn, user}
	var profile Profile
	switch platform {
	case WINDOWS:
		profile = WindowsProfile{connection}
	case APPLE:
		profile = AppleProfile{connection}
	default:
		panic("Support only APPLE or WINDOWS.")
	}

	//TODO If path is not the best practice, than use const blow
	// tmpl, err := template.ParseFiles(profile.TemplatePath())
	tmpl, err := template.New("template.name").Parse(profile.TemplatePath())
	if err != nil {
		panic(err)
	}

	var doc bytes.Buffer
	err = tmpl.Execute(&doc, profile)
	if err != nil {
		panic(err)
	}

	return doc.String()
}

const WindowsTemplate = "[VPN Connection]\n" +
	"Encoding=1\n" +
	"PBVersion=1\n" +
	"Type=2\n" +
	"AutoLogon=0\n" +
	"UseRasCredentials=1\n" +
	"LowDateTime=1332934000\n" +
	"HighDateTime=30576000\n" +
	"DialParamsUID=56078163\n" +
	"Guid=416FEFEA6114F44B8D96246554F43FFB\n" +
	"VpnStrategy=3\n" +
	"ExcludedProtocols=0\n" +
	"LcpExtensions=1\n" +
	"DataEncryption=256\n" +
	"SwCompression=0\n" +
	"NegotiateMultilinkAlways=0\n" +
	"SkipDoubleDialDialog=0\n" +
	"DialMode=0\n" +
	"OverridePref=15\n" +
	"RedialAttempts=3\n" +
	"RedialSeconds=60\n" +
	"IdleDisconnectSeconds=0\n" +
	"RedialOnLinkFailure=1\n" +
	"CallbackMode=0\n" +
	"CustomDialDll=\n" +
	"CustomDialFunc=\n" +
	"CustomRasDialDll=\n" +
	"ForceSecureCompartment=0\n" +
	"DisableIKENameEkuCheck=0\n" +
	"AuthenticateServer=0\n" +
	"ShareMsFilePrint=1\n" +
	"BindMsNetClient=1\n" +
	"SharedPhoneNumbers=0\n" +
	"GlobalDeviceSettings=0\n" +
	"PrerequisiteEntry=\n" +
	"PrerequisitePbk=\n" +
	"PreferredPort=VPN2-0\n" +
	"PreferredDevice=WAN Miniport (L2TP)\n" +
	"PreferredBps=0\n" +
	"PreferredHwFlow=1\n" +
	"PreferredProtocol=1\n" +
	"PreferredCompression=1\n" +
	"PreferredSpeaker=1\n" +
	"PreferredMdmProtocol=0\n" +
	"PreviewUserPw=1\n" +
	"PreviewDomain=1\n" +
	"PreviewPhoneNumber=0\n" +
	"ShowDialingProgress=1\n" +
	"ShowMonitorIconInTaskBar=1\n" +
	"CustomAuthKey=0\n" +
	"AuthRestrictions=544\n" +
	"IpPrioritizeRemote=1\n" +
	"IpInterfaceMetric=0\n" +
	"IpHeaderCompression=0\n" +
	"IpAddress=0.0.0.0\n" +
	"IpDnsAddress=0.0.0.0\n" +
	"IpDns2Address=0.0.0.0\n" +
	"IpWinsAddress=0.0.0.0\n" +
	"IpWins2Address=0.0.0.0\n" +
	"IpAssign=1\n" +
	"IpNameAssign=1\n" +
	"IpDnsFlags=1\n" +
	"IpNBTFlags=1\n" +
	"TcpWindowSize=0\n" +
	"UseFlags=2\n" +
	"IpSecFlags=1\n" +
	"IpDnsSuffix=\n" +
	"Ipv6Assign=1\n" +
	"Ipv6Address=::\n" +
	"Ipv6PrefixLength=0\n" +
	"Ipv6PrioritizeRemote=1\n" +
	"Ipv6InterfaceMetric=0\n" +
	"Ipv6NameAssign=1\n" +
	"Ipv6DnsAddress=::\n" +
	"Ipv6Dns2Address=::\n" +
	"Ipv6Prefix=0000000000000000\n" +
	"Ipv6InterfaceId=0000000000000000\n" +
	"DisableClassBasedDefaultRoute=0\n" +
	"DisableMobility=0\n" +
	"NetworkOutageTime=0\n" +
	"ProvisionType=0\n" +
	"PreSharedKey={{.PreSharedKey}}\n" +

	"NETCOMPONENTS=\n" +
	"ms_msclient=1\n" +
	"ms_server=1\n" +

	"MEDIA=rastapi\n" +
	"Port=VPN2-0\n" +
	"Device=WAN Miniport (L2TP)\n" +

	"DEVICE=vpn\n" +
	"PhoneNumber={{.Host}}\n" +
	"AreaCode=\n" +
	"CountryCode=0\n" +
	"CountryID=0\n" +
	"UseDialingRules=0\n" +
	"Comment=\n" +
	"FriendlyName=\n" +
	"LastSelectedPhone=0\n" +
	"PromoteAlternates=0\n" +
	"TryNextAlternateOnFail=1\n"

const AppleTemplate = "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n" +
	"<!DOCTYPE plist PUBLIC \"-//AppleProfile//DTD PLIST 1.0//EN\" \"http://www.apple.com/DTDs/PropertyList-1.0.dtd\">\n" +
	"<plist version=\"1.0\">\n" +
	"<dict>\n" +
	"    <key>PayloadContent</key>\n" +
	"    <array>\n" +
	"        <dict>\n" +
	"            <key>IPSec</key>\n" +
	"            <dict>\n" +
	"                <key>AuthenticationMethod</key>\n" +
	"                <string>SharedSecret</string>\n" +
	"                <key>LocalIdentifierType</key>\n" +
	"                <string>KeyID</string>\n" +
	"                <key>SharedSecret</key>\n" +
	"                <data>\n" +
	"                {{.PreSharedKey}}\n" +
	"                </data>\n" +
	"            </dict>\n" +
	"            <key>IPv4</key>\n" +
	"            <dict>\n" +
	"                <key>OverridePrimary</key>\n" +
	"                <integer>1</integer>\n" +
	"            </dict>\n" +
	"            <key>PPP</key>\n" +
	"            <dict>\n" +
	"                <key>AuthName</key>\n" +
	"                <string>{{.Username}}</string>\n" +
	"                <key>AuthPassword</key>\n" +
	"                <string>{{.Password}}</string>\n" +
	"                <key>CommRemoteAddress</key>\n" +
	"                <string>{{.Host}}</string>\n" +
	"            </dict>\n" +
	"            <key>PayloadDescription</key>\n" +
	"            <string>配置 VPN 設定</string>\n" +
	"            <key>PayloadDisplayName</key>\n" +
	"            <string>VPN</string>\n" +
	"            <key>PayloadIdentifier</key>\n" +
	"            <string>com.apple.vpn.managed.C3A1F710-0776-425F-85D9-91183D46B573</string>\n" +
	"            <key>PayloadType</key>\n" +
	"            <string>com.apple.vpn.managed</string>\n" +
	"            <key>PayloadUUID</key>\n" +
	"            <string>763C9CD6-C7E8-460C-AB56-30C54077C70B</string>\n" +
	"            <key>PayloadVersion</key>\n" +
	"            <integer>1</integer>\n" +
	"            <key>Proxies</key>\n" +
	"            <dict>\n" +
	"                <key>HTTPEnable</key>\n" +
	"                <integer>0</integer>\n" +
	"                <key>HTTPSEnable</key>\n" +
	"                <integer>0</integer>\n" +
	"            </dict>\n" +
	"            <key>UserDefinedName</key>\n" +
	"            <string>Carter Test on Device</string>\n" +
	"            <key>VPNType</key>\n" +
	"            <string>L2TP</string>\n" +
	"        </dict>\n" +
	"    </array>\n" +
	"    <key>PayloadDisplayName</key>\n" +
	"    <string>Carter VPN Test</string>\n" +
	"    <key>PayloadIdentifier</key>\n" +
	"    <string>cateyesde-macbook-pro.local.9A7D5D7B-525B-4D30-AD87-C3EA25419C63</string>\n" +
	"    <key>PayloadRemovalDisallowed</key>\n" +
	"    <false/>\n" +
	"    <key>PayloadType</key>\n" +
	"    <string>Configuration</string>\n" +
	"    <key>PayloadUUID</key>\n" +
	"    <string>8466A3FD-BF51-4B2C-A128-D5A5E901A8DA</string>\n" +
	"    <key>PayloadVersion</key>\n" +
	"    <integer>1</integer>\n" +
	"</dict>\n" +
	"</plist>\n"
