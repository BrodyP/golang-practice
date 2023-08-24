package maps

import "testing"

//first look
// func TestDictionary(t *testing.T) {
// 	dictionary := map[string]string{"test": "this is test"}
// 	got := Search(dictionary, "test")
// 	want := "this is test"
// 	assertion(t, got, want)

// }

//add assertion
// func TestDictionary(t *testing.T) {
// 	dictionary := Dictionary{"test": "this is test"}
// 	got := dictionary["test"]
// 	want := "this is test"
// 	assertion(t, got, want)
// }

// func assertion(t testing.TB, got string, want string) {
// 	t.Helper()
// 	if got != want {
// 		t.Errorf("got %q want %q given, %q", got, want, "test")
// 	}
// }

//combine test and add errors

func TestDictionary(t *testing.T) {

	assertStrings := func(t testing.TB, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	assertDefinition := func(t testing.TB, dictionary Dictionary, word, definition string) {
		t.Helper()

		got, err := dictionary.Search(word)
		if err != nil {
			t.Fatal("should find added word:", err)
		}
		assertStrings(t, got, definition)
	}

	assertError := func(t testing.TB, got, want error) {
		t.Helper()

		if got != want {
			t.Errorf("got error %q want %q", got, want)
		}
	}

	t.Run("word found", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is test"}

		//refactor
		// got, _ := dictionary.Search("test")
		// want := "this is test"
		// assertStrings(t, want, got)

		assertDefinition(t, dictionary, "test", "this is test")

	})

	t.Run("word not found", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is test"}
		_, err := dictionary.Search("not found")
		assertError(t, err, ErrNotFound)

		//before refactor
		// want := "could not find the word you were looking for"
		// if err == nil {
		// 	t.Error("expect error")
		// }
		// assertStrings(t, want, err.Error())
	})

	t.Run("add", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Add("test", "this is test")

		//refactor 01
		// got, err := dictionary.Search("test")
		// if err != nil {
		// 	t.Fatal("should find added word:", err)
		// }
		// want := "this is test"
		// assertStrings(t, got, want)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, "test", "this is test")
	})

	t.Run("update", func(t *testing.T) {
		definition := "test 01"
		word := "test"
		dictionary := Dictionary{}
		dictionary.Add(word, definition)
		assertDefinition(t, dictionary, word, definition)
		newDefinition := "test 02"
		dictionary.Update(word, newDefinition)
		assertDefinition(t, dictionary, word, newDefinition)

	})

	t.Run("update fail", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Update("word", "test 02")
		assertError(t, err, ErrWordNotExist)
	})

	t.Run("delete", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		dictionary.Add(word, "test 01")
		assertDefinition(t, dictionary, word, "test 01")
		dictionary.Delete(word)
		_, err := dictionary.Search(word)
		assertError(t, err, ErrNotFound)
	})

}
