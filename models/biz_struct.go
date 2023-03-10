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

// BizStruct is an object representing the database table.
type BizStruct struct {
	ID   int    `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name string `boil:"name" json:"name" toml:"name" yaml:"name"`

	R *bizStructR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L bizStructL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var BizStructColumns = struct {
	ID   string
	Name string
}{
	ID:   "id",
	Name: "name",
}

var BizStructTableColumns = struct {
	ID   string
	Name string
}{
	ID:   "biz_struct.id",
	Name: "biz_struct.name",
}

// Generated where

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint) IN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint) NIN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

var BizStructWhere = struct {
	ID   whereHelperint
	Name whereHelperstring
}{
	ID:   whereHelperint{field: "\"alpha\".\"biz_struct\".\"id\""},
	Name: whereHelperstring{field: "\"alpha\".\"biz_struct\".\"name\""},
}

// BizStructRels is where relationship names are stored.
var BizStructRels = struct {
	StructCompanies string
}{
	StructCompanies: "StructCompanies",
}

// bizStructR is where relationships are stored.
type bizStructR struct {
	StructCompanies CompanySlice `boil:"StructCompanies" json:"StructCompanies" toml:"StructCompanies" yaml:"StructCompanies"`
}

// NewStruct creates a new relationship struct
func (*bizStructR) NewStruct() *bizStructR {
	return &bizStructR{}
}

func (r *bizStructR) GetStructCompanies() CompanySlice {
	if r == nil {
		return nil
	}
	return r.StructCompanies
}

// bizStructL is where Load methods for each relationship are stored.
type bizStructL struct{}

var (
	bizStructAllColumns            = []string{"id", "name"}
	bizStructColumnsWithoutDefault = []string{"name"}
	bizStructColumnsWithDefault    = []string{"id"}
	bizStructPrimaryKeyColumns     = []string{"id"}
	bizStructGeneratedColumns      = []string{"id"}
)

type (
	// BizStructSlice is an alias for a slice of pointers to BizStruct.
	// This should almost always be used instead of []BizStruct.
	BizStructSlice []*BizStruct

	bizStructQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	bizStructType                 = reflect.TypeOf(&BizStruct{})
	bizStructMapping              = queries.MakeStructMapping(bizStructType)
	bizStructPrimaryKeyMapping, _ = queries.BindMapping(bizStructType, bizStructMapping, bizStructPrimaryKeyColumns)
	bizStructInsertCacheMut       sync.RWMutex
	bizStructInsertCache          = make(map[string]insertCache)
	bizStructUpdateCacheMut       sync.RWMutex
	bizStructUpdateCache          = make(map[string]updateCache)
	bizStructUpsertCacheMut       sync.RWMutex
	bizStructUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// OneG returns a single bizStruct record from the query using the global executor.
func (q bizStructQuery) OneG(ctx context.Context) (*BizStruct, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single bizStruct record from the query.
func (q bizStructQuery) One(ctx context.Context, exec boil.ContextExecutor) (*BizStruct, error) {
	o := &BizStruct{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for biz_struct")
	}

	return o, nil
}

// AllG returns all BizStruct records from the query using the global executor.
func (q bizStructQuery) AllG(ctx context.Context) (BizStructSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all BizStruct records from the query.
func (q bizStructQuery) All(ctx context.Context, exec boil.ContextExecutor) (BizStructSlice, error) {
	var o []*BizStruct

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to BizStruct slice")
	}

	return o, nil
}

// CountG returns the count of all BizStruct records in the query using the global executor
func (q bizStructQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all BizStruct records in the query.
func (q bizStructQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count biz_struct rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q bizStructQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q bizStructQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if biz_struct exists")
	}

	return count > 0, nil
}

// StructCompanies retrieves all the company's Companies with an executor via struct_id column.
func (o *BizStruct) StructCompanies(mods ...qm.QueryMod) companyQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"alpha\".\"companies\".\"struct_id\"=?", o.ID),
	)

	return Companies(queryMods...)
}

// LoadStructCompanies allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (bizStructL) LoadStructCompanies(ctx context.Context, e boil.ContextExecutor, singular bool, maybeBizStruct interface{}, mods queries.Applicator) error {
	var slice []*BizStruct
	var object *BizStruct

	if singular {
		var ok bool
		object, ok = maybeBizStruct.(*BizStruct)
		if !ok {
			object = new(BizStruct)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeBizStruct)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeBizStruct))
			}
		}
	} else {
		s, ok := maybeBizStruct.(*[]*BizStruct)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeBizStruct)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeBizStruct))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &bizStructR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &bizStructR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`alpha.companies`),
		qm.WhereIn(`alpha.companies.struct_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load companies")
	}

	var resultSlice []*Company
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice companies")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on companies")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for companies")
	}

	if singular {
		object.R.StructCompanies = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &companyR{}
			}
			foreign.R.Struct = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.StructID {
				local.R.StructCompanies = append(local.R.StructCompanies, foreign)
				if foreign.R == nil {
					foreign.R = &companyR{}
				}
				foreign.R.Struct = local
				break
			}
		}
	}

	return nil
}

// AddStructCompaniesG adds the given related objects to the existing relationships
// of the biz_struct, optionally inserting them as new records.
// Appends related to o.R.StructCompanies.
// Sets related.R.Struct appropriately.
// Uses the global database handle.
func (o *BizStruct) AddStructCompaniesG(ctx context.Context, insert bool, related ...*Company) error {
	return o.AddStructCompanies(ctx, boil.GetContextDB(), insert, related...)
}

