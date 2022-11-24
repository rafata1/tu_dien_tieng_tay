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
		},
	}
	return resp, nil
}

func (s *Service) AddWord(word AddWord) (int, error) {
	fmt.Println(word)
	return 0, nil
}
