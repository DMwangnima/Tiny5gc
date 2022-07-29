package pcf

import "Tiny5gc/common"

type Nssai struct {
	SupportedFeatures   string
	DefaultSingleNssais []*common.Snssai
	SingleNssais        []*common.Snssai
}

type SdmSubscription struct {
	CallbackReference     *common.Uri
	MonitoredResourceUris []*common.Uri
}

type AccessAndMobilitySubscriptionData struct {
	SupportedFeatures           *common.SupportedFeatures
	Gpsis                       []*common.Gpsi
	InternalGroupIds            []*common.InternalGroupId
	SubscribedUeAmbr            *common.Ambr
	Nssai                       *Nssai
	RatRestrictions             []*common.RatType
	AreaRestrictions            []*common.AreaRestriction
	CoreNetworkTypeRestrictions []*common.CoreNetworkType
	RfspIndex                   *common.RfspIndex
	SubsRegTimer                *common.DurationSec
	UeUsageType                 *common.UeUsageType
	LadnInformation             []*common.Dnn
	MpsPriority                 bool
	ActiveTime                  *common.DurationSec
	DlPacketCount               *common.DlPacketCount
}

type SnssaiInfo struct {
	SingleNssai *common.Snssai
	DnnInfos    []*DnnInfo
}

type SmfSelectionSubscriptionData struct {
	SupportedFeatures    string
	SubscribedSnssaiInfo []*SnssaiInfo
}

type DnnInfo struct {
	Dnn                 common.Dnn
	DefaultDnnIndicator bool
	LboRoamingAllowed   bool
}

type SessionManagementSubscriptionData struct {
	SingleNssai      *common.Snssai
	DnnConfiguration map[common.Dnn]*DnnConfiguration
}

type DnnConfiguration struct {
	Dnn             common.Dnn
	PduSessionTypes *PduSessionTypes
	SscModes        *SscModes
	LadnIndicator   bool
	FiveGQosProfile *FiveGQosProfile
	// The maximum aggregated uplink and downlink bit rates to be shared across all Non-GBR Qos Flows in each PDU Session
	SessionAMBR     *common.Ambr
	StaticIpAddress *common.IpAddress
}

type FiveGQosProfile struct {
	Fiveqi common.Fiveqi
	Arp    *common.Arp
}

type PduSessionTypes struct {
	DefaultSessionType  *common.PduSessionType
	AllowedSessionTypes []*common.PduSessionType
}

type SscModes struct {
	DefaultSscMode *common.SscMode
	AllowedSscMode []*common.SscMode
}
