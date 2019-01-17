var that;

var unload;

window.onbeforeunload = function (e) {
    if(unload){
        e = window.event || e;
        e.returnValue = ("确定离开当前页面吗？");
    }
};

function getMarkdown() {
    return editormd("content_markdown", {
        placeholder: "此处开始编写您要发布的内容...",
        width: "100%",
        // autoHeight: true,
        height: 350,
        codeFold: true,
        previewCodeHighlight : true,
        htmlDecode: "style,script,iframe",
        syncScrolling: "single",
        path: "/static/js/frame/editormd/lib/",
        emoji: true,
        taskList: true,
        toc: true,         // Using [TOCM]
        tex: true,                   // 开启科学公式TeX语言支持，默认关闭
        flowChart: true,             // 开启流程图支持，默认关闭
        sequenceDiagram: true,       // 开启时序/序列图支持，默认关闭,
        //这个配置在simple.html中并没有，但是为了能够提交表单，使用这个配置可以让构造出来的HTML代码直接在第二个隐藏的textarea域中，方便post提交表单。
        saveHTMLToTextarea: true,
        imageUpload: true,
        imageFormats: ["jpg", "jpeg", "gif", "png", "bmp"],
        imageUploadURL: "/admin/file/upload/editor",
    });
}


var vm = new Vue({
    el: '#app',
    mixins: [baseMixins],
    data: {
        typeMap: {
            ORIGINAL: '原创',
            REPRINT: '转载'
        },
        types: [
            {code: 'ORIGINAL', desc: '原创'},
            {code: 'REPRINT', desc: '转载'},
        ],

        //所有标签集
        tags: null,
        tagsMap: new Map(),
        //需要添加的标签
        newTags: null,
        tag: {
            tagName: ''
        },
        //所有栏目集
        categorys: null,
        categorysMap: new Map(),
        //需要添加的栏目
        newCategorys: null,
        category: {
            categoryName: ''
        },

        articles: {
            id: 0,
            title: '',
            desc: '',
            status: '',
            type: '',
            articlesDetails: {
                content: ''
            },
            //标签集
            tags: null,
            //栏目集
            categorys: null,
        },

        modifyFlag: false,
        descEditFlag: false,
        contentEditFlag: false,

        editorType: {
            markdownActive: true,
            md: 'MD',
            richActive: false,
            rich: 'RICH',
        },
        //markdown对象
        editormd: {},
        selectEditorType: 'MD',
        //rich对象
        editorrich: {}

    },

    mounted: function () {
        that = this;
        this.editormd = getMarkdown();
        this.editormd.on("previewing", this.onpreviewing);
        this.editormd.on("previewed", this.onpreviewed);
        this.editormd.on("load", this.mdInit);

        this.listtag();
        this.listcategory();
    },

    methods: {

        mdInit: function () {
            this.editormd.previewing();
            var id = $("#id");
            this.articles.id = Number(id.val());
            if (id.val() !== '') {
                this.modifyFlag = true;
                that.getbowen.bind(this)();
            }
        },

        //预览触发 on previewing you can custom css .editormd-preview-activeon previewing you can custom css .editormd-preview-active
        onpreviewing: function () {
            console.log("onpreviewing =>", this, this.id, this.settings);
            this.contentEditFlag = false;
        },
        //取消预览触发
        onpreviewed: function () {
            console.log("onpreviewed =>", this, this.id, this.settings);
            this.contentEditFlag = true;
        },

        selecteditor: function (editorType) {
            console.log("selecteditor:", editorType);
            this.selectEditorType = editorType;
        },

        edit: function (id) {
            unload = true;
            if ("desc" === id) {
                this.descEditFlag = true;
            } else {
                this.contentEditFlag = true;
                if (this.selectEditorType === this.editorType.md) {
                    this.editormd.previewed();
                }
            }
            $("#" + id).addClass("no-padding");
            this.editorrich = $("#" + id + "_edit").summernote({
                lang: "zh-CN",
                focus: true,
                placeholder: '请输入博文内容',
                onImageUpload: function (files, editor, welEditable) {
                    console.log("call editorrich onImageUpload");
                    that.sendfile(files, editor, welEditable);
                }
            });
        },

        save: function (id) {
            $("#" + id).removeClass("no-padding");
            console.log("save id:",id);
            var code = $("#" + id + "_edit").code();
            $("#" + id + "_edit").destroy();
            console.log("save code:",code);
            if ("desc" === id) {
                this.articles.desc = code;
                this.descEditFlag = false;
            } else {
                console.log("save selectEditorType:%o,",this.selectEditorType);
                if (this.selectEditorType === this.editorType.rich) {
                    this.articles.articlesDetails.content = code;
                } else {
                    if (this.contentEditFlag) {
                        this.editormd.previewing();
                        this.articles.articlesDetails.content = this.editormd.getMarkdown();
                    }
                }
                this.contentEditFlag = false;
            }

        },

        sendfile: function (files, editor, welEditable) {
            //MD和RICH统一文件名
            let imageData = new FormData();
            imageData.append("editormd-image-file", files[0]);
            this.httpPost('/admin/file/upload/editor', imageData, function (response) {
                if (response.data.success) {
                    that.toast(response.data.message);
                    editor.insertImage(welEditable, response.data.url);
                } else {
                    that.toastErr(response.data.code + "-" + response.data.msg);
                }
            }, {
                'Content-Type': 'multipart/form-data'
            });
        },

        getbowen: function () {
            this.httpPost('/admin/bowen/get', {id: this.articles.id}, function (response) {
                if (response.data.success) {
                    vm.articles = response.data.result;
                    if (vm.articles.articlesDetails.editorType === vm.editorType.rich) {
                        vm.editorType.richActive = true;
                        vm.editorType.markdownActive = false;
                        vm.selectEditorType = vm.editorType.rich
                    } else {
                        vm.editormd.appendMarkdown(vm.articles.articlesDetails.content);
                    }
                } else {
                    that.toastErr(response.data.code + "-" + response.data.msg);
                }
            }, {}, true);
        },

        publish: function (status) {
            if (this.descEditFlag || this.contentEditFlag) {
                that.toastErr("请先保存修改信息");
                return;
            }

            let articles = Object.assign({}, this.articles);
            articles.status = status;
            articles.tags = this.newTags;
            articles.categorys = this.newCategorys;
            articles.articlesDetails.editorType = this.selectEditorType;

            this.httpPost('/admin/bowen/publish', articles, function (response) {
                if (response.data.success) {
                    that.toast("发表成功");
                    articles = response.data.result;
                    unload = false;
                    that.redirectUrl("/admin/bowen/todetails?id=" + articles.id);
                } else {
                    that.toastErr(response.data.code + "-" + response.data.msg);
                }
            }, {}, true);
        },

        modifybowen: function () {
            if (this.descEditFlag || this.contentEditFlag) {
                that.toastErr("请先保存修改信息");
                return;
            }

            this.articles.tags = this.newTags;
            this.articles.categorys = this.newCategorys;

            this.httpPost('/admin/bowen/modify', this.articles, function (response) {
                if (response.data.success) {
                    that.toast("修改发表成功");
                    that.articles = response.data.result;
                    unload = false;
                    that.redirectUrl("/admin/bowen/todetails?id=" + that.articles.id);
                } else {
                    that.toastErr(response.data.code + "-" + response.data.msg);
                }
            }, {}, true);
        },

        listtag: function () {
            this.httpGet('/admin/tag/list', {}, function (response) {
                if (response.data.success) {
                    vm.tags = response.data.result;
                    vm.tags.forEach(function (tag) {
                        vm.tagsMap.set(tag.tagName, tag)
                    });
                } else {
                    that.toastErr(response.data.code + "-" + response.data.msg);
                }
            });
        },

        addtag: function () {
            if (this.tag.tagName === "") {
                that.tips("tagNameTip", "标签名不能为空");
                return
            }
            if (vm.tagsMap.has(this.tag.tagName)) {
                that.tips("tagNameTip", "标签已存在");
                return
            }

            this.httpPost('/admin/tag/add', this.tag, function (response) {
                if (response.data.success) {
                    that.toast("添加成功");
                    that.selecttag(response.data.result, true);
                    that.listtag();
                } else {
                    that.toastErr(response.data.code + "-" + response.data.msg);
                }
            });
        },

        deletetag: function (tag) {
            this.httpPost('/admin/tag/delete', tag, function (response) {
                if (response.data.success) {
                    that.toast("删除成功");
                    that.deletenewtag(tag);
                    that.listtag();
                } else {
                    that.toastErr(response.data.code + "-" + response.data.msg);
                }
            });
        },

        selecttag: function (tag, action) {
            console.log("selecttag tag:%o,action:%o", tag, action);
            if (vm.newTags == null) {
                vm.newTags = Array()
            }
            if (vm.articles.tags == null) {
                vm.articles.tags = Array()
            }
            if (action) {
                for (let i in vm.newTags) {
                    if (vm.newTags[i].id === tag.id) {
                        return;
                    }
                }
                for (let i in vm.articles.tags) {
                    if (vm.articles.tags[i].id === tag.id) {
                        return;
                    }
                }
                vm.newTags.push(tag);
            } else {
                for (let i in vm.articles.tags) {
                    if (vm.articles.tags[i].id === tag.id) {
                        vm.articles.tags.splice(i, 1);
                        return;
                    }
                }
            }

        },

        deletearticlestag: function (tag) {
            if (vm.modifyFlag) {//修改文章标志时删除关系
                this.httpPost('/admin/bowen/tag/delete', {
                    articlesId: vm.articles.id,
                    tagId: tag.id
                }, function (response) {
                    if (response.data.success) {
                        that.toast("删除成功");
                        that.selecttag(tag, false);
                    } else {
                        that.toastErr(response.data.code + "-" + response.data.msg);
                    }
                });
            }
        },

        deletenewtag: function (tag) {
            console.log("del deletenewtag:", tag);
            for (let i in vm.newTags) {
                if (vm.newTags[i].id === tag.id) {
                    console.log("deletenewtag splice id:", tag.id);
                    vm.newTags.splice(i, 1);
                    return;
                }
            }
        },

        listcategory: function () {
            this.httpGet('/admin/category/list', {}, function (response) {
                if (response.data.success) {
                    vm.categorys = response.data.result;
                    vm.categorys.forEach(function (category) {
                        vm.categorysMap.set(category.categoryName, category)
                    });
                } else {
                    that.toastErr(response.data.code + "-" + response.data.msg);
                }
            });
        },

        addcategory: function () {
            if (this.category.categoryName === "") {
                that.tips("categoryNameTip", "栏目名不能为空");
                return
            }
            if (vm.categorysMap.has(this.category.categoryName)) {
                that.tips("categoryNameTip", "栏目已存在");
                return
            }

            this.httpPost('/admin/category/add', this.category, function (response) {
                if (response.data.success) {
                    that.toast("添加成功");
                    that.selectcategory(response.data.result, true);
                    that.listcategory();
                } else {
                    that.toastErr(response.data.code + "-" + response.data.msg);
                }
            });
        },

        deletecategory: function (category) {
            this.httpPost('/admin/category/delete', category, function (response) {
                if (response.data.success) {
                    that.toast("删除成功");
                    that.deletenewcategory(category);
                    that.listcategory();
                } else {
                    that.toastErr(response.data.code + "-" + response.data.msg);
                }
            });
        },

        selectcategory: function (category, action) {
            console.log("selectcategory category:%o,action:%o", category, action);
            if (vm.newCategorys == null) {
                vm.newCategorys = Array()
            }
            if (vm.articles.categorys == null) {
                vm.articles.categorys = Array()
            }
            if (action) {
                for (let i in vm.newCategorys) {
                    if (vm.newCategorys[i].id === category.id) {
                        return;
                    }
                }
                for (let i in vm.articles.categorys) {
                    if (vm.articles.categorys[i].id === category.id) {
                        return;
                    }
                }
                console.log("selectcategory id:", category.id);
                vm.newCategorys.push(category);
            } else {
                console.log("del articles.categorys:", vm.articles.categorys);
                for (let i in vm.articles.categorys) {
                    if (vm.articles.categorys[i].id === category.id) {
                        console.log("articles.categorys splice id:", category.id);
                        vm.articles.categorys.splice(i, 1);
                        return;
                    }
                }
            }

        },

        deletearticlescategory: function (category) {
            if (vm.modifyFlag) {//修改文章标志时删除关系
                this.httpPost('/admin/bowen/category/delete', {
                    articlesId: vm.articles.id,
                    categoryId: category.id
                }, function (response) {
                    if (response.data.success) {
                        that.toast("删除成功");
                        that.selectcategory(category, false);
                    } else {
                        that.toastErr(response.data.code + "-" + response.data.msg);
                    }
                });
            }
        },

        deletenewcategory: function (category) {
            console.log("del newCategorys:", category);
            for (let i in vm.newCategorys) {
                if (vm.newCategorys[i].id === category.id) {
                    console.log("newCategorys splice id:", category.id);
                    vm.newCategorys.splice(i, 1);
                    return;
                }
            }
        },


        selecttype: function (type) {
            vm.articles.type = type.code
        }

    }
});
