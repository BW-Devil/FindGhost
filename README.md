# FindGhost
### :bulb: 简介
FindGhost是一款恶意流量分析系统，会实时分析当前网络中的流量信息，检测出恶意ip或恶意域名。  
对恶意ip或恶意域名会做记录，在网页端可以实时查看记录的信息。  
### :low_brightness: 结构
该项目整体有三部分，嗅探器、分析器、数据库。  
嗅探器用于嗅探网络中的流量，抓取ip等信息传送给分析器。  
分析器分析嗅探器传送过来的数据，将样本数据与数据库中的数据进行比对，分析该流量是否是恶意的。  
数据库其实是一个网络爬虫，会定时抓取netlab360、greensnow等公开的恶意ip库。  
![image](https://user-images.githubusercontent.com/90563485/190836364-e9b7c979-bd6a-4226-98b4-339bf4e0c518.png)
### :computer: 环境
运行环境：Linux  
安装图形化库、编译环境等：  
`Debian / Ubuntu: sudo apt-get install golang gcc libgl1-mesa-dev xorg-dev`  
`Fedora: sudo dnf install golang gcc libXcursor-devel libXrandr-devel mesa-libGL-devel libXi-devel libXinerama-devel libXxf86vm-devel`  
`Arch Linux: sudo pacman -S go xorg-server-devel libxcursor libxrandr libxinerama libxi`  
`Solus: sudo eopkg it -c system.devel golang mesalib-devel libxrandr-devel libxcursor-devel libxi-devel libxinerama-devel`  
`openSUSE: sudo zypper install go gcc libXcursor-devel libXrandr-devel Mesa-libGL-devel libXi-devel libXinerama-devel libXxf86vm-devel`  
`Void Linux: sudo xbps-install -S go base-devel xorg-server-devel libXrandr-devel libXcursor-devel libXinerama-devel`  
### :ghost: 使用
目前只完成了crawler部分，crawler可以直接以api的形式进行使用，也可以使用gui的方式进行使用，gui的编写使用的是fyne框架。  
`./Crawler help`
```shell
NAME:
   Crawler - FindGhost Crawler

USAGE:
   Crawler [global options] command [command options] [arguments...]

VERSION:
   1.0.0

DESCRIPTION:
   FindGhost Crawler

AUTHOR:
   BWFish <weunknowing@gmail.com>

COMMANDS:
   web      start up web program
   gui      start up gui program
   dump     fetch evil ips and domains to file
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)
```
`./Crawler web`：启动api模式    
在网页端访问:  
`http://127.0.0.1:8888/api/ip/x.x.x.x,检测x.x.x.x是否为恶意ip`  
`http://127.0.0.1:8888/api/domain/x.x.x.x,检测x.x.x.x是否为恶意domain`  
说明：访问端口，可以在conf/app.ini中修改HTTP_PORT的值 
![image](https://user-images.githubusercontent.com/90563485/190837051-5d1f2859-cf11-479c-9701-1b9f1a875922.png)
![image](https://user-images.githubusercontent.com/90563485/190837067-282c950a-e8fa-4689-8d77-d9a3c57d2bf2.png)
`./Crawler gui`：启动gui模式  
说明：要先启动api模式让数据库进行更新，再启动gui模式  
![image](https://user-images.githubusercontent.com/90563485/190837193-94d2f46f-1e6a-4376-b791-f840b6353e5e.png)
### :sheep: TODO
* 嗅探器的完善
* 分析器的完善
* 数据库的补充
