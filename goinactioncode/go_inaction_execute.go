package goinactioncode

// Executer performs to run all the samples
type Executer interface {
	Run()
}

// executers represents A map of registered goinaction samples for running.
var executers = make(map[string]Executer)

// RegisterExecuter performs to register sample executer
func RegisterExecuter(name string, executer Executer) {
	executers[name] = executer
}

// RunSample performs to run the specified sample
func RunSample(executerName string) {
	executer, exists := executers[executerName]
	if !exists {
		executer = executers["MatchersSearch"]
	}

	executer.Run()
}
