// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"polaris/ent/history"
	"polaris/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// HistoryUpdate is the builder for updating History entities.
type HistoryUpdate struct {
	config
	hooks    []Hook
	mutation *HistoryMutation
}

// Where appends a list predicates to the HistoryUpdate builder.
func (hu *HistoryUpdate) Where(ps ...predicate.History) *HistoryUpdate {
	hu.mutation.Where(ps...)
	return hu
}

// SetMediaID sets the "media_id" field.
func (hu *HistoryUpdate) SetMediaID(i int) *HistoryUpdate {
	hu.mutation.ResetMediaID()
	hu.mutation.SetMediaID(i)
	return hu
}

// SetNillableMediaID sets the "media_id" field if the given value is not nil.
func (hu *HistoryUpdate) SetNillableMediaID(i *int) *HistoryUpdate {
	if i != nil {
		hu.SetMediaID(*i)
	}
	return hu
}

// AddMediaID adds i to the "media_id" field.
func (hu *HistoryUpdate) AddMediaID(i int) *HistoryUpdate {
	hu.mutation.AddMediaID(i)
	return hu
}

// SetEpisodeID sets the "episode_id" field.
func (hu *HistoryUpdate) SetEpisodeID(i int) *HistoryUpdate {
	hu.mutation.ResetEpisodeID()
	hu.mutation.SetEpisodeID(i)
	return hu
}

// SetNillableEpisodeID sets the "episode_id" field if the given value is not nil.
func (hu *HistoryUpdate) SetNillableEpisodeID(i *int) *HistoryUpdate {
	if i != nil {
		hu.SetEpisodeID(*i)
	}
	return hu
}

// AddEpisodeID adds i to the "episode_id" field.
func (hu *HistoryUpdate) AddEpisodeID(i int) *HistoryUpdate {
	hu.mutation.AddEpisodeID(i)
	return hu
}

// ClearEpisodeID clears the value of the "episode_id" field.
func (hu *HistoryUpdate) ClearEpisodeID() *HistoryUpdate {
	hu.mutation.ClearEpisodeID()
	return hu
}

// SetSourceTitle sets the "source_title" field.
func (hu *HistoryUpdate) SetSourceTitle(s string) *HistoryUpdate {
	hu.mutation.SetSourceTitle(s)
	return hu
}

// SetNillableSourceTitle sets the "source_title" field if the given value is not nil.
func (hu *HistoryUpdate) SetNillableSourceTitle(s *string) *HistoryUpdate {
	if s != nil {
		hu.SetSourceTitle(*s)
	}
	return hu
}

// SetDate sets the "date" field.
func (hu *HistoryUpdate) SetDate(t time.Time) *HistoryUpdate {
	hu.mutation.SetDate(t)
	return hu
}

// SetNillableDate sets the "date" field if the given value is not nil.
func (hu *HistoryUpdate) SetNillableDate(t *time.Time) *HistoryUpdate {
	if t != nil {
		hu.SetDate(*t)
	}
	return hu
}

// SetTargetDir sets the "target_dir" field.
func (hu *HistoryUpdate) SetTargetDir(s string) *HistoryUpdate {
	hu.mutation.SetTargetDir(s)
	return hu
}

// SetNillableTargetDir sets the "target_dir" field if the given value is not nil.
func (hu *HistoryUpdate) SetNillableTargetDir(s *string) *HistoryUpdate {
	if s != nil {
		hu.SetTargetDir(*s)
	}
	return hu
}

// SetSize sets the "size" field.
func (hu *HistoryUpdate) SetSize(i int) *HistoryUpdate {
	hu.mutation.ResetSize()
	hu.mutation.SetSize(i)
	return hu
}

// SetNillableSize sets the "size" field if the given value is not nil.
func (hu *HistoryUpdate) SetNillableSize(i *int) *HistoryUpdate {
	if i != nil {
		hu.SetSize(*i)
	}
	return hu
}

// AddSize adds i to the "size" field.
func (hu *HistoryUpdate) AddSize(i int) *HistoryUpdate {
	hu.mutation.AddSize(i)
	return hu
}

// SetDownloadClientID sets the "download_client_id" field.
func (hu *HistoryUpdate) SetDownloadClientID(i int) *HistoryUpdate {
	hu.mutation.ResetDownloadClientID()
	hu.mutation.SetDownloadClientID(i)
	return hu
}

// SetNillableDownloadClientID sets the "download_client_id" field if the given value is not nil.
func (hu *HistoryUpdate) SetNillableDownloadClientID(i *int) *HistoryUpdate {
	if i != nil {
		hu.SetDownloadClientID(*i)
	}
	return hu
}

// AddDownloadClientID adds i to the "download_client_id" field.
func (hu *HistoryUpdate) AddDownloadClientID(i int) *HistoryUpdate {
	hu.mutation.AddDownloadClientID(i)
	return hu
}

