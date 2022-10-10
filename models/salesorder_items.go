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
	"github.com/volatiletech/strmangle"
)

// SalesorderItem is an object representing the database table.
type SalesorderItem struct {
	ID           int      `boil:"id" json:"id" toml:"id" yaml:"id"`
	SalesorderID null.Int `boil:"salesorder_id" json:"salesorder_id,omitempty" toml:"salesorder_id" yaml:"salesorder_id,omitempty"`
	ProductID    int      `boil:"product_id" json:"product_id" toml:"product_id" yaml:"product_id"`
	Qty          int      `boil:"qty" json:"qty" toml:"qty" yaml:"qty"`

	R *salesorderItemR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L salesorderItemL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var SalesorderItemColumns = struct {
	ID           string
	SalesorderID string
	ProductID    string
	Qty          string
}{
	ID:           "id",
	SalesorderID: "salesorder_id",
	ProductID:    "product_id",
	Qty:          "qty",
}

var SalesorderItemTableColumns = struct {
	ID           string
	SalesorderID string
	ProductID    string
	Qty          string
}{
	ID:           "salesorder_items.id",
	SalesorderID: "salesorder_items.salesorder_id",
	ProductID:    "salesorder_items.product_id",
	Qty:          "salesorder_items.qty",
}

// Generated where

var SalesorderItemWhere = struct {
	ID           whereHelperint
	SalesorderID whereHelpernull_Int
	ProductID    whereHelperint
	Qty          whereHelperint
}{
	ID:           whereHelperint{field: "\"alpha\".\"salesorder_items\".\"id\""},
	SalesorderID: whereHelpernull_Int{field: "\"alpha\".\"salesorder_items\".\"salesorder_id\""},
	ProductID:    whereHelperint{field: "\"alpha\".\"salesorder_items\".\"product_id\""},
	Qty:          whereHelperint{field: "\"alpha\".\"salesorder_items\".\"qty\""},
}

// SalesorderItemRels is where relationship names are stored.
var SalesorderItemRels = struct {
	Product    string
	Salesorder string
}{
	Product:    "Product",
	Salesorder: "Salesorder",
}

// salesorderItemR is where relationships are stored.
type salesorderItemR struct {
	Product    *Product    `boil:"Product" json:"Product" toml:"Product" yaml:"Product"`
	Salesorder *Salesorder `boil:"Salesorder" json:"Salesorder" toml:"Salesorder" yaml:"Salesorder"`
}

// NewStruct creates a new relationship struct
func (*salesorderItemR) NewStruct() *salesorderItemR {
	return &salesorderItemR{}
}

func (r *salesorderItemR) GetProduct() *Product {
	if r == nil {
		return nil
	}
	return r.Product
}

func (r *salesorderItemR) GetSalesorder() *Salesorder {
	if r == nil {
		return nil
	}
	return r.Salesorder
}

// salesorderItemL is where Load methods for each relationship are stored.
type salesorderItemL struct{}

var (
	salesorderItemAllColumns            = []string{"id", "salesorder_id", "product_id", "qty"}
	salesorderItemColumnsWithoutDefault = []string{"product_id", "qty"}
	salesorderItemColumnsWithDefault    = []string{"id", "salesorder_id"}
	salesorderItemPrimaryKeyColumns     = []string{"id"}
	salesorderItemGeneratedColumns      = []string{"id"}
)

type (
	// SalesorderItemSlice is an alias for a slice of pointers to SalesorderItem.
	// This should almost always be used instead of []SalesorderItem.
	SalesorderItemSlice []*SalesorderItem

	salesorderItemQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	salesorderItemType                 = reflect.TypeOf(&SalesorderItem{})
	salesorderItemMapping              = queries.MakeStructMapping(salesorderItemType)
	salesorderItemPrimaryKeyMapping, _ = queries.BindMapping(salesorderItemType, salesorderItemMapping, salesorderItemPrimaryKeyColumns)
	salesorderItemInsertCacheMut       sync.RWMutex
	salesorderItemInsertCache          = make(map[string]insertCache)
	salesorderItemUpdateCacheMut       sync.RWMutex
	salesorderItemUpdateCache          = make(map[string]updateCache)
	salesorderItemUpsertCacheMut       sync.RWMutex
	salesorderItemUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// OneG returns a single salesorderItem record from the query using the global executor.
func (q salesorderItemQuery) OneG(ctx context.Context) (*SalesorderItem, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single salesorderItem record from the query.
func (q salesorderItemQuery) One(ctx context.Context, exec boil.ContextExecutor) (*SalesorderItem, error) {
	o := &SalesorderItem{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for salesorder_items")
	}

	return o, nil
}

// AllG returns all SalesorderItem records from the query using the global executor.
func (q salesorderItemQuery) AllG(ctx context.Context) (SalesorderItemSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all SalesorderItem records from the query.
func (q salesorderItemQuery) All(ctx context.Context, exec boil.ContextExecutor) (SalesorderItemSlice, error) {
	var o []*SalesorderItem

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to SalesorderItem slice")
	}

	return o, nil
}

// CountG returns the count of all SalesorderItem records in the query using the global executor
func (q salesorderItemQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all SalesorderItem records in the query.
func (q salesorderItemQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count salesorder_items rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q salesorderItemQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q salesorderItemQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if salesorder_items exists")
	}

	return count > 0, nil
}

// Product pointed to by the foreign key.
func (o *SalesorderItem) Product(mods ...qm.QueryMod) productQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.ProductID),
	}

	queryMods = append(queryMods, mods...)

	return Products(queryMods...)
}

