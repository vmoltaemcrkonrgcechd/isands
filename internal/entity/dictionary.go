package entity

type DictionaryEntryDTO struct {
	Name string `json:"name"`
}

type DictionaryEntry struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Dictionary []DictionaryEntry
