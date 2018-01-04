'# async-httpclient

## 练习要求：
> * 依据文档图6-1，用中文描述 Reactive 动机
![Figure 6.1. Travel Agency Orchestration Service](images/Figure6.1.png)

> * 使用 go HTTPClient 实现图 6-2 的 Naive Approach

![Figure 6.2. Time consumed to create a response for the client – synchronous way](images/Figure6.2.png)


> * 为每个 HTTP 请求设计一个 goroutine ，利用 Channel 搭建基于消息的异步机制，实现图 6-3

![Figure 6.3. Time consumed to create a response for the client – asynchronous way](images/Figure6.3.png)


> * 对比两种实现，用数据说明 go 异步 REST 服务协作的优势

![compare](images/compare.png)
