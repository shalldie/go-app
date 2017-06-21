<template>
    <ul @touchstart="onTouchStart" @touchmove="onTouchMove" class="select-list">
        <li v-for="(item,index) in list" class="select-list-item"></li>
    </ul>
</template>

<script>
export default {
    props: {
        list: {
            required: true,
            type: Array
        }
    },
    data() {
        return {
            touchs: []
        };
    },
    methods: {
        onTouchStart(ex) {
            this.touches[0] = this.touches[1] = ex.touches[0];
        },
        onTouchMove(ex) {
            // 保存位置
            this.touchs[0] = this.touchs[1];
            this.touchs[1] = {
                time: +new Date,
                pageY: ex.touches[0].pageY
            };
        }
    }
}
</script>

<style lang="scss" scoped>
.select-list {
    position: relative;
    width: 100%;
    height: 238px;
    box-sizing: border-box;

    &-item {
        box-sizing: border-box;
    }
}
</style>