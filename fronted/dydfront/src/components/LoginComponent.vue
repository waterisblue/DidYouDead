<template>
    <transition name="fullLogin" appear>
        <div class="loginWindow">
            <LoginInputComponent inputText="账号" />
            <LoginInputComponent inputText="密码" />
            <div class="btns">
                <div class="loginBtn">
                    <LoginBtnComponent btnText="登录" btnColor="indigo-darken-3" @click="jump('/bubble')" />
                </div>
                <div class="registerBtn" @click="jump('/register')">
                    <LoginBtnComponent btnText="注册" />
                </div>
            </div>
        </div>
    </transition>

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


<style scoped>
.loginWindow {
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

.fullLogin-enter-from,
.fullLogin-appear-from,
.fullLogin-leave-to {
    opacity: 0;
}

/* 过渡结束状态 */
.fullLogin-enter-to,
.fullLogin-appear-to {
    opacity: 1;
}

/* 过渡激活状态 */
.fullLogin-enter-active,
.fullLogin-appear-active,
.fullLogin-leave-active {
    transition: all 13s;
}
</style>