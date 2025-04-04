// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"fs.io/asyncd/ent/enttask"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/errcode"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Common entgql types.
type (
	Cursor         = entgql.Cursor[int]
	PageInfo       = entgql.PageInfo[int]
	OrderDirection = entgql.OrderDirection
)

func orderFunc(o OrderDirection, field string) func(*sql.Selector) {
	if o == entgql.OrderDirectionDesc {
		return Desc(field)
	}
	return Asc(field)
}

const errInvalidPagination = "INVALID_PAGINATION"

func validateFirstLast(first, last *int) (err *gqlerror.Error) {
	switch {
	case first != nil && last != nil:
		err = &gqlerror.Error{
			Message: "Passing both `first` and `last` to paginate a connection is not supported.",
		}
	case first != nil && *first < 0:
		err = &gqlerror.Error{
			Message: "`first` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	case last != nil && *last < 0:
		err = &gqlerror.Error{
			Message: "`last` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	}
	return err
}

func collectedField(ctx context.Context, path ...string) *graphql.CollectedField {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return nil
	}
	field := fc.Field
	oc := graphql.GetOperationContext(ctx)
walk:
	for _, name := range path {
		for _, f := range graphql.CollectFields(oc, field.Selections, nil) {
			if f.Alias == name {
				field = f
				continue walk
			}
		}
		return nil
	}
	return &field
}

func hasCollectedField(ctx context.Context, path ...string) bool {
	if graphql.GetFieldContext(ctx) == nil {
		return true
	}
	return collectedField(ctx, path...) != nil
}

const (
	edgesField      = "edges"
	nodeField       = "node"
	pageInfoField   = "pageInfo"
	totalCountField = "totalCount"
)

func paginateLimit(first, last *int) int {
	var limit int
	if first != nil {
		limit = *first + 1
	} else if last != nil {
		limit = *last + 1
	}
	return limit
}

// EntTaskEdge is the edge representation of EntTask.
type EntTaskEdge struct {
	Node   *EntTask `json:"node"`
	Cursor Cursor   `json:"cursor"`
}

// EntTaskConnection is the connection containing edges to EntTask.
type EntTaskConnection struct {
	Edges      []*EntTaskEdge `json:"edges"`
	PageInfo   PageInfo       `json:"pageInfo"`
	TotalCount int            `json:"totalCount"`
}

func (c *EntTaskConnection) build(nodes []*EntTask, pager *enttaskPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *EntTask
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *EntTask {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *EntTask {
			return nodes[i]
		}
	}
	c.Edges = make([]*EntTaskEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &EntTaskEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// EntTaskPaginateOption enables pagination customization.
type EntTaskPaginateOption func(*enttaskPager) error

// WithEntTaskOrder configures pagination ordering.
func WithEntTaskOrder(order *EntTaskOrder) EntTaskPaginateOption {
	if order == nil {
		order = DefaultEntTaskOrder
	}
	o := *order
	return func(pager *enttaskPager) error {
		if err := o.Direction.Validate(); err != nil {
			return err
		}
		if o.Field == nil {
			o.Field = DefaultEntTaskOrder.Field
		}
		pager.order = &o
		return nil
	}
}

// WithEntTaskFilter configures pagination filter.
func WithEntTaskFilter(filter func(*EntTaskQuery) (*EntTaskQuery, error)) EntTaskPaginateOption {
	return func(pager *enttaskPager) error {
		if filter == nil {
			return errors.New("EntTaskQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type enttaskPager struct {
	reverse bool
	order   *EntTaskOrder
	filter  func(*EntTaskQuery) (*EntTaskQuery, error)
}

func newEntTaskPager(opts []EntTaskPaginateOption, reverse bool) (*enttaskPager, error) {
	pager := &enttaskPager{reverse: reverse}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	if pager.order == nil {
		pager.order = DefaultEntTaskOrder
	}
	return pager, nil
}

func (p *enttaskPager) applyFilter(query *EntTaskQuery) (*EntTaskQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *enttaskPager) toCursor(et *EntTask) Cursor {
	return p.order.Field.toCursor(et)
}

func (p *enttaskPager) applyCursors(query *EntTaskQuery, after, before *Cursor) (*EntTaskQuery, error) {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	for _, predicate := range entgql.CursorsPredicate(after, before, DefaultEntTaskOrder.Field.column, p.order.Field.column, direction) {
		query = query.Where(predicate)
	}
	return query, nil
}

func (p *enttaskPager) applyOrder(query *EntTaskQuery) *EntTaskQuery {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	query = query.Order(p.order.Field.toTerm(direction.OrderTermOption()))
	if p.order.Field != DefaultEntTaskOrder.Field {
		query = query.Order(DefaultEntTaskOrder.Field.toTerm(direction.OrderTermOption()))
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return query
}

func (p *enttaskPager) orderExpr(query *EntTaskQuery) sql.Querier {
	direction := p.order.Direction
	if p.reverse {
		direction = direction.Reverse()
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(p.order.Field.column)
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		b.Ident(p.order.Field.column).Pad().WriteString(string(direction))
		if p.order.Field != DefaultEntTaskOrder.Field {
			b.Comma().Ident(DefaultEntTaskOrder.Field.column).Pad().WriteString(string(direction))
		}
	})
}

// Paginate executes the query and returns a relay based cursor connection to EntTask.
func (et *EntTaskQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...EntTaskPaginateOption,
) (*EntTaskConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newEntTaskPager(opts, last != nil)
	if err != nil {
		return nil, err
	}
	if et, err = pager.applyFilter(et); err != nil {
		return nil, err
	}
	conn := &EntTaskConnection{Edges: []*EntTaskEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			c := et.Clone()
			c.ctx.Fields = nil
			if conn.TotalCount, err = c.Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}
	if et, err = pager.applyCursors(et, after, before); err != nil {
		return nil, err
	}
	limit := paginateLimit(first, last)
	if limit != 0 {
		et.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := et.collectField(ctx, limit == 1, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}
	et = pager.applyOrder(et)
	nodes, err := et.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

// EntTaskOrderField defines the ordering field of EntTask.
type EntTaskOrderField struct {
	// Value extracts the ordering value from the given EntTask.
	Value    func(*EntTask) (ent.Value, error)
	column   string // field or computed.
	toTerm   func(...sql.OrderTermOption) enttask.OrderOption
	toCursor func(*EntTask) Cursor
}

// EntTaskOrder defines the ordering of EntTask.
type EntTaskOrder struct {
	Direction OrderDirection     `json:"direction"`
	Field     *EntTaskOrderField `json:"field"`
}

// DefaultEntTaskOrder is the default ordering of EntTask.
var DefaultEntTaskOrder = &EntTaskOrder{
	Direction: entgql.OrderDirectionAsc,
	Field: &EntTaskOrderField{
		Value: func(et *EntTask) (ent.Value, error) {
			return et.ID, nil
		},
		column: enttask.FieldID,
		toTerm: enttask.ByID,
		toCursor: func(et *EntTask) Cursor {
			return Cursor{ID: et.ID}
		},
	},
}

// ToEdge converts EntTask into EntTaskEdge.
func (et *EntTask) ToEdge(order *EntTaskOrder) *EntTaskEdge {
	if order == nil {
		order = DefaultEntTaskOrder
	}
	return &EntTaskEdge{
		Node:   et,
		Cursor: order.Field.toCursor(et),
	}
}
