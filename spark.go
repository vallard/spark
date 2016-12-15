package spark

type Spark struct {
	token string
}

func New(token string) *Spark {
	return &Spark{
		token: token,
	}
}
