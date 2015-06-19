package lawparser

type legalDocument struct {
	Name          string    `json:"name"`
	Country       string    `json:"country"`
	State         string    `json:"state"`
	Town          string    `json:"town"`
	Order         string    `json:"order"`         //municipal
	LegalCategory string    `json:"legalCategory"` //reglamento
	Topic         string    `json:"topic"`         //transito
	Url           string    `json:"url"`           //transito
	Header        string    `json:"header"`
	Footer        string    `json:"footer"`
	Articles      []article `json:"articles"`
}
type articles []article
type article struct {
	Id        string       `json:"id"`
	Num       string       `json:"num"`
	Parents   []parent     `json:"parents"`
	Headers   []subarticle `json:"headers"`
	Fractions []fraction   `json:"fractions"`
	Footers   []subarticle `json:"footers"`
}

func (self *article) setParent(p parent) {
	self.Parents = append(self.Parents, p)
}

type parent struct {
	Order string `json:"order"`
	Num   string `json:"num"`
	Title string `json:"title"`
}

type subarticle struct {
	Body    string        `json:"body"`
	Deontic string        `json:"deontic"`
	Sub     []subfraction `json:"sub"`
	Punish  punishment    `json:"punish"`
}

type fraction struct {
	subarticle
	Num string `json:"num"`
}

type subfraction struct {
	Num  string `json:"num"`
	Body string `json:"body"`
}
type punishment struct {
	Type string `json:"type"`
	From string `json:"from"`
	To   string `json:"to"`
}
