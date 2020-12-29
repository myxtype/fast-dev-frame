package format

import (
	"testing"
)

func TestNewSubSql(t *testing.T) {
	id := 1
	zone := "+8:00"

	ss := NewSubSql()
	ss.Add("total", "SELECT SUM(amount) FROM `wallet_log` WHERE store_id = ? AND status = 1 AND type IN (1, 2)", id)
	ss.Add("day", "SELECT SUM(amount) FROM `wallet_log` WHERE store_id = ? AND status = 1 AND type IN (1, 2) AND to_days(CONVERT_TZ(created_at, '+0:00', ?)) = to_days(CONVERT_TZ(now(), '+0:00', ?))", id, zone, zone)

	t.Log(ss.Format())
	t.Log(ss.Args()...)
}
