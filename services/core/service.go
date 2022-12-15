package core

import (
	"fmt"
	"github.com/templateOfService/models"
	"github.com/templateOfService/pkg/vietnamese"
	"sort"
	"strings"
	"time"
)

type Service struct {
	repo *Repo
}

var CachedWords []models.Word

func NewService() *Service {
	return &Service{
		repo: NewRepo(),
	}
}

func (s *Service) LoadCache() {
	for {
		words, err := s.repo.GetAllWords()
		if err != nil {
			fmt.Println("Error loading words: ", err.Error())
			return
		}

		definitions, err := s.repo.GetAllDefinitions()
		if err != nil {
			fmt.Println("Error loading definitions: ", err.Error())
			return
		}

		for i := range words {
			words[i].NormalizedWord = vietnamese.RemoveAccent(words[i].Word)
			for _, definition := range definitions {
				if definition.WordID != words[i].ID {
					continue
				}

				words[i].Definitions = append(words[i].Definitions, definition)
				words[i].NormalizedDefinitions += vietnamese.RemoveAccent(definition.Meaning + " " + definition.Examples)
			}
		}

		CachedWords = words
		fmt.Println("Loaded words: ", len(CachedWords))
		time.Sleep(5 * time.Second)
	}
}

func (s *Service) Search(keyword string, prefix string, language string, order string) (SearchRes, error) {
	var res SearchRes
	var notPrioritizedWords []WordRes

	prefix = vietnamese.RemoveAccent(prefix)
	keyword = vietnamese.RemoveAccent(keyword)

	if len(keyword) == 0 && len(prefix) == 0 && len(language) == 0 {
		for _, word := range CachedWords {
			res.Words = append(res.Words, toWordRes(word))
		}
	} else {
		for _, word := range CachedWords {
			if len(language) > 0 && word.Language != language {
				continue
			}
			if len(prefix) > 0 {
				if strings.HasPrefix(word.NormalizedWord, prefix) {
					res.Words = append(res.Words, toWordRes(word))
				}
				continue
			}
			if len(keyword) > 0 {
				if strings.Contains(word.NormalizedWord, keyword) {
					res.Words = append(res.Words, toWordRes(word))
					continue
				}
				if strings.Contains(word.NormalizedDefinitions, keyword) {
					notPrioritizedWords = append(notPrioritizedWords, toWordRes(word))
				}
			}
		}
	}

	res.Words = append(res.Words, notPrioritizedWords...)
	if order == "asc" {
		sort.Slice(res.Words, func(i, j int) bool {
			return res.Words[i].Word < res.Words[j].Word
		})
	}

	if order == "desc" {
		sort.Slice(res.Words, func(i, j int) bool {
			return res.Words[i].Word > res.Words[j].Word
		})
	}

	return res, nil
}

func toWordRes(word models.Word) WordRes {
	var definitions []Definition
	for _, definition := range word.Definitions {
		definitions = append(definitions, Definition{
			Meaning:  definition.Meaning,
			Examples: strings.Split(definition.Examples, "|"),
		})
	}

	return WordRes{
		ID:          word.ID,
		Word:        word.Word,
		Language:    word.Language,
		Type:        word.Type,
		Definitions: definitions,
	}
}

func (s *Service) Upsert(input UpsertWord) (int, error) {
	word, definitions := transform(input)
	if word.ID > 0 {
		err := s.repo.UpdateWord(word)
		if err != nil {
			fmt.Println("Error updating word: ", err.Error())
			return 0, err
		}

		err = s.repo.ModifyDefinitions(word.ID, definitions)
		if err != nil {
			fmt.Println("Error modifying definitions: ", err.Error())
			return 0, err
		}

		return word.ID, nil
	}

	wordID, err := s.repo.InsertWord(word)
	if err != nil {
		fmt.Println("Error inserting word: ", err.Error())
		return 0, err
	}

	for i := range definitions {
		definitions[i].WordID = wordID
	}

	err = s.repo.ModifyDefinitions(wordID, definitions)
	if err != nil {
		fmt.Println("Error modifying definitions: ", err.Error())
		return 0, err
	}

	return wordID, nil
}

func (s *Service) List(ids []int) (SearchRes, error) {
	var res SearchRes
	tracer := make(map[int]struct{})

	for _, id := range ids {
		tracer[id] = struct{}{}
	}

	for _, word := range CachedWords {
		if _, ok := tracer[word.ID]; ok {
			res.Words = append(res.Words, toWordRes(word))
		}
	}
	return res, nil
}

func transform(input UpsertWord) (models.Word, []models.Definition) {
	var definitions []models.Definition
	for _, d := range input.Definitions {
		definitions = append(definitions, models.Definition{
			WordID:   input.ID,
			Meaning:  d.Meaning,
			Examples: strings.Join(d.Examples, "|"),
		})
	}
	return models.Word{
		ID:       input.ID,
		Word:     input.Word,
		Type:     input.Type,
		Language: input.Language,
	}, definitions
}
