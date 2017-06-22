<template>
    <div id="app">
        <select-confirm></select-confirm>
    </div>
</template>

<script>
// import InfoList from './components/InfoList';
// import SelectList from './components/SelectList';
import SelectConfirm from './components/SelectConfirm';

let arr = Array.apply(null, Array(30)).map((item, index) => (index + 10).toString().repeat(4));

export default {
    name: 'app',
    data() {
        return {
            list: arr
        };
    },
    components: {
        SelectConfirm
    }
}
</script>

<style lang="scss">
#app {
    font-family: 'Avenir', Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: center;
    color: #2c3e50;
    height: 500px;
    width: 100%;
    position: fixed;
    left: 0;
    bottom: 0;

    .col-4 {
        width: 33.33%;
        float: left;
    }
}
</style>
