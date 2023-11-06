// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by aliasgen. DO NOT EDIT.

// Package iam aliases all exported identifiers in package
// "cloud.google.com/go/iam/apiv1/iampb".
//
// Deprecated: Please use types in: cloud.google.com/go/iam/apiv1/iampb.
// Please read https://github.com/googleapis/google-cloud-go/blob/main/migration.md
// for more details.
package iam

import (
	src "cloud.google.com/go/iam/apiv1/iampb"
	grpc "google.golang.org/grpc"
)

// Deprecated: Please use consts in: cloud.google.com/go/iam/apiv1/iampb
const (
	AuditConfigDelta_ACTION_UNSPECIFIED = src.AuditConfigDelta_ACTION_UNSPECIFIED
	AuditConfigDelta_ADD                = src.AuditConfigDelta_ADD
	AuditConfigDelta_REMOVE             = src.AuditConfigDelta_REMOVE
	AuditLogConfig_ADMIN_READ           = src.AuditLogConfig_ADMIN_READ
	AuditLogConfig_DATA_READ            = src.AuditLogConfig_DATA_READ
	AuditLogConfig_DATA_WRITE           = src.AuditLogConfig_DATA_WRITE
	AuditLogConfig_LOG_TYPE_UNSPECIFIED = src.AuditLogConfig_LOG_TYPE_UNSPECIFIED
	BindingDelta_ACTION_UNSPECIFIED     = src.BindingDelta_ACTION_UNSPECIFIED
	BindingDelta_ADD                    = src.BindingDelta_ADD
	BindingDelta_REMOVE                 = src.BindingDelta_REMOVE
)

// Deprecated: Please use vars in: cloud.google.com/go/iam/apiv1/iampb
var (
	AuditConfigDelta_Action_name        = src.AuditConfigDelta_Action_name
	AuditConfigDelta_Action_value       = src.AuditConfigDelta_Action_value
	AuditLogConfig_LogType_name         = src.AuditLogConfig_LogType_name
	AuditLogConfig_LogType_value        = src.AuditLogConfig_LogType_value
	BindingDelta_Action_name            = src.BindingDelta_Action_name
	BindingDelta_Action_value           = src.BindingDelta_Action_value
	File_google_iam_v1_iam_policy_proto = src.File_google_iam_v1_iam_policy_proto
	File_google_iam_v1_options_proto    = src.File_google_iam_v1_options_proto
	File_google_iam_v1_policy_proto     = src.File_google_iam_v1_policy_proto
)

// Specifies the audit configuration for a service. The configuration
// determines which permission types are logged, and what identities, if any,
// are exempted from logging. An AuditConfig must have one or more
// AuditLogConfigs. If there are AuditConfigs for both `allServices` and a
// specific service, the union of the two AuditConfigs is used for that
// service: the log_types specified in each AuditConfig are enabled, and the
// exempted_members in each AuditLogConfig are exempted. Example Policy with
// multiple AuditConfigs: { "audit_configs": [ { "service": "allServices",
// "audit_log_configs": [ { "log_type": "DATA_READ", "exempted_members": [
// "user:jose@example.com" ] }, { "log_type": "DATA_WRITE" }, { "log_type":
// "ADMIN_READ" } ] }, { "service": "sampleservice.googleapis.com",
// "audit_log_configs": [ { "log_type": "DATA_READ" }, { "log_type":
// "DATA_WRITE", "exempted_members": [ "user:aliya@example.com" ] } ] } ] } For
// sampleservice, this policy enables DATA_READ, DATA_WRITE and ADMIN_READ
// logging. It also exempts jose@example.com from DATA_READ logging, and
// aliya@example.com from DATA_WRITE logging.
//
// Deprecated: Please use types in: cloud.google.com/go/iam/apiv1/iampb
type AuditConfig = src.AuditConfig

// One delta entry for AuditConfig. Each individual change (only one
// exempted_member in each entry) to a AuditConfig will be a separate entry.
//
// Deprecated: Please use types in: cloud.google.com/go/iam/apiv1/iampb
type AuditConfigDelta = src.AuditConfigDelta

// The type of action performed on an audit configuration in a policy.
//
// Deprecated: Please use types in: cloud.google.com/go/iam/apiv1/iampb
type AuditConfigDelta_Action = src.AuditConfigDelta_Action

// Provides the configuration for logging a type of permissions. Example: {
// "audit_log_configs": [ { "log_type": "DATA_READ", "exempted_members": [
// "user:jose@example.com" ] }, { "log_type": "DATA_WRITE" } ] } This enables
// 'DATA_READ' and 'DATA_WRITE' logging, while exempting jose@example.com from
// DATA_READ logging.
//
// Deprecated: Please use types in: cloud.google.com/go/iam/apiv1/iampb
type AuditLogConfig = src.AuditLogConfig

// The list of valid permission types for which logging can be configured.
// Admin writes are always logged, and are not configurable.
//
// Deprecated: Please use types in: cloud.google.com/go/iam/apiv1/iampb
type AuditLogConfig_LogType = src.AuditLogConfig_LogType

