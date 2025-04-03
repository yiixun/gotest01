package apexi

import (
	"github.com/yiixun/gotest01/v2/pkg/service/cooking"
	"github.com/yiixun/gotest01/v2/pkg/service/fishing"
	"github.com/yiixun/gotest01/v2/pkg/service/hiking"
	"github.com/yiixun/gotest01/v2/pkg/service/shopping"
	"github.com/yiixun/gotest01/v2/pkg/service/transport"
)

// This is a Lookup Config which introduces the structure of the program. All internal services are listed here.
// It is the high level design/architecture of the program.
// So, this file is important, not a dirty work.
// The problem for spring like frameworks are: they introduced too many beans but never check if they are a must.

// apex suggests to program in a good design.

const (
	// Define all service ids, like in microservice env.
	CookingServiceID  = "apex_demo.cooking"
	FishingServiceID  = "apex_demo.fishing"
	HikingServiceID   = "apex_demo.hiking"
	ShoppingServiceID = "apex_demo.shopping"
	VehicleServiceID  = "apex_demo.vehicle"

	GroupA = "GroupA"
)

var (
	// from id => interfaces in pkg/internal. This map is only a place to let program know id and interface mapping.
	Services = map[string]any{
		CookingServiceID:  new(cooking.Cooking),
		FishingServiceID:  new(fishing.Fishing),
		HikingServiceID:   new(hiking.Hiking),
		ShoppingServiceID: new(shopping.Shopping),
		VehicleServiceID:  new(transport.Vehicle),
	}

	// to help wire []interface{}.
	Groups = map[string][]string{
		GroupA: {
			HikingServiceID, VehicleServiceID,
		},
	}
)
