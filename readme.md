# 说明

用于在 Windows 下，特定时间内，屏幕锁定时自动结束软件进程。

自用于下班时自动关闭 QQ 和微信等应用。

配置信息在 config.json 中。software 为进程名称，通过任务管理器->详情信息查看，time 为时间段，只有在该时间段内锁定屏幕才会结束进程。

# 下载

https://github.com/Curtion/stop-software/releases/

# 需求

- 解锁时自动启动被关闭的程序

  - 思路： 通过 wmi 获取进程的 PATH 和 WorkSpace 等参数，后续用这些参数启动程序。

- 托盘运行

# 注意

编译此项目需要 gcc 套件支持，参考：
https://zhuanlan.zhihu.com/p/149305469
