package utils

import (
	"github.com/Hari-Kiri/virest-storage-pool/structures/authenticate"
	"github.com/Hari-Kiri/virest-storage-pool/structures/poolBuild"
	"github.com/Hari-Kiri/virest-storage-pool/structures/poolCreate"
	"github.com/Hari-Kiri/virest-storage-pool/structures/poolDefine"
	"github.com/Hari-Kiri/virest-storage-pool/structures/poolDetail"
	"github.com/Hari-Kiri/virest-storage-pool/structures/poolList"
	"github.com/Hari-Kiri/virest-storage-pool/structures/poolUndefine"
)

// Defined generic type constraint for response model structure.
type ResponseStructure interface {
	authenticate.Response |
		poolList.Response |
		poolDetail.Response |
		poolDefine.Response |
		poolBuild.Response |
		poolCreate.Response |
		poolUndefine.Response
}