// AddStructCompanies adds the given related objects to the existing relationships
// of the biz_struct, optionally inserting them as new records.
// Appends related to o.R.StructCompanies.
// Sets related.R.Struct appropriately.
func (o *BizStruct) AddStructCompanies(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Company) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.StructID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"alpha\".\"companies\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"struct_id"}),
				strmangle.WhereClause("\"", "\"", 2, companyPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.StructID = o.ID
		}
	}

	if o.R == nil {
		o.R = &bizStructR{
			StructCompanies: related,
		}
	} else {
		o.R.StructCompanies = append(o.R.StructCompanies, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &companyR{
				Struct: o,
			}
		} else {
			rel.R.Struct = o
		}
	}
	return nil
}

// BizStructs retrieves all the records using an executor.
func BizStructs(mods ...qm.QueryMod) bizStructQuery {
	mods = append(mods, qm.From("\"alpha\".\"biz_struct\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"alpha\".\"biz_struct\".*"})
	}

	return bizStructQuery{q}
}

// FindBizStructG retrieves a single record by ID.
func FindBizStructG(ctx context.Context, iD int, selectCols ...string) (*BizStruct, error) {
	return FindBizStruct(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindBizStruct retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindBizStruct(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*BizStruct, error) {
	bizStructObj := &BizStruct{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"alpha\".\"biz_struct\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, bizStructObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from biz_struct")
	}

	return bizStructObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *BizStruct) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *BizStruct) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no biz_struct provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(bizStructColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	bizStructInsertCacheMut.RLock()
	cache, cached := bizStructInsertCache[key]
	bizStructInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			bizStructAllColumns,
			bizStructColumnsWithDefault,
			bizStructColumnsWithoutDefault,
			nzDefaults,
		)
		wl = strmangle.SetComplement(wl, bizStructGeneratedColumns)

		cache.valueMapping, err = queries.BindMapping(bizStructType, bizStructMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(bizStructType, bizStructMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"alpha\".\"biz_struct\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"alpha\".\"biz_struct\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into biz_struct")
	}

	if !cached {
		bizStructInsertCacheMut.Lock()
		bizStructInsertCache[key] = cache
		bizStructInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single BizStruct record using the global executor.
// See Update for more documentation.
func (o *BizStruct) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the BizStruct.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *BizStruct) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	bizStructUpdateCacheMut.RLock()
	cache, cached := bizStructUpdateCache[key]
	bizStructUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			bizStructAllColumns,
			bizStructPrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, bizStructGeneratedColumns)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update biz_struct, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"alpha\".\"biz_struct\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, bizStructPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(bizStructType, bizStructMapping, append(wl, bizStructPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update biz_struct row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for biz_struct")
	}

	if !cached {
		bizStructUpdateCacheMut.Lock()
		bizStructUpdateCache[key] = cache
		bizStructUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (q bizStructQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q bizStructQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for biz_struct")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for biz_struct")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o BizStructSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o BizStructSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bizStructPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"alpha\".\"biz_struct\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, bizStructPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in bizStruct slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all bizStruct")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *BizStruct) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *BizStruct) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no biz_struct provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(bizStructColumnsWithDefault, o)

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

	bizStructUpsertCacheMut.RLock()
	cache, cached := bizStructUpsertCache[key]
	bizStructUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			bizStructAllColumns,
			bizStructColumnsWithDefault,
			bizStructColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			bizStructAllColumns,
			bizStructPrimaryKeyColumns,
		)

		insert = strmangle.SetComplement(insert, bizStructGeneratedColumns)
		update = strmangle.SetComplement(update, bizStructGeneratedColumns)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert biz_struct, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(bizStructPrimaryKeyColumns))
			copy(conflict, bizStructPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"alpha\".\"biz_struct\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(bizStructType, bizStructMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(bizStructType, bizStructMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert biz_struct")
	}

	if !cached {
		bizStructUpsertCacheMut.Lock()
		bizStructUpsertCache[key] = cache
		bizStructUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteG deletes a single BizStruct record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *BizStruct) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single BizStruct record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *BizStruct) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no BizStruct provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), bizStructPrimaryKeyMapping)
	sql := "DELETE FROM \"alpha\".\"biz_struct\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from biz_struct")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for biz_struct")
	}

	return rowsAff, nil
}

func (q bizStructQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q bizStructQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no bizStructQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from biz_struct")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for biz_struct")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o BizStructSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o BizStructSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bizStructPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"alpha\".\"biz_struct\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, bizStructPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from bizStruct slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for biz_struct")
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *BizStruct) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no BizStruct provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *BizStruct) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindBizStruct(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *BizStructSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty BizStructSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *BizStructSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := BizStructSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bizStructPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"alpha\".\"biz_struct\".* FROM \"alpha\".\"biz_struct\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, bizStructPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in BizStructSlice")
	}

	*o = slice

	return nil
}

// BizStructExistsG checks if the BizStruct row exists.
func BizStructExistsG(ctx context.Context, iD int) (bool, error) {
	return BizStructExists(ctx, boil.GetContextDB(), iD)
}

// BizStructExists checks if the BizStruct row exists.
func BizStructExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"alpha\".\"biz_struct\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if biz_struct exists")
	}

	return exists, nil
}
