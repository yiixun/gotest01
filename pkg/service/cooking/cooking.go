package cooking

import "github.com/yiixun/gotest01/v2/pkg/dto"

const (
	ServiceID = "apex_demo.cooking"
)

type Cooking interface {
	LightFire()
	AddWater()
	AddFish(f dto.Fish)
}
