<template>
    <p>===集中园区===</p>
    <div class="outer">
        
        <CemeteryLocateTapComponent  tap-image="/static/mu8.jpg" tap-name="傍山谧园" tap-sub-title="免费" tap-detail=""/>
        <CemeteryLocateTapComponent tap-image="/static/mu12.jpg"  tap-name="传统园区" tap-sub-title="免费" tap-detail="" />
        <CemeteryLocateTapComponent tap-image="/static/mu9.jpg" tap-name="绿色伴生园" tap-sub-title="免费" tap-detail=""/>
        <CemeteryLocateTapComponent tap-image="/static/mu13jpg.jpg" tap-name="诗意长眠园" tap-sub-title="免费" tap-detail=""/>
    </div>
    <p>===独立园区===</p>
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
        
    </div>
</template>

<script setup>
import CemeteryLocateTapComponent from '@/components/CemeteryLocateTapComponent.vue'

import { useTimeLineStore } from '@/stores/user';
import { storeToRefs } from 'pinia';

const {timeLineStore, setTimeLineInfo} = useTimeLineStore()

const {timeLineInfo} = storeToRefs(timeLineStore)
setTimeLineInfo('five')
</script>

<style lang="less" scoped>
.outer {
    height: 45vh;
    display: flex;
    
    //  padding: 20px;
     
    // flex-direction: row;
    // justify-content: space-around;
}
p{
font-size: 1.1rem;
color: rgb(103, 160, 95);
}
</style>