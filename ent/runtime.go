// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/mrusme/journalist/ent/feed"
	"github.com/mrusme/journalist/ent/item"
	"github.com/mrusme/journalist/ent/read"
	"github.com/mrusme/journalist/ent/schema"
	"github.com/mrusme/journalist/ent/subscription"
	"github.com/mrusme/journalist/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	feedFields := schema.Feed{}.Fields()
	_ = feedFields
	// feedDescSiteURL is the schema descriptor for site_url field.
	feedDescSiteURL := feedFields[3].Descriptor()
	// feed.SiteURLValidator is a validator for the "site_url" field. It is called by the builders before save.
	feed.SiteURLValidator = feedDescSiteURL.Validators[0].(func(string) error)
	// feedDescFeedURL is the schema descriptor for feed_url field.
	feedDescFeedURL := feedFields[4].Descriptor()
	// feed.FeedURLValidator is a validator for the "feed_url" field. It is called by the builders before save.
	feed.FeedURLValidator = feedDescFeedURL.Validators[0].(func(string) error)
	// feedDescCreatedAt is the schema descriptor for created_at field.
	feedDescCreatedAt := feedFields[11].Descriptor()
	// feed.DefaultCreatedAt holds the default value on creation for the created_at field.
	feed.DefaultCreatedAt = feedDescCreatedAt.Default.(func() time.Time)
	// feedDescUpdatedAt is the schema descriptor for updated_at field.
	feedDescUpdatedAt := feedFields[12].Descriptor()
	// feed.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	feed.DefaultUpdatedAt = feedDescUpdatedAt.Default.(func() time.Time)
	// feed.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	feed.UpdateDefaultUpdatedAt = feedDescUpdatedAt.UpdateDefault.(func() time.Time)
	// feedDescID is the schema descriptor for id field.
	feedDescID := feedFields[0].Descriptor()
	// feed.DefaultID holds the default value on creation for the id field.
	feed.DefaultID = feedDescID.Default.(func() uuid.UUID)
	itemFields := schema.Item{}.Fields()
	_ = itemFields
	// itemDescURL is the schema descriptor for url field.
	itemDescURL := itemFields[4].Descriptor()
	// item.URLValidator is a validator for the "url" field. It is called by the builders before save.
	item.URLValidator = itemDescURL.Validators[0].(func(string) error)
	// itemDescCreatedAt is the schema descriptor for created_at field.
	itemDescCreatedAt := itemFields[15].Descriptor()
	// item.DefaultCreatedAt holds the default value on creation for the created_at field.
	item.DefaultCreatedAt = itemDescCreatedAt.Default.(func() time.Time)
	// itemDescUpdatedAt is the schema descriptor for updated_at field.
	itemDescUpdatedAt := itemFields[16].Descriptor()
	// item.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	item.DefaultUpdatedAt = itemDescUpdatedAt.Default.(func() time.Time)
	// item.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	item.UpdateDefaultUpdatedAt = itemDescUpdatedAt.UpdateDefault.(func() time.Time)
	// itemDescID is the schema descriptor for id field.
	itemDescID := itemFields[0].Descriptor()
	// item.DefaultID holds the default value on creation for the id field.
	item.DefaultID = itemDescID.Default.(func() uuid.UUID)
	readFields := schema.Read{}.Fields()
	_ = readFields
	// readDescCreatedAt is the schema descriptor for created_at field.
	readDescCreatedAt := readFields[3].Descriptor()
	// read.DefaultCreatedAt holds the default value on creation for the created_at field.
	read.DefaultCreatedAt = readDescCreatedAt.Default.(func() time.Time)
	// readDescID is the schema descriptor for id field.
	readDescID := readFields[0].Descriptor()
	// read.DefaultID holds the default value on creation for the id field.
	read.DefaultID = readDescID.Default.(func() uuid.UUID)
	subscriptionFields := schema.Subscription{}.Fields()
	_ = subscriptionFields
	// subscriptionDescGroup is the schema descriptor for group field.
	subscriptionDescGroup := subscriptionFields[3].Descriptor()
	// subscription.GroupValidator is a validator for the "group" field. It is called by the builders before save.
	subscription.GroupValidator = subscriptionDescGroup.Validators[0].(func(string) error)
	// subscriptionDescCreatedAt is the schema descriptor for created_at field.
	subscriptionDescCreatedAt := subscriptionFields[4].Descriptor()
	// subscription.DefaultCreatedAt holds the default value on creation for the created_at field.
	subscription.DefaultCreatedAt = subscriptionDescCreatedAt.Default.(func() time.Time)
	// subscriptionDescID is the schema descriptor for id field.
	subscriptionDescID := subscriptionFields[0].Descriptor()
	// subscription.DefaultID holds the default value on creation for the id field.
	subscription.DefaultID = subscriptionDescID.Default.(func() uuid.UUID)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[1].Descriptor()
	// user.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	user.UsernameValidator = userDescUsername.Validators[0].(func(string) error)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[2].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func(string) error)
	// userDescRole is the schema descriptor for role field.
	userDescRole := userFields[3].Descriptor()
	// user.DefaultRole holds the default value on creation for the role field.
	user.DefaultRole = userDescRole.Default.(string)
	// user.RoleValidator is a validator for the "role" field. It is called by the builders before save.
	user.RoleValidator = userDescRole.Validators[0].(func(string) error)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[4].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[5].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
}
