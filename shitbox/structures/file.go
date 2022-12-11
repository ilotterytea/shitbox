package shitbox

// Shitbox's file data.
type File struct {
	// Literal ID.
	LiteralId string
	// MIME type of the file.
	MIME string
	// Extension of the file.
	Ext string
	// Secret key. It is used to perform some operations with the file.
	SecretKey string
	// A timestamp of when the file was uploaded.
	Timestamp int64
}
