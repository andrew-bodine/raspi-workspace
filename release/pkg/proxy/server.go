package proxy

import (
	"context"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"

	pb "github.com/andrew-bodine/raspi/release/pkg/protobufs"
)

type ReleaseServer interface {
	Serve(listener net.Listener) error
	Stop()

	pb.ReleaseServiceServer
}

type releaseServer struct {
	grpcServer *grpc.Server

	cacher Cacher
	stopCh <-chan struct{}
}

func NewReleaseServer(stopCh <-chan struct{}, lister ReleaseLister) ReleaseServer {
	releaseServer := &releaseServer{
		grpcServer: grpc.NewServer(),
		cacher:     NewCacher(lister),
		stopCh:     stopCh,
	}

	pb.RegisterReleaseServiceServer(releaseServer.grpcServer, releaseServer)

	return releaseServer
}

func (rs *releaseServer) Serve(listener net.Listener) error {
	go rs.waitForStopSignal()
	go rs.cacher.Run()

	return rs.grpcServer.Serve(listener)
}

func (rs *releaseServer) waitForStopSignal() {
	for {
		select {
		case <-rs.stopCh:
			rs.Stop()
			return
		}
	}
}

func (rs *releaseServer) Stop() {
	rs.grpcServer.Stop()
	rs.cacher.Stop()
}

func (rs *releaseServer) GetReleases(ctx context.Context, req *pb.GetReleasesRequest) (*pb.ReleasesResponse, error) {
	log.Println("GetReleases called")

	resp := pb.ReleasesResponse{
		Count:    int32(10),
		Releases: []*pb.Release{},
	}

	return &resp, nil
}

func (rs *releaseServer) WatchReleases(req *pb.WatchReleasesRequest, stream pb.ReleaseService_WatchReleasesServer) error {
	log.Println("WatchReleases called")

	// TODO: Send current releases.

	// TODO: Append stream to servers memory mapping for current connected agents.

	for i := 1; i <= 5; i++ {
		resp := pb.ReleasesResponse{
			Count:    int32(100 * i),
			Releases: []*pb.Release{},
		}

		if err := stream.Send(&resp); err != nil {
			return err
		}

		time.Sleep(time.Second)
	}

	return nil
}
