var that;
var vm = new Vue({
    el: "#app",
    mixins: [searchMixin],
    data: {
        types: {
            ORIGINAL: '原创',
            REPRINT: '转载'
        },

        bean: {
            id: 0,
            type: '',
            content: '',
            paging: {
                pageNum: 1,
                pageSize: 10,
                total: 0,
                pages: 0
            },
        },

        articlesList: null,

        loading: false,//一步加载时的限制
        bottomHight: 50,//滚动条到某个位置才触发时间
    },
    mounted: function () {
        that = this;
        that.bean.id = parseInt($("#id").val());
        that.bean.type = $("#type").val();
        that.bean.content = $("#content").val();
        that.searchText = that.bean.content;
        that.typesearch(1);
    },
    methods: {
        typesearch: function (pageNum) {
            console.log("typesearch pageNum:%o,type:%o,id:%o", pageNum, that.bean.type, that.bean.id);
            if (pageNum < 1 || (this.bean.paging.pages !== 0 && pageNum > this.bean.paging.pages)) {
                return;
            }
            this.bean.paging.pageNum = pageNum;
            this.httpPost("/articles/typesearch", this.bean, function (response) {
                if (response.data.success) {
                    if (that.articlesList === null) {
                        that.articlesList = response.data.result;
                    } else {
                        that.articlesList.push.apply(that.articlesList, response.data.result);
                    }
                    that.bean.paging = response.data.paging;
                } else {
                    // that.toastErr(response.data.code + "-" + response.data.msg);
                }
                vm.loading = false;
            }, {}, true)
        },

        todetails: function (id) {
            return "/articles/" + id
        },

        loadMore: function () {
            that.typesearch(vm.bean.paging.pageNum + 1);
        },

        handleScroll: function () {
            if (getScrollBottomHeight() <= vm.bottomHight && vm.bean.paging.pageNum < vm.bean.paging.pages && vm.loading === false) {
                vm.loading = true;
                that.typesearch(vm.bean.paging.pageNum + 1);
            }
        }
    }
});


//添加滚动事件
window.onload = function () {
    window.addEventListener('scroll', vm.handleScroll)
};