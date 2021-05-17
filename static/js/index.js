var page = 1;
var pageTotal = 4;
var box = new Vue({
    el: '#job-content',
    data: {
        jobs: [{}]
    },
    mounted: function () {
        this.$nextTick(function () {
            $.ajax({
                url: "/page/" + page,
                type: "post",
                dataType: "json",
//                    jsonp: "callback",//传递给请求处理程序或页面的，用以获得jsonp回调函数名的参数名(一般默认为:callback)
//                    jsonpCallback:"flightHandler",//自定义的jsonp回调函数名称，默认为jQuery自动生成的随机函数名，也可以写"?"，jQuery会自
                success: function (msg) {
                    console.log(msg)
                    console.log("hello world")
                    box.jobs = msg.posInfo;
                    // pageTotal = msg.posInfo.pages;

                    if (msg.user != null) {
                        headervue.type = 'user';
                        headervue.person.user = msg.user;
                        console.log(msg.user.Nickname)
                    } else if (msg.hr != null) {
                        headervue.type = 'hr';
                        headervue.person.hr = msg.hr;
                    }


                },
                error: function (msg) {
                    window.location.href = "/user/login";
                }
            });

        });
    }
});

function nextPage() {
    page = page + 1;
    if (page <= pageTotal) {
        $.ajax({
            url: "/page/" + page,
            type: "post",
            dataType: "json",
            success: function (msg) {
                console.log("hello world");
                for (var i = 0; i < msg.posInfo.length; i++) {
                    box.jobs.push(msg.posInfo[i]);
                    console.log(msg.posInfo[i].title)
                }
                // box.jobs.push(msg.posInfo);
            }, error: function () {
                console.log("3err hello world");
                $("#viewMoreButton").empty();
                $("#viewMoreButton").append("没有更多");
            }
        });
    } else {
        $("#viewMoreButton").empty();
        $("#viewMoreButton").append("没有更多");
        $("#viewMoreButton").addClass("disable");
    }
}