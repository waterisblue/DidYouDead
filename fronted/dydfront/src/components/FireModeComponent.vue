<template>
    <div class="outer">
        <!-- <FireModeTapComponent tap-image="/static/jian1.jpg" tap-name="拣灰式火化机" tap-sub-title="价格:4999元(燃油)/4299元(燃气)" tap-detail="拣灰火化机分为燃油和燃气两款。燃油机功率大，时间短，但成本较高；燃气机成本低，耗能少，更环保但用时长。用户可根据尸体的具体大小进行灵活选择，调整燃烧大小和时间，节约成本。但需要专门的装置和操作流程来进行拣灰，增加了操作的复杂性和等待时间；收集的灰渣会有大块残留需要进一步处理或处置，进一步加大火化时长。" />
        <FireModeTapComponent tap-image="/static/lu2.jpg" tap-name="平板式火化机" tap-sub-title="价格：6999元" tap-detail="采用高科技燃烧技术和多重净化技术，燃烧效率高、能源利用率高，同时烟气和废气都通过净化器进行净化，达到国家的排放标准。但要注意，遗体背部与支承平板炕面接触，形成燃烧死角，焚化效率略低。总体用时小于拣灰火化机，是大众首选方式。"/>
        <FireModeTapComponent tap-image="/static/lu.jpg" tap-name="辊道式电热火化炉" tap-sub-title="价格：8999元" tap-detail="优点：采用架空燃烧的方式，增大了遗体燃烧时的表面积，火焰可围绕整个遗体进行燃烧，燃烧死角很小，因而遗体焚化速度快，节约燃料，利于连续焚化，且焚化效果好，噪音小。缺点：可能需要较高的电力供应，且设备成本和维护成本可能较高。"/>
        <FireModeTapComponent tap-image="/static/jian2.jpg" tap-name="生物质火化炉" tap-sub-title="10999元" tap-detail="生物质火化炉能够有效地将生物质燃料转化为热能或电能，提高了能源利用效率，降低了能源消耗成本。生物质燃料来源广泛，且价格相对较为稳定，能够适应不同的生物质燃料类型，具有较高的燃料适应性。但也存在热值较低、技术要求高，普及率低等缺点。"/>
        <FireModeTapComponent tap-image="/static/R5.jpg" tap-name="宠物专用炉" tap-sub-title="599元" tap-detail=" 单独为一只宠物进行火化，结束后您可以得到宠物的全部骨灰。还提供宠物墓地服务，供您选择。"/> -->
        <FireModeTapComponent v-for="item in items" :tap-id="item.id" :tap-image="staticENV + item.imgurl" :tap-name="item.sourcename" :tap-sub-title="item.subtitle" :tap-detail="item.sourcedetail" />
    </div>
</template>

<script setup>
import FireModeTapComponent from '@/components/FireModeTapComponent.vue'
import { useTimeLineStore } from '@/stores/user';
import axios from '../axios';
import { ref } from 'vue';

const { setTimeLineInfo} = useTimeLineStore()
let staticENV = ref(import.meta.env.VITE_STATIC_URL + '/imgfile/')
let items = ref([])

axios.post('/loginafter/getsupplybytype?type=3').then(res => {
    items.value = res.data.data
})

setTimeLineInfo('three')
</script>

<style lang="less" scoped>
.outer {
    height: 75vh;
    display: flex;
    flex-direction: row;
    justify-content: space-around;
}
</style>