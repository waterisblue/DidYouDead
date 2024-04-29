<template>
    <div class="toolBar">
        <v-btn @click="checkUrl('/')">注销</v-btn>
    </div>
    <v-container class="bubbleContainer" fluid>
        <div class="bubbleRow">
            <BubbleComponent innerText="殡仪馆及火化服务" innerId="fireservice"/>
            <BubbleComponent innerText="遗嘱存储" innerId="testament" />
        </div>
        <div class="bubbleRow">
            <BubbleComponent innerText="智能技术服务" innerId="bubble3" />
            <BubbleComponent innerText="客服中心" innerId="bubble4" />
        </div>
    </v-container>
</template>

<script setup>
import BubbleComponent from '@/components/BubbleComponent.vue'
import { useRouter } from 'vue-router';

const router = useRouter()

function checkUrl(path){
    router.push(path)
}
</script>

<style scoped>
.toolBar {
    background-color: aqua;
    height: 5vh;
    display: flex;
    align-items: center;
    justify-content: end;
}
.bubbleContainer {
    min-height: 95vh;
    display: flex;
    flex-direction: column;
    justify-content: space-around;
    background-color: aqua;

    .bubbleRow:nth-child(1){
        flex: 5;
        display: flex;
        flex-direction: row;
        justify-content: space-around;
    }

    .bubbleRow:nth-child(2){
        flex: 5;
        display: flex;
        flex-direction: row;
        justify-content: center;
    }
}
</style>