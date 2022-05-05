package nomad

import (
	"github.com/hashicorp/go-hclog"

	"github.com/hashicorp/nomad/nomad/structs"
)

// KeyRing endpoint serves RPCs for secure variables key management
type KeyRing struct {
	srv       *Server
	logger    hclog.Logger
	encrypter *Encrypter
}

func (k *KeyRing) RotateRootKey(req *structs.KeyringRotateRootKeyRequest) (*structs.KeyringRotateRootKeyResponse, error) {
	return nil, nil
}

func (k *KeyRing) ListRootKeyMeta(req *structs.KeyringListRootKeyMetaRequest) (*structs.KeyringListRootKeyMetaResponse, error) {
	return nil, nil
}

func (k *KeyRing) UpdateRootKey(req *structs.KeyringUpdateRootKeyRequest) (*structs.KeyringUpdateRootKeyResponse, error) {
	return nil, nil
}

func (k *KeyRing) DeleteRootKey(req *structs.KeyringDeleteRootKeyRequest) (*structs.KeyringDeleteRootKeyResponse, error) {
	return nil, nil
}
