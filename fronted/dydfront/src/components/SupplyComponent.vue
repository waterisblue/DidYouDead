<template>
    <v-container class="outerTestament">
        <v-data-table-virtual :headers="supplyHeaders" :items="supply">
            <template v-slot:item.imgurl="{ item }">
                <v-img :src="item.imgurl"></v-img>
            </template>
            <template v-slot:item.type="{ item }">
                {{ CheckType(item.type) }}
            </template>
            <template v-slot:item.delete="{ item }">
                <v-btn color="error" @click="deleteUser(item)">删除</v-btn>
            </template>
        </v-data-table-virtual>
        
        <SnackBarComponent :snackText="snackText" :snackbar="snackBar" snackColor="warning" />
    </v-container>

</template>
<script setup>
import axios from '../axios';
import { ref } from 'vue';
import SnackBarComponent from './SnackBarComponent.vue';

// snackBar
let snackText = ref('')
let snackBar = ref(false)

let supplyHeaders = ref([
    {
        title: '',
        align: 'start',
        sortable: false,
        key: 'imgurl',
    },
    { title: '资源名称', key: 'sourcename' },
    { title: '资源描述', key: 'sourcedetail' },
    { title: '资源类型', key: 'type' },
    { title: '资源价格', key: 'price' },
    { title: '操作', key: 'delete' },
])
let supply = ref([])
axios.post('/loginafter/getsupply').then(res => {
    console.log(res.data.data)
    supply.value = res.data.data
})

function DeleteSupply(){
    
}

function CheckType(type) {
    switch (type) {
        case 1:
            return "殡仪馆"
            break;
        case 2:
            return "墓地"
        default:
            return "其他"
            break;
    }
}

</script>


<style lang="less" scoped>
.outerTestament {
    display: flex;
    flex-direction: column;
    justify-content: center;
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