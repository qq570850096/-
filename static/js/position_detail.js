var href = window.location.href;

var box = new Vue({
    el: '#position_content',
    data: {
        position: [{}],
        department: [{}],
        company: [{}],
        category: [{}],
        reviews: [{}]
    },
    created: function () {
        this.$nextTick(function () {
            $.ajax({
                url: href + "/1",
                type: "post",
                dataType: "json",
                async : false,
                success: function (msg) {
                    box.position = msg.position;
                    box.department = msg.department;
                    box.company = msg.company;
                    box.category = msg.category;

                    if (msg.user != null) {
                        header.type = 'user';
                        header.person.user = msg.user;
                    } else if (msg.hr != null) {
                        header.type = 'hr';
                        header.person.hr = msg.hr;
                    }
                    box.reviews.pop();
                    $(msg.comList).each(function (key, val) {
                        console.log(val);
                        box.reviews.push({
                            userName: val.Nickname,
                            reviewDetail: val.Content.replace(/<[^>]+>/g,"")
                        });
                    });


                },
                error: function (msg) {
                    console.log(msg);
                }
            });
        });
    }

});


var favorFlag;
var emailFlag;
$(document).ready(function () {
    $.ajax({
        url: "/user/applyOrNot/" + posId,
        type: "get",
        dataType: "json",
        success: function (msg) {
            if (msg == 1) {
                $("#apply_tag").css("background-color", "#3992d6").css("border", "1px solid #3992d6").text("投递");
                emailFlag = 0;
            } else {
                $("#apply_tag").css("background-color", "#707070").css("border", "1px solid #707070").text("已投递");
                emailFlag = 1;
            }
        },
        error: function (msg) {
            console.log(msg);
        }
    });
    $.ajax({
        url: "/user/favorOrNot/" + posId,
        type: "get",
        dataType: "json",
        success: function (msg) {
            if (msg == "0") {
                $("#favor_tag").css("background-color", "#3992d6").css("border", "1px solid #3992d6").text("收藏");
                favorFlag = 0;
            } else {
                $("#favor_tag").css("background-color", "#707070").css("border", "1px solid #707070").text("取消收藏");
                favorFlag = 1;
            }
        },
        error: function (msg) {
            console.log(msg);
        }
    });

});

var index = href.lastIndexOf("\/");
var posId = href.substr(index + 1, href.length);
// document.getElementById("apply_tag").setAttribute("href","http://localhost:8080/user/apply/"+posId);
// $("#apply_tag").attr("href", "/user/apply/" + posId);
$("#apply_tag").click(function () {
    $.ajax({
        url: "/user/apply/" + posId,
        type: "get",
        dataType: "json",
        success:function (msg) {
            if (msg == "0") {
                window.alert("您已经投递该职位，无需再次投递")
            } else {
                window.location = "/user/success"
                $("#apply_tag").css("background-color", "#707070").css("border", "1px solid #707070").text("已投递");
            }
            $.get("/page/uprec");
        }

    });
    window.location.reload();
});
$("#favor_tag").click(function () {
    if (favorFlag == 0) {
        $.ajax({
            url: "/user/favor/" + posId,
            type: "get",
            dataType: "text",
            success: function (msg) {
                if (msg == "success") {
                    $("#favor_tag").css("background-color", "#707070").css("border", "1px solid #707070").text("取消收藏");
                    favorFlag = 1;
                } else {
                    console.log(msg);
                }
                $.get("/page/uprec");
            },
            error: function (msg) {
                console.log(msg);
            }
        });
        window.location.reload();
    } else {
        $.ajax({
            url: "/user/disfavor/" + posId,
            type: "get",
            dataType: "text",
            success: function (msg) {
                if (msg == "success") {
                    $("#favor_tag").css("background-color", "#3992d6").css("border", "1px solid #3992d6").text("收藏");
                    favorFlag = 0;

                } else {
                    console.log(msg)
                }
                $.get("/page/uprec");
            },
            error: function (msg) {
                console.log(msg);
            }
        });
        window.location.reload();
    }
});

//隐藏表单项提交
$("#posId").val(posId);

