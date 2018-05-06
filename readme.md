## 基于微信公众号的支付接口

### 实现的接口

> - 统一下单
> - 查询订单
> - 生成前端支付的参数

### 使用步骤
> ```go
> go get github.com/Chasiny/wechat/wepay
> import "github.com/Chasiny/wechat/wepay"   
> ```
>
> ```go
>  //配置公众号
>  	wp := &WePay.WePay{}
>  	wp.Init("填写公众号APPID","公众号对应的商户号id","商户API密钥","支付通知地址（后端接口）")
> ```
>
> ```go
>  //生成订单号
>  	outTradeNO:=WePay.CreateOutTradeNO()
>
>  	//获取前端支付需要的参数，接口参数二为支付金额，单位分
>  	res, err := wp.WebRequest(outTradeNO, 1, "127.0.0.1", 10, "用户openid")
>  	if (err != nil) {
>  		fmt.Println(err.Error())
>  		return
>  	}
>  	fmt.Println(res.Package)
> ```
>
> ```go
>  //查询订单
>  	ans,err:=wp.Query(outTradeNO)
>  	if (err != nil) {
>  		fmt.Println(err.Error())
>  		return
>  	}
> ```
>

### angular2前端调用支付

> ```typescript
> //WeixinJSBridge属于微信浏览器内置对象，在typescript中declare就可以避免编译出错
> declare var WeixinJSBridge:any;
> declare var document:any;
>
> WeixinJSBridge.invoke(
>     'getBrandWCPayRequest', {
>         "appId": p.appId,
>         "timeStamp": p.timeStamp  ,
>         "nonceStr": p.nonceStr,
>         "package": p.package,
>         "signType": p.signType,
>         "paySign": p.paySign,
>       },
>     function (res) {
>       if (res.err_msg == "get_brand_wcpay_request:ok") {
>         alert("支付成功");
>       } 
>       else {
>         alert("支付失败" + JSON.stringify(res));
>       }
>     }
>   );
> ```
>
> 

### by chasiny

