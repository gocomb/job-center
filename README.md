####构想：
1  每个job都有计时器，皆可以固定时间间隔进行周期运行。
2  job可由多种方式触发，如restapi发送开关信息、etcd，rpc等消息触发
3  job的外部触发可以进行job的开始和终止，设置周期参数。
##每个包都总用
1  message包：job所用channel的存储层，外部层序通过调用message包访问job传来的信息，message
2  monitor包，根据message触发一些job控制方法
3  util 封装好的工具包
