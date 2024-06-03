<template>

    <v-container class="outerTestament">
        <v-btn color="warning" style="margin-bottom: 3rem;" @click="router.push('/super/supply/add')">添加资源</v-btn>

        <v-data-table-virtual :headers="supplyHeaders" :items="supply">
            <template v-slot:item.imgurl="{ item }">
                <v-img :src="staticENV + item.imgurl"></v-img>
            </template>
            <template v-slot:item.type="{ item }">
                {{ CheckType(item.type) }}
            </template>
            <template v-slot:item.delete="{ item }">
                <v-btn color="error" @click="DeleteSupply(item)">删除</v-btn>
            </template>
        </v-data-table-virtual>
        <SnackBarComponent :snackText="snackText" :snackbar="snackBar" snackColor="warning" />
    </v-container>

</template>
<script setup>
import axios from '../axios';
import { ref } from 'vue';
import SnackBarComponent from './SnackBarComponent.vue';
import router from '@/router';

// snackBar
let snackText = ref('')
let snackBar = ref(false)


let staticENV = ref(import.meta.env.VITE_STATIC_URL + '/imgfile/')

let supplyHeaders = ref([
    {
        title: '',
        align: 'start',
        sortable: false,
        key: 'imgurl',
    },
    { title: '资源名称', key: 'sourcename' },
    { title: '资源简介', key: 'subtitle' },
    { title: '资源类型', key: 'type' },
    { title: '资源价格', key: 'price' },
    { title: '操作', key: 'delete' },
])
let supply = ref([])
axios.post('/loginafter/getsupply').then(res => {
    console.log(res.data.data)
    supply.value = res.data.data
})

function DeleteSupply(item){
    axios.post('/loginafter/deletesupply', {"id": item.id})
    for(var i = 0; i < supply.value.length; i++){
        if(supply.value[i].id == item.id) {
            supply.value.splice(i, 1)
        }
    }
    snackText.value = account + '删除成功'
    snackBar.value = true
}

function CheckType(type) {
    switch (type) {
        case 1:
            return "殡仪馆"
        case 2:
            return "墓地"
        case 3:
            return "火化方式"
        case 4:
            return "骨灰盒"
        default:
            return "未知"
    }
}

</script>


<style lang="less" scoped>
.outerTestament {
    display: flex;
    flex-direction: column;
    justify-content: start;
    margin-top: 10rem;
    height: 100vh;
}

.btns {
    display: flex;
    flex-direction: row;
    justify-content: right;

    .btn {
        margin: 1rem;
    }
}
</style>