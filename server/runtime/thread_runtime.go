package runtime

import (
	"time"
)

//job的一次以rest触发和周期性触发联合触发机制的完整生命周期
//rest为job开关，置开为立刻触发job，关为关闭job的运行，并不进行周期触发
//rest触发时对上次状态进行逻辑判断，若上次开关状态为开，再将开关置开不触发新的job
func (c *JobRunTime) runOnePeriod(ticker *time.Ticker) {
	c.logger.SetInfo("job routine is run, routine would be init")
	for range ticker.C {
		c.logger.SetInfo("run with the state", "statusSwitchLast is", *c.statusSwitchLast)
		select {
		case i := <-c.statusSwitchOn:
			if i != *c.statusSwitchLast {
				c.logger.SetInfo("into the select thread in statusSwitchOn")
				go c.runOneCycle()
				*c.statusSwitchLast = true
				c.ThreadCount.Done()
			}
		case i := <-c.statusSwitchOff:
			if i != *c.statusSwitchLast {
				c.logger.SetInfo("into the select thread in statusSwitchOff")
				*c.statusSwitchLast = false
				c.ThreadCount.Done()
			}
		default:
			switch *c.statusSwitchLast {
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
func (c *JobRunTime) restTriggerOn(now bool, Last bool) {
	if now != Last {
		c.logger.SetInfo("into the select thread in statusSwitchOn")
		go c.runOneCycle()
		Last = true
		c.ThreadCount.Done()
	}
}
func (c *JobRunTime) restTriggerOff(now bool, Last bool) () {
	if now != Last {
		c.logger.SetInfo("into the select thread in statusSwitchOff")
		*c.statusSwitchLast = false
		c.ThreadCount.Done()
	}


}

/*
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
*/
