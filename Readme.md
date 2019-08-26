# alioss-uploader

## 简介

alioss-uploader可以让你的alisso变成mweb兼容的图床

## 食用方法

1. docker build . -t pk8995/alioss-uploader 或者从docker hub上pull
2. 根据你自己的情况修改config.json中的配置
3. `./startup.sh `
4. mweb配置
   1. name: 写一个你喜欢的名字
   2. API URL: http://127.0.0.1:9002/upload (此处需与config中PostPath)
   3. POST File Name: file
   4. Response URL Path: Url

## Config
    UseInternalEndPoint 为true是将使用internal endpoint 地址上传，如果是在本机使用请设置为false
   

## 免责声明
poc代码，请随意使用后果自负。
