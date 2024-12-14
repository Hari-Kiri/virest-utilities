package utils

import (
	"github.com/Hari-Kiri/virest-storage-pool/structures/poolDefine"
	"github.com/Hari-Kiri/virest-storage-pool/structures/poolList"
	"github.com/Hari-Kiri/virest-storage-pool/structures/poolUndefine"
)

// Defined generic type constraint for request model structure.
type RequestStructure interface {
	poolList.Request |
		poolDefine.Request |
		poolUndefine.Request
}
