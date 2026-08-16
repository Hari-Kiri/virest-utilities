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

Legacy `utils.NewConnect*` helpers remain for compatibility but are deprecated; they now delegate to `utils/hypervisor` and surface non-libvirt failures correctly via `hypervisor.ToLegacy`.
