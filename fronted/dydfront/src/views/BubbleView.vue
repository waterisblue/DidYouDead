<template>
    <v-container class="bubbleContainer" fluid>
        <div class="bubbleRow">
            <BubbleComponent innerText="殡仪馆及火化服务" innerId="bubble1"/>
            <BubbleComponent innerText="遗嘱存储" innerId="bubble2" />
        </div>
        <div class="bubbleRow">
            <BubbleComponent innerText="智能技术服务" innerId="bubble3" />
            <BubbleComponent innerText="客服中心" innerId="bubble4" />
        </div>
    </v-container>
</template>

<script setup>
import BubbleComponent from '@/components/BubbleComponent.vue'
</script>

<style scoped>
.bubbleContainer {
    min-height: 100vh;
    display: flex;
    flex-direction: column;
    justify-content: space-around;
    background-color: aqua;

    .bubbleRow:nth-child(1){
        display: flex;
        flex-direction: row;
        justify-content: space-around;
    }

    .bubbleRow:nth-child(2){
        display: flex;
        flex-direction: row;
        justify-content: center;
    }
}
</style>