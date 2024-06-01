<template>
      <v-layout class="rounded rounded-md">
        <v-app-bar title="DidYouDead">
            <v-btn @click="checkUrl('/bubble')">返回菜单</v-btn>
        </v-app-bar>

        <v-navigation-drawer>
            <v-list density="compact" nav>
                <v-list-item prepend-icon="mdi-account" title="用户管理" value="index" @click="checkUrl('/super')"></v-list-item>
                <v-list-item prepend-icon="mdi-view-dashboard" title="心率检测" value="TestamentDesign" @click="checkUrl('/super/ratescan')"></v-list-item>
                <v-list-item prepend-icon="mdi-home" title="供应商资源管理" value="MyTestament" @click="checkUrl('/super/supply')"></v-list-item>
            </v-list>
        </v-navigation-drawer>

        <v-main class="d-flex align-center justify-center main">
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
import { useRouter } from 'vue-router';

const router = useRouter()
// 用于vue-router
function checkUrl(path) {
    router.push(path)
}

</script>

