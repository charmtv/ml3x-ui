package service

import (
	"encoding/json"
	"testing"

	"github.com/mhsanaei/3x-ui/v2/database/model"
	"github.com/mhsanaei/3x-ui/v2/util/json_util"
)

func TestEnsureAccessLogEnabledSetsDefaultWhenDisabled(t *testing.T) {
	raw := json_util.RawMessage(`{"access":"none","error":"","loglevel":"warning"}`)

	got := ensureAccessLogEnabled(raw)

	var parsed map[string]any
	if err := json.Unmarshal(got, &parsed); err != nil {
		t.Fatalf("invalid json: %v", err)
	}
	if parsed["access"] != "./access.log" {
		t.Fatalf("expected access log to be enabled, got %v", parsed["access"])
	}
	if parsed["loglevel"] != "warning" {
		t.Fatalf("expected existing loglevel to be preserved, got %v", parsed["loglevel"])
	}
}

func TestEnsureAccessLogEnabledPreservesCustomPath(t *testing.T) {
	raw := json_util.RawMessage(`{"access":"/var/log/xray/access.log"}`)

	got := ensureAccessLogEnabled(raw)

	var parsed map[string]any
	if err := json.Unmarshal(got, &parsed); err != nil {
		t.Fatalf("invalid json: %v", err)
	}
	if parsed["access"] != "/var/log/xray/access.log" {
		t.Fatalf("expected custom access log path to be preserved, got %v", parsed["access"])
	}
}

func TestClientsHaveLimitIncludesDeviceLimit(t *testing.T) {
	clients := []model.Client{{Email: "user", DeviceLimit: 1}}

	if !clientsHaveLimit(clients) {
		t.Fatalf("expected deviceLimit to activate limit tracking")
	}
}
