package smf

type PDUSession interface {
	CreateSMContext()
	UpdateSMContext()
	ReleaseSMContext()
	NotifySMContextStatus()
	QuerySMContext()
	Create()
	Update()
	Release()
	NotifyStatus()
}
