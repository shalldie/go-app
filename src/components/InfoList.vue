<template>
    <ul class="info-list">
        <li>---</li>
        <li>---</li>
        <li>---</li>
        <li :class="{active:i==index}" v-for="(item,i) in list">{{item}}</li>
        <li>---</li>
        <li>---</li>
        <li>---</li>
    </ul>
</template>

<script>
export default {
    props: {
        list: {
            type: Array,
            required: true
        }
    },
    data() {
        return {
            index: -1,
            timer: null
        };
    },
    mounted() {
        this.$el.addEventListener('scroll', () => {
            this.throttleGetSelected();
        });
    },
    methods: {
        /**
        * 使用函数节流获取选中项
        */
        throttleGetSelected() {
            clearTimeout(this.timer);
            this.index = -1;
            this.timer = setTimeout(() => {
                this.getSelected();
            }, 500);
        },
        getSelected() {
            let nowTop = this.$el.scrollTop;
            let allTop = this.$el.scrollHeight;
            this.index = ~~(nowTop / allTop) * this.list.length;
            console.log(this.index);
            console.log(nowTop, allTop, this.list.length);
        }
    }
};
</script>

<style lang="scss" scoped>
.info-list {
    margin: 0;
    padding: 0;
    width: 100%;
    height: 100%;
    overflow-x: hidden;
    overflow-y: auto;
    text-align: center;
    font-size: 60px;
    line-height: 80px;
    color: #666;
    background: #f2f2f2;

    li {

        &.active {
            color: black;
            font-weight: bold;
        }
    }
}

 ::-webkit-scrollbar {
    width: 0;
    height: 0
}
</style>
