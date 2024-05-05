<template>
    <v-container class="techOuter" fluid>
        <v-toolbar flat>
            <v-toolbar-title>智能服务</v-toolbar-title>
            <v-btn @click="checkUrl('/bubble')">返回菜单</v-btn>
        </v-toolbar>
        <div class="techMain">
            <TechUpperChipComponent />
            <v-divider></v-divider>
            <TechServiceChartsComponent/>
        </div>
    </v-container>
</template>

<script setup>
import TechUpperChipComponent from '@/components/TechUpperChipComponent.vue'
import TechServiceChartsComponent from '@/components/TechServiceChartsComponent.vue'
import { useRouter } from 'vue-router';

const router = useRouter()

function checkUrl(path) {
    router.push(path)
}
</script>

<style lang="less" scoped>
.techOuter {
    height: 100vh;
    .techMain {
        height: 90vh;
        display: flex;
        flex-direction: column;
    }
}
</style>