// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/traPtitech/Jomon/ent/comment"
	"github.com/traPtitech/Jomon/ent/request"
	"github.com/traPtitech/Jomon/ent/requeststatus"
	"github.com/traPtitech/Jomon/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// DisplayName holds the value of the "display_name" field.
	DisplayName string `json:"display_name,omitempty"`
	// Admin holds the value of the "admin" field.
	Admin bool `json:"admin,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges               UserEdges `json:"edges"`
	comment_user        *uuid.UUID
	request_user        *uuid.UUID
	request_status_user *uuid.UUID
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Group holds the value of the group edge.
	Group []*Group `json:"group,omitempty"`
	// Comment holds the value of the comment edge.
	Comment *Comment `json:"comment,omitempty"`
	// RequestStatus holds the value of the request_status edge.
	RequestStatus *RequestStatus `json:"request_status,omitempty"`
	// Request holds the value of the request edge.
	Request *Request `json:"request,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// GroupOrErr returns the Group value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) GroupOrErr() ([]*Group, error) {
	if e.loadedTypes[0] {
		return e.Group, nil
	}
	return nil, &NotLoadedError{edge: "group"}
}

// CommentOrErr returns the Comment value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) CommentOrErr() (*Comment, error) {
	if e.loadedTypes[1] {
		if e.Comment == nil {
			// The edge comment was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: comment.Label}
		}
		return e.Comment, nil
	}
	return nil, &NotLoadedError{edge: "comment"}
}

// RequestStatusOrErr returns the RequestStatus value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) RequestStatusOrErr() (*RequestStatus, error) {
	if e.loadedTypes[2] {
		if e.RequestStatus == nil {
			// The edge request_status was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: requeststatus.Label}
		}
		return e.RequestStatus, nil
	}
	return nil, &NotLoadedError{edge: "request_status"}
}

// RequestOrErr returns the Request value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) RequestOrErr() (*Request, error) {
	if e.loadedTypes[3] {
		if e.Request == nil {
			// The edge request was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: request.Label}
		}
		return e.Request, nil
	}
	return nil, &NotLoadedError{edge: "request"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldAdmin:
			values[i] = new(sql.NullBool)
		case user.FieldName, user.FieldDisplayName:
			values[i] = new(sql.NullString)
		case user.FieldCreatedAt, user.FieldUpdatedAt, user.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case user.FieldID:
			values[i] = new(uuid.UUID)
		case user.ForeignKeys[0]: // comment_user
			values[i] = new(uuid.UUID)
		case user.ForeignKeys[1]: // request_user
			values[i] = new(uuid.UUID)
		case user.ForeignKeys[2]: // request_status_user
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type User", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				u.ID = *value
			}
		case user.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				u.Name = value.String
			}
		case user.FieldDisplayName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field display_name", values[i])
			} else if value.Valid {
				u.DisplayName = value.String
			}
		case user.FieldAdmin:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field admin", values[i])
			} else if value.Valid {
				u.Admin = value.Bool
			}
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		case user.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				u.UpdatedAt = value.Time
			}
		case user.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				u.DeletedAt = new(time.Time)
				*u.DeletedAt = value.Time
			}
		case user.ForeignKeys[0]:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field comment_user", values[i])
			} else if value != nil {
				u.comment_user = value
			}
		case user.ForeignKeys[1]:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field request_user", values[i])
			} else if value != nil {
				u.request_user = value
			}
		case user.ForeignKeys[2]:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field request_status_user", values[i])
			} else if value != nil {
				u.request_status_user = value
			}
		}
	}
	return nil
}

// QueryGroup queries the "group" edge of the User entity.
func (u *User) QueryGroup() *GroupQuery {
	return (&UserClient{config: u.config}).QueryGroup(u)
}

// QueryComment queries the "comment" edge of the User entity.
func (u *User) QueryComment() *CommentQuery {
	return (&UserClient{config: u.config}).QueryComment(u)
}

// QueryRequestStatus queries the "request_status" edge of the User entity.
func (u *User) QueryRequestStatus() *RequestStatusQuery {
	return (&UserClient{config: u.config}).QueryRequestStatus(u)
}

// QueryRequest queries the "request" edge of the User entity.
func (u *User) QueryRequest() *RequestQuery {
	return (&UserClient{config: u.config}).QueryRequest(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", name=")
	builder.WriteString(u.Name)
	builder.WriteString(", display_name=")
	builder.WriteString(u.DisplayName)
	builder.WriteString(", admin=")
	builder.WriteString(fmt.Sprintf("%v", u.Admin))
	builder.WriteString(", created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(u.UpdatedAt.Format(time.ANSIC))
	if v := u.DeletedAt; v != nil {
		builder.WriteString(", deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
