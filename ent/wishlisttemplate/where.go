// Code generated by ent, DO NOT EDIT.

package wishlisttemplate

import (
	"time"
	"wishlist-wrangler-api/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.WishlistTemplate {
	return predicate.WishlistTemplate(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.WishlistTemplate {
	return predicate.WishlistTemplate(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.WishlistTemplate {
	return predicate.WishlistTemplate(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.WishlistTemplate {
	return predicate.WishlistTemplate(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.WishlistTemplate {
	return predicate.WishlistTemplate(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.WishlistTemplate {
	return predicate.WishlistTemplate(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.WishlistTemplate {
	return predicate.WishlistTemplate(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.WishlistTemplate {
	return predicate.WishlistTemplate(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.WishlistTemplate {
	return predicate.WishlistTemplate(sql.FieldLTE(FieldID, id))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.WishlistTemplate {
	return predicate.WishlistTemplate(sql.FieldEQ(FieldTitle, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.WishlistTemplate {
	return predicate.WishlistTemplate(sql.FieldEQ(FieldCreatedAt, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.WishlistTemplate {
	return predicate.WishlistTemplate(sql.FieldEQ(FieldDescription, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.WishlistTemplate {
	return predicate.WishlistTemplate(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.WishlistTemplate {
	return predicate.WishlistTemplate(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.WishlistTemplate {
	return predicate.WishlistTemplate(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.WishlistTemplate {
	return predicate.WishlistTemplate(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.WishlistTemplate {
	return predicate.WishlistTemplate(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.WishlistTemplate {
	return predicate.WishlistTemplate(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.WishlistTemplate {
	return predicate.WishlistTemplate(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.WishlistTemplate {
	return predicate.WishlistTemplate(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.WishlistTemplate {
	return predicate.WishlistTemplate(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.WishlistTemplate {
	return predicate.WishlistTemplate(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.WishlistTemplate {
	return predicate.WishlistTemplate(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.WishlistTemplate {
	return predicate.WishlistTemplate(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.WishlistTemplate {
	return predicate.WishlistTemplate(sql.FieldContainsFold(FieldTitle, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.WishlistTemplate {
	return predicate.WishlistTemplate(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.WishlistTemplate {
	return predicate.WishlistTemplate(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.WishlistTemplate {
	return predicate.WishlistTemplate(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.WishlistTemplate {
	return predicate.WishlistTemplate(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.WishlistTemplate {
	return predicate.WishlistTemplate(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.WishlistTemplate {
	return predicate.WishlistTemplate(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.WishlistTemplate {
	return predicate.WishlistTemplate(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.WishlistTemplate {
	return predicate.WishlistTemplate(sql.FieldLTE(FieldCreatedAt, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.WishlistTemplate {
	return predicate.WishlistTemplate(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.WishlistTemplate {
	return predicate.WishlistTemplate(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.WishlistTemplate {
	return predicate.WishlistTemplate(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.WishlistTemplate {
	return predicate.WishlistTemplate(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.WishlistTemplate {
	return predicate.WishlistTemplate(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.WishlistTemplate {
	return predicate.WishlistTemplate(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.WishlistTemplate {
	return predicate.WishlistTemplate(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.WishlistTemplate {
	return predicate.WishlistTemplate(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.WishlistTemplate {
	return predicate.WishlistTemplate(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.WishlistTemplate {
	return predicate.WishlistTemplate(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.WishlistTemplate {
	return predicate.WishlistTemplate(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.WishlistTemplate {
	return predicate.WishlistTemplate(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.WishlistTemplate {
	return predicate.WishlistTemplate(sql.FieldContainsFold(FieldDescription, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.WishlistTemplate {
	return predicate.WishlistTemplate(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.WishlistTemplate {
	return predicate.WishlistTemplate(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.WishlistTemplate {
	return predicate.WishlistTemplate(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.WishlistTemplate {
	return predicate.WishlistTemplate(sql.FieldNotIn(FieldStatus, vs...))
}

// HasCreatorId applies the HasEdge predicate on the "creatorId" edge.
func HasCreatorId() predicate.WishlistTemplate {
	return predicate.WishlistTemplate(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CreatorIdTable, CreatorIdColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCreatorIdWith applies the HasEdge predicate on the "creatorId" edge with a given conditions (other predicates).
func HasCreatorIdWith(preds ...predicate.User) predicate.WishlistTemplate {
	return predicate.WishlistTemplate(func(s *sql.Selector) {
		step := newCreatorIdStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSections applies the HasEdge predicate on the "sections" edge.
func HasSections() predicate.WishlistTemplate {
	return predicate.WishlistTemplate(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, SectionsTable, SectionsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSectionsWith applies the HasEdge predicate on the "sections" edge with a given conditions (other predicates).
func HasSectionsWith(preds ...predicate.WishlistTemplateSection) predicate.WishlistTemplate {
	return predicate.WishlistTemplate(func(s *sql.Selector) {
		step := newSectionsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.WishlistTemplate) predicate.WishlistTemplate {
	return predicate.WishlistTemplate(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.WishlistTemplate) predicate.WishlistTemplate {
	return predicate.WishlistTemplate(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.WishlistTemplate) predicate.WishlistTemplate {
	return predicate.WishlistTemplate(sql.NotPredicates(p))
}
