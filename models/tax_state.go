// Code generated by SQLBoiler 4.13.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/v4/types"
	"github.com/volatiletech/strmangle"
)

// TaxState is an object representing the database table.
type TaxState struct {
	ID         int           `boil:"id" json:"id" toml:"id" yaml:"id"`
	StateID    string        `boil:"state_id" json:"state_id" toml:"state_id" yaml:"state_id"`
	CategoryID null.Int      `boil:"category_id" json:"category_id,omitempty" toml:"category_id" yaml:"category_id,omitempty"`
	Since      time.Time     `boil:"since" json:"since" toml:"since" yaml:"since"`
	Rate       types.Decimal `boil:"rate" json:"rate" toml:"rate" yaml:"rate"`

	R *taxStateR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L taxStateL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TaxStateColumns = struct {
	ID         string
	StateID    string
	CategoryID string
	Since      string
	Rate       string
}{
	ID:         "id",
	StateID:    "state_id",
	CategoryID: "category_id",
	Since:      "since",
	Rate:       "rate",
}

var TaxStateTableColumns = struct {
	ID         string
	StateID    string
	CategoryID string
	Since      string
	Rate       string
}{
	ID:         "tax_state.id",
	StateID:    "tax_state.state_id",
	CategoryID: "tax_state.category_id",
	Since:      "tax_state.since",
	Rate:       "tax_state.rate",
}

// Generated where

var TaxStateWhere = struct {
	ID         whereHelperint
	StateID    whereHelperstring
	CategoryID whereHelpernull_Int
	Since      whereHelpertime_Time
	Rate       whereHelpertypes_Decimal
}{
	ID:         whereHelperint{field: "\"alpha\".\"tax_state\".\"id\""},
	StateID:    whereHelperstring{field: "\"alpha\".\"tax_state\".\"state_id\""},
	CategoryID: whereHelpernull_Int{field: "\"alpha\".\"tax_state\".\"category_id\""},
	Since:      whereHelpertime_Time{field: "\"alpha\".\"tax_state\".\"since\""},
	Rate:       whereHelpertypes_Decimal{field: "\"alpha\".\"tax_state\".\"rate\""},
}

// TaxStateRels is where relationship names are stored.
var TaxStateRels = struct {
	Category string
	State    string
}{
	Category: "Category",
	State:    "State",
}

// taxStateR is where relationships are stored.
type taxStateR struct {
	Category *Category `boil:"Category" json:"Category" toml:"Category" yaml:"Category"`
	State    *State    `boil:"State" json:"State" toml:"State" yaml:"State"`
}

// NewStruct creates a new relationship struct
func (*taxStateR) NewStruct() *taxStateR {
	return &taxStateR{}
}

func (r *taxStateR) GetCategory() *Category {
	if r == nil {
		return nil
	}
	return r.Category
}

func (r *taxStateR) GetState() *State {
	if r == nil {
		return nil
	}
	return r.State
}

// taxStateL is where Load methods for each relationship are stored.
type taxStateL struct{}

var (
	taxStateAllColumns            = []string{"id", "state_id", "category_id", "since", "rate"}
	taxStateColumnsWithoutDefault = []string{"state_id", "since", "rate"}
	taxStateColumnsWithDefault    = []string{"id", "category_id"}
	taxStatePrimaryKeyColumns     = []string{"id"}
	taxStateGeneratedColumns      = []string{"id"}
)

type (
	// TaxStateSlice is an alias for a slice of pointers to TaxState.
	// This should almost always be used instead of []TaxState.
	TaxStateSlice []*TaxState

	taxStateQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	taxStateType                 = reflect.TypeOf(&TaxState{})
	taxStateMapping              = queries.MakeStructMapping(taxStateType)
	taxStatePrimaryKeyMapping, _ = queries.BindMapping(taxStateType, taxStateMapping, taxStatePrimaryKeyColumns)
	taxStateInsertCacheMut       sync.RWMutex
	taxStateInsertCache          = make(map[string]insertCache)
	taxStateUpdateCacheMut       sync.RWMutex
	taxStateUpdateCache          = make(map[string]updateCache)
	taxStateUpsertCacheMut       sync.RWMutex
	taxStateUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// OneG returns a single taxState record from the query using the global executor.
func (q taxStateQuery) OneG(ctx context.Context) (*TaxState, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single taxState record from the query.
func (q taxStateQuery) One(ctx context.Context, exec boil.ContextExecutor) (*TaxState, error) {
	o := &TaxState{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for tax_state")
	}

	return o, nil
}

// AllG returns all TaxState records from the query using the global executor.
func (q taxStateQuery) AllG(ctx context.Context) (TaxStateSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all TaxState records from the query.
func (q taxStateQuery) All(ctx context.Context, exec boil.ContextExecutor) (TaxStateSlice, error) {
	var o []*TaxState

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to TaxState slice")
	}

	return o, nil
}

// CountG returns the count of all TaxState records in the query using the global executor
func (q taxStateQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all TaxState records in the query.
func (q taxStateQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count tax_state rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q taxStateQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q taxStateQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if tax_state exists")
	}

	return count > 0, nil
}

// Category pointed to by the foreign key.
func (o *TaxState) Category(mods ...qm.QueryMod) categoryQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.CategoryID),
	}

	queryMods = append(queryMods, mods...)

	return Categories(queryMods...)
}

