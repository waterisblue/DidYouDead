import { ref, computed } from 'vue'
import { defineStore } from 'pinia'

export const useCounterStore = defineStore('counter', () => {
  const count = ref(0)
  const doubleCount = computed(() => count.value * 2)
  function increment() {
    count.value++
  }

  return { count, doubleCount, increment }
})

export const useTimeLineStore = defineStore('timeline', () => {
  const timeLineInfo = ref("one");

  function setTimeLineInfo(path){
      timeLineInfo.value = path
  }


  return {timeLineInfo, setTimeLineInfo}
})

export const useFireServiceChoose = defineStore('fireservice', () => {
  const useFireServiceChoose = ref({})
  function getService(){
    return useFireServiceChoose
  }
  function setService(service){
    useFireServiceChoose.value = service
  }
  return {useFireServiceChoose, getService, setService}
})