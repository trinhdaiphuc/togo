// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/trinhdaiphuc/togo/database/ent/predicate"
	"github.com/trinhdaiphuc/togo/database/ent/task"
	"github.com/trinhdaiphuc/togo/database/ent/user"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetUsername sets the "username" field.
func (uu *UserUpdate) SetUsername(s string) *UserUpdate {
	uu.mutation.SetUsername(s)
	return uu
}

// SetPassword sets the "password" field.
func (uu *UserUpdate) SetPassword(s string) *UserUpdate {
	uu.mutation.SetPassword(s)
	return uu
}

// SetTaskLimit sets the "task_limit" field.
func (uu *UserUpdate) SetTaskLimit(i int) *UserUpdate {
	uu.mutation.ResetTaskLimit()
	uu.mutation.SetTaskLimit(i)
	return uu
}

// SetNillableTaskLimit sets the "task_limit" field if the given value is not nil.
func (uu *UserUpdate) SetNillableTaskLimit(i *int) *UserUpdate {
	if i != nil {
		uu.SetTaskLimit(*i)
	}
	return uu
}

// AddTaskLimit adds i to the "task_limit" field.
func (uu *UserUpdate) AddTaskLimit(i int) *UserUpdate {
	uu.mutation.AddTaskLimit(i)
	return uu
}

// AddUserTaskIDs adds the "user_task" edge to the Task entity by IDs.
func (uu *UserUpdate) AddUserTaskIDs(ids ...int) *UserUpdate {
	uu.mutation.AddUserTaskIDs(ids...)
	return uu
}

// AddUserTask adds the "user_task" edges to the Task entity.
func (uu *UserUpdate) AddUserTask(t ...*Task) *UserUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uu.AddUserTaskIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearUserTask clears all "user_task" edges to the Task entity.
func (uu *UserUpdate) ClearUserTask() *UserUpdate {
	uu.mutation.ClearUserTask()
	return uu
}

// RemoveUserTaskIDs removes the "user_task" edge to Task entities by IDs.
func (uu *UserUpdate) RemoveUserTaskIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveUserTaskIDs(ids...)
	return uu
}

// RemoveUserTask removes "user_task" edges to Task entities.
func (uu *UserUpdate) RemoveUserTask(t ...*Task) *UserUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uu.RemoveUserTaskIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, UserMutation](ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.TaskLimit(); ok {
		if err := user.TaskLimitValidator(v); err != nil {
			return &ValidationError{Name: "task_limit", err: fmt.Errorf(`ent: validator failed for field "User.task_limit": %w`, err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
	}
	if value, ok := uu.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if value, ok := uu.mutation.TaskLimit(); ok {
		_spec.SetField(user.FieldTaskLimit, field.TypeInt, value)
	}
	if value, ok := uu.mutation.AddedTaskLimit(); ok {
		_spec.AddField(user.FieldTaskLimit, field.TypeInt, value)
	}
	if uu.mutation.UserTaskCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserTaskTable,
			Columns: []string{user.UserTaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: task.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedUserTaskIDs(); len(nodes) > 0 && !uu.mutation.UserTaskCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserTaskTable,
			Columns: []string{user.UserTaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: task.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.UserTaskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserTaskTable,
			Columns: []string{user.UserTaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: task.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetUsername sets the "username" field.
func (uuo *UserUpdateOne) SetUsername(s string) *UserUpdateOne {
	uuo.mutation.SetUsername(s)
	return uuo
}

// SetPassword sets the "password" field.
func (uuo *UserUpdateOne) SetPassword(s string) *UserUpdateOne {
	uuo.mutation.SetPassword(s)
	return uuo
}

// SetTaskLimit sets the "task_limit" field.
func (uuo *UserUpdateOne) SetTaskLimit(i int) *UserUpdateOne {
	uuo.mutation.ResetTaskLimit()
	uuo.mutation.SetTaskLimit(i)
	return uuo
}

// SetNillableTaskLimit sets the "task_limit" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableTaskLimit(i *int) *UserUpdateOne {
	if i != nil {
		uuo.SetTaskLimit(*i)
	}
	return uuo
}

// AddTaskLimit adds i to the "task_limit" field.
func (uuo *UserUpdateOne) AddTaskLimit(i int) *UserUpdateOne {
	uuo.mutation.AddTaskLimit(i)
	return uuo
}

// AddUserTaskIDs adds the "user_task" edge to the Task entity by IDs.
func (uuo *UserUpdateOne) AddUserTaskIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddUserTaskIDs(ids...)
	return uuo
}

// AddUserTask adds the "user_task" edges to the Task entity.
func (uuo *UserUpdateOne) AddUserTask(t ...*Task) *UserUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uuo.AddUserTaskIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearUserTask clears all "user_task" edges to the Task entity.
func (uuo *UserUpdateOne) ClearUserTask() *UserUpdateOne {
	uuo.mutation.ClearUserTask()
	return uuo
}

// RemoveUserTaskIDs removes the "user_task" edge to Task entities by IDs.
func (uuo *UserUpdateOne) RemoveUserTaskIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveUserTaskIDs(ids...)
	return uuo
}

// RemoveUserTask removes "user_task" edges to Task entities.
func (uuo *UserUpdateOne) RemoveUserTask(t ...*Task) *UserUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uuo.RemoveUserTaskIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	return withHooks[*User, UserMutation](ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.TaskLimit(); ok {
		if err := user.TaskLimitValidator(v); err != nil {
			return &ValidationError{Name: "task_limit", err: fmt.Errorf(`ent: validator failed for field "User.task_limit": %w`, err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if value, ok := uuo.mutation.TaskLimit(); ok {
		_spec.SetField(user.FieldTaskLimit, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.AddedTaskLimit(); ok {
		_spec.AddField(user.FieldTaskLimit, field.TypeInt, value)
	}
	if uuo.mutation.UserTaskCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserTaskTable,
			Columns: []string{user.UserTaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: task.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedUserTaskIDs(); len(nodes) > 0 && !uuo.mutation.UserTaskCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserTaskTable,
			Columns: []string{user.UserTaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: task.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.UserTaskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserTaskTable,
			Columns: []string{user.UserTaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: task.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
