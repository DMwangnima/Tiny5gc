package udm

import "Tiny5gc/model"

type Nssai struct {
	SupportedFeatures   string
	DefaultSingleNssais []*model.Snssai
	SingleNssais        []*model.Snssai
}

type SdmSubscription struct {
	CallbackReference     *model.Uri
	MonitoredResourceUris []*model.Uri
}

type AccessAndMobilitySubscriptionData struct {
	SupportedFeatures           *model.SupportedFeatures
	Gpsis                       []*model.Gpsi
	InternalGroupIds            []*model.InternalGroupId
	SubscribedUeAmbr            *model.Ambr
	Nssai                       *Nssai
	RatRestrictions             []*model.RatType
	AreaRestrictions            []*model.AreaRestriction
	CoreNetworkTypeRestrictions []*model.CoreNetworkType
	RfspIndex                   *model.RfspIndex
	SubsRegTimer                *model.DurationSec
	UeUsageType                 *model.UeUsageType
	LadnInformation             []*model.Dnn
	MpsPriority                 bool
	ActiveTime                  *model.DurationSec
	DlPacketCount               *model.DlPacketCount
}

type SmfSelectionSubscriptionData struct {
	SupportedFeatures    string
	SubscribedSnssaiInfo []*SnssaiInfo
}

type SnssaiInfo struct {
	SingleNssai *model.Snssai
	DnnInfos    []*DnnInfo
}

type DnnInfo struct {
	Dnn                 model.Dnn
	DefaultDnnIndicator bool
	LboRoamingAllowed   bool
}

type SessionManagementSubscriptionData struct {
	SingleNssai      *model.Snssai
	DnnConfiguration map[model.Dnn]*DnnConfiguration
}

type DnnConfiguration struct {
	Dnn             model.Dnn
	PduSessionTypes *PduSessionTypes
	SscModes        *SscModes
	LadnIndicator   bool
	FiveGQosProfile *FiveGQosProfile
	// The maximum aggregated uplink and downlink bit rates to be shared across all Non-GBR Qos Flows in each PDU Session
	SessionAMBR     *model.Ambr
	StaticIpAddress *model.IpAddress
}

type FiveGQosProfile struct {
	Fiveqi model.Fiveqi
	Arp    *model.Arp
}

type PduSessionTypes struct {
	DefaultSessionType  *model.PduSessionType
	AllowedSessionTypes []*model.PduSessionType
}

type SscModes struct {
	DefaultSscMode *model.SscMode
	AllowedSscMode []*model.SscMode
}
