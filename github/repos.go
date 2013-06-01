package github

type Repo struct {
	Name string `json:"name,omitempty"`
}

type ReposService struct {
	client *Client
}

func (s *ReposService) List(url string) (*[]Repo, error) {
	repos := new([]Repo)

	_, err := s.client.GetUrl(url, repos)

	return repos, err
}
