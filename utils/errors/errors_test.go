package errors_test

import (
	stderrors "errors"
	"fmt"
	"testing"

	viresterrors "github.com/Hari-Kiri/virest-utilities/utils/errors"
	"libvirt.org/go/libvirt"
)

func TestWrap(t *testing.T) {
	if err := viresterrors.Wrap("op", nil); err != nil {
		t.Fatalf("Wrap(nil)=%v", err)
	}

	lv := libvirt.Error{Code: libvirt.ERR_NO_STORAGE_POOL, Message: "missing"}
	wrapped := viresterrors.Wrap("lookup", lv)
	if wrapped == nil {
		t.Fatal("expected wrapped error")
	}
	got, ok := viresterrors.AsLibvirtError(wrapped)
	if !ok || got.Code != libvirt.ERR_NO_STORAGE_POOL {
		t.Fatalf("unwrap failed: ok=%v got=%+v", ok, got)
	}
	if len(wrapped.Error()) < 7 || wrapped.Error()[:7] != "lookup:" {
		t.Fatalf("unexpected message: %v", wrapped)
	}
}

func TestAsLibvirtError(t *testing.T) {
	lv := libvirt.Error{Code: libvirt.ERR_AUTH_FAILED, Message: "denied"}
	got, ok := viresterrors.AsLibvirtError(lv)
	if !ok || got.Code != libvirt.ERR_AUTH_FAILED {
		t.Fatalf("expected libvirt error, got ok=%v err=%+v", ok, got)
	}

	wrapped := fmt.Errorf("connect: %w", lv)
	got, ok = viresterrors.AsLibvirtError(wrapped)
	if !ok || got.Message != "denied" {
		t.Fatalf("expected wrapped libvirt error, got ok=%v err=%+v", ok, got)
	}

	if _, ok := viresterrors.AsLibvirtError(stderrors.New("plain")); ok {
		t.Fatal("plain error should not match")
	}
}

func TestCode(t *testing.T) {
	lv := libvirt.Error{Code: libvirt.ERR_AUTH_FAILED, Message: "denied"}
	code, ok := viresterrors.Code(viresterrors.Wrap("connect", lv))
	if !ok || code != libvirt.ERR_AUTH_FAILED {
		t.Fatalf("code=%v ok=%v", code, ok)
	}
	if _, ok := viresterrors.Code(stderrors.New("plain")); ok {
		t.Fatal("plain error should not yield a code")
	}
	if _, ok := viresterrors.Code(nil); ok {
		t.Fatal("nil should not yield a code")
	}
}

func TestToLegacy(t *testing.T) {
	ve, isErr := viresterrors.ToLegacy(nil)
	if isErr || ve.Message != "" {
		t.Fatalf("nil should be success: %+v %v", ve, isErr)
	}

	ve, isErr = viresterrors.ToLegacy(libvirt.Error{Code: libvirt.ERR_NO_CONNECT, Message: "down"})
	if !isErr || ve.Code != libvirt.ERR_NO_CONNECT {
		t.Fatalf("expected libvirt legacy error: %+v", ve)
	}

	ve, isErr = viresterrors.ToLegacy(stderrors.New("boom"))
	if !isErr || ve.Message != "boom" || ve.Code != libvirt.ERR_INTERNAL_ERROR {
		t.Fatalf("expected synthesized legacy error: %+v", ve)
	}
}
