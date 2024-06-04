<template>
   <p>===除最后一行服务项目外，其他项目均不额外收费===</p>
    <!--  <div class="outer">
        
        <CemeteryLocateTapComponent  tap-image="/static/mu8.jpg" tap-name="傍山谧园" tap-sub-title="免费" tap-detail=""/>
        <CemeteryLocateTapComponent tap-image="/static/mu12.jpg"  tap-name="传统园区" tap-sub-title="免费" tap-detail="" />
        <CemeteryLocateTapComponent tap-image="/static/mu9.jpg" tap-name="绿色伴生园" tap-sub-title="免费" tap-detail=""/>
        <CemeteryLocateTapComponent tap-image="/static/mu13jpg.jpg" tap-name="诗意长眠园" tap-sub-title="免费" tap-detail=""/>
    </div>
    <p>===除最后一行服务项目外，其他项目均不额外收费===</p>
    <div class="outer">
        
        <CemeteryLocateTapComponent  tap-image="/static/mu11jpg.jpg" tap-name="肃穆简约区" tap-sub-title="免费" tap-detail=""/>
        <CemeteryLocateTapComponent tap-image="/static/mu1.jpg"  tap-name="富丽堂皇区" tap-sub-title="免费" tap-detail="" />
        <CemeteryLocateTapComponent tap-image="/static/mu4.jpg" tap-name="新中式园林区" tap-sub-title="免费" tap-detail=""/>
        <CemeteryLocateTapComponent tap-image="/static/mu14.jpg" tap-name="ins纪念区" tap-sub-title="免费" tap-detail=""/>
    </div>
    <p>+++附加服务区+++</p>
    <div class="outer">
        <CemeteryLocateTapComponent  tap-image="/static/hai3.jpg" tap-name="员工集中海葬 " tap-sub-title="5999元" tap-detail=""  />
        <CemeteryLocateTapComponent  tap-image="/static/hai.jpg" tap-name="家属单独海葬 " tap-sub-title="6999元" tap-detail=""/>
        <CemeteryLocateTapComponent tap-image="/static/hai2.jpg"  tap-name="海底墓园 " tap-sub-title="89999元" tap-detail="" />
        <CemeteryLocateTapComponent tap-image="/static/tai3.jpg" tap-name="太空葬 " tap-sub-title="139999元" tap-detail=""/>
    </div> -->
    <div v-for="(chunk, index) in cemetryItems" :key="index" class="outer">
        <CemeteryLocateTapComponent v-for="item in chunk" :key="item.id" :tap-id="item.id" :tap-image="staticENV + item.imgurl" :tap-name="item.sourcename" :tap-sub-title="item.subtitle" :tap-detail="item.sourcedetail" />
    </div>
</template>

<script setup>
import CemeteryLocateTapComponent from '@/components/CemeteryLocateTapComponent.vue'
import { useTimeLineStore } from '@/stores/user';
import axios from '../axios';
import { ref } from 'vue';
import { computed } from 'vue';


const { setTimeLineInfo} = useTimeLineStore()
let staticENV = ref(import.meta.env.VITE_STATIC_URL + '/imgfile/')

setTimeLineInfo('five')


let cemeteries = ref([])
axios.post('/loginafter/getsupplybytype?type=2').then(res => {
    cemeteries.value = res.data.data
})

const cemetryItems = computed(() => {
  const chunkSize = 4;
  const result = [];
  for (let i = 0; i < cemeteries.value.length; i += chunkSize) {
    result.push(cemeteries.value.slice(i, i + chunkSize));
  }
  console.log(result)
  return result;
})
</script>

<style lang="less" scoped>
.outer {
    height: 45vh;
    display: flex;
}
p{
font-size: 1.1rem;
color: rgb(103, 160, 95);
}
</style>