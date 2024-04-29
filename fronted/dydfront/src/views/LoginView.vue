<template>
    <v-container fluid align="center" class="fullScreen">
        <div class="loginScreen">
            <div class="title" align="right">
                摆渡人 - 你的灵魂归宿
            </div>
            <transition name="fullLogin" appear>
                <RouterView />
            </transition>
            <div class="logoWindow">logo</div>
        </div>
    </v-container>
</template>

<script setup>
let name = 'LoginView'
</script>

<style scoped lang="less">
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
    transition: all 3s;
}

.fullScreen {
    min-height: 100vh;
    display: flex;
    flex-direction: row;
    justify-content: center;
    background-image: url('../assets/loginBackground.png');

    .loginScreen {
        width: 120vh;
        display: flex;
        flex-direction: column;
        justify-content: space-around;
    }

    .title {
        font-size: 3rem;
        color: white;
        font-weight: 900;
    }

    .logoWindow {
        min-height: 20vh;
        background-color: aliceblue;
        border-radius: 1rem;
    }
}
</style>
