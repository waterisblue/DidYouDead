<template>
    <video class="videoBackground" playsinline="playsinline" autoplay="autoplay" muted="muted">
        <source src="../assets/indexVideo.mp4" type="video/mp4">
    </video>
    <div class="toolBar">
        <button @click="checkUrl('/')" class="logoutBtn">注销</button>
    </div>
    <v-container class="bubbleContainer" fluid>
        <div class="bubbleRow">
            <BubbleComponent innerText="殡仪馆及火化服务" innerId="fireservice" />
            <BubbleComponent innerText="遗嘱存储" innerId="testament" />
        </div>
        <div class="bubbleRow">
            <BubbleComponent innerText="智能技术服务" innerId="tech" />
            <BubbleComponent v-if="!authority" innerText="客服中心" @click="checkSnackBar()" innerId="bubble" />
            <BubbleComponent v-if="authority" innerText="管理中心" innerId="supervisor" />
        </div>
        <v-snackbar v-model="snackbarHide" vertical>
            <div class="text-subtitle-1 pb-2">欢迎访问DidYouDead客服服务！</div>

            <p>疑问请拨打&nbsp;150-4092-7641&nbsp;咨询详情！</p>

            <template v-slot:actions>
                <v-btn color="indigo" variant="text" @click="snackbarHide = false">
                    Close
                </v-btn>
            </template>
        </v-snackbar>
    </v-container>
</template>

<script setup>
import BubbleComponent from '@/components/BubbleComponent.vue'
import { ref } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter()

function checkUrl(path) {
    router.push(path)
}

// 是否管理员
let authority = false

let snackbarHide = ref(false)

function checkSnackBar(){
    snackbarHide.value = true
}
</script>

<style scoped>
.toolBar {
    height: 5vh;
    display: flex;
    align-items: center;
    justify-content: end;

    .logoutBtn {
        position: absolute;
        right: -3rem;
        top: -4rem;
        width: 8rem;
        height: 8rem;
        border-radius: 50%;
        background-color: transparent;
        border: solid 0.1rem rgba(173, 189, 202, 0.529);

        color: rgba(180, 202, 221, 0.814);
        font-size: large;
        display: flex;
        justify-content: left;
        align-items: end;
        padding: 1.6rem;

        transition: all 1s;
    }

    .logoutBtn:hover {
        border: solid 0.1rem rgb(173, 189, 202);
        background-color: rgba(240, 248, 255, 0.121);
    }
}

.bubbleContainer {
    min-height: 95vh;
    display: flex;
    flex-direction: column;
    justify-content: space-around;

    .bubbleRow:nth-child(1) {
        flex: 5;
        display: flex;
        flex-direction: row;
        justify-content: space-around;
    }

    .bubbleRow:nth-child(2) {
        flex: 5;
        display: flex;
        flex-direction: row;
        justify-content: center;
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