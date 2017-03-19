{{define "navbar"}}
<a class="navbar-brand" href="/">我的博客</a>
<ur class="nav navbar-nav">
    <li {{if .IsIndex}}class="active"{{end}}><a href="/">首页</a></li>
    <li {{if .IsCategory}}class="active"{{end}}><a href="/category">分类</a></li>
    <li {{if .IsTopic}}class="active"{{end}}><a href="/topic">主题</a></li>
</ur>

<div class="pull-right">
    <ur class="nav navbar-nav">
        {{if .IsLogin}}
        <li><a href="/login?exit=true">退出</a></li>
        {{else}}
        <li><a href="/login">登录</a></li>
        {{end}}
    </ur>

</div>
{{end}}