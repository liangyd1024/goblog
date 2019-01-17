var that;
var vm = new Vue({
    el: "#app",
    mixins: [searchMixin],
    data: {

        types: {
            ORIGINAL: '原创',
            REPRINT: '转载'
        },

        articlesRecommendList: null,
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
        palceoffiles:null,

        loading: false,//一步加载时的限制
        bottomHight: 50,//滚动条到某个位置才触发时间
    },
    mounted: function () {
        that = this;
        this.listbowen(1);
        this.listrecommendbowen();
        this.listplaceoffile();
        this.listtag();
        this.listcategory();
    },
    methods: {

        listbowen: function (pageNum) {
            that.loading = true;
            console.log("listbowen pageNum:%o", pageNum);
            if (pageNum < 1 || (this.bean.paging.pages !== 0 && pageNum > this.bean.paging.pages)) {
                return;
            }
            this.bean.paging.pageNum = pageNum;
            this.bean.status = 'PUBLISH';
            this.bean.title = this.searchText;
            this.bean.desc = this.searchText;

            this.httpPost('/articles/list', this.bean, function (response) {
                if (response.data.success) {
                    if (that.articlesList === null) {
                        that.articlesList = response.data.result;
                    } else {
                        that.articlesList.push.apply(that.articlesList, response.data.result);
                    }
                    that.bean.paging = response.data.paging
                } else {
                    // that.toastErr(response.data.code + "-" + response.data.msg);
                }
                that.loading = false;
            });
        },

        todetails: function (id) {
            return "/articles/" + id
        },

        touchMore: function (index) {
            if (vm.loading) {
                return;
            }
            if (index < (vm.bean.paging.pageNum * vm.bean.paging.pageSize -1)) {
                return;
            }
            that.listbowen(vm.bean.paging.pageNum + 1);
        },

        loadMore: function () {
            that.listbowen(vm.bean.paging.pageNum + 1);
        },

        handleScroll: function () {
            if (getScrollBottomHeight() <= vm.bottomHight && vm.bean.paging.pageNum < vm.bean.paging.pages && vm.loading === false) {
                that.listbowen(vm.bean.paging.pageNum + 1);
            }
        }
    }
});

//添加滚动事件
window.onload = function () {
    window.addEventListener('scroll', vm.handleScroll)
};