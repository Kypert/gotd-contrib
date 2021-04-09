// Code generated by 'yaegi extract github.com/gotd/td/telegram/query/messages'. DO NOT EDIT.

package yaegi

import (
	"context"
	"github.com/gotd/td/telegram/query/messages"
	"github.com/gotd/td/tg"
	"reflect"
)

func init() {
	Symbols["github.com/gotd/td/telegram/query/messages"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"NewIterator":     reflect.ValueOf(messages.NewIterator),
		"NewQueryBuilder": reflect.ValueOf(messages.NewQueryBuilder),

		// type definitions
		"Elem":                           reflect.ValueOf((*messages.Elem)(nil)),
		"File":                           reflect.ValueOf((*messages.File)(nil)),
		"GetHistoryQueryBuilder":         reflect.ValueOf((*messages.GetHistoryQueryBuilder)(nil)),
		"GetRecentLocationsQueryBuilder": reflect.ValueOf((*messages.GetRecentLocationsQueryBuilder)(nil)),
		"GetRepliesQueryBuilder":         reflect.ValueOf((*messages.GetRepliesQueryBuilder)(nil)),
		"GetUnreadMentionsQueryBuilder":  reflect.ValueOf((*messages.GetUnreadMentionsQueryBuilder)(nil)),
		"Iterator":                       reflect.ValueOf((*messages.Iterator)(nil)),
		"Query":                          reflect.ValueOf((*messages.Query)(nil)),
		"QueryBuilder":                   reflect.ValueOf((*messages.QueryBuilder)(nil)),
		"QueryFunc":                      reflect.ValueOf((*messages.QueryFunc)(nil)),
		"Request":                        reflect.ValueOf((*messages.Request)(nil)),
		"SearchGlobalQueryBuilder":       reflect.ValueOf((*messages.SearchGlobalQueryBuilder)(nil)),
		"SearchQueryBuilder":             reflect.ValueOf((*messages.SearchQueryBuilder)(nil)),
		"StatsGetMessagePublicForwardsQueryBuilder": reflect.ValueOf((*messages.StatsGetMessagePublicForwardsQueryBuilder)(nil)),

		// interface wrapper definitions
		"_Query": reflect.ValueOf((*_github_com_gotd_td_telegram_query_messages_Query)(nil)),
	}
}

// _github_com_gotd_td_telegram_query_messages_Query is an interface wrapper for Query type
type _github_com_gotd_td_telegram_query_messages_Query struct {
	WQuery func(ctx context.Context, req messages.Request) (tg.MessagesMessagesClass, error)
}

func (W _github_com_gotd_td_telegram_query_messages_Query) Query(ctx context.Context, req messages.Request) (tg.MessagesMessagesClass, error) {
	return W.WQuery(ctx, req)
}
