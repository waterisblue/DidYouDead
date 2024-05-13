<template>
    <v-container fluid class="outScreen">
        <div class="toolBar">
            <v-btn  icon class="backBtn" @click="checkUrl('/bubble')">&lt; </v-btn>
        </div>

        <div class="timeLine" v-if="!fireServiceExist">
            <FireServiceTimeLineComponent />
            <RouterView />
        </div>
        <FireServiceShowComponent v-if="fireServiceExist" />
    </v-container>
</template>

<script setup>
import FireServiceTimeLineComponent from '@/components/FireServiceTimeLineComponent.vue'
import FireServiceShowComponent from '@/components/FireServiceShowComponent.vue'

import { useRouter } from 'vue-router'
import { ref } from 'vue';

const router = useRouter()

function checkUrl(path){
    router.push(path)
}

let fireServiceExist = ref(false)

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