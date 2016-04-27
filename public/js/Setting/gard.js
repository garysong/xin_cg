function form_submit() {
    var dict = $.trim($("#dict").val());
    if (dict == '') {
        $("#dict").addClass("onFocus").focus();
        $("#dictTip").addClass("onError").html("请选择区名称!");
        return false;
    } 

    var street = $.trim($("#Street").val());
    if (street == '') {
        $("#Street").addClass("onFocus").focus();
        $("#streetTip").addClass("onError").html("请输入街道名称!");
        return false;
    }else{
        $("#Street").removeClass("onFocus");
        $("#streetTip").removeClass("onError").addClass("onCorrect").html("正确");
    }

    var childrengard = $.trim($("#Childrengard").val());
    if (childrengard == '') {
        $("#Childrengard").addClass("onFocus").focus();
        $("#childrengardTip").addClass("onError").html("请输入幼儿园名称!");
        return false;
    }else{
        $("#Childrengard").removeClass("onFocus");
        $("#childrengardTip").removeClass("onError").addClass("onCorrect").html("正确");
    }

    var code = $.trim($("#Code").val());
    if (code == '') {
        $("#Code").addClass("onFocus").focus();
        $("#codeTip").addClass("onError").html("请输入编码!");
        return false;
    }else{
        $("#Code").removeClass("onFocus");
        $("#codeTip").removeClass("onError").addClass("onCorrect").html("正确");
    }
    return true;
}


/**
 * 编辑提交检测
 */
 function form_edit_submit() {
    var dict = $.trim($("#dict").val());
    if (dict == '') {
        $("#dict").focus();
       notice_tips("请选择区名称!");
        return false;
    } 

    var code = $.trim($("#Code").val());
    if (code == '') {
        $("#Code").focus();
        notice_tips("请输入编码!");
        return false;
    } 

    return true;
 }

 /**
 * 删除管理员
 */
function delete_gard(id) {
    if (id == '') {
        notice_tips("参数错误!");
        return false;
    }

    art.dialog.confirm('你确定要删除这个幼儿园吗?', function() {
        $.post("/Gard/Delete/",{'id':id,'csrf_token':csrf_token}, function(tmp){
            if (tmp.status == 1) {
                window.location.reload();
                notice_tips("删除成功!");
            } else {
                notice_tips(tmp.message);
            }
        });
    }, function() {
        notice_tips("你取消了删除幼儿园操作!");
    });
}