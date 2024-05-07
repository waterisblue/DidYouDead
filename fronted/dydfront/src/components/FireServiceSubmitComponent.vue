<template>
    <v-container class="outer">
        <div class="detail">
            <div class="detailone">
                <div>姓名: 张三</div>
                <div>性别: 男</div>
                <div>身份证号: 130xxxxxxxxxxx1234</div>
            </div>
            <div class="detailtwo">
                <div>家庭住址: xx</div>
                <div>手机号: 1402131xxxx</div>
                <div>预约时间: 暂空</div>
            </div>
        </div>
        <v-divider></v-divider>
        <div class="selectDetail">
            <div>选择殡仪馆：天堂之翼殡仪馆</div>
            <div>火化方式：无感高温1000C</div>
            <div>骨灰盒样式：蓝色迷笛</div>
            <div>目的位置：xxxxx</div>
        </div>
        <v-divider></v-divider>
        <div class="submitBtns">
            <v-btn class="submitBtn" color="error" @click="checkUrl('/bubble')">取消</v-btn>
            <v-btn class="submitBtn" color="info" @click="checkUrl('/bubble')">提交</v-btn>
        </div>
    </v-container>
</template>

<script setup>

import { useTimeLineStore } from '@/stores/user';
import { storeToRefs } from 'pinia';
import { useRouter } from 'vue-router';
const {timeLineStore, setTimeLineInfo} = useTimeLineStore()

const {timeLineInfo} = storeToRefs(timeLineStore)
setTimeLineInfo('six')

const router = useRouter()

function checkUrl(path){
  router.push(path)
}
</script>

<style scoped lang="less">
.outer {
    display: flex;
    flex-direction: column;
    justify-content: space-around;
    height: 69vh;
    font-size: large;

    .detail {
        margin-top: 1rem;
        flex: 1;
        display: flex;
        flex-direction: column;
        justify-content: space-around;

        .detailone {
            flex: 1;
            display: flex;
            flex-direction: row;
            justify-content: space-between;
        }
        .detailtwo {
            flex: 1;
            display: flex;
            flex-direction: row;
            justify-content: space-between;
        }
    }
    .selectDetail {
        flex: 2;
        display: flex;
        flex-direction: column;
        justify-content: space-around;

    }
    .submitBtns {
        flex: 1;
        display: flex;
        flex-direction: row;
        justify-content: space-around;
        align-items: center;

        .submitBtn {
            width: 20rem;
        }
    }
}
</style>