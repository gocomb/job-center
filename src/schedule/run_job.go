package schedule

import (
	"time"
)

func (c *InitJob) genInitVariable() {

}

func (a *jobOneCycle) runOneCycle() error {
	a.ThreadCount.Add(1)
	err := a.JobFunc()
	return err
}

func (c *JobRunTime) cycleRunTime() {
	//var i = 0
	ticker := time.NewTicker(time.Millisecond * 5000)
	c.logger.SetInfo("main routine is run")
	for range ticker.C {
		c.logger.SetInfo("run with the state", "statusSwitchLast is", c.*statusSwitchLast)
		select {
		case i := <-c.statusSwitchOn:
			if i != *c.statusSwitchLast {
				c.logger.SetInfo("into the select thread in statusSwitchOn")
				go c.runOneCycle()
				c.*statusSwitchLast = true
				c.ThreadCount.Done()
			}
		case i := <-c.statusSwitchOff:
			if i != c.*statusSwitchLast {
				c.logger.SetInfo("into the select thread in statusSwitchOff")
				c.*statusSwitchLast = false
				c.ThreadCount.Done()
			}
		default:
			switch c.*statusSwitchLast {
			case true:
				c.logger.SetInfo("into the select thread in default on")
				go c.runOneCycle()
				c.ThreadCount.Done()
			case false:
				c.logger.SetInfo("into the select thread in default off")
				c.ThreadCount.Done()
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
