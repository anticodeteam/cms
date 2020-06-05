$(document).ready(function() {
    $("#addButton").click(function() {
        $('#modalTable').modal({
            show:true,
            backdrop:'static'
        });
    })
});

$(function () {
    $.ajax({
        type:"post",
        url:"/getknowledge",
        success:function(data){
            var tablestr = "";
            var len = data.length;
            var curruser = $('.currUser').val()
            console.log(data);
            //先拼出一个一级目录，然后查询它是否存在子目录，如存在就插入在它后面，如不存在，则继续拼下一个一级目录依次循环
            for (var i = 0; i < len; i++)   {
                if (data[i].Pid == 0 && data[i].Gid == 0){
                    tablestr += "<tr onclick='showdirectory(\"" + data[i].Id + "\")' style='background-color: #f5eac1'>";
                    tablestr += "<td>" + data[i].Title + "</td>";
                    tablestr += "<td>" + data[i].Creater + "</td>";
                    tablestr += "<td>" + data[i].UpdateTime + "</td>";
                    if(data[i].Creater == curruser){
                        if(data[i].Status == 0){
                            tablestr += "<td><a href='#' onclick='changeKnowledgeStatus("+data[i].Id+",1)'>申请上架</a></td>";
                        }else if(data[i].Status == 1){
                            tablestr += "<td><a href='#' onclick='changeKnowledgeStatus("+data[i].Id+",0)'>取消申请</a></td>";
                        }else{
                            tablestr += "<td><a href='#' onclick='changeKnowledgeStatus("+data[i].Id+",2)'>申请下架</a></td>";
                        }
                    }else{
                        tablestr += "<td></td>";
                    }
                    if(data[i].Isguanzhu == 0){
                        tablestr += "<td>" + "<a href='#' id='moji" + data[i].Id + "\'onclick='addGuanzhuInfo("+ data[i].Id + "," + data[i].Pid +")'>关注</a>" + "</td>";
                    }else{
                        tablestr += "<td>" + "<a href='#' id='moji" + data[i].Id + "\'onclick='deleteGuanzhuInfo("+ data[i].Id + "," + data[i].Pid +")'>取消关注</a>" + "</td>";
                    }
                    tablestr += "<td><img src='/static/img/del.jpg' onclick='deleteKnow("+data[i].Id+")' width='20px' height='20px'></td>"
                    tablestr += "<td><button class='btn btn-default' onclick='openLevel2AddPage("+ data[i].Id +")'>添加二级目录</button></td>"
                    tablestr += "</tr>"
                    for(var j = 0; j < len; j++){
                        if (data[i].Id == data[j].Pid){
                            tablestr += "<tr onclick='jump("+ data[j].Id +")' class='tr"+ data[i].Id +"' style='background-color: #f4f6f9'>";
                            tablestr += "<td>" + data[j].Title      + "</td>";
                            tablestr += "<td>" + data[j].Creater    + "</td>";
                            tablestr += "<td>" + data[j].UpdateTime + "</td>";
                            if(data[j].Creater == curruser){
                                if(data[j].Status == 0){
                                    tablestr += "<td><a href='#' onclick='changeKnowledgeStatus("+data[j].Id+",1)'>申请上架</a></td>";
                                }else if(data[j].Status == 1){
                                    tablestr += "<td><a href='#' onclick='changeKnowledgeStatus("+data[j].Id+",0)'>取消申请</a></td>";
                                }else{
                                    tablestr += "<td><a href='#' onclick='changeKnowledgeStatus("+data[j].Id+",2)'>申请下架</a></td>";
                                }
                            }else{
                                tablestr += "<td></td>";
                            }
                            if(data[j].Isguanzhu == 0){
                                tablestr += "<td>" + "<a href='#' id='moji" + data[j].Id + "\'onclick='addGuanzhuInfo("+ data[j].Id + "," + data[j].Pid +")'>关注</a>" + "</td>";
                            }else{
                                tablestr += "<td>" + "<a href='#' id='moji" + data[j].Id + "\'onclick='deleteGuanzhuInfo("+ data[j].Id + "," + data[j].Pid +")'>取消关注</a>" + "</td>";
                            }
                            tablestr += "<td><img src='/static/img/color_delete.jpg' onclick='deleteKnow("+data[j].Id+")' width='20px' height='20px'></td>"
                            tablestr += "<td></td>"
                            tablestr += "</tr>";
                        }
                    }
                    tablestr += "</td></tr>"
                }
            }
            $("#bodys").html(tablestr);
        }
    });
})

function showdirectory(id){
    if($(".tr"+id).is(":hidden")){
        $(".tr"+id).show();
    }else if(!$("#tr"+id).is(":hidden")){
        $(".tr"+id).hide();
    }
}

function jump(gid) {
    window.location.href="/jump?id="+ gid
}

function  addGuanzhuInfo(id,pid) {
    $.ajax({
        type:"post",
        url:"/addGuanzhu",
        data: {Id:id,Pid:pid},
        dataType : "json",
        success:function(data){
            if (data.flag == true){
                console.log("id:"+id);
                // $("#moji"+id).text('取消关注');
                window.location.reload();
            }
        }
    });
    event.stopPropagation();
}
function deleteGuanzhuInfo(id){
    if(confirm("确定要取消关注吗？"))
    {
        $.ajax({
            type:"post",
            url:"/deleteGuanzhu",
            data: {Id:id},
            success:function(){
                alert("取消成功！");
                // $("#moji"+id).text('关注');
                window.location.reload();
            }
        });
        event.stopPropagation();
        return true;
    }
    else
    {
        return false;
    }
}


function userSaveKnowledge() {
    var konwledgeName = $("#knowledge").val();
    if(konwledgeName){
        $.ajax({
            type:"post",
            url:"/userSaveKnowledge",
            data: {Name:konwledgeName},
            success:function (result) {
                console.log("result:"+result)
                alert("添加成功！");
                window.location.reload();
            }
        })
    }
}

function changeKnowledgeStatus(id,status) {
    $.ajax({
        type:"post",
        url:"/changeThisStatus",
        data:{Id:id,status:status},
        success:function (result) {
            alert("申请已提交，请耐心等待管理员处理！")
            window.location.reload()
        },
        error:function () {
            alert("申请提交失败")
        }
    })
    event.stopPropagation();
}

function openLevel2AddPage(id) {
    $('.knowId').val(id);
    $('#modalTable2').modal({
        show:true,
        backdrop:'static'
    });
    event.stopPropagation();
}

function addLevel2Menu() {
    var id = $('.knowId').val();
    var title = $('#knowlevel2').val()
    $.ajax({
        type:"post",
        url:"/addLevel2Menu",
        data:{ID:id,Title:title},
        dataType: "json",
        success:function () {
            alert("申请已提交，请耐心等待管理员处理！")
            window.location.reload();
        },
        error:function () {
            alert("申请提交失败")
        }
    })
}

function deleteKnow(id){
    $.ajax({
        type:"post",
        url:"/deleteKnow",
        data:{"ID":id},
        dataType:"json",
        success:function () {
            alert("删除成功!")
            window.location.reload();
        },
        error:function () {
            alert("删除失败!")
        }
    })
    event.stopPropagation();
}