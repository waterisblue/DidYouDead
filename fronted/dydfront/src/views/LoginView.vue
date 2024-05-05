<template>
    <video class="videoBackground" playsinline="playsinline" autoplay="autoplay" muted="muted">
        <source src="../assets/loginVideo.mp4" type="video/mp4">
    </video>
    <v-container fluid align="center" class="fullScreen">
        <div class="loginScreen">
            <div class="title" align="right">
                摆渡人 - 你的灵魂归宿
            </div>
            <RouterView />
            <div class="logoWindow"></div>
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
    transition: all 13s;
}

.fullScreen {
    min-height: 100vh;
    display: flex;
    flex-direction: row;
    justify-content: center;

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
        border-radius: 1rem;
        background-image: url('../assets/1T83cq-LogoMakr.png');
        background-repeat: no-repeat;
        background-position: center center;

    }
}

.videoBackground {
    position: fixed;
    /* 或者使用 absolute，视你的布局而定 */
    right: 0;
    bottom: 0;
    min-width: 100%;
    min-height: 100%;
    z-index: -1000;
    /* 确保视频在内容之下 */
    background-size: cover;
    float: left;
}
</style>
