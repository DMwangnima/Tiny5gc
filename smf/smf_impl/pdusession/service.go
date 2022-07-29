package pdusession

import (
	"Tiny5gc/model"
	smfModel "Tiny5gc/model/smf"
	"sync"
)

type smContextKey string

func newSmContextKey(supi model.Supi, id model.PduSessionId) smContextKey {

}

type Service struct {
	mu         sync.Mutex
	smContexts map[smContextKey]*SMContext
}

func (s *Service) CreateSMContext(req *smfModel.SmContextCreateData) (resp *smfModel.SmContextCreatedData, err error) {
	if req.Supi == nil {
		// emergency registered and UICCless
	}
	var smContext *SMContext
	var supi model.Supi
	var id model.PduSessionId
	key := newSmContextKey(supi, id)

	s.mu.Lock()
	smContext = s.smContexts[key]
	if smContext != nil {
	}
	s.mu.Unlock()
}

func (s *Service) UpdateSMContext() {
	//TODO implement me
	panic("implement me")
}

func (s *Service) ReleaseSMContext() {
	//TODO implement me
	panic("implement me")
}

func (s *Service) NotifySMContextStatus() {
	//TODO implement me
	panic("implement me")
}

func (s *Service) QuerySMContext() {
	//TODO implement me
	panic("implement me")
}

func (s *Service) Create() {
	//TODO implement me
	panic("implement me")
}

func (s *Service) Update() {
	//TODO implement me
	panic("implement me")
}

func (s *Service) Release() {
	//TODO implement me
	panic("implement me")
}

func (s *Service) NotifyStatus() {
	//TODO implement me
	panic("implement me")
}
