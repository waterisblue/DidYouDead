<template>
    <div class="outer">
        <OneTabCardComponent />
        <OneTabCardComponent />
        <OneTabCardComponent />
        <OneTabCardComponent />
        <OneTabCardComponent />
    </div>
</template>

<script setup>
import OneTabCardComponent from '@/components/OneTapCardComponent.vue'
</script>

<style lang="less" scoped>
.outer {
    height: 75vh;
    display: flex;
    flex-direction: row;
    justify-content: space-around;
}
</style>