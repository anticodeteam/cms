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
                        <h3 class="card-title"></h3>
                        <div class="card-tools">
                            <div class="input-group input-group-sm" style="width: 150px;">
                                <input type="text" name="table_search" class="form-control float-right" placeholder="Search">
                                <div class="input-group-append">
                                    <button type="submit" class="btn btn-default"><i class="fas fa-search"></i></button>
                                </div>
                            </div>
                        </div>
                    </div>
                    <!-- /.card-body -->
                </div>
                <!-- /.card -->
            </div>
        </div>
        <button type="button" class="btn btn-block btn-outline-primary" style="width: 10%" onclick="add()">添加知识文章</button>
        <p></p>
        <ul  id="article">
            {{range .List}}
            <li><a href="#" onclick="jump({{.Id}})">{{.Title}}</a></li>
            {{end}}
        </ul>

        <input type="hidden" id="fid" value="{{.Fid}}">
    </div>
    <!-- /.control-sidebar -->
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
    // window.onload = function () {
    //     $.ajax({
    //         type:"post",
    //         url:"/getArticleList",
    //         success:function(data){
    //             console.log("data:"+data)
    //             var str = ""
    //             for (var i = 0;i<data.length;i++){
    //                 str += "<li><a href='#' onclick='add("+ data[i] +")'>"
    //             }
    //         }
    //     });
    // }
</script>


<script type="text/javascript">
    function add() {
        var id = $("#fid").val()
        window.location.href="/jumpAddArticle?id=" +id;
    }

    function jump(id) {
        window.location.href="/articlePage?id="+ id;
    }
</script>

</body>
</html>