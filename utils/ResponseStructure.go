package utils

import (
	"github.com/Hari-Kiri/virest-storage-pool/structures/poolDefine"
	"github.com/Hari-Kiri/virest-storage-pool/structures/poolList"
	"github.com/Hari-Kiri/virest-storage-pool/structures/poolUndefine"
)

// Defined generic type constraint for response model structure.
type ResponseStructure interface {
	poolList.Response |
		poolDefine.Response |
		poolUndefine.Response
}
