package github

type Repo struct {
	Name         string `json:"name,omitempty"`
	LanguagesUrl string `json:"languages_url,omitempty"`
}

type ReposService struct {
	client    *Client
	languages *LanguagesService
}

func (s *ReposService) List(url string) (repos *[]Repo, err error) {
	repos = new([]Repo)
	_, err = s.client.GetUrl(url, repos)
	return
}
