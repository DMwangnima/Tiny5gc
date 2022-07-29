package smf

type PDUSessionType uint8

const (
	IPV4 PDUSessionType = iota
	IPV6
	ETHERNET
	UNSTRUCTURED
)

type SNSSAI struct {
	Sst SST
	Sd  SD
}

type SSCMode uint8

const (
	// The UPF acting as anchor UPF is maintained for IP continuity
	SSCMode1 SSCMode = iota + 1
	// Release the old PDUSession and UPF, then establish a new PDUSession to the same DN with selecting a new UPF possibly
	SSCMode2
	// Establish a new PDUSession via a new anchor UPF to the same DN and select a good opportunity to release the previous PDUSession and anchor UPF
	SSCMode3
)

type CreatePDUSessionReq struct {
	PDUSessionId       uint64
	PDUSessionType     PDUSessionType
	SlicingInformation SNSSAI
	SSCMode            SSCMode
}

type PDUSessionService interface {
	CreateSMContext()
	UpdateSMContext()
	ReleaseSMContext()
	Create()
	Update()
	Release()
}
