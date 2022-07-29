package amf

func RegisterManagement() {

}

func ConnectManagement() {

}

// how to select a SMF?
func SessionManagement() {

}

type SMInformation struct {
	pduType               PDUType
	sscMode               SSCMode
	protocolConfiguration interface{}
}

type EstablishPDUSessionReq struct {
	snssai        SNSSAI
	dnn           DNN
	pduSessionId  uint64
	smInformation SMInformation
}

func EstablishPDUSession() {
	// select smf to serve the process
}

type CommunicationService interface {
	UEContextTransfer()
	RegistrationCompleteNotify()
	N1MessageNotify()
	N1N2MessageTransfer()
}

type EventExposureService interface {
	Subscribe()
	Unsubscribe()
	Notify()
}

type MTService interface {
	EnableUEReachability()
}
