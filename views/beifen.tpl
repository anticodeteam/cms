<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <title>AdminLTE 3 | Dashboard</title>
    <!-- Tell the browser to be responsive to screen width -->
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <!-- Font Awesome -->
    <link rel="stylesheet" href="/static/plugins/fontawesome-free/css/all.min.css">
    <!-- Ionicons -->
    <link rel="stylesheet" href="https://code.ionicframework.com/ionicons/2.0.1/css/ionicons.min.css">
    <!-- DataTables -->
    <link rel="stylesheet" href="/static/plugins/datatables-bs4/css/dataTables.bootstrap4.min.css">
    <link rel="stylesheet" href="/static/plugins/datatables-responsive/css/responsive.bootstrap4.min.css">
    <!-- Tempusdominus Bbootstrap 4 -->
    <link rel="stylesheet" href="/static/plugins/tempusdominus-bootstrap-4/css/tempusdominus-bootstrap-4.min.css">
    <!-- iCheck -->
    <link rel="stylesheet" href="/static/plugins/icheck-bootstrap/icheck-bootstrap.min.css">
    <!-- JQVMap -->
    <link rel="stylesheet" href="/static/plugins/jqvmap/jqvmap.min.css">
    <!-- Theme style -->
    <link rel="stylesheet" href="/static/css/adminlte.min.css">
    <!-- overlayScrollbars -->
    <link rel="stylesheet" href="/static/plugins/overlayScrollbars/css/OverlayScrollbars.min.css">
    <!-- Daterange picker -->
    <link rel="stylesheet" href="/static/plugins/daterangepicker/daterangepicker.css">
    <!-- summernote -->
    <link rel="stylesheet" href="/static/plugins/summernote/summernote-bs4.css">
    <!-- Google Font: Source Sans Pro -->
    <link href="https://fonts.googleapis.com/css?family=Source+Sans+Pro:300,400,400i,700" rel="stylesheet">
