// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/NpoolPlatform/third-login-gateway/pkg/db/ent/auth"
	"github.com/NpoolPlatform/third-login-gateway/pkg/db/ent/thirdparty"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entql"
	"entgo.io/ent/schema/field"
)

// schemaGraph holds a representation of ent/schema at runtime.
var schemaGraph = func() *sqlgraph.Schema {
	graph := &sqlgraph.Schema{Nodes: make([]*sqlgraph.Node, 2)}
	graph.Nodes[0] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   auth.Table,
			Columns: auth.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: auth.FieldID,
			},
		},
		Type: "Auth",
		Fields: map[string]*sqlgraph.FieldSpec{
			auth.FieldCreatedAt:    {Type: field.TypeUint32, Column: auth.FieldCreatedAt},
			auth.FieldUpdatedAt:    {Type: field.TypeUint32, Column: auth.FieldUpdatedAt},
			auth.FieldDeletedAt:    {Type: field.TypeUint32, Column: auth.FieldDeletedAt},
			auth.FieldAppID:        {Type: field.TypeUUID, Column: auth.FieldAppID},
			auth.FieldThirdPartyID: {Type: field.TypeUUID, Column: auth.FieldThirdPartyID},
			auth.FieldAppKey:       {Type: field.TypeString, Column: auth.FieldAppKey},
			auth.FieldAppSecret:    {Type: field.TypeString, Column: auth.FieldAppSecret},
			auth.FieldRedirectURL:  {Type: field.TypeString, Column: auth.FieldRedirectURL},
		},
	}
	graph.Nodes[1] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   thirdparty.Table,
			Columns: thirdparty.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: thirdparty.FieldID,
			},
		},
		Type: "ThirdParty",
		Fields: map[string]*sqlgraph.FieldSpec{
			thirdparty.FieldCreatedAt: {Type: field.TypeUint32, Column: thirdparty.FieldCreatedAt},
			thirdparty.FieldUpdatedAt: {Type: field.TypeUint32, Column: thirdparty.FieldUpdatedAt},
			thirdparty.FieldDeletedAt: {Type: field.TypeUint32, Column: thirdparty.FieldDeletedAt},
			thirdparty.FieldBrandName: {Type: field.TypeString, Column: thirdparty.FieldBrandName},
			thirdparty.FieldLogo:      {Type: field.TypeString, Column: thirdparty.FieldLogo},
			thirdparty.FieldDomain:    {Type: field.TypeString, Column: thirdparty.FieldDomain},
		},
	}
	return graph
}()

// predicateAdder wraps the addPredicate method.
// All update, update-one and query builders implement this interface.
type predicateAdder interface {
	addPredicate(func(s *sql.Selector))
}

// addPredicate implements the predicateAdder interface.
func (aq *AuthQuery) addPredicate(pred func(s *sql.Selector)) {
	aq.predicates = append(aq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the AuthQuery builder.
func (aq *AuthQuery) Filter() *AuthFilter {
	return &AuthFilter{aq}
}

// addPredicate implements the predicateAdder interface.
func (m *AuthMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the AuthMutation builder.
func (m *AuthMutation) Filter() *AuthFilter {
	return &AuthFilter{m}
}

// AuthFilter provides a generic filtering capability at runtime for AuthQuery.
type AuthFilter struct {
	predicateAdder
}

// Where applies the entql predicate on the query filter.
func (f *AuthFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[0].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *AuthFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(auth.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *AuthFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(auth.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *AuthFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(auth.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *AuthFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(auth.FieldDeletedAt))
}

// WhereAppID applies the entql [16]byte predicate on the app_id field.
func (f *AuthFilter) WhereAppID(p entql.ValueP) {
	f.Where(p.Field(auth.FieldAppID))
}

// WhereThirdPartyID applies the entql [16]byte predicate on the third_party_id field.
func (f *AuthFilter) WhereThirdPartyID(p entql.ValueP) {
	f.Where(p.Field(auth.FieldThirdPartyID))
}

// WhereAppKey applies the entql string predicate on the app_key field.
func (f *AuthFilter) WhereAppKey(p entql.StringP) {
	f.Where(p.Field(auth.FieldAppKey))
}

// WhereAppSecret applies the entql string predicate on the app_secret field.
func (f *AuthFilter) WhereAppSecret(p entql.StringP) {
	f.Where(p.Field(auth.FieldAppSecret))
}

// WhereRedirectURL applies the entql string predicate on the redirect_url field.
func (f *AuthFilter) WhereRedirectURL(p entql.StringP) {
	f.Where(p.Field(auth.FieldRedirectURL))
}

// addPredicate implements the predicateAdder interface.
func (tpq *ThirdPartyQuery) addPredicate(pred func(s *sql.Selector)) {
	tpq.predicates = append(tpq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the ThirdPartyQuery builder.
func (tpq *ThirdPartyQuery) Filter() *ThirdPartyFilter {
	return &ThirdPartyFilter{tpq}
}

// addPredicate implements the predicateAdder interface.
func (m *ThirdPartyMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the ThirdPartyMutation builder.
func (m *ThirdPartyMutation) Filter() *ThirdPartyFilter {
	return &ThirdPartyFilter{m}
}

// ThirdPartyFilter provides a generic filtering capability at runtime for ThirdPartyQuery.
type ThirdPartyFilter struct {
	predicateAdder
}

// Where applies the entql predicate on the query filter.
func (f *ThirdPartyFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[1].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *ThirdPartyFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(thirdparty.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *ThirdPartyFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(thirdparty.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *ThirdPartyFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(thirdparty.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *ThirdPartyFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(thirdparty.FieldDeletedAt))
}

// WhereBrandName applies the entql string predicate on the brand_name field.
func (f *ThirdPartyFilter) WhereBrandName(p entql.StringP) {
	f.Where(p.Field(thirdparty.FieldBrandName))
}

// WhereLogo applies the entql string predicate on the logo field.
func (f *ThirdPartyFilter) WhereLogo(p entql.StringP) {
	f.Where(p.Field(thirdparty.FieldLogo))
}

// WhereDomain applies the entql string predicate on the domain field.
func (f *ThirdPartyFilter) WhereDomain(p entql.StringP) {
	f.Where(p.Field(thirdparty.FieldDomain))
}
