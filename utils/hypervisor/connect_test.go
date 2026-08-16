package hypervisor_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/Hari-Kiri/virest-utilities/utils/hypervisor"
	"libvirt.org/go/libvirt"
)

func TestAsLibvirtError(t *testing.T) {
	lv := libvirt.Error{Code: libvirt.ERR_AUTH_FAILED, Message: "denied"}
	got, ok := hypervisor.AsLibvirtError(lv)
	if !ok || got.Code != libvirt.ERR_AUTH_FAILED {
		t.Fatalf("expected libvirt error, got ok=%v err=%+v", ok, got)
	}

	wrapped := fmt.Errorf("connect: %w", lv)
	got, ok = hypervisor.AsLibvirtError(wrapped)
	if !ok || got.Message != "denied" {
		t.Fatalf("expected wrapped libvirt error, got ok=%v err=%+v", ok, got)
	}

	if _, ok := hypervisor.AsLibvirtError(errors.New("plain")); ok {
		t.Fatal("plain error should not match")
	}
}

func TestToLegacy(t *testing.T) {
	ve, isErr := hypervisor.ToLegacy(nil)
	if isErr || ve.Message != "" {
		t.Fatalf("nil should be success: %+v %v", ve, isErr)
	}

	ve, isErr = hypervisor.ToLegacy(libvirt.Error{Code: libvirt.ERR_NO_CONNECT, Message: "down"})
	if !isErr || ve.Code != libvirt.ERR_NO_CONNECT {
		t.Fatalf("expected libvirt legacy error: %+v", ve)
	}

	ve, isErr = hypervisor.ToLegacy(errors.New("boom"))
	if !isErr || ve.Message != "boom" || ve.Code != libvirt.ERR_INTERNAL_ERROR {
		t.Fatalf("expected synthesized legacy error: %+v", ve)
	}
}
