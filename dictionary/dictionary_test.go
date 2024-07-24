package dictionary

import "testing"

func TestSearch(t *testing.T) {

	dictionary := Dictionary{"test": "this is just a test"}
	t.Run("known word", func(t *testing.T) {

		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
	})

	t.Run("word not found", func(t *testing.T) {
		_, err := dictionary.Search("test2")
		want := ErrWordNotFound

		assertError(t, err, ErrWordNotFound)
		assertStrings(t, err.Error(), want.Error())

	})

}

func TestAdd(t *testing.T) {

	dict := Dictionary{"test": "this is a test"}

	t.Run("add word", func(t *testing.T) {
		word := "test"
		definition := "this is a test"
		_, err := dict.Add("test2", "this is another test")

		assertError(t, err, nil)
		assertDefinition(t, dict, word, definition)
	})

	t.Run("add existing word", func(t *testing.T) {
		word := "test"
		definition := "this is a test"
		_, err := dict.Add(word, definition)

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dict, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	dict := Dictionary{"test": "this is a test"}

	t.Run("new word update", func(t *testing.T) {
		word := "test2"
		definition := "this is a test2"
		got := dict.Update(word, definition)
		want := ErrWordDoesNotExist
		assertError(t, got, want)
	})
	t.Run("existing word update", func(t *testing.T) {
		word := "test"
		newDefinition := "brand new test"
		err := dict.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dict, word, newDefinition)
	})
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("expected error, got nil")
	}

	assertStrings(t, got.Error(), want.Error())
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dict Dictionary, word, definition string) {
	t.Helper()
	got, _ := dict.Search(word)
	want := definition

	assertStrings(t, got, want)

}
