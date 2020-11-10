# 序列化对象自动生成工具
支持类型:  protobuf

## 1.google protobuf
### 1.1 安装 protoc
#### 1.1.1 直接下载release版本
&emsp;&emsp;protobuf/protoc/ 目录下有mac和win平台的生成工具，其他release版本下载地址：https://github.com/protocolbuffers/protobuf/releases
#### 1.1.2 自行编译
&emsp;&emsp;参考 https://github.com/protocolbuffers/protobuf
#### 1.1.3 HomeBrew
&emsp;&emsp;mac平台也可以brew安装

### 1.2 指令
&emsp;&emsp;参考 .vscode/launch.json
#### 1.2.1 -cmd
&emsp;&emsp;install:安装plugin : "golang"; "javascript";
&emsp;&emsp;gen:生成文件
#### 1.2.2 -sc
&emsp;&emsp;生成对应语言: "golang"; "javascript"; "typescript"
#### 1.2.3 -p
&emsp;&emsp;该工程文件路径
#### 1.2.4 -o
&emsp;&emsp;输出路径，默认 protobuf/gen/{sc}/