<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>上传页面</title>
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
<p>{{.titleid}}</p>
    <div style="margin-left: 300px">
        <form id="fform" method="POST"  action="upload" enctype="multipart/form-data">
            <input id="myfile" name="file" type="file" />
            <input type="submit" value="保存" onclick="add()"  class="btn btn-block btn-outline-primary" style="width: 150px;margin-top: 10px"/>
        </form>
        <input type="hidden" id="titleid" value="{{.titleid}}">
    </div>

</body>
<script>
   function add() {
    var titleid = $("#titleid").val()
       var filename = $("#myfile")[0].value.split("\\")[$("#myfile")[0].value.split("\\").length-1]
        $.ajax({
            type:"post",
            url:"/addfilename",
            data:{titleid:titleid,filename:filename},
            success:function (data) {
                alert("上传成功！")
            }

        })
    }
</script>
</html>