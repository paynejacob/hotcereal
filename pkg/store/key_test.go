package store

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTypeKey_String(t *testing.T) {
	sut := TypeKey{
		Body:          "hotcerealStore",
		PackageLength: 9,
		TypeLength:    5,
	}

	assert.Equal(t, "hotcerealStore", sut.String())
}

func TestTypeKey_Bytes(t *testing.T) {
	sut := TypeKey{
		Body:          "hotcerealStore",
		PackageLength: 9,
		TypeLength:    5,
	}

	assert.Equal(t, []byte("hotcerealStore"), sut.Bytes())
}
func TestTypeKey_Package(t *testing.T) {
	sut := TypeKey{
		Body:          "hotcerealStore",
		PackageLength: 9,
		TypeLength:    5,
	}

	assert.Equal(t, "hotcereal", sut.Package())
}
func TestTypeKey_Type(t *testing.T) {
	sut := TypeKey{
		Body:          "hotcerealStore",
		PackageLength: 9,
		TypeLength:    5,
	}

	assert.Equal(t, "Store", sut.Type())
}

func TestObjectKey_String(t *testing.T) {
	sut := ObjectKey{
		TypeKey: TypeKey{
			Body:          "hotcerealStore123",
			PackageLength: 9,
			TypeLength:    5,
		},
		IdLength: 3,
	}

	assert.Equal(t, "hotcerealStore123", sut.String())
}

func TestObjectKey_Bytes(t *testing.T) {
	sut := ObjectKey{
		TypeKey: TypeKey{
			Body:          "hotcerealStore123",
			PackageLength: 9,
			TypeLength:    5,
		},
		IdLength: 3,
	}

	assert.Equal(t, []byte("hotcerealStore123"), sut.Bytes())
}
func TestObjectKey_Package(t *testing.T) {
	sut := ObjectKey{
		TypeKey: TypeKey{
			Body:          "hotcerealStore123",
			PackageLength: 9,
			TypeLength:    5,
		},
		IdLength: 3,
	}

	assert.Equal(t, "hotcereal", sut.Package())
}
func TestObjectKey_Type(t *testing.T) {
	sut := ObjectKey{
		TypeKey: TypeKey{
			Body:          "hotcerealStore123",
			PackageLength: 9,
			TypeLength:    5,
		},
		IdLength: 3,
	}

	assert.Equal(t, "Store", sut.Type())
}

func TestObjectKey_Id(t *testing.T) {
	sut := ObjectKey{
		TypeKey: TypeKey{
			Body:          "hotcerealStore123",
			PackageLength: 9,
			TypeLength:    5,
		},
		IdLength: 3,
	}

	assert.Equal(t, "123", sut.Id())
}

func TestFieldKey_String(t *testing.T) {
	sut := FieldKey{
		ObjectKey: ObjectKey{
			TypeKey: TypeKey{
				Body:          "hotcerealStore123file",
				PackageLength: 9,
				TypeLength:    5,
			},
			IdLength: 3,
		},
		FieldLength: 4,
	}

	assert.Equal(t, "hotcerealStore123file", sut.String())
}

func TestFieldKey_Bytes(t *testing.T) {
	sut := FieldKey{
		ObjectKey: ObjectKey{
			TypeKey: TypeKey{
				Body:          "hotcerealStore123file",
				PackageLength: 9,
				TypeLength:    5,
			},
			IdLength: 3,
		},
		FieldLength: 4,
	}

	assert.Equal(t, []byte("hotcerealStore123file"), sut.Bytes())
}
func TestFieldKey_Package(t *testing.T) {
	sut := FieldKey{
		ObjectKey: ObjectKey{
			TypeKey: TypeKey{
				Body:          "hotcerealStore123file",
				PackageLength: 9,
				TypeLength:    5,
			},
			IdLength: 3,
		},
		FieldLength: 4,
	}

	assert.Equal(t, "hotcereal", sut.Package())
}
func TestFieldKey_Type(t *testing.T) {
	sut := FieldKey{
		ObjectKey: ObjectKey{
			TypeKey: TypeKey{
				Body:          "hotcerealStore123file",
				PackageLength: 9,
				TypeLength:    5,
			},
			IdLength: 3,
		},
		FieldLength: 4,
	}

	assert.Equal(t, "Store", sut.Type())
}

func TestFieldKey_Id(t *testing.T) {
	sut := FieldKey{
		ObjectKey: ObjectKey{
			TypeKey: TypeKey{
				Body:          "hotcerealStore123file",
				PackageLength: 9,
				TypeLength:    5,
			},
			IdLength: 3,
		},
		FieldLength: 4,
	}

	assert.Equal(t, "123", sut.Id())
}

func TestFieldKey_Field(t *testing.T) {
	sut := FieldKey{
		ObjectKey: ObjectKey{
			TypeKey: TypeKey{
				Body:          "hotcerealStore123file",
				PackageLength: 9,
				TypeLength:    5,
			},
			IdLength: 3,
		},
		FieldLength: 4,
	}

	assert.Equal(t, "file", sut.Field())
}
