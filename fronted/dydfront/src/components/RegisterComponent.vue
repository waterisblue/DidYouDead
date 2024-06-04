<template>
    <div class="registerWindow">
        <LoginInputComponent @send-input-data="(username) => user.username = username" :thisRules="accountRules" show-password="false"
            inputText="账号" />
        <LoginInputComponent @send-input-data="(password) => user.password = password" :thisRules="passwordRules"
            inputText="密码" />
        <LoginInputComponent @send-input-data="(password) => user.secondPassword = password" :thisRules="passwordRules"
            inputText="确认密码" />
        <div class="btns">
            <div class="registerBtn">
                <LoginBtnComponent btnText="注册" @click="registerFun()" />
            </div>
            <div class="registerBtn">
                <LoginBtnComponent btn-color="info" btnText="已有账号，点击登录" @click="router.push('/')" />
            </div>
        </div>
        <v-snackbar :timeout="2000" color="error" v-model="snackbar">
            {{ snackText }}

            <template v-slot:actions>
                <v-btn color="pink" variant="text" @click="snackbar = false">
                    Close
                </v-btn>
            </template>
        </v-snackbar>
    </div>
</template>

<script setup>
import LoginInputComponent from '../components/LoginInputComponent.vue'
import LoginBtnComponent from '../components/LoginBtnComponent.vue'
import { useRouter } from 'vue-router';
import axios from '../axios';
import { ref } from 'vue';

const router = useRouter()

let snackbar = ref(false)
let snackText = ref('')

let user = {
    username: '',
    password: ''
}

let accountRules = ref([
    value => (value && value.length >= 7) || '账号必须大于七位！',
    value => value.length <= 20 || '账号必须小于二十位！'
])

let passwordRules = ref([
    value => (value && value.length >= 7) || '密码必须大于七位！',
    value => value.length <= 20 || '密码必须小于二十位！'
])

function registerFun() {
    if (user.username.length > 20 || user.username.length < 7 || user.password.length > 20 || user.password.length < 7) {
        snackText.value = '输入的账号或密码格式有误'
        snackbar.value = true
        return
    }

    if (user.password != user.secondPassword) {
        snackText.value = '两次输入的密码不同'
        snackbar.value = true
        return
    }
    axios.post('/registeruser', user).then(res => {
        if(res.data.code != 200){
            snackText.value = res.data.msg
            snackbar.value = true
            return
        }
        router.push('/')
    }).catch(err => {
        snackText.value = '注册失败，出现错误'
        snackbar.value = true
    })
}
</script>


<style>
.registerWindow {
    min-height: 60vh;
    border-radius: 1rem;

    background-color: rgba(215, 231, 231, 0.8);


    padding: 1rem;
    display: flex;
    flex-direction: column;
    justify-content: space-around;

    .btns {
        display: flex;
        justify-content: space-around;
    }
}
</style>