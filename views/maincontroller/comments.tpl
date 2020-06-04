<!DOCTYPE html>
<html>
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
    <title></title>
    <meta charset="utf-8" />
    <style type="text/css">
        #tab {
            border-right:1px solid #000000;
            border-bottom:1px solid #000000;
            width:200px;
        }
        #tab td{
            border-left:1px solid #000000;
            border-top:1px solid #000000;
        }
        #comments{
            margin-left: 280px;
        }
        body {
            height: 100%;
            width: 100%;
            display: flex;
            align-items: center;
            justify-content: center;
        }

        .comBtn {
            position: relative;
            outline: none;
            text-decoration: none;
            border-radius: 50px;
            display: flex;
            justify-content: center;
            align-items: center;
            cursor: pointer;
            text-transform: uppercase;
            height: 40px;
            width: 100px;
            opacity: 1;
            background-color: #ffffff;
            border: 1px solid rgba(22, 76, 167, 0.6);
        }
        .comBtn span {
            color: #164ca7;
            font-size: 12px;
            font-weight: 500;
            letter-spacing: 0.7px;
        }
        .comBtn:hover {
            animation: rotate 0.7s ease-in-out both;
        }
        .comBtn:hover span {
            animation: storm 0.7s ease-in-out both;
            animation-delay: 0.06s;
        }

        @keyframes rotate {
            0% {
                transform: rotate(0deg) translate3d(0, 0, 0);
            }
            25% {
                transform: rotate(3deg) translate3d(0, 0, 0);
            }
            50% {
                transform: rotate(-3deg) translate3d(0, 0, 0);
            }
            75% {
                transform: rotate(1deg) translate3d(0, 0, 0);
            }
            100% {
                transform: rotate(0deg) translate3d(0, 0, 0);
            }
        }
        @keyframes storm {
            0% {
                transform: translate3d(0, 0, 0) translateZ(0);
            }
            25% {
                transform: translate3d(4px, 0, 0) translateZ(0);
            }
            50% {
                transform: translate3d(-3px, 0, 0) translateZ(0);
            }
            75% {
                transform: translate3d(2px, 0, 0) translateZ(0);
            }
            100% {
                transform: translate3d(0, 0, 0) translateZ(0);
            }
        }

    </style>

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
<script>
    $(function () {
        $("#Button1").click(function () {
            //获取昵称
            var name = $("#name").val();
            //获取内容
            var next = $("#next").val();
            //添加
            var ul_li = $(" <tr><td>" + name + "</td><td>" + next + "</td><td>"+"<button>删除</button>"+"</td></tr>");
            var b = $("#tab");
            b.append(ul_li);
        });
    });
    function Addcomment() {
        //获取昵称
        var name = $("#name").val();
        //获取内容
        var next = $("#next").val();
        var id = $("#Knoid").val()
        console.log(name+next+id)
        window.location.href="/addcomments?name="+name+"&comment=s"+next+"&id="+id;
    }
    function DeleteComments() {
        var name = $("#name").val();
        console.log(name)
        window.location.href="/DeleteComment?name="+name;
    }
    function GetComment() {
        var id = $("#Knoid").val()
        //window.location.href="/getcomments?id="+id;
        $.ajax({
            type:"post",
            url:"/getcomments?id="+id,
            success:function (data) {

                $(".com").slideToggle("slow")
                var str = "";
                for (var i=0;i<data.length;i++){
                    str += "<ul style=\"list-style:none;\"><li style=\'width: 180px;\'>"+"用户："+ data[i]["Name"] +
                        "<button  style=\"margin-left: 15px;border-radius: 5px;background-color: #1fc8e3\" " +
                        "onclick=\"DeleteComments()\" "+">删除</button>" + "</li><li>"+"内容："+data[i]["Comments"]+"</li></ul>"
                }
                $("#tabCom").html(str);

            }
        })
    }
</script>
</html>
