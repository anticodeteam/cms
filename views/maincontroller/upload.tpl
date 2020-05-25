<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>上传下载页面</title>
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
<style type="text/css">
    .filename{
        background-color: #a1efef;
        border:1px #1b1e21 solid;
        margin-top: 5px;
        border-radius: 5px;
        display: none
    }
    .filename{
        /*height: 120px;*/
        width: 25%;
    }
</style>
<body>
<div style="margin-left: 300px;margin-top: 20px">
    <form id="fform" method="POST"  action="upload" enctype="multipart/form-data">
        <input id="myfile" name="file" type="file" />
        <input type="submit" value="上传" onclick="upload()"  class="btn btn-block btn-outline-primary" style="width: 150px;margin-top: 10px"/>
    </form>
</div>
<div style="margin-left: 300px;margin-top: 20px">
{{/*    <input id="myfile" name="file" type="file" src="../../SDPATH"/>*/}}
        <button class="btn btn-block btn-outline-primary" style="width: 150px;margin-top: 10px" id="GetName" onclick="getfilename()">获取文件</button>
{{/*        <button class="btn btn-block btn-outline-primary" style="width: 150px;margin-top: 10px" id="GetName"><a href="{{.src}}" download="">下载文件</a></button>*/}}
    <div class="filename">
        <table id="tab">

        </table>
    </div>

{{/*    <button class="btn btn-block btn-outline-primary" style="width: 150px;margin-top: 10px" onclick="download()">下载</button>*/}}
    <input type="hidden" id="titleid" value="{{.titleid}}">
</div>

</body>
<script>
    //上传
    function upload() {
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
    // 获取文件名
     function getfilename() {
         $.ajax({
             type:"post",
             url:"/Getfilename",
             //data:{titleid:titleid,filename:filename},
             success:function (data) {
                 $(".filename").slideToggle("slow")
                 console.log(data)
                 var str = "";
                 for (var i=0;i<data.length;i++){
                     str += "<tr><td style=\'width: 180px;\'>" + data[i] +
                         "</td><td><button class=\"btn btn-block btn-outline-primary\" style=\"width: 60px;margin-top: 5px\" " +
                          "onclick=\"download('"+data[i]+"')\" value="+data[i]+">下载</button>" + "</td></tr>"
                 }
                 $("#tab").html(str);
             }
         })
     }
    //下载文件
    function download(filename) {
        window.location.href = "/Download?filename="+filename
    }
</script>
