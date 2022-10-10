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

// Stock is an object representing the database table.
type Stock struct {
	ID        int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	DepotID   null.Int  `boil:"depot_id" json:"depot_id,omitempty" toml:"depot_id" yaml:"depot_id,omitempty"`
	ProductID int       `boil:"product_id" json:"product_id" toml:"product_id" yaml:"product_id"`
	Since     time.Time `boil:"since" json:"since" toml:"since" yaml:"since"`
	Qty       int       `boil:"qty" json:"qty" toml:"qty" yaml:"qty"`

	R *stockR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L stockL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var StockColumns = struct {
	ID        string
	DepotID   string
	ProductID string
	Since     string
	Qty       string
}{
	ID:        "id",
	DepotID:   "depot_id",
	ProductID: "product_id",
	Since:     "since",
	Qty:       "qty",
}

var StockTableColumns = struct {
	ID        string
	DepotID   string
	ProductID string
	Since     string
	Qty       string
}{
	ID:        "stock.id",
	DepotID:   "stock.depot_id",
	ProductID: "stock.product_id",
	Since:     "stock.since",
	Qty:       "stock.qty",
}

// Generated where

var StockWhere = struct {
	ID        whereHelperint
	DepotID   whereHelpernull_Int
	ProductID whereHelperint
	Since     whereHelpertime_Time
	Qty       whereHelperint
}{
	ID:        whereHelperint{field: "\"alpha\".\"stock\".\"id\""},
	DepotID:   whereHelpernull_Int{field: "\"alpha\".\"stock\".\"depot_id\""},
	ProductID: whereHelperint{field: "\"alpha\".\"stock\".\"product_id\""},
	Since:     whereHelpertime_Time{field: "\"alpha\".\"stock\".\"since\""},
	Qty:       whereHelperint{field: "\"alpha\".\"stock\".\"qty\""},
}

// StockRels is where relationship names are stored.
var StockRels = struct {
	Depot   string
	Product string
}{
	Depot:   "Depot",
	Product: "Product",
}

// stockR is where relationships are stored.
type stockR struct {
	Depot   *Depot   `boil:"Depot" json:"Depot" toml:"Depot" yaml:"Depot"`
	Product *Product `boil:"Product" json:"Product" toml:"Product" yaml:"Product"`
}

// NewStruct creates a new relationship struct
func (*stockR) NewStruct() *stockR {
	return &stockR{}
}

func (r *stockR) GetDepot() *Depot {
	if r == nil {
		return nil
	}
	return r.Depot
}

func (r *stockR) GetProduct() *Product {
	if r == nil {
		return nil
	}
	return r.Product
}

// stockL is where Load methods for each relationship are stored.
type stockL struct{}

var (
	stockAllColumns            = []string{"id", "depot_id", "product_id", "since", "qty"}
	stockColumnsWithoutDefault = []string{"product_id", "since", "qty"}
	stockColumnsWithDefault    = []string{"id", "depot_id"}
	stockPrimaryKeyColumns     = []string{"id"}
	stockGeneratedColumns      = []string{"id"}
)

type (
	// StockSlice is an alias for a slice of pointers to Stock.
	// This should almost always be used instead of []Stock.
	StockSlice []*Stock

	stockQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	stockType                 = reflect.TypeOf(&Stock{})
	stockMapping              = queries.MakeStructMapping(stockType)
	stockPrimaryKeyMapping, _ = queries.BindMapping(stockType, stockMapping, stockPrimaryKeyColumns)
	stockInsertCacheMut       sync.RWMutex
	stockInsertCache          = make(map[string]insertCache)
	stockUpdateCacheMut       sync.RWMutex
	stockUpdateCache          = make(map[string]updateCache)
	stockUpsertCacheMut       sync.RWMutex
	stockUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// OneG returns a single stock record from the query using the global executor.
func (q stockQuery) OneG(ctx context.Context) (*Stock, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single stock record from the query.
func (q stockQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Stock, error) {
	o := &Stock{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for stock")
	}

	return o, nil
}

// AllG returns all Stock records from the query using the global executor.
func (q stockQuery) AllG(ctx context.Context) (StockSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all Stock records from the query.
func (q stockQuery) All(ctx context.Context, exec boil.ContextExecutor) (StockSlice, error) {
	var o []*Stock

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Stock slice")
	}

	return o, nil
}

// CountG returns the count of all Stock records in the query using the global executor
func (q stockQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all Stock records in the query.
func (q stockQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count stock rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q stockQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q stockQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if stock exists")
	}

	return count > 0, nil
}

// Depot pointed to by the foreign key.
func (o *Stock) Depot(mods ...qm.QueryMod) depotQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.DepotID),
	}

	queryMods = append(queryMods, mods...)

	return Depots(queryMods...)
}

