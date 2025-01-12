package utils

import (
	"github.com/Hari-Kiri/virest-storage-pool/structures/findStoragePoolSources"
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
	"github.com/Hari-Kiri/virest-storage-volume/structures/volumeCreate"
	"github.com/Hari-Kiri/virest-storage-volume/structures/volumeDelete"
	"github.com/Hari-Kiri/virest-storage-volume/structures/volumeListAll"
)

// Defined generic type constraint for request model structure.
type RequestStructure interface {
	findStoragePoolSources.Request |
		getUid.Request |
		getGid.Request |
		poolList.Request |
		poolInfo.Request |
		poolDetail.Request |
		poolDefine.Request |
		poolBuild.Request |
		poolCreate.Request |
		poolAutostart.Request |
		poolDestroy.Request |
		poolDelete.Request |
		poolUndefine.Request |
		poolRefresh.Request |
		poolCapabilities.Request |
		volumeListAll.Request |
		volumeCreate.Request |
		volumeDelete.Request
}
