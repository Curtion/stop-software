# 说明

用于在 Windows 下，特定时间内，屏幕锁定时自动结束软件进程。

自用于下班时自动关闭 QQ 和微信等应用。

# 配置

配置信息在 config.json 中。

- stopSoftware 为进程名称，锁定屏幕自动关闭这些进程，通过任务管理器->详情信息查看。
- startSoftware 为软件路径，在解锁屏幕后自动启动这些软件。
- time 为时间段，有两个作用。
  - 在锁定屏幕时会判断当前时间是否在这个时间段内，如果在，则关闭 stopSoftware 的进程，反之相反。
  - 在解锁屏幕时会判断当前时间是否在这个时间段内，如果不在，则启动 startSoftware 的软件，反之相反。

# 下载

https://github.com/Curtion/stop-software/releases/

# 需求

- 托盘运行

# 注意

编译此项目需要 gcc 套件支持，参考：
https://zhuanlan.zhihu.com/p/149305469
