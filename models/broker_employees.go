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

// BrokerEmployee is an object representing the database table.
type BrokerEmployee struct {
	ID       int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	Email    string    `boil:"email" json:"email" toml:"email" yaml:"email"`
	RoleID   null.Int  `boil:"role_id" json:"role_id,omitempty" toml:"role_id" yaml:"role_id,omitempty"`
	Title    string    `boil:"title" json:"title" toml:"title" yaml:"title"`
	Forename string    `boil:"forename" json:"forename" toml:"forename" yaml:"forename"`
	Surname  string    `boil:"surname" json:"surname" toml:"surname" yaml:"surname"`
	Active   bool      `boil:"active" json:"active" toml:"active" yaml:"active"`
	Created  time.Time `boil:"created" json:"created" toml:"created" yaml:"created"`

	R *brokerEmployeeR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L brokerEmployeeL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var BrokerEmployeeColumns = struct {
	ID       string
	Email    string
	RoleID   string
	Title    string
	Forename string
	Surname  string
	Active   string
	Created  string
}{
	ID:       "id",
	Email:    "email",
	RoleID:   "role_id",
	Title:    "title",
	Forename: "forename",
	Surname:  "surname",
	Active:   "active",
	Created:  "created",
}

var BrokerEmployeeTableColumns = struct {
	ID       string
	Email    string
	RoleID   string
	Title    string
	Forename string
	Surname  string
	Active   string
	Created  string
}{
	ID:       "broker_employees.id",
	Email:    "broker_employees.email",
	RoleID:   "broker_employees.role_id",
	Title:    "broker_employees.title",
	Forename: "broker_employees.forename",
	Surname:  "broker_employees.surname",
	Active:   "broker_employees.active",
	Created:  "broker_employees.created",
}

// Generated where

type whereHelpernull_Int struct{ field string }

func (w whereHelpernull_Int) EQ(x null.Int) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Int) NEQ(x null.Int) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Int) LT(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Int) LTE(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Int) GT(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Int) GTE(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelpernull_Int) IN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelpernull_Int) NIN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

func (w whereHelpernull_Int) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Int) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

type whereHelperbool struct{ field string }

func (w whereHelperbool) EQ(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperbool) NEQ(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperbool) LT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperbool) LTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperbool) GT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperbool) GTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var BrokerEmployeeWhere = struct {
	ID       whereHelperint
	Email    whereHelperstring
	RoleID   whereHelpernull_Int
	Title    whereHelperstring
	Forename whereHelperstring
	Surname  whereHelperstring
	Active   whereHelperbool
	Created  whereHelpertime_Time
}{
	ID:       whereHelperint{field: "\"alpha\".\"broker_employees\".\"id\""},
	Email:    whereHelperstring{field: "\"alpha\".\"broker_employees\".\"email\""},
	RoleID:   whereHelpernull_Int{field: "\"alpha\".\"broker_employees\".\"role_id\""},
	Title:    whereHelperstring{field: "\"alpha\".\"broker_employees\".\"title\""},
	Forename: whereHelperstring{field: "\"alpha\".\"broker_employees\".\"forename\""},
	Surname:  whereHelperstring{field: "\"alpha\".\"broker_employees\".\"surname\""},
	Active:   whereHelperbool{field: "\"alpha\".\"broker_employees\".\"active\""},
	Created:  whereHelpertime_Time{field: "\"alpha\".\"broker_employees\".\"created\""},
}

// BrokerEmployeeRels is where relationship names are stored.
var BrokerEmployeeRels = struct {
	Role string
}{
	Role: "Role",
}

// brokerEmployeeR is where relationships are stored.
type brokerEmployeeR struct {
	Role *BrokerRole `boil:"Role" json:"Role" toml:"Role" yaml:"Role"`
}

// NewStruct creates a new relationship struct
func (*brokerEmployeeR) NewStruct() *brokerEmployeeR {
	return &brokerEmployeeR{}
}

func (r *brokerEmployeeR) GetRole() *BrokerRole {
	if r == nil {
		return nil
	}
	return r.Role
}

// brokerEmployeeL is where Load methods for each relationship are stored.
type brokerEmployeeL struct{}

