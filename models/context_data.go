package models

type ContextData struct {
	IsLogged         bool   // know if user has a session
	User             User   // data of the user logged in
	VisitedUser      User   // use only for the home page of user /user it holds the data of the visited user
	ActiveTab        string // use to know which active tab in view currently /user
	VisitedUserRepos []Repo // use for /user?tab=repositories view, store list of repos.
	VisitedRepo      Repo   // when visitig /user/repo
	Readme           string // readme to be rendered on view if necessary, use on /user, /user/repo and file explorer
}
