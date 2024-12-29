package utils

import (
	"github.com/Hari-Kiri/virest-storage-pool/structures/authenticate"
	"github.com/Hari-Kiri/virest-storage-pool/structures/getGid"
	"github.com/Hari-Kiri/virest-storage-pool/structures/getUid"
	"github.com/Hari-Kiri/virest-storage-pool/structures/poolAutostart"
	"github.com/Hari-Kiri/virest-storage-pool/structures/poolBuild"
	"github.com/Hari-Kiri/virest-storage-pool/structures/poolCapabilities"
	"github.com/Hari-Kiri/virest-storage-pool/structures/poolCreate"
	"github.com/Hari-Kiri/virest-storage-pool/structures/poolDefine"
	"github.com/Hari-Kiri/virest-storage-pool/structures/poolDelete"
	"github.com/Hari-Kiri/virest-storage-pool/structures/poolDestroy"
	"github.com/Hari-Kiri/virest-storage-pool/structures/poolDetail"
	"github.com/Hari-Kiri/virest-storage-pool/structures/poolInfo"
	"github.com/Hari-Kiri/virest-storage-pool/structures/poolList"
	"github.com/Hari-Kiri/virest-storage-pool/structures/poolRefresh"
	"github.com/Hari-Kiri/virest-storage-pool/structures/poolUndefine"
)

// Defined generic type constraint for response model structure.
type ResponseStructure interface {
	authenticate.Response |
		getUid.Response |
		getGid.Response |
		poolList.Response |
		poolInfo.Response |
		poolDetail.Response |
		poolDefine.Response |
		poolBuild.Response |
		poolCreate.Response |
		poolAutostart.Response |
		poolDestroy.Response |
		poolDelete.Response |
		poolUndefine.Response |
		poolRefresh.Response |
		poolCapabilities.Response
}
