// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: workspace_create.sql

package database

import (
	"context"
	"database/sql"
	"encoding/json"
)

const insertWorkspace = `-- name: InsertWorkspace :exec
INSERT INTO
    ` + "`" + `workspaces` + "`" + ` (
        id,
        name,
        slug,
        tenant_id,
        plan,
        features,
        beta_features
    )
VALUES
    (
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?
    )
`

type InsertWorkspaceParams struct {
	ID           string
	Name         string
	Slug         sql.NullString
	TenantID     string
	Plan         NullWorkspacesPlan
	Features     json.RawMessage
	BetaFeatures json.RawMessage
}

func (q *Queries) InsertWorkspace(ctx context.Context, arg InsertWorkspaceParams) error {
	_, err := q.db.ExecContext(ctx, insertWorkspace,
		arg.ID,
		arg.Name,
		arg.Slug,
		arg.TenantID,
		arg.Plan,
		arg.Features,
		arg.BetaFeatures,
	)
	return err
}
