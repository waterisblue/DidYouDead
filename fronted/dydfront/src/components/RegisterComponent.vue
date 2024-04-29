<template>
    <div class="registerWindow">
        <LoginInputComponent inputText="账号" />
        <LoginInputComponent inputText="密码" />
        <LoginInputComponent inputText="确认密码" />
        <div class="btns">
            <div class="registerBtn">
                <LoginBtnComponent btnText="注册" @click="registerFun()" />
            </div>
        </div>
    </div>
</template>

<script setup>
import LoginInputComponent from '../components/LoginInputComponent.vue'
import LoginBtnComponent from '../components/LoginBtnComponent.vue'
import { useRouter } from 'vue-router';

const router = useRouter()

function registerFun(){
    router.push('/')
}
</script>


<style>
.registerWindow {
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