// Salesorder pointed to by the foreign key.
func (o *SalesorderItem) Salesorder(mods ...qm.QueryMod) salesorderQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.SalesorderID),
	}

	queryMods = append(queryMods, mods...)

	return Salesorders(queryMods...)
}

// LoadProduct allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (salesorderItemL) LoadProduct(ctx context.Context, e boil.ContextExecutor, singular bool, maybeSalesorderItem interface{}, mods queries.Applicator) error {
	var slice []*SalesorderItem
	var object *SalesorderItem

	if singular {
		var ok bool
		object, ok = maybeSalesorderItem.(*SalesorderItem)
		if !ok {
			object = new(SalesorderItem)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeSalesorderItem)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeSalesorderItem))
			}
		}
	} else {
		s, ok := maybeSalesorderItem.(*[]*SalesorderItem)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeSalesorderItem)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeSalesorderItem))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &salesorderItemR{}
		}
		args = append(args, object.ProductID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &salesorderItemR{}
			}

			for _, a := range args {
				if a == obj.ProductID {
					continue Outer
				}
			}

			args = append(args, obj.ProductID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`alpha.products`),
		qm.WhereIn(`alpha.products.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Product")
	}

	var resultSlice []*Product
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Product")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for products")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for products")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Product = foreign
		if foreign.R == nil {
			foreign.R = &productR{}
		}
		foreign.R.SalesorderItems = append(foreign.R.SalesorderItems, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ProductID == foreign.ID {
				local.R.Product = foreign
				if foreign.R == nil {
					foreign.R = &productR{}
				}
				foreign.R.SalesorderItems = append(foreign.R.SalesorderItems, local)
				break
			}
		}
	}

	return nil
}

// LoadSalesorder allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (salesorderItemL) LoadSalesorder(ctx context.Context, e boil.ContextExecutor, singular bool, maybeSalesorderItem interface{}, mods queries.Applicator) error {
	var slice []*SalesorderItem
	var object *SalesorderItem

	if singular {
		var ok bool
		object, ok = maybeSalesorderItem.(*SalesorderItem)
		if !ok {
			object = new(SalesorderItem)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeSalesorderItem)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeSalesorderItem))
			}
		}
	} else {
		s, ok := maybeSalesorderItem.(*[]*SalesorderItem)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeSalesorderItem)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeSalesorderItem))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &salesorderItemR{}
		}
		if !queries.IsNil(object.SalesorderID) {
			args = append(args, object.SalesorderID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &salesorderItemR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.SalesorderID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.SalesorderID) {
				args = append(args, obj.SalesorderID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`alpha.salesorders`),
		qm.WhereIn(`alpha.salesorders.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Salesorder")
	}

	var resultSlice []*Salesorder
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Salesorder")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for salesorders")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for salesorders")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Salesorder = foreign
		if foreign.R == nil {
			foreign.R = &salesorderR{}
		}
		foreign.R.SalesorderItems = append(foreign.R.SalesorderItems, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.SalesorderID, foreign.ID) {
				local.R.Salesorder = foreign
				if foreign.R == nil {
					foreign.R = &salesorderR{}
				}
				foreign.R.SalesorderItems = append(foreign.R.SalesorderItems, local)
				break
			}
		}
	}

	return nil
}

