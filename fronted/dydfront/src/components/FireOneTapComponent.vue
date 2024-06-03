<template>
    <div class="outer">
        <OneTabCardComponent v-for="item in items" :tap-id="item.id" :tap-image="staticENV + item.imgurl" :tap-name="item.sourcename" :tap-sub-title="item.subtitle" :tap-detail="item.sourcedetail" />
    </div>
</template>

<script setup>
import OneTabCardComponent from '@/components/OneTapCardComponent.vue'

import { useTimeLineStore } from '@/stores/user';
import axios from '../axios';
import { ref } from 'vue';
let staticENV = ref(import.meta.env.VITE_STATIC_URL + '/imgfile/')
let items = ref([])

axios.post('/loginafter/getsupplybytype?type=1').then(res => {
    items.value = res.data.data
})


const { setTimeLineInfo} = useTimeLineStore()
setTimeLineInfo('one')
</script>

<style lang="less" scoped>
.outer {
    height: 75vh;
    display: flex;
    flex-direction: row;;
    justify-content: space-around;
}
</style>