var that;

var vm = new Vue({
    el: "#app",
    mixins: [baseMixins],
    data: {
        user: {
            userName: '',
            userPwd: '',
            captcha: ''
        },
        image: ''
    },

    mounted: function () {
        that = this;
        that.captcha();
    },

    methods: {
        login: function () {
            if (that.user.userName === "") {
                that.tips("userName", "用户名不能为空");
                return;
            }
            if (that.user.userPwd === "") {
                that.tips("userPwd", "用户密码不能为空");
                return;
            }
            if (that.user.captcha === "") {
                that.tips("captcha", "验证码不能为空");
                return;
            }
            this.httpPost('/admin/user/login', that.user, function (response) {
                if (response.data.success) {
                    that.redirectUrl("/admin");
                } else {
                    that.toastErr(response.data.code + "-" + response.data.msg);
                    that.user.captcha = "";
                    that.captcha();
                }
            }, {}, true);
        },

        captcha: function () {
            this.httpGet('/admin/user/login/captcha', {}, function (response) {
                if (response.data.success) {
                    that.image = response.data.result
                }
            })
        }
    }
});