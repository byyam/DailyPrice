# DailyPrice
Push stock/bitcoin price on duty.

### 使用引导

1. 创建一个企业微信群聊，右键群->创建群机器人，保存机器人的key。
2. 注册一个[coin market](https://coinmarketcap.com/api/)账号，保存api key。
3. 将企业微信的key填入pushkey的配置，coin market的key填入coin market key的配置项。
4. 在`cmd`目录执行`go build`命令进行编译操作，完成后执行`./cmd`程序即可完成推送。

