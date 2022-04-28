package structs

import "time"

// DirEntry is the metadata envelope for a Secure Variable
type DirEntry struct {
	Namespace   string
	Path        string
	CreatedAt   time.Time
	CreateIndex uint64
	ModifyIndex uint64
	ModifyTime  time.Time

	// reserved for post-1.4.0 work
	// LockIndex      uint64
	// Session        string
	// DeletedAt      time.Time
	// Version        uint64
	// CustomMetaData map[string]string

	EncryptedData   *DirEntryData     // removed during serialization
	UnencryptedData map[string]string // empty until serialized
}

// DirEntryData is the secret data for a Secure Variable
type DirEntryData struct {
	Data  []byte // includes nonce
	KeyID uint64 // ID of root key used to encrypt this entry
}

// SecureVariablesQuota is used to track the total size of secure
// variables entries per namespace. The total length of
// DirEntry.EncryptedData will be added to the SecureVariablesQuota
// table in the same transaction as a write, update, or delete.
type SecureVariablesQuota struct {
	Namespace   string
	Size        uint64
	ModifyIndex uint64
}

// EnvelopeKey is used to wrap a RootKey before persisting it to
// disk. It is never stored in raft.
type EnvelopeKey struct {
	KeyID     uint64 // auto-incrementing per-server
	CreatedAt time.Time
	Algorithm EncryptionAlgorithm
	Key       []byte // encoded as base64 blob
	RootKey   RootKey
}

// RootKey is used to encrypt and decrypt secure variables. It is
// never stored in raft.
type RootKey struct {
	Meta RootKeyMeta
	Key  []byte // encrypted by EnvelopeKey and encoded as base64 blob
}

// RootKeyMeta is the metadata used to refer to a RootKey. It is
// stored in raft.
type RootKeyMeta struct {
	Active           bool
	KeyID            string // UUID
	Algorithm        EncryptionAlgorithm
	EncryptionsCount uint64
	CreatedAt        time.Time
	CreateIndex      uint64
}

// EncryptionAlgorithm chooses which algorithm is used for
// encrypting / decrypting entries with this key
type EncryptionAlgorithm string

const (
	EncryptionAlgorithmXChaCha20 EncryptionAlgorithm = "xchacha20"
	EncryptionAlgorithmAES256GCM EncryptionAlgorithm = "aes256-gcm"
)

type KeyringRotateKeyRequest struct {
	Algorithm EncryptionAlgorithm
	Full      bool
	WriteRequest
}

type KeyringRotateKeyResponse struct {
	KeyID string
	WriteMeta
}

type KeyringUpdateRequest struct {
	RootKey RootKey
	WriteRequest
}

type KeyringUpdateResponse struct {
	WriteMeta
}

type KeyringListRootKeyMeta struct {
	QueryOptions
}

type KeyringListRootKeyMetaResponse struct {
	Keyring []RootKeyMeta
	QueryMeta
}

type KeyringDeleteRequest struct {
	KeyID string
	WriteRequest
}

type KeyringDeleteResponse struct {
	WriteMeta
}
