package hypervisor_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/Hari-Kiri/virest-utilities/utils/hypervisor"
	"libvirt.org/go/libvirt"
)

func TestWrap(t *testing.T) {
	if err := hypervisor.Wrap("op", nil); err != nil {
		t.Fatalf("Wrap(nil)=%v", err)
	}

	lv := libvirt.Error{Code: libvirt.ERR_NO_STORAGE_POOL, Message: "missing"}
	wrapped := hypervisor.Wrap("lookup", lv)
	if wrapped == nil {
		t.Fatal("expected wrapped error")
	}
	got, ok := hypervisor.AsLibvirtError(wrapped)
	if !ok || got.Code != libvirt.ERR_NO_STORAGE_POOL {
		t.Fatalf("unwrap failed: ok=%v got=%+v", ok, got)
	}
	if wrapped.Error() != "lookup: missing" && wrapped.Error()[:7] != "lookup:" {
		// Message formatting depends on libvirt.Error.Error(); ensure prefix at least.
		if fmt.Sprintf("%v", wrapped)[:7] != "lookup:" {
			t.Fatalf("unexpected message: %v", wrapped)
		}
	}
}

func TestCode(t *testing.T) {
	lv := libvirt.Error{Code: libvirt.ERR_AUTH_FAILED, Message: "denied"}
	code, ok := hypervisor.Code(hypervisor.Wrap("connect", lv))
	if !ok || code != libvirt.ERR_AUTH_FAILED {
		t.Fatalf("code=%v ok=%v", code, ok)
	}
	if _, ok := hypervisor.Code(errors.New("plain")); ok {
		t.Fatal("plain error should not yield a code")
	}
	if _, ok := hypervisor.Code(nil); ok {
		t.Fatal("nil should not yield a code")
	}
}
