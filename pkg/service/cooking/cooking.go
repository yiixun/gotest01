package cooking

import "github.com/yiixun/gotest01/v2/pkg/dto"

type Cooking interface {
	LightFire()
	AddWater()
	AddFish(f dto.Fish)
}
