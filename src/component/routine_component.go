//AppName: goblog
//Version: V1.0.0
//User: marco
//Date: 2019/1/1

package component

func GoRoutine(exec func())  {
	go exec()
}
