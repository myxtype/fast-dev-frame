package exporter

import (
	"io/ioutil"
	"testing"
)

func TestNewExporter(t *testing.T) {
	writer, err := NewExporter([]string{"ID", "昵称"})
	if err != nil {
		t.Fatal(err)
	}
	defer writer.Close()

	writer.Append([]string{"1", "@我不认输～2"})
	writer.Append([]string{"2", "很爱你"})

	f, err := writer.WriteAllSeekZero()
	if err != nil {
		t.Fatal(err)
	}

	// Upload file
	// todo
	data, err := ioutil.ReadAll(f)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s", data)
}
