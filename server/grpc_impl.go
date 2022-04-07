package server

import (
	pb "my-fabric-offchain-storage/api"
)

type server struct {
	pb.UnimplementedDataReceiverServer
}
