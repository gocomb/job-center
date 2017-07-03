package schedule

import (
	"util"
	"time"
)

func(a Job) runOneCycle()error{
	err:=a.JobFunc()
	return err
}
func collectMainInOnCycle() {
	//var i = 0
	ticker := time.NewTicker(time.Millisecond * 5000)
	util.DebugPrint("main routine is run")
	for range ticker.C {
		ThreadCount.Add(1)
		util.DebugPrint("run with the state", "statusSwitchLast is", *statusSwitchLast)
		select {
		case i := <-statusSwitchOn:
			if i != *statusSwitchLast {
				util.DebugPrint("into the select thread in statusSwitchOn")
				go runOneCycle()
				*statusSwitchLast = true
				ThreadCount.Done()
			}
		case i := <-statusSwitchOff:
			if i != *statusSwitchLast {
				util.DebugPrint("into the select thread in statusSwitchOff")
				*statusSwitchLast = false
				ThreadCount.Done()
			}
		default:
			switch *statusSwitchLast {
			case true:
				util.DebugPrint("into the select thread in default on")
				go runOneCycle()
				ThreadCount.Done()
			case false:
				util.DebugPrint("into the select thread in default off")
				ThreadCount.Done()
			}
			/* i = i + 1
			ThreadCount.Add(1)
			util.DebugPrint("thread is run at", i, routineSwitch, *Switch)
			go runOneCycle()
			util.DebugPrint("oneCycle thread is run ", routineSwitch, *Switch)
			go turnState()
			util.DebugPrint("Information synchronization is run ", routineSwitch, *Switch)
			go util.Tick("freedom", 100, "Loop read rest", getStateFromRest)
			util.DebugPrint("collect rest is run ", routineSwitch, *Switch)
			ThreadCount.Wait()
		} else {

		} */
		}
	}
}
