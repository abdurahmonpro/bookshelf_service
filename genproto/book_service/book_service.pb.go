// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: book_service.proto

package book_service

import (
	empty "github.com/golang/protobuf/ptypes/empty"
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

var File_book_service_proto protoreflect.FileDescriptor

var file_book_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0a, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xc5, 0x03, 0x0a, 0x0b,
	0x42, 0x6f, 0x6f, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x1a,
	0x12, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42,
	0x6f, 0x6f, 0x6b, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44,
	0x12, 0x14, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x42, 0x6f, 0x6f, 0x6b, 0x50, 0x4b, 0x1a, 0x12, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x07,
	0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1d, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x18, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x1a, 0x12, 0x2e, 0x62,
	0x6f, 0x6f, 0x6b, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x6f, 0x6f, 0x6b,
	0x22, 0x00, 0x12, 0x42, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x74, 0x63,
	0x68, 0x12, 0x1d, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x74, 0x63, 0x68, 0x42, 0x6f, 0x6f, 0x6b,
	0x1a, 0x12, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x42, 0x6f, 0x6f, 0x6b, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x12, 0x14, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x42, 0x6f, 0x6f, 0x6b, 0x50, 0x4b, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x12, 0x41, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x42, 0x79, 0x54, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x19, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x42, 0x79, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x1a, 0x12, 0x2e,
	0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x6f, 0x6f,
	0x6b, 0x22, 0x00, 0x42, 0x17, 0x5a, 0x15, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var file_book_service_proto_goTypes = []interface{}{
	(*CreateBook)(nil),       // 0: book_service.CreateBook
	(*BookPK)(nil),           // 1: book_service.BookPK
	(*BookListRequest)(nil),  // 2: book_service.BookListRequest
	(*UpdateBook)(nil),       // 3: book_service.UpdateBook
	(*UpdatePatchBook)(nil),  // 4: book_service.UpdatePatchBook
	(*BookByTitle)(nil),      // 5: book_service.BookByTitle
	(*Book)(nil),             // 6: book_service.Book
	(*BookListResponse)(nil), // 7: book_service.BookListResponse
	(*empty.Empty)(nil),      // 8: google.protobuf.Empty
}
var file_book_service_proto_depIdxs = []int32{
	0, // 0: book_service.BookService.Create:input_type -> book_service.CreateBook
	1, // 1: book_service.BookService.GetByID:input_type -> book_service.BookPK
	2, // 2: book_service.BookService.GetList:input_type -> book_service.BookListRequest
	3, // 3: book_service.BookService.Update:input_type -> book_service.UpdateBook
	4, // 4: book_service.BookService.UpdatePatch:input_type -> book_service.UpdatePatchBook
	1, // 5: book_service.BookService.Delete:input_type -> book_service.BookPK
	5, // 6: book_service.BookService.GetBookByTitle:input_type -> book_service.BookByTitle
	6, // 7: book_service.BookService.Create:output_type -> book_service.Book
	6, // 8: book_service.BookService.GetByID:output_type -> book_service.Book
	7, // 9: book_service.BookService.GetList:output_type -> book_service.BookListResponse
	6, // 10: book_service.BookService.Update:output_type -> book_service.Book
	6, // 11: book_service.BookService.UpdatePatch:output_type -> book_service.Book
	8, // 12: book_service.BookService.Delete:output_type -> google.protobuf.Empty
	6, // 13: book_service.BookService.GetBookByTitle:output_type -> book_service.Book
	7, // [7:14] is the sub-list for method output_type
	0, // [0:7] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_book_service_proto_init() }
func file_book_service_proto_init() {
	if File_book_service_proto != nil {
		return
	}
	file_book_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_book_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_book_service_proto_goTypes,
		DependencyIndexes: file_book_service_proto_depIdxs,
	}.Build()
	File_book_service_proto = out.File
	file_book_service_proto_rawDesc = nil
	file_book_service_proto_goTypes = nil
	file_book_service_proto_depIdxs = nil
}
