package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Primera suma que hacemos", func(t *testing.T) {
		got := Suma(5, 10)
		want := 15

		if got != want {
			t.Errorf("Obtenemos %q queremos obtener %q", got, want)
		}

	})
}
