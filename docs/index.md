# 接口文档


## 接口说明

`基础地址` **http://127.0.0.1:8000/v1/cmb/**


::: tip

1. 所有接口均为 `POST` 请求
2. 所有请求参数均为 `string` 类型

:::


::: warning

关键字段说明:
1. `sub_account_no` 记账单元编号  主账号 + 记账单元编号 = 收款卡号 (充值的时候 平台公司名称 + 收款卡号 )
2. `sub_account_name` 主要标记 名称 例如商户名称 A 就可以设置A  或者包工头 B 就可以设置 B 主要为了银联端查看区分
3. `biz_no` 业务单号 每次提交新的支付 (每次交易唯一)  相对与业务单号
4. `bus_code` 银行端业务编号 
5. `bus_no`  银行端编号
6. `bus_name` 银行端编号名称
7. `amount` 金额 (字段 为string 类型 规格 小数位为2 例如  1 需要为1.00 整数后面补0 )


:::

**基础请求参数**

```json
{
  "code": 200,
  "msg": "",
  "data": {}
}
```


## 交易管家
### 获取代发交易模式

 > 获取模式代发,超网代发支付模式使用


POST payroll/mods


> Body 请求参数

```json
{
  "bus_code": "N03020" 
}
```

请求参数

| 字段       | 说明   | 备注        |
|----------|------|-----------|
| bus_code | 交易模式 | 默认 N03020 |


> 返回示例

> 200 Response

```json
{
  "code": 200,
  "msg": "",
  "data": [
    {
      "mod_name":"测试1",
      "mod_no": "000001",
      "eft_at": "20230708",
      "status": "A",
      "exp_at": "20240708"
    }
  ]
}
```
返回参数说明

| 字段       | 说明   | 备注                  |
|----------|------|---------------------|
| mod_no   | 模式编号 | 模式编号(用于超网代发选择)      |
| mod_name | 模式名称 |
| status   | 状态   | 状态 A 有效 B 关闭        |
| eft_at   | 生效时间 |                     |
| exp_at   | 失效时间 | 正常可以忽略,到期网银端会提示手动续约 |

`异常返回`



### 添加单个记账单元

POST  sub_account/add

```json
{
  "sub_account_name": "测试1",
  "sub_account_no": "记账单元编号"
}
```

请求参数

| 字段               | 说明               | 备注  |
|------------------|------------------|-----|
| sub_account_name | 记账单元名称           |     |
| sub_account_no   | 可以自定义 (为空系统自动生成) |     |

> 返回示例

> 200 Response

```json

{
  "sub_account_name": "测试商户",
  "sub_account_no": "记账单元编号"
}

```

返回参数说明 



`异常返回`



### 关闭记账单元

>关闭记账单元 (可以不对接) 主要测试验收使用

`POST` sub_account/close


```json

{
  "sub_account_no": "00001"
}

```

返回参数

| 字段             | 说明     | 备注  |
|----------------|--------|-----|
| sub_account_no | 记账单元编号 |     |



```json

```

返回参数说明



`异常返回`

### 获取单天交易

>该接口用于查询记账子单元的当日交易明细包括交易时间、借贷金额、币种、收付方名称、收付方账号、交易摘要、记账子单元编号、记账子单元名称、记账流水号以及是否关联付方匹配等信息

`POST`  sub_account/trans/today



```json
{
  "sub_account_no": ""
}

```

返回参数

| 字段             | 说明     | 备注      |
|----------------|--------|---------|
| sub_account_no | 记账单元编号 | 为空查询所有  |
| more_key       | 续传key  | 为空无更多内容 |

```json

{
  "trans": [
    {}
  ],
  "more_key": ""
}

```

返回参数说明

| 字段       | 说明    | 备注 |
|----------|-------|--|
| trans    | 交易列表  |  |
| more_key | 续传key | 为空无更多内容 |


`异常返回`

###  获取历史交易

> 该接口用于查询记账子单元的历史交易明细，即已记入记账子单元的历史交易，包括交易日期、交易时间、借贷金额、币种、收付方名称、收付方账号、交易摘要、记账子单元编号、记账子单元名称、记账方式、记账流水号以及是否关联付方匹配等信息。

`POST` sub_account/trans/history



```json

```

返回参数


```json

```

返回参数说明



`异常返回`

### 获取记账单元余额

> 获取对应记账单元余额 (默认批量获取 ,指定对应编号获取对应记账单元)

POST   sub_account/balances

请求参数

| 字段          | 说明     | 备注     |
|-------------|--------|--------|
| sub_account | 记账单元编号 | 空为所有余额 |


返回参数


```json

```

返回参数说明

`异常返回`

## 代发代扣

### 超网代发

> 批量或者单个进行超网代发

`POST` payroll/sup



```json

{
  "total_info": {
    "total_num": "总笔数如果 一个代发1",
    "total_amount": "总金额",
    "biz_no":"业务单号(自己制单生成的业务单号)",
    "sub_account_no": "子单元编号(如果为空 从默认记账单元支付)",
    "mod_no": "交易号码 从查询业务单号获取,(这个可以固定)g"
    
  },
  "detail_info": [
    {
      "name": "姓名",
      "bank_card_no": "卡号",
      "trs_no": "交易序号",
      "amount": "交易金额(eg:1.01,1.00)",
      "remark": "回单备注"
    }
  ]
}
```
返回参数


```json

```

返回参数说明



`异常返回`


### 代发结果查询-按业务参考号

