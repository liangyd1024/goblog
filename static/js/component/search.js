var searchMixin = {
    mixins:[baseMixins],
    data:{
        searchText:''
    },
    methods: {
        tosearch: function (type, id, content = '') {
            console.log("call search type:%o,id:%o", type, id);
            if(id === 0 && content === ''){
                this.tips("searchform","请输入搜索关键字");
                return;
            }
            window.location.href = "/articles/tosearch?type=" + type + "&id=" + id + "&content=" + content;
        },

        listtag: function () {
            this.httpGet('/articles/tag/list', {}, function (response) {
                if (response.data.success) {
                    vm.tags = response.data.result;
                } else {
                    that.toastErr(response.data.code + "-" + response.data.msg);
                }
            });
        },

        listcategory: function () {
            this.httpGet('/articles/category/list', {}, function (response) {
                if (response.data.success) {
                    vm.categorys = response.data.result;
                } else {
                    that.toastErr(response.data.code + "-" + response.data.msg);
                }
            });
        },

        listrecommendbowen:function(){
            this.httpGet('/articles/recommend/list', {}, function (response) {
                if (response.data.success) {
                    vm.articlesRecommendList = response.data.result;
                } else {
                    that.toastErr(response.data.code + "-" + response.data.msg);
                }
            });
        },

        listplaceoffile: function () {
            this.httpGet('/articles/place-of-file/list', {}, function (response) {
                if (response.data.success) {
                    vm.palceoffiles = response.data.result;
                } else {
                    that.toastErr(response.data.code + "-" + response.data.msg);
                }
            });
        },
    }
};