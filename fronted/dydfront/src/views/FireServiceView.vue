<template>
    <v-container fluid class="outScreen">
        <div class="toolBar">
            <v-btn @click="checkUrl('bubble')">&lt; 返回</v-btn>
        </div>

        <div class="timeLine">
            <FireServiceTimeLineComponent />
            <div>
                123
            </div>
        </div>
    </v-container>
</template>

<script setup>
import FireServiceTimeLineComponent from '@/components/FireServiceTimeLineComponent.vue'
import { useRouter } from 'vue-router';

const router = useRouter()

function checkUrl(path){
    router.push(path)
}
</script>

<style scoped lang="less">
.outScreen {
    background-color: aliceblue;
    min-height: 100vh;
    display: flex;
    flex-direction: column;
    .toolBar {
        flex: 1;
        display: flex;
        flex-direction: row;
        align-items: center;
        justify-content: start;
    }
    .timeLine {
        flex: 50;
    }
}
</style>