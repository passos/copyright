# 版权保护交易 PRD



## Overview

本系统的目标是对版权进行所有权证明和存在证明的认证和存证，同时支持对被保护内容的自助转载授权。通过使用区块链的防篡改特性来实现对文章和图片内容的存证。



## 系统结构

![系统结构](/Users/simon/Documents/区块链/课程/DAPP301/Copyright/img/系统结构.png)

- 内容认证
  - 原始数据内容存储在中心数据库中
  - 原始数据的元数据、内容 Hash、认证时间戳，经过密钥签名后上链
  - 每个数据内容的密钥单独生成，认证完成后发送给用户（可通过用户密码二次加密），服务端只保存公钥，不保存私钥
  - 认证结果生成一个唯一 hash (DNA)，是可信时间戳和内容 hash 的计算结果
  - 以 DNA 为 key，原始内容存储到数据库中
- 内容验证
  - 通过 DNA + 签名密钥，验证时间戳和内容 hash
- 内容提取
  - 内容外链为一个 URL，其中包含内容 DNA，例如 https://yuanben.io/image/63LGYLAR6ASOJWV1D5181ZZFW1PWVI5VL9OW5HNU5UMGY33ULL
  - 以每个内容的 DNA 为 key，数据库提取对应原始内容
- 转载授权
  - 每次转载授权，根据转载协议付费或者免费使用
  - 转载内容的链接中包含一个key，对应到授权协议和原始内容
  - 通过外链地址的请求，可以统计到内容的引用源站(Refer URL)，授权用户等信息。在此基础上可以实现转载监控
- 中心数据库
  - 存储原始内容和元数据，包括区块链地址、DNA、签名密钥等
  - 用户数据（用户名、密码、实名认证信息等）



## 所有权证明和存在性证明逻辑链

![证明逻辑链](/Users/simon/Documents/区块链/课程/DAPP301/Copyright/img/证明逻辑链.png)



### 参考

http://yuanbenlian.mydoc.io/docs/api.md?t=268053

Article metadata fields https://github.com/primasio/dtcp/blob/master/draft/3_article.md

http://dublincore.org/

## 系统模块

### 账户注册、登录

邮箱注册，使用第三方认证接口实名认证(身份证号+姓名)

其它略...

### 认证

![认证原创流程](/Users/simon/Documents/区块链/课程/DAPP301/Copyright/img/认证原创流程.png)

入口

![认证原创 00](/Users/simon/Documents/区块链/课程/DAPP301/Copyright/img/认证原创 00.png)

认证文章

![认证原创 01-1](/Users/simon/Documents/区块链/课程/DAPP301/Copyright/img/认证原创 01-1.png)

认证图片

![认证原创 01-2](/Users/simon/Documents/区块链/课程/DAPP301/Copyright/img/认证原创 01-2.png)

上传完成

![认证原创 02](/Users/simon/Documents/区块链/课程/DAPP301/Copyright/img/认证原创 02.png)

选择转载协议

![认证原创 03](/Users/simon/Documents/区块链/课程/DAPP301/Copyright/img/认证原创 03.png)

完成认证

![认证原创 04](/Users/simon/Documents/区块链/课程/DAPP301/Copyright/img/认证原创 04.png)

点击区块链地址

![认证原创 05-1](/Users/simon/Documents/区块链/课程/DAPP301/Copyright/img/认证原创 05-1.png)

点击区块链条目

![认证原创 05-2](/Users/simon/Documents/区块链/课程/DAPP301/Copyright/img/认证原创 05-2.png)

认证详情

![认证原创 06](/Users/simon/Documents/区块链/课程/DAPP301/Copyright/img/认证原创 06.png)

认证内容

![认证原创 07](/Users/simon/Documents/区块链/课程/DAPP301/Copyright/img/认证原创 07.png)

内容列表

![内容列表](/Users/simon/Documents/区块链/课程/DAPP301/Copyright/img/内容列表.png)



### 验证

验证输入

![认证原创 08](/Users/simon/Documents/区块链/课程/DAPP301/Copyright/img/认证原创 08.png)

验证结果

![认证原创 09](/Users/simon/Documents/区块链/课程/DAPP301/Copyright/img/认证原创 09.png)

## 后端 API 接口设计

> 所有 API 前缀 /api/v1

### /account

POST 创建账户

### /account/:id

支持 RESTful 的 GET/POST/PUT/DELETE 操作

### /content

POST 发布内容

### /content/:DNA

GET 获取内容

### /metadata

POST 发布元数据

### /metadata/:DNA

GET 获取元数据

### /verify

POST DNA, 签名私钥 返回验证结果