</head>
<body class="hold-transition sidebar-mini layout-fixed">
<div class="wrapper">

    <!-- Navbar -->
    <nav class="main-header navbar navbar-expand navbar-white navbar-light">
        <!-- Left navbar links -->
        <ul class="navbar-nav">
            <li class="nav-item">
                <a class="nav-link" data-widget="pushmenu" href="#" role="button"><i class="fas fa-bars"></i></a>
            </li>
            <li class="nav-item d-none d-sm-inline-block">
                <a href="/homeAction" class="nav-link">主页</a>
            </li>
            <li class="nav-item d-none d-sm-inline-block">
                <a href="#" class="nav-link">Contact</a>
            </li>
        </ul>

        <!-- SEARCH FORM -->
        <form class="form-inline ml-3">
            <div class="input-group input-group-sm">
                <input class="form-control form-control-navbar" type="search" placeholder="Search" aria-label="Search">
                <div class="input-group-append">
                    <button class="btn btn-navbar" type="submit">
                        <i class="fas fa-search"></i>
                    </button>
                </div>
            </div>
        </form>

        <!-- Right navbar links -->
        <ul class="navbar-nav ml-auto">
            <!-- Messages Dropdown Menu -->
            <li class="nav-item dropdown">
                <a class="nav-link" data-toggle="dropdown" href="#">
                    <i class="far fa-comments"></i>
                    <span class="badge badge-danger navbar-badge">3</span>
                </a>
                <div class="dropdown-menu dropdown-menu-lg dropdown-menu-right">
                    <a href="#" class="dropdown-item">
                        <!-- Message Start -->
                        <div class="media">
                            <img src="/static/img/user1-128x128.jpg" alt="User Avatar" class="img-size-50 mr-3 img-circle">
                            <div class="media-body">
                                <h3 class="dropdown-item-title">
                                    Brad Diesel
                                    <span class="float-right text-sm text-danger"><i class="fas fa-star"></i></span>
                                </h3>
                                <p class="text-sm">Call me whenever you can...</p>
                                <p class="text-sm text-muted"><i class="far fa-clock mr-1"></i> 4 Hours Ago</p>
                            </div>
                        </div>
                        <!-- Message End -->
                    </a>
                    <div class="dropdown-divider"></div>
                    <a href="#" class="dropdown-item">
                        <!-- Message Start -->
                        <div class="media">
                            <img src="/static/img/user8-128x128.jpg" alt="User Avatar" class="img-size-50 img-circle mr-3">
                            <div class="media-body">
                                <h3 class="dropdown-item-title">
                                    John Pierce
                                    <span class="float-right text-sm text-muted"><i class="fas fa-star"></i></span>
                                </h3>
                                <p class="text-sm">I got your message bro</p>
                                <p class="text-sm text-muted"><i class="far fa-clock mr-1"></i> 4 Hours Ago</p>
                            </div>
                        </div>
                        <!-- Message End -->
                    </a>
                    <div class="dropdown-divider"></div>
                    <a href="#" class="dropdown-item">
                        <!-- Message Start -->
                        <div class="media">
                            <img src="/static/img/user3-128x128.jpg" alt="User Avatar" class="img-size-50 img-circle mr-3">
                            <div class="media-body">
                                <h3 class="dropdown-item-title">
                                    Nora Silvester
                                    <span class="float-right text-sm text-warning"><i class="fas fa-star"></i></span>
                                </h3>
                                <p class="text-sm">The subject goes here</p>
                                <p class="text-sm text-muted"><i class="far fa-clock mr-1"></i> 4 Hours Ago</p>
                            </div>
                        </div>
                        <!-- Message End -->
                    </a>
                    <div class="dropdown-divider"></div>
                    <a href="#" class="dropdown-item dropdown-footer">See All Messages</a>
                </div>
            </li>
            <!-- Notifications Dropdown Menu -->
            <li class="nav-item dropdown">
                <a class="nav-link" data-toggle="dropdown" href="#">
                    <i class="far fa-bell"></i>
                    <span class="badge badge-warning navbar-badge">15</span>
                </a>
                <div class="dropdown-menu dropdown-menu-lg dropdown-menu-right">
                    <span class="dropdown-item dropdown-header">15 Notifications</span>
                    <div class="dropdown-divider"></div>
                    <a href="#" class="dropdown-item">
                        <i class="fas fa-envelope mr-2"></i> 4 new messages
                        <span class="float-right text-muted text-sm">3 mins</span>
                    </a>
                    <div class="dropdown-divider"></div>
                    <a href="#" class="dropdown-item">
                        <i class="fas fa-users mr-2"></i> 8 friend requests
                        <span class="float-right text-muted text-sm">12 hours</span>
                    </a>
                    <div class="dropdown-divider"></div>
                    <a href="#" class="dropdown-item">
                        <i class="fas fa-file mr-2"></i> 3 new reports
                        <span class="float-right text-muted text-sm">2 days</span>
                    </a>
                    <div class="dropdown-divider"></div>
                    <a href="#" class="dropdown-item dropdown-footer">See All Notifications</a>
                </div>
            </li>
            <li class="nav-item">
                <a class="nav-link" data-widget="control-sidebar" data-slide="true" href="#" role="button">
                    <i class="fas fa-th-large"></i>
                </a>
            </li>
        </ul>
    </nav>
    <!-- /.navbar -->

    <!-- Main Sidebar Container -->
    <aside class="main-sidebar sidebar-dark-primary elevation-4">
        <!-- Brand Logo -->
        <a href="" class="brand-link">
            <img src="/static/img/logo2.png" alt="AdminLTE Logo" class="brand-image img-circle elevation-3"
                 style="opacity: .8">
            <span class="brand-text font-weight-light">CMS</span>
        </a>

        <!-- Sidebar -->
        <div class="sidebar">
            <!-- Sidebar user panel (optional) -->


            <!-- Sidebar Menu -->
            <nav class="mt-2">
                <ul class="nav nav-pills nav-sidebar flex-column" data-widget="treeview" role="menu" data-accordion="false">
                    <!-- Add icons to the links using the .nav-icon class
                         with font-awesome or any other icon font library -->
                    <li class="nav-item">
                        <a href="#" class="nav-link">
                            <i class="nav-icon fas fa-th"></i>
                            <p>
                                欢迎
                            </p>
                        </a>
                    </li>
                    <li class="nav-item has-treeview">
                        <a href="#" class="nav-link">
                            <i class="nav-icon fas fa-copy"></i>
                            <p>
                                系统设置
                                <i class="right fas fa-angle-left"></i>
                            </p>
                        </a>
                        <ul class="nav nav-treeview">
                            <li class="nav-item">
                                <a href="/treeview" class="nav-link">
                                    <i class="far fa-circle nav-icon"></i>
                                    <p>客户端目录</p>
                                </a>
                            </li>
                        </ul>
                    </li>

                </ul>
            </nav>
            <!-- /.sidebar-menu -->
        </div>
        <!-- /.sidebar -->
    </aside>

    <!-- Content Wrapper. Contains page content -->
    <div class="content-wrapper">
        <!-- Content Header (Page header) -->
        <div class="content-header">
            <div class="container-fluid">
                <div class="row mb-2">
                    <div class="col-sm-6">
                        <h1 class="m-0 text-dark">系统设置</h1>
                    </div><!-- /.col -->
                    <div class="col-sm-6">
                        <ol class="breadcrumb float-sm-right">
                            <li class="breadcrumb-item"><a href="/homeAction">主页</a></li>
                            <li class="breadcrumb-item active">目录设置</li>
                        </ol>
                    </div><!-- /.col -->
                </div><!-- /.row -->
            </div><!-- /.container-fluid -->
        </div>
        <!-- /.content-header -->

        <!-- Main content -->

        <!-- 主目录 -->
        <section class="content">
            <div class="row">
                <div class="col-12">
                    <div class="card">
                        <div class="card-header">
                            <h3 class="card-title">大目录设置</h3>
                        </div>
                        <!-- /.card-header -->
                        <div class="card-body">
                            <div class="btn-group">
                                <button type="button" class="btn btn-default" id="addButton"  data-name="addButton">添加</button>
                                <button type="button" class="btn btn-default">编辑</button>
                                <button type="button" class="btn btn-default">删除</button>
                            </div>
                            <a href="/gettree">跳转</a>
                            <p></p>
                            <table id="table1" class="table table-bordered table-hover">
                                <thead>
                                <tr>
                                    <th>#</th>
                                    <th>id</th>
                                    <th>目录名</th>
                                </tr>
                                </thead>
                                <tbody>
                                {{range .List}}
                                    <tr onclick="isShow({{.Id}},'{{.Title}}')">
                                        <td><input type="checkbox"></td>
                                        <th >{{.Id}}</th>
                                        <th>{{.Title}}</th>
                                    </tr>
                                {{end}}
                            </table>
                        </div>
                        <!-- /.card-body -->
                    </div>
                </div>
                <!-- /.col -->
            </div>
            <!-- /.row -->
        </section>

        <input type="hidden" id="parentsId" >
        <input type="hidden" id="parentsTitle" >

        <!-- 大分类目录  -->
        <section class="content" id="example1" style="display: none">
            <div class="row">
                <div class="col-12">
                    <div class="card">
                        <div class="card-header">
                            <h3 class="card-title"  id="directory"></h3>
                        </div>
                        <!-- /.card-header -->
                        <div class="card-body">
                            <div class="btn-group">
                                <button type="button" class="btn btn-default" id="addButton2"  data-name="addButton2">添加</button>
                                <button type="button" class="btn btn-default">编辑</button>
                                <button type="button" class="btn btn-default">删除</button>
                            </div>
                            <p></p>
                            <table id="table2" class="table table-bordered table-hover" >
                                <thead>
                                <tr>
                                    <th>#</th>
                                    <th>目录名</th>
                                </tr>
                                </thead>
                                <tbody id="bodys">
                            </table>
                        </div>
                        <!-- /.card-body -->
                    </div>
                </div>
                <!-- /.col -->
            </div>
            <!-- /.row -->
        </section>

        <!--  小分类目录 -->
        <section class="content" id="example3" style="display: none">
            <div class="row">
                <div class="col-12">
                    <div class="card">
                        <div class="card-header">
                            <h3 class="card-title"  id="directory2"></h3>
                        </div>
                        <!-- /.card-header -->
                        <div class="card-body">
                            <div class="btn-group">
                                <button type="button" class="btn btn-default" id="addButton3"  data-name="addButton3">添加</button>
                                <button type="button" class="btn btn-default">编辑</button>
                                <button type="button" class="btn btn-default">删除</button>
                            </div>
                            <p></p>
                            <table id="table3" class="table table-bordered table-hover" >
                                <thead>
                                <tr>
                                    <th>#</th>
                                    <th>目录名</th>
                                </tr>
                                </thead>
                                <tbody id="secondbodys">
                            </table>
                        </div>
                        <!-- /.card-body -->
                    </div>
                </div>
                <!-- /.col -->
            </div>
            <!-- /.row -->
        </section>


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
                                        <input type="text" class="form-control" id="projcatename" name="projcatename"></div>
                                </div>
                            </div>
                        </div>
                        <div class="modal-footer">
                            <button type="button" class="btn btn-default" data-dismiss="modal">取消</button>
                            <button type="button" class="btn btn-primary" onclick="save()">保存</button>
                        </div>
                    </div>
                </div>
            </div>
        </div>

        <div class="container form-horizontal">
            <div class="modal fade" id="modalTable2">
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
                                        <input type="text" class="form-control" id="childname" name="childname"></div>
                                </div>
                            </div>
                        </div>
                        <div class="modal-footer">
                            <button type="button" class="btn btn-default" data-dismiss="modal">取消</button>
                            <button type="button" class="btn btn-primary" onclick="save2()">保存</button>
                        </div>
                    </div>
                </div>
            </div>
        </div>


        <div class="container form-horizontal">
            <div class="modal fade" id="modalTable3">
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
                                        <input type="text" class="form-control" id="smallname" name="smallname"></div>
                                </div>
                            </div>
                        </div>
                        <div class="modal-footer">
                            <button type="button" class="btn btn-default" data-dismiss="modal">取消</button>
                            <button type="button" class="btn btn-primary" onclick="save3()">保存</button>
                        </div>
                    </div>
                </div>
            </div>
        </div>



    </div>
    <!-- /.content-wrapper -->
    <footer class="main-footer">
        <strong>Copyright &copy; 2008-2020 <a href="https://www.tci-sz.com/">transcosmos.sz</a>.</strong>
        All rights reserved.
        <div class="float-right d-none d-sm-inline-block">
            <b>Version</b> 1.0.0
        </div>
    </footer>

    <!-- Control Sidebar -->
    <aside class="control-sidebar control-sidebar-dark">
        <!-- Control sidebar content goes here -->
    </aside>
    <!-- /.control-sidebar -->
