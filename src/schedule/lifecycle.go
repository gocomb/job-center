package schedule

import (
	"time"
	"util"
)

func (c *JobRunTime) restDelayPeriod(ticker *time.Ticker) {
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
		}
	}
}
func (c *JobRunTime) restConstantPeriod(ticker *time.Ticker) {
	c.logger.SetInfo("main routine is run")
	var i = 0
	i = i + 1
	c.ThreadCount.Add(1)
	c.logger.SetInfo("thread is run at", i, c.routineSwitch, c.*Switch)
	go c.runOneCycle()
	c.logger.SetInfo("oneCycle thread is run ", c.routineSwitch, c.*Switch)
	go c.turnState()
	c.logger.SetInfo("Information synchronization is run ", c.routineSwitch, c.*Switch)
	go util.Tick("freedom", 100, "Loop read rest", c.getStateFromRest)
	c.logger.SetInfo("collect rest is run ", c.routineSwitch, c.*Switch)
	c.ThreadCount.Wait()

}
