var page = 1;
var pageTotal = 5;

var box = new Vue({
    el: '#posItem',
    data: {
        orderBy: "salaryUp",
        keyword: "",
        jobs: [{}]
    },
    mounted: function () {

        this.$nextTick(function () {

            if (GetQueryString('orderBy') != null) {
                box.orderBy = GetQueryString('orderBy');
            } else {
                box.orderBy = 'salaryUp';
            }
            if (box.orderBy == "hits") {
                $("#byHits").parent().addClass("active");
            } else if (box.orderBy == "releaseDate") {
                $("#byRelease").parent().addClass("active");
            } else {
                $("#bySalary").parent().addClass("active");
            }

            box.keyword = decodeURIComponent(escape(GetQueryString('keyword')));
            $("#keyword").val(box.keyword);

            var searchItem = {
                keyword: box.keyword,
                orderBy: box.orderBy,
                page: 1
            };

            $.ajax({
                url: "/search",
                type: "post",
                data: searchItem,
                dataType: "json",
                success: function (data) {
                    pageTotal = data.pagetotal;
                    box.jobs = data.posInfo;
                    console.log(data.posInfo);

                    if (data.user != null) {
                        headervue.type = 'user';
                        headervue.person.user = data.user;
                    } else if (data.hr != null) {
                        headervue.type = 'hr';
                        headervue.person.hr = data.hr;
                    }

                },
                error: function (msg) {
                    console.log(msg);
                }
            });

        });
    }
});


$("#bySalary").click(function () {
    window.location.href = "?keyword=" + box.keyword + "&orderBy=salaryUp";
});

$("#byRelease").click(function () {
    window.location.href = "?keyword=" + box.keyword + "&orderBy=releaseDate";
});

$("#byHits").click(function () {
    window.location.href = "?keyword=" + box.keyword + "&orderBy=hits";
});

function nextPage() {
    page = page + 1;
    if (page <= pageTotal) {
        var searchItem = {
            keyword: box.keyword,
            orderBy: box.orderBy,
            page: page
        };
        $.ajax({
            url: "/search",
            type: "post",
            data: searchItem,
            dataType: "json",
            success: function (data) {

                $.each(data.posInfo, function (key, val) {
                    console.log(val);
                    box.jobs.push(val);
                });

            }, error: function () {
                $("#viewMoreButton").empty();
                $("#viewMoreButton").append("????????????");
            }
        });
    } else {
        $("#viewMoreButton").empty();
        $("#viewMoreButton").append("????????????");
        $("#viewMoreButton").addClass("disable");
    }
}

//??????url???orderBy??????
function GetQueryString(name) {
    var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)");
    var r = window.location.search.substr(1).match(reg);
    if (r != null)
        return unescape(r[2]);
    return null;
}
