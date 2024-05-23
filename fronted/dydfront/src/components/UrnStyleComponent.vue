<template>
    <div class="outer">
        <UrnStyleTapComponent tap-image="/static/gu1.jpg" tap-name="青花瓷" tap-sub-title="" tap-detail="" />
        <UrnStyleTapComponent tap-image="/static/gu3.jpg" tap-name="原木狗狗屋"  tap-sub-title="" tap-detail=""/>
        <UrnStyleTapComponent tap-image="/static/guC.jpg" tap-name="星河仙途"  tap-sub-title="" tap-detail=""/>
        <UrnStyleTapComponent tap-image="/static/gu5.jpg" tap-name="熠熠生辉"  tap-sub-title="" tap-detail=""/>
        <UrnStyleTapComponent tap-image="/static/gu6.jpg" tap-name="半壶纱"  tap-sub-title="" tap-detail=""/>
        
    </div>
    <div class="outer2">
        <UrnStyleTapComponent tap-image="/static/gu7.jpg" tap-name="传统汉白玉" tap-sub-title="" tap-detail="" />
        <UrnStyleTapComponent tap-image="/static/gu8.jpg" tap-name="猫猫之家" tap-sub-title="" tap-detail=""/>
        <UrnStyleTapComponent tap-image="/static/gu9.jpg" tap-name="莲芷晴" tap-sub-title="" tap-detail=""/>
        <UrnStyleTapComponent tap-image="/static/gu11.jpg" tap-name="珍爱纪念馆" tap-sub-title="" tap-detail=""/>
        <UrnStyleTapComponent tap-image="/static/gu12.jpg" tap-name="环保诗意园" tap-sub-title="" tap-detail=""/>
        
    </div>
    <div class="outer2">
        <UrnStyleTapComponent tap-image="/static/gu13.jpg" tap-name="精致釉下彩" tap-sub-title="" tap-detail="" />
        <UrnStyleTapComponent tap-image="/static/gu0.jpg" tap-name="太空专用真空杯" tap-sub-title="" tap-detail=""/>
        <UrnStyleTapComponent tap-image="/static/gu15.jpg" tap-name="猫猫不倒翁" tap-sub-title="" tap-detail=""/>
        <UrnStyleTapComponent tap-image="/static/gu16.jpg" tap-name="ins简约风" tap-sub-title="" tap-detail=""/>
        <UrnStyleTapComponent tap-image="/static/gu17.jpg" tap-name="卡通罐罐（可备注小动物）" tap-sub-title="" tap-detail=""/>
        
    </div>
</template>

<script setup>
import UrnStyleTapComponent from '@/components/UrnStyleTapComponent.vue'

import { useTimeLineStore } from '@/stores/user';
import { storeToRefs } from 'pinia';

const {timeLineStore, setTimeLineInfo} = useTimeLineStore()

const {timeLineInfo} = storeToRefs(timeLineStore)
setTimeLineInfo('four')
</script>

<style lang="less" scoped>
.outer {
    height: 45vh;
    display: flex;
    // flex-direction: row;
    // justify-content: space-around;
}
.outer2{
    height: 45vh;
    display: flex;
    // flex-direction: row;
    // justify-content: space-around;
}
</style>