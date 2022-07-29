package redis_udm

import (
	"Tiny5gc/model"
	"Tiny5gc/model/udm"
)

type RedisSubscriberDataManagement struct {
	//redisCli
}

func (r RedisSubscriberDataManagement) GetNSSAI(supi *model.Supi, plmnId *model.PlmnId) (*udm.Nssai, error) {
	//TODO implement me
	panic("implement me")
}

func (r RedisSubscriberDataManagement) GetAccessAndMobilitySubscriptionData(supi *model.Supi, plmnId *model.PlmnId) (*udm.AccessAndMobilitySubscriptionData, error) {
	//TODO implement me
	panic("implement me")
}

func (r RedisSubscriberDataManagement) GetSmfSelectionSubscriptionData(supi *model.Supi, plmnId *model.PlmnId) (*udm.SmfSelectionSubscriptionData, error) {
	//TODO implement me
	panic("implement me")
}

func (r RedisSubscriberDataManagement) GetSessionManagementSubscriptionData(supi *model.Supi, plmnId *model.PlmnId, snssai *model.Snssai, dnn *model.Dnn) (*udm.SessionManagementSubscriptionData, error) {
	//TODO implement me
	panic("implement me")
}
