<template>
    <transition name="fullLogin" appear>
        <div class="loginWindow">
            <LoginInputComponent @send-input-data="(username) => user.username = username" inputText="账号" show-password=true />
            <LoginInputComponent @send-input-data="(password) => user.password = password" inputText="密码" /> 
                <div class="btns">
                <div class="loginBtn">
                    <LoginBtnComponent btnText="登录" btnColor="indigo-darken-3" @click="loginFunc" />
                </div>
                <div class="registerBtn" @click="router.push('/register')">
                    <LoginBtnComponent btnText="注册" />
                </div>
            </div>
        </div>
    </transition>
    <v-snackbar :timeout="2000" color="error" v-model="snackbar">
            {{ snackText }}

            <template v-slot:actions>
                <v-btn color="pink" variant="text" @click="snackbar = false">
                    Close
                </v-btn>
            </template>
        </v-snackbar>
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

var user = {
    username: '',
    password: ''
}
function loginFunc(path) {
    axios.post('/grantAuth', user).then(res => {
        if(res.data.code != 200){
            snackText.value = res.data.msg
            snackbar.value = true
            return
        }
        
        localStorage.setItem('token', res.data.data.token)
        router.push('/bubble')
    }).catch(err => {
        snackText.value = res.data.msg
        snackbar.value = true
        return
    })
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