// Associates `members`, or principals, with a `role`.
//
// Deprecated: Please use types in: cloud.google.com/go/iam/apiv1/iampb
type Binding = src.Binding

// One delta entry for Binding. Each individual change (only one member in
// each entry) to a binding will be a separate entry.
//
// Deprecated: Please use types in: cloud.google.com/go/iam/apiv1/iampb
type BindingDelta = src.BindingDelta

// The type of action performed on a Binding in a policy.
//
// Deprecated: Please use types in: cloud.google.com/go/iam/apiv1/iampb
type BindingDelta_Action = src.BindingDelta_Action

// Request message for `GetIamPolicy` method.
//
// Deprecated: Please use types in: cloud.google.com/go/iam/apiv1/iampb
type GetIamPolicyRequest = src.GetIamPolicyRequest

// Encapsulates settings provided to GetIamPolicy.
//
// Deprecated: Please use types in: cloud.google.com/go/iam/apiv1/iampb
type GetPolicyOptions = src.GetPolicyOptions

// IAMPolicyClient is the client API for IAMPolicy service. For semantics
// around ctx use and closing/ending streaming RPCs, please refer to
// https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
//
// Deprecated: Please use types in: cloud.google.com/go/iam/apiv1/iampb
type IAMPolicyClient = src.IAMPolicyClient

// IAMPolicyServer is the server API for IAMPolicy service.
//
// Deprecated: Please use types in: cloud.google.com/go/iam/apiv1/iampb
type IAMPolicyServer = src.IAMPolicyServer

// An Identity and Access Management (IAM) policy, which specifies access
// controls for Google Cloud resources. A `Policy` is a collection of
// `bindings`. A `binding` binds one or more `members`, or principals, to a
// single `role`. Principals can be user accounts, service accounts, Google
// groups, and domains (such as G Suite). A `role` is a named list of
// permissions; each `role` can be an IAM predefined role or a user-created
// custom role. For some types of Google Cloud resources, a `binding` can also
// specify a `condition`, which is a logical expression that allows access to a
// resource only if the expression evaluates to `true`. A condition can add
// constraints based on attributes of the request, the resource, or both. To
// learn which resources support conditions in their IAM policies, see the [IAM
// documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
// **JSON example:** { "bindings": [ { "role":
// "roles/resourcemanager.organizationAdmin", "members": [
// "user:mike@example.com", "group:admins@example.com", "domain:google.com",
// "serviceAccount:my-project-id@appspot.gserviceaccount.com" ] }, { "role":
// "roles/resourcemanager.organizationViewer", "members": [
// "user:eve@example.com" ], "condition": { "title": "expirable access",
// "description": "Does not grant access after Sep 2020", "expression":
// "request.time < timestamp('2020-10-01T00:00:00.000Z')", } } ], "etag":
// "BwWWja0YfJA=", "version": 3 } **YAML example:** bindings: - members: -
// user:mike@example.com - group:admins@example.com - domain:google.com -
// serviceAccount:my-project-id@appspot.gserviceaccount.com role:
// roles/resourcemanager.organizationAdmin - members: - user:eve@example.com
// role: roles/resourcemanager.organizationViewer condition: title: expirable
// access description: Does not grant access after Sep 2020 expression:
// request.time < timestamp('2020-10-01T00:00:00.000Z') etag: BwWWja0YfJA=
// version: 3 For a description of IAM and its features, see the [IAM
// documentation](https://cloud.google.com/iam/docs/).
//
// Deprecated: Please use types in: cloud.google.com/go/iam/apiv1/iampb
type Policy = src.Policy

// The difference delta between two policies.
//
// Deprecated: Please use types in: cloud.google.com/go/iam/apiv1/iampb
type PolicyDelta = src.PolicyDelta

// Request message for `SetIamPolicy` method.
//
// Deprecated: Please use types in: cloud.google.com/go/iam/apiv1/iampb
type SetIamPolicyRequest = src.SetIamPolicyRequest

// Request message for `TestIamPermissions` method.
//
// Deprecated: Please use types in: cloud.google.com/go/iam/apiv1/iampb
type TestIamPermissionsRequest = src.TestIamPermissionsRequest

// Response message for `TestIamPermissions` method.
//
// Deprecated: Please use types in: cloud.google.com/go/iam/apiv1/iampb
type TestIamPermissionsResponse = src.TestIamPermissionsResponse

// UnimplementedIAMPolicyServer can be embedded to have forward compatible
// implementations.
//
// Deprecated: Please use types in: cloud.google.com/go/iam/apiv1/iampb
type UnimplementedIAMPolicyServer = src.UnimplementedIAMPolicyServer

// Deprecated: Please use funcs in: cloud.google.com/go/iam/apiv1/iampb
func NewIAMPolicyClient(cc grpc.ClientConnInterface) IAMPolicyClient {
	return src.NewIAMPolicyClient(cc)
}

// Deprecated: Please use funcs in: cloud.google.com/go/iam/apiv1/iampb
func RegisterIAMPolicyServer(s *grpc.Server, srv IAMPolicyServer) {
	src.RegisterIAMPolicyServer(s, srv)
}
