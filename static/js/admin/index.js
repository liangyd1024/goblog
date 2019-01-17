var that;
var vm = new Vue({
    el: '#wrapper',
    mixins: [baseMixins],
    mounted: function () {
        that = this;
    },
    methods: {
        logout: function () {
            this.httpGet('/admin/user/logout', {}, function (response) {
                if (response.data.success && response.data.result) {
                    that.back('/admin');
                }
            });
        }
    }
});