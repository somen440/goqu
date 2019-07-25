package goqu

import (
	"fmt"
	"time"

	"github.com/doug-martin/goqu/v8/exp"
)

type (
	SQLFragmentType   int
	SQLDialectOptions struct {
		// Set to true if the dialect supports ORDER BY expressions in DELETE statements (DEFAULT=false)
		SupportsOrderByOnDelete bool
		// Set to true if the dialect supports ORDER BY expressions in UPDATE statements (DEFAULT=false)
		SupportsOrderByOnUpdate bool
		// Set to true if the dialect supports LIMIT expressions in DELETE statements (DEFAULT=false)
		SupportsLimitOnDelete bool
		// Set to true if the dialect supports LIMIT expressions in UPDATE statements (DEFAULT=false)
		SupportsLimitOnUpdate bool
		// Set to true if the dialect supports RETURN expressions (DEFAULT=true)
		SupportsReturn bool
		// Set to true if the dialect supports Conflict Target (DEFAULT=true)
		SupportsConflictTarget bool
		// Set to true if the dialect supports Conflict Target (DEFAULT=true)
		SupportsConflictUpdateWhere bool
		// Set to true if the dialect supports Insert Ignore syntax (DEFAULT=false)
		SupportsInsertIgnoreSyntax bool
		// Set to true if the dialect supports Common Table Expressions (DEFAULT=true)
		SupportsWithCTE bool
		// Set to true if the dialect supports recursive Common Table Expressions (DEFAULT=true)
		SupportsWithCTERecursive bool
		// Set to true if multiple tables are supported in UPDATE statement. (DEFAULT=true)
		SupportsMultipleUpdateTables bool
		// Set to false if the dialect does not require expressions to be wrapped in parens (DEFAULT=true)
		WrapCompoundsInParens bool

		// Set to true if the dialect requires join tables in UPDATE to be in a FROM clause (DEFAULT=true).
		UseFromClauseForMultipleUpdateTables bool

		// The UPDATE fragment to use when generating sql. (DEFAULT=[]byte("UPDATE"))
		UpdateClause []byte
		// The INSERT fragment to use when generating sql. (DEFAULT=[]byte("INSERT INTO"))
		InsertClause []byte
		// The INSERT IGNORE INTO fragment to use when generating sql. (DEFAULT=[]byte("INSERT IGNORE INTO"))
		InsertIgnoreClause []byte
		// The SELECT fragment to use when generating sql. (DEFAULT=[]byte("SELECT"))
		SelectClause []byte
		// The DELETE fragment to use when generating sql. (DEFAULT=[]byte("DELETE"))
		DeleteClause []byte
		// The TRUNCATE fragment to use when generating sql. (DEFAULT=[]byte("TRUNCATE"))
		TruncateClause []byte
		// The WITH fragment to use when generating sql. (DEFAULT=[]byte("WITH "))
		WithFragment []byte
		// The RECURSIVE fragment to use when generating sql (after WITH). (DEFAULT=[]byte("RECURSIVE "))
		RecursiveFragment []byte
		// The CASCADE fragment to use when generating sql. (DEFAULT=[]byte(" CASCADE"))
		CascadeFragment []byte
		// The RESTRICT fragment to use when generating sql. (DEFAULT=[]byte(" RESTRICT"))
		RestrictFragment []byte
		// The SQL fragment to use when generating insert sql and using
		// DEFAULT VALUES (e.g. postgres="DEFAULT VALUES", mysql="", sqlite3=""). (DEFAULT=[]byte(" DEFAULT VALUES"))
		DefaultValuesFragment []byte
		// The SQL fragment to use when generating insert sql and listing columns using a VALUES clause
		// (DEFAULT=[]byte(" VALUES "))
		ValuesFragment []byte
		// The SQL fragment to use when generating truncate sql and using the IDENTITY clause
		// (DEFAULT=[]byte(" IDENTITY"))
		IdentityFragment []byte
		// The SQL fragment to use when generating update sql and using the SET clause (DEFAULT=[]byte(" SET "))
		SetFragment []byte
		// The SQL DISTINCT keyword (DEFAULT=[]byte(" DISTINCT "))
		DistinctFragment []byte
		// The SQL RETURNING clause (DEFAULT=[]byte(" RETURNING "))
		ReturningFragment []byte
		// The SQL FROM clause fragment (DEFAULT=[]byte(" FROM"))
		FromFragment []byte
		// The SQL USING join clause fragment (DEFAULT=[]byte(" USING "))
		UsingFragment []byte
		// The SQL ON join clause fragment (DEFAULT=[]byte(" ON "))
		OnFragment []byte
		// The SQL WHERE clause fragment (DEFAULT=[]byte(" WHERE "))
		WhereFragment []byte
		// The SQL GROUP BY clause fragment(DEFAULT=[]byte(" GROUP BY "))
		GroupByFragment []byte
		// The SQL HAVING clause fragment(DELiFAULT=[]byte(" HAVING "))
		HavingFragment []byte
		// The SQL ORDER BY clause fragment(DEFAULT=[]byte(" ORDER BY "))
		OrderByFragment []byte
		// The SQL LIMIT BY clause fragment(DEFAULT=[]byte(" LIMIT "))
		LimitFragment []byte
		// The SQL OFFSET BY clause fragment(DEFAULT=[]byte(" OFFSET "))
		OffsetFragment []byte
		// The SQL FOR UPDATE fragment(DEFAULT=[]byte(" FOR UPDATE "))
		ForUpdateFragment []byte
		// The SQL FOR NO KEY UPDATE fragment(DEFAULT=[]byte(" FOR NO KEY UPDATE "))
		ForNoKeyUpdateFragment []byte
		// The SQL FOR SHARE fragment(DEFAULT=[]byte(" FOR SHARE "))
		ForShareFragment []byte
		// The SQL FOR KEY SHARE fragment(DEFAULT=[]byte(" FOR KEY SHARE "))
		ForKeyShareFragment []byte
		// The SQL NOWAIT fragment(DEFAULT=[]byte("NOWAIT"))
		NowaitFragment []byte
		// The SQL SKIP LOCKED fragment(DEFAULT=[]byte("SKIP LOCKED"))
		SkipLockedFragment []byte
		// The SQL AS fragment when aliasing an Expression(DEFAULT=[]byte(" AS "))
		AsFragment []byte
		// The quote rune to use when quoting identifiers(DEFAULT='"')
		QuoteRune rune
		// The NULL literal to use when interpolating nulls values (DEFAULT=[]byte("NULL"))
		Null []byte
		// The TRUE literal to use when interpolating bool true values (DEFAULT=[]byte("TRUE"))
		True []byte
		// The FALSE literal to use when interpolating bool false values (DEFAULT=[]byte("FALSE"))
		False []byte
		// The ASC fragment when specifying column order (DEFAULT=[]byte(" ASC"))
		AscFragment []byte
		// The DESC fragment when specifying column order (DEFAULT=[]byte(" DESC"))
		DescFragment []byte
		// The NULLS FIRST fragment when specifying column order (DEFAULT=[]byte(" NULLS FIRST"))
		NullsFirstFragment []byte
		// The NULLS LAST fragment when specifying column order (DEFAULT=[]byte(" NULLS LAST"))
		NullsLastFragment []byte
		// The AND keyword used when joining ExpressionLists (DEFAULT=[]byte(" AND "))
		AndFragment []byte
		// The OR keyword used when joining ExpressionLists (DEFAULT=[]byte(" OR "))
		OrFragment []byte
		// The UNION keyword used when creating compound statements (DEFAULT=[]byte(" UNION "))
		UnionFragment []byte
		// The UNION ALL keyword used when creating compound statements (DEFAULT=[]byte(" UNION ALL "))
		UnionAllFragment []byte
		// The INTERSECT keyword used when creating compound statements (DEFAULT=[]byte(" INTERSECT "))
		IntersectFragment []byte
		// The INTERSECT ALL keyword used when creating compound statements (DEFAULT=[]byte(" INTERSECT ALL "))
		IntersectAllFragment []byte
		// The CAST keyword to use when casting a value (DEFAULT=[]byte("CAST"))
		CastFragment []byte
		// The quote rune to use when quoting string literals (DEFAULT='\'')
		StringQuote rune
		// The operator to use when setting values in an update statement (DEFAULT='=')
		SetOperatorRune rune
		// The placeholder rune to use when generating a non interpolated statement (DEFAULT='?')
		PlaceHolderRune rune
		// Empty string (DEFAULT="")
		EmptyString string
		// Comma rune (DEFAULT=',')
		CommaRune rune
		// Space rune (DEFAULT=' ')
		SpaceRune rune
		// Left paren rune (DEFAULT='(')
		LeftParenRune rune
		// Right paren rune (DEFAULT=')')
		RightParenRune rune
		// Star rune (DEFAULT='*')
		StarRune rune
		// Period rune (DEFAULT='.')
		PeriodRune rune
		// Set to true to include positional argument numbers when creating a prepared statement (Default=false)
		IncludePlaceholderNum bool
		// The time format to use when serializing time.Time (DEFAULT=time.RFC3339Nano)
		TimeFormat string
		// A map used to look up BooleanOperations and their SQL equivalents
		// (Default= map[exp.BooleanOperation][]byte{
		// 		exp.EqOp:             []byte("="),
		// 		exp.NeqOp:            []byte("!="),
		// 		exp.GtOp:             []byte(">"),
		// 		exp.GteOp:            []byte(">="),
		// 		exp.LtOp:             []byte("<"),
		// 		exp.LteOp:            []byte("<="),
		// 		exp.InOp:             []byte("IN"),
		// 		exp.NotInOp:          []byte("NOT IN"),
		// 		exp.IsOp:             []byte("IS"),
		// 		exp.IsNotOp:          []byte("IS NOT"),
		// 		exp.LikeOp:           []byte("LIKE"),
		// 		exp.NotLikeOp:        []byte("NOT LIKE"),
		// 		exp.ILikeOp:          []byte("ILIKE"),
		// 		exp.NotILikeOp:       []byte("NOT ILIKE"),
		// 		exp.RegexpLikeOp:     []byte("~"),
		// 		exp.RegexpNotLikeOp:  []byte("!~"),
		// 		exp.RegexpILikeOp:    []byte("~*"),
		// 		exp.RegexpNotILikeOp: []byte("!~*"),
		// })
		BooleanOperatorLookup map[exp.BooleanOperation][]byte
		// A map used to look up RangeOperations and their SQL equivalents
		// (Default=map[exp.RangeOperation][]byte{
		// 		exp.BetweenOp:    []byte("BETWEEN"),
		// 		exp.NotBetweenOp: []byte("NOT BETWEEN"),
		// 	})
		RangeOperatorLookup map[exp.RangeOperation][]byte
		// A map used to look up JoinTypes and their SQL equivalents
		// (Default= map[exp.JoinType][]byte{
		// 		exp.InnerJoinType:        []byte(" INNER JOIN "),
		// 		exp.FullOuterJoinType:    []byte(" FULL OUTER JOIN "),
		// 		exp.RightOuterJoinType:   []byte(" RIGHT OUTER JOIN "),
		// 		exp.LeftOuterJoinType:    []byte(" LEFT OUTER JOIN "),
		// 		exp.FullJoinType:         []byte(" FULL JOIN "),
		// 		exp.RightJoinType:        []byte(" RIGHT JOIN "),
		// 		exp.LeftJoinType:         []byte(" LEFT JOIN "),
		// 		exp.NaturalJoinType:      []byte(" NATURAL JOIN "),
		// 		exp.NaturalLeftJoinType:  []byte(" NATURAL LEFT JOIN "),
		// 		exp.NaturalRightJoinType: []byte(" NATURAL RIGHT JOIN "),
		// 		exp.NaturalFullJoinType:  []byte(" NATURAL FULL JOIN "),
		// 		exp.CrossJoinType:        []byte(" CROSS JOIN "),
		// 	})
		JoinTypeLookup map[exp.JoinType][]byte
		// Whether or not to use literal TRUE or FALSE for IS statements (e.g. IS TRUE or IS 0)
		UseLiteralIsBools bool
		// EscapedRunes is a map of a rune and the corresponding escape sequence in bytes. Used when escaping text
		// types.
		// (Default= map[rune][]byte{
		// 		'\'': []byte("''"),
		// 	})
		EscapedRunes map[rune][]byte

		// The SQL fragment to use for CONFLICT (Default=[]byte(" ON CONFLICT"))
		ConflictFragment []byte
		// The SQL fragment to use for CONFLICT DO NOTHING (Default=[]byte(" DO NOTHING"))
		ConflictDoNothingFragment []byte
		// The SQL fragment to use for CONFLICT DO UPDATE (Default=[]byte(" DO UPDATE SET"))
		ConflictDoUpdateFragment []byte

		// The order of SQL fragments when creating a SELECT statement
		// (Default=[]SQLFragmentType{
		// 		CommonTableSQLFragment,
		// 		SelectSQLFragment,
		// 		FromSQLFragment,
		// 		JoinSQLFragment,
		// 		WhereSQLFragment,
		// 		GroupBySQLFragment,
		// 		HavingSQLFragment,
		// 		CompoundsSQLFragment,
		// 		OrderSQLFragment,
		// 		LimitSQLFragment,
		// 		OffsetSQLFragment,
		// 		ForSQLFragment,
		// 	})
		SelectSQLOrder []SQLFragmentType

		// The order of SQL fragments when creating an UPDATE statement
		// (Default=[]SQLFragmentType{
		// 		CommonTableSQLFragment,
		// 		UpdateBeginSQLFragment,
		// 		SourcesSQLFragment,
		// 		UpdateSQLFragment,
		// 		WhereSQLFragment,
		// 		OrderSQLFragment,
		// 		LimitSQLFragment,
		// 		ReturningSQLFragment,
		// 	})
		UpdateSQLOrder []SQLFragmentType

		// The order of SQL fragments when creating an INSERT statement
		// (Default=[]SQLFragmentType{
		// 		CommonTableSQLFragment,
		// 		InsertBeingSQLFragment,
		// 		SourcesSQLFragment,
		// 		InsertSQLFragment,
		// 		ReturningSQLFragment,
		// 	})
		InsertSQLOrder []SQLFragmentType

		// The order of SQL fragments when creating a DELETE statement
		// (Default=[]SQLFragmentType{
		// 		CommonTableSQLFragment,
		// 		DeleteBeginSQLFragment,
		// 		FromSQLFragment,
		// 		WhereSQLFragment,
		// 		OrderSQLFragment,
		// 		LimitSQLFragment,
		// 		ReturningSQLFragment,
		// 	})
		DeleteSQLOrder []SQLFragmentType

		// The order of SQL fragments when creating a TRUNCATE statement
		// (Default=[]SQLFragmentType{
		// 		TruncateSQLFragment,
		// 	})
		TruncateSQLOrder []SQLFragmentType
	}
)

