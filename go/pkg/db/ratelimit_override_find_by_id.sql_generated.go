// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: ratelimit_override_find_by_id.sql

package db

import (
	"context"
)

const findRatelimitOverrideById = `-- name: FindRatelimitOverrideById :one
SELECT id, workspace_id, namespace_id, identifier, ` + "`" + `limit` + "`" + `, duration, async, sharding, created_at_m, updated_at_m, deleted_at_m FROM ratelimit_overrides
WHERE
    workspace_id = ?
    AND id = ?
`

type FindRatelimitOverrideByIdParams struct {
	WorkspaceID string `db:"workspace_id"`
	OverrideID  string `db:"override_id"`
}

// FindRatelimitOverrideById
//
//	SELECT id, workspace_id, namespace_id, identifier, `limit`, duration, async, sharding, created_at_m, updated_at_m, deleted_at_m FROM ratelimit_overrides
//	WHERE
//	    workspace_id = ?
//	    AND id = ?
func (q *Queries) FindRatelimitOverrideById(ctx context.Context, db DBTX, arg FindRatelimitOverrideByIdParams) (RatelimitOverride, error) {
	row := db.QueryRowContext(ctx, findRatelimitOverrideById, arg.WorkspaceID, arg.OverrideID)
	var i RatelimitOverride
	err := row.Scan(
		&i.ID,
		&i.WorkspaceID,
		&i.NamespaceID,
		&i.Identifier,
		&i.Limit,
		&i.Duration,
		&i.Async,
		&i.Sharding,
		&i.CreatedAtM,
		&i.UpdatedAtM,
		&i.DeletedAtM,
	)
	return i, err
}