</div>
<!-- ./wrapper -->

<!-- jQuery -->
<script src="/static/plugins/jquery/jquery.min.js"></script>
<!-- jQuery UI 1.11.4 -->
<script src="/static/plugins/jquery-ui/jquery-ui.min.js"></script>
<!-- Resolve conflict in jQuery UI tooltip with Bootstrap tooltip -->
<script>
    $.widget.bridge('uibutton', $.ui.button)
</script>
<!-- Bootstrap 4 -->
<script src="/static/plugins/bootstrap/js/bootstrap.bundle.min.js"></script>
<!-- DataTables -->
<script src="/static/plugins/datatables/jquery.dataTables.min.js"></script>
<script src="/static/plugins/datatables-bs4/js/dataTables.bootstrap4.min.js"></script>
<script src="/static/plugins/datatables-responsive/js/dataTables.responsive.min.js"></script>
<script src="/static/plugins/datatables-responsive/js/responsive.bootstrap4.min.js"></script>
<!-- ChartJS -->
<script src="/static/plugins/chart.js/Chart.min.js"></script>
<!-- Sparkline -->
<script src="/static/plugins/sparklines/sparkline.js"></script>
<!-- JQVMap -->
<script src="/static/plugins/jqvmap/jquery.vmap.min.js"></script>
<script src="/static/plugins/jqvmap/maps/jquery.vmap.usa.js"></script>
<!-- jQuery Knob Chart -->
<script src="/static/plugins/jquery-knob/jquery.knob.min.js"></script>
<!-- daterangepicker -->
<script src="/static/plugins/moment/moment.min.js"></script>
<script src="/static/plugins/daterangepicker/daterangepicker.js"></script>
<!-- Tempusdominus Bootstrap 4 -->
<script src="/static/plugins/tempusdominus-bootstrap-4/js/tempusdominus-bootstrap-4.min.js"></script>
<!-- Summernote -->
<script src="/static/plugins/summernote/summernote-bs4.min.js"></script>
<!-- overlayScrollbars -->
<script src="/static/plugins/overlayScrollbars/js/jquery.overlayScrollbars.min.js"></script>
<!-- AdminLTE App -->
<script src="/static/js/adminlte.js"></script>
<!-- AdminLTE dashboard demo (This is only for demo purposes) -->
<script src="/static/js/pages/dashboard.js"></script>
<!-- AdminLTE for demo purposes -->
<script src="/static/js/demo.js"></script>
<script>
    $(function () {
        $('#table1').DataTable({
            "paging": true,
            "lengthChange": true,
            "searching": false,
            "ordering": false,
            "info": true,
            "autoWidth": true,
            "responsive": true,
        });
    });

    $(function () {
        $('#table2').DataTable({
            "paging": true,
            "lengthChange": false,
            "searching": false,
            "ordering": false,
            "info": true,
            "autoWidth": false,
            "responsive": true,
        });
    });

    $(function () {
        $('#table3').DataTable({
            "paging": true,
            "lengthChange": false,
            "searching": false,
            "ordering": false,
            "info": true,
            "autoWidth": false,
            "responsive": true,
        });
    });
