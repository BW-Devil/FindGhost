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
抓包库：libpcap  
数据库：mongodb  
### :ghost: 使用
#### 安装图形化库、编译环境等：
```shell
Debian/Ubuntu: sudo apt-get install golang gcc libgl1-mesa-dev xorg-dev
Fedora: sudo dnf install golang gcc libXcursor-devel libXrandr-devel mesa-libGL-devel libXi-devel libXinerama-devel libXxf86vm-devel
Arch Linux: sudo pacman -S go xorg-server-devel libxcursor libxrandr libxinerama libxi
Solus: sudo eopkg it -c system.devel golang mesalib-devel libxrandr-devel libxcursor-devel libxi-devel libxinerama-devel
openSUSE: sudo zypper install go gcc libXcursor-devel libXrandr-devel Mesa-libGL-devel libXi-devel libXinerama-devel libXxf86vm-devel
Void Linux: sudo xbps-install -S go base-devel xorg-server-devel libXrandr-devel libXcursor-devel libXinerama-devel
```  
#### 启动各个组件
1.进入crawler文件夹，使用命令`./Crawler web`启动爬虫，更新数据  
2.进入analyser文件夹，使用命令`./Analyser start`启动分析器  
3.进入catcher文件夹，使用命令`./Catcher catch`启动嗅探器  
启动后，打开网页随便输入一个网址，测试一下流量的获取  
![image](https://user-images.githubusercontent.com/90563485/192080528-2ffc4dfe-1558-461d-9045-ad69bd08a1e7.png)
![image](https://user-images.githubusercontent.com/90563485/192080538-851d6f9d-d9a4-4b84-9429-d53e3702891a.png)
在浏览器输入：`http://127.0.0.1:7777/`，查看捕获的恶意ip,恶意dns,恶意http等信息  
![image](https://user-images.githubusercontent.com/90563485/192080576-7949ef75-0d02-4945-b50b-f12098469a8c.png)
![image](https://user-images.githubusercontent.com/90563485/192080584-3660a993-2ea3-4562-88c3-30eece957313.png)
### :sheep: TODO
* ~~嗅探器的完善~~  
* ~~分析器的完善~~  
* ~~数据库的补充~~  
* 数据库的改进  
* 网页界面的优化  
