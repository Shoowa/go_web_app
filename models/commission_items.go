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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// CommissionItem is an object representing the database table.
type CommissionItem struct {
	ID           int `boil:"id" json:"id" toml:"id" yaml:"id"`
	CommissionID int `boil:"commission_id" json:"commission_id" toml:"commission_id" yaml:"commission_id"`
	CategoryID   int `boil:"category_id" json:"category_id" toml:"category_id" yaml:"category_id"`
	Budget       int `boil:"budget" json:"budget" toml:"budget" yaml:"budget"`
	Qty          int `boil:"qty" json:"qty" toml:"qty" yaml:"qty"`

	R *commissionItemR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L commissionItemL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CommissionItemColumns = struct {
	ID           string
	CommissionID string
	CategoryID   string
	Budget       string
	Qty          string
}{
	ID:           "id",
	CommissionID: "commission_id",
	CategoryID:   "category_id",
	Budget:       "budget",
	Qty:          "qty",
}

var CommissionItemTableColumns = struct {
	ID           string
	CommissionID string
	CategoryID   string
	Budget       string
	Qty          string
}{
	ID:           "commission_items.id",
	CommissionID: "commission_items.commission_id",
	CategoryID:   "commission_items.category_id",
	Budget:       "commission_items.budget",
	Qty:          "commission_items.qty",
}

// Generated where

var CommissionItemWhere = struct {
	ID           whereHelperint
	CommissionID whereHelperint
	CategoryID   whereHelperint
	Budget       whereHelperint
	Qty          whereHelperint
}{
	ID:           whereHelperint{field: "\"alpha\".\"commission_items\".\"id\""},
	CommissionID: whereHelperint{field: "\"alpha\".\"commission_items\".\"commission_id\""},
	CategoryID:   whereHelperint{field: "\"alpha\".\"commission_items\".\"category_id\""},
	Budget:       whereHelperint{field: "\"alpha\".\"commission_items\".\"budget\""},
	Qty:          whereHelperint{field: "\"alpha\".\"commission_items\".\"qty\""},
}

// CommissionItemRels is where relationship names are stored.
var CommissionItemRels = struct {
	Category   string
	Commission string
}{
	Category:   "Category",
	Commission: "Commission",
}

// commissionItemR is where relationships are stored.
type commissionItemR struct {
	Category   *Category   `boil:"Category" json:"Category" toml:"Category" yaml:"Category"`
	Commission *Commission `boil:"Commission" json:"Commission" toml:"Commission" yaml:"Commission"`
}

// NewStruct creates a new relationship struct
func (*commissionItemR) NewStruct() *commissionItemR {
	return &commissionItemR{}
}

func (r *commissionItemR) GetCategory() *Category {
	if r == nil {
		return nil
	}
	return r.Category
}

func (r *commissionItemR) GetCommission() *Commission {
	if r == nil {
		return nil
	}
	return r.Commission
}

// commissionItemL is where Load methods for each relationship are stored.
type commissionItemL struct{}

var (
	commissionItemAllColumns            = []string{"id", "commission_id", "category_id", "budget", "qty"}
	commissionItemColumnsWithoutDefault = []string{"commission_id", "category_id", "budget", "qty"}
	commissionItemColumnsWithDefault    = []string{"id"}
	commissionItemPrimaryKeyColumns     = []string{"id"}
	commissionItemGeneratedColumns      = []string{"id"}
)

type (
	// CommissionItemSlice is an alias for a slice of pointers to CommissionItem.
	// This should almost always be used instead of []CommissionItem.
	CommissionItemSlice []*CommissionItem

	commissionItemQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	commissionItemType                 = reflect.TypeOf(&CommissionItem{})
	commissionItemMapping              = queries.MakeStructMapping(commissionItemType)
	commissionItemPrimaryKeyMapping, _ = queries.BindMapping(commissionItemType, commissionItemMapping, commissionItemPrimaryKeyColumns)
	commissionItemInsertCacheMut       sync.RWMutex
	commissionItemInsertCache          = make(map[string]insertCache)
	commissionItemUpdateCacheMut       sync.RWMutex
	commissionItemUpdateCache          = make(map[string]updateCache)
	commissionItemUpsertCacheMut       sync.RWMutex
	commissionItemUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// OneG returns a single commissionItem record from the query using the global executor.
func (q commissionItemQuery) OneG(ctx context.Context) (*CommissionItem, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single commissionItem record from the query.
func (q commissionItemQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CommissionItem, error) {
	o := &CommissionItem{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for commission_items")
	}

	return o, nil
}

// AllG returns all CommissionItem records from the query using the global executor.
func (q commissionItemQuery) AllG(ctx context.Context) (CommissionItemSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all CommissionItem records from the query.
func (q commissionItemQuery) All(ctx context.Context, exec boil.ContextExecutor) (CommissionItemSlice, error) {
	var o []*CommissionItem

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CommissionItem slice")
	}

	return o, nil
}

// CountG returns the count of all CommissionItem records in the query using the global executor
func (q commissionItemQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all CommissionItem records in the query.
func (q commissionItemQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count commission_items rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q commissionItemQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q commissionItemQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if commission_items exists")
	}

	return count > 0, nil
}

// Category pointed to by the foreign key.
func (o *CommissionItem) Category(mods ...qm.QueryMod) categoryQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.CategoryID),
	}

	queryMods = append(queryMods, mods...)

	return Categories(queryMods...)
}

