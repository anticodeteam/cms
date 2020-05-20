<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <title>AdminLTE 3 | Dashboard</title>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <!-- Font Awesome -->
    <link rel="stylesheet" href="/static/plugins/fontawesome-free/css/all.min.css">
    <!-- Ionicons -->
    <link rel="stylesheet" href="https://code.ionicframework.com/ionicons/2.0.1/css/ionicons.min.css">
    <!-- Theme style -->
    <link rel="stylesheet" href="/static/css/adminlte.min.css">
    <!-- Google Font: Source Sans Pro -->
    <link href="https://fonts.googleapis.com/css?family=Source+Sans+Pro:300,400,400i,700" rel="stylesheet">
</head>
<body class="hold-transition sidebar-mini layout-fixed">
<div class="wrapper">


    <!-- Content Wrapper. Contains page content -->
    <div class="content-wrapper">
        <!-- Content Header (Page header) -->
        <div class="content-header">
            <div class="container-fluid">
                <div class="row mb-2">
                    <div class="col-sm-6">
{{/*                        <h1 class="m-0 text-dark">{{.BigTitle}}</h1>*/}}
                    </div><!-- /.col -->
                </div><!-- /.row -->
            </div><!-- /.container-fluid -->
        </div>
        <!-- /.content-header -->

        <!-- Main content -->

        <div class="row">
            <div class="col-12">
                <div class="card">
                    <div class="card-header">
                        <button class="btn btn-default" id="addButton">添加新的知识库选项</button>
                        <h3 class="card-title">{{.BigTitle}}</h3>
                        <div class="card-tools">
                            <div class="input-group input-group-sm" style="width: 150px;">
                                <input type="text" name="table_search" class="form-control float-right" placeholder="Search">

                                <div class="input-group-append">
                                    <button type="submit" class="btn btn-default"><i class="fas fa-search"></i></button>
                                </div>
                            </div>
                        </div>
                    </div>
                    <!-- /.card-header -->
                    <div class="card-body table-responsive p-0" style="height: 100%;">
                        <table class="table table-head-fixed text-nowrap">
                            <thead>
                            <tr>
                                <th>名称</th>
                                <th>创建者</th>
                                <th>更新时间</th>
                                <th></th>
                            </tr>
                            </thead>
                            <tbody  id="bodys">
                            </tbody>
                        </table>
                    </div>
                    <!-- /.card-body -->
                </div>

                <!-- /.card -->
            </div>
        </div>
    </div>
    <!-- /.control-sidebar -->

    <div class="container form-horizontal">
        <div class="modal fade" id="modalTable">
            <div class="modal-dialog">
                <div class="modal-content">
                    <div class="modal-header">
                        <h3 class="modal-title">添加</h3>
                        <button type="button" class="close" data-dismiss="modal">
                            <span aria-hidden="true">&times;</span>
                        </button>
                    </div>
                    <div class="modal-body">
                        <div class="modal-body-content">
                            <div class="form-group must">
                                <label class="col-sm-3 control-label">名称</label>
                                <div class="col-sm-7">
                                    <input type="text" class="form-control" id="knowledge" name="knowledge"></div>
                            </div>
                        </div>
                    </div>
                    <div class="modal-footer">
                        <button type="button" class="btn btn-default" data-dismiss="modal">取消</button>
                        <button type="button" class="btn btn-primary" onclick="userSaveKnowledge()">保存</button>
                    </div>
                </div>
            </div>
        </div>
    </div>

</div>
<!-- ./wrapper -->

<!-- jQuery -->
<script src="/static/plugins/jquery/jquery.min.js"></script>
<!-- Bootstrap 4 -->
<script src="/static/plugins/bootstrap/js/bootstrap.bundle.min.js"></script>
<!-- AdminLTE App -->
<script src="/static/js/adminlte.min.js"></script>
<!-- AdminLTE for demo purposes -->
<script src="/static/js/demo.js"></script>

<script type="text/javascript">
    $(document).ready(function() {
        $("#addButton").click(function() {
            $('#modalTable').modal({
                show:true,
                backdrop:'static'
            });
        })
    });

    window.onload = function () {
        $.ajax({
            type:"post",
            url:"/getknowledge",
            success:function(data){
                var tablestr = "";
                var len = data.length;
                console.log(data);
                console.log("长度："+len);
                //先拼出一个一级目录，然后查询它是否存在子目录，如存在就插入在它后面，如不存在，则继续拼下一个一级目录依次循环
                for (var i = 0; i < len; i++)   {
                    if (data[i].Pid == 0 && data[i].Gid == 0){
                        tablestr += "<tr onclick='showdirectory(\"" + data[i].Id + "\")'>";
                        tablestr += "<td>" + data[i].Title + "</td>";
                        tablestr += "<td>" + data[i].Creater + "</td>";
                        tablestr += "<td>" + data[i].UpdateTime + "</td>";
                        if(data[i].Isguanzhu == 0){
                            tablestr += "<td>" + "<a href='#' id='moji" + data[i].Id + "\'    onclick='addGuanzhuInfo("+ data[i].Id + "," + data[i].Pid +")'>关注</a>" + "</td>";
                        }else{
                            tablestr += "<td>" + "<a href='#' id='moji" + data[i].Id + "\'    onclick='deleteGuanzhuInfo("+ data[i].Id + "," + data[i].Pid +")'>取消关注</a>" + "</td>";

                        }

                        tablestr += "</tr><tr><td colspan='4'><table style='width: 100%' id='div"+ data[i].Id +"'>";
                        for(var j = 0; j < len; j++){
                            if (data[i].Id == data[j].Pid){
                                tablestr += "<tr>";
                                tablestr += "<td onclick='jump("+ data[j].Id +")'>" + data[j].Title + "</td>";
                                tablestr += "<td onclick='jump("+ data[j].Id +")'>" + data[j].Creater + "</td>";
                                tablestr += "<td onclick='jump("+ data[j].Id +")'>" + data[j].UpdateTime + "</td>";
                                //tablestr += "<td>" + "<a href='#' id='moji" + data[j].Id + "\'    onclick='addGuanzhuInfo("+ data[j].Id + "," + data[j].Pid +")'>关注</a>" + "</td>";
                                tablestr += "</tr>";
                            }
                        }
                        tablestr += "</table></td></tr>"
                    }
                }
                $("#bodys").html(tablestr)
            }
        });
    }
</script>


<script type="text/javascript">
    function showdirectory(id){
        if($("#div"+id).is(":hidden")){
            $("#div"+id).show();

        }else if(!$("#div"+id).is(":hidden")){
            $("#div"+id).hide();
        }
    }

    function jump(gid) {
        // $.ajax({
        //     type:"post",
        //     url:"/jump",
        //     data:{id:gid},
        //     success:function () {
        //         window.location.href="/jump"
        //     }
        // });
       // $.post("/jump", {id:gid} );
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
                    alert("添加成功，等待管理员审批！");
                    window.location.reload();
                }
            })
        }
    }
</script>
</body>
</html>