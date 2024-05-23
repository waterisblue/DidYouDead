<template>
    <div class="outer">
        <OneTabCardComponent tap-image="/static/dongjiao.jpg" tap-name="北京市东郊殡仪馆" tap-sub-title="联系电话：010-65422869" tap-detail="特色介绍：北京市东郊殡仪馆始建于1952年，位于北京市朝阳区平房北街133号，占地面积为2000平方米,是一个集遗体接运、遗体冷藏、遗体整容、遗体告别、遗体火化和骨灰寄存等多种殡仪服务功能于一体的综合性殡仪馆。该殡仪馆秉持“民政为民，民政爱民”理念，坚持公益至上、真情为民，主要提供24小时遗体接运服务，以及接待洽谈、遗体守灵、追思祭祀、丧葬用品销售等多种服务项目。" />
        <OneTabCardComponent tap-image="/static/babaos.jpg" tap-name="北京市八宝山殡仪馆" tap-sub-title="联系电话：010-37622847" tap-detail="特色介绍：八宝山殡仪馆是北京市内知名的殡仪服务机构，位于北京市石景山区，提供多种殡仪服务，包括遗体接运、告别仪式、火化等。以其温馨的服务和典雅的环境著称，旨在为您提供一个平和的告别空间。该馆的个性化服务确保每位逝者都能得到他们应得的尊重和纪念，让家属在这个艰难时刻感受到安慰和支持。"/>
        <OneTabCardComponent tap-image="/static/mentou.jpg" tap-name="门头沟区殡仪馆" tap-sub-title="联系电话：010-89345039" tap-detail="特色介绍：门头沟区殡仪馆位于北京市门头沟区167号，以宁静的天体主题为背景，为您的亲人提供一个独特的最后旅程。我们相信，每个生命都如同繁星般璀璨，应当以最尊贵的方式送别。我们的服务旨在帮助家庭以一种庄重而又个性化的方式纪念他们的爱人。"/>
        <OneTabCardComponent tap-image="/static/fangshan.jpg" tap-name="房山区殡仪馆" tap-sub-title="联系电话：010-87645335" tap-detail="特色介绍：房山区殡仪馆位于北京市房山区老汽车站对面，提供全面的殡仪服务，包括遗体接运、告别仪式、火化等，为逝者家属提供便利。该馆以其专注于细节和对家庭的深切同情而自豪，致力于创造一个温馨的环境，让您在此缅怀逝者，并与家人和朋友共同分享生命中的美好回忆。"/>
        <OneTabCardComponent tap-image="/static/tongz4.jpg" tap-name="通州区殡仪馆" tap-sub-title="联系电话：010-69542145" tap-detail=" 特色介绍：通州区殡仪馆位于风景如画的郊外，提供一个与自然和谐共存的安息地。我们的理念是致力于为当地居民提供优质的殡仪服务，满足家属的需求，让逝者安息。将每一次告别变成对生命的庆祝，通过我们美丽的花园和宁静的湖泊，为逝者创造一个永恒的休憩之所。"/>
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