<!DOCTYPE html>
<html>
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
    <title></title>
    <meta charset="utf-8" />
    <link rel="stylesheet" href="/static/css/comments.css">
</head>
<body>
<div id="comments">
    <button onclick="GetComment()" class="comBtn"><span>他人评论</span></button>
    <div class="com">
        <fieldset>
            <legend>评论区</legend>
                <table id="tabCom"></table>
        </fieldset>
    </div>
    <fieldset>
        <legend><span>个人评论</span></legend>
        <table  id="tab">
            <tr>
                <td class="tdstyle">昵称</td>
                <td class="tdstyle">操作</td>
                <td class="commnet">内容</td>
            </tr>
        </table>
        昵称：<input id="name" type="text" />
        <br>
        内容: <br><textarea id="next" cols="100" rows="5"></textarea>
        <br>
        <input id="Button1" type="button" onclick="Addcomment()" value="评论" />
    </fieldset>
</div>
<div>
    <input type="hidden" id="Knoid" value="{{.konwledgeId}}">
</div>
<pre style="width:100%;height:100%" id="data-test"></pre>
</body>
<script src="https://code.jquery.com/jquery-3.4.0.min.js"></script>
<script src="/static/js/comments.js"></script>
</html>
