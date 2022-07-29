package smf

import "Tiny5gc/model"

type SmContextCreateData struct {
	Supi                    *model.Supi
	Pei                     *model.Pei
	Gpsi                    *model.Gpsi
	PduSessionId            *model.PduSessionId
	Dnn                     *model.Dnn
	Snssai                  *model.Snssai
	HplmnSnssai             *model.Snssai
	AmfId                   *model.AmfId
	requestType             *model.RequestType
	N1SmMsg                 *model.RefToBinaryData
	AnType                  *model.AccessType
	UeLocation              *model.UserLocation
	UeTimeZone              *model.TimeZone
	SmContextStatusUri      *model.Uri
	HSmfId                  *model.HSmfId
	IndicationFlags         *model.IndicationFlags
	OldPduSessionId         *model.PduSessionId
	PduSessionsActivateList []*model.PduSessionId
	UeEpsPdnConnection      *model.UeEpsPdnConnection
	HoState                 *model.HoState
	PcfId                   *model.NfInstanceId
	SupportedFeatures       *model.SupportedFeatures
}

type SmContextCreatedData struct {
}