> 根据代发经办请求报文中的业务参考号查询批次的概要信息，同一业务参考号可能对应多笔，但只有一笔是成功的。

`POST` payroll/pay/query



```json

```

返回参数


```json

```

返回参数说明



`异常返回`

### 超网代发 查询交易明细信息

> 批量或者单个进行超网代发

`POST` payroll/pay/detail



```json

```

返回参数


```json

```

返回参数说明



`异常返回`




### 代发退票明细查询

> 用于代发退票明细的查询。

`POST` payroll/refund



```json

```

返回参数


```json

```

返回参数说明



`异常返回`


### 代发明细对账单查询请求

> 用于代发明细对账单查询，采用异步的方式，生成PDF文件，该请求需要开通N66010:代发业务明细对账单业务功能，且当前查询用户需要有该查询账号对应业务的可经办权限。

`POST` payroll/receipt/apply



```json

```

返回参数


```json

```

返回参数说明



`异常返回`


### 代发明细对账单处理结果查询请求

> 用于查询代发明细对账单请求处理结果。 获取回单地址

`POST` payroll/receipt/download



```json

{
  "task_id": "0001"
}

```

返回参数

| 参数      | 说明     | 备注  |
|---------|--------|-----|
| task_id | 申请返回ID |     |


```json

{
  "download_url": ""
}

```

返回参数说明

| 参数           | 说明   | 备注  |     |
|--------------|------|-----|-----|
| download_url | 下载地址 |     |     |



`异常返回`



## 账务处理

### 可经办业务模式查询

> 用于对公支付时业务选择

`POST` p2p/mods



```json

{
  "bus_code": ""
}

```
参数说明

| 字段       | 说明   | 备注     |
|----------|------|--------|
| bus_code | 业务代码 | 为空默认值为 |


返回参数


```json

```

返回参数说明

| 字段       | 说明   | 备注                  |
|----------|------|---------------------|
| mod_no   | 模式编号 | 模式编号(用于超网代发选择)      |
| mod_name | 模式名称 |
| status   | 状态   | 状态 A 有效 B 关闭        |
| eft_at   | 生效时间 |                     |
| exp_at   | 失效时间 | 正常可以忽略,到期网银端会提示手动续约 |


`异常返回`


### 账户详情

> 查询账户的详细信息，包括账户的实时余额，账户性质，开户分行，开户日期，利率，到期日，存期等要素。

`POST` account/info



```json

```

返回参数


```json

{
  "available_balance": "100",
  "account": "账户"
}
```

返回参数说明

| 字段                | 说明   | 备注  |
|-------------------|------|-----|
| available_balance | 可用余额 |     |
| account           | 账户   |     |



 
`异常返回`


### 查询账户历史余额

> 用于查询账户历史某一段时间内的余额。


`POST` account/balance



```json

```

返回参数


```json

```

返回参数说明



`异常返回`


### 单笔回单查询

> 该接口主要用于根据流水号查询单笔交易回单。日期需为回单产生的日期，流水号可通过账户交易信息查询trsQryByBreakPoint请求返回接口transSequenceIdn字段获取。


`POST` pay/receipt/query

```json

```

返回参数


```json

```

返回参数说明


`异常返回`







### 企银支付单笔经办

>  用于发起单笔支付经办 公对公支付


`POST` pay/p2p

```json

```

返回参数


```json
{
  "bank_name": "招商银行",
  "bank_link_no": "21",
  "op_bank_name": "开户行名称",
  "account_name": "账号名称",
  "account": "账户",
  "amount": "10000",
  "remark": "提现退费",
  "mod_no": "",
  "bus_no": "",
  "sub_account_no": "00000001"
}
```
请求参数说明

| 字段             | 说明           | 备注           |
|----------------|--------------|--------------|
| bank_name      | 银行名称         |              |
| bank_link_no   | 银联号          |              |
| op_bank_name   | 开户行          |              |
| account_name   | 账号名称         |              |
| account        | 卡号           |              |
| remark         | 备注(回单上可以显示)  |              |
| bus_no         | 业务流水号        |              |
| mod_no         | 请求支付业务模式进行选择 |              |
| sub_account_no | 记账单元号码       | 如果为空默认从主账户支付 |


返回参数说明


`异常返回`

### 企银支付业务查询

>  用于发起单笔支付经办 获取支付结果查询 通过支付的bus_no 进行查询


`POST` pay/mods

```json
{
  "biz_no": ""
}

```
参数说明

| 字段     | 说明          | 备注  |
|--------|-------------|-----|
| biz_no | 支付请求输入的业务单号 |     |


返回参数


```json

```

返回参数说明


`异常返回`

### 支付退票明细查询

>  用于对银行退票明细的查询。当收方是他行，结算通道为：R 实时-超网时，不支持退票查询。


`POST` pay/refund

```json

```

返回参数


```json

```

返回参数说明


`异常返回`



## 构建与部署 


### 直接下载打包文件模式
1. 下载打包好文件
2. 创建配置文件夹 configs
3. 设置配置文件

```bash
wget '打包文件路径'
go-cmb-service --configs=./configs/cmb.yaml
```

### 自己下载源码构建模式


```bash 
git clone https://github.com/ahKevinXy/go-cmb-service.git

# linux 平台构建方式
 CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o  go-cmb-service
  
# mac 平台构建方式 
 CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o  go-cmb-service

# windows 平台构建方式
 CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o  go-cmb-service
 
 # 创建配置文件
 
 go-cmb-service --configs=./configs/cmb.yaml

```