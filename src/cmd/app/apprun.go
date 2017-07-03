package app

import (
	"sync"
	"time"
	"common"
)

//var Switch *bool
//var switchTemp bool
//var routineSwitch chan bool
var ThreadCount sync.WaitGroup
var statusSwitchOn chan bool
var statusSwitchOff chan bool
var statusSwitchLast *bool

func init() {
	//Switch = new(bool)
	//*Switch = true
	//switchTemp = *Switch
	statusSwitchLast = new(bool)
	*statusSwitchLast = true
}

func Run() (err error) {
	statusSwitchOn = make(chan bool)
	statusSwitchOff = make(chan bool)
	//routineSwitch = make(chan bool)
	collectMainInOnCycle()
	return nil
}

func TurnStatus(status bool) {
	switch status {
	case true:
		common.DebugPrint("turn the SwitchOn")
		statusSwitchOn <- true
	case false:
		common.DebugPrint("turn the SwitchOff")
		statusSwitchOff <- false
	}
}

func runOneCycle() {
	//routineSwitch <- *Switch
	collect.RunOneCycle()
}
func getState() (s bool) {
	//s = *Switch
	return s
}

/*func getStateFromRest() error {
	if switchTemp != *Switch {
		switchTemp = *Switch
		routineSwitch <- *Switch
		common.DebugPrint("run with the rest", "routineSwitch is", routineSwitch, "*Switch is ")

	}
	return nil
}
func turnState() {
	bool := <-routineSwitch
	common.DebugPrint("turnState bool is ", bool)
	if bool == true {
		ThreadCount.Done()
	}
}
*/
func collectMainInOnCycle() {
	//var i = 0
	ticker := time.NewTicker(time.Millisecond * 5000)
	common.DebugPrint("main routine is run")
	for range ticker.C {
		ThreadCount.Add(1)
		common.DebugPrint("run with the state", "statusSwitchLast is", *statusSwitchLast)
		select {
		case i := <-statusSwitchOn:
			if i != *statusSwitchLast {
				common.DebugPrint("into the select thread in statusSwitchOn")
				go runOneCycle()
				*statusSwitchLast = true
				ThreadCount.Done()
			}
		case i := <-statusSwitchOff:
			if i != *statusSwitchLast {
				common.DebugPrint("into the select thread in statusSwitchOff")
				*statusSwitchLast = false
				ThreadCount.Done()
			}
		default:
			switch *statusSwitchLast {
			case true:
				common.DebugPrint("into the select thread in default on")
				go runOneCycle()
				ThreadCount.Done()
			case false:
				common.DebugPrint("into the select thread in default off")
				ThreadCount.Done()
			}
			/* i = i + 1
			ThreadCount.Add(1)
			common.DebugPrint("thread is run at", i, routineSwitch, *Switch)
			go runOneCycle()
			common.DebugPrint("oneCycle thread is run ", routineSwitch, *Switch)
			go turnState()
			common.DebugPrint("Information synchronization is run ", routineSwitch, *Switch)
			go common.Tick("freedom", 100, "Loop read rest", getStateFromRest)
			common.DebugPrint("collect rest is run ", routineSwitch, *Switch)
			ThreadCount.Wait()
		} else {

		} */
		}
	}
}