// Commission pointed to by the foreign key.
func (o *CommissionItem) Commission(mods ...qm.QueryMod) commissionQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.CommissionID),
	}

	queryMods = append(queryMods, mods...)

	return Commissions(queryMods...)
}

// LoadCategory allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (commissionItemL) LoadCategory(ctx context.Context, e boil.ContextExecutor, singular bool, maybeCommissionItem interface{}, mods queries.Applicator) error {
	var slice []*CommissionItem
	var object *CommissionItem

	if singular {
		var ok bool
		object, ok = maybeCommissionItem.(*CommissionItem)
		if !ok {
			object = new(CommissionItem)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeCommissionItem)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeCommissionItem))
			}
		}
	} else {
		s, ok := maybeCommissionItem.(*[]*CommissionItem)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeCommissionItem)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeCommissionItem))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &commissionItemR{}
		}
		args = append(args, object.CategoryID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &commissionItemR{}
			}

			for _, a := range args {
				if a == obj.CategoryID {
					continue Outer
				}
			}

			args = append(args, obj.CategoryID)

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
		foreign.R.CommissionItems = append(foreign.R.CommissionItems, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.CategoryID == foreign.ID {
				local.R.Category = foreign
				if foreign.R == nil {
					foreign.R = &categoryR{}
				}
				foreign.R.CommissionItems = append(foreign.R.CommissionItems, local)
				break
			}
		}
	}

	return nil
}

