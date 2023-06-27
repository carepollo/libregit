package models

type ContextData struct {
	IsLogged    bool   // know if user has a session
	User        User   // data of the user logged in
	VisitedUser User   // use only for the home page of user /user it holds the data of the visited user
	ActiveTab   string // use only for the home page of user /user is to know which tab is currently active
	// Readme string // readme to be rendered on view if necessary, use on /user, /user/repo and file explorer
}