// State pointed to by the foreign key.
func (o *TaxState) State(mods ...qm.QueryMod) stateQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.StateID),
	}

	queryMods = append(queryMods, mods...)

	return States(queryMods...)
}

// LoadCategory allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (taxStateL) LoadCategory(ctx context.Context, e boil.ContextExecutor, singular bool, maybeTaxState interface{}, mods queries.Applicator) error {
	var slice []*TaxState
	var object *TaxState

	if singular {
		var ok bool
		object, ok = maybeTaxState.(*TaxState)
		if !ok {
			object = new(TaxState)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeTaxState)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeTaxState))
			}
		}
	} else {
		s, ok := maybeTaxState.(*[]*TaxState)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeTaxState)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeTaxState))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &taxStateR{}
		}
		if !queries.IsNil(object.CategoryID) {
			args = append(args, object.CategoryID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &taxStateR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.CategoryID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.CategoryID) {
				args = append(args, obj.CategoryID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`alpha.categories`),
		qm.WhereIn(`alpha.categories.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Category")
	}

	var resultSlice []*Category
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Category")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for categories")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for categories")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Category = foreign
		if foreign.R == nil {
			foreign.R = &categoryR{}
		}
		foreign.R.TaxStates = append(foreign.R.TaxStates, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.CategoryID, foreign.ID) {
				local.R.Category = foreign
				if foreign.R == nil {
					foreign.R = &categoryR{}
				}
				foreign.R.TaxStates = append(foreign.R.TaxStates, local)
				break
			}
		}
	}

	return nil
}

// LoadState allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (taxStateL) LoadState(ctx context.Context, e boil.ContextExecutor, singular bool, maybeTaxState interface{}, mods queries.Applicator) error {
	var slice []*TaxState
	var object *TaxState

	if singular {
		var ok bool
		object, ok = maybeTaxState.(*TaxState)
		if !ok {
			object = new(TaxState)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeTaxState)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeTaxState))
			}
		}
	} else {
		s, ok := maybeTaxState.(*[]*TaxState)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeTaxState)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeTaxState))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &taxStateR{}
		}
		args = append(args, object.StateID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &taxStateR{}
			}

			for _, a := range args {
				if a == obj.StateID {
					continue Outer
				}
			}

			args = append(args, obj.StateID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`alpha.state`),
		qm.WhereIn(`alpha.state.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load State")
	}

	var resultSlice []*State
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice State")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for state")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for state")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.State = foreign
		if foreign.R == nil {
			foreign.R = &stateR{}
		}
		foreign.R.TaxStates = append(foreign.R.TaxStates, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.StateID == foreign.ID {
				local.R.State = foreign
				if foreign.R == nil {
					foreign.R = &stateR{}
				}
				foreign.R.TaxStates = append(foreign.R.TaxStates, local)
				break
			}
		}
	}

	return nil
}

