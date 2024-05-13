<template>
    <v-container class="outer">
        <div class="detail">
            <div class="detailone">
                <div>姓名: {{ service.Name }}</div>
                <div>性别: {{ service.Sex }}</div>
                <div>身份证号: {{service.IdNum}}</div>
            </div>
            <div class="detailtwo">
                <div>家庭住址: {{service.Locate}}</div>
                <div>手机号: {{service.PhoneNum}}</div>
                <div>预约时间: {{service.OrderTime}}</div>
            </div>
        </div>
        <v-divider></v-divider>
        <div class="selectDetail">
            <div>选择殡仪馆：{{service.FuneralParlor}}</div>
            <div>火化方式：{{service.FireService}}</div>
            <div>骨灰盒样式：{{service.UrnStyle}}</div>
            <div>墓地位置：{{service.Cemetery}}</div>
        </div>
        <v-divider></v-divider>
        <div class="submitBtns">
            <v-btn class="submitBtn" color="error" @click="checkUrl('/bubble')">取消</v-btn>
            <v-btn class="submitBtn" color="info" @click="checkUrl('/bubble')">提交</v-btn>
        </div>
    </v-container>
</template>

<script setup>

import { useTimeLineStore, useFireServiceChoose } from '@/stores/user';
import { storeToRefs } from 'pinia';
import { useRouter } from 'vue-router';
const {timeLineStore, setTimeLineInfo} = useTimeLineStore()
const { getService, setService } = useFireServiceChoose()
let service = getService()
console.log(service)
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