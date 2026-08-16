# virest-utilities

ViRest utility tools.

## Hypervisor connection

Prefer the idiomatic connect helpers in `utils/hypervisor` (returns `(virest.Connection, error)` and does not require importing the parent `utils` package, which pulls in libguestfs via disk helpers):

```go
import "github.com/Hari-Kiri/virest-utilities/utils/hypervisor"

conn, err := hypervisor.Connect("qemu:///system")
if err != nil {
    return err
}
defer conn.Close()
```

Also available: `ConnectReadOnly`, `ConnectWithAuth`, `ConnectWithAuthDefault`.

Legacy `utils.NewConnect*` helpers remain for compatibility but are deprecated; they now delegate to `utils/hypervisor` and surface non-libvirt failures correctly via `utils/errors.ToLegacy`.

## Errors

Use `utils/errors` to annotate failures and recover libvirt details (import under an alias such as `viresterrors` to avoid clashing with the standard library `errors` package):

```go
import viresterrors "github.com/Hari-Kiri/virest-utilities/utils/errors"

err := viresterrors.Wrap("define pool", someErr)
if err != nil {
    if code, ok := viresterrors.Code(err); ok {
        // use libvirt.ErrorNumber code
        _ = code
    }
}
```

Also available: `AsLibvirtError`, `ToLegacy`.
