//模态框组件
var smallModal = {
    props: ['id', 'title', 'content', 'click-event'],
    template: '<div class="modal inmodal fade" v-bind:id="id" tabindex="-1" role="dialog"  aria-hidden="true">' +
        '<div class="modal-dialog modal-sm">' +
        '<div class="modal-content">' +
        '<div class="modal-header">' +
        '<button type="button" class="close" data-dismiss="modal"><span aria-hidden="true">&times;</span><span class="sr-only">Close</span></button>' +
        '<h5 class="modal-title">{{title}}</h5>' +
        '</div>' +
        '<div class="modal-body">' +
        '   <p><strong>Hi,</strong> {{content}} </p>' +
        '</div>' +
        '<div class="modal-footer">' +
        '<button type="button" class="btn btn-white" data-dismiss="modal">关闭</button>' +
        '<button type="button" class="btn btn-primary" data-dismiss="modal" @click="confirm">确认</button>' +
        '</div>' +
        '</div>' +
        '</div>' +
        '</div>',
    methods: {
        confirm: function (args) {
            this.$emit("click-event", args);
        },
    }
};


Vue.component("small-modal", smallModal);