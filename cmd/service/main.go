package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	v1 "templates-service/pkg/api/v1"
)

var _ (v1.TemplateServiceServer) = (*Service)(nil)

type Service struct {
	v1.UnimplementedTemplateServiceServer
}

func (s Service) CreateTemplate(ctx context.Context, req *v1.CreateTemplateRequest) (*v1.CreateTemplateResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) CreateTemplates(ctx context.Context, req *v1.CreateTemplatesRequest) (*v1.CreateTemplatesResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) UpdateTemplate(ctx context.Context, req *v1.UpdateTemplateRequest) (*v1.UpdateTemplateResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) DeleteTemplate(ctx context.Context, req *v1.DeleteTemplateRequest) (*v1.DeleteTemplateResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) DeleteTemplates(ctx context.Context, req *v1.DeleteTemplatesRequest) (*v1.DeleteTemplatesResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) GetTemplate(ctx context.Context, req *v1.GetTemplateRequest) (*v1.GetTemplateResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) ListTemplates(ctx context.Context, req *v1.ListTemplatesRequest) (*v1.ListTemplatesResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) CreateTemplateGroup(ctx context.Context, req *v1.CreateTemplateGroupRequest) (*v1.CreateTemplateGroupResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) UpdateTemplateGroup(ctx context.Context, req *v1.UpdateTemplateGroupRequest) (*v1.UpdateTemplateGroupResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) DeleteTemplateGroup(ctx context.Context, req *v1.DeleteTemplateGroupRequest) (*v1.DeleteTemplateGroupResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) GetTemplateGroup(ctx context.Context, req *v1.GetTemplateGroupRequest) (*v1.GetTemplateGroupResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) ListTemplateGroups(ctx context.Context, req *v1.ListTemplateGroupsRequest) (*v1.ListTemplateGroupsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	v1.RegisterTemplateServiceServer(grpcServer, &Service{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
