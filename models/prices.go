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

// Price is an object representing the database table.
type Price struct {
	ID        int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	ProductID int       `boil:"product_id" json:"product_id" toml:"product_id" yaml:"product_id"`
	Since     time.Time `boil:"since" json:"since" toml:"since" yaml:"since"`
	Price     int       `boil:"price" json:"price" toml:"price" yaml:"price"`

	R *priceR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L priceL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PriceColumns = struct {
	ID        string
	ProductID string
	Since     string
	Price     string
}{
	ID:        "id",
	ProductID: "product_id",
	Since:     "since",
	Price:     "price",
}

var PriceTableColumns = struct {
	ID        string
	ProductID string
	Since     string
	Price     string
}{
	ID:        "prices.id",
	ProductID: "prices.product_id",
	Since:     "prices.since",
	Price:     "prices.price",
}

// Generated where

var PriceWhere = struct {
	ID        whereHelperint
	ProductID whereHelperint
	Since     whereHelpertime_Time
	Price     whereHelperint
}{
	ID:        whereHelperint{field: "\"alpha\".\"prices\".\"id\""},
	ProductID: whereHelperint{field: "\"alpha\".\"prices\".\"product_id\""},
	Since:     whereHelpertime_Time{field: "\"alpha\".\"prices\".\"since\""},
	Price:     whereHelperint{field: "\"alpha\".\"prices\".\"price\""},
}

// PriceRels is where relationship names are stored.
var PriceRels = struct {
	Product string
}{
	Product: "Product",
}

// priceR is where relationships are stored.
type priceR struct {
	Product *Product `boil:"Product" json:"Product" toml:"Product" yaml:"Product"`
}

// NewStruct creates a new relationship struct
func (*priceR) NewStruct() *priceR {
	return &priceR{}
}

func (r *priceR) GetProduct() *Product {
	if r == nil {
		return nil
	}
	return r.Product
}

// priceL is where Load methods for each relationship are stored.
type priceL struct{}

var (
	priceAllColumns            = []string{"id", "product_id", "since", "price"}
	priceColumnsWithoutDefault = []string{"product_id", "since", "price"}
	priceColumnsWithDefault    = []string{"id"}
	pricePrimaryKeyColumns     = []string{"id"}
	priceGeneratedColumns      = []string{"id"}
)

type (
	// PriceSlice is an alias for a slice of pointers to Price.
	// This should almost always be used instead of []Price.
	PriceSlice []*Price

	priceQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	priceType                 = reflect.TypeOf(&Price{})
	priceMapping              = queries.MakeStructMapping(priceType)
	pricePrimaryKeyMapping, _ = queries.BindMapping(priceType, priceMapping, pricePrimaryKeyColumns)
	priceInsertCacheMut       sync.RWMutex
	priceInsertCache          = make(map[string]insertCache)
	priceUpdateCacheMut       sync.RWMutex
	priceUpdateCache          = make(map[string]updateCache)
	priceUpsertCacheMut       sync.RWMutex
	priceUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// OneG returns a single price record from the query using the global executor.
func (q priceQuery) OneG(ctx context.Context) (*Price, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single price record from the query.
func (q priceQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Price, error) {
	o := &Price{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for prices")
	}

	return o, nil
}

// AllG returns all Price records from the query using the global executor.
func (q priceQuery) AllG(ctx context.Context) (PriceSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all Price records from the query.
func (q priceQuery) All(ctx context.Context, exec boil.ContextExecutor) (PriceSlice, error) {
	var o []*Price

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Price slice")
	}

	return o, nil
}

// CountG returns the count of all Price records in the query using the global executor
func (q priceQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all Price records in the query.
func (q priceQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count prices rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q priceQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q priceQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if prices exists")
	}

	return count > 0, nil
}

// Product pointed to by the foreign key.
func (o *Price) Product(mods ...qm.QueryMod) productQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.ProductID),
	}

	queryMods = append(queryMods, mods...)

	return Products(queryMods...)
}

// LoadProduct allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (priceL) LoadProduct(ctx context.Context, e boil.ContextExecutor, singular bool, maybePrice interface{}, mods queries.Applicator) error {
	var slice []*Price
	var object *Price

	if singular {
		var ok bool
		object, ok = maybePrice.(*Price)
		if !ok {
			object = new(Price)
			ok = queries.SetFromEmbeddedStruct(&object, &maybePrice)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybePrice))
			}
		}
	} else {
		s, ok := maybePrice.(*[]*Price)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybePrice)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybePrice))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &priceR{}
		}
		args = append(args, object.ProductID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &priceR{}
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
		foreign.R.Prices = append(foreign.R.Prices, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ProductID == foreign.ID {
				local.R.Product = foreign
				if foreign.R == nil {
					foreign.R = &productR{}
				}
				foreign.R.Prices = append(foreign.R.Prices, local)
				break
			}
		}
	}

	return nil
}

