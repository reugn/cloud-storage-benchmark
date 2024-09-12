package storage

// Client represents a client to interact with a storage service.
type Client interface {
	Write(string) error
	Read(string) error
	Delete(string) error
}
