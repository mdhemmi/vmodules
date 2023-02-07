package getfuncname

import "runtime"

func GetCurrentFuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
	//return fmt.Sprintf("%s", runtime.FuncForPC(pc).Name())
}
