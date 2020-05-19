<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>权限分配</title>
    <script src="https://code.jquery.com/jquery-3.1.1.min.js"></script>
</head>
<body>
    <div>
        <table>
            <thead>
                <td>Uid</td>
                <td>Pwd</td>
                <td>Auth</td>
            </thead>
                {{range .List}}
                 <tr>
                     <td>{{.Uid}}</td>
                     <td>{{.Pwd}}</td>
                     <td>
                         {{if eq .Auth "1"}}
                             <input type="checkbox" value="{{.Auth}}" onclick="selectAssign({{.Uid}})" checked>
                         {{else}}
                             <input type="checkbox" value="{{.Auth}}" onclick="selectAssign({{.Uid}})" >
                         {{end}}
                     </td>
                 </tr>
                {{end}}
        </table>
    </div>
</body>
<script>
    function selectAssign(s) {
        $.ajax({
            type:"post",
            url:"/Assign?Uid="+s,
            success:function(result){
                alert("分配成功！")
            },
            error:function () {
                alert("分配失败！")
            }
        })
    }
</script>
</html>