// LoadCommission allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (commissionItemL) LoadCommission(ctx context.Context, e boil.ContextExecutor, singular bool, maybeCommissionItem interface{}, mods queries.Applicator) error {
	var slice []*CommissionItem
	var object *CommissionItem

	if singular {
		var ok bool
		object, ok = maybeCommissionItem.(*CommissionItem)
		if !ok {
			object = new(CommissionItem)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeCommissionItem)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeCommissionItem))
			}
		}
	} else {
		s, ok := maybeCommissionItem.(*[]*CommissionItem)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeCommissionItem)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeCommissionItem))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &commissionItemR{}
		}
		args = append(args, object.CommissionID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &commissionItemR{}
			}

			for _, a := range args {
				if a == obj.CommissionID {
					continue Outer
				}
			}

			args = append(args, obj.CommissionID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`alpha.commissions`),
		qm.WhereIn(`alpha.commissions.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Commission")
	}

	var resultSlice []*Commission
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Commission")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for commissions")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for commissions")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Commission = foreign
		if foreign.R == nil {
			foreign.R = &commissionR{}
		}
		foreign.R.CommissionItems = append(foreign.R.CommissionItems, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.CommissionID == foreign.ID {
				local.R.Commission = foreign
				if foreign.R == nil {
					foreign.R = &commissionR{}
				}
				foreign.R.CommissionItems = append(foreign.R.CommissionItems, local)
				break
			}
		}
	}

	return nil
}

// SetCategoryG of the commissionItem to the related item.
// Sets o.R.Category to related.
// Adds o to related.R.CommissionItems.
// Uses the global database handle.
func (o *CommissionItem) SetCategoryG(ctx context.Context, insert bool, related *Category) error {
	return o.SetCategory(ctx, boil.GetContextDB(), insert, related)
}

// SetCategory of the commissionItem to the related item.
// Sets o.R.Category to related.
// Adds o to related.R.CommissionItems.
func (o *CommissionItem) SetCategory(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Category) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"alpha\".\"commission_items\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"category_id"}),
		strmangle.WhereClause("\"", "\"", 2, commissionItemPrimaryKeyColumns),
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

	o.CategoryID = related.ID
	if o.R == nil {
		o.R = &commissionItemR{
			Category: related,
		}
	} else {
		o.R.Category = related
	}

	if related.R == nil {
		related.R = &categoryR{
			CommissionItems: CommissionItemSlice{o},
		}
	} else {
		related.R.CommissionItems = append(related.R.CommissionItems, o)
	}

	return nil
}

// SetCommissionG of the commissionItem to the related item.
// Sets o.R.Commission to related.
// Adds o to related.R.CommissionItems.
// Uses the global database handle.
func (o *CommissionItem) SetCommissionG(ctx context.Context, insert bool, related *Commission) error {
	return o.SetCommission(ctx, boil.GetContextDB(), insert, related)
}

// SetCommission of the commissionItem to the related item.
// Sets o.R.Commission to related.
// Adds o to related.R.CommissionItems.
func (o *CommissionItem) SetCommission(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Commission) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"alpha\".\"commission_items\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"commission_id"}),
		strmangle.WhereClause("\"", "\"", 2, commissionItemPrimaryKeyColumns),
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

	o.CommissionID = related.ID
	if o.R == nil {
		o.R = &commissionItemR{
			Commission: related,
		}
	} else {
		o.R.Commission = related
	}

	if related.R == nil {
		related.R = &commissionR{
			CommissionItems: CommissionItemSlice{o},
		}
	} else {
		related.R.CommissionItems = append(related.R.CommissionItems, o)
	}

	return nil
}

// CommissionItems retrieves all the records using an executor.
func CommissionItems(mods ...qm.QueryMod) commissionItemQuery {
	mods = append(mods, qm.From("\"alpha\".\"commission_items\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"alpha\".\"commission_items\".*"})
	}

	return commissionItemQuery{q}
}

// FindCommissionItemG retrieves a single record by ID.
func FindCommissionItemG(ctx context.Context, iD int, selectCols ...string) (*CommissionItem, error) {
	return FindCommissionItem(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindCommissionItem retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCommissionItem(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*CommissionItem, error) {
	commissionItemObj := &CommissionItem{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"alpha\".\"commission_items\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, commissionItemObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from commission_items")
	}

	return commissionItemObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *CommissionItem) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CommissionItem) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no commission_items provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(commissionItemColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	commissionItemInsertCacheMut.RLock()
	cache, cached := commissionItemInsertCache[key]
	commissionItemInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			commissionItemAllColumns,
			commissionItemColumnsWithDefault,
			commissionItemColumnsWithoutDefault,
			nzDefaults,
		)
		wl = strmangle.SetComplement(wl, commissionItemGeneratedColumns)

		cache.valueMapping, err = queries.BindMapping(commissionItemType, commissionItemMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(commissionItemType, commissionItemMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"alpha\".\"commission_items\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"alpha\".\"commission_items\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into commission_items")
	}

	if !cached {
		commissionItemInsertCacheMut.Lock()
		commissionItemInsertCache[key] = cache
		commissionItemInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single CommissionItem record using the global executor.
// See Update for more documentation.
func (o *CommissionItem) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the CommissionItem.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CommissionItem) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	commissionItemUpdateCacheMut.RLock()
	cache, cached := commissionItemUpdateCache[key]
	commissionItemUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			commissionItemAllColumns,
			commissionItemPrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, commissionItemGeneratedColumns)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update commission_items, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"alpha\".\"commission_items\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, commissionItemPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(commissionItemType, commissionItemMapping, append(wl, commissionItemPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update commission_items row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for commission_items")
	}

	if !cached {
		commissionItemUpdateCacheMut.Lock()
		commissionItemUpdateCache[key] = cache
		commissionItemUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (q commissionItemQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q commissionItemQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for commission_items")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for commission_items")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o CommissionItemSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CommissionItemSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), commissionItemPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"alpha\".\"commission_items\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, commissionItemPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in commissionItem slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all commissionItem")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *CommissionItem) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CommissionItem) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no commission_items provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(commissionItemColumnsWithDefault, o)

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

	commissionItemUpsertCacheMut.RLock()
	cache, cached := commissionItemUpsertCache[key]
	commissionItemUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			commissionItemAllColumns,
			commissionItemColumnsWithDefault,
			commissionItemColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			commissionItemAllColumns,
			commissionItemPrimaryKeyColumns,
		)

		insert = strmangle.SetComplement(insert, commissionItemGeneratedColumns)
		update = strmangle.SetComplement(update, commissionItemGeneratedColumns)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert commission_items, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(commissionItemPrimaryKeyColumns))
			copy(conflict, commissionItemPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"alpha\".\"commission_items\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(commissionItemType, commissionItemMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(commissionItemType, commissionItemMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert commission_items")
	}

	if !cached {
		commissionItemUpsertCacheMut.Lock()
		commissionItemUpsertCache[key] = cache
		commissionItemUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteG deletes a single CommissionItem record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *CommissionItem) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single CommissionItem record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CommissionItem) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CommissionItem provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), commissionItemPrimaryKeyMapping)
	sql := "DELETE FROM \"alpha\".\"commission_items\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from commission_items")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for commission_items")
	}

	return rowsAff, nil
}

func (q commissionItemQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q commissionItemQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no commissionItemQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from commission_items")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for commission_items")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o CommissionItemSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CommissionItemSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), commissionItemPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"alpha\".\"commission_items\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, commissionItemPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from commissionItem slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for commission_items")
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *CommissionItem) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no CommissionItem provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *CommissionItem) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCommissionItem(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CommissionItemSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty CommissionItemSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CommissionItemSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CommissionItemSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), commissionItemPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"alpha\".\"commission_items\".* FROM \"alpha\".\"commission_items\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, commissionItemPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CommissionItemSlice")
	}

	*o = slice

	return nil
}

// CommissionItemExistsG checks if the CommissionItem row exists.
func CommissionItemExistsG(ctx context.Context, iD int) (bool, error) {
	return CommissionItemExists(ctx, boil.GetContextDB(), iD)
}

// CommissionItemExists checks if the CommissionItem row exists.
func CommissionItemExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"alpha\".\"commission_items\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if commission_items exists")
	}

	return exists, nil
}
