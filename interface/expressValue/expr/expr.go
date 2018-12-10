package expr

type Expr interface {
	Eval(env Env) float64
}

type Env map[Var]float64

type Var string

type literal float64

type unary struct {
	op rune
	x Expr
}
// + -
type binary struct {
	op rune
	x, y Expr
}

//+ - * /
type call struct {
	fun string
	args[]Expr
}

//pow	sin   sqrt
func (v Var) Eval(env Env) float64 {
	return env[v]
}

func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
}