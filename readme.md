# 说明

用于在 Windows 下，特定时间内，屏幕锁定时自动结束软件进程。

自用于下班时自动关闭 QQ 和微信等应用。

# 配置

配置信息在 config.json 中。

- stopSoftware 为进程名称，锁定屏幕自动关闭这些进程，通过任务管理器->详情信息查看。
- startSoftware 为软件路径，在解锁屏幕后自动启动这些软件。
- stopTime 为时间段，在锁定屏幕时会判断当前时间是否在这个时间段内，如果在，则关闭 stopSoftware 指定的进程。
- startTime 为时间段，在解锁屏幕时会判断当前时间是否在这个时间段内，如果在，则启动 startSoftware 指定的软件。

# 下载

https://github.com/Curtion/stop-software/releases/

- stopSoftware.tar.gz 是一个控制台应用，启动后会有一个控制台界面。
- stopSoftware_hidden.tar.gz 以隐藏控制台的方式运行，但关闭应用需要使用任务管理器。

# TODO

- 托盘运行
- 开机自启

# 注意

编译此项目需要 gcc 套件支持，参考：
https://zhuanlan.zhihu.com/p/149305469
