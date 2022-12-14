package core

import "fmt"

type Service struct {
	repo *Repo
}

func NewService() *Service {
	return &Service{
		repo: NewRepo(),
	}
}

func (s *Service) Search(keyword string, language string) (SearchRes, error) {
	resp := SearchRes{
		Words: []WordRes{
			{
				ID:   1,
				Word: "Ốc rưởn",
				Type: "danh từ",
				Definitions: []Definition{
					{
						Meaning:  "he he",
						Examples: []string{"thằng bé cười ốc rưởn - he he", "cười ốc rưởn - hehe"},
					},
					{
						Meaning:  "ho ho",
						Examples: []string{"thằng bé cười ốc rưởn - ho ho", "cười ốc rưởn - hoho"},
					},
				},
			},
			{
				ID:   2,
				Word: "Ốc ốc",
				Type: "danh từ",
				Definitions: []Definition{
					{
						Meaning:  "he he",
						Examples: []string{"thằng bé cười ốc rưởn - he he", "cười ốc rưởn - hehe"},
					},
					{
						Meaning:  "ho ho",
						Examples: []string{"thằng bé cười ốc rưởn - ho ho", "cười ốc rưởn - hoho"},
					},
				},
			},
		},
	}
	return resp, nil
}

func (s *Service) Upsert(word UpsertWord) (int, error) {
	fmt.Println(word)
	return 0, nil
}
