fabric-ca v1.4.3 国密改造
国密改造参照https://blog.csdn.net/dyj5841619/article/details/90642808
在fabric-ca 1.4.3上修改并完成调试； 并修改此项目vendor第三方库，所以请不要替换vendor目录

1 编译设置
设置export GO111MODULE=off 关闭go modele功能

2 编译fabric-ca-server
make fabric-ca-server
二进制生层bin/fabric-ca-server


3 编译fabric-ca-client
make fabric-ca-client

至此，二进制fabric编译完成
