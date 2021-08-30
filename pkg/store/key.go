package store

type Key interface {
	String() string
	Bytes() []byte
}

type TypeKey struct {
	Body          string
	PackageLength int
	TypeLength    int
}

func (k TypeKey) String() string {
	return k.Body
}

func (k TypeKey) Bytes() []byte {
	return []byte(k.Body)
}

func (k TypeKey) Package() string {
	return k.Body[0:k.PackageLength]
}

func (k TypeKey) Type() string {
	return k.Body[k.PackageLength : k.PackageLength+k.TypeLength]
}

type ObjectKey struct {
	TypeKey
	IdLength int
}

func (k ObjectKey) Id() string {
	return k.Body[k.PackageLength+k.TypeLength : k.PackageLength+k.TypeLength+k.IdLength]
}

type FieldKey struct {
	ObjectKey
	FieldLength int
}

func (k FieldKey) Field() string {
	return k.Body[k.PackageLength+k.TypeLength+k.IdLength:]
}
