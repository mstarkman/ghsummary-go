package github

type Repo struct {
	Name         string `json:"name,omitempty"`
	LanguagesUrl string `json:"languages_url,omitempty"`
}

type ReposService struct {
	client *Client
}

func (s *ReposService) List(url string) (*[]Repo, error) {
	repos := new([]Repo)

	_, err := s.client.GetUrl(url, repos)

	return repos, err
}
