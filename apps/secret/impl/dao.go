package impl

import (
	"context"
	"fmt"
	"github.com/lifangjunone/cmdb/apps/secret"
)

func (i *impl) deleteSecret(ctx context.Context, ins *secret.Secret) error {
	if ins == nil {
		return fmt.Errorf("secret is nil")
	}

	stmt, err := i.db.PrepareContext(ctx, deleteSecretSQL)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, ins.Id)
	if err != nil {
		return err
	}

	return nil
}
