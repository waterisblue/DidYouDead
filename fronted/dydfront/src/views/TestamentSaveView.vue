<template>
    <v-layout class="rounded rounded-md">
        <v-app-bar title="DidYouDead">
            <v-btn @click="checkUrl('/bubble')">返回菜单</v-btn>
        </v-app-bar>

        <v-navigation-drawer>
            <v-list density="compact" nav>
                <v-list-item prepend-icon="mdi-view-dashboard" title="首页" value="index" @click="checkUrl('/testament')"></v-list-item>
                <v-list-item prepend-icon="mdi-forum" title="遗嘱制定" value="TestamentDesign" @click="checkUrl('/testament/design')"></v-list-item>
                <v-list-item prepend-icon="mdi-forum" title="我的遗嘱" value="MyTestament" @click="checkUrl('/testament/my')"></v-list-item>
            </v-list>
        </v-navigation-drawer>

        <v-main class="d-flex align-center justify-center main">
            <span v-if="indexCheck">建立一个属于你的遗嘱</span>
            <RouterView />
        </v-main>
    </v-layout>
</template>

<style scoped>
.main {
    height: 100vh;
}
</style>

<script setup>
import { ref } from 'vue';
import { useRoute, useRouter } from 'vue-router';

const router = useRouter()
let indexCheck = ref(false)

const route = useRoute()
if(route.path === '/testament'){
    indexCheck.value = true
}
// 用于vue-router
function checkUrl(path) {
    if(path != '/testament')
        indexCheck.value = false
    else
        indexCheck.value = true
    
    router.push(path)
}

</script>