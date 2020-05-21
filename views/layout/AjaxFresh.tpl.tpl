<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
</head>
<script src="/static/plugins/jquery/jquery.min.js"></script>
{{.HtmlHead}}
<body>
<div class="userWrap">
     <ul class="userMenu">
             <li class="current" data-id="center">用户中心</li>
             <li data-id="account">账户信息</li>
             <li data-id="trade">交易记录</li>
             <li data-id="info">消息中心</li>
     </ul>
     <div id="content"></div>
</div>

<div class="container">
    {{.LayoutContent}}
</div>
<div>
</div>
<script src="http://netdna.bootstrapcdn.com/bootstrap/3.0.3/js/bootstrap.min.js"></script>
{{.Scripts}}


</body>
<script>

</script>
</html>