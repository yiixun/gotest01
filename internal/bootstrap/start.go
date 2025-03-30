package bootstrap

import (
	"github.com/yiixun/gotest01/v2/internal/services/cooking"
	"github.com/yiixun/gotest01/v2/internal/services/fishing"
	"github.com/yiixun/gotest01/v2/pkg/apex"
)

func Initialize() {

	// apex start:
	apex.Bootstrap()

	// service start
	apex.Bootup(
		cooking.NewCookingImpl,
		fishing.NewFishingImp,
	)

}
