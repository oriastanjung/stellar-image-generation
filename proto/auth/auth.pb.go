// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.1
// source: auth/auth.proto

package auth

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_auth_auth_proto protoreflect.FileDescriptor

var file_auth_auth_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x61, 0x75, 0x74, 0x68, 0x1a, 0x13, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x61, 0x75,
	0x74, 0x68, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13,
	0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x32, 0xd2, 0x05, 0x0a, 0x11, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x12, 0x42, 0x0a, 0x0b, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x17, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0a, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x13, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0a, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x53,
	0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x13, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x6c, 0x6f, 0x67,
	0x69, 0x6e, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x49, 0x0a, 0x0a, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x1b, 0x2e, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6a, 0x0a,
	0x15, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x46, 0x6f, 0x72, 0x67, 0x65, 0x74, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x26, 0x2e, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x46, 0x6f, 0x72, 0x67, 0x65, 0x74, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27,
	0x2e, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x46, 0x6f, 0x72, 0x67, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x67, 0x0a, 0x14, 0x52, 0x65, 0x73,
	0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x42, 0x79, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x25, 0x2e, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73,
	0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x42, 0x79, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x61, 0x64, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x42, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x4d, 0x0a, 0x12, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x56,
	0x69, 0x61, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x1d, 0x2e, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x52, 0x0a, 0x1a, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x56, 0x69,
	0x61, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x12,
	0x1c, 0x2e, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x72, 0x69, 0x61, 0x73, 0x74, 0x61, 0x6e, 0x6a, 0x75, 0x6e, 0x67,
	0x2f, 0x73, 0x74, 0x65, 0x6c, 0x6c, 0x61, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_auth_auth_proto_goTypes = []any{
	(*SignUpRequest)(nil),                 // 0: register.SignUpRequest
	(*LoginRequest)(nil),                  // 1: login.LoginRequest
	(*VerifyUserRequest)(nil),             // 2: addition.VerifyUserRequest
	(*RequestForgetPasswordRequest)(nil),  // 3: addition.RequestForgetPasswordRequest
	(*ResetPasswordByTokenRequest)(nil),   // 4: addition.ResetPasswordByTokenRequest
	(*emptypb.Empty)(nil),                 // 5: google.protobuf.Empty
	(*LoginGoogleRequest)(nil),            // 6: addition.LoginGoogleRequest
	(*SignUpResponse)(nil),                // 7: register.SignUpResponse
	(*LoginResponse)(nil),                 // 8: login.LoginResponse
	(*VerifyUserResponse)(nil),            // 9: addition.VerifyUserResponse
	(*RequestForgetPasswordResponse)(nil), // 10: addition.RequestForgetPasswordResponse
	(*ResetPasswordByTokenResponse)(nil),  // 11: addition.ResetPasswordByTokenResponse
	(*LoginGoogleResponse)(nil),           // 12: addition.LoginGoogleResponse
}
var file_auth_auth_proto_depIdxs = []int32{
	0,  // 0: auth.AuthServiceRoutes.SignUpAdmin:input_type -> register.SignUpRequest
	1,  // 1: auth.AuthServiceRoutes.LoginAdmin:input_type -> login.LoginRequest
	0,  // 2: auth.AuthServiceRoutes.SignUpUser:input_type -> register.SignUpRequest
	1,  // 3: auth.AuthServiceRoutes.LoginUser:input_type -> login.LoginRequest
	2,  // 4: auth.AuthServiceRoutes.VerifyUser:input_type -> addition.VerifyUserRequest
	3,  // 5: auth.AuthServiceRoutes.RequestForgetPassword:input_type -> addition.RequestForgetPasswordRequest
	4,  // 6: auth.AuthServiceRoutes.ResetPasswordByToken:input_type -> addition.ResetPasswordByTokenRequest
	5,  // 7: auth.AuthServiceRoutes.LoginUserViaGoogle:input_type -> google.protobuf.Empty
	6,  // 8: auth.AuthServiceRoutes.LoginUserViaGoogleCallback:input_type -> addition.LoginGoogleRequest
	7,  // 9: auth.AuthServiceRoutes.SignUpAdmin:output_type -> register.SignUpResponse
	8,  // 10: auth.AuthServiceRoutes.LoginAdmin:output_type -> login.LoginResponse
	7,  // 11: auth.AuthServiceRoutes.SignUpUser:output_type -> register.SignUpResponse
	8,  // 12: auth.AuthServiceRoutes.LoginUser:output_type -> login.LoginResponse
	9,  // 13: auth.AuthServiceRoutes.VerifyUser:output_type -> addition.VerifyUserResponse
	10, // 14: auth.AuthServiceRoutes.RequestForgetPassword:output_type -> addition.RequestForgetPasswordResponse
	11, // 15: auth.AuthServiceRoutes.ResetPasswordByToken:output_type -> addition.ResetPasswordByTokenResponse
	12, // 16: auth.AuthServiceRoutes.LoginUserViaGoogle:output_type -> addition.LoginGoogleResponse
	8,  // 17: auth.AuthServiceRoutes.LoginUserViaGoogleCallback:output_type -> login.LoginResponse
	9,  // [9:18] is the sub-list for method output_type
	0,  // [0:9] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_auth_auth_proto_init() }
func file_auth_auth_proto_init() {
	if File_auth_auth_proto != nil {
		return
	}
	file_auth_register_proto_init()
	file_auth_login_proto_init()
	file_auth_addition_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_auth_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_auth_proto_goTypes,
		DependencyIndexes: file_auth_auth_proto_depIdxs,
	}.Build()
	File_auth_auth_proto = out.File
	file_auth_auth_proto_rawDesc = nil
	file_auth_auth_proto_goTypes = nil
	file_auth_auth_proto_depIdxs = nil
}