const (
	CommonTableSQLFragment = iota
	SelectSQLFragment
	FromSQLFragment
	JoinSQLFragment
	WhereSQLFragment
	GroupBySQLFragment
	HavingSQLFragment
	CompoundsSQLFragment
	OrderSQLFragment
	LimitSQLFragment
	OffsetSQLFragment
	ForSQLFragment
	UpdateBeginSQLFragment
	SourcesSQLFragment
	IntoSQLFragment
	UpdateSQLFragment
	UpdateFromSQLFragment
	ReturningSQLFragment
	InsertBeingSQLFragment
	InsertSQLFragment
	DeleteBeginSQLFragment
	TruncateSQLFragment
)

// nolint:gocyclo
func (sf SQLFragmentType) String() string {
	switch sf {
	case CommonTableSQLFragment:
		return "CommonTableSQLFragment"
	case SelectSQLFragment:
		return "SelectSQLFragment"
	case FromSQLFragment:
		return "FromSQLFragment"
	case JoinSQLFragment:
		return "JoinSQLFragment"
	case WhereSQLFragment:
		return "WhereSQLFragment"
	case GroupBySQLFragment:
		return "GroupBySQLFragment"
	case HavingSQLFragment:
		return "HavingSQLFragment"
	case CompoundsSQLFragment:
		return "CompoundsSQLFragment"
	case OrderSQLFragment:
		return "OrderSQLFragment"
	case LimitSQLFragment:
		return "LimitSQLFragment"
	case OffsetSQLFragment:
		return "OffsetSQLFragment"
	case ForSQLFragment:
		return "ForSQLFragment"
	case UpdateBeginSQLFragment:
		return "UpdateBeginSQLFragment"
	case SourcesSQLFragment:
		return "SourcesSQLFragment"
	case IntoSQLFragment:
		return "IntoSQLFragment"
	case UpdateSQLFragment:
		return "UpdateSQLFragment"
	case UpdateFromSQLFragment:
		return "UpdateFromSQLFragment"
	case ReturningSQLFragment:
		return "ReturningSQLFragment"
	case InsertBeingSQLFragment:
		return "InsertBeingSQLFragment"
	case DeleteBeginSQLFragment:
		return "DeleteBeginSQLFragment"
	case TruncateSQLFragment:
		return "TruncateSQLFragment"
	}
	return fmt.Sprintf("%d", sf)
}

