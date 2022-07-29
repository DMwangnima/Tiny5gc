package udm

type Subscription struct {
	User                      User
	DNN                       string
	IsAllowedInsertUPFInVPLMN bool
	AllowedPDUSessionType     []PDUSessionType
	AllowedSSCMode            []SSCMode
	DefaultPDUSessionType     PDUSessionType
	DefaultSSCMode            SSCMode
}
