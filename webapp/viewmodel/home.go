package viewmodel

type Home struct {
	Title string
	Active string
}

func NewHome() Home {
	return Home {
		Title: "Lemonade Stand Supply",
		Active: "home",
	}
}