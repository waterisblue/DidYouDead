<template>
    <transition name="bubbleAni" appear>
        <div class="bubble" :id="innerId" @click="bubbleClick(innerId)">
            {{innerText}}
        </div>
    </transition>
</template>

<script setup>
    import { defineProps } from 'vue';
    import { useRouter } from 'vue-router';

    const router = useRouter()
    let name = 'BubbleComponent'
    
    const props = defineProps({
        innerText: String,
        innerId: String
    })

    function bubbleClick(path) {
        let bubble = document.getElementById(props.innerId)
        bubble.style.width = '80rem'
        bubble.style.height = '80rem'
        bubble.style.opacity = '0'
        setTimeout(()=>{router.push(props.innerId)}, 1000)
    }
</script>
<style scoped>
    .bubble {
            background-color: rgba(255, 255, 255, 0.7);
            border-radius: 50%;
            display: flex;
            justify-content: center;
            align-items: center;
            color: #b0b4b8;
            box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
            cursor: pointer;
            animation: floatUp 4s ease-in-out infinite;
            animation-delay: calc(var(--i) * -0.5s);
            text-align: center;

            width: 20rem;
            height: 20rem;
            font-size: 3rem;
            transition: all 2s;
    }

    .bubbleAni-enter-from,
    .bubbleAni-appear-from {
        width: 1rem;
        height: 1rem;
        font-size: 0.1rem;
    }

    /* 过渡结束状态 */
    .bubbleAni-enter-to,
    .bubbleAni-appear-to {
        width: 20rem;
        height: 20rem;
        font-size: 3rem;
    }

    /* 过渡激活状态 */
    .bubbleAni-enter-active,
    .bubbleAni-appear-active,
    .bubbleAni-leave-active {
        transition: all 2s;
    }


    .bubble:hover {
        width: 22rem;
        height: 22rem;
        font-size: 3.5rem;
    }
</style>