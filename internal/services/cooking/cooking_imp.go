package cooking

import (
	"github.com/yiixun/gotest01/v2/internal/services/apexi"
	"github.com/yiixun/gotest01/v2/pkg/apex"
	"github.com/yiixun/gotest01/v2/pkg/config"
	"github.com/yiixun/gotest01/v2/pkg/service/fishing"
	"github.com/yiixun/gotest01/v2/pkg/service/hiking"
	"github.com/yiixun/gotest01/v2/pkg/service/shopping"
	"github.com/yiixun/gotest01/v2/pkg/service/transport"
)

var (
	impl *CookingImpl = nil
)

type CookingImpl struct {
	fishing  fishing.Fishing
	shopping shopping.Shopping

	// premitive types
	conf  string
	seats int

	// Optional: check nil before use
	vehicle transport.Vehicle
	hiking  hiking.Hiking
}

func NewCookingImpl() {
	impl = new(CookingImpl)

	impl.fishing = apex.Get(apexi.FishingServiceID).(fishing.Fishing)
	// Wire to make code short
	apex.Wire(apexi.ShoppingServiceID, &impl.shopping)
	apex.Wire(config.Conf, &impl.conf)
	apex.Wire(config.DefaultLunchSeats, &impl.seats)

	// Optional fields are wired after all constructors complete. There is no dependency check.
	apex.WireOptions(
		func() {
			apex.WireOpt(apexi.VehicleServiceID, &impl.vehicle)
			apex.WireOpt(apexi.HikingServiceID, &impl.hiking)
		},
	)

	// Register happen before WireOptions. But optionals will be wired before service get invoked.
	err := apex.Register(apexi.CookingServiceID, impl)
	if err != nil {
		panic("Failed to register " + apexi.CookingServiceID)
	}
}
