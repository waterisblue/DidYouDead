<template>
    <v-container fluid class="outScreen">
        <div class="toolBar">
            <v-btn icon class="backBtn" @click="checkUrl('/bubble')">&lt; </v-btn>
        </div>

        <div class="timeLine" v-if="!fireServiceExist">
            <FireServiceTimeLineComponent />
            <RouterView />
        </div>
        <FireServiceShowComponent v-if="fireServiceExist" :fire-service="fireService"
            @change-plan="(exist) => { fireServiceExist = exist }" @pay-service="payService" />
    </v-container>
</template>

<script setup>
import FireServiceTimeLineComponent from '@/components/FireServiceTimeLineComponent.vue'
import FireServiceShowComponent from '@/components/FireServiceShowComponent.vue'

import { useRouter } from 'vue-router'
import { ref } from 'vue';
import axios from '../axios';
import { onMounted } from 'vue';
const router = useRouter()

let fireServiceExist = ref(false)
let fireService = ref(null)


onMounted(() => {
    // 监听路由变化
    const updateQueryParamX = () => {
        let exist = router.currentRoute.value.query.exist

        if (exist != null) {
            location.href = window.location.origin + window.location.pathname;
        }
    };

    // 初始化时读取一次
    updateQueryParamX();

    // 监听路由变化
    router.afterEach(updateQueryParamX);
})
axios.post('/loginafter/getFireServiceByUserName').then(res => {
    if (res.data.code == 200) {
        fireService.value = res.data.data.fireService
        fireServiceExist.value = true
    }
})



function payService(id) {
    let backUrl = import.meta.env.VITE_BACK_URL
    window.open(`${backUrl}/pay?amount=${fireService.value.Amount}.00&subject="智能设备产品支付"&id=${id}&typeId=0`, '_blank');
}

function checkUrl(path) {
    router.push(path)
}
</script>

<style scoped lang="less">
.outScreen {
    background-color: aliceblue;
    background-image: url();
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