</script>

<script type="text/javascript">
    $(document).ready(function() {
        $("#addButton").click(function() {
            $('#modalTable').modal({
                show:true,
                backdrop:'static'
            });
        })
    })


    $(document).ready(function() {
        $("#addButton2").click(function() {
            $('#modalTable2').modal({
                show:true,
                backdrop:'static'
            });
        })
    })

    $(document).ready(function() {
        $("#addButton3").click(function() {
            $('#modalTable3').modal({
                show:true,
                backdrop:'static'
            });
        })
    })
</script>

<script type="text/javascript">
    function save(){            //添加一级目录
        var projcatename = $('#projcatename').val();
        if (projcatename){
            $.ajax({
                type:"post",
                url:"/addtree",
                //data: {title:projcatename,code:projcatecode},
                data: {title:projcatename},
                success:function(data){
                    console.log("data:"+data);
                    alert("添加成功！");
                    window.location.reload();
                }
            });
        }
    }


    function save2(){           //添加二级目录
        var childname = $('#childname').val();
        var pId = $("#parentsId").val();
        if (childname){
            $.ajax({
                type:"post",
                url:"/addnewtree",
                data: {title:childname,code:pId},
                // data: {title:childname},
                success:function(){
                    alert("添加成功！");
                    window.location.reload();
                }
            });
        }
    }

    function save3(){           //添加三级目录
        var smallname = $('#smallname').val();
        var pId = $("#pid").val();
        if (smallname){
            $.ajax({
                type:"post",
                url:"/addmorenewtree",
                data: {title:smallname,code:pId},
                // data: {title:childname},
                success:function(){
                    alert("添加成功！");
                    window.location.reload();
                }
            });
        }
    }

