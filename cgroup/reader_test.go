package cgroup

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReaderGetStats(t *testing.T) {
	reader, err := NewReader("testdata/docker", true)
	if err != nil {
		t.Fatal(err)
	}

	stats, err := reader.GetStatsForProcess(985)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "testdata/docker/sys/fs/cgroup/blkio/docker/b29faf21b7eff959f64b4192c34d5d67a707fe8561e9eaa608cb27693fba4242", stats.BlockIO.Path)
	assert.Equal(t, "testdata/docker/sys/fs/cgroup/cpu/docker/b29faf21b7eff959f64b4192c34d5d67a707fe8561e9eaa608cb27693fba4242", stats.CPU.Path)
	assert.Equal(t, "testdata/docker/sys/fs/cgroup/cpuacct/docker/b29faf21b7eff959f64b4192c34d5d67a707fe8561e9eaa608cb27693fba4242", stats.CPUAccounting.Path)
	assert.Equal(t, "testdata/docker/sys/fs/cgroup/memory/docker/b29faf21b7eff959f64b4192c34d5d67a707fe8561e9eaa608cb27693fba4242", stats.Memory.Path)

	json, err := json.MarshalIndent(stats, "", "  ")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(json))
}
