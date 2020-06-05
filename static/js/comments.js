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