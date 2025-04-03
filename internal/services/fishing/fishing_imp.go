package fishing

import (
	"github.com/yiixun/gotest01/v2/internal/services/apexi"
	"github.com/yiixun/gotest01/v2/pkg/apex"
)

var (
	impl *FishingImp = nil
)

type FishingImp struct {
}

func NewFishingImp() {
	impl = new(FishingImp)

	apex.Register(apexi.FishingServiceID, impl)
}

func (f *FishingImp) ThrowPole(strength int32) {}
func (f *FishingImp) PullPole(strength int32)  {}