</script>

<script type="text/javascript">
    function isShow(parentsid,directory){
        $("#example1").css("display","block");
        $("#example3").css("display","none");
        $("#parentsId").val(parentsid);
        var directorys = directory + "（一级目录）";
        $("#directory").html(directorys);
        $.ajax({
            type:"post",
            url:"/getmulu",
            //data: {title:projcatename,code:projcatecode},
            data: {pid:parentsid},
            success:function(data){
                console.log(data);
                var tablestr = "";
                if (data == null){
                    tablestr = tablestr + "<td>" + "暂无数据" + "</td>" + "</tr>";
                    $("#bodys").html(tablestr)
                }else{
                    var len = data.length;

                    for (var i = 0; i < len; i++){
                        tablestr +=   "<tr onclick='ShowSecond(\"" +   data[i].Title + "\")'>" ;              //onclick带中文时要记得用对引号用转义字符
                        tablestr +=   "<td >" + "<input type='checkbox'>" + "</td>";
                        tablestr +=   "<td>" + data[i].Title + "</td>" + "</tr>"
                    }
                    $("#bodys").html(tablestr)

                }
            }
        });
    }
</script>


<script type="text/javascript">
    function ShowSecond(directory){
        $("#example3").css("display","block");
        $("#parentsTitle").val(directory);
        console.log("test2:"+directory);
        var directorys = directory + "（二级目录）";
        $("#directory2").html(directorys);
        $.ajax({
            type:"post",
            url:"/getmoremulu",
            data: {title:directory},
            success:function(data){
                console.log(data);
                var tablestr = "";
                if (data == null){
                    tablestr = tablestr + "<td>" + "暂无数据" + "</td>" + "</tr>";
                    $("#secondbodys").html(tablestr)
                }else{
                    var len = data.length;
                    for (var i = 0; i < len; i++){
                        tablestr +=   "<tr>";              //onclick
                        tablestr +=   "<td >" + "<input type='checkbox'>" + "</td>";
                        tablestr +=   "<td>" + data[i].Title + "</td>" + "</tr>"
                    }
                    $("#secondbodys").html(tablestr)

                }
            }
        });
    }
</script>

</body>
</html>
