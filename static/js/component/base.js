var baseMixins = {

    data: {
        isMobile: !!isMobile()
    },

    methods: {

        //时间格式化函数，此处仅针对yyyy-MM-dd hh:mm:ss 的格式进行格式化
        dateFormat: function (time, format = true) {
            var date = new Date(time);
            var year = date.getFullYear();
            /* 在日期格式中，月份是从0开始的，因此要加0
             * 使用三元表达式在小于10的前面加0，以达到格式统一  如 09:11:05
             * */
            var month = date.getMonth() + 1 < 10 ? "0" + (date.getMonth() + 1) : date.getMonth() + 1;
            var day = date.getDate() < 10 ? "0" + date.getDate() : date.getDate();
            if (format) {
                var hours = date.getHours() < 10 ? "0" + date.getHours() : date.getHours();
                var minutes = date.getMinutes() < 10 ? "0" + date.getMinutes() : date.getMinutes();
                var seconds = date.getSeconds() < 10 ? "0" + date.getSeconds() : date.getSeconds();
                return year + "-" + month + "-" + day + " " + hours + ":" + minutes + ":" + seconds;
            } else {
                return year + "-" + month + "-" + day;
            }
        },

        redirectUrl: function (url) {
            if (url) {
                window.location.href = url;
            }
        },

        back: function () {
            history.back(-1);
        },

        toast: function (content, type = 'success') {
            toastr.options = {
                "closeButton": true,
                "debug": true,
                "progressBar": true,
                "positionClass": "toast-top-right",
                "onclick": null,
                "showDuration": "400",
                "hideDuration": "1000",
                "timeOut": "2000",
                "extendedTimeOut": "1000",
                "showEasing": "swing",
                "hideEasing": "linear",
                "showMethod": "fadeIn",
                "hideMethod": "fadeOut"
            };
            toastr[type](content || "操作成功", "提示");
        },

        tips: function (id, content, color = "#asdsde") {
            layer.tips(content, '#' + id, {
                tips: [1, color],
                time: 1500
            });
        },

        toastWarn: function (content) {
            this.toast(content, "warning");
        },

        toastErr: function (content) {
            this.toast(content, "error");
        },

        http: function (method, url, data, callback, headers = {}, mask = false) {
            var that = this;
            console.log("call req url:%o,data:%o,mask:%o", url, data, mask);
            if (mask) {
                var index = parent.layer.load(1000, {shade: [0.5, '#eee']});
            }
            axios({
                method: method,
                url: url,
                data: data,
                header: headers,
                timeout: 30000,
                // baseURL: 'https://some-domain.com/api/',
            }).then(function (response) {
                console.log("call resp url:%o,response:%o", url, response);
                if (callback) {
                    callback(response);
                }
                if (mask) {
                    parent.layer.close(index);
                }
            }).catch(function (error) {
                console.log("call err url:%o,error:%o", url, error);
                that.toastErr('请求超时!');
                if (mask) {
                    parent.layer.close(index);
                }
            })
        },

        httpGet: function (url, data = {}, callback, headers, mask) {
            this.http("GET", url, data, callback, headers, mask);
        },

        httpPost: function (url, data = {}, callback, headers, mask) {
            this.http("POST", url, data, callback, headers, mask);
        }
    }
};


//滚动条到底部的距离
function getScrollBottomHeight() {
    return getPageHeight() - getScrollTop() - getWindowHeight();

}

//页面高度
function getPageHeight() {
    return document.querySelector("html").scrollHeight
}

//滚动条顶 高度
function getScrollTop() {
    var scrollTop = 0, bodyScrollTop = 0, documentScrollTop = 0;
    if (document.body) {
        bodyScrollTop = document.body.scrollTop;
    }
    if (document.documentElement) {
        documentScrollTop = document.documentElement.scrollTop;
    }
    scrollTop = (bodyScrollTop - documentScrollTop > 0) ? bodyScrollTop : documentScrollTop;
    return scrollTop;
}

//获取窗口高度
function getWindowHeight() {
    var windowHeight = 0;
    if (document.compatMode === "CSS1Compat") {
        windowHeight = document.documentElement.clientHeight;
    } else {
        windowHeight = document.body.clientHeight;
    }
    return windowHeight;
}

//是否移动端
function isMobile() {
    let flag = navigator.userAgent.match(/(phone|pad|pod|iPhone|iPod|ios|iPad|Android|Mobile|BlackBerry|IEMobile|MQQBrowser|JUC|Fennec|wOSBrowser|BrowserNG|WebOS|Symbian|Windows Phone)/i);
    return flag;
}