var (
	brokerEmployeeAllColumns            = []string{"id", "email", "role_id", "title", "forename", "surname", "active", "created"}
	brokerEmployeeColumnsWithoutDefault = []string{"email", "title", "forename", "surname"}
	brokerEmployeeColumnsWithDefault    = []string{"id", "role_id", "active", "created"}
	brokerEmployeePrimaryKeyColumns     = []string{"id"}
	brokerEmployeeGeneratedColumns      = []string{"id"}
)

type (
	// BrokerEmployeeSlice is an alias for a slice of pointers to BrokerEmployee.
	// This should almost always be used instead of []BrokerEmployee.
	BrokerEmployeeSlice []*BrokerEmployee

	brokerEmployeeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	brokerEmployeeType                 = reflect.TypeOf(&BrokerEmployee{})
	brokerEmployeeMapping              = queries.MakeStructMapping(brokerEmployeeType)
	brokerEmployeePrimaryKeyMapping, _ = queries.BindMapping(brokerEmployeeType, brokerEmployeeMapping, brokerEmployeePrimaryKeyColumns)
	brokerEmployeeInsertCacheMut       sync.RWMutex
	brokerEmployeeInsertCache          = make(map[string]insertCache)
	brokerEmployeeUpdateCacheMut       sync.RWMutex
	brokerEmployeeUpdateCache          = make(map[string]updateCache)
	brokerEmployeeUpsertCacheMut       sync.RWMutex
	brokerEmployeeUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// OneG returns a single brokerEmployee record from the query using the global executor.
func (q brokerEmployeeQuery) OneG(ctx context.Context) (*BrokerEmployee, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single brokerEmployee record from the query.
func (q brokerEmployeeQuery) One(ctx context.Context, exec boil.ContextExecutor) (*BrokerEmployee, error) {
	o := &BrokerEmployee{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for broker_employees")
	}

	return o, nil
}

// AllG returns all BrokerEmployee records from the query using the global executor.
func (q brokerEmployeeQuery) AllG(ctx context.Context) (BrokerEmployeeSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all BrokerEmployee records from the query.
func (q brokerEmployeeQuery) All(ctx context.Context, exec boil.ContextExecutor) (BrokerEmployeeSlice, error) {
	var o []*BrokerEmployee

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to BrokerEmployee slice")
	}

	return o, nil
}

// CountG returns the count of all BrokerEmployee records in the query using the global executor
func (q brokerEmployeeQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all BrokerEmployee records in the query.
func (q brokerEmployeeQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count broker_employees rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q brokerEmployeeQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q brokerEmployeeQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if broker_employees exists")
	}

	return count > 0, nil
}

// Role pointed to by the foreign key.
func (o *BrokerEmployee) Role(mods ...qm.QueryMod) brokerRoleQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.RoleID),
	}

	queryMods = append(queryMods, mods...)

	return BrokerRoles(queryMods...)
}

