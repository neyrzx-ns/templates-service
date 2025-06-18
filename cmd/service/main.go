package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"templates-service/internal/presentation/rpc"
	v1 "templates-service/pkg/api/v1"
)

type TemplateChannelType uint8

const (
	TEMPLATE_CHANNELTYPE_UNSPECIFIED TemplateChannelType = iota
	TEMPLATE_CHANNELTYPE_IN_APP
	TEMPLATE_CHANNELTYPE_MESSENGER
	TEMPLATE_CHANNELTYPE_EMAIL
	TEMPLATE_CHANNELTYPE_SMS
)

type createTemplateDTO struct {
	Name                string              `json:"name,omitempty"`
	Description         string              `json:"description,omitempty"`
	Text                string              `json:"text,omitempty"`
	TemplateChannelType TemplateChannelType `json:"template_channel_type,omitempty"`
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	v1.RegisterTemplateServiceServer(grpcServer, &rpc.Service{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	//
	//dto := createTemplateDTO{
	//	Name:                "Информационное уведомление (PUSH)",
	//	Description:         "Информирование о внутреннем событии.",
	//	Text:                "Вам уведомление от {{.today}} о событии: {{.event_name}}",
	//	TemplateChannelType: TEMPLATE_CHANNELTYPE_IN_APP,
	//}
	//
	//dto2 := createTemplateDTO{
	//	Name:                "Info notification (PUSH)",
	//	Description:         "Информирование о внутреннем событии.",
	//	Text:                "Вам уведомление от {{.today}} о событии: {{.event_name}}",
	//	TemplateChannelType: TEMPLATE_CHANNELTYPE_IN_APP,
	//}
	//
	//templateBuilder := entity.NewTemplateBuilder[valueobjects.PushNotificationSpecification]()
	//template, err := templateBuilder.
	//	SetName(dto.Name).
	//	SetDescription(dto.Description).
	//	SetText(dto.Text).
	//	Build()
	//
	//if err != nil {
	//	fmt.Println(err.Error())
	//} else {
	//	fmt.Println("template1:", template)
	//}
	//
	//templateBuilder2 := entity.NewTemplateBuilder[valueobjects.PushNotificationSpecification]()
	//template2, err2 := templateBuilder2.
	//	SetName(dto2.Name).
	//	SetDescription(dto2.Description).
	//	SetText(dto2.Text).
	//	Build()
	//
	//if err2 != nil {
	//	fmt.Println(err2.Error())
	//} else {
	//	fmt.Println("template2:", template2)
	//}
}
