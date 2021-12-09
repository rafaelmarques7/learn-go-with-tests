package maps

import "testing"

func TestDictionary(t *testing.T) {
	dict := Dictionary{"apple": "apple is a type of fruit"}

	t.Run("Search word that exists in dictionary", func(t *testing.T) {
		word, definition := "apple", "apple is a type of fruit"
		assertDefinition(t, dict, word, definition)
	})

	t.Run("Search word that does not exist in dictionary", func(t *testing.T) {
		_, err := dict.Search("banana")
		assertError(t, err, ErrWordNotFound)
	})

	t.Run("Adding words to dictionary", func(t *testing.T) {
		word := "pizza"
		definition := "a type of food"

		_, err := dict.Search(word)
		assertError(t, err, ErrWordNotFound)

		dict.AddWord(word, definition)
		assertDefinition(t, dict, word, definition)
	})

	t.Run("Adding words which already exist results in error", func(t *testing.T) {
		word := "apple"
		definition := "new definition"
		err := dict.AddWord(word, definition)
		assertError(t, err, ErrWordExists)
	})

	t.Run("Updating the definition of an existing word in the dictionary", func(t *testing.T) {
		word := "apple"
		definition := "apple is a type of fruit, but could also refer to the computer companny Apple"
		err := dict.UpdateWord(word, definition)
		assertNoError(t, err)
		assertDefinition(t, dict, word, definition)
	})

	t.Run("Updating the definition of a word that does not exist results in error", func(t *testing.T) {
		word := "banana"
		definition := "banana is a type of fruit"
		err := dict.UpdateWord(word, definition)
		assertError(t, err, ErrCanNotUpdate)
	})

	t.Run("Deleting a word from the dictionary", func(t *testing.T) {
		word := "apple"
		errDelete := dict.Delete(word)
		assertNoError(t, errDelete)

		_, errSearch := dict.Search(word)
		assertError(t, errSearch, ErrWordNotFound)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("strings do not match. got %q, want %q", got, want)
	}
}

func assertError(t testing.TB, errGotten, errExpected error) {
	t.Helper()
	if errGotten == nil {
		t.Fatalf("Expected error (%q) but did not get one", errExpected)
	}
	if errExpected != errGotten {
		t.Errorf("Error message does not match. Got %q, want %q", errGotten, errExpected)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("Got an error but no error was expected")
	}
}

func assertDefinition(t testing.TB, dict Dictionary, word, expectedDefinition string) {
	t.Helper()
	dictDefinition, err := dict.Search(word)
	assertNoError(t, err)
	assertStrings(t, dictDefinition, expectedDefinition)
}
