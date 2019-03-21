// 对结构体进行&取地址操作时，视为对该类型进行一次new实例化操作
ins := &T{}

type Command struct {
	Name    string
	Var     *int
	Command string
}

var version int = 1

cmd := &Command{}
cmd.Name = "version"
cmd.Var = &version
cmd.Comment = "show version"

// 使用函数封装上面的初始化过程：

func newCommand(name string, varref *int, commen string) *Command {
	return &Command{
		Name : name,
		Var : varref,
		Comment : comment,
	}
}

cmd = newCommand(
	"version",
	&version,
	"show version",
)