package bitmex

import (
    "testing"
)

func TestGetInstruments(t *testing.T) {
    client := New("", "")
    _, err := client.GetInstruments()
    if err != nil {
        t.Fatal(err)
    }
}

func TestGetActiveInstruments(t *testing.T) {
    client := New("", "")
    _, err := client.GetActiveInstruments()
    if err != nil {
        t.Fatal(err)
    }
}

func TestGetActiveInstrumentsAndIndices(t *testing.T) {
    client := New("", "")
    _, err := client.GetActiveInstrumentsAndIndices()
    if err != nil {
        t.Fatal(err)
    }
}

func TestGetIndices(t *testing.T) {
    client := New("", "")
    instruments, err := client.GetIndices()
    if err != nil {
        t.Fatal(err)
    }

    for _, i := range instruments {
        t.Logf("%+v\n", i.Symbol)
    }
}
