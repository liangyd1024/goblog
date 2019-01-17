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
    });
    console.log("======initMarkdown=======");
}

var that;

var vm = new Vue({
    el: '#app',
    mixins: [baseMixins],
    data: {
        articles: {
            id: 0,
            title: '',
            desc: '',
            status: '',
            articlesDetails: {
                content: ''
            },
        },
        comment:{
            articlesId:'',
            paging: {
                pageNum: 1,
                pageSize: 20,
                total: 0,
                pages: 0
            },
        },

        commentList: null,
    },

    mounted: function () {
        that = this;
        this.articles.id = Number($("#id").val());
        this.comment.articlesId = this.articles.id;
        this.getbowen().then(function () {
            if (vm.articles.articlesDetails.editorType === 'MD') {
                initMarkdown();
            }
        });
        this.listcomment(1);
    },

    methods: {
        async getbowen() {
            var response = await new Promise(function (resolve) {
                that.httpPost('/admin/bowen/get', {id: that.articles.id}, function (response) {
                    resolve(response);
                }, {}, true);
            });
            if (response.data.success) {
                vm.articles = response.data.result;
            } else {
                that.toastErr(response.data.code + "-" + response.data.msg);
            }
        },

        listcomment: function (pageNum) {
            console.log("listcomment pageNum:%o", pageNum);
            if (pageNum < 1 || (this.comment.paging.pages !== 0 && pageNum > this.comment.paging.pages)) {
                return;
            }
            that.comment.paging.pageNum = pageNum;
            that.httpPost('/admin/bowen/comment/list', that.comment, function (response) {
                if (response.data.success) {
                    if (that.commentList === null) {
                        that.commentList = response.data.result;
                    } else {
                        that.commentList.push.apply(that.commentList, response.data.result);
                    }
                    that.comment.paging = response.data.paging;
                }
            });
        },

    }
});
