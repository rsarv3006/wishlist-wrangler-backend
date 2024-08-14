// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"wishlist-wrangler-api/ent/predicate"
	"wishlist-wrangler-api/ent/wishlist"
	"wishlist-wrangler-api/ent/wishlistsection"
	"wishlist-wrangler-api/ent/wishlisttemplatesection"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// WishlistSectionQuery is the builder for querying WishlistSection entities.
type WishlistSectionQuery struct {
	config
	ctx                         *QueryContext
	order                       []wishlistsection.OrderOption
	inters                      []Interceptor
	predicates                  []predicate.WishlistSection
	withWishlist                *WishlistQuery
	withWishlistTemplateSection *WishlistTemplateSectionQuery
	withFKs                     bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the WishlistSectionQuery builder.
func (wsq *WishlistSectionQuery) Where(ps ...predicate.WishlistSection) *WishlistSectionQuery {
	wsq.predicates = append(wsq.predicates, ps...)
	return wsq
}

// Limit the number of records to be returned by this query.
func (wsq *WishlistSectionQuery) Limit(limit int) *WishlistSectionQuery {
	wsq.ctx.Limit = &limit
	return wsq
}

// Offset to start from.
func (wsq *WishlistSectionQuery) Offset(offset int) *WishlistSectionQuery {
	wsq.ctx.Offset = &offset
	return wsq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (wsq *WishlistSectionQuery) Unique(unique bool) *WishlistSectionQuery {
	wsq.ctx.Unique = &unique
	return wsq
}

// Order specifies how the records should be ordered.
func (wsq *WishlistSectionQuery) Order(o ...wishlistsection.OrderOption) *WishlistSectionQuery {
	wsq.order = append(wsq.order, o...)
	return wsq
}

// QueryWishlist chains the current query on the "wishlist" edge.
func (wsq *WishlistSectionQuery) QueryWishlist() *WishlistQuery {
	query := (&WishlistClient{config: wsq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wsq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(wishlistsection.Table, wishlistsection.FieldID, selector),
			sqlgraph.To(wishlist.Table, wishlist.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, wishlistsection.WishlistTable, wishlistsection.WishlistColumn),
		)
		fromU = sqlgraph.SetNeighbors(wsq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryWishlistTemplateSection chains the current query on the "wishlistTemplateSection" edge.
func (wsq *WishlistSectionQuery) QueryWishlistTemplateSection() *WishlistTemplateSectionQuery {
	query := (&WishlistTemplateSectionClient{config: wsq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wsq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(wishlistsection.Table, wishlistsection.FieldID, selector),
			sqlgraph.To(wishlisttemplatesection.Table, wishlisttemplatesection.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, wishlistsection.WishlistTemplateSectionTable, wishlistsection.WishlistTemplateSectionColumn),
		)
		fromU = sqlgraph.SetNeighbors(wsq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first WishlistSection entity from the query.
// Returns a *NotFoundError when no WishlistSection was found.
func (wsq *WishlistSectionQuery) First(ctx context.Context) (*WishlistSection, error) {
	nodes, err := wsq.Limit(1).All(setContextOp(ctx, wsq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{wishlistsection.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (wsq *WishlistSectionQuery) FirstX(ctx context.Context) *WishlistSection {
	node, err := wsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first WishlistSection ID from the query.
// Returns a *NotFoundError when no WishlistSection ID was found.
func (wsq *WishlistSectionQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = wsq.Limit(1).IDs(setContextOp(ctx, wsq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{wishlistsection.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (wsq *WishlistSectionQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := wsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single WishlistSection entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one WishlistSection entity is found.
// Returns a *NotFoundError when no WishlistSection entities are found.
func (wsq *WishlistSectionQuery) Only(ctx context.Context) (*WishlistSection, error) {
	nodes, err := wsq.Limit(2).All(setContextOp(ctx, wsq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{wishlistsection.Label}
	default:
		return nil, &NotSingularError{wishlistsection.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (wsq *WishlistSectionQuery) OnlyX(ctx context.Context) *WishlistSection {
	node, err := wsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only WishlistSection ID in the query.
// Returns a *NotSingularError when more than one WishlistSection ID is found.
// Returns a *NotFoundError when no entities are found.
func (wsq *WishlistSectionQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = wsq.Limit(2).IDs(setContextOp(ctx, wsq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{wishlistsection.Label}
	default:
		err = &NotSingularError{wishlistsection.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (wsq *WishlistSectionQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := wsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of WishlistSections.
func (wsq *WishlistSectionQuery) All(ctx context.Context) ([]*WishlistSection, error) {
	ctx = setContextOp(ctx, wsq.ctx, ent.OpQueryAll)
	if err := wsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*WishlistSection, *WishlistSectionQuery]()
	return withInterceptors[[]*WishlistSection](ctx, wsq, qr, wsq.inters)
}

// AllX is like All, but panics if an error occurs.
func (wsq *WishlistSectionQuery) AllX(ctx context.Context) []*WishlistSection {
	nodes, err := wsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of WishlistSection IDs.
func (wsq *WishlistSectionQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if wsq.ctx.Unique == nil && wsq.path != nil {
		wsq.Unique(true)
	}
	ctx = setContextOp(ctx, wsq.ctx, ent.OpQueryIDs)
	if err = wsq.Select(wishlistsection.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (wsq *WishlistSectionQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := wsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (wsq *WishlistSectionQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, wsq.ctx, ent.OpQueryCount)
	if err := wsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, wsq, querierCount[*WishlistSectionQuery](), wsq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (wsq *WishlistSectionQuery) CountX(ctx context.Context) int {
	count, err := wsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (wsq *WishlistSectionQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, wsq.ctx, ent.OpQueryExist)
	switch _, err := wsq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (wsq *WishlistSectionQuery) ExistX(ctx context.Context) bool {
	exist, err := wsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the WishlistSectionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (wsq *WishlistSectionQuery) Clone() *WishlistSectionQuery {
	if wsq == nil {
		return nil
	}
	return &WishlistSectionQuery{
		config:                      wsq.config,
		ctx:                         wsq.ctx.Clone(),
		order:                       append([]wishlistsection.OrderOption{}, wsq.order...),
		inters:                      append([]Interceptor{}, wsq.inters...),
		predicates:                  append([]predicate.WishlistSection{}, wsq.predicates...),
		withWishlist:                wsq.withWishlist.Clone(),
		withWishlistTemplateSection: wsq.withWishlistTemplateSection.Clone(),
		// clone intermediate query.
		sql:  wsq.sql.Clone(),
		path: wsq.path,
	}
}

// WithWishlist tells the query-builder to eager-load the nodes that are connected to
// the "wishlist" edge. The optional arguments are used to configure the query builder of the edge.
func (wsq *WishlistSectionQuery) WithWishlist(opts ...func(*WishlistQuery)) *WishlistSectionQuery {
	query := (&WishlistClient{config: wsq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wsq.withWishlist = query
	return wsq
}

// WithWishlistTemplateSection tells the query-builder to eager-load the nodes that are connected to
// the "wishlistTemplateSection" edge. The optional arguments are used to configure the query builder of the edge.
func (wsq *WishlistSectionQuery) WithWishlistTemplateSection(opts ...func(*WishlistTemplateSectionQuery)) *WishlistSectionQuery {
	query := (&WishlistTemplateSectionClient{config: wsq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wsq.withWishlistTemplateSection = query
	return wsq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Type wishlistsection.Type `json:"type,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.WishlistSection.Query().
//		GroupBy(wishlistsection.FieldType).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (wsq *WishlistSectionQuery) GroupBy(field string, fields ...string) *WishlistSectionGroupBy {
	wsq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &WishlistSectionGroupBy{build: wsq}
	grbuild.flds = &wsq.ctx.Fields
	grbuild.label = wishlistsection.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Type wishlistsection.Type `json:"type,omitempty"`
//	}
//
//	client.WishlistSection.Query().
//		Select(wishlistsection.FieldType).
//		Scan(ctx, &v)
func (wsq *WishlistSectionQuery) Select(fields ...string) *WishlistSectionSelect {
	wsq.ctx.Fields = append(wsq.ctx.Fields, fields...)
	sbuild := &WishlistSectionSelect{WishlistSectionQuery: wsq}
	sbuild.label = wishlistsection.Label
	sbuild.flds, sbuild.scan = &wsq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a WishlistSectionSelect configured with the given aggregations.
func (wsq *WishlistSectionQuery) Aggregate(fns ...AggregateFunc) *WishlistSectionSelect {
	return wsq.Select().Aggregate(fns...)
}

func (wsq *WishlistSectionQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range wsq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, wsq); err != nil {
				return err
			}
		}
	}
	for _, f := range wsq.ctx.Fields {
		if !wishlistsection.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if wsq.path != nil {
		prev, err := wsq.path(ctx)
		if err != nil {
			return err
		}
		wsq.sql = prev
	}
	return nil
}

func (wsq *WishlistSectionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*WishlistSection, error) {
	var (
		nodes       = []*WishlistSection{}
		withFKs     = wsq.withFKs
		_spec       = wsq.querySpec()
		loadedTypes = [2]bool{
			wsq.withWishlist != nil,
			wsq.withWishlistTemplateSection != nil,
		}
	)
	if wsq.withWishlist != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, wishlistsection.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*WishlistSection).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &WishlistSection{config: wsq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, wsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := wsq.withWishlist; query != nil {
		if err := wsq.loadWishlist(ctx, query, nodes, nil,
			func(n *WishlistSection, e *Wishlist) { n.Edges.Wishlist = e }); err != nil {
			return nil, err
		}
	}
	if query := wsq.withWishlistTemplateSection; query != nil {
		if err := wsq.loadWishlistTemplateSection(ctx, query, nodes,
			func(n *WishlistSection) { n.Edges.WishlistTemplateSection = []*WishlistTemplateSection{} },
			func(n *WishlistSection, e *WishlistTemplateSection) {
				n.Edges.WishlistTemplateSection = append(n.Edges.WishlistTemplateSection, e)
			}); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (wsq *WishlistSectionQuery) loadWishlist(ctx context.Context, query *WishlistQuery, nodes []*WishlistSection, init func(*WishlistSection), assign func(*WishlistSection, *Wishlist)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*WishlistSection)
	for i := range nodes {
		if nodes[i].wishlist_sections == nil {
			continue
		}
		fk := *nodes[i].wishlist_sections
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(wishlist.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "wishlist_sections" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (wsq *WishlistSectionQuery) loadWishlistTemplateSection(ctx context.Context, query *WishlistTemplateSectionQuery, nodes []*WishlistSection, init func(*WishlistSection), assign func(*WishlistSection, *WishlistTemplateSection)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*WishlistSection)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.WishlistTemplateSection(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(wishlistsection.WishlistTemplateSectionColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.wishlist_section_wishlist_template_section
		if fk == nil {
			return fmt.Errorf(`foreign-key "wishlist_section_wishlist_template_section" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "wishlist_section_wishlist_template_section" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (wsq *WishlistSectionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := wsq.querySpec()
	_spec.Node.Columns = wsq.ctx.Fields
	if len(wsq.ctx.Fields) > 0 {
		_spec.Unique = wsq.ctx.Unique != nil && *wsq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, wsq.driver, _spec)
}

func (wsq *WishlistSectionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(wishlistsection.Table, wishlistsection.Columns, sqlgraph.NewFieldSpec(wishlistsection.FieldID, field.TypeUUID))
	_spec.From = wsq.sql
	if unique := wsq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if wsq.path != nil {
		_spec.Unique = true
	}
	if fields := wsq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, wishlistsection.FieldID)
		for i := range fields {
			if fields[i] != wishlistsection.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := wsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := wsq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := wsq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := wsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (wsq *WishlistSectionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(wsq.driver.Dialect())
	t1 := builder.Table(wishlistsection.Table)
	columns := wsq.ctx.Fields
	if len(columns) == 0 {
		columns = wishlistsection.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if wsq.sql != nil {
		selector = wsq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if wsq.ctx.Unique != nil && *wsq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range wsq.predicates {
		p(selector)
	}
	for _, p := range wsq.order {
		p(selector)
	}
	if offset := wsq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := wsq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WishlistSectionGroupBy is the group-by builder for WishlistSection entities.
type WishlistSectionGroupBy struct {
	selector
	build *WishlistSectionQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (wsgb *WishlistSectionGroupBy) Aggregate(fns ...AggregateFunc) *WishlistSectionGroupBy {
	wsgb.fns = append(wsgb.fns, fns...)
	return wsgb
}

// Scan applies the selector query and scans the result into the given value.
func (wsgb *WishlistSectionGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, wsgb.build.ctx, ent.OpQueryGroupBy)
	if err := wsgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WishlistSectionQuery, *WishlistSectionGroupBy](ctx, wsgb.build, wsgb, wsgb.build.inters, v)
}

func (wsgb *WishlistSectionGroupBy) sqlScan(ctx context.Context, root *WishlistSectionQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(wsgb.fns))
	for _, fn := range wsgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*wsgb.flds)+len(wsgb.fns))
		for _, f := range *wsgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*wsgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wsgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// WishlistSectionSelect is the builder for selecting fields of WishlistSection entities.
type WishlistSectionSelect struct {
	*WishlistSectionQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (wss *WishlistSectionSelect) Aggregate(fns ...AggregateFunc) *WishlistSectionSelect {
	wss.fns = append(wss.fns, fns...)
	return wss
}

// Scan applies the selector query and scans the result into the given value.
func (wss *WishlistSectionSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, wss.ctx, ent.OpQuerySelect)
	if err := wss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WishlistSectionQuery, *WishlistSectionSelect](ctx, wss.WishlistSectionQuery, wss, wss.inters, v)
}

func (wss *WishlistSectionSelect) sqlScan(ctx context.Context, root *WishlistSectionQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(wss.fns))
	for _, fn := range wss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*wss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
