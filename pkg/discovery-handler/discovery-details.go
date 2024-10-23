package discoveryhandler

import "github.com/gntouts/akri-discovery-handler-template/pkg/pb"

type Device interface {
	ToProtobuf() *pb.Device
}

type DiscoveryDetails interface {
	Scan() ([]Device, error)
}

func ParseDiscoveryDetails(input string) (DiscoveryDetails, error) {
	// Implement your parsing logic here
	return nil, nil
}
