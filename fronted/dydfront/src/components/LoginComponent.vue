<template>
    <div class="loginWindow">
        <LoginInputComponent inputText="账号" />
        <LoginInputComponent inputText="密码" />
        <div class="btns">
            <div class="loginBtn">
                <LoginBtnComponent btnText="登录" btnColor="indigo-darken-3" @click="jump('/bubble')" />
            </div>
            <div class="registerBtn" @click="jump('/register')">
                <LoginBtnComponent btnText="注册"/>
            </div>
        </div>
    </div>
</template>

<script setup>
import LoginInputComponent from '../components/LoginInputComponent.vue'
import LoginBtnComponent from '../components/LoginBtnComponent.vue'

import { useRouter } from 'vue-router';

const router = useRouter()

function jump(path) {
    router.push(path)
}
</script>


<style>
.loginWindow {
    min-height: 60vh;
    border-radius: 1rem;

    background-color: rgba(215, 231, 231, 0.384);


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