package message

import ."github.com/job-center/pkg/cache"
//Generate a new message pool
func (p *Pool) NewPool() {
	p.Cache = Register("myRepo")
}

//Push a new message to pool
func (p *Pool) PushMessage(key string, val interface{}) {
	p.Put(key, val)
}

//Fetch the specific message from pool
func (p *Pool) FetchMessage(key string, val interface{}) {
	p.Get(key)
}

//Traverse the pool to fetch all fatal message
func (p *Pool) FetchAllFatalMessage() []interface{} {
	return p.MatchQueryKey("fatal")
}

//Traverse the pool to fetch all error message
func (p *Pool) FetchAllErrorMessage() []interface{} {
	return p.MatchQueryKey("error")
}

//Traverse the pool to fetch all warn message
func (p *Pool) FetchAllWarnMessage() []interface{} {
	return p.MatchQueryKey("warn")
}

//Traverse the pool to fetch all debug message
func (p *Pool) FetchAllDebugMessage() []interface{} {
	return p.MatchQueryKey("debug")
}

//Traverse the pool to fetch all info message
func (p *Pool) FetchAllInfoMessage() []interface{} {
	return p.MatchQueryKey("info")
}
