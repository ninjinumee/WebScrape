



var app = new Vue({
    el: '#app',
    data: {
        message: ''
    },
    methods: {
        postApp: function () {
            // バリデートの判定
            this.$validator.validateAll().then((result) => {
                if (result) {
                    console.log('abc')
                }
            });
        },
    }
})