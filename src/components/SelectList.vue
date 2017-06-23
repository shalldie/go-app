<template>
    <ul @touchstart="onTouchStart" @touchmove="onTouchMove" class="select-list">
        <li v-for="(item,index) in list" class="select-list-item" :key="index">{{item}}</li>
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
            diffY: 0,
            touches: []
        };
    },
    methods: {
        onTouchStart(ex) {
            this.touches[0] = this.touches[1] = ex.touches[0];
        },
        onTouchMove(ex) {
            // 保存位置
            this.touches[0] = this.touches[1];
            this.touches[1] = {
                time: +new Date,
                pageY: ex.touches[0].pageY
            };
            this.setPosition();
        },
        onTouchEnd(ex) {
            this.setEasePosition();
        },
        setPosition() {
            let diff = this.touches[1].pageY - this.touches[0].pageY;
            this.diffY += diff;
            this.$el.style.transition = 'transform 0';
            this.$el.style.transform = `translate3d(0,${this.diffY}px,0)`;
        },
        setEasePosition() {
            let diff = this.touches[1].pageY - this.touches[0].pageY;   // 最后时刻的位移
            let time = this.touches[1].time - this.touches[0].time;     // 位移花费的时间
            let speed = diff / time; // 速度
            const downSpeed = 2;    // 减速度
            let timeSpan = Math.abs(~~(speed / downSpeed));  // 缓动耗费时间
            let distance = ~~(speed / 2) * timeSpan; // 缓动移动距离， 平均速度*时间
            this.$el.style.transition = `transform ${timeSpan}ms`;
            this.$el.style.transform = `translate3d(0,${0}px,0)`;
        }
    }
}
</script>

<style lang="scss" scoped>
.select-list {
    list-style: none;
    margin: 0;
    padding: 0;
    position: relative;
    width: 100%;
    height: 170px;
    box-sizing: border-box;

    &-item {
        box-sizing: border-box;
        font-size: 18px;
        height: 34px;
        line-height: 34px;
        text-align: center;
    }
}
</style>