// SetProductG of the price to the related item.
// Sets o.R.Product to related.
// Adds o to related.R.Prices.
// Uses the global database handle.
func (o *Price) SetProductG(ctx context.Context, insert bool, related *Product) error {
	return o.SetProduct(ctx, boil.GetContextDB(), insert, related)
}

// SetProduct of the price to the related item.
// Sets o.R.Product to related.
// Adds o to related.R.Prices.
func (o *Price) SetProduct(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Product) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"alpha\".\"prices\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"product_id"}),
		strmangle.WhereClause("\"", "\"", 2, pricePrimaryKeyColumns),
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
		o.R = &priceR{
			Product: related,
		}
	} else {
		o.R.Product = related
	}

	if related.R == nil {
		related.R = &productR{
			Prices: PriceSlice{o},
		}
	} else {
		related.R.Prices = append(related.R.Prices, o)
	}

	return nil
}

// Prices retrieves all the records using an executor.
func Prices(mods ...qm.QueryMod) priceQuery {
	mods = append(mods, qm.From("\"alpha\".\"prices\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"alpha\".\"prices\".*"})
	}

	return priceQuery{q}
}

// FindPriceG retrieves a single record by ID.
func FindPriceG(ctx context.Context, iD int, selectCols ...string) (*Price, error) {
	return FindPrice(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindPrice retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPrice(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Price, error) {
	priceObj := &Price{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"alpha\".\"prices\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, priceObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from prices")
	}

	return priceObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Price) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Price) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no prices provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(priceColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	priceInsertCacheMut.RLock()
	cache, cached := priceInsertCache[key]
	priceInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			priceAllColumns,
			priceColumnsWithDefault,
			priceColumnsWithoutDefault,
			nzDefaults,
		)
		wl = strmangle.SetComplement(wl, priceGeneratedColumns)

		cache.valueMapping, err = queries.BindMapping(priceType, priceMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(priceType, priceMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"alpha\".\"prices\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"alpha\".\"prices\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into prices")
	}

	if !cached {
		priceInsertCacheMut.Lock()
		priceInsertCache[key] = cache
		priceInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single Price record using the global executor.
// See Update for more documentation.
func (o *Price) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the Price.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Price) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	priceUpdateCacheMut.RLock()
	cache, cached := priceUpdateCache[key]
	priceUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			priceAllColumns,
			pricePrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, priceGeneratedColumns)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update prices, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"alpha\".\"prices\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, pricePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(priceType, priceMapping, append(wl, pricePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update prices row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for prices")
	}

	if !cached {
		priceUpdateCacheMut.Lock()
		priceUpdateCache[key] = cache
		priceUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (q priceQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q priceQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for prices")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for prices")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o PriceSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PriceSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), pricePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"alpha\".\"prices\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, pricePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in price slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all price")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Price) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Price) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no prices provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(priceColumnsWithDefault, o)

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

	priceUpsertCacheMut.RLock()
	cache, cached := priceUpsertCache[key]
	priceUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			priceAllColumns,
			priceColumnsWithDefault,
			priceColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			priceAllColumns,
			pricePrimaryKeyColumns,
		)

		insert = strmangle.SetComplement(insert, priceGeneratedColumns)
		update = strmangle.SetComplement(update, priceGeneratedColumns)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert prices, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(pricePrimaryKeyColumns))
			copy(conflict, pricePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"alpha\".\"prices\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(priceType, priceMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(priceType, priceMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert prices")
	}

	if !cached {
		priceUpsertCacheMut.Lock()
		priceUpsertCache[key] = cache
		priceUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteG deletes a single Price record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Price) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single Price record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Price) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Price provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), pricePrimaryKeyMapping)
	sql := "DELETE FROM \"alpha\".\"prices\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from prices")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for prices")
	}

	return rowsAff, nil
}

func (q priceQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q priceQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no priceQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from prices")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for prices")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o PriceSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PriceSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), pricePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"alpha\".\"prices\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, pricePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from price slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for prices")
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Price) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no Price provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Price) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindPrice(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PriceSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty PriceSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PriceSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PriceSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), pricePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"alpha\".\"prices\".* FROM \"alpha\".\"prices\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, pricePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in PriceSlice")
	}

	*o = slice

	return nil
}

// PriceExistsG checks if the Price row exists.
func PriceExistsG(ctx context.Context, iD int) (bool, error) {
	return PriceExists(ctx, boil.GetContextDB(), iD)
}

// PriceExists checks if the Price row exists.
func PriceExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"alpha\".\"prices\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if prices exists")
	}

	return exists, nil
}
