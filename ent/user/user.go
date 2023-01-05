// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDisplayName holds the string denoting the display_name field in the database.
	FieldDisplayName = "display_name"
	// FieldAdmin holds the string denoting the admin field in the database.
	FieldAdmin = "admin"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// EdgeGroupUser holds the string denoting the group_user edge name in mutations.
	EdgeGroupUser = "group_user"
	// EdgeGroupOwner holds the string denoting the group_owner edge name in mutations.
	EdgeGroupOwner = "group_owner"
	// EdgeComment holds the string denoting the comment edge name in mutations.
	EdgeComment = "comment"
	// EdgeRequestStatus holds the string denoting the request_status edge name in mutations.
	EdgeRequestStatus = "request_status"
	// EdgeRequest holds the string denoting the request edge name in mutations.
	EdgeRequest = "request"
	// EdgeFile holds the string denoting the file edge name in mutations.
	EdgeFile = "file"
	// EdgeRequestTarget holds the string denoting the request_target edge name in mutations.
	EdgeRequestTarget = "request_target"
	// Table holds the table name of the user in the database.
	Table = "users"
	// GroupUserTable is the table that holds the group_user relation/edge. The primary key declared below.
	GroupUserTable = "group_user"
	// GroupUserInverseTable is the table name for the Group entity.
	// It exists in this package in order to avoid circular dependency with the "group" package.
	GroupUserInverseTable = "groups"
	// GroupOwnerTable is the table that holds the group_owner relation/edge. The primary key declared below.
	GroupOwnerTable = "group_owner"
	// GroupOwnerInverseTable is the table name for the Group entity.
	// It exists in this package in order to avoid circular dependency with the "group" package.
	GroupOwnerInverseTable = "groups"
	// CommentTable is the table that holds the comment relation/edge.
	CommentTable = "comments"
	// CommentInverseTable is the table name for the Comment entity.
	// It exists in this package in order to avoid circular dependency with the "comment" package.
	CommentInverseTable = "comments"
	// CommentColumn is the table column denoting the comment relation/edge.
	CommentColumn = "comment_user"
	// RequestStatusTable is the table that holds the request_status relation/edge.
	RequestStatusTable = "request_status"
	// RequestStatusInverseTable is the table name for the RequestStatus entity.
	// It exists in this package in order to avoid circular dependency with the "requeststatus" package.
	RequestStatusInverseTable = "request_status"
	// RequestStatusColumn is the table column denoting the request_status relation/edge.
	RequestStatusColumn = "request_status_user"
	// RequestTable is the table that holds the request relation/edge.
	RequestTable = "requests"
	// RequestInverseTable is the table name for the Request entity.
	// It exists in this package in order to avoid circular dependency with the "request" package.
	RequestInverseTable = "requests"
	// RequestColumn is the table column denoting the request relation/edge.
	RequestColumn = "request_user"
	// FileTable is the table that holds the file relation/edge.
	FileTable = "files"
	// FileInverseTable is the table name for the File entity.
	// It exists in this package in order to avoid circular dependency with the "file" package.
	FileInverseTable = "files"
	// FileColumn is the table column denoting the file relation/edge.
	FileColumn = "file_user"
	// RequestTargetTable is the table that holds the request_target relation/edge.
	RequestTargetTable = "request_targets"
	// RequestTargetInverseTable is the table name for the RequestTarget entity.
	// It exists in this package in order to avoid circular dependency with the "requesttarget" package.
	RequestTargetInverseTable = "request_targets"
	// RequestTargetColumn is the table column denoting the request_target relation/edge.
	RequestTargetColumn = "request_target_user"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDisplayName,
	FieldAdmin,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
}

var (
	// GroupUserPrimaryKey and GroupUserColumn2 are the table columns denoting the
	// primary key for the group_user relation (M2M).
	GroupUserPrimaryKey = []string{"group_id", "user_id"}
	// GroupOwnerPrimaryKey and GroupOwnerColumn2 are the table columns denoting the
	// primary key for the group_owner relation (M2M).
	GroupOwnerPrimaryKey = []string{"group_id", "user_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultAdmin holds the default value on creation for the "admin" field.
	DefaultAdmin bool
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