// SetProductG of the salesorderItem to the related item.
// Sets o.R.Product to related.
// Adds o to related.R.SalesorderItems.
// Uses the global database handle.
func (o *SalesorderItem) SetProductG(ctx context.Context, insert bool, related *Product) error {
	return o.SetProduct(ctx, boil.GetContextDB(), insert, related)
}

// SetProduct of the salesorderItem to the related item.
// Sets o.R.Product to related.
// Adds o to related.R.SalesorderItems.
func (o *SalesorderItem) SetProduct(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Product) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"alpha\".\"salesorder_items\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"product_id"}),
		strmangle.WhereClause("\"", "\"", 2, salesorderItemPrimaryKeyColumns),
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

	o.ProductID = related.ID
	if o.R == nil {
		o.R = &salesorderItemR{
			Product: related,
		}
	} else {
		o.R.Product = related
	}

	if related.R == nil {
		related.R = &productR{
			SalesorderItems: SalesorderItemSlice{o},
		}
	} else {
		related.R.SalesorderItems = append(related.R.SalesorderItems, o)
	}

	return nil
}

// SetSalesorderG of the salesorderItem to the related item.
// Sets o.R.Salesorder to related.
// Adds o to related.R.SalesorderItems.
// Uses the global database handle.
func (o *SalesorderItem) SetSalesorderG(ctx context.Context, insert bool, related *Salesorder) error {
	return o.SetSalesorder(ctx, boil.GetContextDB(), insert, related)
}

// SetSalesorder of the salesorderItem to the related item.
// Sets o.R.Salesorder to related.
// Adds o to related.R.SalesorderItems.
func (o *SalesorderItem) SetSalesorder(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Salesorder) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"alpha\".\"salesorder_items\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"salesorder_id"}),
		strmangle.WhereClause("\"", "\"", 2, salesorderItemPrimaryKeyColumns),
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

	queries.Assign(&o.SalesorderID, related.ID)
	if o.R == nil {
		o.R = &salesorderItemR{
			Salesorder: related,
		}
	} else {
		o.R.Salesorder = related
	}

	if related.R == nil {
		related.R = &salesorderR{
			SalesorderItems: SalesorderItemSlice{o},
		}
	} else {
		related.R.SalesorderItems = append(related.R.SalesorderItems, o)
	}

	return nil
}

// RemoveSalesorderG relationship.
// Sets o.R.Salesorder to nil.
// Removes o from all passed in related items' relationships struct.
// Uses the global database handle.
func (o *SalesorderItem) RemoveSalesorderG(ctx context.Context, related *Salesorder) error {
	return o.RemoveSalesorder(ctx, boil.GetContextDB(), related)
}

// RemoveSalesorder relationship.
// Sets o.R.Salesorder to nil.
// Removes o from all passed in related items' relationships struct.
func (o *SalesorderItem) RemoveSalesorder(ctx context.Context, exec boil.ContextExecutor, related *Salesorder) error {
	var err error

	queries.SetScanner(&o.SalesorderID, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("salesorder_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	if o.R != nil {
		o.R.Salesorder = nil
	}
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.SalesorderItems {
		if queries.Equal(o.SalesorderID, ri.SalesorderID) {
			continue
		}

		ln := len(related.R.SalesorderItems)
		if ln > 1 && i < ln-1 {
			related.R.SalesorderItems[i] = related.R.SalesorderItems[ln-1]
		}
		related.R.SalesorderItems = related.R.SalesorderItems[:ln-1]
		break
	}
	return nil
}

// SalesorderItems retrieves all the records using an executor.
func SalesorderItems(mods ...qm.QueryMod) salesorderItemQuery {
	mods = append(mods, qm.From("\"alpha\".\"salesorder_items\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"alpha\".\"salesorder_items\".*"})
	}

	return salesorderItemQuery{q}
}

