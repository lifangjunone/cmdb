package impl

const (
	sqlInsertResource = `INSERT INTO resource (
       id, resource_type, vendor, region, zone, create_at, expire_at, category, type,
       name,description,status,update_at,sync_at,sync_account,public_ip,
       private_ip,pay_type,describe_hash,resource_hash,secret_id,domain,
       namespace,env,usage_mode
    ) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);`

	sqlUpdateResource = `UPDATE resource SET expire_at=?,category=?,type=?,name=?,description=?,
   	   status=?,update_at=?,sync_at=?,sync_account=?,
       public_ip=?,private_ip=?,pay_type=?,describe_hash=?,resource_hash=?,
       secret_id=?,namespace=?,env=?,usage_mode=?
    WHERE id = ?`

	sqlDeleteResource = `DELETE FROM resource WHERE id = ?;`

	sqlQueryResource = `SELECT r.* FROM resource r %s JOIN resource_tag t ON r.id = t.resource_id`

	sqlCountResource = `SELECT COUNT(DISTINCT r.id) FROM resource r %s JOIN resource_tag t ON r.id = t.resource_id`
)
