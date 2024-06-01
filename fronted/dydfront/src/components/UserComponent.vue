<template>
    <v-container class="outerTestament">
        <v-data-table-virtual :headers="userHeaders" :items="user">
            <template v-slot:item.Administrator="{ item }">
                <v-checkbox v-model="item.Administrator" @click="grantUser(item.Username, item.Administrator)"></v-checkbox>
            </template>
            <template v-slot:item.delete="{ item }">
                <v-btn color="error" @click="deleteUser(item.Username)">删除</v-btn>
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

let userHeaders = ref([
    {
        title: '用户名',
        align: 'start',
        sortable: false,
        key: 'Username',
    },
    { title: '是否管理员', key: 'Administrator' },
    { title: '操作', key: 'delete' },
])
let user = ref([])
axios.post('/loginafter/getalluser').then(res => {
    user.value = res.data.data
})

function deleteUser(account) {
    axios.post('/loginafter/deleteuser', {"username": account})
    for(var i = 0; i < user.value.length; i++){
        if(user.value[i].Username == account) {
            user.value.splice(i, 1)
        }
    }
    snackText.value = account + '删除成功'
    snackBar.value = true
    
}

function grantUser(account, grant) {
    axios.post('/loginafter/grant', {
        "username": account,
        "administrator": grant
    })
    snackText.value = account + '权限更改成功'
    snackBar.value = true
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