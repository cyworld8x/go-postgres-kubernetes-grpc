// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: pkg/pb/proto/rpc_ticket_service .proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_pkg_pb_proto_rpc_ticket_service__proto protoreflect.FileDescriptor

var file_pkg_pb_proto_rpc_ticket_service__proto_rawDesc = []byte{
	0x0a, 0x26, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72,
	0x70, 0x63, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x21, 0x70, 0x6b,
	0x67, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
	0xd8, 0x01, 0x0a, 0x0d, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x46, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x73, 0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70,
	0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0a, 0x53, 0x65, 0x6c,
	0x6c, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x6c,
	0x6c, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x6c, 0x6c, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x07, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x49, 0x6e, 0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e,
	0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e,
	0x70, 0x62, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_pkg_pb_proto_rpc_ticket_service__proto_goTypes = []interface{}{
	(*CreateTicketsRequest)(nil),  // 0: pb.CreateTicketsRequest
	(*SellTicketRequest)(nil),     // 1: pb.SellTicketRequest
	(*CheckinTicketRequest)(nil),  // 2: pb.CheckinTicketRequest
	(*CreateTicketsResponse)(nil), // 3: pb.CreateTicketsResponse
	(*SellTicketResponse)(nil),    // 4: pb.SellTicketResponse
	(*CheckinTicketResponse)(nil), // 5: pb.CheckinTicketResponse
}
var file_pkg_pb_proto_rpc_ticket_service__proto_depIdxs = []int32{
	0, // 0: pb.TicketService.CreateTickets:input_type -> pb.CreateTicketsRequest
	1, // 1: pb.TicketService.SellTicket:input_type -> pb.SellTicketRequest
	2, // 2: pb.TicketService.CheckIn:input_type -> pb.CheckinTicketRequest
	3, // 3: pb.TicketService.CreateTickets:output_type -> pb.CreateTicketsResponse
	4, // 4: pb.TicketService.SellTicket:output_type -> pb.SellTicketResponse
	5, // 5: pb.TicketService.CheckIn:output_type -> pb.CheckinTicketResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_pb_proto_rpc_ticket_service__proto_init() }
func file_pkg_pb_proto_rpc_ticket_service__proto_init() {
	if File_pkg_pb_proto_rpc_ticket_service__proto != nil {
		return
	}
	file_pkg_pb_proto_ticket_service_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_pb_proto_rpc_ticket_service__proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_pb_proto_rpc_ticket_service__proto_goTypes,
		DependencyIndexes: file_pkg_pb_proto_rpc_ticket_service__proto_depIdxs,
	}.Build()
	File_pkg_pb_proto_rpc_ticket_service__proto = out.File
	file_pkg_pb_proto_rpc_ticket_service__proto_rawDesc = nil
	file_pkg_pb_proto_rpc_ticket_service__proto_goTypes = nil
	file_pkg_pb_proto_rpc_ticket_service__proto_depIdxs = nil
}
