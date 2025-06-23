package scriptservicetype

import (
	"fmt"
	_ "modernc.org/sqlite"
	"scrunt-back/models/scrunt"
)

func UpdateScriptServiceType(scriptServiceTypeId, scriptId, serviceTypeId int, name string) (int64, error) {
	statement, err := scrunt.Db.Prepare(`UPDATE script_service_types SET script_id=?, service_type_id=?, name=? WHERE id = ?`)
	if err != nil {
		return 0, err
	}

	result, err := statement.Exec(scriptId, serviceTypeId, name, scriptServiceTypeId)
	if err != nil {
		fmt.Println("updateServiceKey.UpdateServiceKey - Error:", err)
		return 0, err
	} else {
		return result.RowsAffected()
	}
}
