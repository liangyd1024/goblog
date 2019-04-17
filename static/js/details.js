var that;
var vm = new Vue({
    el: "#app",
    mixins: [searchMixin],
    data: {
        praiseFlag: false,

        types: {
            ORIGINAL: '原创',
            REPRINT: '转载'
        },

        articlesRecommendList: null,
        articles: {},

        tags: null,
        categorys: null,
        palceoffiles:null,

        commentList: null,
        bean: {
            articlesId: 0,
            commentator: '',
            email: '',
            content: '',
            parentId: 0,
            paging: {
                pageNum: 1,
                pageSize: 20,
                total: 0,
                pages: 0
            },
        },

        loading: false,//一步加载时的限制
        bottomHight: 50,//滚动条到某个位置才触发时间

        replyMouse: false,
        replyIndex: -1,
        replyForwardIndex: -1,
        replyFlag: false,
        replyComment: {},
    },

    mounted: function () {
        that = this;
        that.articles.id = Number($("#id").val());
        that.bean.commentator = $("#commentCommentator").val();
        that.bean.email = $("#commentEmail").val();
        that.getbowen().then(function () {
            if (vm.articles.articlesDetails.editorType === 'MD') {
                initMarkdown();
            }
            //更新浏览数
            that.browse();
            that.listcomment(1);
        });
        that.listrecommendbowen();
        this.listplaceoffile();
        that.listtag();
        that.listcategory();

        that.tips('praise', '喜欢就戳这里');
    },

    methods: {

        async getbowen() {
            var response = await new Promise(function (resolve) {
                that.httpPost('/articles/get', {id: that.articles.id}, function (response) {
                    resolve(response);
                }, {}, true);
            });
            if (response.data.success) {
                vm.articles = response.data.result;
                vm.bean.articlesId = vm.articles.id;
                document.title = vm.articles.title;
            } else {
                // that.toastErr(response.data.code + "-" + response.data.msg);
            }
        },

        todetails: function (id) {
            return "/articles/" + id
        },

        browse: function () {
            that.httpPost('/articles/browse', {id: that.articles.id}, function (response) {
                if (response.data.success) {
                    if (response.data.result) {
                        that.articles.browseNum = that.articles.browseNum + 1;
                    }
                }
            });
        },

        praise: function () {
            if (that.praiseFlag) {
                that.tips('praise', '已经赞过啦', 'red');
                return;
            }
            that.httpPost('/articles/praise', {id: that.articles.id}, function (response) {
                if (response.data.success) {
                    if (response.data.result) {
                        that.articles.praiseNum = that.articles.praiseNum + 1;
                        that.praiseFlag = true;
                        that.tips('praise', '+1');
                    } else {
                        that.tips('praise', '已经赞过啦', 'red');
                    }
                }
            });
        },

        listcomment: function (pageNum) {
            console.log("listcomment pageNum:%o", pageNum);
            if (pageNum < 1 || (this.bean.paging.pages !== 0 && pageNum > this.bean.paging.pages)) {
                return;
            }
            that.bean.paging.pageNum = pageNum;
            that.httpPost('/articles/comment/list', that.bean, function (response) {
                if (response.data.success) {
                    if (that.commentList === null) {
                        that.commentList = response.data.result;
                    } else {
                        that.commentList.push.apply(that.commentList, response.data.result);
                    }
                    that.bean.paging = response.data.paging;
                }
                vm.loading = false;
            });
        },

        pubcomment: function () {
            if (that.bean.commentator === '' || that.bean.email === '' || that.bean.content === '') {
                return;
            }
            that.httpPost('/articles/comment/pub', that.bean, function (response) {
                if (response.data.success) {
                    that.tips('pubcomment', '评论成功', 'black');
                    that.bean.content = '';
                    that.commentList.unshift(response.data.result);
                    that.articles.commentNum = that.articles.commentNum + 1;
                    that.bean.paging.total = that.bean.paging.total + 1;
                } else {
                    that.toastErr(response.data.code + '-' + response.data.msg);
                }
            });
        },

        loadMore: function () {
            that.listcomment(vm.bean.paging.pageNum + 1);
        },

        handleScroll: function () {
            if (getScrollBottomHeight() <= vm.bottomHight && vm.bean.paging.pageNum < vm.bean.paging.pages && vm.loading === false) {
                vm.loading = true;
                that.listcomment(vm.bean.paging.pageNum + 1);
            }
        },

        commentItem: function (index) {
            return "comment-item-" + index;
        },

        mouseEnter: function (index) {
            this.replyMouse = true;
            this.replyIndex = index;
        },

        mouseLeave: function () {
            this.replyMouse = false;
        },

        replyClick: function (comment, index) {
            window.location.href = "#comment-area";
            this.replyFlag = true;
            this.replyComment = comment;
            this.tips("pubClick", "点我发表评论");
            this.replyForwardIndex = index;
        },

        pubClick: function () {
            this.replyFlag = false;
            this.replyComment = {};
            that.bean.parentId = 0;
        },

        replycomment: function () {
            if (that.bean.commentator === '' || that.bean.email === '' || that.bean.content === '') {
                return;
            }
            that.bean.parentId = that.replyComment.id;
            that.httpPost('/articles/comment/reply', that.bean, function (response) {
                if (response.data.success) {
                    that.tips('replycomment', '回复成功', 'black');
                    that.bean.content = '';
                    this.replyFlag = false;
                    this.replyComment = {};
                    that.replyComment.replyCommentList.unshift(response.data.result);
                    that.articles.commentNum = that.articles.commentNum + 1;
                    setTimeout(function () {
                        window.location.href = "#comment-item-" + that.replyForwardIndex;
                    },500);
                } else {
                    that.toastErr(response.data.code + '-' + response.data.msg);
                }
            });
        }

    }
});


function initMarkdown() {
    editormd.markdownToHTML('content_markdown', {
        htmlDecode: "style,script,iframe",
        path: "/static/js/frame/editormd/lib/",
        emoji: true,
        taskList: true,
        tex: true,  // 默认不解析
        flowChart: true,  // 默认不解析
        sequenceDiagram: true,  // 默认不解析
        tocm: true,         // Using [TOCM]
        // tocContainer    : "#custom-toc-container"
    });
    console.log("======initMarkdown=======");
}

//添加滚动事件
window.onload = function () {
    window.addEventListener('scroll', vm.handleScroll)
};
