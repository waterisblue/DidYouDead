<template>
    <div class="outer">
        <OneTabCardComponent tap-image="/static/onetap.jpg" tap-name="永恒之翼殡仪馆" tap-sub-title="1288 永恒路，和谐市，安宁区，邮编 310015" tap-detail="永恒之翼殡仪馆提供一个庄严而宁静的环境，让您在此向逝去的亲人致以最深的敬意。我们的专业团队将确保每个纪念仪式都是独特而尊重的，帮助您以最贴心的方式来缅怀和告别。" />
        <OneTabCardComponent tap-image="/static/twotap.jpg" tap-name="安祥殿堂" tap-sub-title="772 宁静大道，慰藉镇，温馨岭，邮编 210045" tap-detail="安祥殿堂以其温馨的服务和典雅的环境著称，旨在为您提供一个平和的告别空间。我们的个性化服务确保每位逝者都能得到他们应得的尊重和纪念，让家属在这个艰难时刻感受到安慰和支持。"/>
        <OneTabCardComponent tap-image="/static/threetap.jpg" tap-name="星河仙途" tap-sub-title="555 星辰街，天际村，星光县，邮编 420123" tap-detail="星河仙途殡仪馆以宁静的天体主题为背景，为您的亲人提供一个独特的最后旅程。我们相信，每个生命都如同繁星般璀璨，应当以最尊贵的方式送别。我们的服务旨在帮助家庭以一种庄重而又个性化的方式纪念他们的爱人。"/>
        <OneTabCardComponent tap-image="/static/fourtap.jpg" tap-name="珍爱纪念馆" tap-sub-title="939 记忆巷，回忆市，珍爱区，邮编 520234" tap-detail="珍爱纪念馆以其专注于细节和对家庭的深切同情而自豪。我们提供全方位的殡葬服务，并致力于创造一个温馨的环境，让您在此缅怀逝者，并与家人和朋友共同分享生命中的美好回忆。"/>
        <OneTabCardComponent tap-image="/static/fivetap.jpg" tap-name="诗意长眠园" tap-sub-title="88 诗意路，风景区，自然市，邮编 330789" tap-detail=" 诗意长眠园位于风景如画的郊外，提供一个与自然和谐共存的安息地。我们的理念是将每一次告别变成对生命的庆祝，通过我们美丽的花园和宁静的湖泊，为逝者创造一个永恒的休憩之所。"/>
    </div>
</template>

<script setup>
import OneTabCardComponent from '@/components/OneTapCardComponent.vue'

import { useTimeLineStore } from '@/stores/user';
import { storeToRefs } from 'pinia';

const {timeLineStore, setTimeLineInfo} = useTimeLineStore()

const {timeLineInfo} = storeToRefs(timeLineStore)
setTimeLineInfo('one')
</script>

<style lang="less" scoped>
.outer {
    height: 75vh;
    display: flex;
    flex-direction: row;
    justify-content: space-around;
}
</style>