// ClearDownloadClientID clears the value of the "download_client_id" field.
func (hu *HistoryUpdate) ClearDownloadClientID() *HistoryUpdate {
	hu.mutation.ClearDownloadClientID()
	return hu
}

// SetStatus sets the "status" field.
func (hu *HistoryUpdate) SetStatus(h history.Status) *HistoryUpdate {
	hu.mutation.SetStatus(h)
	return hu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (hu *HistoryUpdate) SetNillableStatus(h *history.Status) *HistoryUpdate {
	if h != nil {
		hu.SetStatus(*h)
	}
	return hu
}

// SetSaved sets the "saved" field.
func (hu *HistoryUpdate) SetSaved(s string) *HistoryUpdate {
	hu.mutation.SetSaved(s)
	return hu
}

// SetNillableSaved sets the "saved" field if the given value is not nil.
func (hu *HistoryUpdate) SetNillableSaved(s *string) *HistoryUpdate {
	if s != nil {
		hu.SetSaved(*s)
	}
	return hu
}

// ClearSaved clears the value of the "saved" field.
func (hu *HistoryUpdate) ClearSaved() *HistoryUpdate {
	hu.mutation.ClearSaved()
	return hu
}

// Mutation returns the HistoryMutation object of the builder.
func (hu *HistoryUpdate) Mutation() *HistoryMutation {
	return hu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (hu *HistoryUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, hu.sqlSave, hu.mutation, hu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (hu *HistoryUpdate) SaveX(ctx context.Context) int {
	affected, err := hu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (hu *HistoryUpdate) Exec(ctx context.Context) error {
	_, err := hu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hu *HistoryUpdate) ExecX(ctx context.Context) {
	if err := hu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (hu *HistoryUpdate) check() error {
	if v, ok := hu.mutation.Status(); ok {
		if err := history.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "History.status": %w`, err)}
		}
	}
	return nil
}

func (hu *HistoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := hu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(history.Table, history.Columns, sqlgraph.NewFieldSpec(history.FieldID, field.TypeInt))
	if ps := hu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := hu.mutation.MediaID(); ok {
		_spec.SetField(history.FieldMediaID, field.TypeInt, value)
	}
	if value, ok := hu.mutation.AddedMediaID(); ok {
		_spec.AddField(history.FieldMediaID, field.TypeInt, value)
	}
	if value, ok := hu.mutation.EpisodeID(); ok {
		_spec.SetField(history.FieldEpisodeID, field.TypeInt, value)
	}
	if value, ok := hu.mutation.AddedEpisodeID(); ok {
		_spec.AddField(history.FieldEpisodeID, field.TypeInt, value)
	}
	if hu.mutation.EpisodeIDCleared() {
		_spec.ClearField(history.FieldEpisodeID, field.TypeInt)
	}
	if value, ok := hu.mutation.SourceTitle(); ok {
		_spec.SetField(history.FieldSourceTitle, field.TypeString, value)
	}
	if value, ok := hu.mutation.Date(); ok {
		_spec.SetField(history.FieldDate, field.TypeTime, value)
	}
	if value, ok := hu.mutation.TargetDir(); ok {
		_spec.SetField(history.FieldTargetDir, field.TypeString, value)
	}
	if value, ok := hu.mutation.Size(); ok {
		_spec.SetField(history.FieldSize, field.TypeInt, value)
	}
	if value, ok := hu.mutation.AddedSize(); ok {
		_spec.AddField(history.FieldSize, field.TypeInt, value)
	}
	if value, ok := hu.mutation.DownloadClientID(); ok {
		_spec.SetField(history.FieldDownloadClientID, field.TypeInt, value)
	}
	if value, ok := hu.mutation.AddedDownloadClientID(); ok {
		_spec.AddField(history.FieldDownloadClientID, field.TypeInt, value)
	}
	if hu.mutation.DownloadClientIDCleared() {
		_spec.ClearField(history.FieldDownloadClientID, field.TypeInt)
	}
	if value, ok := hu.mutation.Status(); ok {
		_spec.SetField(history.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := hu.mutation.Saved(); ok {
		_spec.SetField(history.FieldSaved, field.TypeString, value)
	}
	if hu.mutation.SavedCleared() {
		_spec.ClearField(history.FieldSaved, field.TypeString)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, hu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{history.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	hu.mutation.done = true
	return n, nil
}

// HistoryUpdateOne is the builder for updating a single History entity.
type HistoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *HistoryMutation
}

// SetMediaID sets the "media_id" field.
func (huo *HistoryUpdateOne) SetMediaID(i int) *HistoryUpdateOne {
	huo.mutation.ResetMediaID()
	huo.mutation.SetMediaID(i)
	return huo
}

// SetNillableMediaID sets the "media_id" field if the given value is not nil.
func (huo *HistoryUpdateOne) SetNillableMediaID(i *int) *HistoryUpdateOne {
	if i != nil {
		huo.SetMediaID(*i)
	}
	return huo
}

// AddMediaID adds i to the "media_id" field.
func (huo *HistoryUpdateOne) AddMediaID(i int) *HistoryUpdateOne {
	huo.mutation.AddMediaID(i)
	return huo
}

// SetEpisodeID sets the "episode_id" field.
func (huo *HistoryUpdateOne) SetEpisodeID(i int) *HistoryUpdateOne {
	huo.mutation.ResetEpisodeID()
	huo.mutation.SetEpisodeID(i)
	return huo
}

// SetNillableEpisodeID sets the "episode_id" field if the given value is not nil.
func (huo *HistoryUpdateOne) SetNillableEpisodeID(i *int) *HistoryUpdateOne {
	if i != nil {
		huo.SetEpisodeID(*i)
	}
	return huo
}

// AddEpisodeID adds i to the "episode_id" field.
func (huo *HistoryUpdateOne) AddEpisodeID(i int) *HistoryUpdateOne {
	huo.mutation.AddEpisodeID(i)
	return huo
}

// ClearEpisodeID clears the value of the "episode_id" field.
func (huo *HistoryUpdateOne) ClearEpisodeID() *HistoryUpdateOne {
	huo.mutation.ClearEpisodeID()
	return huo
}

// SetSourceTitle sets the "source_title" field.
func (huo *HistoryUpdateOne) SetSourceTitle(s string) *HistoryUpdateOne {
	huo.mutation.SetSourceTitle(s)
	return huo
}

// SetNillableSourceTitle sets the "source_title" field if the given value is not nil.
func (huo *HistoryUpdateOne) SetNillableSourceTitle(s *string) *HistoryUpdateOne {
	if s != nil {
		huo.SetSourceTitle(*s)
	}
	return huo
}

// SetDate sets the "date" field.
func (huo *HistoryUpdateOne) SetDate(t time.Time) *HistoryUpdateOne {
	huo.mutation.SetDate(t)
	return huo
}

// SetNillableDate sets the "date" field if the given value is not nil.
func (huo *HistoryUpdateOne) SetNillableDate(t *time.Time) *HistoryUpdateOne {
	if t != nil {
		huo.SetDate(*t)
	}
	return huo
}

// SetTargetDir sets the "target_dir" field.
func (huo *HistoryUpdateOne) SetTargetDir(s string) *HistoryUpdateOne {
	huo.mutation.SetTargetDir(s)
	return huo
}

// SetNillableTargetDir sets the "target_dir" field if the given value is not nil.
func (huo *HistoryUpdateOne) SetNillableTargetDir(s *string) *HistoryUpdateOne {
	if s != nil {
		huo.SetTargetDir(*s)
	}
	return huo
}

// SetSize sets the "size" field.
func (huo *HistoryUpdateOne) SetSize(i int) *HistoryUpdateOne {
	huo.mutation.ResetSize()
	huo.mutation.SetSize(i)
	return huo
}

// SetNillableSize sets the "size" field if the given value is not nil.
func (huo *HistoryUpdateOne) SetNillableSize(i *int) *HistoryUpdateOne {
	if i != nil {
		huo.SetSize(*i)
	}
	return huo
}

// AddSize adds i to the "size" field.
func (huo *HistoryUpdateOne) AddSize(i int) *HistoryUpdateOne {
	huo.mutation.AddSize(i)
	return huo
}

// SetDownloadClientID sets the "download_client_id" field.
func (huo *HistoryUpdateOne) SetDownloadClientID(i int) *HistoryUpdateOne {
	huo.mutation.ResetDownloadClientID()
	huo.mutation.SetDownloadClientID(i)
	return huo
}

// SetNillableDownloadClientID sets the "download_client_id" field if the given value is not nil.
func (huo *HistoryUpdateOne) SetNillableDownloadClientID(i *int) *HistoryUpdateOne {
	if i != nil {
		huo.SetDownloadClientID(*i)
	}
	return huo
}

// AddDownloadClientID adds i to the "download_client_id" field.
func (huo *HistoryUpdateOne) AddDownloadClientID(i int) *HistoryUpdateOne {
	huo.mutation.AddDownloadClientID(i)
	return huo
}

// ClearDownloadClientID clears the value of the "download_client_id" field.
func (huo *HistoryUpdateOne) ClearDownloadClientID() *HistoryUpdateOne {
	huo.mutation.ClearDownloadClientID()
	return huo
}

// SetStatus sets the "status" field.
func (huo *HistoryUpdateOne) SetStatus(h history.Status) *HistoryUpdateOne {
	huo.mutation.SetStatus(h)
	return huo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (huo *HistoryUpdateOne) SetNillableStatus(h *history.Status) *HistoryUpdateOne {
	if h != nil {
		huo.SetStatus(*h)
	}
	return huo
}

// SetSaved sets the "saved" field.
func (huo *HistoryUpdateOne) SetSaved(s string) *HistoryUpdateOne {
	huo.mutation.SetSaved(s)
	return huo
}

// SetNillableSaved sets the "saved" field if the given value is not nil.
func (huo *HistoryUpdateOne) SetNillableSaved(s *string) *HistoryUpdateOne {
	if s != nil {
		huo.SetSaved(*s)
	}
	return huo
}

// ClearSaved clears the value of the "saved" field.
func (huo *HistoryUpdateOne) ClearSaved() *HistoryUpdateOne {
	huo.mutation.ClearSaved()
	return huo
}

// Mutation returns the HistoryMutation object of the builder.
func (huo *HistoryUpdateOne) Mutation() *HistoryMutation {
	return huo.mutation
}

// Where appends a list predicates to the HistoryUpdate builder.
func (huo *HistoryUpdateOne) Where(ps ...predicate.History) *HistoryUpdateOne {
	huo.mutation.Where(ps...)
	return huo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (huo *HistoryUpdateOne) Select(field string, fields ...string) *HistoryUpdateOne {
	huo.fields = append([]string{field}, fields...)
	return huo
}

// Save executes the query and returns the updated History entity.
func (huo *HistoryUpdateOne) Save(ctx context.Context) (*History, error) {
	return withHooks(ctx, huo.sqlSave, huo.mutation, huo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (huo *HistoryUpdateOne) SaveX(ctx context.Context) *History {
	node, err := huo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (huo *HistoryUpdateOne) Exec(ctx context.Context) error {
	_, err := huo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (huo *HistoryUpdateOne) ExecX(ctx context.Context) {
	if err := huo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (huo *HistoryUpdateOne) check() error {
	if v, ok := huo.mutation.Status(); ok {
		if err := history.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "History.status": %w`, err)}
		}
	}
	return nil
}

func (huo *HistoryUpdateOne) sqlSave(ctx context.Context) (_node *History, err error) {
	if err := huo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(history.Table, history.Columns, sqlgraph.NewFieldSpec(history.FieldID, field.TypeInt))
	id, ok := huo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "History.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := huo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, history.FieldID)
		for _, f := range fields {
			if !history.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != history.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := huo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := huo.mutation.MediaID(); ok {
		_spec.SetField(history.FieldMediaID, field.TypeInt, value)
	}
	if value, ok := huo.mutation.AddedMediaID(); ok {
		_spec.AddField(history.FieldMediaID, field.TypeInt, value)
	}
	if value, ok := huo.mutation.EpisodeID(); ok {
		_spec.SetField(history.FieldEpisodeID, field.TypeInt, value)
	}
	if value, ok := huo.mutation.AddedEpisodeID(); ok {
		_spec.AddField(history.FieldEpisodeID, field.TypeInt, value)
	}
	if huo.mutation.EpisodeIDCleared() {
		_spec.ClearField(history.FieldEpisodeID, field.TypeInt)
	}
	if value, ok := huo.mutation.SourceTitle(); ok {
		_spec.SetField(history.FieldSourceTitle, field.TypeString, value)
	}
	if value, ok := huo.mutation.Date(); ok {
		_spec.SetField(history.FieldDate, field.TypeTime, value)
	}
	if value, ok := huo.mutation.TargetDir(); ok {
		_spec.SetField(history.FieldTargetDir, field.TypeString, value)
	}
	if value, ok := huo.mutation.Size(); ok {
		_spec.SetField(history.FieldSize, field.TypeInt, value)
	}
	if value, ok := huo.mutation.AddedSize(); ok {
		_spec.AddField(history.FieldSize, field.TypeInt, value)
	}
	if value, ok := huo.mutation.DownloadClientID(); ok {
		_spec.SetField(history.FieldDownloadClientID, field.TypeInt, value)
	}
	if value, ok := huo.mutation.AddedDownloadClientID(); ok {
		_spec.AddField(history.FieldDownloadClientID, field.TypeInt, value)
	}
	if huo.mutation.DownloadClientIDCleared() {
		_spec.ClearField(history.FieldDownloadClientID, field.TypeInt)
	}
	if value, ok := huo.mutation.Status(); ok {
		_spec.SetField(history.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := huo.mutation.Saved(); ok {
		_spec.SetField(history.FieldSaved, field.TypeString, value)
	}
	if huo.mutation.SavedCleared() {
		_spec.ClearField(history.FieldSaved, field.TypeString)
	}
	_node = &History{config: huo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, huo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{history.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	huo.mutation.done = true
	return _node, nil
}