// SetCategoryG of the taxState to the related item.
// Sets o.R.Category to related.
// Adds o to related.R.TaxStates.
// Uses the global database handle.
func (o *TaxState) SetCategoryG(ctx context.Context, insert bool, related *Category) error {
	return o.SetCategory(ctx, boil.GetContextDB(), insert, related)
}

// SetCategory of the taxState to the related item.
// Sets o.R.Category to related.
// Adds o to related.R.TaxStates.
func (o *TaxState) SetCategory(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Category) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"alpha\".\"tax_state\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"category_id"}),
		strmangle.WhereClause("\"", "\"", 2, taxStatePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.CategoryID, related.ID)
	if o.R == nil {
		o.R = &taxStateR{
			Category: related,
		}
	} else {
		o.R.Category = related
	}

	if related.R == nil {
		related.R = &categoryR{
			TaxStates: TaxStateSlice{o},
		}
	} else {
		related.R.TaxStates = append(related.R.TaxStates, o)
	}

	return nil
}

// RemoveCategoryG relationship.
// Sets o.R.Category to nil.
// Removes o from all passed in related items' relationships struct.
// Uses the global database handle.
func (o *TaxState) RemoveCategoryG(ctx context.Context, related *Category) error {
	return o.RemoveCategory(ctx, boil.GetContextDB(), related)
}

// RemoveCategory relationship.
// Sets o.R.Category to nil.
// Removes o from all passed in related items' relationships struct.
func (o *TaxState) RemoveCategory(ctx context.Context, exec boil.ContextExecutor, related *Category) error {
	var err error

	queries.SetScanner(&o.CategoryID, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("category_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	if o.R != nil {
		o.R.Category = nil
	}
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.TaxStates {
		if queries.Equal(o.CategoryID, ri.CategoryID) {
			continue
		}

		ln := len(related.R.TaxStates)
		if ln > 1 && i < ln-1 {
			related.R.TaxStates[i] = related.R.TaxStates[ln-1]
		}
		related.R.TaxStates = related.R.TaxStates[:ln-1]
		break
	}
	return nil
}

// SetStateG of the taxState to the related item.
// Sets o.R.State to related.
// Adds o to related.R.TaxStates.
// Uses the global database handle.
func (o *TaxState) SetStateG(ctx context.Context, insert bool, related *State) error {
	return o.SetState(ctx, boil.GetContextDB(), insert, related)
}

// SetState of the taxState to the related item.
// Sets o.R.State to related.
// Adds o to related.R.TaxStates.
func (o *TaxState) SetState(ctx context.Context, exec boil.ContextExecutor, insert bool, related *State) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"alpha\".\"tax_state\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"state_id"}),
		strmangle.WhereClause("\"", "\"", 2, taxStatePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.StateID = related.ID
	if o.R == nil {
		o.R = &taxStateR{
			State: related,
		}
	} else {
		o.R.State = related
	}

	if related.R == nil {
		related.R = &stateR{
			TaxStates: TaxStateSlice{o},
		}
	} else {
		related.R.TaxStates = append(related.R.TaxStates, o)
	}

	return nil
}

// TaxStates retrieves all the records using an executor.
func TaxStates(mods ...qm.QueryMod) taxStateQuery {
	mods = append(mods, qm.From("\"alpha\".\"tax_state\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"alpha\".\"tax_state\".*"})
	}

	return taxStateQuery{q}
}

