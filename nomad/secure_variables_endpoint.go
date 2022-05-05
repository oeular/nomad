package nomad

import (
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/nomad/nomad/structs"
)

// SecureVariables endpoint serves RPCs for storing and retrieving
// encrypted variables
type SecureVariables struct {
	srv       *Server
	logger    hclog.Logger
	encrypter *Encrypter
}

func (sv *SecureVariables) Create(req *structs.SecureVariablesUpsertRequest) (*structs.SecureVariablesUpsertResponse, error) {
	return nil, nil
}

func (sv *SecureVariables) List(req *structs.SecureVariablesListRequest) (*structs.SecureVariablesListResponse, error) {
	return nil, nil
}

func (sv *SecureVariables) Read(req *structs.SecureVariablesReadRequest) (*structs.SecureVariablesReadResponse, error) {
	return nil, nil
}

func (sv *SecureVariables) Update(req *structs.SecureVariablesUpsertRequest) (*structs.SecureVariablesUpsertResponse, error) {
	return nil, nil
}

func (sv *SecureVariables) Delete(req *structs.SecureVariablesDeleteRequest) (*structs.SecureVariablesDeleteResponse, error) {
	return nil, nil
}
