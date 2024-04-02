// Package story defines the entity object for user story.
package story

import (
	"fmt"
	"time"

	"github.com/w40141/UsdmApi/domain/shelf"
	"github.com/w40141/UsdmApi/domain/vo"
)

// ParentOfEpisode implements request.ParentOfEpisode.
func (*E) ParentOfEpisode() error {
	panic("unimplemented")
}

// Description implements request.Storyer.
func (s E) Description() string {
	return s.description.String()
}

// ID implements request.Storyer.
func (s E) ID() vo.ID {
	return s.id
}

// Reason implements request.Storyer.
func (s E) Reason() string {
	return s.reason.String()
}

// Title implements request.Storyer.
func (s E) Title() string {
	return s.title.String()
}

// ParentID getter
func (s E) ParentID() string {
	return s.parentID.String()
}

// CreateAt getter
func (s E) CreateAt() (time.Time, bool) {
	return *s.createAt, s.createAt != nil
}

// UpdateAt getter
func (s E) UpdateAt() (time.Time, bool) {
	return *s.updateAt, s.updateAt != nil
}

// ParentOfScene implements request.Storyer.
func (*E) ParentOfScene() error {
	panic("unimplemented")
}

// Update updates a Story.
func (s *E) Update(
	title string,
	description string,
	reason string,
	parent shelf.ParentOfStory,
) (E, error) {
	if s == nil {
		return E{}, fmt.Errorf("story is nil")
	}
	return New(
		s.id.String(),
		title,
		description,
		reason,
		s.bookID.String(),
		parent.ID().String(),
	)
}

// Delete deletes a Story.
func (s *E) Delete(
	participant shelf.Participanter,
) (D, error) {
	if s == nil {
		return D{}, fmt.Errorf("story is nil")
	}
	if !participant.CanDelete() {
		return D{}, fmt.Errorf("participant can not delete")
	}
	return D{
		id: s.id,
	}, nil
}
