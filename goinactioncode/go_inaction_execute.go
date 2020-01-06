package goinactioncode

import (
	"goloveu/goinactioncode/executers"
)



// RunSample performs to run the specified sample
func RunSample(executerName string) {
	executer := executers.RetrieveExecuter(executerName)
	executer.Run()
}
