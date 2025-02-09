// Code generated by ent, DO NOT EDIT.

package feed

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the feed type in the database.
	Label = "feed"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldFeedTitle holds the string denoting the feed_title field in the database.
	FieldFeedTitle = "feed_title"
	// FieldFeedDescription holds the string denoting the feed_description field in the database.
	FieldFeedDescription = "feed_description"
	// FieldFeedLink holds the string denoting the feed_link field in the database.
	FieldFeedLink = "feed_link"
	// FieldFeedFeedLink holds the string denoting the feed_feed_link field in the database.
	FieldFeedFeedLink = "feed_feed_link"
	// FieldFeedUpdated holds the string denoting the feed_updated field in the database.
	FieldFeedUpdated = "feed_updated"
	// FieldFeedPublished holds the string denoting the feed_published field in the database.
	FieldFeedPublished = "feed_published"
	// FieldFeedAuthorName holds the string denoting the feed_author_name field in the database.
	FieldFeedAuthorName = "feed_author_name"
	// FieldFeedAuthorEmail holds the string denoting the feed_author_email field in the database.
	FieldFeedAuthorEmail = "feed_author_email"
	// FieldFeedLanguage holds the string denoting the feed_language field in the database.
	FieldFeedLanguage = "feed_language"
	// FieldFeedImageTitle holds the string denoting the feed_image_title field in the database.
	FieldFeedImageTitle = "feed_image_title"
	// FieldFeedImageURL holds the string denoting the feed_image_url field in the database.
	FieldFeedImageURL = "feed_image_url"
	// FieldFeedCopyright holds the string denoting the feed_copyright field in the database.
	FieldFeedCopyright = "feed_copyright"
	// FieldFeedGenerator holds the string denoting the feed_generator field in the database.
	FieldFeedGenerator = "feed_generator"
	// FieldFeedCategories holds the string denoting the feed_categories field in the database.
	FieldFeedCategories = "feed_categories"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// EdgeItems holds the string denoting the items edge name in mutations.
	EdgeItems = "items"
	// EdgeSubscribedUsers holds the string denoting the subscribed_users edge name in mutations.
	EdgeSubscribedUsers = "subscribed_users"
	// EdgeSubscriptions holds the string denoting the subscriptions edge name in mutations.
	EdgeSubscriptions = "subscriptions"
	// Table holds the table name of the feed in the database.
	Table = "feeds"
	// ItemsTable is the table that holds the items relation/edge.
	ItemsTable = "items"
	// ItemsInverseTable is the table name for the Item entity.
	// It exists in this package in order to avoid circular dependency with the "item" package.
	ItemsInverseTable = "items"
	// ItemsColumn is the table column denoting the items relation/edge.
	ItemsColumn = "feed_items"
	// SubscribedUsersTable is the table that holds the subscribed_users relation/edge. The primary key declared below.
	SubscribedUsersTable = "subscriptions"
	// SubscribedUsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	SubscribedUsersInverseTable = "users"
	// SubscriptionsTable is the table that holds the subscriptions relation/edge.
	SubscriptionsTable = "subscriptions"
	// SubscriptionsInverseTable is the table name for the Subscription entity.
	// It exists in this package in order to avoid circular dependency with the "subscription" package.
	SubscriptionsInverseTable = "subscriptions"
	// SubscriptionsColumn is the table column denoting the subscriptions relation/edge.
	SubscriptionsColumn = "feed_id"
)

// Columns holds all SQL columns for feed fields.
var Columns = []string{
	FieldID,
	FieldURL,
	FieldUsername,
	FieldPassword,
	FieldFeedTitle,
	FieldFeedDescription,
	FieldFeedLink,
	FieldFeedFeedLink,
	FieldFeedUpdated,
	FieldFeedPublished,
	FieldFeedAuthorName,
	FieldFeedAuthorEmail,
	FieldFeedLanguage,
	FieldFeedImageTitle,
	FieldFeedImageURL,
	FieldFeedCopyright,
	FieldFeedGenerator,
	FieldFeedCategories,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
}

var (
	// SubscribedUsersPrimaryKey and SubscribedUsersColumn2 are the table columns denoting the
	// primary key for the subscribed_users relation (M2M).
	SubscribedUsersPrimaryKey = []string{"user_id", "feed_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// URLValidator is a validator for the "url" field. It is called by the builders before save.
	URLValidator func(string) error
	// DefaultUsername holds the default value on creation for the "username" field.
	DefaultUsername string
	// DefaultPassword holds the default value on creation for the "password" field.
	DefaultPassword string
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
