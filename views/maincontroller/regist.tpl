<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>注册页面</title>
    <link rel="stylesheet" href="/static/plugins/fontawesome-free/css/all.min.css">
    <!-- Ionicons -->
    <link rel="stylesheet" href="https://code.ionicframework.com/ionicons/2.0.1/css/ionicons.min.css">
    <!-- icheck bootstrap -->
    <link rel="stylesheet" href="/static/plugins/icheck-bootstrap/icheck-bootstrap.min.css">
    <!-- Theme style -->
    <link rel="stylesheet" href="/static/css/adminlte.min.css">
    <!-- Google Font: Source Sans Pro -->
    <link href="https://fonts.googleapis.com/css?family=Source+Sans+Pro:300,400,400i,700" rel="stylesheet">
    <!--Def-->
    <link rel="stylesheet" href="/static/css/Def.css">

    <script src="https://code.jquery.com/jquery-3.1.1.min.js"></script>
</head>
<body>
<div style="margin:100px 30%">
    <div class="card card-info">
        <div class="card-header">
            <h3 class="card-title">注册信息</h3>
        </div>
        <!-- /.card-header -->
        <!-- form start -->
        <form class="form-horizontal" action="/registDBInfoAction" method="post">

            <div class="card-body">
                <div class="form-group row">
                    <label for="UserName" class="col-sm-2 col-form-label">用户名</label>
                    <div class="col-sm-10">
                        <input id="UserName" type="text" class="form-control" name="Uid" placeholder="请输入用户名" onfocus="this.removeAttribute('readonly')">
                    </div>
                </div>
                <div class="form-group row">
                    <label for="inputPassword" class="col-sm-2 col-form-label">密码</label>
                    <div class="col-sm-10">
                        <input id="inputPassword" type="password" class="form-control" name="Pwd" placeholder="请输入密码" onfocus="this.removeAttribute('readonly')">
                    </div>
                </div>

            </div>
            <!-- /.card-body -->
            <div class="card-footer white" style="float:right">
                <button type="submit" class="btn btn-info">Sign in</button>
                <input  type="button" class="btn btn-info" value="Cancel" onclick="cancel()">
            </div>
            <!-- /.card-footer -->
        </form>
    </div>
</div>
</body>
<script>
    $(function () {
        $("#UserName").val()
        $('#UserName').attr('readonly','true');
        $("#inputPassword").val()
        $('#inputPassword').attr('readonly','true');

        if({{.result}}){
            alert({{.result}})
        }
    })

    function cancel() {
        window.location.href = "/RegisCancel"
    }
</script>
</html>