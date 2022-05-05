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

func (sv *SecureVariables) Create(req *structs.SecureVariablesCreateRequest) (*structs.SecureVariablesCreateResponse, error) {
	return nil, nil
}

func (sv *SecureVariables) List(req *structs.SecureVariablesListRequest) (*structs.SecureVariablesListResponse, error) {
	return nil, nil
}

func (sv *SecureVariables) Read(req *structs.SecureVariablesReadRequest) (*structs.SecureVariablesReadResponse, error) {
	return nil, nil
}

func (sv *SecureVariables) Update(req *structs.SecureVariablesUpdateRequest) (*structs.SecureVariablesUpdateResponse, error) {
	return nil, nil
}

func (sv *SecureVariables) Destroy(req *structs.SecureVariablesDestroyRequest) (*structs.SecureVariablesDestroyResponse, error) {
	return nil, nil
}