// Product pointed to by the foreign key.
func (o *Stock) Product(mods ...qm.QueryMod) productQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.ProductID),
	}

	queryMods = append(queryMods, mods...)

	return Products(queryMods...)
}

// LoadDepot allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (stockL) LoadDepot(ctx context.Context, e boil.ContextExecutor, singular bool, maybeStock interface{}, mods queries.Applicator) error {
	var slice []*Stock
	var object *Stock

	if singular {
		var ok bool
		object, ok = maybeStock.(*Stock)
		if !ok {
			object = new(Stock)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeStock)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeStock))
			}
		}
	} else {
		s, ok := maybeStock.(*[]*Stock)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeStock)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeStock))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &stockR{}
		}
		if !queries.IsNil(object.DepotID) {
			args = append(args, object.DepotID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &stockR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.DepotID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.DepotID) {
				args = append(args, obj.DepotID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`alpha.depots`),
		qm.WhereIn(`alpha.depots.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Depot")
	}

	var resultSlice []*Depot
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Depot")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for depots")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for depots")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Depot = foreign
		if foreign.R == nil {
			foreign.R = &depotR{}
		}
		foreign.R.Stocks = append(foreign.R.Stocks, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.DepotID, foreign.ID) {
				local.R.Depot = foreign
				if foreign.R == nil {
					foreign.R = &depotR{}
				}
				foreign.R.Stocks = append(foreign.R.Stocks, local)
				break
			}
		}
	}

	return nil
}

// LoadProduct allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (stockL) LoadProduct(ctx context.Context, e boil.ContextExecutor, singular bool, maybeStock interface{}, mods queries.Applicator) error {
	var slice []*Stock
	var object *Stock

	if singular {
		var ok bool
		object, ok = maybeStock.(*Stock)
		if !ok {
			object = new(Stock)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeStock)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeStock))
			}
		}
	} else {
		s, ok := maybeStock.(*[]*Stock)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeStock)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeStock))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &stockR{}
		}
		args = append(args, object.ProductID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &stockR{}
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
		foreign.R.Stocks = append(foreign.R.Stocks, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ProductID == foreign.ID {
				local.R.Product = foreign
				if foreign.R == nil {
					foreign.R = &productR{}
				}
				foreign.R.Stocks = append(foreign.R.Stocks, local)
				break
			}
		}
	}

	return nil
}

// SetDepotG of the stock to the related item.
// Sets o.R.Depot to related.
// Adds o to related.R.Stocks.
// Uses the global database handle.
func (o *Stock) SetDepotG(ctx context.Context, insert bool, related *Depot) error {
	return o.SetDepot(ctx, boil.GetContextDB(), insert, related)
}

// SetDepot of the stock to the related item.
// Sets o.R.Depot to related.
// Adds o to related.R.Stocks.
func (o *Stock) SetDepot(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Depot) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"alpha\".\"stock\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"depot_id"}),
		strmangle.WhereClause("\"", "\"", 2, stockPrimaryKeyColumns),
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

	queries.Assign(&o.DepotID, related.ID)
	if o.R == nil {
		o.R = &stockR{
			Depot: related,
		}
	} else {
		o.R.Depot = related
	}

	if related.R == nil {
		related.R = &depotR{
			Stocks: StockSlice{o},
		}
	} else {
		related.R.Stocks = append(related.R.Stocks, o)
	}

	return nil
}

// RemoveDepotG relationship.
// Sets o.R.Depot to nil.
// Removes o from all passed in related items' relationships struct.
// Uses the global database handle.
func (o *Stock) RemoveDepotG(ctx context.Context, related *Depot) error {
	return o.RemoveDepot(ctx, boil.GetContextDB(), related)
}

// RemoveDepot relationship.
// Sets o.R.Depot to nil.
// Removes o from all passed in related items' relationships struct.
func (o *Stock) RemoveDepot(ctx context.Context, exec boil.ContextExecutor, related *Depot) error {
	var err error

	queries.SetScanner(&o.DepotID, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("depot_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	if o.R != nil {
		o.R.Depot = nil
	}
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.Stocks {
		if queries.Equal(o.DepotID, ri.DepotID) {
			continue
		}

		ln := len(related.R.Stocks)
		if ln > 1 && i < ln-1 {
			related.R.Stocks[i] = related.R.Stocks[ln-1]
		}
		related.R.Stocks = related.R.Stocks[:ln-1]
		break
	}
	return nil
}

// SetProductG of the stock to the related item.
// Sets o.R.Product to related.
// Adds o to related.R.Stocks.
// Uses the global database handle.
func (o *Stock) SetProductG(ctx context.Context, insert bool, related *Product) error {
	return o.SetProduct(ctx, boil.GetContextDB(), insert, related)
}

// SetProduct of the stock to the related item.
// Sets o.R.Product to related.
// Adds o to related.R.Stocks.
func (o *Stock) SetProduct(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Product) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"alpha\".\"stock\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"product_id"}),
		strmangle.WhereClause("\"", "\"", 2, stockPrimaryKeyColumns),
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
		o.R = &stockR{
			Product: related,
		}
	} else {
		o.R.Product = related
	}

	if related.R == nil {
		related.R = &productR{
			Stocks: StockSlice{o},
		}
	} else {
		related.R.Stocks = append(related.R.Stocks, o)
	}

	return nil
}

// Stocks retrieves all the records using an executor.
func Stocks(mods ...qm.QueryMod) stockQuery {
	mods = append(mods, qm.From("\"alpha\".\"stock\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"alpha\".\"stock\".*"})
	}

	return stockQuery{q}
}

// FindStockG retrieves a single record by ID.
func FindStockG(ctx context.Context, iD int, selectCols ...string) (*Stock, error) {
	return FindStock(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindStock retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindStock(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Stock, error) {
	stockObj := &Stock{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"alpha\".\"stock\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, stockObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from stock")
	}

	return stockObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Stock) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Stock) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no stock provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(stockColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	stockInsertCacheMut.RLock()
	cache, cached := stockInsertCache[key]
	stockInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			stockAllColumns,
			stockColumnsWithDefault,
			stockColumnsWithoutDefault,
			nzDefaults,
		)
		wl = strmangle.SetComplement(wl, stockGeneratedColumns)

		cache.valueMapping, err = queries.BindMapping(stockType, stockMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(stockType, stockMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"alpha\".\"stock\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"alpha\".\"stock\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into stock")
	}

	if !cached {
		stockInsertCacheMut.Lock()
		stockInsertCache[key] = cache
		stockInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single Stock record using the global executor.
// See Update for more documentation.
func (o *Stock) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the Stock.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Stock) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	stockUpdateCacheMut.RLock()
	cache, cached := stockUpdateCache[key]
	stockUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			stockAllColumns,
			stockPrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, stockGeneratedColumns)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update stock, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"alpha\".\"stock\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, stockPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(stockType, stockMapping, append(wl, stockPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update stock row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for stock")
	}

	if !cached {
		stockUpdateCacheMut.Lock()
		stockUpdateCache[key] = cache
		stockUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (q stockQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q stockQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for stock")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for stock")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o StockSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o StockSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stockPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"alpha\".\"stock\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, stockPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in stock slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all stock")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Stock) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Stock) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no stock provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(stockColumnsWithDefault, o)

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

	stockUpsertCacheMut.RLock()
	cache, cached := stockUpsertCache[key]
	stockUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			stockAllColumns,
			stockColumnsWithDefault,
			stockColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			stockAllColumns,
			stockPrimaryKeyColumns,
		)

		insert = strmangle.SetComplement(insert, stockGeneratedColumns)
		update = strmangle.SetComplement(update, stockGeneratedColumns)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert stock, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(stockPrimaryKeyColumns))
			copy(conflict, stockPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"alpha\".\"stock\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(stockType, stockMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(stockType, stockMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert stock")
	}

	if !cached {
		stockUpsertCacheMut.Lock()
		stockUpsertCache[key] = cache
		stockUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteG deletes a single Stock record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Stock) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single Stock record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Stock) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Stock provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), stockPrimaryKeyMapping)
	sql := "DELETE FROM \"alpha\".\"stock\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from stock")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for stock")
	}

	return rowsAff, nil
}

func (q stockQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q stockQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no stockQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from stock")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for stock")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o StockSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o StockSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stockPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"alpha\".\"stock\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, stockPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from stock slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for stock")
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Stock) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no Stock provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Stock) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindStock(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *StockSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty StockSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *StockSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := StockSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stockPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"alpha\".\"stock\".* FROM \"alpha\".\"stock\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, stockPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in StockSlice")
	}

	*o = slice

	return nil
}

// StockExistsG checks if the Stock row exists.
func StockExistsG(ctx context.Context, iD int) (bool, error) {
	return StockExists(ctx, boil.GetContextDB(), iD)
}

// StockExists checks if the Stock row exists.
func StockExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"alpha\".\"stock\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if stock exists")
	}

	return exists, nil
}