// LoadRole allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (brokerEmployeeL) LoadRole(ctx context.Context, e boil.ContextExecutor, singular bool, maybeBrokerEmployee interface{}, mods queries.Applicator) error {
	var slice []*BrokerEmployee
	var object *BrokerEmployee

	if singular {
		var ok bool
		object, ok = maybeBrokerEmployee.(*BrokerEmployee)
		if !ok {
			object = new(BrokerEmployee)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeBrokerEmployee)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeBrokerEmployee))
			}
		}
	} else {
		s, ok := maybeBrokerEmployee.(*[]*BrokerEmployee)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeBrokerEmployee)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeBrokerEmployee))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &brokerEmployeeR{}
		}
		if !queries.IsNil(object.RoleID) {
			args = append(args, object.RoleID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &brokerEmployeeR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.RoleID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.RoleID) {
				args = append(args, obj.RoleID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`alpha.broker_roles`),
		qm.WhereIn(`alpha.broker_roles.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load BrokerRole")
	}

	var resultSlice []*BrokerRole
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice BrokerRole")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for broker_roles")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for broker_roles")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Role = foreign
		if foreign.R == nil {
			foreign.R = &brokerRoleR{}
		}
		foreign.R.RoleBrokerEmployees = append(foreign.R.RoleBrokerEmployees, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.RoleID, foreign.ID) {
				local.R.Role = foreign
				if foreign.R == nil {
					foreign.R = &brokerRoleR{}
				}
				foreign.R.RoleBrokerEmployees = append(foreign.R.RoleBrokerEmployees, local)
				break
			}
		}
	}

	return nil
}

// SetRoleG of the brokerEmployee to the related item.
// Sets o.R.Role to related.
// Adds o to related.R.RoleBrokerEmployees.
// Uses the global database handle.
func (o *BrokerEmployee) SetRoleG(ctx context.Context, insert bool, related *BrokerRole) error {
	return o.SetRole(ctx, boil.GetContextDB(), insert, related)
}

// SetRole of the brokerEmployee to the related item.
// Sets o.R.Role to related.
// Adds o to related.R.RoleBrokerEmployees.
func (o *BrokerEmployee) SetRole(ctx context.Context, exec boil.ContextExecutor, insert bool, related *BrokerRole) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"alpha\".\"broker_employees\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"role_id"}),
		strmangle.WhereClause("\"", "\"", 2, brokerEmployeePrimaryKeyColumns),
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

	queries.Assign(&o.RoleID, related.ID)
	if o.R == nil {
		o.R = &brokerEmployeeR{
			Role: related,
		}
	} else {
		o.R.Role = related
	}

	if related.R == nil {
		related.R = &brokerRoleR{
			RoleBrokerEmployees: BrokerEmployeeSlice{o},
		}
	} else {
		related.R.RoleBrokerEmployees = append(related.R.RoleBrokerEmployees, o)
	}

	return nil
}

// RemoveRoleG relationship.
// Sets o.R.Role to nil.
// Removes o from all passed in related items' relationships struct.
// Uses the global database handle.
func (o *BrokerEmployee) RemoveRoleG(ctx context.Context, related *BrokerRole) error {
	return o.RemoveRole(ctx, boil.GetContextDB(), related)
}

// RemoveRole relationship.
// Sets o.R.Role to nil.
// Removes o from all passed in related items' relationships struct.
func (o *BrokerEmployee) RemoveRole(ctx context.Context, exec boil.ContextExecutor, related *BrokerRole) error {
	var err error

	queries.SetScanner(&o.RoleID, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("role_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	if o.R != nil {
		o.R.Role = nil
	}
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.RoleBrokerEmployees {
		if queries.Equal(o.RoleID, ri.RoleID) {
			continue
		}

		ln := len(related.R.RoleBrokerEmployees)
		if ln > 1 && i < ln-1 {
			related.R.RoleBrokerEmployees[i] = related.R.RoleBrokerEmployees[ln-1]
		}
		related.R.RoleBrokerEmployees = related.R.RoleBrokerEmployees[:ln-1]
		break
	}
	return nil
}

// BrokerEmployees retrieves all the records using an executor.
func BrokerEmployees(mods ...qm.QueryMod) brokerEmployeeQuery {
	mods = append(mods, qm.From("\"alpha\".\"broker_employees\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"alpha\".\"broker_employees\".*"})
	}

	return brokerEmployeeQuery{q}
}

// FindBrokerEmployeeG retrieves a single record by ID.
func FindBrokerEmployeeG(ctx context.Context, iD int, selectCols ...string) (*BrokerEmployee, error) {
	return FindBrokerEmployee(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindBrokerEmployee retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindBrokerEmployee(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*BrokerEmployee, error) {
	brokerEmployeeObj := &BrokerEmployee{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"alpha\".\"broker_employees\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, brokerEmployeeObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from broker_employees")
	}

	return brokerEmployeeObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *BrokerEmployee) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *BrokerEmployee) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no broker_employees provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(brokerEmployeeColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	brokerEmployeeInsertCacheMut.RLock()
	cache, cached := brokerEmployeeInsertCache[key]
	brokerEmployeeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			brokerEmployeeAllColumns,
			brokerEmployeeColumnsWithDefault,
			brokerEmployeeColumnsWithoutDefault,
			nzDefaults,
		)
		wl = strmangle.SetComplement(wl, brokerEmployeeGeneratedColumns)

		cache.valueMapping, err = queries.BindMapping(brokerEmployeeType, brokerEmployeeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(brokerEmployeeType, brokerEmployeeMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"alpha\".\"broker_employees\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"alpha\".\"broker_employees\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into broker_employees")
	}

	if !cached {
		brokerEmployeeInsertCacheMut.Lock()
		brokerEmployeeInsertCache[key] = cache
		brokerEmployeeInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single BrokerEmployee record using the global executor.
// See Update for more documentation.
func (o *BrokerEmployee) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the BrokerEmployee.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *BrokerEmployee) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	brokerEmployeeUpdateCacheMut.RLock()
	cache, cached := brokerEmployeeUpdateCache[key]
	brokerEmployeeUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			brokerEmployeeAllColumns,
			brokerEmployeePrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, brokerEmployeeGeneratedColumns)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update broker_employees, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"alpha\".\"broker_employees\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, brokerEmployeePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(brokerEmployeeType, brokerEmployeeMapping, append(wl, brokerEmployeePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update broker_employees row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for broker_employees")
	}

	if !cached {
		brokerEmployeeUpdateCacheMut.Lock()
		brokerEmployeeUpdateCache[key] = cache
		brokerEmployeeUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (q brokerEmployeeQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q brokerEmployeeQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for broker_employees")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for broker_employees")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o BrokerEmployeeSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o BrokerEmployeeSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), brokerEmployeePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"alpha\".\"broker_employees\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, brokerEmployeePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in brokerEmployee slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all brokerEmployee")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *BrokerEmployee) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *BrokerEmployee) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no broker_employees provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(brokerEmployeeColumnsWithDefault, o)

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

	brokerEmployeeUpsertCacheMut.RLock()
	cache, cached := brokerEmployeeUpsertCache[key]
	brokerEmployeeUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			brokerEmployeeAllColumns,
			brokerEmployeeColumnsWithDefault,
			brokerEmployeeColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			brokerEmployeeAllColumns,
			brokerEmployeePrimaryKeyColumns,
		)

		insert = strmangle.SetComplement(insert, brokerEmployeeGeneratedColumns)
		update = strmangle.SetComplement(update, brokerEmployeeGeneratedColumns)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert broker_employees, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(brokerEmployeePrimaryKeyColumns))
			copy(conflict, brokerEmployeePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"alpha\".\"broker_employees\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(brokerEmployeeType, brokerEmployeeMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(brokerEmployeeType, brokerEmployeeMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert broker_employees")
	}

	if !cached {
		brokerEmployeeUpsertCacheMut.Lock()
		brokerEmployeeUpsertCache[key] = cache
		brokerEmployeeUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteG deletes a single BrokerEmployee record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *BrokerEmployee) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single BrokerEmployee record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *BrokerEmployee) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no BrokerEmployee provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), brokerEmployeePrimaryKeyMapping)
	sql := "DELETE FROM \"alpha\".\"broker_employees\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from broker_employees")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for broker_employees")
	}

	return rowsAff, nil
}

func (q brokerEmployeeQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q brokerEmployeeQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no brokerEmployeeQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from broker_employees")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for broker_employees")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o BrokerEmployeeSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o BrokerEmployeeSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), brokerEmployeePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"alpha\".\"broker_employees\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, brokerEmployeePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from brokerEmployee slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for broker_employees")
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *BrokerEmployee) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no BrokerEmployee provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *BrokerEmployee) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindBrokerEmployee(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *BrokerEmployeeSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty BrokerEmployeeSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *BrokerEmployeeSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := BrokerEmployeeSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), brokerEmployeePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"alpha\".\"broker_employees\".* FROM \"alpha\".\"broker_employees\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, brokerEmployeePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in BrokerEmployeeSlice")
	}

	*o = slice

	return nil
}

// BrokerEmployeeExistsG checks if the BrokerEmployee row exists.
func BrokerEmployeeExistsG(ctx context.Context, iD int) (bool, error) {
	return BrokerEmployeeExists(ctx, boil.GetContextDB(), iD)
}

// BrokerEmployeeExists checks if the BrokerEmployee row exists.
func BrokerEmployeeExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"alpha\".\"broker_employees\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if broker_employees exists")
	}

	return exists, nil
}
