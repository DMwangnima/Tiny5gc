package amf

// Ts 23502-010 clause 4.2.2.2.2 general registration

type N2Header struct {
	is5G     bool
	location Location
	cellId   CellIdentity
	ratType  RATType
}

type RegisterReq struct {
	header N2Header
	typ    RegistrationType
	// whatever it takes, AMF would receive PUID from UE's registration request or old AMF or identity req.
	puId PermanentUserId
	// this parameter would be set if serving AMF has changed and new AMF would ask the old AMF for UE's SUPI and MM context
	tuId  TermporaryUserId
	pars  SecurityParameter
	nssai NSSAI
}
