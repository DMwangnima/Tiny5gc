package udm

import (
	"Tiny5gc/model"
	"Tiny5gc/model/udm"
)

type SubscriberDataManagement interface {
	GetNSSAI(supi *model.Supi, plmnId *model.PlmnId) (*udm.Nssai, error)
	GetAccessAndMobilitySubscriptionData(supi *model.Supi, plmnId *model.PlmnId) (*udm.AccessAndMobilitySubscriptionData, error)
	GetSmfSelectionSubscriptionData(supi *model.Supi, plmnId *model.PlmnId) (*udm.SmfSelectionSubscriptionData, error)
	GetSessionManagementSubscriptionData(supi *model.Supi, plmnId *model.PlmnId, snssai *model.Snssai, dnn *model.Dnn) (*udm.SessionManagementSubscriptionData, error)
	//GetUEContextSMFData(supi *model.Supi) (*UEContextSMFData, error)
}
