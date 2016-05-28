function form_submit() {
    
    return true;
}

function form_edit_submit() {
    return true    
}


function delete_equipment(id) {
    if (id == '') {
        notice_tips("参数错误!");
        return false;
    }

    art.dialog.confirm('你确定要删除这个消防设施吗?', function() {
        $.post("/Equipment/delete/",{'id':id,'csrf_token':csrf_token}, function(tmp){
            if (tmp.status == 1) {
                window.location.reload();
                notice_tips("删除成功!");
            } else {
                notice_tips(tmp.message);
            }
        });
    }, function() {
        notice_tips("你取消了删除消防设施操作!");
    });
}