package fishing

import (
	"github.com/yiixun/gotest01/v2/pkg/apex"
	"github.com/yiixun/gotest01/v2/pkg/service/fishing"
)

var (
	impl *FishingImp = nil
)

type FishingImp struct {
}

func NewFishingImp() {
	impl = new(FishingImp)

	apex.Register(fishing.ServiceID, impl)
}

func (f *FishingImp) ThrowPole(strength int32) {}
func (f *FishingImp) PullPole(strength int32)  {}