func DefaultDialectOptions() *SQLDialectOptions {
	return &SQLDialectOptions{
		SupportsOrderByOnDelete:     false,
		SupportsOrderByOnUpdate:     false,
		SupportsLimitOnDelete:       false,
		SupportsLimitOnUpdate:       false,
		SupportsReturn:              true,
		SupportsConflictUpdateWhere: true,
		SupportsInsertIgnoreSyntax:  false,
		SupportsConflictTarget:      true,
		SupportsWithCTE:             true,
		SupportsWithCTERecursive:    true,
		WrapCompoundsInParens:       true,

		SupportsMultipleUpdateTables:         true,
		UseFromClauseForMultipleUpdateTables: true,

		UpdateClause:              []byte("UPDATE"),
		InsertClause:              []byte("INSERT INTO"),
		InsertIgnoreClause:        []byte("INSERT IGNORE INTO"),
		SelectClause:              []byte("SELECT"),
		DeleteClause:              []byte("DELETE"),
		TruncateClause:            []byte("TRUNCATE"),
		WithFragment:              []byte("WITH "),
		RecursiveFragment:         []byte("RECURSIVE "),
		CascadeFragment:           []byte(" CASCADE"),
		RestrictFragment:          []byte(" RESTRICT"),
		DefaultValuesFragment:     []byte(" DEFAULT VALUES"),
		ValuesFragment:            []byte(" VALUES "),
		IdentityFragment:          []byte(" IDENTITY"),
		SetFragment:               []byte(" SET "),
		DistinctFragment:          []byte(" DISTINCT "),
		ReturningFragment:         []byte(" RETURNING "),
		FromFragment:              []byte(" FROM"),
		UsingFragment:             []byte(" USING "),
		OnFragment:                []byte(" ON "),
		WhereFragment:             []byte(" WHERE "),
		GroupByFragment:           []byte(" GROUP BY "),
		HavingFragment:            []byte(" HAVING "),
		OrderByFragment:           []byte(" ORDER BY "),
		LimitFragment:             []byte(" LIMIT "),
		OffsetFragment:            []byte(" OFFSET "),
		ForUpdateFragment:         []byte(" FOR UPDATE "),
		ForNoKeyUpdateFragment:    []byte(" FOR NO KEY UPDATE "),
		ForShareFragment:          []byte(" FOR SHARE "),
		ForKeyShareFragment:       []byte(" FOR KEY SHARE "),
		NowaitFragment:            []byte("NOWAIT"),
		SkipLockedFragment:        []byte("SKIP LOCKED"),
		AsFragment:                []byte(" AS "),
		AscFragment:               []byte(" ASC"),
		DescFragment:              []byte(" DESC"),
		NullsFirstFragment:        []byte(" NULLS FIRST"),
		NullsLastFragment:         []byte(" NULLS LAST"),
		AndFragment:               []byte(" AND "),
		OrFragment:                []byte(" OR "),
		UnionFragment:             []byte(" UNION "),
		UnionAllFragment:          []byte(" UNION ALL "),
		IntersectFragment:         []byte(" INTERSECT "),
		IntersectAllFragment:      []byte(" INTERSECT ALL "),
		ConflictFragment:          []byte(" ON CONFLICT"),
		ConflictDoUpdateFragment:  []byte(" DO UPDATE SET "),
		ConflictDoNothingFragment: []byte(" DO NOTHING"),
		CastFragment:              []byte("CAST"),
		Null:                      []byte("NULL"),
		True:                      []byte("TRUE"),
		False:                     []byte("FALSE"),

		PlaceHolderRune: '?',
		QuoteRune:       '"',
		StringQuote:     '\'',
		SetOperatorRune: '=',
		CommaRune:       ',',
		SpaceRune:       ' ',
		LeftParenRune:   '(',
		RightParenRune:  ')',
		StarRune:        '*',
		PeriodRune:      '.',
		EmptyString:     "",

		BooleanOperatorLookup: map[exp.BooleanOperation][]byte{
			exp.EqOp:             []byte("="),
			exp.NeqOp:            []byte("!="),
			exp.GtOp:             []byte(">"),
			exp.GteOp:            []byte(">="),
			exp.LtOp:             []byte("<"),
			exp.LteOp:            []byte("<="),
			exp.InOp:             []byte("IN"),
			exp.NotInOp:          []byte("NOT IN"),
			exp.IsOp:             []byte("IS"),
			exp.IsNotOp:          []byte("IS NOT"),
			exp.LikeOp:           []byte("LIKE"),
			exp.NotLikeOp:        []byte("NOT LIKE"),
			exp.ILikeOp:          []byte("ILIKE"),
			exp.NotILikeOp:       []byte("NOT ILIKE"),
			exp.RegexpLikeOp:     []byte("~"),
			exp.RegexpNotLikeOp:  []byte("!~"),
			exp.RegexpILikeOp:    []byte("~*"),
			exp.RegexpNotILikeOp: []byte("!~*"),
		},
		RangeOperatorLookup: map[exp.RangeOperation][]byte{
			exp.BetweenOp:    []byte("BETWEEN"),
			exp.NotBetweenOp: []byte("NOT BETWEEN"),
		},
		JoinTypeLookup: map[exp.JoinType][]byte{
			exp.InnerJoinType:        []byte(" INNER JOIN "),
			exp.FullOuterJoinType:    []byte(" FULL OUTER JOIN "),
			exp.RightOuterJoinType:   []byte(" RIGHT OUTER JOIN "),
			exp.LeftOuterJoinType:    []byte(" LEFT OUTER JOIN "),
			exp.FullJoinType:         []byte(" FULL JOIN "),
			exp.RightJoinType:        []byte(" RIGHT JOIN "),
			exp.LeftJoinType:         []byte(" LEFT JOIN "),
			exp.NaturalJoinType:      []byte(" NATURAL JOIN "),
			exp.NaturalLeftJoinType:  []byte(" NATURAL LEFT JOIN "),
			exp.NaturalRightJoinType: []byte(" NATURAL RIGHT JOIN "),
			exp.NaturalFullJoinType:  []byte(" NATURAL FULL JOIN "),
			exp.CrossJoinType:        []byte(" CROSS JOIN "),
		},

		TimeFormat:        time.RFC3339Nano,
		UseLiteralIsBools: true,

		EscapedRunes: map[rune][]byte{
			'\'': []byte("''"),
		},

		SelectSQLOrder: []SQLFragmentType{
			CommonTableSQLFragment,
			SelectSQLFragment,
			FromSQLFragment,
			JoinSQLFragment,
			WhereSQLFragment,
			GroupBySQLFragment,
			HavingSQLFragment,
			CompoundsSQLFragment,
			OrderSQLFragment,
			LimitSQLFragment,
			OffsetSQLFragment,
			ForSQLFragment,
		},
		UpdateSQLOrder: []SQLFragmentType{
			CommonTableSQLFragment,
			UpdateBeginSQLFragment,
			SourcesSQLFragment,
			UpdateSQLFragment,
			UpdateFromSQLFragment,
			WhereSQLFragment,
			OrderSQLFragment,
			LimitSQLFragment,
			ReturningSQLFragment,
		},
		InsertSQLOrder: []SQLFragmentType{
			CommonTableSQLFragment,
			InsertBeingSQLFragment,
			IntoSQLFragment,
			InsertSQLFragment,
			ReturningSQLFragment,
		},
		DeleteSQLOrder: []SQLFragmentType{
			CommonTableSQLFragment,
			DeleteBeginSQLFragment,
			FromSQLFragment,
			WhereSQLFragment,
			OrderSQLFragment,
			LimitSQLFragment,
			ReturningSQLFragment,
		},
		TruncateSQLOrder: []SQLFragmentType{
			TruncateSQLFragment,
		},
	}
}