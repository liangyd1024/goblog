var that;

var vm = new Vue({
    el: '#app',
    mixins: [baseMixins,searchMixin],
    data: {

        types: {
            ORIGINAL: '原创',
            REPRINT: '转载'
        },
        articles: {},
        articlesList: null,

        bean: {
            title: '',
            desc: '',
            content: '',
            paging: {
                pageNum: 1,
                pageSize: 5,
                total: 0,
                pages: 0
            },
            status: ''
        },
        tags: null,
        categorys: null,

        //草稿数
        articlesInitNum: 0,
        //发布数
        articlesPublishNum: 0,

        //原创数
        articlesOriginalNum: 0,
        //转载数
        articlesReprintNum: 0,
        //其他
        articlesOtherNum: 0,
    },

    mounted: function () {
        that = this;
        this.collectstatus();
        this.collecttype();
        this.listbowen(1);
        this.listtag();
        this.listcategory();
    },

    methods: {

        collectstatus: function () {
            this.httpGet('/admin/bowen/collect/status', {}, function (response) {
                if (response.data.success) {
                    vm.articlesInitNum = 0;
                    vm.articlesPublishNum = 0;
                    for (index in response.data.result) {
                        if (response.data.result[index].status === 'INIT') {
                            vm.articlesInitNum = response.data.result[index].total
                        } else if (response.data.result[index].status === 'PUBLISH') {
                            vm.articlesPublishNum = response.data.result[index].total
                        }
                    }
                } else {
                    that.toastErr(response.data.code + "-" + response.data.msg);
                }
            });
        },

        collecttype: function () {
            this.httpGet('/admin/bowen/collect/type', {}, function (response) {
                if (response.data.success) {
                    vm.articlesOriginalNum = 0;
                    vm.articlesReprintNum = 0;
                    for (index in response.data.result) {
                        var num = response.data.result[index].total;
                        var type = response.data.result[index].type;
                        if (type === 'ORIGINAL') {
                            vm.articlesOriginalNum = num;
                        } else if (type === 'REPRINT') {
                            vm.articlesReprintNum = num;
                        } else {
                            vm.articlesOtherNum = num;
                        }
                    }
                } else {
                    that.toastErr(response.data.code + "-" + response.data.msg);
                }
            });
        },


        listtag: function () {
            this.httpGet('/admin/tag/list', {}, function (response) {
                if (response.data.success) {
                    vm.tags = response.data.result;
                } else {
                    that.toastErr(response.data.code + "-" + response.data.msg);
                }
            });
        },

        listcategory: function () {
            this.httpGet('/admin/category/list', {}, function (response) {
                if (response.data.success) {
                    vm.categorys = response.data.result;
                } else {
                    that.toastErr(response.data.code + "-" + response.data.msg);
                }
            });
        },

        listbowen: function (pageNum, status = '', type = '') {
            var that = this;
            console.log("listbowen pageNum:%o,status:%o,type:%o", pageNum, status, type);
            if (pageNum < 1 || (this.bean.paging.pages !== 0 && pageNum > this.bean.paging.pages)) {
                return;
            }

            this.bean.paging.pageNum = pageNum;

            var url = "/admin/bowen/list";
            if(this.searchText !== ""){
                url = "/articles/typesearch";
                this.bean.type = "articles";
                this.bean.content = this.searchText;
            }else{
                this.bean.title = this.searchText;
                this.bean.desc = this.searchText;
                this.bean.status = status;
                this.bean.type = type;
            }

            this.httpPost(url, this.bean, function (response) {
                if (response.data.success) {
                    that.articlesList = response.data.result;
                    that.bean.paging = response.data.paging
                } else {
                    that.toastErr(response.data.code + "-" + response.data.msg);
                }
            });
        },

        assign: function (articles) {
            this.articles = articles
        },

        topublish: function () {
            return "/admin/bowen/topublish"
        },

        tomodify: function (id) {
            return "/admin/bowen/tomodify?id=" + id
        },

        todetails: function (id) {
            return "/admin/bowen/todetails?id=" + id
        },

        publishbowen: function () {
            this.httpPost('/admin/bowen/modifystatus', {id: this.articles.id, status: 'PUBLISH'}, function (response) {
                if (response.data.success) {
                    that.toast("发布成功");
                    that.articles.status = 'PUBLISH';
                    that.collectstatus();
                } else {
                    that.toastErr(response.data.code + "-" + response.data.msg);
                }
            });
        },

        deletebowen: function () {
            this.httpPost('/admin/bowen/delete', this.articles, function (response) {
                if (response.data.success) {
                    that.toast("删除成功");
                    that.listbowen.bind(that)(vm.bean.paging.pageNum);
                    that.collectstatus();
                    that.listtag();
                    that.listcategory();
                } else {
                    that.toastErr(response.data.code + "-" + response.data.msg);
                }
            });
        },

        listtagbowen: function (tag) {
            tag.paging = that.bean.paging;
            this.httpPost('/admin/tag/bowen/list', tag, function (response) {
                if (response.data.success) {
                    that.articlesList = response.data.result;
                } else {
                    that.toastErr(response.data.code + "-" + response.data.msg);
                }
            });
        },

        listcategorybowen: function (category) {
            category.paging = that.bean.paging;
            this.httpPost('/admin/category/bowen/list', category, function (response) {
                if (response.data.success) {
                    that.articlesList = response.data.result;
                } else {
                    that.toastErr(response.data.code + "-" + response.data.msg);
                }
            });
        }
    }
});
