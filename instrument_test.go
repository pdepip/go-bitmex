package bitmex

import (
    "testing"
)

func TestGetInstruments(t *testing.T) {
    client := New("", "")
    instruments, err := client.GetInstruments()
    if err != nil {
        t.Fatal(err)
    }

    t.Logf("%+v\n", instruments)
}

func TestGetActiveInstruments(t *testing.T) {
    client := New("", "")
    instruments, err := client.GetActiveInstruments()
    if err != nil {
        t.Fatal(err)
    }

    t.Logf("%+v\n", instruments)
}

