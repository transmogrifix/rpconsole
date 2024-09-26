// Code generated by protoc-gen-connect-gateway. DO NOT EDIT.
//
// Source: redpanda/api/console/v1alpha1/debug_bundle.proto

package consolev1alpha1connect

import (
	context "context"
	fmt "fmt"

	runtime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	connect_gateway "go.vallahaye.net/connect-gateway"

	v1alpha1 "github.com/redpanda-data/console/backend/pkg/protogen/redpanda/api/console/v1alpha1"
)

// DebugBundleServiceGatewayServer implements the gRPC server API for the DebugBundleService
// service.
type DebugBundleServiceGatewayServer struct {
	v1alpha1.UnimplementedDebugBundleServiceServer
	createDebugBundle     connect_gateway.UnaryHandler[v1alpha1.CreateDebugBundleRequest, v1alpha1.CreateDebugBundleResponse]
	getDebugBundleStatus  connect_gateway.UnaryHandler[v1alpha1.GetDebugBundleStatusRequest, v1alpha1.GetDebugBundleStatusResponse]
	deleteDebugBundle     connect_gateway.UnaryHandler[v1alpha1.DeleteDebugBundleRequest, v1alpha1.DeleteDebugBundleResponse]
	deleteDebugBundleFile connect_gateway.UnaryHandler[v1alpha1.DeleteDebugBundleFileRequest, v1alpha1.DeleteDebugBundleFileResponse]
}

// NewDebugBundleServiceGatewayServer constructs a Connect-Gateway gRPC server for the
// DebugBundleService service.
func NewDebugBundleServiceGatewayServer(svc DebugBundleServiceHandler, opts ...connect_gateway.HandlerOption) *DebugBundleServiceGatewayServer {
	return &DebugBundleServiceGatewayServer{
		createDebugBundle:     connect_gateway.NewUnaryHandler(DebugBundleServiceCreateDebugBundleProcedure, svc.CreateDebugBundle, opts...),
		getDebugBundleStatus:  connect_gateway.NewUnaryHandler(DebugBundleServiceGetDebugBundleStatusProcedure, svc.GetDebugBundleStatus, opts...),
		deleteDebugBundle:     connect_gateway.NewUnaryHandler(DebugBundleServiceDeleteDebugBundleProcedure, svc.DeleteDebugBundle, opts...),
		deleteDebugBundleFile: connect_gateway.NewUnaryHandler(DebugBundleServiceDeleteDebugBundleFileProcedure, svc.DeleteDebugBundleFile, opts...),
	}
}

func (s *DebugBundleServiceGatewayServer) CreateDebugBundle(ctx context.Context, req *v1alpha1.CreateDebugBundleRequest) (*v1alpha1.CreateDebugBundleResponse, error) {
	return s.createDebugBundle(ctx, req)
}

func (s *DebugBundleServiceGatewayServer) GetDebugBundleStatus(ctx context.Context, req *v1alpha1.GetDebugBundleStatusRequest) (*v1alpha1.GetDebugBundleStatusResponse, error) {
	return s.getDebugBundleStatus(ctx, req)
}

func (s *DebugBundleServiceGatewayServer) DeleteDebugBundle(ctx context.Context, req *v1alpha1.DeleteDebugBundleRequest) (*v1alpha1.DeleteDebugBundleResponse, error) {
	return s.deleteDebugBundle(ctx, req)
}

func (s *DebugBundleServiceGatewayServer) DeleteDebugBundleFile(ctx context.Context, req *v1alpha1.DeleteDebugBundleFileRequest) (*v1alpha1.DeleteDebugBundleFileResponse, error) {
	return s.deleteDebugBundleFile(ctx, req)
}

// RegisterDebugBundleServiceHandlerGatewayServer registers the Connect handlers for the
// DebugBundleService "svc" to "mux".
func RegisterDebugBundleServiceHandlerGatewayServer(mux *runtime.ServeMux, svc DebugBundleServiceHandler, opts ...connect_gateway.HandlerOption) {
	if err := v1alpha1.RegisterDebugBundleServiceHandlerServer(context.TODO(), mux, NewDebugBundleServiceGatewayServer(svc, opts...)); err != nil {
		panic(fmt.Errorf("connect-gateway: %w", err))
	}
}