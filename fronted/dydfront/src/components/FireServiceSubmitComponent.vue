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
                <div>预约时间: {{convertTime(service.OrderTime)}}</div>
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
            <v-btn class="submitBtn" color="info" @click="submitFireService" :active="submitActive">提交 合计{{ amount }}元</v-btn>
        </div>
        <SnackBarComponent :snackText="snackText" :snackbar="snackBar" snackColor="waring" />
    </v-container>
</template>

<script setup>
import { useTimeLineStore, useFireServiceChoose } from '@/stores/user';
import axios from '../axios';
import SnackBarComponent from '@/components/SnackBarComponent.vue'
import { useRouter } from 'vue-router';
import { ref } from 'vue';

const { setTimeLineInfo} = useTimeLineStore()
const { getService } = useFireServiceChoose()




let service = getService()
setTimeLineInfo('six')

const router = useRouter()

function checkUrl(path){
  router.push(path)
}
// snackBar
let snackText = ref('')
let snackBar = ref(false)
let amount = ref(-1)
let submitActive = ref(false)
service = getService()

console.log(service)
axios.post('/loginafter/calAmount', {
    "funeralParlorId": service.FuneralParlorId,
    "fireServiceId": service.FireServiceId,
    "urnStyleId": service.UrnStyleId,
    "cemeteryId": service.CemeteryId
}).then(res => {
    amount.value = res.data.data.amount
    submitActive.value = true
})


function submitFireService(){
    
    let thatSnackBar = snackBar
    axios.post('/loginafter/uploadFireService', {
        Age: parseInt(service.Age),
        Cemetery: service.Cemetery,
        FireService: service.FireService,
        FuneralParlor: service.FuneralParlor,
        IdNum: service.IdNum,
        Locate: service.Locate,
        Name: service.Name,
        OrderTime: service.OrderTime,
        PhoneNum: service.PhoneNum,
        Sex: service.Sex,
        UrnStyle: service.UrnStyle,
        Name: service.Name,
        Amount: amount.value
    }).then(res => {
        if(res.data.code == 200){
            thatSnackBar.value = true
            snackText.value = "火化方式上传成功！"
            checkUrl('/fireService?exist=true')
            // setTimeout(() => {}, 1500)
        }
    })
}


function convertTime(time){
    if (!time) {
        return ""
    }
    return time.toLocaleString()
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