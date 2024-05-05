<template>
    <v-container fluid class="outScreen">
        <div class="toolBar">
            <v-btn  icon class="backBtn" @click="checkUrl('bubble')">&lt; </v-btn>
        </div>

        <div class="timeLine">
            <FireServiceTimeLineComponent />
            <!-- <FireOneTapComponent /> -->
            <FireTwoTapComponent />
        </div>
    </v-container>
</template>

<script setup>
import FireServiceTimeLineComponent from '@/components/FireServiceTimeLineComponent.vue'
import FireOneTapComponent from '@/components/FireOneTapComponent.vue'
import FireTwoTapComponent from '@/components/FireTwoTapComponent.vue'

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
    .backBtn {
        font-size: 1.5rem;
        color: rgb(133, 154, 173);
        display: flex;
        justify-content: center;
        align-items: center;
    }
}
</style>