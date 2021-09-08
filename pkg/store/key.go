package store

const (
	TypeKeySuffix int = iota
	ObjectKeySuffix
	FieldKeySuffix
)

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
	return k.Body[0:k.PackageLength+k.TypeLength] + string(rune(TypeKeySuffix))
}

func (k TypeKey) Bytes() []byte {
	return []byte(k.String())
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

func (k ObjectKey) String() string {
	return k.TypeKey.String() + k.Id() + string(rune(ObjectKeySuffix))
}

func (k ObjectKey) Bytes() []byte {
	return []byte(k.String())
}

func (k ObjectKey) Id() string {
	return k.Body[k.PackageLength+k.TypeLength : k.PackageLength+k.TypeLength+k.IdLength]
}

type FieldKey struct {
	ObjectKey
	FieldLength int
}

func (k FieldKey) String() string {
	return k.ObjectKey.String() + k.Field() + string(rune(FieldKeySuffix))
}

func (k FieldKey) Bytes() []byte {
	return []byte(k.String())
}

func (k FieldKey) Field() string {
	return k.Body[k.PackageLength+k.TypeLength+k.IdLength : k.PackageLength+k.TypeLength+k.IdLength+k.FieldLength]
}
