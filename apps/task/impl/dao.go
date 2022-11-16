package impl

import (
	"context"
	"fmt"

	"github.com/lifangjunone/cmdb/apps/task"
)

func (s *impl) insert(ctx context.Context, t *task.Task) error {
	stmt, err := s.db.Prepare(insertTaskSQL)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		t.Id, t.Data.Region, t.Data.ResourceType, t.Data.SecretId, t.SecretDescription, t.Data.Timeout,
		t.Status.Stage, t.Status.Message, t.Status.StartAt, t.Status.EndAt, t.Status.TotalSucceed, t.Status.TotalFailed,
	)
	if err != nil {
		return fmt.Errorf("save task info error, %s", err)
	}
	return nil
}

func (s *impl) update(ctx context.Context, t *task.Task) error {
	stmt, err := s.db.Prepare(updateTaskSQL)
	if err != nil {
		return err
	}
	defer stmt.Close()

	status := t.Status

	_, err = stmt.Exec(
		status.Stage, status.Message, status.EndAt, status.TotalSucceed, status.TotalFailed, t.Id,
	)
	if err != nil {
		return fmt.Errorf("update task info error, %s", err)
	}

	return nil
}
