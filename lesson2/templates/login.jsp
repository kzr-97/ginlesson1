 <%@ page language="java" import="java.util.*" pageEncoding="UTF-8"%>
 <%@ include file="common.jsp"%>
<html>
  <head>
    <title>${webTitle}</title>
    <meta http-equiv="pragma" content="no-cache">
    <meta http-equiv="cache-control" content="no-cache">
    <meta http-equiv="expires" content="0">    
  </head>
  
  <body>
	    <jsp:include page="top.jsp"></jsp:include>
		
		<div class="layui-container fly-marginTop">
		  <div class="fly-panel fly-panel-user" pad20>
		    <div class="layui-tab layui-tab-brief" lay-filter="user">
		      <ul class="layui-tab-title">
		        <li class="layui-this">登入</li>
		        <li><a href="showReg">注册</a></li>
		      </ul>
		      <div class="layui-form layui-tab-content" id="LAY_ucm" style="padding: 20px 0;">
		        <div class="layui-tab-item layui-show">
		          <div class="layui-form layui-form-pane">
		              <div class="layui-form-item">
		                <label for="uname" class="layui-form-label">用户名</label>
		                <div class="layui-input-inline">
		                  <input type="text" id="uname" name="uname" required lay-verify="required" autocomplete="off" class="layui-input">
		                </div>
		              </div>
		              <div class="layui-form-item">
		                <label for="upass" class="layui-form-label">密码</label>
		                <div class="layui-input-inline">
		                  <input type="password" id="upass" name="upass" required lay-verify="required" autocomplete="off" class="layui-input">
		                </div>
		              </div>
		              <div class="layui-form-item">
		                <button class="layui-btn" onclick="login()">立即登录</button>
		              </div>
		          </div>
		        </div>
		      </div>
		    </div>
		  </div>
		</div>
	</body>
<script>
	function login(){     
		 var uname=document.getElementById("uname").value;  
		 var upass=document.getElementById("upass").value; 
		 if(uname==""){
		 	layer.msg('请输入用户名 ');
		 	return false;
		 }
		 if(upass==""){
		 	layer.msg('请输入密码 ');
		 	return false;
		 }
		 $.ajax({  
		        type: "POST",      
		        url: "login", //servlet的名字
		        data: "staff="+uname+"&pwd="+upass,
		        dataType: "json",
		        success: function(data){   
				    if(data.status==0){     
				    	location.href="listBlock";
				    }else{
			    		layer.msg(data.message);
				    }
		 		}     
		});     
	}  
</script>
</html>
