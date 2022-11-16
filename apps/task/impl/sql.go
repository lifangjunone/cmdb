package impl

const (
	insertTaskSQL = `
	INSERT INTO task (
		id,region,resource_type,secret_id,secret_desc,timeout,status,
		message,start_at,end_at,total_succeed,total_failed
	) VALUES (?,?,?,?,?,?,?,?,?,?,?,?);
	`

	updateTaskSQL = `
	UPDATE task SET status=?,message=?,end_at=?,
	total_succeed=?,total_failed=? WHERE id = ?
	`

	queryTaskSQL = `SELECT * FROM task`
)
