package utils

import (
	"github.com/Hari-Kiri/virest-storage-pool/structures/poolAutostart"
	"github.com/Hari-Kiri/virest-storage-pool/structures/poolBuild"
	"github.com/Hari-Kiri/virest-storage-pool/structures/poolCreate"
	"github.com/Hari-Kiri/virest-storage-pool/structures/poolDefine"
	"github.com/Hari-Kiri/virest-storage-pool/structures/poolDelete"
	"github.com/Hari-Kiri/virest-storage-pool/structures/poolDestroy"
	"github.com/Hari-Kiri/virest-storage-pool/structures/poolDetail"
	"github.com/Hari-Kiri/virest-storage-pool/structures/poolList"
	"github.com/Hari-Kiri/virest-storage-pool/structures/poolUndefine"
)

// Defined generic type constraint for request model structure.
type RequestStructure interface {
	poolList.Request |
		poolDetail.Request |
		poolDefine.Request |
		poolBuild.Request |
		poolCreate.Request |
		poolAutostart.Request |
		poolDestroy.Request |
		poolDelete.Request |
		poolUndefine.Request
}
