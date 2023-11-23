// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: groupEventMemberQueries.sql

package crud

import (
	"context"
)

const createGroupEventMember = `-- name: CreateGroupEventMember :one
INSERT INTO group_event_member (
  user_id, event_id, status_
) VALUES (
  ?, ?, ?
)
RETURNING id, user_id, event_id, status_
`

type CreateGroupEventMemberParams struct {
	UserID  int64
	EventID int64
	Status  int64
}

func (q *Queries) CreateGroupEventMember(ctx context.Context, arg CreateGroupEventMemberParams) (GroupEventMember, error) {
	row := q.db.QueryRowContext(ctx, createGroupEventMember, arg.UserID, arg.EventID, arg.Status)
	var i GroupEventMember
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.EventID,
		&i.Status,
	)
	return i, err
}

const deleteGroupEventMember = `-- name: DeleteGroupEventMember :exec
DELETE FROM group_event_member
WHERE event_id = ? AND user_id = ?
`

type DeleteGroupEventMemberParams struct {
	EventID int64
	UserID  int64
}

func (q *Queries) DeleteGroupEventMember(ctx context.Context, arg DeleteGroupEventMemberParams) error {
	_, err := q.db.ExecContext(ctx, deleteGroupEventMember, arg.EventID, arg.UserID)
	return err
}

const execUpdateGroupEventMember = `-- name: ExecUpdateGroupEventMember :exec
UPDATE group_event_member
SET status_ = CASE
    WHEN status_ = 0 THEN 1
    ELSE status_
END
WHERE event_id = ? AND user_id = ?
`

type ExecUpdateGroupEventMemberParams struct {
	EventID int64
	UserID  int64
}

func (q *Queries) ExecUpdateGroupEventMember(ctx context.Context, arg ExecUpdateGroupEventMemberParams) error {
	_, err := q.db.ExecContext(ctx, execUpdateGroupEventMember, arg.EventID, arg.UserID)
	return err
}

const getGroupEventMember = `-- name: GetGroupEventMember :one
SELECT COUNT(*) FROM group_event_member
WHERE event_id = ? AND user_id = ? LIMIT 1
`

type GetGroupEventMemberParams struct {
	EventID int64
	UserID  int64
}

func (q *Queries) GetGroupEventMember(ctx context.Context, arg GetGroupEventMemberParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, getGroupEventMember, arg.EventID, arg.UserID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getGroupEventMembers = `-- name: GetGroupEventMembers :many
SELECT id, user_id, event_id, status_ FROM group_event_member
WHERE event_id = ?
`

func (q *Queries) GetGroupEventMembers(ctx context.Context, eventID int64) ([]GroupEventMember, error) {
	rows, err := q.db.QueryContext(ctx, getGroupEventMembers, eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GroupEventMember
	for rows.Next() {
		var i GroupEventMember
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.EventID,
			&i.Status,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getGroupEventMembersGoing = `-- name: GetGroupEventMembersGoing :many
SELECT id, user_id, event_id, status_ FROM group_event_member
WHERE event_id = ? AND status_ = 1
`

func (q *Queries) GetGroupEventMembersGoing(ctx context.Context, eventID int64) ([]GroupEventMember, error) {
	rows, err := q.db.QueryContext(ctx, getGroupEventMembersGoing, eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GroupEventMember
	for rows.Next() {
		var i GroupEventMember
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.EventID,
			&i.Status,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getGroupEventsByUserAccepted = `-- name: GetGroupEventsByUserAccepted :many
SELECT id, user_id, event_id, status_ FROM group_event_member
WHERE user_id = ? AND status_ = 2
`

func (q *Queries) GetGroupEventsByUserAccepted(ctx context.Context, userID int64) ([]GroupEventMember, error) {
	rows, err := q.db.QueryContext(ctx, getGroupEventsByUserAccepted, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GroupEventMember
	for rows.Next() {
		var i GroupEventMember
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.EventID,
			&i.Status,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getGroupEventsByUserNoReply = `-- name: GetGroupEventsByUserNoReply :many
SELECT id, user_id, event_id, status_ FROM group_event_member
WHERE user_id = ? AND status_ = 0
`

func (q *Queries) GetGroupEventsByUserNoReply(ctx context.Context, userID int64) ([]GroupEventMember, error) {
	rows, err := q.db.QueryContext(ctx, getGroupEventsByUserNoReply, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GroupEventMember
	for rows.Next() {
		var i GroupEventMember
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.EventID,
			&i.Status,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateGroupEventMember = `-- name: UpdateGroupEventMember :one
UPDATE group_event_member
set status_ = ?
WHERE event_id = ? AND user_id = ?
RETURNING id, user_id, event_id, status_
`

type UpdateGroupEventMemberParams struct {
	Status  int64
	EventID int64
	UserID  int64
}

func (q *Queries) UpdateGroupEventMember(ctx context.Context, arg UpdateGroupEventMemberParams) (GroupEventMember, error) {
	row := q.db.QueryRowContext(ctx, updateGroupEventMember, arg.Status, arg.EventID, arg.UserID)
	var i GroupEventMember
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.EventID,
		&i.Status,
	)
	return i, err
}