// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"fs.io/asyncd/ent/enttask"
	"fs.io/asyncd/ent/enttaskhandler"
	"fs.io/asyncd/ent/predicate"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeEntTask        = "EntTask"
	TypeEntTaskHandler = "EntTaskHandler"
)

// EntTaskMutation represents an operation that mutates the EntTask nodes in the graph.
type EntTaskMutation struct {
	config
	op            Op
	typ           string
	id            *int
	handler       *string
	parameter     *string
	priority      *int
	addpriority   *int
	created_at    *time.Time
	updated_at    *time.Time
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*EntTask, error)
	predicates    []predicate.EntTask
}

var _ ent.Mutation = (*EntTaskMutation)(nil)

// enttaskOption allows management of the mutation configuration using functional options.
type enttaskOption func(*EntTaskMutation)

// newEntTaskMutation creates new mutation for the EntTask entity.
func newEntTaskMutation(c config, op Op, opts ...enttaskOption) *EntTaskMutation {
	m := &EntTaskMutation{
		config:        c,
		op:            op,
		typ:           TypeEntTask,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withEntTaskID sets the ID field of the mutation.
func withEntTaskID(id int) enttaskOption {
	return func(m *EntTaskMutation) {
		var (
			err   error
			once  sync.Once
			value *EntTask
		)
		m.oldValue = func(ctx context.Context) (*EntTask, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().EntTask.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withEntTask sets the old EntTask of the mutation.
func withEntTask(node *EntTask) enttaskOption {
	return func(m *EntTaskMutation) {
		m.oldValue = func(context.Context) (*EntTask, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m EntTaskMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m EntTaskMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *EntTaskMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *EntTaskMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().EntTask.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetHandler sets the "handler" field.
func (m *EntTaskMutation) SetHandler(s string) {
	m.handler = &s
}

// Handler returns the value of the "handler" field in the mutation.
func (m *EntTaskMutation) Handler() (r string, exists bool) {
	v := m.handler
	if v == nil {
		return
	}
	return *v, true
}

// OldHandler returns the old "handler" field's value of the EntTask entity.
// If the EntTask object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EntTaskMutation) OldHandler(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldHandler is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldHandler requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldHandler: %w", err)
	}
	return oldValue.Handler, nil
}

// ResetHandler resets all changes to the "handler" field.
func (m *EntTaskMutation) ResetHandler() {
	m.handler = nil
}

// SetParameter sets the "parameter" field.
func (m *EntTaskMutation) SetParameter(s string) {
	m.parameter = &s
}

// Parameter returns the value of the "parameter" field in the mutation.
func (m *EntTaskMutation) Parameter() (r string, exists bool) {
	v := m.parameter
	if v == nil {
		return
	}
	return *v, true
}

// OldParameter returns the old "parameter" field's value of the EntTask entity.
// If the EntTask object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EntTaskMutation) OldParameter(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldParameter is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldParameter requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldParameter: %w", err)
	}
	return oldValue.Parameter, nil
}

// ResetParameter resets all changes to the "parameter" field.
func (m *EntTaskMutation) ResetParameter() {
	m.parameter = nil
}

// SetPriority sets the "priority" field.
func (m *EntTaskMutation) SetPriority(i int) {
	m.priority = &i
	m.addpriority = nil
}

// Priority returns the value of the "priority" field in the mutation.
func (m *EntTaskMutation) Priority() (r int, exists bool) {
	v := m.priority
	if v == nil {
		return
	}
	return *v, true
}

// OldPriority returns the old "priority" field's value of the EntTask entity.
// If the EntTask object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EntTaskMutation) OldPriority(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPriority is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPriority requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPriority: %w", err)
	}
	return oldValue.Priority, nil
}

// AddPriority adds i to the "priority" field.
func (m *EntTaskMutation) AddPriority(i int) {
	if m.addpriority != nil {
		*m.addpriority += i
	} else {
		m.addpriority = &i
	}
}

// AddedPriority returns the value that was added to the "priority" field in this mutation.
func (m *EntTaskMutation) AddedPriority() (r int, exists bool) {
	v := m.addpriority
	if v == nil {
		return
	}
	return *v, true
}

// ResetPriority resets all changes to the "priority" field.
func (m *EntTaskMutation) ResetPriority() {
	m.priority = nil
	m.addpriority = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *EntTaskMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *EntTaskMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the EntTask entity.
// If the EntTask object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EntTaskMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *EntTaskMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *EntTaskMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *EntTaskMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the EntTask entity.
// If the EntTask object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EntTaskMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *EntTaskMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// Where appends a list predicates to the EntTaskMutation builder.
func (m *EntTaskMutation) Where(ps ...predicate.EntTask) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the EntTaskMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *EntTaskMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.EntTask, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *EntTaskMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *EntTaskMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (EntTask).
func (m *EntTaskMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *EntTaskMutation) Fields() []string {
	fields := make([]string, 0, 5)
	if m.handler != nil {
		fields = append(fields, enttask.FieldHandler)
	}
	if m.parameter != nil {
		fields = append(fields, enttask.FieldParameter)
	}
	if m.priority != nil {
		fields = append(fields, enttask.FieldPriority)
	}
	if m.created_at != nil {
		fields = append(fields, enttask.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, enttask.FieldUpdatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *EntTaskMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case enttask.FieldHandler:
		return m.Handler()
	case enttask.FieldParameter:
		return m.Parameter()
	case enttask.FieldPriority:
		return m.Priority()
	case enttask.FieldCreatedAt:
		return m.CreatedAt()
	case enttask.FieldUpdatedAt:
		return m.UpdatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *EntTaskMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case enttask.FieldHandler:
		return m.OldHandler(ctx)
	case enttask.FieldParameter:
		return m.OldParameter(ctx)
	case enttask.FieldPriority:
		return m.OldPriority(ctx)
	case enttask.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case enttask.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown EntTask field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *EntTaskMutation) SetField(name string, value ent.Value) error {
	switch name {
	case enttask.FieldHandler:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetHandler(v)
		return nil
	case enttask.FieldParameter:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetParameter(v)
		return nil
	case enttask.FieldPriority:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPriority(v)
		return nil
	case enttask.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case enttask.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown EntTask field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *EntTaskMutation) AddedFields() []string {
	var fields []string
	if m.addpriority != nil {
		fields = append(fields, enttask.FieldPriority)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *EntTaskMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case enttask.FieldPriority:
		return m.AddedPriority()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *EntTaskMutation) AddField(name string, value ent.Value) error {
	switch name {
	case enttask.FieldPriority:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddPriority(v)
		return nil
	}
	return fmt.Errorf("unknown EntTask numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *EntTaskMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *EntTaskMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *EntTaskMutation) ClearField(name string) error {
	return fmt.Errorf("unknown EntTask nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *EntTaskMutation) ResetField(name string) error {
	switch name {
	case enttask.FieldHandler:
		m.ResetHandler()
		return nil
	case enttask.FieldParameter:
		m.ResetParameter()
		return nil
	case enttask.FieldPriority:
		m.ResetPriority()
		return nil
	case enttask.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case enttask.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	}
	return fmt.Errorf("unknown EntTask field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *EntTaskMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *EntTaskMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *EntTaskMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *EntTaskMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *EntTaskMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *EntTaskMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *EntTaskMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown EntTask unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *EntTaskMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown EntTask edge %s", name)
}

// EntTaskHandlerMutation represents an operation that mutates the EntTaskHandler nodes in the graph.
type EntTaskHandlerMutation struct {
	config
	op            Op
	typ           string
	id            *int
	name          *string
	signature     *string
	created_at    *time.Time
	updated_at    *time.Time
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*EntTaskHandler, error)
	predicates    []predicate.EntTaskHandler
}

var _ ent.Mutation = (*EntTaskHandlerMutation)(nil)

// enttaskhandlerOption allows management of the mutation configuration using functional options.
type enttaskhandlerOption func(*EntTaskHandlerMutation)

// newEntTaskHandlerMutation creates new mutation for the EntTaskHandler entity.
func newEntTaskHandlerMutation(c config, op Op, opts ...enttaskhandlerOption) *EntTaskHandlerMutation {
	m := &EntTaskHandlerMutation{
		config:        c,
		op:            op,
		typ:           TypeEntTaskHandler,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withEntTaskHandlerID sets the ID field of the mutation.
func withEntTaskHandlerID(id int) enttaskhandlerOption {
	return func(m *EntTaskHandlerMutation) {
		var (
			err   error
			once  sync.Once
			value *EntTaskHandler
		)
		m.oldValue = func(ctx context.Context) (*EntTaskHandler, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().EntTaskHandler.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withEntTaskHandler sets the old EntTaskHandler of the mutation.
func withEntTaskHandler(node *EntTaskHandler) enttaskhandlerOption {
	return func(m *EntTaskHandlerMutation) {
		m.oldValue = func(context.Context) (*EntTaskHandler, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m EntTaskHandlerMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m EntTaskHandlerMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *EntTaskHandlerMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *EntTaskHandlerMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().EntTaskHandler.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetName sets the "name" field.
func (m *EntTaskHandlerMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *EntTaskHandlerMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the EntTaskHandler entity.
// If the EntTaskHandler object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EntTaskHandlerMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *EntTaskHandlerMutation) ResetName() {
	m.name = nil
}

// SetSignature sets the "signature" field.
func (m *EntTaskHandlerMutation) SetSignature(s string) {
	m.signature = &s
}

// Signature returns the value of the "signature" field in the mutation.
func (m *EntTaskHandlerMutation) Signature() (r string, exists bool) {
	v := m.signature
	if v == nil {
		return
	}
	return *v, true
}

// OldSignature returns the old "signature" field's value of the EntTaskHandler entity.
// If the EntTaskHandler object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EntTaskHandlerMutation) OldSignature(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldSignature is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldSignature requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldSignature: %w", err)
	}
	return oldValue.Signature, nil
}

// ResetSignature resets all changes to the "signature" field.
func (m *EntTaskHandlerMutation) ResetSignature() {
	m.signature = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *EntTaskHandlerMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *EntTaskHandlerMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the EntTaskHandler entity.
// If the EntTaskHandler object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EntTaskHandlerMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *EntTaskHandlerMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *EntTaskHandlerMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *EntTaskHandlerMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the EntTaskHandler entity.
// If the EntTaskHandler object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *EntTaskHandlerMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *EntTaskHandlerMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// Where appends a list predicates to the EntTaskHandlerMutation builder.
func (m *EntTaskHandlerMutation) Where(ps ...predicate.EntTaskHandler) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the EntTaskHandlerMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *EntTaskHandlerMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.EntTaskHandler, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *EntTaskHandlerMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *EntTaskHandlerMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (EntTaskHandler).
func (m *EntTaskHandlerMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *EntTaskHandlerMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m.name != nil {
		fields = append(fields, enttaskhandler.FieldName)
	}
	if m.signature != nil {
		fields = append(fields, enttaskhandler.FieldSignature)
	}
	if m.created_at != nil {
		fields = append(fields, enttaskhandler.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, enttaskhandler.FieldUpdatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *EntTaskHandlerMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case enttaskhandler.FieldName:
		return m.Name()
	case enttaskhandler.FieldSignature:
		return m.Signature()
	case enttaskhandler.FieldCreatedAt:
		return m.CreatedAt()
	case enttaskhandler.FieldUpdatedAt:
		return m.UpdatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *EntTaskHandlerMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case enttaskhandler.FieldName:
		return m.OldName(ctx)
	case enttaskhandler.FieldSignature:
		return m.OldSignature(ctx)
	case enttaskhandler.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case enttaskhandler.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown EntTaskHandler field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *EntTaskHandlerMutation) SetField(name string, value ent.Value) error {
	switch name {
	case enttaskhandler.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case enttaskhandler.FieldSignature:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetSignature(v)
		return nil
	case enttaskhandler.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case enttaskhandler.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown EntTaskHandler field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *EntTaskHandlerMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *EntTaskHandlerMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *EntTaskHandlerMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown EntTaskHandler numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *EntTaskHandlerMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *EntTaskHandlerMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *EntTaskHandlerMutation) ClearField(name string) error {
	return fmt.Errorf("unknown EntTaskHandler nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *EntTaskHandlerMutation) ResetField(name string) error {
	switch name {
	case enttaskhandler.FieldName:
		m.ResetName()
		return nil
	case enttaskhandler.FieldSignature:
		m.ResetSignature()
		return nil
	case enttaskhandler.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case enttaskhandler.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	}
	return fmt.Errorf("unknown EntTaskHandler field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *EntTaskHandlerMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *EntTaskHandlerMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *EntTaskHandlerMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *EntTaskHandlerMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *EntTaskHandlerMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *EntTaskHandlerMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *EntTaskHandlerMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown EntTaskHandler unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *EntTaskHandlerMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown EntTaskHandler edge %s", name)
}
