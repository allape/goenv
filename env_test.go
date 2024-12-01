package goenv

import (
	"fmt"
	"os"
	"testing"
)

type AliasString string

func (a AliasString) Print() {
	fmt.Println(a)
}

func TestAlias(t *testing.T) {
	err := os.Setenv("GOENV_ALIAS", "alias")
	if err != nil {
		t.Error(err)
	}
	v := Getenv("GOENV_ALIAS", AliasString(""))
	v.Print()
}

func TestGetenv(t *testing.T) {
	err := os.Setenv("GOENV_TEST", "true")
	if err != nil {
		t.Error(err)
	}
	v := Getenv("GOENV_TEST", false)
	if v != true {
		t.Errorf("Expected true, got %v", v)
	}

	err = os.Setenv("GOENV_TEST_INT8", "8")
	if err != nil {
		t.Error(err)
	}
	vint8 := Getenv("GOENV_TEST_INT8", int8(0))
	if vint8 != 8 {
		t.Errorf("Expected 8, got %v", v)
	}

	err = os.Setenv("GOENV_TEST_INT16", "16")
	if err != nil {
		t.Error(err)
	}
	vint16 := Getenv("GOENV_TEST_INT16", int16(0))
	if vint16 != 16 {
		t.Errorf("Expected 16, got %v", v)
	}

	err = os.Setenv("GOENV_TEST_INT32", "32")
	if err != nil {
		t.Error(err)
	}
	vint32 := Getenv("GOENV_TEST_INT32", int32(0))
	if vint32 != 32 {
		t.Errorf("Expected 32, got %v", v)
	}

	err = os.Setenv("GOENV_TEST_INT64", "64")
	if err != nil {
		t.Error(err)
	}
	vint64 := Getenv("GOENV_TEST_INT64", int64(0))
	if vint64 != 64 {
		t.Errorf("Expected 64, got %v", v)
	}

	err = os.Setenv("GOENV_TEST_UINT8", "8")
	if err != nil {
		t.Error(err)
	}
	vuint8 := Getenv("GOENV_TEST_UINT8", uint8(0))
	if vuint8 != 8 {
		t.Errorf("Expected 8, got %v", v)
	}

	err = os.Setenv("GOENV_TEST_UINT16", "16")
	if err != nil {
		t.Error(err)
	}
	vuint16 := Getenv("GOENV_TEST_UINT16", uint16(0))
	if vuint16 != 16 {
		t.Errorf("Expected 16, got %v", v)
	}

	err = os.Setenv("GOENV_TEST_UINT32", "32")
	if err != nil {
		t.Error(err)
	}
	vuint32 := Getenv("GOENV_TEST_UINT32", uint32(0))
	if vuint32 != 32 {
		t.Errorf("Expected 32, got %v", v)
	}

	err = os.Setenv("GOENV_TEST_UINT64", "64")
	if err != nil {
		t.Error(err)
	}
	vuint64 := Getenv("GOENV_TEST_UINT64", uint64(0))
	if vuint64 != 64 {
		t.Errorf("Expected 64, got %v", v)
	}

	err = os.Setenv("GOENV_TEST_FLOAT32", "32.32")
	if err != nil {
		t.Error(err)
	}
	vfloat32 := Getenv("GOENV_TEST_FLOAT32", float32(0))
	if vfloat32 != 32.32 {
		t.Errorf("Expected 32.32, got %v", v)
	}

	err = os.Setenv("GOENV_TEST_FLOAT64", "64.64")
	if err != nil {
		t.Error(err)
	}
	vfloat64 := Getenv("GOENV_TEST_FLOAT64", float64(0))
	if vfloat64 != 64.64 {
		t.Errorf("Expected 64.64, got %v", v)
	}

	err = os.Setenv("GOENV_TEST_COMPLEX64", "64.64")
	if err != nil {
		t.Error(err)
	}
	vcomplex64 := Getenv("GOENV_TEST_COMPLEX64", complex64(0))
	if vcomplex64 != 64.64 {
		t.Errorf("Expected 64.64, got %v", v)
	}

	err = os.Setenv("GOENV_TEST_COMPLEX128", "128.128")
	if err != nil {
		t.Error(err)
	}
	vcomplex128 := Getenv("GOENV_TEST_COMPLEX128", complex128(0))
	if vcomplex128 != 128.128 {
		t.Errorf("Expected 128.128, got %v", v)
	}

	err = os.Setenv("GOENV_TEST_STRING", "string")
	if err != nil {
		t.Error(err)
	}
	vstring := Getenv("GOENV_TEST_STRING", "")
	if vstring != "string" {
		t.Errorf("Expected string, got %v", v)
	}

	err = os.Unsetenv("GOENV_TEST_DEFAULT")
	if err != nil {
		t.Error(err)
	}
	vdefault := Getenv("GOENV_TEST_DEFAULT", "default")
	if vdefault != "default" {
		t.Errorf("Expected default, got %v", v)
	}
}
