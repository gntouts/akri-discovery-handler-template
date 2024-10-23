package discoveryhandler

import (
	"context"
	"fmt"

	"github.com/gntouts/akri-discovery-handler-template/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func RegisterDiscoveryHandler(agentSocket string, discoveryHandlerName string, discoverySocket string, shared bool) error {
	creds := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.NewClient(fmt.Sprintf("unix://%s", agentSocket), creds)
	if err != nil {
		return err
	}
	defer conn.Close()
	client := pb.NewRegistrationClient(conn)
	_, err = client.RegisterDiscoveryHandler(
		context.Background(),
		&pb.RegisterDiscoveryHandlerRequest{
			Name:         discoveryHandlerName,
			Endpoint:     discoverySocket,
			Shared:       shared, // Specifies whether this device could be used by multiple nodes
			EndpointType: pb.RegisterDiscoveryHandlerRequest_UDS,
		})
	return err
}