// FindSalesorderItemG retrieves a single record by ID.
func FindSalesorderItemG(ctx context.Context, iD int, selectCols ...string) (*SalesorderItem, error) {
	return FindSalesorderItem(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindSalesorderItem retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindSalesorderItem(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*SalesorderItem, error) {
	salesorderItemObj := &SalesorderItem{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"alpha\".\"salesorder_items\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, salesorderItemObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from salesorder_items")
	}

	return salesorderItemObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *SalesorderItem) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *SalesorderItem) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no salesorder_items provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(salesorderItemColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	salesorderItemInsertCacheMut.RLock()
	cache, cached := salesorderItemInsertCache[key]
	salesorderItemInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			salesorderItemAllColumns,
			salesorderItemColumnsWithDefault,
			salesorderItemColumnsWithoutDefault,
			nzDefaults,
		)
		wl = strmangle.SetComplement(wl, salesorderItemGeneratedColumns)

		cache.valueMapping, err = queries.BindMapping(salesorderItemType, salesorderItemMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(salesorderItemType, salesorderItemMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"alpha\".\"salesorder_items\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"alpha\".\"salesorder_items\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into salesorder_items")
	}

	if !cached {
		salesorderItemInsertCacheMut.Lock()
		salesorderItemInsertCache[key] = cache
		salesorderItemInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single SalesorderItem record using the global executor.
// See Update for more documentation.
func (o *SalesorderItem) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the SalesorderItem.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *SalesorderItem) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	salesorderItemUpdateCacheMut.RLock()
	cache, cached := salesorderItemUpdateCache[key]
	salesorderItemUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			salesorderItemAllColumns,
			salesorderItemPrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, salesorderItemGeneratedColumns)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update salesorder_items, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"alpha\".\"salesorder_items\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, salesorderItemPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(salesorderItemType, salesorderItemMapping, append(wl, salesorderItemPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update salesorder_items row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for salesorder_items")
	}

	if !cached {
		salesorderItemUpdateCacheMut.Lock()
		salesorderItemUpdateCache[key] = cache
		salesorderItemUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (q salesorderItemQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q salesorderItemQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for salesorder_items")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for salesorder_items")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o SalesorderItemSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o SalesorderItemSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), salesorderItemPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"alpha\".\"salesorder_items\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, salesorderItemPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in salesorderItem slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all salesorderItem")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *SalesorderItem) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *SalesorderItem) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no salesorder_items provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(salesorderItemColumnsWithDefault, o)

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

	salesorderItemUpsertCacheMut.RLock()
	cache, cached := salesorderItemUpsertCache[key]
	salesorderItemUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			salesorderItemAllColumns,
			salesorderItemColumnsWithDefault,
			salesorderItemColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			salesorderItemAllColumns,
			salesorderItemPrimaryKeyColumns,
		)

		insert = strmangle.SetComplement(insert, salesorderItemGeneratedColumns)
		update = strmangle.SetComplement(update, salesorderItemGeneratedColumns)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert salesorder_items, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(salesorderItemPrimaryKeyColumns))
			copy(conflict, salesorderItemPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"alpha\".\"salesorder_items\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(salesorderItemType, salesorderItemMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(salesorderItemType, salesorderItemMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert salesorder_items")
	}

	if !cached {
		salesorderItemUpsertCacheMut.Lock()
		salesorderItemUpsertCache[key] = cache
		salesorderItemUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteG deletes a single SalesorderItem record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *SalesorderItem) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single SalesorderItem record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *SalesorderItem) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no SalesorderItem provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), salesorderItemPrimaryKeyMapping)
	sql := "DELETE FROM \"alpha\".\"salesorder_items\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from salesorder_items")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for salesorder_items")
	}

	return rowsAff, nil
}

func (q salesorderItemQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q salesorderItemQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no salesorderItemQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from salesorder_items")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for salesorder_items")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o SalesorderItemSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o SalesorderItemSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), salesorderItemPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"alpha\".\"salesorder_items\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, salesorderItemPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from salesorderItem slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for salesorder_items")
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *SalesorderItem) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no SalesorderItem provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *SalesorderItem) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindSalesorderItem(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *SalesorderItemSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty SalesorderItemSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *SalesorderItemSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := SalesorderItemSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), salesorderItemPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"alpha\".\"salesorder_items\".* FROM \"alpha\".\"salesorder_items\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, salesorderItemPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in SalesorderItemSlice")
	}

	*o = slice

	return nil
}

// SalesorderItemExistsG checks if the SalesorderItem row exists.
func SalesorderItemExistsG(ctx context.Context, iD int) (bool, error) {
	return SalesorderItemExists(ctx, boil.GetContextDB(), iD)
}

// SalesorderItemExists checks if the SalesorderItem row exists.
func SalesorderItemExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"alpha\".\"salesorder_items\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if salesorder_items exists")
	}

	return exists, nil
}