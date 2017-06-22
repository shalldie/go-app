<template>
    <div class="select-confirm">
        <select-list :list="list"></select-list>
    </div>
</template>

<script>
import SelectList from './SelectList';

let list = Array.from({ length: 10 }, (item, index) => index.toString().repeat(10));

export default {
    data() {
        return { list };
    },
    computed: {

    },
    components: {
        SelectList
    }
}
</script>

<style lang="scss">
.select-confirm {
    position: fixed;
    left: 0;
    bottom: 0;
    width: 100%;
    height: 170px;
    border-top: 1px solid #ddd;
    .col-4 {
        float: left;
        width: 33.333%;
        height: 100%;
    }
}
</style>