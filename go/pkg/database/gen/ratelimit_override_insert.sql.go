// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: ratelimit_override_insert.sql

package gen

import (
	"context"
)

const insertOverride = `-- name: InsertOverride :exec
INSERT INTO
    ` + "`" + `ratelimit_overrides` + "`" + ` (
        id,
        workspace_id,
        namespace_id,
        identifier,
        ` + "`" + `limit` + "`" + `,
        duration,
        async,
        created_at
    )
VALUES
    (
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        false,
        now()
    )
`

type InsertOverrideParams struct {
	ID          string `db:"id"`
	WorkspaceID string `db:"workspace_id"`
	NamespaceID string `db:"namespace_id"`
	Identifier  string `db:"identifier"`
	Limit       int32  `db:"limit"`
	Duration    int32  `db:"duration"`
}

// InsertOverride
//
//	INSERT INTO
//	    `ratelimit_overrides` (
//	        id,
//	        workspace_id,
//	        namespace_id,
//	        identifier,
//	        `limit`,
//	        duration,
//	        async,
//	        created_at
//	    )
//	VALUES
//	    (
//	        ?,
//	        ?,
//	        ?,
//	        ?,
//	        ?,
//	        ?,
//	        false,
//	        now()
//	    )
func (q *Queries) InsertOverride(ctx context.Context, arg InsertOverrideParams) error {
	_, err := q.db.ExecContext(ctx, insertOverride,
		arg.ID,
		arg.WorkspaceID,
		arg.NamespaceID,
		arg.Identifier,
		arg.Limit,
		arg.Duration,
	)
	return err
}
