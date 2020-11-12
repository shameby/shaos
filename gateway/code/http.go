package code

type HTTPCode int

// 参数级错误
const (
	ValidatorErr HTTPCode = iota + 4001000
	ParamsIllegal
)

// 业务级错误
const (
	ModulesFuncErr HTTPCode = iota + 4002000
)