// FindTaxStateG retrieves a single record by ID.
func FindTaxStateG(ctx context.Context, iD int, selectCols ...string) (*TaxState, error) {
	return FindTaxState(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindTaxState retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTaxState(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*TaxState, error) {
	taxStateObj := &TaxState{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"alpha\".\"tax_state\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, taxStateObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from tax_state")
	}

	return taxStateObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *TaxState) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *TaxState) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no tax_state provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(taxStateColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	taxStateInsertCacheMut.RLock()
	cache, cached := taxStateInsertCache[key]
	taxStateInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			taxStateAllColumns,
			taxStateColumnsWithDefault,
			taxStateColumnsWithoutDefault,
			nzDefaults,
		)
		wl = strmangle.SetComplement(wl, taxStateGeneratedColumns)

		cache.valueMapping, err = queries.BindMapping(taxStateType, taxStateMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(taxStateType, taxStateMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"alpha\".\"tax_state\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"alpha\".\"tax_state\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into tax_state")
	}

	if !cached {
		taxStateInsertCacheMut.Lock()
		taxStateInsertCache[key] = cache
		taxStateInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single TaxState record using the global executor.
// See Update for more documentation.
func (o *TaxState) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the TaxState.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *TaxState) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	taxStateUpdateCacheMut.RLock()
	cache, cached := taxStateUpdateCache[key]
	taxStateUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			taxStateAllColumns,
			taxStatePrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, taxStateGeneratedColumns)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update tax_state, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"alpha\".\"tax_state\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, taxStatePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(taxStateType, taxStateMapping, append(wl, taxStatePrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update tax_state row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for tax_state")
	}

	if !cached {
		taxStateUpdateCacheMut.Lock()
		taxStateUpdateCache[key] = cache
		taxStateUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (q taxStateQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q taxStateQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for tax_state")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for tax_state")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o TaxStateSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TaxStateSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), taxStatePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"alpha\".\"tax_state\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, taxStatePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in taxState slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all taxState")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *TaxState) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *TaxState) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no tax_state provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(taxStateColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	taxStateUpsertCacheMut.RLock()
	cache, cached := taxStateUpsertCache[key]
	taxStateUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			taxStateAllColumns,
			taxStateColumnsWithDefault,
			taxStateColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			taxStateAllColumns,
			taxStatePrimaryKeyColumns,
		)

		insert = strmangle.SetComplement(insert, taxStateGeneratedColumns)
		update = strmangle.SetComplement(update, taxStateGeneratedColumns)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert tax_state, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(taxStatePrimaryKeyColumns))
			copy(conflict, taxStatePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"alpha\".\"tax_state\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(taxStateType, taxStateMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(taxStateType, taxStateMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert tax_state")
	}

	if !cached {
		taxStateUpsertCacheMut.Lock()
		taxStateUpsertCache[key] = cache
		taxStateUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteG deletes a single TaxState record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *TaxState) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single TaxState record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *TaxState) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no TaxState provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), taxStatePrimaryKeyMapping)
	sql := "DELETE FROM \"alpha\".\"tax_state\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from tax_state")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for tax_state")
	}

	return rowsAff, nil
}

func (q taxStateQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q taxStateQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no taxStateQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from tax_state")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for tax_state")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o TaxStateSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TaxStateSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), taxStatePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"alpha\".\"tax_state\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, taxStatePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from taxState slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for tax_state")
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *TaxState) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no TaxState provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *TaxState) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTaxState(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TaxStateSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty TaxStateSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TaxStateSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TaxStateSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), taxStatePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"alpha\".\"tax_state\".* FROM \"alpha\".\"tax_state\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, taxStatePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in TaxStateSlice")
	}

	*o = slice

	return nil
}

// TaxStateExistsG checks if the TaxState row exists.
func TaxStateExistsG(ctx context.Context, iD int) (bool, error) {
	return TaxStateExists(ctx, boil.GetContextDB(), iD)
}

// TaxStateExists checks if the TaxState row exists.
func TaxStateExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"alpha\".\"tax_state\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if tax_state exists")
	}

	return exists, nil
}
