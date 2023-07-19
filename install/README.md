## 基本运行环境
- windows10+
## 安装手册
### 运行时环境准备
- 解压安装包到目标目录`${target_path}`
- 打开`${target_path}`目录
- 在`${target_path}`下安装`Docker Desktop Installer.exe`
- 以超级管理员方式打开 PowerShell`(“开始”菜单 >“PowerShell” >单击右键 >“以管理员身份运行”)`,运行一下命令
```
dism.exe /online /enable-feature /featurename:VirtualMachinePlatform /all /norestart
```
- 在`${target_path}`下安装`wsl_update_x64.msi`
- 以超级管理员方式打开 PowerShell`(“开始”菜单 >“PowerShell” >单击右键 >“以管理员身份运行”)`,运行一下命令
```
wsl --set-default-version 2
```
- 重启电脑
### 服务启动
- 启动`Docker Desktop`
- 以超级管理员方式打开 PowerShell`(“开始”菜单 >“PowerShell” >单击右键 >“以管理员身份运行”)`,执行load镜像命令
```
cd ${target_path} 
docker load -i ${xxxx}.tar
```
- 以超级管理员方式打开 PowerShell`(“开始”菜单 >“PowerShell” >单击右键 >“以管理员身份运行”)`,执行启动命令
```
cd ${target_path} 
docker compose up -d 
```
