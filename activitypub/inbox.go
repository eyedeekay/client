package activitypub

import "fmt"

type (
	// InboxStream contains all activities received by the actor.
	// The server SHOULD filter content according to the requester's permission.
	// In general, the owner of an inbox is likely to be able to access all of their inbox contents.
	// Depending on access control, some other content may be public, whereas other content may
	// require authentication for non-owner users, if they can access the inbox at all.
	InboxStream Inbox

	// Inbox is a type alias for an Ordered Collection
	Inbox OrderedCollection
)

// InboxNew initializes a new Inbox
func InboxNew() *Inbox {
	id := ObjectID("inbox")

	i := Inbox{ID: id, Type: OrderedCollectionType}
	i.Name = NaturalLanguageValueNew()
	i.Content = NaturalLanguageValueNew()
	i.Summary = NaturalLanguageValueNew()

	i.TotalItems = 0

	return &i
}

// Append adds an element to an InboxStream
func (i *InboxStream) Append(o Item) error {
	if i == nil {
		return fmt.Errorf("nil ")
	}
	i.OrderedItems = append(i.OrderedItems, o)
	i.TotalItems++
	return nil
}

// Append adds an element to an Inbox
func (i *Inbox) Append(ob Item) error {
	i.OrderedItems = append(i.OrderedItems, ob)
	i.TotalItems++
	return nil
}

// GetID returns the ObjectID corresponding to InboxStream
func (i InboxStream) GetID() *ObjectID {
	return i.Collection().GetID()
}

// GetType returns the InboxStream's type
func (i InboxStream) GetType() ActivityVocabularyType {
	return i.Type
}

// IsLink returns false for an InboxStream object
func (i InboxStream) IsLink() bool {
	return false
}

// IsObject returns true for a InboxStream object
func (i InboxStream) IsObject() bool {
	return true
}

// GetID returns the ObjectID corresponding to Inbox
func (i Inbox) GetID() *ObjectID {
	return i.Collection().GetID()
}

// GetType returns the Inbox's type
func (i Inbox) GetType() ActivityVocabularyType {
	return i.Type
}

// IsLink returns false for an Inbox object
func (i Inbox) IsLink() bool {
	return false
}

// IsObject returns true for a Inbox object
func (i Inbox) IsObject() bool {
	return true
}

// UnmarshalJSON
func (i *InboxStream) UnmarshalJSON(data []byte) error {
	c := OrderedCollection(*i)
	err := c.UnmarshalJSON(data)

	*i = InboxStream(c)

	return err
}

// UnmarshalJSON
func (i *Inbox) UnmarshalJSON(data []byte) error {
	c := OrderedCollection(*i)
	err := c.UnmarshalJSON(data)

	*i = Inbox(c)

	return err
}

// Collection returns the underlying Collection type
func (i Inbox) Collection() CollectionInterface {
	c := OrderedCollection(i)
	return &c
}

// Collection returns the underlying Collection type
func (i InboxStream) Collection() CollectionInterface {
	c := OrderedCollection(i)
	return &c
}
